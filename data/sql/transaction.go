package sql

import (
	"context"
	"database/sql"

	"github.com/hadroncorp/geck/data/persistence"
)

type Transaction struct {
	Tx *sql.Tx
}

var _ persistence.Transaction = (*Transaction)(nil)

func (t Transaction) Commit(_ context.Context) error {
	return t.Tx.Commit()
}

func (t Transaction) Rollback(_ context.Context) error {
	return t.Tx.Rollback()
}
