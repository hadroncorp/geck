package task

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hadroncorp/geck/data/persistence"
	"github.com/hadroncorp/geck/systemerror"
	"github.com/hadroncorp/geck/transport"
)

type ControllerHTTP struct {
	Service                   Service
	TransactionContextFactory persistence.TransactionContextFactory
}

var _ transport.VersionedControllerHTTP = ControllerHTTP{}

func NewControllerHTTP(service Service, txFactory persistence.TransactionContextFactory) ControllerHTTP {
	return ControllerHTTP{
		Service:                   service,
		TransactionContextFactory: txFactory,
	}
}

func (h ControllerHTTP) SetRoutes(_ *echo.Echo) {
}

func (h ControllerHTTP) SetVersionedRoutes(g *echo.Group) {
	g.POST("/tasks", h.create, transport.WithPersistentTransaction(h.TransactionContextFactory))
	g.GET("/tasks/:task_id", h.get)
	g.DELETE("/tasks/:task_id", h.delete)
}

func (h ControllerHTTP) create(c echo.Context) error {
	cmd := CreateCommand{}
	if err := c.Bind(&cmd); err != nil {
		return err
	}

	if err := h.Service.Create(c.Request().Context(), cmd); err != nil {
		return err
	}
	panic(systemerror.NewDomain("INSUFFICIENT_CREDIT", "some message", nil))
	return c.JSON(http.StatusCreated, transport.Data{
		Data: cmd,
	})
}

func (h ControllerHTTP) get(c echo.Context) error {
	entity, err := h.Service.Get(c.Request().Context(), c.Param("task_id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, transport.Data{
		Data: ConvertView(entity),
	})
}

func (h ControllerHTTP) delete(c echo.Context) error {
	return h.Service.Delete(c.Request().Context(), c.Param("task_id"))
}
