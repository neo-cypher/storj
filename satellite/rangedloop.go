// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package satellite

import (
	"context"
	"errors"
	"net"
	"runtime/pprof"

	"github.com/spacemonkeygo/monkit/v3"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"storj.io/private/debug"
	"storj.io/storj/private/lifecycle"
	"storj.io/storj/satellite/accounting/nodetally"
	"storj.io/storj/satellite/audit"
	"storj.io/storj/satellite/durability"
	"storj.io/storj/satellite/gc/piecetracker"
	"storj.io/storj/satellite/gracefulexit"
	"storj.io/storj/satellite/metabase"
	"storj.io/storj/satellite/metabase/rangedloop"
	"storj.io/storj/satellite/metrics"
	"storj.io/storj/satellite/nodeselection"
	"storj.io/storj/satellite/overlay"
	"storj.io/storj/satellite/repair/checker"
)

// RangedLoop is the satellite ranged loop process.
//
// architecture: Peer
type RangedLoop struct {
	Log *zap.Logger
	DB  DB

	Servers  *lifecycle.Group
	Services *lifecycle.Group

	Audit struct {
		Observer rangedloop.Observer
	}

	Debug struct {
		Listener net.Listener
		Server   *debug.Server
	}

	Metrics struct {
		Observer rangedloop.Observer
	}

	Overlay struct {
		Service *overlay.Service
	}

	Repair struct {
		Observer *checker.Observer
	}

	GracefulExit struct {
		Observer rangedloop.Observer
	}

	Accounting struct {
		NodeTallyObserver *nodetally.Observer
	}

	PieceTracker struct {
		Observer *piecetracker.Observer
	}

	DurabilityReport struct {
		Observer []*durability.Report
	}

	RangedLoop struct {
		Service *rangedloop.Service
	}
}

