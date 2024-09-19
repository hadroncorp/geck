package transport

import (
	"net/http"

	"github.com/hadroncorp/geck/systemerror"
)

// took from: https://cloud.google.com/apis/design/errors
var statusCodeHTTPMap = map[systemerror.Status]int{
	systemerror.StatusInvalidArgument:    http.StatusBadRequest,
	systemerror.StatusFailedPrecondition: http.StatusPreconditionFailed,
	systemerror.StatusOutOfRange:         http.StatusBadRequest,
	systemerror.StatusUnauthenticated:    http.StatusUnauthorized,
	systemerror.StatusPermissionDenied:   http.StatusForbidden,
	systemerror.StatusNotFound:           http.StatusNotFound,
	systemerror.StatusAborted:            http.StatusConflict,
	systemerror.StatusAlreadyExists:      http.StatusConflict,
	systemerror.StatusResourceExhausted:  http.StatusTooManyRequests,
	systemerror.StatusCancelled:          499,
	systemerror.StatusDataLoss:           http.StatusInternalServerError,
	systemerror.StatusUnknown:            http.StatusInternalServerError,
	systemerror.StatusInternal:           http.StatusInternalServerError,
	systemerror.StatusNotImplemented:     http.StatusNotImplemented,
	systemerror.StatusUnavailable:        http.StatusServiceUnavailable,
	systemerror.StatusDeadlineExceeded:   http.StatusGatewayTimeout,
	systemerror.StatusDomain:             http.StatusNotAcceptable,
}
