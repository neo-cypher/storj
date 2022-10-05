// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package metabase_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"storj.io/common/testcontext"
	"storj.io/common/uuid"
	"storj.io/storj/satellite/metabase"
	"storj.io/storj/satellite/metabase/metabasetest"
)

func TestListObjects(t *testing.T) {
	metabasetest.Run(t, func(ctx *testcontext.Context, t *testing.T, db *metabase.DB) {
		obj := metabasetest.RandObjectStream()

		t.Run("ProjectID missing", func(t *testing.T) {
			defer metabasetest.DeleteAll{}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts:     metabase.ListObjects{},
				ErrClass: &metabase.ErrInvalidRequest,
				ErrText:  "ProjectID missing",
			}.Check(ctx, t, db)

			metabasetest.Verify{}.Check(ctx, t, db)
		})

		t.Run("BucketName missing", func(t *testing.T) {
			defer metabasetest.DeleteAll{}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID: obj.ProjectID,
					Limit:     -1,
				},
				ErrClass: &metabase.ErrInvalidRequest,
				ErrText:  "BucketName missing",
			}.Check(ctx, t, db)

			metabasetest.Verify{}.Check(ctx, t, db)
		})

		t.Run("Invalid limit", func(t *testing.T) {
			defer metabasetest.DeleteAll{}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:  obj.ProjectID,
					BucketName: obj.BucketName,
					Limit:      -1,
				},
				ErrClass: &metabase.ErrInvalidRequest,
				ErrText:  "Invalid limit: -1",
			}.Check(ctx, t, db)

			metabasetest.Verify{}.Check(ctx, t, db)
		})

		t.Run("Status is invalid", func(t *testing.T) {
			defer metabasetest.DeleteAll{}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:  obj.ProjectID,
					BucketName: obj.BucketName,
					Limit:      3,
				},
				ErrClass: &metabase.ErrInvalidRequest,
				ErrText:  "Status is invalid",
			}.Check(ctx, t, db)

			metabasetest.Verify{}.Check(ctx, t, db)
		})

		t.Run("no objects", func(t *testing.T) {
			defer metabasetest.DeleteAll{}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:  obj.ProjectID,
					BucketName: obj.BucketName,
					Status:     metabase.Committed,
				},
				Result: metabase.ListObjectsResult{},
			}.Check(ctx, t, db)

			metabasetest.Verify{}.Check(ctx, t, db)
		})

		t.Run("less objects than limit", func(t *testing.T) {
			defer metabasetest.DeleteAll{}.Check(ctx, t, db)
			numberOfObjects := 3
			limit := 10
			expected := make([]metabase.ObjectEntry, numberOfObjects)
			objects := createObjects(ctx, t, db, numberOfObjects, uuid.UUID{1}, "mybucket")
			for i, obj := range objects {
				if delimiterIndex := strings.Index(string(obj.ObjectKey), string(metabase.Delimiter)); delimiterIndex > -1 {
					expected[i] = metabase.ObjectEntry{
						IsPrefix:  true,
						ObjectKey: obj.ObjectKey[:delimiterIndex+1],
						Status:    3,
					}
				} else {
					expected[i] = objectEntryFromRaw(obj)
				}
			}
			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             uuid.UUID{1},
					BucketName:            "mybucket",
					Recursive:             false,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,
					Limit:                 limit,
				},
				Result: metabase.ListObjectsResult{
					Objects: expected,
					More:    false,
				}}.Check(ctx, t, db)
			metabasetest.Verify{Objects: objects}.Check(ctx, t, db)
		})

		t.Run("more objects than limit", func(t *testing.T) {
			defer metabasetest.DeleteAll{}.Check(ctx, t, db)
			numberOfObjects := 10
			limit := 3
			expected := make([]metabase.ObjectEntry, limit)
			objects := createObjects(ctx, t, db, numberOfObjects, uuid.UUID{1}, "mybucket")
			for i, obj := range objects[:limit] {
				expected[i] = objectEntryFromRaw(obj)
			}
			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             uuid.UUID{1},
					BucketName:            "mybucket",
					Recursive:             true,
					Limit:                 limit,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,
				},
				Result: metabase.ListObjectsResult{
					Objects: expected,
					More:    true,
				}}.Check(ctx, t, db)
			metabasetest.Verify{Objects: objects}.Check(ctx, t, db)
		})

		t.Run("objects in one bucket in project with 2 buckets", func(t *testing.T) {
			defer metabasetest.DeleteAll{}.Check(ctx, t, db)
			numberOfObjectsPerBucket := 5

			expected := make([]metabase.ObjectEntry, numberOfObjectsPerBucket)

			objectsBucketA := createObjects(ctx, t, db, numberOfObjectsPerBucket, uuid.UUID{1}, "bucket-a")
			objectsBucketB := createObjects(ctx, t, db, numberOfObjectsPerBucket, uuid.UUID{1}, "bucket-b")

			for i, obj := range objectsBucketA {
				expected[i] = objectEntryFromRaw(obj)
			}

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             uuid.UUID{1},
					BucketName:            "bucket-a",
					Recursive:             true,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,
				},
				Result: metabase.ListObjectsResult{
					Objects: expected,
				}}.Check(ctx, t, db)

			metabasetest.Verify{
				Objects: append(objectsBucketA, objectsBucketB...),
			}.Check(ctx, t, db)
		})

		t.Run("objects in one bucket with same bucketName in another project", func(t *testing.T) {
			defer metabasetest.DeleteAll{}.Check(ctx, t, db)
			numberOfObjectsPerBucket := 5

			expected := make([]metabase.ObjectEntry, numberOfObjectsPerBucket)

			objectsProject1 := createObjects(ctx, t, db, numberOfObjectsPerBucket, uuid.UUID{1}, "mybucket")
			objectsProject2 := createObjects(ctx, t, db, numberOfObjectsPerBucket, uuid.UUID{2}, "mybucket")
			for i, obj := range objectsProject1 {
				expected[i] = objectEntryFromRaw(obj)
			}

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             uuid.UUID{1},
					BucketName:            "mybucket",
					Recursive:             true,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,
				},
				Result: metabase.ListObjectsResult{
					Objects: expected,
				}}.Check(ctx, t, db)

			metabasetest.Verify{
				Objects: append(objectsProject1, objectsProject2...),
			}.Check(ctx, t, db)
		})

		t.Run("recursive", func(t *testing.T) {
			defer metabasetest.DeleteAll{}.Check(ctx, t, db)
			projectID, bucketName := uuid.UUID{1}, "bucky"

			objects := createObjectsWithKeys(ctx, t, db, projectID, bucketName, []metabase.ObjectKey{
				"a",
				"b/1",
				"b/2",
				"b/3",
				"c",
				"c/",
				"c//",
				"c/1",
				"g",
			})

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Recursive:             true,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						objects["a"],
						objects["b/1"],
						objects["b/2"],
						objects["b/3"],
						objects["c"],
						objects["c/"],
						objects["c//"],
						objects["c/1"],
						objects["g"],
					},
				}}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Recursive:             true,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,

					Cursor: metabase.ListObjectsCursor{Key: "a", Version: 10},
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						objects["b/1"],
						objects["b/2"],
						objects["b/3"],
						objects["c"],
						objects["c/"],
						objects["c//"],
						objects["c/1"],
						objects["g"],
					}},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Recursive:             true,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,

					Cursor: metabase.ListObjectsCursor{Key: "b", Version: 0},
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						objects["b/1"],
						objects["b/2"],
						objects["b/3"],
						objects["c"],
						objects["c/"],
						objects["c//"],
						objects["c/1"],
						objects["g"],
					},
				}}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Recursive:             true,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,

					Prefix: "b/",
				},
				Result: metabase.ListObjectsResult{
					Objects: withoutPrefix("b/",
						objects["b/1"],
						objects["b/2"],
						objects["b/3"],
					),
				}}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Recursive:             true,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,

					Prefix: "b/",
					Cursor: metabase.ListObjectsCursor{Key: "a"},
				},
				Result: metabase.ListObjectsResult{
					Objects: withoutPrefix("b/",
						objects["b/1"],
						objects["b/2"],
						objects["b/3"],
					),
				}}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Recursive:             true,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,

					Prefix: "b/",
					Cursor: metabase.ListObjectsCursor{Key: "b/2", Version: -3},
				},
				Result: metabase.ListObjectsResult{
					Objects: withoutPrefix("b/",
						objects["b/2"],
						objects["b/3"],
					),
				}}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Recursive:             true,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,

					Prefix: "b/",
					Cursor: metabase.ListObjectsCursor{Key: "c/"},
				},
				Result: metabase.ListObjectsResult{},
			}.Check(ctx, t, db)
		})

		t.Run("non-recursive", func(t *testing.T) {
			defer metabasetest.DeleteAll{}.Check(ctx, t, db)
			projectID, bucketName := uuid.UUID{1}, "bucky"

			objects := createObjectsWithKeys(ctx, t, db, projectID, bucketName, []metabase.ObjectKey{
				"a",
				"b/1",
				"b/2",
				"b/3",
				"c",
				"c/",
				"c//",
				"c/1",
				"g",
			})

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						objects["a"],
						prefixEntry("b/", metabase.Committed),
						objects["c"],
						prefixEntry("c/", metabase.Committed),
						objects["g"],
					}},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,

					Cursor: metabase.ListObjectsCursor{Key: "a", Version: 10},
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						prefixEntry("b/", metabase.Committed),
						objects["c"],
						prefixEntry("c/", metabase.Committed),
						objects["g"],
					}},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,

					Cursor: metabase.ListObjectsCursor{Key: "b", Version: 0},
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						prefixEntry("b/", metabase.Committed),
						objects["c"],
						prefixEntry("c/", metabase.Committed),
						objects["g"],
					}},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,

					Prefix: "b/",
				},
				Result: metabase.ListObjectsResult{
					Objects: withoutPrefix("b/",
						objects["b/1"],
						objects["b/2"],
						objects["b/3"],
					)},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,

					Prefix: "b/",
					Cursor: metabase.ListObjectsCursor{Key: "a"},
				},
				Result: metabase.ListObjectsResult{
					Objects: withoutPrefix("b/",
						objects["b/1"],
						objects["b/2"],
						objects["b/3"],
					)},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,

					Prefix: "b/",
					Cursor: metabase.ListObjectsCursor{Key: "b/2", Version: -3},
				},
				Result: metabase.ListObjectsResult{
					Objects: withoutPrefix("b/",
						objects["b/2"],
						objects["b/3"],
					)},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,

					Prefix: "b/",
					Cursor: metabase.ListObjectsCursor{Key: "c/"},
				},
				Result: metabase.ListObjectsResult{},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,

					Prefix: "c/",
					Cursor: metabase.ListObjectsCursor{Key: "c/"},
				},
				Result: metabase.ListObjectsResult{
					Objects: withoutPrefix("c/",
						objects["c/"],
						prefixEntry("c//", metabase.Committed),
						objects["c/1"],
					)},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:             projectID,
					BucketName:            bucketName,
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,

					Prefix: "c//",
				},
				Result: metabase.ListObjectsResult{
					Objects: withoutPrefix("c//",
						objects["c//"],
					)},
			}.Check(ctx, t, db)
		})

	})
}

