package storage

import "errors"

// transactional SQL error
var (
	ErrSQLTransactionNoTrx        = errors.New("error on getting transaction context. no transaction available")
	ErrSQLTransactionFailedTrx    = errors.New("error on performing transactional request")
	ErrSQLTransactionFailedCommit = errors.New("failed to commit transaction")
)
