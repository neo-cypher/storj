// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package gracefulexit_test

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"golang.org/x/sync/errgroup"

	"storj.io/common/errs2"
	"storj.io/common/identity"
	"storj.io/common/memory"
	"storj.io/common/pb"
	"storj.io/common/rpc/rpcstatus"
	"storj.io/common/signing"
	"storj.io/common/storj"
	"storj.io/common/sync2"
	"storj.io/common/testcontext"
	"storj.io/common/testrand"
	"storj.io/storj/private/testplanet"
	"storj.io/storj/satellite"
	"storj.io/storj/satellite/metabase"
	"storj.io/storj/satellite/metainfo"
	"storj.io/storj/satellite/overlay"
	"storj.io/storj/satellite/reputation"
	"storj.io/storj/storagenode"
	"storj.io/storj/storagenode/blobstore/testblobs"
	"storj.io/storj/storagenode/gracefulexit"
	"storj.io/storj/storagenode/satellites"
	"storj.io/uplink/private/piecestore"
)

const numObjects = 6
const numMultipartObjects = 6

// exitProcessClient is used so we can pass the graceful exit process clients regardless of implementation.
type exitProcessClient interface {
	Send(*pb.StorageNodeMessage) error
	Recv() (*pb.SatelliteMessage, error)
}

func TestSuccessOld(t *testing.T) {
	testTransfers(t, numObjects, numMultipartObjects, func(t *testing.T, ctx *testcontext.Context, nodeFullIDs map[storj.NodeID]*identity.FullIdentity, satellite *testplanet.Satellite, processClient exitProcessClient, exitingNode *storagenode.Peer, numPieces int) {
		var pieceID storj.PieceID
		failedCount := 0
		deletedCount := 0
		for {
			response, err := processClient.Recv()
			if errs.Is(err, io.EOF) {
				// Done
				break
			}
			require.NoError(t, err)

			switch m := response.GetMessage().(type) {
			case *pb.SatelliteMessage_TransferPiece:
				require.NotNil(t, m)

				// pick the first one to fail
				if pieceID.IsZero() {
					pieceID = m.TransferPiece.OriginalPieceId
				}

				if failedCount > 0 || pieceID != m.TransferPiece.OriginalPieceId {

					pieceReader, err := exitingNode.Storage2.Store.Reader(ctx, satellite.ID(), m.TransferPiece.OriginalPieceId)
					require.NoError(t, err)

					header, err := pieceReader.GetPieceHeader()
					require.NoError(t, err)

					orderLimit := header.OrderLimit
					originalPieceHash := &pb.PieceHash{
						PieceId:       orderLimit.PieceId,
						Hash:          header.GetHash(),
						PieceSize:     pieceReader.Size(),
						Timestamp:     header.GetCreationTime(),
						Signature:     header.GetSignature(),
						HashAlgorithm: header.GetHashAlgorithm(),
					}

					newPieceHash := &pb.PieceHash{
						PieceId:       m.TransferPiece.AddressedOrderLimit.Limit.PieceId,
						Hash:          originalPieceHash.Hash,
						PieceSize:     originalPieceHash.PieceSize,
						HashAlgorithm: originalPieceHash.HashAlgorithm,
						Timestamp:     time.Now(),
					}

					receivingNodeID := nodeFullIDs[m.TransferPiece.AddressedOrderLimit.Limit.StorageNodeId]
					require.NotNil(t, receivingNodeID)
					signer := signing.SignerFromFullIdentity(receivingNodeID)

					signedNewPieceHash, err := signing.SignPieceHash(ctx, signer, newPieceHash)
					require.NoError(t, err)

					success := &pb.StorageNodeMessage{
						Message: &pb.StorageNodeMessage_Succeeded{
							Succeeded: &pb.TransferSucceeded{
								OriginalPieceId:      m.TransferPiece.OriginalPieceId,
								OriginalPieceHash:    originalPieceHash,
								OriginalOrderLimit:   &orderLimit,
								ReplacementPieceHash: signedNewPieceHash,
							},
						},
					}
					err = processClient.Send(success)
					require.NoError(t, err)
				} else {
					failedCount++
					failed := &pb.StorageNodeMessage{
						Message: &pb.StorageNodeMessage_Failed{
							Failed: &pb.TransferFailed{
								OriginalPieceId: m.TransferPiece.OriginalPieceId,
								Error:           pb.TransferFailed_UNKNOWN,
							},
						},
					}
					err = processClient.Send(failed)
					require.NoError(t, err)
				}
			case *pb.SatelliteMessage_DeletePiece:
				deletedCount++
			case *pb.SatelliteMessage_ExitCompleted:
				signee := signing.SigneeFromPeerIdentity(satellite.Identity.PeerIdentity())
				err = signing.VerifyExitCompleted(ctx, signee, m.ExitCompleted)
				require.NoError(t, err)
			default:
				require.FailNow(t, "should not reach this case: %#v", m)
			}
		}

		// check that the exit has completed and we have the correct transferred/failed values
		progress, err := satellite.DB.GracefulExit().GetProgress(ctx, exitingNode.ID())
		require.NoError(t, err)

		require.EqualValues(t, numPieces, progress.PiecesTransferred)
		require.EqualValues(t, numPieces, deletedCount)
		// even though we failed 1, it eventually succeeded, so the count should be 0
		require.EqualValues(t, 0, progress.PiecesFailed)
	})
}

func TestSuccess(t *testing.T) {
	const steps = 5
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 4,
		Reconfigure: testplanet.Reconfigure{
			Satellite: func(log *zap.Logger, index int, config *satellite.Config) {
				config.GracefulExit.TimeBased = true
			},
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]

		// check that there are no exiting nodes.
		exitingNodes, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 0)

		exitingNode := planet.StorageNodes[0]

		simTime := time.Now()
		satellite.GracefulExit.Endpoint.SetNowFunc(func() time.Time { return simTime })
		doneTime := simTime.AddDate(0, 0, satellite.Config.GracefulExit.GracefulExitDurationInDays)
		interval := doneTime.Sub(simTime) / steps

		// we should get NotReady responses until after the GE time has elapsed.
		for simTime.Before(doneTime) {
			response, err := callProcess(ctx, exitingNode, satellite)
			require.NoError(t, err)
			require.IsType(t, (*pb.SatelliteMessage_NotReady)(nil), response.GetMessage())

			// check that the exiting node is still currently exiting.
			exitingNodes, err = satellite.DB.OverlayCache().GetExitingNodes(ctx)
			require.NoError(t, err)
			require.Len(t, exitingNodes, 1)
			require.Equal(t, exitingNode.ID(), exitingNodes[0].NodeID)

			simTime = simTime.Add(interval)
		}
		simTime = doneTime.Add(time.Second)

		// now we should get a successful finish message
		response, err := callProcess(ctx, exitingNode, satellite)
		require.NoError(t, err)
		require.IsType(t, (*pb.SatelliteMessage_ExitCompleted)(nil), response.GetMessage())

		// verify signature on exit receipt and we're done
		m := response.GetMessage().(*pb.SatelliteMessage_ExitCompleted)
		signee := signing.SigneeFromPeerIdentity(satellite.Identity.PeerIdentity())
		err = signing.VerifyExitCompleted(ctx, signee, m.ExitCompleted)
		require.NoError(t, err)
	})
}

func TestConcurrentConnections(t *testing.T) {
	successThreshold := 4
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: successThreshold + 1,
		UplinkCount:      1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: testplanet.Combine(
				testplanet.ReconfigureRS(2, 3, successThreshold, successThreshold),
				func(log *zap.Logger, index int, config *satellite.Config) {
					// This test can be removed entirely when we are using time-based GE everywhere.
					config.GracefulExit.TimeBased = false
				},
			),
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		uplinkPeer := planet.Uplinks[0]
		satellite := planet.Satellites[0]

		err := uplinkPeer.Upload(ctx, satellite, "testbucket", "test/path1", testrand.Bytes(5*memory.KiB))
		require.NoError(t, err)

		// check that there are no exiting nodes.
		exitingNodeIDs, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodeIDs, 0)

		exitingNode, err := findNodeToExit(ctx, planet, 2)
		require.NoError(t, err)

		var group errgroup.Group
		concurrentCalls := 4

		var mainStarted sync2.Fence
		defer mainStarted.Release()

		for i := 0; i < concurrentCalls; i++ {
			group.Go(func() (err error) {
				// connect to satellite so we initiate the exit.
				conn, err := exitingNode.Dialer.DialNodeURL(ctx, satellite.NodeURL())
				require.NoError(t, err)
				defer func() {
					err = errs.Combine(err, conn.Close())
				}()

				client := pb.NewDRPCSatelliteGracefulExitClient(conn)

				if !mainStarted.Wait(ctx) {
					return ctx.Err()
				}

				c, err := client.Process(ctx)
				require.NoError(t, err)
				defer func() {
					err = errs.Combine(err, c.Close())
				}()

				_, err = c.Recv()
				require.Error(t, err)
				require.True(t, errs2.IsRPC(err, rpcstatus.Aborted))
				return nil
			})
		}

		// connect to satellite so we initiate the exit ("main" call)
		conn, err := exitingNode.Dialer.DialNodeURL(ctx, satellite.NodeURL())
		require.NoError(t, err)
		defer ctx.Check(conn.Close)

		client := pb.NewDRPCSatelliteGracefulExitClient(conn)

		{ // this connection will immediately return since graceful exit has not been initiated yet
			c, err := client.Process(ctx)
			require.NoError(t, err)
			response, err := c.Recv()
			require.NoError(t, err)
			switch m := response.GetMessage().(type) {
			case *pb.SatelliteMessage_NotReady:
			default:
				require.FailNow(t, "should not reach this case: %#v", m)
			}
			require.NoError(t, c.Close())
		}

		// run the satellite ranged loop to build the transfer queue.
		_, err = satellite.RangedLoop.RangedLoop.Service.RunOnce(ctx)
		require.NoError(t, err)

		{ // this connection should not close immediately, since there are pieces to transfer
			c, err := client.Process(ctx)
			require.NoError(t, err)

			_, err = c.Recv()
			require.NoError(t, err)

			// deferring here to ensure that the other connections see the in-use connection.
			defer ctx.Check(c.Close)
		}
		// start receiving from concurrent connections
		mainStarted.Release()

		err = group.Wait()
		require.NoError(t, err)
	})
}

