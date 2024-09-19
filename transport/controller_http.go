package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/actuator"
	"github.com/hadroncorp/geck/observability/logging"
)

type ControllerHTTP interface {
	SetRoutes(e *echo.Echo)
}

type VersionedControllerHTTP interface {
	ControllerHTTP
	SetVersionedRoutes(g *echo.Group)
}

type ActuatorControllerHTTP struct {
	Manager *actuator.Manager
	Logger  logging.Logger
	Config  ConfigActuatorHTTP
}

var _ ControllerHTTP = (*ActuatorControllerHTTP)(nil)

type NewActuatorControllerHTTPParams struct {
	fx.In
	Manager *actuator.Manager
	Logger  logging.Logger
	Config  ConfigActuatorHTTP
}

func NewActuatorControllerHTTP(params NewActuatorControllerHTTPParams) ActuatorControllerHTTP {
	return ActuatorControllerHTTP{
		Manager: params.Manager,
		Logger:  params.Logger,
		Config:  params.Config,
	}
}

func (a ActuatorControllerHTTP) SetRoutes(e *echo.Echo) {
	// K8s/AWS ECS/Load Balancer/Proxy probes
	e.GET("/healthz", a.getLiveness)
	e.GET("/readiness", a.getReadiness)
	// Java Spring Boot like actuators
	// These endpoints are protected as they expose system sensitive data and
	// attackers can potentially use it against running system.
	//
	// Health and readiness endpoints are still available without requiring any kind of authN and authZ,
	// specially made for container orchestrator/load balancer/proxy tasks which require health-checking and more.
	e.GET("/actuator/health", a.getHealth,
		HasAnyAuthoritiesEcho(a.Config.ActuatorRoleAllowlist...))
	e.GET("/actuator/info", a.getInfo,
		HasAnyAuthoritiesEcho(a.Config.ActuatorRoleAllowlist...))
}

func (a ActuatorControllerHTTP) getLiveness(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (a ActuatorControllerHTTP) getReadiness(c echo.Context) error {
	state, err := a.Manager.Health(c.Request().Context())
	if err != nil || state.Status != actuator.StatusUp {
		return c.NoContent(http.StatusServiceUnavailable)
	}
	return c.NoContent(http.StatusOK)
}

func (a ActuatorControllerHTTP) getHealth(c echo.Context) error {
	state, err := a.Manager.Health(c.Request().Context())
	if err != nil || state.Status != actuator.StatusUp {
		a.Logger.WithError(err).Write("health check failed")
		return c.JSON(http.StatusServiceUnavailable, Data{
			Data: state,
		})
	}
	return c.JSON(http.StatusOK, Data{
		Data: state,
	})
}

func (a ActuatorControllerHTTP) getInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, Data{
		Data: a.Manager.Info(c.Request().Context()),
	})
}
