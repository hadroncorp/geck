package systemerror

import "fmt"

// Status is the system status of an Error. Implements fmt.Stringer, returning a Status full name as string
// (e.g. StatusInvalidArgument -> "INVALID_ARGUMENT").
//
// Systems should use these to map transport error codes (e.g. StatusNotFound -> http.StatusNotFound).
type Status uint64

var _ fmt.Stringer = Status(0)

const (
	StatusInvalidArgument Status = iota + 1
	StatusFailedPrecondition
	StatusOutOfRange
	StatusUnauthenticated
	StatusPermissionDenied
	StatusNotFound
	StatusAborted
	StatusAlreadyExists
	StatusResourceExhausted
	StatusCancelled
	StatusDataLoss
	StatusUnknown
	StatusInternal
	StatusNotImplemented
	StatusUnavailable
	StatusDeadlineExceeded
	StatusDomain
)

var statusStringMap = map[Status]string{
	StatusInvalidArgument:    "INVALID_ARGUMENT",
	StatusFailedPrecondition: "FAILED_PRECONDITION",
	StatusOutOfRange:         "OUT_OF_RANGE",
	StatusUnauthenticated:    "UNAUTHENTICATED",
	StatusPermissionDenied:   "PERMISSION_DENIED",
	StatusNotFound:           "NOT_FOUND",
	StatusAborted:            "ABORTED",
	StatusAlreadyExists:      "ALREADY_EXISTS",
	StatusResourceExhausted:  "RESOURCE_EXHAUSTED",
	StatusCancelled:          "CANCELLED",
	StatusDataLoss:           "DATA_LOSS",
	StatusUnknown:            "UNKNOWN",
	StatusInternal:           "INTERNAL",
	StatusNotImplemented:     "NOT_IMPLEMENTED",
	StatusUnavailable:        "UNAVAILABLE",
	StatusDeadlineExceeded:   "DEADLINE_EXCEEDED",
	StatusDomain:             "DOMAIN",
}

func (s Status) String() string {
	return statusStringMap[s]
}
