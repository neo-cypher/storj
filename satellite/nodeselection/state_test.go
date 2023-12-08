// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package nodeselection_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"storj.io/common/testcontext"
	"storj.io/common/testrand"
	"storj.io/storj/satellite/nodeselection"
)

func TestState_SelectNonDistinct(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	reputableNodes := joinNodes(
		createRandomNodes(2, "1.0.1", false),
		createRandomNodes(3, "1.0.2", false),
	)
	newNodes := joinNodes(
		createRandomNodes(2, "1.0.3", false),
		createRandomNodes(3, "1.0.4", false),
	)

	state := nodeselection.NewState(reputableNodes, newNodes)

	{ // select 5 non-distinct subnet reputable nodes
		const selectCount = 5
		selected, err := state.Select(ctx, nodeselection.Request{
			Count:       selectCount,
			NewFraction: 0,
		})
		require.NoError(t, err)
		require.Len(t, selected, selectCount)
	}

	{ // select 6 non-distinct subnet reputable and new nodes (50%)
		const selectCount = 6
		const newFraction = 0.5
		selected, err := state.Select(ctx, nodeselection.Request{
			Count:       selectCount,
			NewFraction: newFraction,
		})
		require.NoError(t, err)
		require.Len(t, selected, selectCount)
		require.Len(t, intersectLists(selected, reputableNodes), selectCount*(1-newFraction))
		require.Len(t, intersectLists(selected, newNodes), selectCount*newFraction)
	}

	{ // select 10 distinct subnet reputable and new nodes (100%), falling back to 5 reputable
		const selectCount = 10
		const newFraction = 1.0
		selected, err := state.Select(ctx, nodeselection.Request{
			Count:       selectCount,
			NewFraction: newFraction,
		})
		require.NoError(t, err)
		require.Len(t, selected, selectCount)
		require.Len(t, intersectLists(selected, reputableNodes), 5)
		require.Len(t, intersectLists(selected, newNodes), 5)
	}
}

func TestState_SelectDistinct(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	reputableNodes := joinNodes(
		createRandomNodes(2, "1.0.1", true),
		createRandomNodes(3, "1.0.2", true),
	)
	newNodes := joinNodes(
		createRandomNodes(2, "1.0.3", true),
		createRandomNodes(3, "1.0.4", true),
	)

	state := nodeselection.NewState(reputableNodes, newNodes)

	{ // select 2 distinct subnet reputable nodes
		const selectCount = 2
		selected, err := state.Select(ctx, nodeselection.Request{
			Count:       selectCount,
			NewFraction: 0,
		})
		require.NoError(t, err)
		require.Len(t, selected, selectCount)
	}

	{ // try to select 5 distinct subnet reputable nodes, but there are only two 2 in the state
		const selectCount = 5
		selected, err := state.Select(ctx, nodeselection.Request{
			Count:       selectCount,
			NewFraction: 0,
		})
		require.Error(t, err)
		require.Len(t, selected, 2)
	}

	{ // select 4 distinct subnet reputable and new nodes (50%)
		const selectCount = 4
		const newFraction = 0.5
		selected, err := state.Select(ctx, nodeselection.Request{
			Count:       selectCount,
			NewFraction: newFraction,
		})
		require.NoError(t, err)
		require.Len(t, selected, selectCount)
		require.Len(t, intersectLists(selected, reputableNodes), selectCount*(1-newFraction))
		require.Len(t, intersectLists(selected, newNodes), selectCount*newFraction)
	}
}

func TestState_Select_Concurrent(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	reputableNodes := joinNodes(
		createRandomNodes(2, "1.0.1", false),
		createRandomNodes(3, "1.0.2", false),
	)
	newNodes := joinNodes(
		createRandomNodes(2, "1.0.3", false),
		createRandomNodes(3, "1.0.4", false),
	)

	state := nodeselection.NewState(reputableNodes, newNodes)

	var group errgroup.Group
	group.Go(func() error {
		const selectCount = 5
		nodes, err := state.Select(ctx, nodeselection.Request{
			Count:       selectCount,
			NewFraction: 0.5,
		})
		require.Len(t, nodes, selectCount)
		return err
	})

	group.Go(func() error {
		const selectCount = 4
		nodes, err := state.Select(ctx, nodeselection.Request{
			Count:       selectCount,
			NewFraction: 0.5,
		})
		require.Len(t, nodes, selectCount)
		return err
	})
	require.NoError(t, group.Wait())
}

// createRandomNodes creates n random nodes all in the subnet.
func createRandomNodes(n int, subnet string, shareNets bool) []*nodeselection.SelectedNode {
	xs := make([]*nodeselection.SelectedNode, n)
	for i := range xs {
		addr := subnet + "." + strconv.Itoa(i) + ":8080"
		xs[i] = &nodeselection.SelectedNode{
			ID:         testrand.NodeID(),
			LastNet:    addr,
			LastIPPort: addr,
		}
		if shareNets {
			xs[i].LastNet = subnet
		} else {
			xs[i].LastNet = addr
		}
	}
	return xs
}

// joinNodes appends all slices into a single slice.
func joinNodes(lists ...[]*nodeselection.SelectedNode) []*nodeselection.SelectedNode {
	xs := []*nodeselection.SelectedNode{}
	for _, list := range lists {
		xs = append(xs, list...)
	}
	return xs
}

// intersectLists returns nodes that exist in both lists compared by ID.
func intersectLists(as, bs []*nodeselection.SelectedNode) []*nodeselection.SelectedNode {
	var xs []*nodeselection.SelectedNode

next:
	for _, a := range as {
		for _, b := range bs {
			if a.ID == b.ID {
				xs = append(xs, a)
				continue next
			}
		}
	}

	return xs
}