// NewRangedLoop creates a new satellite ranged loop process.
func NewRangedLoop(log *zap.Logger, db DB, metabaseDB *metabase.DB, config *Config, atomicLogLevel *zap.AtomicLevel) (_ *RangedLoop, err error) {
	peer := &RangedLoop{
		Log: log,
		DB:  db,

		Servers:  lifecycle.NewGroup(log.Named("servers")),
		Services: lifecycle.NewGroup(log.Named("services")),
	}

	{ // setup debug
		var err error
		if config.Debug.Addr != "" {
			peer.Debug.Listener, err = net.Listen("tcp", config.Debug.Addr)
			if err != nil {
				withoutStack := errors.New(err.Error())
				peer.Log.Debug("failed to start debug endpoints", zap.Error(withoutStack))
			}
		}
		debugConfig := config.Debug
		debugConfig.ControlTitle = "RangedLoop"
		peer.Debug.Server = debug.NewServerWithAtomicLevel(log.Named("debug"), peer.Debug.Listener, monkit.Default, debugConfig, atomicLogLevel)
		peer.Servers.Add(lifecycle.Item{
			Name:  "debug",
			Run:   peer.Debug.Server.Run,
			Close: peer.Debug.Server.Close,
		})
	}

	{ // setup audit observer
		peer.Audit.Observer = audit.NewObserver(log.Named("audit"), db.VerifyQueue(), config.Audit)
	}

	{ // setup metrics observer
		peer.Metrics.Observer = metrics.NewObserver()
	}

	{ // setup gracefulexit
		if config.GracefulExit.Enabled && !config.GracefulExit.TimeBased {
			peer.GracefulExit.Observer = gracefulexit.NewObserver(
				peer.Log.Named("gracefulexit:observer"),
				peer.DB.GracefulExit(),
				peer.DB.OverlayCache(),
				metabaseDB,
				config.GracefulExit,
			)
		}
	}

	{ // setup node tally observer
		peer.Accounting.NodeTallyObserver = nodetally.NewObserver(
			log.Named("accounting:nodetally"),
			db.StoragenodeAccounting(),
			metabaseDB)
	}

	{ // setup piece tracker observer
		peer.PieceTracker.Observer = piecetracker.NewObserver(
			log.Named("piecetracker"),
			metabaseDB,
			peer.DB.OverlayCache(),
			config.PieceTracker,
		)
	}

	{ // setup
		classes := map[string]func(node *nodeselection.SelectedNode) string{
			"email": func(node *nodeselection.SelectedNode) string {
				return node.Email
			},
			"wallet": func(node *nodeselection.SelectedNode) string {
				return node.Wallet
			},
			"net": func(node *nodeselection.SelectedNode) string {
				return node.LastNet
			},
			"country": func(node *nodeselection.SelectedNode) string {
				return node.CountryCode.String()
			},
		}
		for class, f := range classes {
			peer.DurabilityReport.Observer = append(peer.DurabilityReport.Observer, durability.NewDurability(db.OverlayCache(), metabaseDB, class, f, config.Metainfo.RS.Total, config.Metainfo.RS.Repair, config.Metainfo.RS.Repair-config.Metainfo.RS.Min, config.RangedLoop.AsOfSystemInterval))
		}
	}

	{ // setup overlay
		placement, err := config.Placement.Parse()
		if err != nil {
			return nil, err
		}

		peer.Overlay.Service, err = overlay.NewService(peer.Log.Named("overlay"), peer.DB.OverlayCache(), peer.DB.NodeEvents(), placement.CreateFilters, config.Console.ExternalAddress, config.Console.SatelliteName, config.Overlay)
		if err != nil {
			return nil, errs.Combine(err, peer.Close())
		}
		peer.Services.Add(lifecycle.Item{
			Name:  "overlay",
			Run:   peer.Overlay.Service.Run,
			Close: peer.Overlay.Service.Close,
		})
	}

	{ // setup repair
		placement, err := config.Placement.Parse()
		if err != nil {
			return nil, err
		}

		if len(config.Checker.RepairExcludedCountryCodes) == 0 {
			config.Checker.RepairExcludedCountryCodes = config.Overlay.RepairExcludedCountryCodes
		}

		peer.Repair.Observer = checker.NewObserver(
			peer.Log.Named("repair:checker"),
			peer.DB.RepairQueue(),
			peer.Overlay.Service,
			placement.CreateFilters,
			config.Checker,
		)
	}

	{ // setup ranged loop
		observers := []rangedloop.Observer{
			rangedloop.NewLiveCountObserver(metabaseDB, config.RangedLoop.SuspiciousProcessedRatio, config.RangedLoop.AsOfSystemInterval),
			peer.Metrics.Observer,
		}

		if config.Audit.UseRangedLoop {
			observers = append(observers, peer.Audit.Observer)
		}

		if config.Tally.UseRangedLoop {
			observers = append(observers, peer.Accounting.NodeTallyObserver)
		}

		if peer.GracefulExit.Observer != nil && config.GracefulExit.UseRangedLoop {
			observers = append(observers, peer.GracefulExit.Observer)
		}

		if config.Repairer.UseRangedLoop {
			observers = append(observers, peer.Repair.Observer)
		}

		if config.PieceTracker.UseRangedLoop {
			observers = append(observers, peer.PieceTracker.Observer)
		}

		if config.DurabilityReport.Enabled {
			for _, observer := range peer.DurabilityReport.Observer {
				observers = append(observers, observer)
			}
		}

		segments := rangedloop.NewMetabaseRangeSplitter(metabaseDB, config.RangedLoop.AsOfSystemInterval, config.RangedLoop.BatchSize)
		peer.RangedLoop.Service = rangedloop.NewService(log.Named("rangedloop"), config.RangedLoop, segments, observers)

		peer.Services.Add(lifecycle.Item{
			Name: "rangeloop",
			Run:  peer.RangedLoop.Service.Run,
		})
	}

	return peer, nil
}

// Run runs satellite ranged loop until it's either closed or it errors.
func (peer *RangedLoop) Run(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	group, ctx := errgroup.WithContext(ctx)

	pprof.Do(ctx, pprof.Labels("subsystem", "rangedloop"), func(ctx context.Context) {
		peer.Servers.Run(ctx, group)
		peer.Services.Run(ctx, group)

		pprof.Do(ctx, pprof.Labels("name", "subsystem-wait"), func(ctx context.Context) {
			err = group.Wait()
		})
	})
	return err
}

// Close closes all the resources.
func (peer *RangedLoop) Close() error {
	return errs.Combine(
		peer.Servers.Close(),
		peer.Services.Close(),
	)
}
