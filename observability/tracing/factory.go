package tracing

import (
	"context"

	"github.com/hadroncorp/geck/identifier"
)

type TraceFactory interface {
	NewTracedContext(ctx context.Context) context.Context
}

// TraceFactoryTemplate is the default implementation of TraceFactory. Inject identifier.Factory to set
// identifier generation algorithm. Use NewTracedContext if you do not want to use a custom identifier generator.
type TraceFactoryTemplate struct {
	FactoryID identifier.Factory
}

var _ TraceFactory = TraceFactoryTemplate{}

func NewTraceFactoryTemplate(factory identifier.Factory) TraceFactoryTemplate {
	return TraceFactoryTemplate{
		FactoryID: factory,
	}
}

func (t TraceFactoryTemplate) NewTracedContext(ctx context.Context) context.Context {
	id, _ := t.FactoryID.NewIdentifier()
	return context.WithValue(ctx, SpanContextKey, id)
}
