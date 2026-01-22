package database

import (
	"context"
	"database/sql"
	"go-ecommerce-backend-api/global"

	"github.com/jmoiron/sqlx"
)

type txKey struct {
}

type DBTX interface {
	sqlx.ExtContext
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type Transactor interface {
	WithTransaction(ctx context.Context, fn func(txCtx context.Context) error) error
}

type transactorImpl struct {
	db *sqlx.DB
}

func NewTransactor() Transactor {
	return &transactorImpl{
		db: global.Pdbx,
	}
}

func (t *transactorImpl) WithTransaction(ctx context.Context, fn func(txCtx context.Context) error) error {
	tx, err := t.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	txCtx := context.WithValue(ctx, txKey{}, tx)
	err = fn(txCtx)
	return err
}

func GetExecutor(ctx context.Context, db *sqlx.DB) DBTX {
	if tx, ok := ctx.Value(txKey{}).(*sqlx.Tx); ok {
		return tx
	}
	return db
}

