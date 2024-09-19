package applicationfx

import (
	"github.com/rs/zerolog"
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/application"
)

var ApplicationModule = fx.Module("application",
	fx.Provide(
		application.NewConfig,
	),
	fx.Invoke(
		func(logger zerolog.Logger, cfg application.Config) {
			logger.Info().
				Str("name", cfg.ApplicationName).
				Str("version", cfg.Version).
				Str("environment", cfg.Environment).
				Msg("starting application")
		},
	),
)
