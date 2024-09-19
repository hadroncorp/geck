package cachingfx

import (
	"github.com/caarlos0/env/v11"
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/data/caching"
)

var CacheEmbeddedModule = fx.Module("cache_embedded",
	fx.Provide(
		env.ParseAs[caching.BigCacheConfig],
		caching.NewBigCache,
		fx.Annotate(
			caching.NewCacheEmbedded,
			fx.As(new(caching.Cache)),
		),
	),
)
