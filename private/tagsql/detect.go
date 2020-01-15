// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package tagsql

import (
	"database/sql"
	"reflect"

	"github.com/zeebo/errs"
)

// Currently lib/pq has known issues with contexts in general.
// For lib/pq context methods will be completely disabled.
//
// A few issues:
//   https://github.com/lib/pq/issues/874
//   https://github.com/lib/pq/issues/908
//   https://github.com/lib/pq/issues/731
//
// mattn/go-sqlite3 seems to work with contexts on the most part,
// except in transactions. For them, we need to disable.
//   https://github.com/mattn/go-sqlite3/issues/769
//
// Currently we don't have data on whether github.com/jackc/pgx supports them properly.

// ContextSupport returns the level of context support a driver has.
type ContextSupport byte

// Constants for defining context level support.
const (
	SupportBasic        ContextSupport = 1 << 0
	SupportTransactions ContextSupport = 1 << 1

	SupportNone ContextSupport = 0
	SupportAll  ContextSupport = SupportBasic | SupportTransactions
)

// Basic returns true when driver supports basic contexts.
func (v ContextSupport) Basic() bool {
	return v&SupportBasic == SupportBasic
}

// Transactions returns true when driver supports contexts inside transactions.
func (v ContextSupport) Transactions() bool {
	return v&SupportTransactions == SupportTransactions
}

// DetectContextSupport detects *sql.DB driver without importing the specific packages.
func DetectContextSupport(db *sql.DB) (ContextSupport, error) {
	// We're using reflect so we don't have to import these packages
	// into the binary.
	typ := reflect.TypeOf(db.Driver())
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	switch {
	case typ.PkgPath() == "github.com/mattn/go-sqlite3" && typ.Name() == "SQLiteDriver":
		return SupportBasic, nil
	case typ.PkgPath() == "github.com/lib/pq" && typ.Name() == "Driver" ||
		// internally uses lib/pq
		typ.PkgPath() == "storj.io/storj/private/dbutil/cockroachutil" && typ.Name() == "Driver":
		return SupportNone, nil
	default:
		return SupportNone, errs.New("sql driver %q %q unsupported", typ.PkgPath(), typ.Name())
	}
}
