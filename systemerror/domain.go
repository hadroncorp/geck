package systemerror

import "errors"

// ErrDomain static error for generic domain operations.
var ErrDomain = errors.New("domain error")

// NewDomain allocates a SystemError with StatusDomain and ErrDomain.
//
// A domain error has been detected.
func NewDomain(reason, message string, metadata map[string]string) SystemError {
	return SystemError{
		ErrStatus:   StatusDomain,
		ErrReason:   reason,
		ErrMessage:  message,
		ErrMetadata: metadata,
		StaticError: ErrDomain,
	}
}
