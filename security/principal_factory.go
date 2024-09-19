package security

import (
	"context"
)

type PrincipalFactory[T any] interface {
	NewContextWithPrincipal(parent context.Context, args T) (context.Context, error)
}
