// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package attribution_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"storj.io/common/pb"
	"storj.io/common/testcontext"
	"storj.io/common/testrand"
	"storj.io/common/uuid"
	"storj.io/storj/satellite"
	"storj.io/storj/satellite/accounting"
	"storj.io/storj/satellite/attribution"
	"storj.io/storj/satellite/satellitedb/satellitedbtest"
)

const (
	remoteSize int64 = 5368709120
	inlineSize int64 = 1024
	egressSize int64 = 2048
)

type AttributionTestData struct {
	name       string
	partnerID  uuid.UUID
	userAgent  []byte
	projectID  uuid.UUID
	bucketName []byte
	bucketID   []byte

	start        time.Time
	end          time.Time
	dataInterval time.Time
	bwStart      time.Time

	hours       int64
	hoursOfData int
	padding     int

	remoteSize int64
	inlineSize int64
	egressSize int64

	dataCounter        int
	expectedTotalBytes int64
	expectedEgress     int64
}

func (testData *AttributionTestData) init() {
	testData.bucketID = []byte(testData.projectID.String() + "/" + string(testData.bucketName))
	testData.hours = int64(testData.end.Sub(testData.start).Hours())
	testData.hoursOfData = int(testData.hours) + (testData.padding * 2)
	testData.dataInterval = time.Date(testData.start.Year(), testData.start.Month(), testData.start.Day(), -testData.padding, -1, 0, 0, testData.start.Location())
	testData.bwStart = time.Date(testData.start.Year(), testData.start.Month(), testData.start.Day(), -testData.padding, 0, 0, 0, testData.start.Location())
}

func TestDB(t *testing.T) {
	satellitedbtest.Run(t, func(ctx *testcontext.Context, t *testing.T, db satellite.DB) {
		attributionDB := db.Attribution()
		project1, project2 := testrand.UUID(), testrand.UUID()
		partner1, partner2 := testrand.UUID(), testrand.UUID()
		agent1, agent2 := []byte("agent1"), []byte("agent2")

		infos := []*attribution.Info{
			{project1, []byte("alpha"), partner1, agent1, time.Time{}},
			{project1, []byte("beta"), partner2, agent2, time.Time{}},
			{project2, []byte("alpha"), partner2, agent2, time.Time{}},
			{project2, []byte("beta"), partner1, agent1, time.Time{}},
		}

		for _, info := range infos {
			got, err := attributionDB.Insert(ctx, info)
			require.NoError(t, err)

			got.CreatedAt = time.Time{}
			assert.Equal(t, info, got)
		}

		for _, info := range infos {
			got, err := attributionDB.Get(ctx, info.ProjectID, info.BucketName)
			require.NoError(t, err)
			assert.Equal(t, info.PartnerID, got.PartnerID)
			assert.Equal(t, info.UserAgent, got.UserAgent)
		}
	})
}

func TestQueryAttribution(t *testing.T) {
	satellitedbtest.Run(t, func(ctx *testcontext.Context, t *testing.T, db satellite.DB) {
		now := time.Now()

		projectID := testrand.UUID()
		partnerID := testrand.UUID()
		userAgent := []byte("agent1")
		alphaBucket := []byte("alpha")
		betaBucket := []byte("beta")
		testData := []AttributionTestData{
			{
				name:       "new partnerID, userAgent, projectID, alpha",
				partnerID:  testrand.UUID(),
				userAgent:  []byte("agent2"),
				projectID:  projectID,
				bucketName: alphaBucket,

				remoteSize: remoteSize,
				inlineSize: inlineSize,
				egressSize: egressSize,

				start:   time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()),
				end:     time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()),
				padding: 2,
			},
			{
				name:       "partnerID, userAgent, new projectID, alpha",
				partnerID:  partnerID,
				userAgent:  userAgent,
				projectID:  testrand.UUID(),
				bucketName: alphaBucket,

				remoteSize: remoteSize / 2,
				inlineSize: inlineSize / 2,
				egressSize: egressSize / 2,

				start:   time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()),
				end:     time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()),
				padding: 2,
			},
			{
				name:       "new partnerID, userAgent, projectID, beta",
				partnerID:  testrand.UUID(),
				userAgent:  []byte("agent3"),
				projectID:  projectID,
				bucketName: betaBucket,

				remoteSize: remoteSize / 3,
				inlineSize: inlineSize / 3,
				egressSize: egressSize / 3,

				start:   time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()),
				end:     time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()),
				padding: 2,
			},
			{
				name:       "partnerID, userAgent new projectID, beta",
				partnerID:  partnerID,
				userAgent:  userAgent,
				projectID:  testrand.UUID(),
				bucketName: betaBucket,

				remoteSize: remoteSize / 4,
				inlineSize: inlineSize / 4,
				egressSize: egressSize / 4,

				start:   time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()),
				end:     time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()),
				padding: 2,
			}}

		for _, td := range testData {
			td := td
			td.init()
			info := attribution.Info{td.projectID, td.bucketName, td.partnerID, td.userAgent, time.Time{}}
			_, err := db.Attribution().Insert(ctx, &info)
			require.NoError(t, err)

			for i := 0; i < td.hoursOfData; i++ {
				createData(ctx, t, db, &td)
			}

			verifyData(ctx, t, db.Attribution(), &td)
		}
	})
}

