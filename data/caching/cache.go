package caching

import (
	"context"
)

type Cache interface {
	Set(ctx context.Context, key string, value []byte) error
	SetMany(ctx context.Context, keyValues map[string][]byte) error
	Append(ctx context.Context, key string, value []byte) error
	Add(ctx context.Context, key string, value []byte) error
	List(ctx context.Context, key string) ([][]byte, error)
	Get(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
	DeleteMany(ctx context.Context, keys []string) error
}
