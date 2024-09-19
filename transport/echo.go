package transport

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/application"
	"github.com/hadroncorp/geck/observability/logging"
)

type NewEchoParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Config    ConfigHTTP
	Logger    logging.Logger
}

func NewEcho(params NewEchoParams) *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = HandleEchoError
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := e.Start(params.Config.Address); err != nil && !errors.Is(err, http.ErrServerClosed) {
					params.Logger.WithError(err).WriteWithCtx(ctx, "failed to start server")
				}
			}()
			params.Logger.Info().WithField("address", params.Config.Address).WriteWithCtx(ctx, "started http server")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
	return e
}

type RegisterMiddlewaresEchoParams struct {
	fx.In

	Echo             *echo.Echo
	Logger           logging.Logger
	GroupMiddlewares [][]echo.MiddlewareFunc `group:"middlewares_groups_http"`
	Middlewares      []echo.MiddlewareFunc   `group:"middlewares_http"`
}

func RegisterMiddlewaresEcho(params RegisterMiddlewaresEchoParams) {
	params.Logger.Info().WithField("total_middleware_groups", len(params.GroupMiddlewares)).Write("registering http middleware groups")
	params.Logger.Info().WithField("total_middlewares", len(params.Middlewares)).Write("registering http middlewares")
	groupCount := 0
	for _, middlewareGroup := range params.GroupMiddlewares {
		for _, middleware := range middlewareGroup {
			params.Echo.Use(middleware)
			groupCount++
		}
	}
	params.Echo.Use(params.Middlewares...)
	params.Logger.Info().WithField("total_middlewares", groupCount+len(params.Middlewares)).Write("registered http middlewares")
}

type RegisterControllersEchoParams struct {
	fx.In

	Echo                 *echo.Echo
	Config               application.Config
	Logger               logging.Logger
	RootControllers      []ControllerHTTP          `group:"root_controllers_http"`
	VersionedControllers []VersionedControllerHTTP `group:"versioned_controllers_http"`
}

func RegisterControllersEcho(params RegisterControllersEchoParams) {
	params.Logger.Info().
		WithField("total_controllers", len(params.RootControllers)).
		Write("registering http root controllers")
	for _, controller := range params.RootControllers {
		controller.SetRoutes(params.Echo)
	}
	basePath := fmt.Sprintf("/%s", params.Config.Semver.Major)
	g := params.Echo.Group(basePath)
	params.Logger.Info().
		WithField("total_controllers", len(params.VersionedControllers)).
		WithField("base_path", basePath).
		Write("registering http versioned controllers")
	for _, controller := range params.VersionedControllers {
		controller.SetRoutes(params.Echo)
		controller.SetVersionedRoutes(g)
	}
}