func TestQueryAllAttribution(t *testing.T) {
	satellitedbtest.Run(t, func(ctx *testcontext.Context, t *testing.T, db satellite.DB) {
		now := time.Now()

		projectID := testrand.UUID()
		partnerID := testrand.UUID()
		userAgent := []byte("agent1")
		alphaBucket := []byte("alpha")
		betaBucket := []byte("beta")
		testData := []AttributionTestData{
			{
				name:       "new partnerID, userAgent, projectID, alpha",
				partnerID:  testrand.UUID(),
				userAgent:  []byte("agent2"),
				projectID:  projectID,
				bucketName: alphaBucket,

				remoteSize: remoteSize,
				inlineSize: inlineSize,
				egressSize: egressSize,

				start:   time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()),
				end:     time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()),
				padding: 2,
			},
			{
				name:       "partnerID, userAgent, new projectID, alpha",
				partnerID:  partnerID,
				userAgent:  userAgent,
				projectID:  testrand.UUID(),
				bucketName: alphaBucket,

				remoteSize: remoteSize / 2,
				inlineSize: inlineSize / 2,
				egressSize: egressSize / 2,

				start:   time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()),
				end:     time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()),
				padding: 2,
			},
			{
				name:       "new partnerID, userAgent, projectID, beta",
				partnerID:  testrand.UUID(),
				userAgent:  []byte("agent3"),
				projectID:  projectID,
				bucketName: betaBucket,

				remoteSize: remoteSize / 3,
				inlineSize: inlineSize / 3,
				egressSize: egressSize / 3,

				start:   time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()),
				end:     time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()),
				padding: 2,
			},
			{
				name:       "partnerID, userAgent new projectID, beta",
				partnerID:  partnerID,
				userAgent:  userAgent,
				projectID:  testrand.UUID(),
				bucketName: betaBucket,

				remoteSize: remoteSize / 4,
				inlineSize: inlineSize / 4,
				egressSize: egressSize / 4,

				start:   time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()),
				end:     time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()),
				padding: 2,
			}}

		for i, td := range testData {
			td := td
			td.init()

			info := attribution.Info{td.projectID, td.bucketName, td.partnerID, td.userAgent, time.Time{}}
			_, err := db.Attribution().Insert(ctx, &info)
			require.NoError(t, err)

			for i := 0; i < td.hoursOfData; i++ {
				createData(ctx, t, db, &td)
			}
			testData[i] = td
		}
		verifyAllData(ctx, t, db.Attribution(), testData)
	})
}

func verifyData(ctx *testcontext.Context, t *testing.T, attributionDB attribution.DB, testData *AttributionTestData) {
	results, err := attributionDB.QueryAttribution(ctx, testData.partnerID, testData.userAgent, testData.start, testData.end)
	require.NoError(t, err)
	require.NotEqual(t, 0, len(results), "Results must not be empty.")
	count := 0
	for _, r := range results {
		projectID, err := uuid.FromBytes(r.ProjectID)
		require.NoError(t, err)

		// The query returns results by partnerID, so we need to filter out by projectID
		if projectID != testData.projectID {
			continue
		}
		count++

		assert.Equal(t, testData.partnerID[:], r.PartnerID, testData.name)
		assert.Equal(t, testData.userAgent, r.UserAgent, testData.name)
		assert.Equal(t, testData.projectID[:], r.ProjectID, testData.name)
		assert.Equal(t, testData.bucketName, r.BucketName, testData.name)
		assert.Equal(t, float64(testData.expectedTotalBytes/testData.hours), r.TotalBytesPerHour, testData.name)
		assert.Equal(t, testData.expectedEgress, r.EgressData, testData.name)
	}
	require.NotEqual(t, 0, count, "Results were returned, but did not match all of the the projectIDs.")
}

