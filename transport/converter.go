package transport

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/hadroncorp/geck/internal/reflection"
	"github.com/hadroncorp/geck/observability/logging"
	"github.com/hadroncorp/geck/systemerror"
)

func convertErrorEcho(err error) Error {
	var echoHttpErr *echo.HTTPError
	ok := errors.As(err, &echoHttpErr)
	if ok {
		return Error{
			Code:     echoHttpErr.Code,
			Message:  fmt.Sprintf("%v", echoHttpErr.Message),
			Status:   http.StatusText(echoHttpErr.Code),
			Internal: echoHttpErr.Internal,
		}
	}
	return Error{
		Code:     http.StatusInternalServerError,
		Message:  http.StatusText(http.StatusInternalServerError),
		Status:   "INTERNAL_SERVER_ERROR",
		Internal: err,
	}
}

func convertSystemErrorEcho(err error) Error {
	var sysErr systemerror.Error
	ok := errors.As(err, &sysErr)
	if !ok {
		return convertErrorEcho(err)
	}
	return Error{
		Code:    statusCodeHTTPMap[sysErr.Status()],
		Message: sysErr.LocalizedMessage(),
		Status:  sysErr.Status().String(),
		Details: []ErrorDetail{
			{
				Type:     reflection.NewTypeFullNameAny(sysErr),
				Reason:   sysErr.Reason(),
				Metadata: sysErr.Metadata(),
			},
		},
		Internal: err,
	}
}

func convertContainerErrorsEcho(srcErr error) Errors {
	var containerErr systemerror.Container
	ok := errors.As(srcErr, &containerErr)
	if !ok {
		sysErr := convertSystemErrorEcho(srcErr)
		return Errors{
			Code:   sysErr.Code,
			Errors: []Error{sysErr},
		}
	}

	srcErrs := containerErr.Unwrap()
	errs := Errors{
		Code:   0,
		Errors: make([]Error, 0, len(srcErrs)),
	}
	for _, err := range srcErrs {
		sysErr := convertSystemErrorEcho(err)
		if sysErr.Code > errs.Code {
			errs.Code = sysErr.Code
		}
		errs.Errors = append(errs.Errors, sysErr)
	}
	return errs
}

func convertErrorEchoJWT(ctx context.Context, logger logging.Logger, err error) error {
	var tokenErr *echojwt.TokenError
	ok := errors.As(err, &tokenErr)
	if ok {
		logger.WithError(tokenErr.Err).WriteWithCtx(ctx, "got jwt error")
	}
	return systemerror.NewUnauthenticated()
}