func TestListObjectsSkipCursor(t *testing.T) {
	metabasetest.Run(t, func(ctx *testcontext.Context, t *testing.T, db *metabase.DB) {
		projectID, bucketName := uuid.UUID{1}, "bucky"

		t.Run("no prefix", func(t *testing.T) {
			defer metabasetest.DeleteAll{}.Check(ctx, t, db)

			createObjectsWithKeys(ctx, t, db, projectID, bucketName, []metabase.ObjectKey{
				"08/test",
				"09/test",
				"10/test",
			})

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:  projectID,
					BucketName: bucketName,
					Recursive:  false,
					Prefix:     "",
					Cursor: metabase.ListObjectsCursor{
						Key:     metabase.ObjectKey("08/"),
						Version: 1,
					},
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						prefixEntry(metabase.ObjectKey("09/"), metabase.Committed),
						prefixEntry(metabase.ObjectKey("10/"), metabase.Committed),
					}},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:  projectID,
					BucketName: bucketName,
					Recursive:  false,
					Prefix:     "",
					Cursor: metabase.ListObjectsCursor{
						Key:     metabase.ObjectKey("08"),
						Version: 1,
					},
					Status:                metabase.Committed,
					IncludeSystemMetadata: true,
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						prefixEntry(metabase.ObjectKey("08/"), metabase.Committed),
						prefixEntry(metabase.ObjectKey("09/"), metabase.Committed),
						prefixEntry(metabase.ObjectKey("10/"), metabase.Committed),
					}},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:  projectID,
					BucketName: bucketName,
					Recursive:  false,
					Prefix:     "",
					Cursor: metabase.ListObjectsCursor{
						Key:     metabase.ObjectKey("08/a/x"),
						Version: 1,
					},
					Status:                metabase.Committed,
					IncludeSystemMetadata: true,
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						prefixEntry(metabase.ObjectKey("09/"), metabase.Committed),
						prefixEntry(metabase.ObjectKey("10/"), metabase.Committed),
					}},
			}.Check(ctx, t, db)
		})

		t.Run("prefix", func(t *testing.T) {
			defer metabasetest.DeleteAll{}.Check(ctx, t, db)

			createObjectsWithKeys(ctx, t, db, projectID, bucketName, []metabase.ObjectKey{
				"2017/05/08/test",
				"2017/05/09/test",
				"2017/05/10/test",
			})

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:  projectID,
					BucketName: bucketName,
					Recursive:  false,
					Prefix:     metabase.ObjectKey("2017/05/"),
					Cursor: metabase.ListObjectsCursor{
						Key:     metabase.ObjectKey("2017/05/08"),
						Version: 1,
					},
					Status:                metabase.Committed,
					IncludeSystemMetadata: true,
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						prefixEntry(metabase.ObjectKey("08/"), metabase.Committed),
						prefixEntry(metabase.ObjectKey("09/"), metabase.Committed),
						prefixEntry(metabase.ObjectKey("10/"), metabase.Committed),
					}},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:  projectID,
					BucketName: bucketName,
					Recursive:  false,
					Prefix:     metabase.ObjectKey("2017/05/"),
					Cursor: metabase.ListObjectsCursor{
						Key:     metabase.ObjectKey("2017/05/08/"),
						Version: 1,
					},
					Status:                metabase.Committed,
					IncludeSystemMetadata: true,
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						prefixEntry(metabase.ObjectKey("09/"), metabase.Committed),
						prefixEntry(metabase.ObjectKey("10/"), metabase.Committed),
					}},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:  projectID,
					BucketName: bucketName,
					Recursive:  false,
					Prefix:     metabase.ObjectKey("2017/05/"),
					Cursor: metabase.ListObjectsCursor{
						Key:     metabase.ObjectKey("2017/05/08/a/x"),
						Version: 1,
					},
					Status:                metabase.Committed,
					IncludeSystemMetadata: true,
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						prefixEntry(metabase.ObjectKey("09/"), metabase.Committed),
						prefixEntry(metabase.ObjectKey("10/"), metabase.Committed),
					}},
			}.Check(ctx, t, db)
		})

		t.Run("batch-size", func(t *testing.T) {
			defer metabasetest.DeleteAll{}.Check(ctx, t, db)

			afterDelimiter := metabase.ObjectKey(metabase.Delimiter + 1)

			objects := createObjectsWithKeys(ctx, t, db, projectID, bucketName, []metabase.ObjectKey{
				"2017/05/08",
				"2017/05/08/a",
				"2017/05/08/b",
				"2017/05/08/c",
				"2017/05/08/d",
				"2017/05/08/e",
				"2017/05/08" + afterDelimiter,
				"2017/05/09/a",
				"2017/05/09/b",
				"2017/05/09/c",
				"2017/05/09/d",
				"2017/05/09/e",
				"2017/05/10/a",
				"2017/05/10/b",
				"2017/05/10/c",
				"2017/05/10/d",
				"2017/05/10/e",
			})

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:  projectID,
					BucketName: bucketName,
					Recursive:  false,
					Prefix:     metabase.ObjectKey("2017/05/"),
					Cursor: metabase.ListObjectsCursor{
						Key:     metabase.ObjectKey("2017/05/08"),
						Version: 1,
					},
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						prefixEntry(metabase.ObjectKey("08/"), metabase.Committed),
						withoutPrefix1("2017/05/", objects["2017/05/08"+afterDelimiter]),
						prefixEntry(metabase.ObjectKey("09/"), metabase.Committed),
						prefixEntry(metabase.ObjectKey("10/"), metabase.Committed),
					}},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:  projectID,
					BucketName: bucketName,
					Recursive:  false,
					//BatchSize:  3,
					Prefix: metabase.ObjectKey("2017/05/"),
					Cursor: metabase.ListObjectsCursor{
						Key:     metabase.ObjectKey("2017/05/08/"),
						Version: 1,
					},
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						withoutPrefix1("2017/05/", objects["2017/05/08"+afterDelimiter]),
						prefixEntry(metabase.ObjectKey("09/"), metabase.Committed),
						prefixEntry(metabase.ObjectKey("10/"), metabase.Committed),
					}},
			}.Check(ctx, t, db)

			metabasetest.ListObjects{
				Opts: metabase.ListObjects{
					ProjectID:  projectID,
					BucketName: bucketName,
					Recursive:  false,
					Prefix:     metabase.ObjectKey("2017/05/"),
					Cursor: metabase.ListObjectsCursor{
						Key:     metabase.ObjectKey("2017/05/08/a/x"),
						Version: 1,
					},
					Status:                metabase.Committed,
					IncludeCustomMetadata: true,
					IncludeSystemMetadata: true,
				},
				Result: metabase.ListObjectsResult{
					Objects: []metabase.ObjectEntry{
						withoutPrefix1("2017/05/", objects["2017/05/08"+afterDelimiter]),
						prefixEntry(metabase.ObjectKey("09/"), metabase.Committed),
						prefixEntry(metabase.ObjectKey("10/"), metabase.Committed),
					}},
			}.Check(ctx, t, db)
		})
	})
}

