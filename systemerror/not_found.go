package systemerror

import (
	"errors"
	"fmt"

	"github.com/hadroncorp/geck/internal/reflection"
)

// ErrNotFound the specified element was not found.
var ErrNotFound = errors.New("not found")

// NewNotFound allocates a SystemError with StatusNotFound and ErrNotFound.
//
// The specified element was not found.
func NewNotFound(reason, message string, metadata map[string]string) SystemError {
	return SystemError{
		ErrStatus:   StatusNotFound,
		ErrReason:   reason,
		ErrMessage:  message,
		ErrMetadata: metadata,
		StaticError: ErrNotFound,
	}
}

// NewResourceNotFound allocates a SystemError with StatusNotFound and ErrNotFound.
//
// The specified resource was not found.
// Aimed for resource-oriented systems.
//
//   - T : Resource type.
func NewResourceNotFound[T any](key string) SystemError {
	return SystemError{
		ErrStatus:  StatusNotFound,
		ErrReason:  "RESOURCE_NOT_FOUND",
		ErrMessage: fmt.Sprintf("resource '%s' not found", reflection.NewTypeName[T]()),
		ErrMetadata: map[string]string{
			"resource_key": key,
		},
		StaticError: ErrNotFound,
	}
}