func TestRecvTimeout(t *testing.T) {
	successThreshold := 4
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: successThreshold + 1,
		UplinkCount:      1,
		Reconfigure: testplanet.Reconfigure{
			StorageNodeDB: func(index int, db storagenode.DB, log *zap.Logger) (storagenode.DB, error) {
				return testblobs.NewSlowDB(log.Named("slowdb"), db), nil
			},
			Satellite: testplanet.Combine(
				func(log *zap.Logger, index int, config *satellite.Config) {
					// This config value will create a very short timeframe allowed for receiving
					// data from storage nodes. This will cause context to cancel with timeout.
					config.GracefulExit.RecvTimeout = 10 * time.Millisecond
					// This test can be removed entirely when we are using time-based GE everywhere.
					config.GracefulExit.TimeBased = false
				},
				testplanet.ReconfigureRS(2, 3, successThreshold, successThreshold),
			),
			StorageNode: func(index int, config *storagenode.Config) {
				config.GracefulExit = gracefulexit.Config{
					ChoreInterval:          2 * time.Minute,
					NumWorkers:             2,
					NumConcurrentTransfers: 2,
					MinBytesPerSecond:      128,
					MinDownloadTimeout:     2 * time.Minute,
				}
			},
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]
		ul := planet.Uplinks[0]

		err := ul.Upload(ctx, satellite, "testbucket", "test/path1", testrand.Bytes(5*memory.KiB))
		require.NoError(t, err)

		exitingNode, err := findNodeToExit(ctx, planet, 1)
		require.NoError(t, err)
		exitingNode.GracefulExit.Chore.Loop.Pause()

		exitStatusReq := overlay.ExitStatusRequest{
			NodeID:          exitingNode.ID(),
			ExitInitiatedAt: time.Now(),
		}
		_, err = satellite.Overlay.DB.UpdateExitStatus(ctx, &exitStatusReq)
		require.NoError(t, err)

		// run the satellite ranged loop to build the transfer queue.
		_, err = satellite.RangedLoop.RangedLoop.Service.RunOnce(ctx)
		require.NoError(t, err)

		// check that the satellite knows the storage node is exiting.
		exitingNodes, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 1)
		require.Equal(t, exitingNode.ID(), exitingNodes[0].NodeID)

		queueItems, err := satellite.DB.GracefulExit().GetIncomplete(ctx, exitingNode.ID(), 10, 0)
		require.NoError(t, err)
		require.Len(t, queueItems, 1)

		storageNodeDB := exitingNode.DB.(*testblobs.SlowDB)
		// make uploads on storage node slower than the timeout for transferring bytes to another node
		delay := 200 * time.Millisecond
		storageNodeDB.SetLatency(delay)

		// run the SN chore again to start processing transfers.
		worker := gracefulexit.NewWorker(zaptest.NewLogger(t), exitingNode.GracefulExit.Service, exitingNode.PieceTransfer.Service, exitingNode.Dialer, satellite.NodeURL(), exitingNode.Config.GracefulExit)
		err = worker.Run(ctx)
		require.Error(t, err)
		require.True(t, errs2.IsRPC(err, rpcstatus.DeadlineExceeded))
	})
}

func TestInvalidStorageNodeSignature(t *testing.T) {
	testTransfers(t, 1, 0, func(t *testing.T, ctx *testcontext.Context, nodeFullIDs map[storj.NodeID]*identity.FullIdentity, satellite *testplanet.Satellite, processClient exitProcessClient, exitingNode *storagenode.Peer, numPieces int) {
		response, err := processClient.Recv()
		require.NoError(t, err)

		switch m := response.GetMessage().(type) {
		case *pb.SatelliteMessage_TransferPiece:
			require.NotNil(t, m)
			pieceReader, err := exitingNode.Storage2.Store.Reader(ctx, satellite.ID(), m.TransferPiece.OriginalPieceId)
			require.NoError(t, err)

			header, err := pieceReader.GetPieceHeader()
			require.NoError(t, err)

			orderLimit := header.OrderLimit

			originalPieceHash := &pb.PieceHash{
				PieceId:   orderLimit.PieceId,
				Hash:      header.GetHash(),
				PieceSize: pieceReader.Size(),
				Timestamp: header.GetCreationTime(),
				Signature: header.GetSignature(),
			}

			newPieceHash := &pb.PieceHash{
				PieceId:   m.TransferPiece.AddressedOrderLimit.Limit.PieceId,
				Hash:      originalPieceHash.Hash,
				PieceSize: originalPieceHash.PieceSize,
				Timestamp: time.Now(),
			}

			wrongSigner := signing.SignerFromFullIdentity(exitingNode.Identity)

			signedNewPieceHash, err := signing.SignPieceHash(ctx, wrongSigner, newPieceHash)
			require.NoError(t, err)

			message := &pb.StorageNodeMessage{
				Message: &pb.StorageNodeMessage_Succeeded{
					Succeeded: &pb.TransferSucceeded{
						OriginalPieceId:      m.TransferPiece.OriginalPieceId,
						OriginalPieceHash:    originalPieceHash,
						OriginalOrderLimit:   &orderLimit,
						ReplacementPieceHash: signedNewPieceHash,
					},
				},
			}
			err = processClient.Send(message)
			require.NoError(t, err)
		default:
			require.FailNow(t, "should not reach this case: %#v", m)
		}

		response, err = processClient.Recv()
		require.NoError(t, err)

		switch m := response.GetMessage().(type) {
		case *pb.SatelliteMessage_ExitFailed:
			require.NotNil(t, m)
			require.NotNil(t, m.ExitFailed)
			require.Equal(t, m.ExitFailed.Reason, pb.ExitFailed_VERIFICATION_FAILED)

			node, err := satellite.DB.OverlayCache().Get(ctx, m.ExitFailed.NodeId)
			require.NoError(t, err)
			require.NotNil(t, node.Disqualified)
		default:
			require.FailNow(t, "should not reach this case: %#v", m)
		}

		// check that the exit has completed and we have the correct transferred/failed values
		progress, err := satellite.DB.GracefulExit().GetProgress(ctx, exitingNode.ID())
		require.NoError(t, err)

		require.Equal(t, int64(0), progress.PiecesTransferred)
		require.Equal(t, int64(1), progress.PiecesFailed)
	})
}

func TestExitDisqualifiedNodeFailOnStartOld(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 2,
		UplinkCount:      1,
		Reconfigure: testplanet.Reconfigure{Satellite: func(log *zap.Logger, index int, config *satellite.Config) {
			// This test can be removed entirely when we are using time-based GE everywhere.
			config.GracefulExit.TimeBased = false
		}},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]
		exitingNode := planet.StorageNodes[0]

		_, err := satellite.DB.OverlayCache().DisqualifyNode(ctx, exitingNode.ID(), time.Now(), overlay.DisqualificationReasonUnknown)
		require.NoError(t, err)

		conn, err := exitingNode.Dialer.DialNodeURL(ctx, satellite.NodeURL())
		require.NoError(t, err)
		defer ctx.Check(conn.Close)

		client := pb.NewDRPCSatelliteGracefulExitClient(conn)
		processClient, err := client.Process(ctx)
		require.NoError(t, err)

		// Process endpoint should return immediately if node is disqualified
		response, err := processClient.Recv()
		require.True(t, errs2.IsRPC(err, rpcstatus.FailedPrecondition))
		require.Nil(t, response)

		require.NoError(t, processClient.Close())

		// disqualified node should fail graceful exit
		exitStatus, err := satellite.Overlay.DB.GetExitStatus(ctx, exitingNode.ID())
		require.NoError(t, err)
		require.NotNil(t, exitStatus.ExitFinishedAt)
		require.False(t, exitStatus.ExitSuccess)
	})

}

func TestExitDisqualifiedNodeFailOnStart(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 2,
		UplinkCount:      1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: func(log *zap.Logger, index int, config *satellite.Config) {
				config.GracefulExit.TimeBased = true
			},
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]
		exitingNode := planet.StorageNodes[0]

		_, err := satellite.DB.OverlayCache().DisqualifyNode(ctx, exitingNode.ID(), time.Now(), overlay.DisqualificationReasonUnknown)
		require.NoError(t, err)

		conn, err := exitingNode.Dialer.DialNodeURL(ctx, satellite.NodeURL())
		require.NoError(t, err)
		defer ctx.Check(conn.Close)

		client := pb.NewDRPCSatelliteGracefulExitClient(conn)
		processClient, err := client.Process(ctx)
		require.NoError(t, err)

		// Process endpoint should return immediately if node is disqualified
		response, err := processClient.Recv()
		require.True(t, errs2.IsRPC(err, rpcstatus.FailedPrecondition))
		require.Nil(t, response)

		require.NoError(t, processClient.Close())

		// make sure GE was not initiated for the disqualified node
		exitStatus, err := satellite.Overlay.DB.GetExitStatus(ctx, exitingNode.ID())
		require.NoError(t, err)
		require.Nil(t, exitStatus.ExitInitiatedAt)
		require.False(t, exitStatus.ExitSuccess)
	})
}