func BenchmarkNonRecursiveObjectsListing(b *testing.B) {
	metabasetest.Bench(b, func(ctx *testcontext.Context, b *testing.B, db *metabase.DB) {
		baseObj := metabasetest.RandObjectStream()

		batchsize := 5
		for i := 0; i < 500; i++ {
			metabasetest.CreateObject(ctx, b, db, metabasetest.RandObjectStream(), 0)
		}

		for i := 0; i < 10; i++ {
			baseObj.ObjectKey = metabase.ObjectKey("foo/" + strconv.Itoa(i))
			metabasetest.CreateObject(ctx, b, db, baseObj, 0)

			baseObj.ObjectKey = metabase.ObjectKey("foo/prefixA/" + strconv.Itoa(i))
			metabasetest.CreateObject(ctx, b, db, baseObj, 0)

			baseObj.ObjectKey = metabase.ObjectKey("foo/prefixB/" + strconv.Itoa(i))
			metabasetest.CreateObject(ctx, b, db, baseObj, 0)
		}

		for i := 0; i < 50; i++ {
			baseObj.ObjectKey = metabase.ObjectKey("boo/foo" + strconv.Itoa(i) + "/object")
			metabasetest.CreateObject(ctx, b, db, baseObj, 0)
		}

		b.Run("listing no prefix", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := db.ListObjects(ctx, metabase.ListObjects{
					ProjectID:  baseObj.ProjectID,
					BucketName: baseObj.BucketName,
					Status:     metabase.Committed,
					Limit:      batchsize,
				})
				require.NoError(b, err)
				for result.More {
					result, err = db.ListObjects(ctx, metabase.ListObjects{
						ProjectID:  baseObj.ProjectID,
						BucketName: baseObj.BucketName,
						Cursor:     metabase.ListObjectsCursor{Key: result.Objects[len(result.Objects)-1].ObjectKey},
						Status:     metabase.Committed,
						Limit:      batchsize,
					})
					require.NoError(b, err)
				}
			}
		})

		b.Run("listing with prefix", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := db.ListObjects(ctx, metabase.ListObjects{
					ProjectID:  baseObj.ProjectID,
					BucketName: baseObj.BucketName,
					Prefix:     "foo/",
					Status:     metabase.Committed,
					Limit:      batchsize,
				})
				require.NoError(b, err)
				for result.More {
					cursorKey := "foo/" + result.Objects[len(result.Objects)-1].ObjectKey
					result, err = db.ListObjects(ctx, metabase.ListObjects{
						ProjectID:  baseObj.ProjectID,
						BucketName: baseObj.BucketName,
						Prefix:     "foo/",
						Cursor:     metabase.ListObjectsCursor{Key: cursorKey},
						Status:     metabase.Committed,
						Limit:      batchsize,
					})
					require.NoError(b, err)
				}
			}
		})

		b.Run("listing only prefix", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := db.ListObjects(ctx, metabase.ListObjects{
					ProjectID:  baseObj.ProjectID,
					BucketName: baseObj.BucketName,
					Prefix:     "boo/",
					Status:     metabase.Committed,
					Limit:      batchsize,
				})
				require.NoError(b, err)
				for result.More {
					cursorKey := "boo/" + result.Objects[len(result.Objects)-1].ObjectKey
					result, err = db.ListObjects(ctx, metabase.ListObjects{
						ProjectID:  baseObj.ProjectID,
						BucketName: baseObj.BucketName,
						Prefix:     "boo/",
						Cursor:     metabase.ListObjectsCursor{Key: cursorKey},
						Status:     metabase.Committed,
						Limit:      batchsize,
					})
					require.NoError(b, err)

				}
			}
		})
	})
}
