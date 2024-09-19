package systemerror

import "github.com/samber/lo"

// Container is an aggregation structure that holds a slice of errors.
// Commonly already implemented by the returning value from errors.Join routine.
type Container interface {
	Unwrap() []error
}

// Error is an extension of the standard lib's error interface, bringing high-level error metadata like Status fields.
// Extends built-in error interface.
type Error interface {
	error
	// Unwrap returns static error (if any). Use it to run error type checks.
	//
	// For example:
	//	ok := errors.Is(err, ErrOutOfRange) // returns true if StaticError is ErrOutOfRange.
	Unwrap() error
	// Status returns the system ErrStatus code of this error.
	Status() Status
	// Reason returns domain-owned ErrReason code. Similar to ErrStatus but code might not be known by external systems.
	Reason() string
	// Message returns error ErrMessage.
	Message() string
	// LocalizedMessage returns locale-based error message. Returns Message as fallback value if localized ErrMessage is empty.
	LocalizedMessage() string
	// Metadata returns custom parameters related to the error.
	Metadata() map[string]string
}

// SystemError is a generic error structure that can be extended or directly used to generate errors in a unified way.
//
// Implements Error (from this same package) and fmt.Unwrap interfaces.
type SystemError struct {
	// ErrStatus the system ErrStatus code of this error.
	ErrStatus Status
	// ErrReason domain-owned ErrReason code. Similar to ErrStatus but code might not be known by external systems.
	ErrReason string
	// ErrMessage error message.
	ErrMessage string
	// ErrLocalizedMessage locale-based error ErrMessage.
	ErrLocalizedMessage string
	// ErrMetadata custom parameters related to the error.
	ErrMetadata map[string]string
	// StaticError parent static error
	StaticError error
}

var (
	_ Error = SystemError{}
)

// Error returns error as string.
func (e SystemError) Error() string {
	return e.ErrMessage
}

// Unwrap returns static error (if any). Use it to run error type checks.
//
// For example:
//
//	ok := errors.Is(err, ErrOutOfRange) // returns true if StaticError is ErrOutOfRange.
func (e SystemError) Unwrap() error {
	return e.StaticError
}

// Status returns the system ErrStatus code of this error.
func (e SystemError) Status() Status {
	return e.ErrStatus
}

// Reason returns domain-owned ErrReason code. Similar to ErrStatus but code might not be known by external systems.
func (e SystemError) Reason() string {
	return e.ErrReason
}

// Message returns error message.
func (e SystemError) Message() string {
	return e.ErrMessage
}

// LocalizedMessage returns locale-based error message. Returns Message as fallback value if localized error is empty.
func (e SystemError) LocalizedMessage() string {
	return lo.CoalesceOrEmpty(e.ErrLocalizedMessage, e.ErrMessage)
}

// Metadata returns custom parameters related to the error.
func (e SystemError) Metadata() map[string]string {
	return e.ErrMetadata
}
