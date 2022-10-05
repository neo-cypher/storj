// Copyright (C) 2019 Storj Labs, Incache.
// See LICENSE for copying information.

package overlay

import (
	"context"
	"time"

	"go.uber.org/zap"

	"storj.io/common/storj"
	"storj.io/common/sync2"
)

// DownloadSelectionDB implements the database for download selection cache.
//
// architecture: Database
type DownloadSelectionDB interface {
	// SelectAllStorageNodesDownload returns nodes that are ready for downloading
	SelectAllStorageNodesDownload(ctx context.Context, onlineWindow time.Duration, asOf AsOfSystemTimeConfig) ([]*SelectedNode, error)
}

// DownloadSelectionCacheConfig contains configuration for the selection cache.
type DownloadSelectionCacheConfig struct {
	Staleness      time.Duration
	OnlineWindow   time.Duration
	AsOfSystemTime AsOfSystemTimeConfig
}

// DownloadSelectionCache keeps a list of all the storage nodes that are qualified to download data from.
// The cache will sync with the nodes table in the database and get refreshed once the staleness time has past.
type DownloadSelectionCache struct {
	log    *zap.Logger
	db     DownloadSelectionDB
	config DownloadSelectionCacheConfig

	cache sync2.ReadCache
}

// NewDownloadSelectionCache creates a new cache that keeps a list of all the storage nodes that are qualified to download data from.
func NewDownloadSelectionCache(log *zap.Logger, db DownloadSelectionDB, config DownloadSelectionCacheConfig) (*DownloadSelectionCache, error) {
	cache := &DownloadSelectionCache{
		log:    log,
		db:     db,
		config: config,
	}
	return cache, cache.cache.Init(config.Staleness/2, config.Staleness, cache.read)
}

// Run runs the background task for cache.
func (cache *DownloadSelectionCache) Run(ctx context.Context) (err error) {
	return cache.cache.Run(ctx)
}

// Refresh populates the cache with all of the reputableNodes and newNode nodes
// This method is useful for tests.
func (cache *DownloadSelectionCache) Refresh(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)
	_, err = cache.cache.RefreshAndGet(ctx, time.Now())
	return err
}

// read loads the latest download selection state.
func (cache *DownloadSelectionCache) read(ctx context.Context) (_ interface{}, err error) {
	defer mon.Task()(&ctx)(&err)

	onlineNodes, err := cache.db.SelectAllStorageNodesDownload(ctx, cache.config.OnlineWindow, cache.config.AsOfSystemTime)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	mon.IntVal("refresh_cache_size_online").Observe(int64(len(onlineNodes)))

	return NewDownloadSelectionCacheState(onlineNodes), nil
}

// GetNodeIPs gets the last node ip:port from the cache, refreshing when needed.
func (cache *DownloadSelectionCache) GetNodeIPs(ctx context.Context, nodes []storj.NodeID) (_ map[storj.NodeID]string, err error) {
	defer mon.Task()(&ctx)(&err)

	stateAny, err := cache.cache.Get(ctx, time.Now())
	if err != nil {
		return nil, Error.Wrap(err)
	}
	state := stateAny.(*DownloadSelectionCacheState)

	return state.IPs(nodes), nil
}

// GetNodes gets nodes by ID from the cache, and refreshes the cache if it is stale.
func (cache *DownloadSelectionCache) GetNodes(ctx context.Context, nodes []storj.NodeID) (_ map[storj.NodeID]*SelectedNode, err error) {
	defer mon.Task()(&ctx)(&err)

	stateAny, err := cache.cache.Get(ctx, time.Now())
	if err != nil {
		return nil, Error.Wrap(err)
	}
	state := stateAny.(*DownloadSelectionCacheState)

	return state.Nodes(nodes), nil
}

// Size returns how many nodes are in the cache.
func (cache *DownloadSelectionCache) Size(ctx context.Context) (int, error) {
	stateAny, err := cache.cache.Get(ctx, time.Now())
	if stateAny == nil || err != nil {
		return 0, Error.Wrap(err)
	}
	state := stateAny.(*DownloadSelectionCacheState)
	return state.Size(), nil
}

// DownloadSelectionCacheState contains state of download selection cache.
type DownloadSelectionCacheState struct {
	// byID returns IP based on storj.NodeID
	byID map[storj.NodeID]*SelectedNode // TODO: optimize, avoid pointery structures for performance
}

// NewDownloadSelectionCacheState creates a new state from the nodes.
func NewDownloadSelectionCacheState(nodes []*SelectedNode) *DownloadSelectionCacheState {
	byID := map[storj.NodeID]*SelectedNode{}
	for _, n := range nodes {
		byID[n.ID] = n
	}
	return &DownloadSelectionCacheState{
		byID: byID,
	}
}

// Size returns how many nodes are in the state.
func (state *DownloadSelectionCacheState) Size() int {
	return len(state.byID)
}

// IPs returns node ip:port for nodes that are in state.
func (state *DownloadSelectionCacheState) IPs(nodes []storj.NodeID) map[storj.NodeID]string {
	xs := make(map[storj.NodeID]string, len(nodes))
	for _, nodeID := range nodes {
		if n, exists := state.byID[nodeID]; exists {
			xs[nodeID] = n.LastIPPort
		}
	}
	return xs
}

// Nodes returns node ip:port for nodes that are in state.
func (state *DownloadSelectionCacheState) Nodes(nodes []storj.NodeID) map[storj.NodeID]*SelectedNode {
	xs := make(map[storj.NodeID]*SelectedNode, len(nodes))
	for _, nodeID := range nodes {
		if n, exists := state.byID[nodeID]; exists {
			xs[nodeID] = n.Clone() // TODO: optimize the clones
		}
	}
	return xs
}
