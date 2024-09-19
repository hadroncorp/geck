package sql

import (
	"context"

	"github.com/hadroncorp/geck/actuator"
	"github.com/hadroncorp/geck/internal/reflection"
)

type Actuator struct {
	Client Client
}

var _ actuator.Actuator = (*Actuator)(nil)

func NewActuator(client Client) Actuator {
	return Actuator{
		Client: client,
	}
}

func (a Actuator) State(ctx context.Context) (actuator.State, error) {
	row := a.Client.QueryRowContext(ctx, "SELECT version()")
	if err := row.Err(); err != nil {
		return actuator.State{
			Status:      actuator.StatusDown,
			Description: err.Error(),
		}, nil
	}
	var version string
	if err := row.Scan(&version); err != nil {
		return actuator.State{
			Status:      actuator.StatusDown,
			Description: err.Error(),
		}, nil
	}

	return actuator.State{
		Status: actuator.StatusUp,
		Details: map[string]any{
			"driver_name": reflection.NewTypeFullNameAny(a.Client.Driver()),
			"version":     version,
		},
	}, nil
}
