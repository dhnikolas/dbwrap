package dbwrap

import (
	"context"
	"database/sql"

	"github.com/dhnikolas/dbwrap/errors"
)

type txk string

const txKey txk = "wrappedTx"

func (w *Wrapper) WithTransaction(ctx context.Context, options *sql.TxOptions, fn func(ctx context.Context) error) error {
	if _, ok := ctx.Value(txKey).(*sql.Tx); ok {
		// Транзакция уже начата
		return fn(ctx)
	}
	tx, err := w.db.BeginTx(ctx, options)
	if err != nil {
		return errors.NewBeginTransactionErr(err)
	}
	ctx = context.WithValue(ctx, txKey, tx)

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(ctx)
	return err
}