func TestExitDisqualifiedNodeFailEventually(t *testing.T) {
	testTransfers(t, numObjects, numMultipartObjects, func(t *testing.T, ctx *testcontext.Context, nodeFullIDs map[storj.NodeID]*identity.FullIdentity, satellite *testplanet.Satellite, processClient exitProcessClient, exitingNode *storagenode.Peer, numPieces int) {
		var disqualifiedError error
		isDisqualified := false
		for {
			response, err := processClient.Recv()
			if errs.Is(err, io.EOF) {
				// Done
				break
			}
			if errs2.IsRPC(err, rpcstatus.FailedPrecondition) {
				disqualifiedError = err
				break
			}

			if !isDisqualified {
				_, err := satellite.DB.OverlayCache().DisqualifyNode(ctx, exitingNode.ID(), time.Now(), overlay.DisqualificationReasonUnknown)
				require.NoError(t, err)
			}

			switch m := response.GetMessage().(type) {
			case *pb.SatelliteMessage_TransferPiece:
				require.NotNil(t, m)

				pieceReader, err := exitingNode.Storage2.Store.Reader(ctx, satellite.ID(), m.TransferPiece.OriginalPieceId)
				require.NoError(t, err)

				header, err := pieceReader.GetPieceHeader()
				require.NoError(t, err)

				orderLimit := header.OrderLimit
				originalPieceHash := &pb.PieceHash{
					PieceId:       orderLimit.PieceId,
					Hash:          header.GetHash(),
					HashAlgorithm: header.HashAlgorithm,
					PieceSize:     pieceReader.Size(),
					Timestamp:     header.GetCreationTime(),
					Signature:     header.GetSignature(),
				}

				newPieceHash := &pb.PieceHash{
					PieceId:       m.TransferPiece.AddressedOrderLimit.Limit.PieceId,
					Hash:          originalPieceHash.Hash,
					HashAlgorithm: header.HashAlgorithm,
					PieceSize:     originalPieceHash.PieceSize,
					Timestamp:     time.Now(),
				}

				receivingNodeID := nodeFullIDs[m.TransferPiece.AddressedOrderLimit.Limit.StorageNodeId]
				require.NotNil(t, receivingNodeID)
				signer := signing.SignerFromFullIdentity(receivingNodeID)

				signedNewPieceHash, err := signing.SignPieceHash(ctx, signer, newPieceHash)
				require.NoError(t, err)

				success := &pb.StorageNodeMessage{
					Message: &pb.StorageNodeMessage_Succeeded{
						Succeeded: &pb.TransferSucceeded{
							OriginalPieceId:      m.TransferPiece.OriginalPieceId,
							OriginalPieceHash:    originalPieceHash,
							OriginalOrderLimit:   &orderLimit,
							ReplacementPieceHash: signedNewPieceHash,
						},
					},
				}
				err = processClient.Send(success)
				require.NoError(t, err)
			case *pb.SatelliteMessage_DeletePiece:
				continue
			default:
				require.FailNow(t, "should not reach this case: %#v", m)
			}
		}
		// check that the exit has failed due to node has been disqualified
		require.True(t, errs2.IsRPC(disqualifiedError, rpcstatus.FailedPrecondition))

		// check that the exit has completed and we have the correct transferred/failed values
		progress, err := satellite.DB.GracefulExit().GetProgress(ctx, exitingNode.ID())
		require.NoError(t, err)

		require.EqualValues(t, numPieces, progress.PiecesTransferred)

		// disqualified node should fail graceful exit
		exitStatus, err := satellite.Overlay.DB.GetExitStatus(ctx, exitingNode.ID())
		require.NoError(t, err)
		require.NotNil(t, exitStatus.ExitFinishedAt)
		require.False(t, exitStatus.ExitSuccess)
	})
}

func TestFailureHashMismatch(t *testing.T) {
	testTransfers(t, 1, 0, testFailureHashMismatch)
}

func TestFailureHashMismatchMultipart(t *testing.T) {
	testTransfers(t, 0, 1, testFailureHashMismatch)
}

func testFailureHashMismatch(t *testing.T, ctx *testcontext.Context, nodeFullIDs map[storj.NodeID]*identity.FullIdentity, satellite *testplanet.Satellite, processClient exitProcessClient, exitingNode *storagenode.Peer, numPieces int) {
	response, err := processClient.Recv()
	require.NoError(t, err)

	switch m := response.GetMessage().(type) {
	case *pb.SatelliteMessage_TransferPiece:
		require.NotNil(t, m)
		pieceReader, err := exitingNode.Storage2.Store.Reader(ctx, satellite.ID(), m.TransferPiece.OriginalPieceId)
		require.NoError(t, err)

		header, err := pieceReader.GetPieceHeader()
		require.NoError(t, err)

		orderLimit := header.OrderLimit
		originalPieceHash := &pb.PieceHash{
			PieceId:   orderLimit.PieceId,
			Hash:      header.GetHash(),
			PieceSize: pieceReader.Size(),
			Timestamp: header.GetCreationTime(),
			Signature: header.GetSignature(),
		}

		newPieceHash := &pb.PieceHash{
			PieceId:   m.TransferPiece.AddressedOrderLimit.Limit.PieceId,
			Hash:      originalPieceHash.Hash[:1],
			PieceSize: originalPieceHash.PieceSize,
			Timestamp: time.Now(),
		}

		receivingNodeID := nodeFullIDs[m.TransferPiece.AddressedOrderLimit.Limit.StorageNodeId]
		require.NotNil(t, receivingNodeID)
		signer := signing.SignerFromFullIdentity(receivingNodeID)

		signedNewPieceHash, err := signing.SignPieceHash(ctx, signer, newPieceHash)
		require.NoError(t, err)

		message := &pb.StorageNodeMessage{
			Message: &pb.StorageNodeMessage_Succeeded{
				Succeeded: &pb.TransferSucceeded{
					OriginalPieceId:      m.TransferPiece.OriginalPieceId,
					OriginalPieceHash:    originalPieceHash,
					OriginalOrderLimit:   &orderLimit,
					ReplacementPieceHash: signedNewPieceHash,
				},
			},
		}
		err = processClient.Send(message)
		require.NoError(t, err)
	default:
		require.FailNow(t, "should not reach this case: %#v", m)
	}

	response, err = processClient.Recv()
	require.NoError(t, err)

	switch m := response.GetMessage().(type) {
	case *pb.SatelliteMessage_ExitFailed:
		require.NotNil(t, m)
		require.NotNil(t, m.ExitFailed)
		require.Equal(t, m.ExitFailed.Reason, pb.ExitFailed_VERIFICATION_FAILED)

		node, err := satellite.DB.OverlayCache().Get(ctx, m.ExitFailed.NodeId)
		require.NoError(t, err)
		require.NotNil(t, node.Disqualified)
	default:
		require.FailNow(t, "should not reach this case: %#v", m)
	}

	// check that the exit has completed and we have the correct transferred/failed values
	progress, err := satellite.DB.GracefulExit().GetProgress(ctx, exitingNode.ID())
	require.NoError(t, err)

	require.Equal(t, int64(0), progress.PiecesTransferred)
	require.Equal(t, int64(1), progress.PiecesFailed)
}

func TestFailureUnknownError(t *testing.T) {
	testTransfers(t, 1, 0, func(t *testing.T, ctx *testcontext.Context, nodeFullIDs map[storj.NodeID]*identity.FullIdentity, satellite *testplanet.Satellite, processClient exitProcessClient, exitingNode *storagenode.Peer, numPieces int) {
		response, err := processClient.Recv()
		require.NoError(t, err)

		switch m := response.GetMessage().(type) {
		case *pb.SatelliteMessage_TransferPiece:
			require.NotNil(t, m)
			message := &pb.StorageNodeMessage{
				Message: &pb.StorageNodeMessage_Failed{
					Failed: &pb.TransferFailed{
						Error:           pb.TransferFailed_UNKNOWN,
						OriginalPieceId: m.TransferPiece.OriginalPieceId,
					},
				},
			}
			err = processClient.Send(message)
			require.NoError(t, err)
		default:
			require.FailNow(t, "should not reach this case: %#v", m)
		}

		response, err = processClient.Recv()
		require.NoError(t, err)

		switch m := response.GetMessage().(type) {
		case *pb.SatelliteMessage_TransferPiece:
			require.NotNil(t, m)
		default:
			require.FailNow(t, "should not reach this case: %#v", m)
		}

		// check that the exit has completed and we have the correct transferred/failed values
		progress, err := satellite.DB.GracefulExit().GetProgress(ctx, exitingNode.ID())
		require.NoError(t, err)

		require.Equal(t, int64(0), progress.PiecesTransferred)
		require.Equal(t, int64(0), progress.PiecesFailed)
	})
}

func TestFailureUplinkSignature(t *testing.T) {
	testTransfers(t, 1, 0, func(t *testing.T, ctx *testcontext.Context, nodeFullIDs map[storj.NodeID]*identity.FullIdentity, satellite *testplanet.Satellite, processClient exitProcessClient, exitingNode *storagenode.Peer, numPieces int) {
		response, err := processClient.Recv()
		require.NoError(t, err)

		switch m := response.GetMessage().(type) {
		case *pb.SatelliteMessage_TransferPiece:
			require.NotNil(t, m)
			pieceReader, err := exitingNode.Storage2.Store.Reader(ctx, satellite.ID(), m.TransferPiece.OriginalPieceId)
			require.NoError(t, err)

			header, err := pieceReader.GetPieceHeader()
			require.NoError(t, err)

			orderLimit := header.OrderLimit
			orderLimit.UplinkPublicKey = storj.PiecePublicKey{}

			originalPieceHash := &pb.PieceHash{
				PieceId:   orderLimit.PieceId,
				Hash:      header.GetHash(),
				PieceSize: pieceReader.Size(),
				Timestamp: header.GetCreationTime(),
				Signature: header.GetSignature(),
			}

			newPieceHash := &pb.PieceHash{
				PieceId:   m.TransferPiece.AddressedOrderLimit.Limit.PieceId,
				Hash:      originalPieceHash.Hash,
				PieceSize: originalPieceHash.PieceSize,
				Timestamp: time.Now(),
			}

			receivingNodeID := nodeFullIDs[m.TransferPiece.AddressedOrderLimit.Limit.StorageNodeId]
			require.NotNil(t, receivingNodeID)
			signer := signing.SignerFromFullIdentity(receivingNodeID)

			signedNewPieceHash, err := signing.SignPieceHash(ctx, signer, newPieceHash)
			require.NoError(t, err)

			message := &pb.StorageNodeMessage{
				Message: &pb.StorageNodeMessage_Succeeded{
					Succeeded: &pb.TransferSucceeded{
						OriginalPieceId:      m.TransferPiece.OriginalPieceId,
						OriginalPieceHash:    originalPieceHash,
						OriginalOrderLimit:   &orderLimit,
						ReplacementPieceHash: signedNewPieceHash,
					},
				},
			}
			err = processClient.Send(message)
			require.NoError(t, err)
		default:
			require.FailNow(t, "should not reach this case: %#v", m)
		}

		response, err = processClient.Recv()
		require.NoError(t, err)

		switch m := response.GetMessage().(type) {
		case *pb.SatelliteMessage_ExitFailed:
			require.NotNil(t, m)
			require.NotNil(t, m.ExitFailed)
			require.Equal(t, m.ExitFailed.Reason, pb.ExitFailed_VERIFICATION_FAILED)

			node, err := satellite.DB.OverlayCache().Get(ctx, m.ExitFailed.NodeId)
			require.NoError(t, err)
			require.NotNil(t, node.Disqualified)
		default:
			require.FailNow(t, "should not reach this case: %#v", m)
		}

		// check that the exit has completed and we have the correct transferred/failed values
		progress, err := satellite.DB.GracefulExit().GetProgress(ctx, exitingNode.ID())
		require.NoError(t, err)

		require.Equal(t, int64(0), progress.PiecesTransferred)
		require.Equal(t, int64(1), progress.PiecesFailed)
	})
}

