package systemerror

import "errors"

// ErrUnauthenticated no principal was found or its session has expired.
var ErrUnauthenticated = errors.New("unauthenticated")

// NewUnauthenticated allocates a new SystemError using StatusUnauthenticated and ErrUnauthenticated.
//
// The principal was not found or its session has expired.
func NewUnauthenticated() SystemError {
	return SystemError{
		ErrStatus:   StatusUnauthenticated,
		ErrReason:   "PRINCIPAL_UNAUTHENTICATED",
		ErrMessage:  "principal is not authenticated",
		StaticError: ErrUnauthenticated,
	}
}
