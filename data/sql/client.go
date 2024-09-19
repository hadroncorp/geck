package sql

import (
	"context"
	"database/sql"
	"database/sql/driver"

	"github.com/hadroncorp/geck/data/persistence"
	"github.com/hadroncorp/geck/observability/logging"
)

type Client interface {
	PingContext(ctx context.Context) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Driver() driver.Driver
}

type LoggerClient struct {
	Logger logging.Logger
	Next   Client
}

var _ Client = (*LoggerClient)(nil)

func NewLoggerClient(logger logging.Logger, db Client) LoggerClient {
	return LoggerClient{
		Logger: logger,
		Next:   db,
	}
}

func (l LoggerClient) PingContext(ctx context.Context) error {
	l.Logger.Debug().WriteWithCtx(ctx, "pinging database")
	return l.Next.PingContext(ctx)
}

func (l LoggerClient) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	l.Logger.Debug().WithField("statement", query).
		WithField("total_args", len(args)).
		WriteWithCtx(ctx, "executing query")
	return l.Next.ExecContext(ctx, query, args...)
}

func (l LoggerClient) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	l.Logger.Debug().WithField("statement", query).WriteWithCtx(ctx, "preparing query")
	return l.Next.PrepareContext(ctx, query)
}

func (l LoggerClient) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	l.Logger.Debug().WithField("statement", query).
		WithField("total_args", len(args)).
		WriteWithCtx(ctx, "querying statement")
	return l.Next.QueryContext(ctx, query, args...)
}

func (l LoggerClient) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	l.Logger.Debug().WithField("statement", query).
		WithField("total_args", len(args)).
		WriteWithCtx(ctx, "querying statement")
	return l.Next.QueryRowContext(ctx, query, args...)
}

func (l LoggerClient) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	isolationLvl := 0
	readOnly := false
	if opts != nil {
		isolationLvl = int(opts.Isolation)
		readOnly = opts.ReadOnly
	}
	l.Logger.Debug().
		WithField("isolation_level", isolationLvl).
		WithField("read_only", readOnly).
		WriteWithCtx(ctx, "starting transaction")
	return l.Next.BeginTx(ctx, opts)
}

func (l LoggerClient) Driver() driver.Driver {
	return l.Next.Driver()
}

type TransactionalClient struct {
	TransactionContextFactory TransactionContextFactory
	Logger                    logging.Logger
	Next                      Client
}

var _ Client = (*TransactionalClient)(nil)

func NewTransactionalClient(factory TransactionContextFactory, logger logging.Logger, next Client) TransactionalClient {
	return TransactionalClient{
		TransactionContextFactory: factory,
		Logger:                    logger,
		Next:                      next,
	}
}

func (t TransactionalClient) PingContext(ctx context.Context) error {
	return t.Next.PingContext(ctx)
}

func (t TransactionalClient) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	txRaw, err := persistence.GetTxFromContext(ctx)
	if err != nil {
		t.Logger.Warn().WithField("error", err).WriteWithCtx(ctx, "error getting transaction, using fallback client")
		return t.Next.ExecContext(ctx, query, args...)
	}

	tx, ok := txRaw.(Transaction)
	if !ok {
		t.Logger.WithError(err).WriteWithCtx(ctx, "error casting transaction structure, using fallback client")
		return t.Next.ExecContext(ctx, query, args...)
	}
	t.Logger.Debug().WriteWithCtx(ctx, "executing transactional query")
	return tx.Tx.ExecContext(ctx, query, args...)
}

func (t TransactionalClient) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	txRaw, err := persistence.GetTxFromContext(ctx)
	if err != nil {
		t.Logger.Warn().WithField("error", err).WriteWithCtx(ctx, "error getting transaction, using fallback client")
		return t.Next.PrepareContext(ctx, query)
	}

	tx, ok := txRaw.(Transaction)
	if !ok {
		t.Logger.WithError(err).WriteWithCtx(ctx, "error casting transaction structure, using fallback client")
		return t.Next.PrepareContext(ctx, query)
	}
	t.Logger.Debug().WriteWithCtx(ctx, "executing transactional query")
	return tx.Tx.PrepareContext(ctx, query)
}

func (t TransactionalClient) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	txRaw, err := persistence.GetTxFromContext(ctx)
	if err != nil {
		t.Logger.Warn().WithField("error", err).WriteWithCtx(ctx, "error getting transaction, using fallback client")
		return t.Next.QueryContext(ctx, query, args...)
	}

	tx, ok := txRaw.(Transaction)
	if !ok {
		t.Logger.WithError(err).WriteWithCtx(ctx, "error casting transaction structure, using fallback client")
		return t.Next.QueryContext(ctx, query, args...)
	}
	t.Logger.Debug().WriteWithCtx(ctx, "executing transactional query")
	return tx.Tx.QueryContext(ctx, query, args...)
}

func (t TransactionalClient) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	txRaw, err := persistence.GetTxFromContext(ctx)
	if err != nil {
		t.Logger.Warn().WithField("error", err).WriteWithCtx(ctx, "error getting transaction, using fallback client")
		return t.Next.QueryRowContext(ctx, query, args...)
	}

	tx, ok := txRaw.(Transaction)
	if !ok {
		t.Logger.WithError(err).WriteWithCtx(ctx, "error casting transaction structure, using fallback client")
		return t.Next.QueryRowContext(ctx, query, args...)
	}
	t.Logger.Debug().WriteWithCtx(ctx, "executing transactional query")
	return tx.Tx.QueryRowContext(ctx, query, args...)
}

func (t TransactionalClient) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return t.Next.BeginTx(ctx, opts)
}

func (t TransactionalClient) Driver() driver.Driver {
	return t.Next.Driver()
}