func TestSuccessSegmentUpdate(t *testing.T) {
	testTransfers(t, 1, 0, testSuccessSegmentUpdate)
}

func TestSuccessSegmentUpdateMultipart(t *testing.T) {
	testTransfers(t, 0, 1, testSuccessSegmentUpdate)
}

func testSuccessSegmentUpdate(t *testing.T, ctx *testcontext.Context, nodeFullIDs map[storj.NodeID]*identity.FullIdentity, satellite *testplanet.Satellite, processClient exitProcessClient, exitingNode *storagenode.Peer, numPieces int) {
	var recNodeID storj.NodeID

	response, err := processClient.Recv()
	require.NoError(t, err)

	switch m := response.GetMessage().(type) {
	case *pb.SatelliteMessage_TransferPiece:
		require.NotNil(t, m)

		pieceReader, err := exitingNode.Storage2.Store.Reader(ctx, satellite.ID(), m.TransferPiece.OriginalPieceId)
		require.NoError(t, err)

		header, err := pieceReader.GetPieceHeader()
		require.NoError(t, err)

		orderLimit := header.OrderLimit
		originalPieceHash := &pb.PieceHash{
			PieceId:       orderLimit.PieceId,
			Hash:          header.GetHash(),
			HashAlgorithm: header.HashAlgorithm,
			PieceSize:     pieceReader.Size(),
			Timestamp:     header.GetCreationTime(),
			Signature:     header.GetSignature(),
		}

		newPieceHash := &pb.PieceHash{
			PieceId:       m.TransferPiece.AddressedOrderLimit.Limit.PieceId,
			Hash:          originalPieceHash.Hash,
			HashAlgorithm: header.HashAlgorithm,
			PieceSize:     originalPieceHash.PieceSize,
			Timestamp:     time.Now(),
		}

		receivingIdentity := nodeFullIDs[m.TransferPiece.AddressedOrderLimit.Limit.StorageNodeId]
		require.NotNil(t, receivingIdentity)

		// get the receiving node piece count before processing
		recNodeID = receivingIdentity.ID

		signer := signing.SignerFromFullIdentity(receivingIdentity)

		signedNewPieceHash, err := signing.SignPieceHash(ctx, signer, newPieceHash)
		require.NoError(t, err)

		success := &pb.StorageNodeMessage{
			Message: &pb.StorageNodeMessage_Succeeded{
				Succeeded: &pb.TransferSucceeded{
					OriginalPieceId:      m.TransferPiece.OriginalPieceId,
					OriginalPieceHash:    originalPieceHash,
					OriginalOrderLimit:   &orderLimit,
					ReplacementPieceHash: signedNewPieceHash,
				},
			},
		}
		err = processClient.Send(success)
		require.NoError(t, err)
	default:
		require.FailNow(t, "did not get a TransferPiece message")
	}

	response, err = processClient.Recv()
	require.NoError(t, err)

	switch response.GetMessage().(type) {
	case *pb.SatelliteMessage_DeletePiece:
		// expect the delete piece message
	default:
		require.FailNow(t, "did not get a DeletePiece message")
	}
	// check that the exit has completed and we have the correct transferred/failed values
	progress, err := satellite.DB.GracefulExit().GetProgress(ctx, exitingNode.ID())
	require.NoError(t, err)

	require.EqualValues(t, numPieces, progress.PiecesTransferred)
	// even though we failed 1, it eventually succeeded, so the count should be 0
	require.EqualValues(t, 0, progress.PiecesFailed)

	segments, err := satellite.Metabase.DB.TestingAllSegments(ctx)
	require.NoError(t, err)
	require.Len(t, segments, 1)
	found := 0
	require.True(t, len(segments[0].Pieces) > 0)
	for _, piece := range segments[0].Pieces {
		require.NotEqual(t, exitingNode.ID(), piece.StorageNode)
		if piece.StorageNode == recNodeID {
			found++
		}
	}
	require.Equal(t, 1, found)
}

func TestUpdateSegmentFailure_DuplicatedNodeID(t *testing.T) {
	testTransfers(t, 1, 0, testUpdateSegmentFailureDuplicatedNodeID)
}

func TestUpdateSegmentFailure_DuplicatedNodeIDMultipart(t *testing.T) {
	testTransfers(t, 0, 1, testUpdateSegmentFailureDuplicatedNodeID)
}

func testUpdateSegmentFailureDuplicatedNodeID(t *testing.T, ctx *testcontext.Context, nodeFullIDs map[storj.NodeID]*identity.FullIdentity, satellite *testplanet.Satellite, processClient exitProcessClient, exitingNode *storagenode.Peer, numPieces int) {
	response, err := processClient.Recv()
	require.NoError(t, err)

	var firstRecNodeID storj.NodeID
	var pieceID storj.PieceID
	switch m := response.GetMessage().(type) {
	case *pb.SatelliteMessage_TransferPiece:
		firstRecNodeID = m.TransferPiece.AddressedOrderLimit.Limit.StorageNodeId
		pieceID = m.TransferPiece.OriginalPieceId

		pieceReader, err := exitingNode.Storage2.Store.Reader(ctx, satellite.ID(), pieceID)
		require.NoError(t, err)

		header, err := pieceReader.GetPieceHeader()
		require.NoError(t, err)

		orderLimit := header.OrderLimit
		originalPieceHash := &pb.PieceHash{
			PieceId:       orderLimit.PieceId,
			Hash:          header.GetHash(),
			HashAlgorithm: header.HashAlgorithm,
			PieceSize:     pieceReader.Size(),
			Timestamp:     header.GetCreationTime(),
			Signature:     header.GetSignature(),
		}

		newPieceHash := &pb.PieceHash{
			PieceId:       m.TransferPiece.AddressedOrderLimit.Limit.PieceId,
			Hash:          originalPieceHash.Hash,
			HashAlgorithm: header.HashAlgorithm,
			PieceSize:     originalPieceHash.PieceSize,
			Timestamp:     time.Now(),
		}

		receivingNodeIdentity := nodeFullIDs[m.TransferPiece.AddressedOrderLimit.Limit.StorageNodeId]
		require.NotNil(t, receivingNodeIdentity)
		signer := signing.SignerFromFullIdentity(receivingNodeIdentity)

		signedNewPieceHash, err := signing.SignPieceHash(ctx, signer, newPieceHash)
		require.NoError(t, err)

		success := &pb.StorageNodeMessage{
			Message: &pb.StorageNodeMessage_Succeeded{
				Succeeded: &pb.TransferSucceeded{
					OriginalPieceId:      pieceID,
					OriginalPieceHash:    originalPieceHash,
					OriginalOrderLimit:   &orderLimit,
					ReplacementPieceHash: signedNewPieceHash,
				},
			},
		}

		// update segment to include the new receiving node before responding to satellite
		segments, err := satellite.Metabase.DB.TestingAllSegments(ctx)
		require.NoError(t, err)
		require.Len(t, segments, 1)
		require.True(t, len(segments[0].Pieces) > 0)

		pieceToRemove := make(metabase.Pieces, 1)
		pieceToAdd := make(metabase.Pieces, 1)
		pieces := segments[0].Pieces

		for _, piece := range pieces {
			if pieceToRemove[0] == (metabase.Piece{}) && piece.StorageNode != exitingNode.ID() {
				pieceToRemove[0] = piece
				continue
			}
		}

		pieceToAdd[0] = metabase.Piece{
			Number:      pieceToRemove[0].Number,
			StorageNode: firstRecNodeID,
		}

		err = satellite.GracefulExit.Endpoint.UpdatePiecesCheckDuplicates(ctx, segments[0], pieceToAdd, pieceToRemove, false)
		require.NoError(t, err)

		err = processClient.Send(success)
		require.NoError(t, err)
	default:
		require.FailNow(t, "should not reach this case: %#v", m)
	}

	response, err = processClient.Recv()
	require.NoError(t, err)

	switch m := response.GetMessage().(type) {
	case *pb.SatelliteMessage_TransferPiece:
		// validate we get a new node to transfer too
		require.True(t, m.TransferPiece.OriginalPieceId == pieceID)
		require.True(t, m.TransferPiece.AddressedOrderLimit.Limit.StorageNodeId != firstRecNodeID)
	default:
		require.FailNow(t, "should not reach this case: %#v", m)
	}

	// check exiting node is still in the segment
	segments, err := satellite.Metabase.DB.TestingAllSegments(ctx)
	require.NoError(t, err)
	require.Len(t, segments, 1)

	require.True(t, len(segments[0].Pieces) > 0)

	pieces := segments[0].Pieces
	pieceMap := make(map[storj.NodeID]int)
	for _, piece := range pieces {
		pieceMap[piece.StorageNode]++
	}

	exitingNodeID := exitingNode.ID()
	count, ok := pieceMap[exitingNodeID]
	require.True(t, ok)
	require.Equal(t, 1, count)
	count, ok = pieceMap[firstRecNodeID]
	require.True(t, ok)
	require.Equal(t, 1, count)
}