func verifyAllData(ctx *testcontext.Context, t *testing.T, attributionDB attribution.DB, testData []AttributionTestData) {
	results, err := attributionDB.QueryAllAttribution(ctx, testData[0].start, testData[0].end)
	require.NoError(t, err)
	require.NotEqual(t, 0, len(results), "Results must not be empty.")
	count := 0
	for _, tt := range testData {
		for _, r := range results {
			projectID, err := uuid.FromBytes(r.ProjectID)
			require.NoError(t, err)

			// The query returns results by partnerID, so we need to filter out by projectID
			if projectID != tt.projectID || !bytes.Equal(r.BucketName, tt.bucketName) {
				continue
			}
			count++

			assert.Equal(t, tt.partnerID[:], r.PartnerID, tt.name)
			assert.Equal(t, tt.userAgent, r.UserAgent, tt.name)
			assert.Equal(t, tt.projectID[:], r.ProjectID, tt.name)
			assert.Equal(t, tt.bucketName, r.BucketName, tt.name)
			assert.Equal(t, float64(tt.expectedTotalBytes/tt.hours), r.TotalBytesPerHour, tt.name)
			assert.Equal(t, tt.expectedEgress, r.EgressData, tt.name)
		}
	}

	require.Equal(t, len(testData), len(results))
	require.NotEqual(t, 0, count, "Results were returned, but did not match all of the the projectIDs.")
}

func createData(ctx *testcontext.Context, t *testing.T, db satellite.DB, testData *AttributionTestData) {
	projectAccoutingDB := db.ProjectAccounting()
	orderDB := db.Orders()

	// split the expected egress size into two separate bucket_bandwidth_rollup rows to test attribution query summation
	err := orderDB.UpdateBucketBandwidthSettle(ctx, testData.projectID, testData.bucketName, pb.PieceAction_GET, testData.egressSize/2, 0, testData.bwStart)
	require.NoError(t, err)

	err = orderDB.UpdateBucketBandwidthSettle(ctx, testData.projectID, testData.bucketName, pb.PieceAction_GET, testData.egressSize/2, 0, testData.bwStart.Add(2*time.Hour))
	require.NoError(t, err)

	// Only GET should be counted. So this should not effect results
	err = orderDB.UpdateBucketBandwidthSettle(ctx, testData.projectID, testData.bucketName, pb.PieceAction_GET_AUDIT, testData.egressSize, 0, testData.bwStart)
	require.NoError(t, err)

	testData.bwStart = testData.bwStart.Add(time.Hour)

	testData.dataCounter++
	testData.dataInterval = testData.dataInterval.Add(time.Minute * 30)
	_, err = createTallyData(ctx, projectAccoutingDB, testData.projectID, testData.bucketName, testData.remoteSize*int64(testData.dataCounter), testData.inlineSize*int64(testData.dataCounter), testData.dataInterval)
	require.NoError(t, err)

	testData.dataCounter++
	testData.dataInterval = testData.dataInterval.Add(time.Minute * 30)
	tally, err := createTallyData(ctx, projectAccoutingDB, testData.projectID, testData.bucketName, testData.remoteSize*int64(testData.dataCounter), testData.inlineSize*int64(testData.dataCounter), testData.dataInterval)
	require.NoError(t, err)

	if (testData.dataInterval.After(testData.start) || testData.dataInterval.Equal(testData.start)) && testData.dataInterval.Before(testData.end) {
		testData.expectedTotalBytes += tally.TotalBytes
		testData.expectedEgress += testData.egressSize
	}
}

func createTallyData(ctx *testcontext.Context, projectAccoutingDB accounting.ProjectAccounting, projectID uuid.UUID, bucket []byte, remote int64, inline int64, dataIntervalStart time.Time) (_ accounting.BucketStorageTally, err error) {
	tally := accounting.BucketStorageTally{
		BucketName:    string(bucket),
		ProjectID:     projectID,
		IntervalStart: dataIntervalStart,

		ObjectCount: 0,

		TotalSegmentCount: 0,

		TotalBytes:   inline + remote,
		MetadataSize: 0,
	}
	err = projectAccoutingDB.CreateStorageTally(ctx, tally)
	if err != nil {
		return accounting.BucketStorageTally{}, err
	}
	return tally, nil
}
