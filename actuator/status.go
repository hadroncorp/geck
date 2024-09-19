package actuator

import (
	"encoding"
	"fmt"
)

// Status is the status of a system component. Implements both fmt.Stringer and encoding.TextMarshaler interfaces.
type Status int8

var (
	_ fmt.Stringer           = (*Status)(nil)
	_ encoding.TextMarshaler = (*Status)(nil)
)

const (
	// StatusUnknown component status is not known at the moment.
	StatusUnknown Status = iota
	// StatusUp component is healthy and available.
	StatusUp
	// StatusDown component is unhealthy and unavailable.
	StatusDown
)

var statusTextMap = map[Status]string{
	StatusUnknown: "UNKNOWN",
	StatusUp:      "UP",
	StatusDown:    "DOWN",
}

func (s Status) String() string {
	return statusTextMap[s]
}

func (s Status) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}