func TestExitDisabled(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 2,
		UplinkCount:      1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: func(log *zap.Logger, index int, config *satellite.Config) {
				config.GracefulExit.Enabled = false
			},
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]
		exitingNode := planet.StorageNodes[0]

		require.Nil(t, satellite.GracefulExit.Endpoint)

		conn, err := exitingNode.Dialer.DialNodeURL(ctx, satellite.NodeURL())
		require.NoError(t, err)
		defer ctx.Check(conn.Close)

		client := pb.NewDRPCSatelliteGracefulExitClient(conn)
		processClient, err := client.Process(ctx)
		require.NoError(t, err)

		// Process endpoint should return immediately if GE is disabled
		response, err := processClient.Recv()
		require.Error(t, err)
		// drpc will return "Unknown"
		require.True(t, errs2.IsRPC(err, rpcstatus.Unknown))
		require.Nil(t, response)
	})
}

func TestSegmentChangedOrDeleted(t *testing.T) {
	successThreshold := 4
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: successThreshold + 1,
		UplinkCount:      1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: testplanet.Combine(
				testplanet.ReconfigureRS(2, 3, successThreshold, successThreshold),
				func(log *zap.Logger, index int, config *satellite.Config) {
					// This test can be removed entirely when we are using time-based GE everywhere.
					config.GracefulExit.TimeBased = false
				},
			),
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		uplinkPeer := planet.Uplinks[0]
		satellite := planet.Satellites[0]

		err := uplinkPeer.Upload(ctx, satellite, "testbucket", "test/path0", testrand.Bytes(5*memory.KiB))
		require.NoError(t, err)
		err = uplinkPeer.Upload(ctx, satellite, "testbucket", "test/path1", testrand.Bytes(5*memory.KiB))
		require.NoError(t, err)

		// check that there are no exiting nodes.
		exitingNodes, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 0)

		exitingNode, err := findNodeToExit(ctx, planet, 2)
		require.NoError(t, err)

		exitRequest := &overlay.ExitStatusRequest{
			NodeID:          exitingNode.ID(),
			ExitInitiatedAt: time.Now(),
		}

		_, err = satellite.DB.OverlayCache().UpdateExitStatus(ctx, exitRequest)
		require.NoError(t, err)
		err = satellite.DB.GracefulExit().IncrementProgress(ctx, exitingNode.ID(), 0, 0, 0)
		require.NoError(t, err)

		exitingNodes, err = satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 1)
		require.Equal(t, exitingNode.ID(), exitingNodes[0].NodeID)

		// run the satellite ranged loop to build the transfer queue.
		_, err = satellite.RangedLoop.RangedLoop.Service.RunOnce(ctx)
		require.NoError(t, err)

		// make sure all the pieces are in the transfer queue
		incomplete, err := satellite.DB.GracefulExit().GetIncomplete(ctx, exitingNode.ID(), 10, 0)
		require.NoError(t, err)
		require.Len(t, incomplete, 2)

		// updating the first object and deleting the second. this will cause a root piece ID change which will result in
		// a successful graceful exit instead of a request to transfer pieces since the root piece IDs will have changed.
		err = uplinkPeer.Upload(ctx, satellite, "testbucket", "test/path0", testrand.Bytes(5*memory.KiB))
		require.NoError(t, err)
		err = uplinkPeer.DeleteObject(ctx, satellite, "testbucket", "test/path1")
		require.NoError(t, err)

		// reconnect to the satellite.
		conn, err := exitingNode.Dialer.DialNodeURL(ctx, satellite.NodeURL())
		require.NoError(t, err)
		defer ctx.Check(conn.Close)

		client := pb.NewDRPCSatelliteGracefulExitClient(conn)

		c, err := client.Process(ctx)
		require.NoError(t, err)
		defer ctx.Check(c.CloseSend)

		response, err := c.Recv()
		require.NoError(t, err)

		// we expect an exit completed b/c there is nothing to do here
		switch m := response.GetMessage().(type) {
		case *pb.SatelliteMessage_ExitCompleted:
			signee := signing.SigneeFromPeerIdentity(satellite.Identity.PeerIdentity())
			err = signing.VerifyExitCompleted(ctx, signee, m.ExitCompleted)
			require.NoError(t, err)

			exitStatus, err := satellite.DB.OverlayCache().GetExitStatus(ctx, exitingNode.ID())
			require.NoError(t, err)
			require.NotNil(t, exitStatus.ExitFinishedAt)
			require.True(t, exitStatus.ExitSuccess)
		default:
			require.FailNow(t, "should not reach this case: %#v", m)
		}

		queueItems, err := satellite.DB.GracefulExit().GetIncomplete(ctx, exitingNode.ID(), 2, 0)
		require.NoError(t, err)
		require.Len(t, queueItems, 0)
	})
}

func TestSegmentChangedOrDeletedMultipart(t *testing.T) {
	successThreshold := 4
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: successThreshold + 1,
		UplinkCount:      1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: testplanet.Combine(
				testplanet.ReconfigureRS(2, 3, successThreshold, successThreshold),
				func(log *zap.Logger, index int, config *satellite.Config) {
					// This test can be removed entirely when we are using time-based GE everywhere.
					config.GracefulExit.TimeBased = false
				},
			),
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		uplinkPeer := planet.Uplinks[0]
		satellite := planet.Satellites[0]

		project, err := uplinkPeer.GetProject(ctx, satellite)
		require.NoError(t, err)
		defer func() { require.NoError(t, project.Close()) }()

		_, err = project.EnsureBucket(ctx, "testbucket")
		require.NoError(t, err)

		// TODO: activate when an object part can be overwritten
		// info0, err := multipart.NewMultipartUpload(ctx, project, "testbucket", "test/path0", nil)
		// require.NoError(t, err)
		// _, err = multipart.PutObjectPart(ctx, project, "testbucket", "test/path0", info0.StreamID, 1,
		//	etag.NewHashReader(bytes.NewReader(testrand.Bytes(5*memory.KiB)), sha256.New()))
		// require.NoError(t, err)

		info1, err := project.BeginUpload(ctx, "testbucket", "test/path1", nil)
		require.NoError(t, err)

		upload, err := project.UploadPart(ctx, "testbucket", "test/path1", info1.UploadID, 1)
		require.NoError(t, err)
		_, err = upload.Write(testrand.Bytes(5 * memory.KiB))
		require.NoError(t, err)
		require.NoError(t, upload.Commit())

		// check that there are no exiting nodes.
		exitingNodes, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 0)

		exitingNode, err := findNodeToExit(ctx, planet, 2)
		require.NoError(t, err)

		exitRequest := &overlay.ExitStatusRequest{
			NodeID:          exitingNode.ID(),
			ExitInitiatedAt: time.Now(),
		}

		_, err = satellite.DB.OverlayCache().UpdateExitStatus(ctx, exitRequest)
		require.NoError(t, err)
		err = satellite.DB.GracefulExit().IncrementProgress(ctx, exitingNode.ID(), 0, 0, 0)
		require.NoError(t, err)

		exitingNodes, err = satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 1)
		require.Equal(t, exitingNode.ID(), exitingNodes[0].NodeID)

		// run the satellite ranged loop to build the transfer queue.
		_, err = satellite.RangedLoop.RangedLoop.Service.RunOnce(ctx)
		require.NoError(t, err)

		// make sure all the pieces are in the transfer queue
		incomplete, err := satellite.DB.GracefulExit().GetIncomplete(ctx, exitingNode.ID(), 10, 0)
		require.NoError(t, err)
		require.Len(t, incomplete, 1)
		// TODO: change to this when an object part can be overwritten
		// require.Len(t, incomplete, 2)

		// updating the first object and deleting the second. this will cause a root piece ID change which will result in
		// a successful graceful exit instead of a request to transfer pieces since the root piece IDs will have changed.
		// TODO: activate when an object part can be overwritten
		// _, err = multipart.PutObjectPart(ctx, project, "testbucket", "test/path0", info0.StreamID, 1, bytes.NewReader(testrand.Bytes(5*memory.KiB)))
		// require.NoError(t, err)
		err = project.AbortUpload(ctx, "testbucket", "test/path1", info1.UploadID)
		require.NoError(t, err)

		// reconnect to the satellite.
		conn, err := exitingNode.Dialer.DialNodeURL(ctx, satellite.NodeURL())
		require.NoError(t, err)
		defer ctx.Check(conn.Close)

		client := pb.NewDRPCSatelliteGracefulExitClient(conn)

		c, err := client.Process(ctx)
		require.NoError(t, err)
		defer ctx.Check(c.CloseSend)

		response, err := c.Recv()
		require.NoError(t, err)

		// we expect an exit completed b/c there is nothing to do here
		switch m := response.GetMessage().(type) {
		case *pb.SatelliteMessage_ExitCompleted:
			signee := signing.SigneeFromPeerIdentity(satellite.Identity.PeerIdentity())
			err = signing.VerifyExitCompleted(ctx, signee, m.ExitCompleted)
			require.NoError(t, err)

			exitStatus, err := satellite.DB.OverlayCache().GetExitStatus(ctx, exitingNode.ID())
			require.NoError(t, err)
			require.NotNil(t, exitStatus.ExitFinishedAt)
			require.True(t, exitStatus.ExitSuccess)
		default:
			require.FailNow(t, "should not reach this case: %#v", m)
		}

		queueItems, err := satellite.DB.GracefulExit().GetIncomplete(ctx, exitingNode.ID(), 2, 0)
		require.NoError(t, err)
		require.Len(t, queueItems, 0)
	})
}

