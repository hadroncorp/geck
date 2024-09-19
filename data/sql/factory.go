package sql

import (
	"context"
	"database/sql"

	"github.com/hadroncorp/geck/data/persistence"
)

type TransactionContextFactory struct {
	Client Client
	Config ConfigTransactionFactory
}

var _ persistence.TransactionContextFactory = (*TransactionContextFactory)(nil)

func NewTransactionContextFactory(client Client, cfg ConfigTransactionFactory) TransactionContextFactory {
	return TransactionContextFactory{
		Client: client,
		Config: cfg,
	}
}

func (t TransactionContextFactory) NewContext(parent context.Context) (context.Context, error) {
	_, err := persistence.GetTxFromContext(parent)
	if err == nil {
		return parent, nil // re-use ctx
	}

	tx, err := t.Client.BeginTx(parent, &sql.TxOptions{
		Isolation: sql.IsolationLevel(t.Config.IsolationLevel),
		ReadOnly:  t.Config.ReadOnly,
	})
	if err != nil {
		return nil, err
	}
	return context.WithValue(parent, persistence.TransactionContextKey, Transaction{Tx: tx}), nil
}
