package tracing

import "errors"

// ErrSpanNotFound no trace id was found.
var ErrSpanNotFound = errors.New("span not found")
