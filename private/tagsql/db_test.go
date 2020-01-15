// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package tagsql_test

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"

	"storj.io/common/testcontext"
	"storj.io/storj/private/dbutil/cockroachutil"
	"storj.io/storj/private/dbutil/pgutil"
	"storj.io/storj/private/dbutil/pgutil/pgtest"
	"storj.io/storj/private/tagsql"
)

func run(t *testing.T, fn func(*testcontext.Context, *testing.T, *sql.DB, tagsql.ContextSupport)) {
	t.Helper()

	t.Run("mattn-sqlite3", func(t *testing.T) {
		ctx := testcontext.New(t)
		defer ctx.Cleanup()

		db, err := sql.Open("sqlite3", ":memory:")
		if err != nil {
			t.Fatal(err)
		}
		defer ctx.Check(db.Close)

		fn(ctx, t, db, tagsql.SupportBasic)
	})

	t.Run("lib-pq-postgres", func(t *testing.T) {
		ctx := testcontext.New(t)
		defer ctx.Cleanup()

		if *pgtest.ConnStr == "" {
			t.Skipf("postgresql flag missing, example:\n-postgres-test-db=%s", pgtest.DefaultConnStr)
		}

		db, err := pgutil.OpenUnique(ctx, *pgtest.ConnStr, "detect")
		require.NoError(t, err)
		defer ctx.Check(db.Close)

		fn(ctx, t, db.DB, tagsql.SupportNone)
	})

	t.Run("lib-pq-cockroach", func(t *testing.T) {
		ctx := testcontext.New(t)
		defer ctx.Cleanup()

		if *pgtest.CrdbConnStr == "" {
			t.Skipf("postgresql flag missing, example:\n-cockroach-test-db=%s", pgtest.DefaultCrdbConnStr)
		}

		db, err := cockroachutil.OpenUnique(ctx, *pgtest.CrdbConnStr, "detect")
		require.NoError(t, err)
		defer ctx.Check(db.Close)

		fn(ctx, t, db.DB, tagsql.SupportNone)
	})
}
