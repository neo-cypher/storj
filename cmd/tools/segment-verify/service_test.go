// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package main_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeebo/errs"

	"storj.io/common/pb"
	"storj.io/common/storj"
	"storj.io/common/testcontext"
	"storj.io/common/uuid"
	segmentverify "storj.io/storj/cmd/tools/segment-verify"
	"storj.io/storj/private/testplanet"
	"storj.io/storj/satellite/metabase"
	"storj.io/storj/satellite/overlay"
)

var maxUUID = uuid.UUID{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

func TestService_EmptyRange(t *testing.T) {
	ctx := testcontext.New(t)
	log := testplanet.NewLogger(t)

	config := segmentverify.ServiceConfig{
		NotFoundPath: ctx.File("not-found.csv"),
		RetryPath:    ctx.File("retry.csv"),
	}

	metabase := newMetabaseMock(map[metabase.NodeAlias]storj.NodeID{})
	verifier := &verifierMock{allSuccess: true}

	service, err := segmentverify.NewService(log.Named("segment-verify"), metabase, verifier, metabase, config)
	require.NoError(t, err)
	require.NotNil(t, service)

	defer ctx.Check(service.Close)

	err = service.ProcessRange(ctx, uuid.UUID{}, maxUUID)
	require.NoError(t, err)
}

func TestService_Success(t *testing.T) {
	ctx := testcontext.New(t)
	log := testplanet.NewLogger(t)

	config := segmentverify.ServiceConfig{
		NotFoundPath:      ctx.File("not-found.csv"),
		RetryPath:         ctx.File("retry.csv"),
		PriorityNodesPath: ctx.File("priority-nodes.txt"),

		Check:       3,
		BatchSize:   100,
		Concurrency: 3,
	}

	// the node 1 is going to be priority
	err := ioutil.WriteFile(config.PriorityNodesPath, []byte((storj.NodeID{1}).String()+"\n"), 0755)
	require.NoError(t, err)

	func() {
		nodes := map[metabase.NodeAlias]storj.NodeID{}
		for i := 1; i <= 0xFF; i++ {
			nodes[metabase.NodeAlias(i)] = storj.NodeID{byte(i)}
		}

		segments := []metabase.VerifySegment{
			{
				StreamID:    uuid.UUID{0x10, 0x10},
				AliasPieces: metabase.AliasPieces{{Number: 1, Alias: 8}, {Number: 3, Alias: 9}, {Number: 5, Alias: 10}, {Number: 0, Alias: 1}},
			},
			{
				StreamID:    uuid.UUID{0x20, 0x20},
				AliasPieces: metabase.AliasPieces{{Number: 0, Alias: 2}, {Number: 1, Alias: 3}, {Number: 7, Alias: 4}},
			},
			{ // this won't get processed due to the high limit
				StreamID:    uuid.UUID{0x30, 0x30},
				AliasPieces: metabase.AliasPieces{{Number: 0, Alias: 2}, {Number: 1, Alias: 3}, {Number: 7, Alias: 4}},
			},
		}

		metabase := newMetabaseMock(nodes, segments...)
		verifier := &verifierMock{allSuccess: true}

		service, err := segmentverify.NewService(log.Named("segment-verify"), metabase, verifier, metabase, config)
		require.NoError(t, err)
		require.NotNil(t, service)

		defer ctx.Check(service.Close)

		err = service.ProcessRange(ctx, uuid.UUID{0x10, 0x10}, uuid.UUID{0x30, 0x30})
		require.NoError(t, err)

		for node, list := range verifier.processed {
			assert.True(t, isUnique(list), "each node should process only once: %v %#v", node, list)
		}

		// node 1 is a priority node in the segments[0]
		assert.Len(t, verifier.processed[nodes[1]], 1)
		// we should get two other checks against the nodes in segments[8-10]
		assert.Equal(t, 2,
			len(verifier.processed[nodes[8]])+len(verifier.processed[nodes[9]])+len(verifier.processed[nodes[10]]),
		)
		// these correspond to checks against segment[1]
		assert.Len(t, verifier.processed[nodes[2]], 1)
		assert.Len(t, verifier.processed[nodes[3]], 1)
		assert.Len(t, verifier.processed[nodes[4]], 1)
	}()

	retryCSV, err := ioutil.ReadFile(config.RetryPath)
	require.NoError(t, err)
	require.Equal(t, "stream id,position,found,not found,retry\n", string(retryCSV))

	notFoundCSV, err := ioutil.ReadFile(config.NotFoundPath)
	require.NoError(t, err)
	require.Equal(t, "stream id,position,found,not found,retry\n", string(notFoundCSV))
}

func TestService_Failures(t *testing.T) {
	ctx := testcontext.New(t)
	log := testplanet.NewLogger(t)

	config := segmentverify.ServiceConfig{
		NotFoundPath:      ctx.File("not-found.csv"),
		RetryPath:         ctx.File("retry.csv"),
		PriorityNodesPath: ctx.File("priority-nodes.txt"),

		Check:       2,
		BatchSize:   100,
		Concurrency: 3,
	}

	// the node 1 is going to be priority
	err := ioutil.WriteFile(config.PriorityNodesPath, []byte((storj.NodeID{1}).String()+"\n"), 0755)
	require.NoError(t, err)

	func() {
		nodes := map[metabase.NodeAlias]storj.NodeID{}
		for i := 1; i <= 0xFF; i++ {
			nodes[metabase.NodeAlias(i)] = storj.NodeID{byte(i)}
		}

		segments := []metabase.VerifySegment{
			{
				StreamID:    uuid.UUID{0x10, 0x10},
				AliasPieces: metabase.AliasPieces{{Number: 1, Alias: 8}, {Number: 3, Alias: 9}, {Number: 5, Alias: 10}, {Number: 0, Alias: 1}},
			},
			{
				StreamID:    uuid.UUID{0x20, 0x20},
				AliasPieces: metabase.AliasPieces{{Number: 0, Alias: 2}, {Number: 1, Alias: 3}, {Number: 7, Alias: 4}},
			},
			{
				StreamID:    uuid.UUID{0x30, 0x30},
				AliasPieces: metabase.AliasPieces{{Number: 0, Alias: 2}, {Number: 1, Alias: 3}, {Number: 7, Alias: 4}},
			},
		}

		metabase := newMetabaseMock(nodes, segments...)
		verifier := &verifierMock{
			offline:  []storj.NodeID{{0x02}, {0x08}, {0x09}, {0x0A}},
			success:  []uuid.UUID{segments[0].StreamID, segments[2].StreamID},
			notFound: []uuid.UUID{segments[1].StreamID},
		}

		service, err := segmentverify.NewService(log.Named("segment-verify"), metabase, verifier, metabase, config)
		require.NoError(t, err)
		require.NotNil(t, service)

		defer ctx.Check(service.Close)

		err = service.ProcessRange(ctx, uuid.UUID{}, maxUUID)
		require.NoError(t, err)

		for node, list := range verifier.processed {
			assert.True(t, isUnique(list), "each node should process only once: %v %#v", node, list)
		}
	}()

	retryCSV, err := ioutil.ReadFile(config.RetryPath)
	require.NoError(t, err)
	require.Equal(t, ""+
		"stream id,position,found,not found,retry\n"+
		"10100000-0000-0000-0000-000000000000,0,1,0,1\n",
		string(retryCSV))

	notFoundCSV, err := ioutil.ReadFile(config.NotFoundPath)
	require.NoError(t, err)
	require.Equal(t, ""+
		"stream id,position,found,not found,retry\n"+
		"20200000-0000-0000-0000-000000000000,0,0,2,0\n",
		string(notFoundCSV))
}

func isUnique(segments []*segmentverify.Segment) bool {
	type segmentID struct {
		StreamID uuid.UUID
		Position metabase.SegmentPosition
	}
	seen := map[segmentID]bool{}
	for _, seg := range segments {
		id := segmentID{StreamID: seg.StreamID, Position: seg.Position}
		if seen[id] {
			return false
		}
		seen[id] = true
	}
	return true
}

type metabaseMock struct {
	nodeIDToAlias map[storj.NodeID]metabase.NodeAlias
	aliasToNodeID map[metabase.NodeAlias]storj.NodeID
	segments      []metabase.VerifySegment
}

func newMetabaseMock(nodes map[metabase.NodeAlias]storj.NodeID, segments ...metabase.VerifySegment) *metabaseMock {
	mock := &metabaseMock{
		nodeIDToAlias: map[storj.NodeID]metabase.NodeAlias{},
		aliasToNodeID: nodes,
		segments:      segments,
	}
	for n, id := range nodes {
		mock.nodeIDToAlias[id] = n
	}
	return mock
}

func (db *metabaseMock) Get(ctx context.Context, nodeID storj.NodeID) (*overlay.NodeDossier, error) {
	return &overlay.NodeDossier{
		Node: pb.Node{
			Id: nodeID,
			Address: &pb.NodeAddress{
				Transport: pb.NodeTransport_TCP_TLS_GRPC,
				Address:   fmt.Sprintf("nodeid:%v", nodeID),
			},
		},
	}, nil
}

func (db *metabaseMock) SelectAllStorageNodesDownload(ctx context.Context, onlineWindow time.Duration, asOf overlay.AsOfSystemTimeConfig) ([]*overlay.SelectedNode, error) {
	var xs []*overlay.SelectedNode
	for nodeID := range db.nodeIDToAlias {
		xs = append(xs, &overlay.SelectedNode{
			ID: nodeID,
			Address: &pb.NodeAddress{
				Transport: pb.NodeTransport_TCP_TLS_GRPC,
				Address:   fmt.Sprintf("nodeid:%v", nodeID),
			},
			LastNet:     "nodeid",
			LastIPPort:  fmt.Sprintf("nodeid:%v", nodeID),
			CountryCode: 0,
		})
	}
	return xs, nil
}

func (db *metabaseMock) ConvertNodesToAliases(ctx context.Context, nodeIDs []storj.NodeID) ([]metabase.NodeAlias, error) {
	xs := make([]metabase.NodeAlias, len(nodeIDs))
	for i, id := range nodeIDs {
		alias, ok := db.nodeIDToAlias[id]
		if !ok {
			return nil, errs.New("id %v not found", id)
		}
		xs[i] = alias
	}
	return xs, nil
}

func (db *metabaseMock) ConvertAliasesToNodes(ctx context.Context, aliases []metabase.NodeAlias) ([]storj.NodeID, error) {
	xs := make([]storj.NodeID, len(aliases))
	for i, alias := range aliases {
		id, ok := db.aliasToNodeID[alias]
		if !ok {
			return nil, errs.New("alias %v not found", alias)
		}
		xs[i] = id
	}
	return xs, nil
}

func (db *metabaseMock) DeleteSegmentByPosition(ctx context.Context, opts metabase.GetSegmentByPosition) error {
	for i, s := range db.segments {
		if opts.StreamID == s.StreamID && opts.Position == s.Position {
			db.segments = append(db.segments[:i], db.segments[i+1:]...)
			return nil
		}
	}
	return metabase.ErrSegmentNotFound.New("%v", opts)
}

func (db *metabaseMock) GetSegmentByPosition(ctx context.Context, opts metabase.GetSegmentByPosition) (segment metabase.Segment, err error) {
	for _, s := range db.segments {
		if opts.StreamID == s.StreamID && opts.Position == s.Position {
			var pieces metabase.Pieces
			for _, p := range s.AliasPieces {
				pieces = append(pieces, metabase.Piece{
					Number:      p.Number,
					StorageNode: db.aliasToNodeID[p.Alias],
				})
			}

			return metabase.Segment{
				StreamID: s.StreamID,
				Position: s.Position,
				Pieces:   pieces,
			}, nil
		}
	}

	return metabase.Segment{}, metabase.ErrSegmentNotFound.New("%v", opts)
}

func (db *metabaseMock) ListVerifySegments(ctx context.Context, opts metabase.ListVerifySegments) (result metabase.ListVerifySegmentsResult, err error) {
	r := metabase.ListVerifySegmentsResult{}

	for _, s := range db.segments {
		if s.StreamID.Less(opts.CursorStreamID) {
			continue
		}
		if s.StreamID == opts.CursorStreamID && !opts.CursorPosition.Less(s.Position) {
			continue
		}

		r.Segments = append(r.Segments, s)
		if len(r.Segments) >= opts.Limit {
			break
		}
	}

	return r, nil
}

type verifierMock struct {
	allSuccess bool
	fail       error
	offline    []storj.NodeID
	success    []uuid.UUID
	notFound   []uuid.UUID

	mu        sync.Mutex
	processed map[storj.NodeID][]*segmentverify.Segment
}

func (v *verifierMock) Verify(ctx context.Context, target storj.NodeURL, segments []*segmentverify.Segment, _ bool) error {
	v.mu.Lock()
	if v.processed == nil {
		v.processed = map[storj.NodeID][]*segmentverify.Segment{}
	}
	v.processed[target.ID] = append(v.processed[target.ID], segments...)
	v.mu.Unlock()

	for _, n := range v.offline {
		if n == target.ID {
			return segmentverify.ErrNodeOffline.New("node did not respond %v", target)
		}
	}
	if v.fail != nil {
		return errs.Wrap(v.fail)
	}

	if v.allSuccess {
		for _, seg := range segments {
			seg.Status.MarkFound()
		}
		return nil
	}

	for _, seg := range v.success {
		for _, t := range segments {
			if t.StreamID == seg {
				t.Status.MarkFound()
			}
		}
	}
	for _, seg := range v.notFound {
		for _, t := range segments {
			if t.StreamID == seg {
				t.Status.MarkNotFound()
			}
		}
	}

	return nil
}