func TestFailureNotFound(t *testing.T) {
	testTransfers(t, 1, 0, func(t *testing.T, ctx *testcontext.Context, nodeFullIDs map[storj.NodeID]*identity.FullIdentity, satellite *testplanet.Satellite, processClient exitProcessClient, exitingNode *storagenode.Peer, numPieces int) {
		response, err := processClient.Recv()
		require.NoError(t, err)

		switch m := response.GetMessage().(type) {
		case *pb.SatelliteMessage_TransferPiece:
			require.NotNil(t, m)

			message := &pb.StorageNodeMessage{
				Message: &pb.StorageNodeMessage_Failed{
					Failed: &pb.TransferFailed{
						OriginalPieceId: m.TransferPiece.OriginalPieceId,
						Error:           pb.TransferFailed_NOT_FOUND,
					},
				},
			}
			err = processClient.Send(message)
			require.NoError(t, err)
		default:
			require.FailNow(t, "should not reach this case: %#v", m)
		}

		response, err = processClient.Recv()
		require.NoError(t, err)

		switch m := response.GetMessage().(type) {
		case *pb.SatelliteMessage_ExitFailed:
			require.NotNil(t, m)
			require.NotNil(t, m.ExitFailed)
			require.Equal(t, m.ExitFailed.Reason, pb.ExitFailed_OVERALL_FAILURE_PERCENTAGE_EXCEEDED)

			node, err := satellite.DB.OverlayCache().Get(ctx, m.ExitFailed.NodeId)
			require.NoError(t, err)
			require.NotNil(t, node.Disqualified)
		default:
			require.FailNow(t, "should not reach this case: %#v", m)
		}

		// check that node is no longer in the segment

		segments, err := satellite.Metabase.DB.TestingAllSegments(ctx)
		require.NoError(t, err)
		require.Len(t, segments, 1)

		for _, piece := range segments[0].Pieces {
			require.NotEqual(t, piece.StorageNode, exitingNode.ID())
		}

		// check that the exit has completed and we have the correct transferred/failed values
		progress, err := satellite.DB.GracefulExit().GetProgress(ctx, exitingNode.ID())
		require.NoError(t, err)

		require.Equal(t, int64(0), progress.PiecesTransferred)
		require.Equal(t, int64(1), progress.PiecesFailed)
	})

}

func TestFailureStorageNodeIgnoresTransferMessages(t *testing.T) {
	var maxOrderLimitSendCount = 3
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 5,
		UplinkCount:      1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: testplanet.Combine(
				func(log *zap.Logger, index int, config *satellite.Config) {
					// We don't care whether a node gracefully exits or not in this test,
					// so we set the max failures percentage extra high.
					config.GracefulExit.OverallMaxFailuresPercentage = 101
					config.GracefulExit.MaxOrderLimitSendCount = maxOrderLimitSendCount
					// This test can be removed entirely when we are using time-based GE everywhere.
					config.GracefulExit.TimeBased = false
				},
				testplanet.ReconfigureRS(2, 3, 4, 4),
			),
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		uplinkPeer := planet.Uplinks[0]
		satellite := planet.Satellites[0]

		nodeFullIDs := make(map[storj.NodeID]*identity.FullIdentity)
		for _, node := range planet.StorageNodes {
			nodeFullIDs[node.ID()] = node.Identity
		}

		err := uplinkPeer.Upload(ctx, satellite, "testbucket", "test/path", testrand.Bytes(5*memory.KiB))
		require.NoError(t, err)

		// check that there are no exiting nodes.
		exitingNodes, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 0)

		exitingNode, err := findNodeToExit(ctx, planet, 1)
		require.NoError(t, err)

		// connect to satellite so we initiate the exit.
		conn, err := exitingNode.Dialer.DialNodeURL(ctx, satellite.NodeURL())
		require.NoError(t, err)
		defer ctx.Check(conn.Close)

		client := pb.NewDRPCSatelliteGracefulExitClient(conn)

		c, err := client.Process(ctx)
		require.NoError(t, err)

		response, err := c.Recv()
		require.NoError(t, err)

		// should get a NotReady since the metainfo loop would not be finished at this point.
		switch m := response.GetMessage().(type) {
		case *pb.SatelliteMessage_NotReady:
			// now check that the exiting node is initiated.
			exitingNodes, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
			require.NoError(t, err)
			require.Len(t, exitingNodes, 1)

			require.Equal(t, exitingNode.ID(), exitingNodes[0].NodeID)
		default:
			require.FailNow(t, "should not reach this case: %#v", m)
		}
		// close the old client
		require.NoError(t, c.CloseSend())

		// run the satellite ranged loop to build the transfer queue.
		_, err = satellite.RangedLoop.RangedLoop.Service.RunOnce(ctx)
		require.NoError(t, err)

		// make sure all the pieces are in the transfer queue
		_, err = satellite.DB.GracefulExit().GetIncomplete(ctx, exitingNode.ID(), 1, 0)
		require.NoError(t, err)

		var messageCount int

		// We need to label this outer loop so that we're able to exit it from the inner loop.
		// The outer loop is for sending the request from node to satellite multiple times.
		// The inner loop is for reading the response.
	MessageLoop:
		for {
			var unknownMsgSent bool
			c, err := client.Process(ctx)
			require.NoError(t, err)

			for {
				response, err := c.Recv()
				if unknownMsgSent {
					require.Error(t, err)
					break
				} else {
					require.NoError(t, err)
				}

				switch m := response.GetMessage().(type) {
				case *pb.SatelliteMessage_ExitCompleted:
					break MessageLoop
				case *pb.SatelliteMessage_TransferPiece:
					messageCount++
					unknownMsgSent = true
					// We send an unknown message because we want to fail the
					// transfer message request we get from the satellite.
					// This allows us to keep the conn open but repopulate
					// the pending queue.
					err = c.Send(&pb.StorageNodeMessage{})
					require.NoError(t, err)
					require.NoError(t, c.CloseSend())
				default:
					require.FailNow(t, "should not reach this case: %#v", m)
				}
			}
		}
		require.Equal(t, messageCount, maxOrderLimitSendCount)

		// make sure not responding piece not in queue
		incompletes, err := satellite.DB.GracefulExit().GetIncomplete(ctx, exitingNode.ID(), 10, 0)
		require.NoError(t, err)
		require.Len(t, incompletes, 0)

		// check that the exit has completed and we have the correct transferred/failed values
		progress, err := satellite.DB.GracefulExit().GetProgress(ctx, exitingNode.ID())
		require.NoError(t, err)
		require.EqualValues(t, 1, progress.PiecesFailed)
		status, err := satellite.DB.OverlayCache().GetExitStatus(ctx, exitingNode.ID())
		require.NoError(t, err)
		require.NotNil(t, status.ExitFinishedAt)
	})
}

func TestIneligibleNodeAge(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 5,
		Reconfigure: testplanet.Reconfigure{
			Satellite: func(log *zap.Logger, index int, config *satellite.Config) {
				// Set the required node age to 1 month.
				config.GracefulExit.NodeMinAgeInMonths = 1
				config.GracefulExit.TimeBased = true
			},
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]

		// check that there are no exiting nodes.
		exitingNodes, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 0)

		exitingNode := planet.StorageNodes[0]

		// try to initiate GE; expect to get a node ineligible error
		response, err := callProcess(ctx, exitingNode, satellite)
		require.Error(t, err)
		require.Nil(t, response)
		require.True(t, errs2.IsRPC(err, rpcstatus.FailedPrecondition))

		// check that there are still no exiting nodes
		exitingNodes, err = satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 0)
	})
}

func TestIneligibleNodeAgeOld(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 5,
		UplinkCount:      1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: testplanet.Combine(
				func(log *zap.Logger, index int, config *satellite.Config) {
					// Set the required node age to 1 month.
					config.GracefulExit.NodeMinAgeInMonths = 1
				},
				testplanet.ReconfigureRS(2, 3, 4, 4),
			),
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		uplinkPeer := planet.Uplinks[0]
		satellite := planet.Satellites[0]

		nodeFullIDs := make(map[storj.NodeID]*identity.FullIdentity)
		for _, node := range planet.StorageNodes {
			nodeFullIDs[node.ID()] = node.Identity
		}

		err := uplinkPeer.Upload(ctx, satellite, "testbucket", "test/path", testrand.Bytes(5*memory.KiB))
		require.NoError(t, err)

		// check that there are no exiting nodes.
		exitingNodes, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 0)

		exitingNode, err := findNodeToExit(ctx, planet, 1)
		require.NoError(t, err)

		// connect to satellite so we initiate the exit.
		conn, err := exitingNode.Dialer.DialNodeURL(ctx, satellite.NodeURL())
		require.NoError(t, err)
		defer ctx.Check(conn.Close)

		client := pb.NewDRPCSatelliteGracefulExitClient(conn)

		c, err := client.Process(ctx)
		require.NoError(t, err)

		_, err = c.Recv()
		// expect the node ineligible error here
		require.Error(t, err)
		require.True(t, errs2.IsRPC(err, rpcstatus.FailedPrecondition))

		// check that there are still no exiting nodes
		exitingNodes, err = satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 0)

		// close the old client
		require.NoError(t, c.CloseSend())
	})
}

