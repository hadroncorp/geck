package loggingfx

import (
	"log"
	"os"

	"github.com/rs/zerolog"
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/application"
	"github.com/hadroncorp/geck/observability/logging"
)

var StdLoggerModule = fx.Module("logger_std",
	fx.Provide(
		func() *log.Logger {
			return log.New(os.Stdout, "", 0)
		},
		fx.Annotate(
			logging.NewStdLoggerAdapter,
			fx.As(new(logging.Logger)),
		),
	),
)

var ZerologLoggerModule = fx.Module("logger_zerolog",
	fx.Provide(
		logging.NewZerologDefaultLogger,
		fx.Annotate(
			logging.NewZerologLoggerAdapter,
			fx.As(new(logging.Logger)),
		),
	),
)

var ZerologAppLoggerModule = fx.Module("logger_zerolog_app",
	fx.Provide(
		func(cfg application.Config) zerolog.Logger {
			return logging.NewApplicationZerologLogger(cfg, os.Stdout)
		},
		fx.Annotate(
			logging.NewZerologLoggerAdapter,
			fx.As(new(logging.Logger)),
		),
	),
)

func DecorateLoggerWithModule(moduleName string) any {
	return func(logger logging.Logger) logging.Logger {
		return logger.Module(moduleName)
	}
}
