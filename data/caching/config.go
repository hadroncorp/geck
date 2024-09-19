package caching

import "time"

type BigCacheConfig struct {
	ItemTTL time.Duration `env:"BIG_CACHE_ITEM_TTL" envDefault:"5m"`
}