func testTransfers(t *testing.T, objects int, multipartObjects int, verifier func(t *testing.T, ctx *testcontext.Context, nodeFullIDs map[storj.NodeID]*identity.FullIdentity, satellite *testplanet.Satellite, processClient exitProcessClient, exitingNode *storagenode.Peer, numPieces int)) {
	const successThreshold = 4
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: successThreshold + 1,
		UplinkCount:      1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: testplanet.Combine(
				testplanet.ReconfigureRS(2, 3, successThreshold, successThreshold),
				func(log *zap.Logger, index int, config *satellite.Config) {
					// These tests can be removed entirely when we are using time-based GE everywhere.
					config.GracefulExit.TimeBased = false
				},
			),
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		uplinkPeer := planet.Uplinks[0]
		satellite := planet.Satellites[0]

		project, err := uplinkPeer.GetProject(ctx, satellite)
		require.NoError(t, err)
		defer func() { require.NoError(t, project.Close()) }()

		_, err = project.EnsureBucket(ctx, "testbucket")
		require.NoError(t, err)

		nodeFullIDs := make(map[storj.NodeID]*identity.FullIdentity)
		for _, node := range planet.StorageNodes {
			nodeFullIDs[node.ID()] = node.Identity
		}

		hashes := []pb.PieceHashAlgorithm{pb.PieceHashAlgorithm_BLAKE3, pb.PieceHashAlgorithm_SHA256}
		for i := 0; i < objects; i++ {
			err := uplinkPeer.Upload(piecestore.WithPieceHashAlgo(ctx, hashes[i%len(hashes)]), satellite, "testbucket", "test/path-"+strconv.Itoa(i), testrand.Bytes(5*memory.KiB))
			require.NoError(t, err)
		}

		for i := 0; i < multipartObjects; i++ {
			objectName := "test/multipart" + strconv.Itoa(i)

			info, err := project.BeginUpload(ctx, "testbucket", objectName, nil)
			require.NoError(t, err)

			upload, err := project.UploadPart(ctx, "testbucket", objectName, info.UploadID, 1)
			require.NoError(t, err)
			_, err = upload.Write(testrand.Bytes(5 * memory.KiB))
			require.NoError(t, err)
			require.NoError(t, upload.Commit())
		}

		// check that there are no exiting nodes.
		exitingNodes, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 0)

		exitingNode, err := findNodeToExit(ctx, planet, objects)
		require.NoError(t, err)

		// connect to satellite so we initiate the exit.
		conn, err := exitingNode.Dialer.DialNodeURL(ctx, satellite.NodeURL())
		require.NoError(t, err)
		defer ctx.Check(conn.Close)

		client := pb.NewDRPCSatelliteGracefulExitClient(conn)

		c, err := client.Process(ctx)
		require.NoError(t, err)

		response, err := c.Recv()
		require.NoError(t, err)

		// should get a NotReady since the metainfo loop would not be finished at this point.
		switch m := response.GetMessage().(type) {
		case *pb.SatelliteMessage_NotReady:
			// now check that the exiting node is initiated.
			exitingNodes, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
			require.NoError(t, err)
			require.Len(t, exitingNodes, 1)

			require.Equal(t, exitingNode.ID(), exitingNodes[0].NodeID)
		default:
			require.FailNow(t, "should not reach this case: %#v", m)
		}
		// close the old client
		require.NoError(t, c.CloseSend())

		// run the satellite ranged loop to build the transfer queue.
		_, err = satellite.RangedLoop.RangedLoop.Service.RunOnce(ctx)
		require.NoError(t, err)

		// make sure all the pieces are in the transfer queue
		incompleteTransfers, err := satellite.DB.GracefulExit().GetIncomplete(ctx, exitingNode.ID(), objects+multipartObjects, 0)
		require.NoError(t, err)

		// connect to satellite again to start receiving transfers
		c, err = client.Process(ctx)
		require.NoError(t, err)
		defer ctx.Check(c.CloseSend)

		verifier(t, ctx, nodeFullIDs, satellite, c, exitingNode.Peer, len(incompleteTransfers))
	})
}

func findNodeToExit(ctx context.Context, planet *testplanet.Planet, objects int) (*testplanet.StorageNode, error) {
	satellite := planet.Satellites[0]

	pieceCountMap := make(map[storj.NodeID]int, len(planet.StorageNodes))
	for _, node := range planet.StorageNodes {
		pieceCountMap[node.ID()] = 0
	}

	segments, err := satellite.Metabase.DB.TestingAllSegments(ctx)
	if err != nil {
		return nil, err
	}
	for _, segment := range segments {
		for _, piece := range segment.Pieces {
			pieceCountMap[piece.StorageNode]++
		}
	}

	var exitingNodeID storj.NodeID
	maxCount := 0
	for k, v := range pieceCountMap {
		if exitingNodeID.IsZero() {
			exitingNodeID = k
			maxCount = v
			continue
		}
		if v > maxCount {
			exitingNodeID = k
			maxCount = v
		}
	}

	return planet.FindNode(exitingNodeID), nil
}

func TestUpdatePiecesCheckDuplicates(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 3, UplinkCount: 1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: testplanet.ReconfigureRS(1, 1, 3, 3),
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]
		uplinkPeer := planet.Uplinks[0]
		path := "test/path"

		err := uplinkPeer.Upload(ctx, satellite, "test1", path, testrand.Bytes(5*memory.KiB))
		require.NoError(t, err)

		segments, err := satellite.Metabase.DB.TestingAllSegments(ctx)
		require.NoError(t, err)
		require.Len(t, segments, 1)

		pieces := segments[0].Pieces
		require.False(t, hasDuplicates(pieces))

		// Remove second piece in the list and replace it with
		// a piece on the first node.
		// This way we can ensure that we use a valid piece num.
		removePiece := metabase.Piece{
			Number:      pieces[1].Number,
			StorageNode: pieces[1].StorageNode,
		}
		addPiece := metabase.Piece{
			Number:      pieces[1].Number,
			StorageNode: pieces[0].StorageNode,
		}

		// test no duplicates
		err = satellite.GracefulExit.Endpoint.UpdatePiecesCheckDuplicates(ctx, segments[0], metabase.Pieces{addPiece}, metabase.Pieces{removePiece}, true)
		require.True(t, metainfo.ErrNodeAlreadyExists.Has(err))
	})
}

func TestNodeAlreadyExited(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 4,
		Reconfigure: testplanet.Reconfigure{
			Satellite: func(log *zap.Logger, index int, config *satellite.Config) {
				config.GracefulExit.TimeBased = true
			},
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]

		// check that there are no exiting nodes.
		exitingNodes, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 0)

		exitingNode := planet.StorageNodes[0]

		simTime := time.Now()
		satellite.GracefulExit.Endpoint.SetNowFunc(func() time.Time { return simTime })
		doneTime := simTime.AddDate(0, 0, satellite.Config.GracefulExit.GracefulExitDurationInDays)

		// initiate GE
		response, err := callProcess(ctx, exitingNode, satellite)
		require.NoError(t, err)
		require.IsType(t, (*pb.SatelliteMessage_NotReady)(nil), response.GetMessage())

		// jump to when GE will be done
		simTime = doneTime.Add(time.Second)

		// should get ExitCompleted now
		response, err = callProcess(ctx, exitingNode, satellite)
		require.NoError(t, err)
		require.IsType(t, (*pb.SatelliteMessage_ExitCompleted)(nil), response.GetMessage())

		// now that the node has successfully exited, try doing it again! we expect to get the
		// ExitCompleted message again.
		response, err = callProcess(ctx, exitingNode, satellite)
		require.NoError(t, err)
		require.IsType(t, (*pb.SatelliteMessage_ExitCompleted)(nil), response.GetMessage())

		// check that node is not marked as exiting still
		exitingNodes, err = satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 0)
	})
}

func TestNodeSuspended(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 4,
		Reconfigure: testplanet.Reconfigure{
			Satellite: func(log *zap.Logger, index int, config *satellite.Config) {
				config.GracefulExit.TimeBased = true
			},
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]

		// check that there are no exiting nodes.
		exitingNodes, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 0)

		// mark a node as suspended
		exitingNode := planet.StorageNodes[0]
		err = satellite.Reputation.Service.TestSuspendNodeUnknownAudit(ctx, exitingNode.ID(), time.Now())
		require.NoError(t, err)

		// initiate GE
		response, err := callProcess(ctx, exitingNode, satellite)
		require.Error(t, err)
		require.ErrorContains(t, err, "node is suspended")
		require.Nil(t, response)
	})
}

func TestManyNodesGracefullyExiting(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 8,
		UplinkCount:      1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: testplanet.Combine(
				testplanet.ReconfigureRS(2, 3, 4, 4),
				func(log *zap.Logger, index int, config *satellite.Config) {
					config.GracefulExit.TimeBased = true
				},
			),
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]
		uplink := planet.Uplinks[0]

		satellite.RangedLoop.RangedLoop.Service.Loop.Stop()
		satellite.Repair.Repairer.Loop.Pause()

		// upload several objects; enough that we can reasonably expect every node to have several pieces
		const numObjects = 32
		objectData := make([][]byte, numObjects)
		for i := 0; i < numObjects; i++ {
			objectData[i] = testrand.Bytes(64 * memory.KiB)
			err := uplink.Upload(ctx, satellite, "testbucket", fmt.Sprintf("test/path/obj%d", i), objectData[i])
			require.NoError(t, err, i)
		}

		// Make half of the nodes initiate GE
		for i := 0; i < len(planet.StorageNodes)/2; i++ {
			response, err := callProcess(ctx, planet.StorageNodes[i], satellite)
			require.NoError(t, err, i)
			require.IsType(t, (*pb.SatelliteMessage_NotReady)(nil), response.GetMessage())
		}

		// run the satellite ranged loop to build the transfer queue.
		_, err := satellite.RangedLoop.RangedLoop.Service.RunOnce(ctx)
		require.NoError(t, err)

		// we expect ~78% of segments to be in the repair queue (the chance that a
		// segment still has at least 3 pieces in not-exiting nodes). but since things
		// will fluctuate, let's just expect half
		count, err := satellite.DB.RepairQueue().Count(ctx)
		require.NoError(t, err)
		require.GreaterOrEqual(t, count, numObjects/2)

		// perform the repairs, which should get every piece so that it will still be
		// reconstructable without the exiting nodes.
		satellite.Repair.Repairer.Loop.Restart()
		satellite.Repair.Repairer.Loop.TriggerWait()
		satellite.Repair.Repairer.Loop.Pause()
		satellite.Repair.Repairer.WaitForPendingRepairs()

		// turn off the exiting nodes entirely
		for i := 0; i < len(planet.StorageNodes)/2; i++ {
			err = planet.StopNodeAndUpdate(ctx, planet.StorageNodes[i])
			require.NoError(t, err)
		}

		// expect that we can retrieve and verify all objects
		for i, obj := range objectData {
			gotData, err := uplink.Download(ctx, satellite, "testbucket", fmt.Sprintf("test/path/obj%d", i))
			require.NoError(t, err, i)
			require.Equal(t, string(obj), string(gotData))
		}
	})
}

