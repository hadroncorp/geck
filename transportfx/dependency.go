package transportfx

import (
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/observability/logging"
	"github.com/hadroncorp/geck/observability/loggingfx"
	"github.com/hadroncorp/geck/transport"
)

func AsControllerHTTP(t any) any {
	return fx.Annotate(t,
		fx.As(new(transport.ControllerHTTP)),
		fx.ResultTags(`group:"root_controllers_http"`),
	)
}

func AsVersionedControllerHTTP(t any) any {
	return fx.Annotate(t,
		fx.As(new(transport.VersionedControllerHTTP)),
		fx.ResultTags(`group:"versioned_controllers_http"`),
	)
}

func AsMiddlewareHTTP(t any) any {
	return fx.Annotate(t,
		fx.ResultTags(`group:"middlewares_http"`),
	)
}

func AsMiddlewaresHTTP(t any) any {
	return fx.Annotate(t,
		fx.ResultTags(`group:"middlewares_groups_http"`),
	)
}

var TransportModuleHTTP = fx.Module("transport_http",
	fx.Decorate(
		loggingfx.DecorateLoggerWithModule("transport.http"),
	),
	fx.Provide(
		transport.NewConfigHTTP,
		transport.NewEcho,
		// middlewares
		AsMiddlewaresHTTP(transport.NewDefaultEchoMiddlewareGroup),
		// controllers
		transport.NewConfigActuatorHTTP,
		AsControllerHTTP(transport.NewActuatorControllerHTTP),
	),
	fx.Invoke(
		transport.RegisterMiddlewaresEcho,
		transport.RegisterControllersEcho,
		func(logger logging.Logger, cfg transport.ConfigHTTP) {
			logger.Debug().
				WithField("routes", cfg.AuthenticationWhitelist).
				Write("skipping security in specified routes")
		},
	),
)

var TransportJWTModuleHTTP = fx.Module("transport_http_jwt",
	fx.Provide(
		transport.NewEchoJWTAuthenticatorConfig,
		AsMiddlewareHTTP(transport.NewEchoJWTAuthenticator),
	),
)
