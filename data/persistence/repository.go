package persistence

import (
	"context"

	"github.com/hadroncorp/geck/data"
)

type WriteRepository[T Persistable] interface {
	Save(ctx context.Context, entity T) error
	SaveMany(ctx context.Context, entities []T) error
	Remove(ctx context.Context, entity T) error
}

type ReadRepository[T any, K comparable] interface {
	FindByKey(ctx context.Context, key K) (*T, error)
}

type PagingRepository[T any] interface {
	FindAll(ctx context.Context, criteria data.Criteria) (data.Page[T], error)
}

type CrudRepository[T Persistable, K comparable] interface {
	WriteRepository[T]
	ReadRepository[T, K]
}

type PagingCrudRepository[T Persistable, K comparable] interface {
	WriteRepository[T]
	ReadRepository[T, K]
	PagingRepository[T]
}