func TestNodeFailingGracefulExitWithLowOnlineScore(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 4,
		Reconfigure: testplanet.Reconfigure{
			Satellite: func(log *zap.Logger, index int, config *satellite.Config) {
				config.Reputation.AuditHistory.WindowSize = 24 * time.Hour
				config.Reputation.AuditHistory.TrackingPeriod = 3 * 24 * time.Hour
				config.Reputation.FlushInterval = 0
				config.GracefulExit.MinimumOnlineScore = 0.6
				config.GracefulExit.TimeBased = true
				config.GracefulExit.GracefulExitDurationInDays = 30
			},
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]
		exitingNode := planet.StorageNodes[0]

		exitingNode.GracefulExit.Chore.Loop.Pause()
		exitingNode.Contact.Chore.Pause(ctx)

		simTime := time.Now()
		satellite.GracefulExit.Endpoint.SetNowFunc(func() time.Time { return simTime })
		doneTime := simTime.AddDate(0, 0, satellite.Config.GracefulExit.GracefulExitDurationInDays)

		// initiate GE
		response, err := callProcess(ctx, exitingNode, satellite)
		require.NoError(t, err)
		require.IsType(t, (*pb.SatelliteMessage_NotReady)(nil), response.GetMessage())

		// set the audit history for that node to reflect a poor online score
		last := reputation.AuditSuccess
		for simTime.Before(doneTime) {
			// alternate between Success and Offline to get a ~50% score
			if last == reputation.AuditSuccess {
				last = reputation.AuditOffline
			} else {
				last = reputation.AuditSuccess
			}
			_, err := satellite.DB.Reputation().Update(ctx, reputation.UpdateRequest{
				NodeID:       exitingNode.ID(),
				AuditOutcome: last,
				Config:       satellite.Config.Reputation,
			}, simTime)
			require.NoError(t, err)

			// GE shouldn't fail until the end of the period, so node has a chance to get score back up
			response, err := callProcess(ctx, exitingNode, satellite)
			require.NoError(t, err)
			require.IsTypef(t, (*pb.SatelliteMessage_NotReady)(nil), response.GetMessage(), "simTime=%s, doneTime=%s", simTime, doneTime)

			simTime = simTime.Add(time.Hour)
		}
		err = satellite.Reputation.Service.TestFlushAllNodeInfo(ctx)
		require.NoError(t, err)

		simTime = doneTime.Add(time.Second)
		response, err = callProcess(ctx, exitingNode, satellite)
		require.NoError(t, err)
		msg := response.GetMessage()
		require.IsType(t, (*pb.SatelliteMessage_ExitFailed)(nil), msg)
		failure := msg.(*pb.SatelliteMessage_ExitFailed)

		// validate signature on failure message
		signee := signing.SigneeFromPeerIdentity(satellite.Identity.PeerIdentity())
		err = signing.VerifyExitFailed(ctx, signee, failure.ExitFailed)
		require.Equal(t, exitingNode.ID(), failure.ExitFailed.NodeId)
		// truncate to micros since the Failed time has gone through the database
		expectedFailTime := simTime.Truncate(time.Microsecond)
		require.Falsef(t, failure.ExitFailed.Failed.Before(expectedFailTime),
			"failure time should have been at or after %s: %s", simTime, failure.ExitFailed.Failed)
		require.Equal(t, satellite.ID(), failure.ExitFailed.SatelliteId)
		require.Equal(t, pb.ExitFailed_INACTIVE_TIMEFRAME_EXCEEDED, failure.ExitFailed.Reason)
		require.NoError(t, err)
	})
}

func TestSuspendedNodesFailGracefulExit(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 4,
		Reconfigure: testplanet.Reconfigure{
			Satellite: func(log *zap.Logger, index int, config *satellite.Config) {
				config.Reputation.FlushInterval = 0
				config.GracefulExit.TimeBased = true
			},
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]
		exitingNode := planet.StorageNodes[0]

		simTime := time.Now()
		satellite.GracefulExit.Endpoint.SetNowFunc(func() time.Time { return simTime })
		doneTime := simTime.AddDate(0, 0, satellite.Config.GracefulExit.GracefulExitDurationInDays)

		// initiate GE
		response, err := callProcess(ctx, exitingNode, satellite)
		require.NoError(t, err)
		require.IsType(t, (*pb.SatelliteMessage_NotReady)(nil), response.GetMessage())

		// suspend the node
		err = satellite.Reputation.Service.TestSuspendNodeUnknownAudit(ctx, exitingNode.ID(), simTime)
		require.NoError(t, err)

		// expect failure when the time is up
		simTime = doneTime.Add(time.Second)

		response, err = callProcess(ctx, exitingNode, satellite)
		require.NoError(t, err)
		msg := response.GetMessage()
		require.IsType(t, (*pb.SatelliteMessage_ExitFailed)(nil), msg)
		failure := msg.(*pb.SatelliteMessage_ExitFailed)
		require.Equal(t, pb.ExitFailed_OVERALL_FAILURE_PERCENTAGE_EXCEEDED, failure.ExitFailed.Reason)
	})
}

func TestNodeStartedGracefulExitBeforeUpgradeToTimeBased(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 8,
		UplinkCount:      1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: testplanet.Combine(
				testplanet.ReconfigureRS(4, 5, 6, 6),
				func(log *zap.Logger, index int, config *satellite.Config) {
					config.Reputation.FlushInterval = 0
					config.GracefulExit.GracefulExitDurationInDays = 1
					config.GracefulExit.EndpointBatchSize = 1
					// starts off as false; we will recreate the API server halfway through with true
					config.GracefulExit.TimeBased = false
				},
			),
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]
		exitingNode := planet.StorageNodes[0]
		uplink := planet.Uplinks[0]

		simTime := time.Now()
		satellite.GracefulExit.Endpoint.SetNowFunc(func() time.Time { return simTime })

		// Get some data on the node.
		for i := 0; i < 10; i++ {
			path := fmt.Sprintf("test/path/%d", i)
			err := uplink.Upload(ctx, satellite, "test-bucket", path, testrand.Bytes(24*memory.KiB))
			require.NoError(t, err, path)
		}
		_, piecesContentSize, err := exitingNode.Storage2.BlobsCache.SpaceUsedBySatellite(ctx, satellite.ID())
		require.NoError(t, err)
		require.NotZero(t, piecesContentSize)

		// Pause the GE chore on the node so it doesn't start transferring pieces too quickly
		// and finish before we even do the test.
		exitingNode.GracefulExit.Chore.Loop.Pause()

		// Initiate GE on the node (this affects only the storagenode DB).
		err = exitingNode.DB.Satellites().InitiateGracefulExit(ctx, satellite.ID(), time.Now(), piecesContentSize)
		require.NoError(t, err)

		// Allow the node to start talking to the satellite about its graceful exit. This will make
		// the Process() call to the API, but the transfer queue won't get built yet because the
		// chore is stopped. The worker should exit almost immediately after determining that the
		// server doesn't have any pieces for it to transfer yet.
		exitingNode.GracefulExit.Chore.Loop.TriggerWait()
		err = exitingNode.GracefulExit.Chore.TestWaitForNoWorkers(ctx)
		require.NoError(t, err)

		// Run the ranged loop once to build the transfer queue.
		_, err = satellite.RangedLoop.RangedLoop.Service.RunOnce(ctx)
		require.NoError(t, err)

		// Ensure there are elements in the queue.
		queue, err := satellite.DB.GracefulExit().GetIncomplete(ctx, exitingNode.ID(), 1, 0)
		require.NoError(t, err)
		require.Len(t, queue, 1)

		// Change the setting of TimeBased on the API server. It shouldn't be necessary
		// to restart the subsystem because TimeBased does not affect how an API server
		// object is initialized. It _does_ affect whether the gracefulexit ranged loop
		// observer gets registered, but having the non-time-based ranged loop observer
		// registered should not break anything.
		satellite.GracefulExit.Endpoint.TestSetTimeBased(true)

		// Allow the node to proceed with graceful exit. Again, the worker should exit almost
		// immediately after determining that the server doesn't have any pieces for it to
		// transfer yet.
		exitingNode.GracefulExit.Chore.Loop.TriggerWait()
		err = exitingNode.GracefulExit.Chore.TestWaitForNoWorkers(ctx)
		require.NoError(t, err)

		// It shouldn't be done yet, because graceful exit is time-based now
		exits, err := exitingNode.DB.Satellites().ListGracefulExits(ctx)
		require.NoError(t, err)
		require.Len(t, exits, 1)
		nodeInfo, err := satellite.Overlay.Service.Get(ctx, exitingNode.ID())
		require.NoError(t, err)
		require.False(t, nodeInfo.ExitStatus.ExitSuccess)
		require.NotNil(t, nodeInfo.ExitStatus.ExitInitiatedAt)
		require.Nil(t, nodeInfo.ExitStatus.ExitFinishedAt)

		// Move time forward 2 days so that the node should be done
		simTime = simTime.Add(48 * time.Hour)

		// Have the node check in to the graceful exit endpoint again. This time, the API server
		// should respond that the node has completed GE.
		exitingNode.GracefulExit.Chore.Loop.TriggerWait()
		err = exitingNode.GracefulExit.Chore.TestWaitForNoWorkers(ctx)
		require.NoError(t, err)

		// The node should know it's done
		exits, err = exitingNode.DB.Satellites().ListGracefulExits(ctx)
		require.NoError(t, err)
		require.Len(t, exits, 1)
		require.Equal(t, satellites.ExitSucceeded, exits[0].Status)
		require.NotNil(t, exits[0].FinishedAt)
		// and the satellite should know it's done.
		nodeInfo, err = satellite.Overlay.Service.Get(ctx, exitingNode.ID())
		require.NoError(t, err)
		require.True(t, nodeInfo.ExitStatus.ExitSuccess)
		require.NotNil(t, nodeInfo.ExitStatus.ExitFinishedAt)
	})
}

func hasDuplicates(pieces metabase.Pieces) bool {
	nodePieceCounts := make(map[storj.NodeID]int)
	for _, piece := range pieces {
		nodePieceCounts[piece.StorageNode]++
	}

	for _, count := range nodePieceCounts {
		if count > 1 {
			return true
		}
	}

	return false
}

func callProcess(ctx *testcontext.Context, exitingNode *testplanet.StorageNode, satellite *testplanet.Satellite) (*pb.SatelliteMessage, error) {
	conn, err := exitingNode.Dialer.DialNodeURL(ctx, satellite.NodeURL())
	if err != nil {
		return nil, err
	}
	defer ctx.Check(conn.Close)

	client := pb.NewDRPCSatelliteGracefulExitClient(conn)

	c, err := client.Process(ctx)
	if err != nil {
		return nil, err
	}
	defer ctx.Check(c.CloseSend)

	return c.Recv()
}
