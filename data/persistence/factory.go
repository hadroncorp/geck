package persistence

import (
	"context"
)

type TransactionContextFactory interface {
	NewContext(parent context.Context) (context.Context, error)
}
