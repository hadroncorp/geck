package actuator

import (
	"context"
)

// State is an informational structure containing insights about a certain system component ant its current state (UP, DOWN).
//
// For example, a database.
type State struct {
	// Status indicates if the component is currently available or not.
	Status Status `json:"status"`
	// Description component descriptive information.
	Description string `json:"description,omitempty"`
	// Details component additional metadata (e.g. version, metrics).
	Details any `json:"details"`
}

// Actuator is a system agent used to indicate state of a certain system component (e.g. a database, host disk).
type Actuator interface {
	// State returns the current state of the target component. Returns error if communication with component
	// has failed (not the same as State.Status).
	State(ctx context.Context) (State, error)
}
