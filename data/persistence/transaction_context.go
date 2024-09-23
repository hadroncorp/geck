package persistence

import (
	"context"
	"errors"

	"github.com/hadroncorp/geck/systemerror"
)

type transactionContextType string

const TransactionContextKey transactionContextType = "geck.persistence.tx"

type Transaction interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

func GetTxFromContext(ctx context.Context) (Transaction, error) {
	tx, ok := ctx.Value(TransactionContextKey).(Transaction)
	if !ok {
		return nil, ErrTxContextNotFound
	}
	return tx, nil
}

func CloseTransaction(ctx context.Context, srcErr error) error {
	tx, err := GetTxFromContext(ctx)
	if err != nil && errors.Is(err, ErrTxContextNotFound) {
		return srcErr
	} else if err != nil {
		return errors.Join(srcErr, err)
	}
	if srcErr == nil {
		return tx.Commit(ctx)
	}
	errRollback := tx.Rollback(ctx)
	if srcErrors, ok := srcErr.(systemerror.Container); ok {
		return errors.Join(append(srcErrors.Unwrap(), errRollback)...)
	}
	return errors.Join(srcErr, errRollback)
}
