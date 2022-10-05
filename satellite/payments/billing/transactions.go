// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package billing

import (
	"context"
	"time"

	"github.com/zeebo/errs"

	"storj.io/common/currency"
	"storj.io/common/uuid"
)

// TransactionStatus indicates transaction status.
type TransactionStatus string

// ErrInsufficientFunds represents err when a user balance is too low for some transaction.
var ErrInsufficientFunds = errs.New("Insufficient funds for this transaction")

// ErrNoWallet represents err when there is no wallet in the DB.
var ErrNoWallet = errs.New("no wallet in the database")

// ErrNoTransactions represents err when there is no billing transactions in the DB.
var ErrNoTransactions = errs.New("no transactions in the database")

const (
	// TransactionStatusPending indicates that status of this transaction is pending.
	TransactionStatusPending = "pending"
	// TransactionStatusCancelled indicates that status of this transaction is cancelled.
	TransactionStatusCancelled = "cancelled"
	// TransactionStatusCompleted indicates that status of this transaction is complete.
	TransactionStatusCompleted = "complete"
)

// TransactionType indicates transaction type.
type TransactionType string

const (
	// TransactionTypeCredit indicates that type of this transaction is credit.
	TransactionTypeCredit = "credit"
	// TransactionTypeDebit indicates that type of this transaction is debit.
	TransactionTypeDebit = "debit"
	// TransactionTypeUnknown indicates that type of this transaction is unknown.
	TransactionTypeUnknown = "unknown"
)

// TransactionsDB is an interface which defines functionality
// of DB which stores billing transactions.
//
// architecture: Database
type TransactionsDB interface {
	// Insert inserts the provided transaction.
	Insert(ctx context.Context, tx Transaction) (txID int64, err error)
	// UpdateStatus updates the status of the transaction.
	UpdateStatus(ctx context.Context, txID int64, status TransactionStatus) error
	// UpdateMetadata updates the metadata of the transaction.
	UpdateMetadata(ctx context.Context, txID int64, metadata []byte) error
	// LastTransaction returns the timestamp and metadata of the last known transaction for given source and type.
	LastTransaction(ctx context.Context, txSource string, txType TransactionType) (time.Time, []byte, error)
	// List returns all transactions for the specified user.
	List(ctx context.Context, userID uuid.UUID) ([]Transaction, error)
	// GetBalance returns the current usable balance for the specified user.
	GetBalance(ctx context.Context, userID uuid.UUID) (currency.Amount, error)
}

// PaymentType is an interface which defines functionality required for all billing payment types. Payment types can
// include but are not limited to Bitcoin, Ether, credit or debit card, ACH transfer, or even physical transfer of live
// goats. In each case, a source, type, and method to get new transactions must be defined by the service, though
// metadata specific to each payment type is also supported (i.e. goat hair type).
type PaymentType interface {
	// Source the source of the payment
	Source() string
	// Type the type of the payment
	Type() TransactionType
	// GetNewTransactions returns new transactions that occurred after the provided last transaction received.
	GetNewTransactions(ctx context.Context, lastTransactionTime time.Time, metadata []byte) ([]Transaction, error)
}

// Transaction defines billing related transaction info that is stored in the DB.
type Transaction struct {
	ID          int64
	UserID      uuid.UUID
	Amount      currency.Amount
	Description string
	Source      string
	Status      TransactionStatus
	Type        TransactionType
	Metadata    []byte
	Timestamp   time.Time
	CreatedAt   time.Time
}
