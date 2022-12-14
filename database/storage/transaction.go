package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lukmanlukmin/golib/database/connection"
	cErr "github.com/lukmanlukmin/golib/errors"
	"github.com/lukmanlukmin/golib/log"
)

const txKey ctxTrxKey = "IsTransaction"

type (
	ctxTrxKey string

	BaseStorage struct {
		Storage *connection.Store
	}

	SQLExec interface {
		sqlx.Execer
		sqlx.ExecerContext
		NamedExec(query string, arg interface{}) (sql.Result, error)
		NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	}
	SQLQuery interface {
		sqlx.Queryer
		GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
		SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
		PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	}

	SQLQueryExec interface {
		SQLExec
		SQLQuery
		Rebind(query string) string
	}
)

func NewBaseStorage(store *connection.Store) *BaseStorage {
	return &BaseStorage{
		Storage: store,
	}
}

type TransactionFunc func(ctx context.Context) error

// wrapper function to handle transaction process
func (r *BaseStorage) WithTransaction(ctx context.Context, fn TransactionFunc) error {

	tx := getTxFromContext(ctx)
	// if no available transaction, set to default transaction
	if tx == nil {
		newCtx, err := r.setDefaultTrx(ctx)
		if err != nil {
			return err
		}
		ctx = newCtx
	}

	defer func(trx *sqlx.Tx) {
		if err := trx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.WithError(err).Errorln("failed on rollback transaction")
		}
	}(tx)

	if err := fn(ctx); err != nil {
		return cErr.CustomError{
			Message: fmt.Sprintf("%s: %v", ErrSQLTransactionFailedTrx.Error(), err),
		}
	}

	if err := tx.Commit(); err != nil {
		return cErr.CustomError{
			Message: fmt.Sprintf("%s: %v", ErrSQLTransactionFailedCommit.Error(), err),
		}
	}
	return nil
}

// set transaction to context
func (r *BaseStorage) SetTxToContext(ctx context.Context, tx *sqlx.Tx) context.Context {
	ctx = context.WithValue(ctx, txKey, tx)
	return ctx
}

// set default transaction config to context
func (r *BaseStorage) setDefaultTrx(ctx context.Context) (context.Context, error) {
	tx, err := r.Storage.Master.BeginTxx(ctx, &sql.TxOptions{})
	return r.SetTxToContext(ctx, tx), err
}

// get transaction from context. return null if not available
func getTxFromContext(ctx context.Context) *sqlx.Tx {
	if tx, ok := ctx.Value(txKey).(*sqlx.Tx); ok {
		return tx
	}
	return nil
}
