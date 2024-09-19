package tracing

import (
	"context"

	"github.com/hadroncorp/geck/identifier"
)

var (
	spanFactoryID identifier.FactoryUUID
)

type SpanContextType string

const SpanContextKey SpanContextType = "geck.tracing.span_id"

// NewTracedContext appends a generated span id to the given context.Context.
//
// Uses identifier.FactoryUUID as identifier generation algorithm.
func NewTracedContext(ctx context.Context) context.Context {
	id, _ := spanFactoryID.NewIdentifier()
	return context.WithValue(ctx, SpanContextKey, id)
}

// GetSpanFromContext retrieves a span identifier from the given context.Context. Produces ErrSpanNotFound if
// not found or id cannot be cast.
func GetSpanFromContext(ctx context.Context) (string, error) {
	span, ok := ctx.Value(SpanContextKey).(string)
	if !ok {
		return "", ErrSpanNotFound
	}
	return span, nil
}
