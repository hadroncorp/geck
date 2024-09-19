package caching

import (
	"context"

	"github.com/allegro/bigcache/v3"
	"go.uber.org/fx"
)

func NewBigCache(lifecycle fx.Lifecycle, cfg BigCacheConfig) (*bigcache.BigCache, error) {
	bc, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cfg.ItemTTL))
	if err != nil {
		return nil, err
	}
	lifecycle.Append(fx.Hook{
		OnStart: nil,
		OnStop: func(ctx context.Context) error {
			return bc.Close()
		},
	})
	return bc, nil
}
