package systemerror

import (
	"errors"
	"fmt"

	"github.com/hadroncorp/geck/internal/reflection"
)

// ErrAlreadyExists the specified element already exists.
var ErrAlreadyExists = errors.New("already exists")

// NewAlreadyExists allocates a SystemError with StatusAlreadyExists and ErrAlreadyExists.
//
// The element already exists.
func NewAlreadyExists(reason, message string, metadata map[string]string) SystemError {
	return SystemError{
		ErrStatus:   StatusAlreadyExists,
		ErrReason:   reason,
		ErrMessage:  message,
		ErrMetadata: metadata,
		StaticError: ErrAlreadyExists,
	}
}

// NewResourceAlreadyExists allocates a SystemError with StatusAlreadyExists and ErrAlreadyExists.
//
// The resource already exists.
// Aimed for resource-oriented systems.
func NewResourceAlreadyExists[T any](key string) SystemError {
	return SystemError{
		ErrStatus:  StatusAlreadyExists,
		ErrReason:  "RESOURCE_ALREADY_EXISTS",
		ErrMessage: fmt.Sprintf("resource already exists: %s", reflection.NewTypeName[T]()),
		ErrMetadata: map[string]string{
			"key": key,
		},
		StaticError: ErrAlreadyExists,
	}
}
