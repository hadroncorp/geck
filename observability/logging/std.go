package logging

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hadroncorp/geck/observability/tracing"
)

// StdEvent is standard library (log.Logger) the implementation of Event.
type StdEvent struct {
	level  Level
	logger *log.Logger
	module string
	fields map[string]any
}

var _ Event = &StdEvent{}

// WithField appends a field to the context.
func (s *StdEvent) WithField(field string, val any) Event {
	s.fields[field] = val
	return s
}

// Write writes a new log entry into the Logger instance (most probably will write to an underlying io.Writer instance).
func (s *StdEvent) Write(msg string) {
	buf := strings.Builder{}
	i := 0
	for k, v := range s.fields {
		switch val := v.(type) {
		case string:
			buf.WriteString(fmt.Sprintf("%s:%q", k, val))
		default:
			buf.WriteString(fmt.Sprintf("%s:%v", k, v))
		}

		if i < len(s.fields)-1 {
			buf.WriteByte(' ')
		}
		i++
	}

	var lvl string
	switch s.level {
	case DebugLevel:
		lvl = "DEBUG"
	case InfoLevel:
		lvl = "INFO"
	case WarnLevel:
		lvl = "WARN"
	case TraceLevel:
		lvl = "TRACE"
	case ErrorLevel:
		lvl = "ERROR"
	default:
		lvl = "DEBUG"
	}

	s.logger.Printf("%s %s message:%q", lvl, buf.String(), msg)
}

// WriteWithCtx writes a new log entry into the Logger instance (most probably will write to an underlying io.Writer instance).
//
// Uses context.Context to retrieve (and possibly append) useful information like trace identifiers.
func (s *StdEvent) WriteWithCtx(ctx context.Context, msg string) {
	if span, _ := tracing.GetSpanFromContext(ctx); span != "" {
		s.WithField("span_id", span)
	}
	s.Write(msg)
}

// StdLoggerAdapter is the standard library (log.Logger) implementation of Logger.
type StdLoggerAdapter struct {
	ModuleName string
	Logger     *log.Logger
}

var _ Logger = StdLoggerAdapter{}

// NewStdLoggerAdapter allocates a new StdLoggerAdapter instance.
func NewStdLoggerAdapter(l *log.Logger) StdLoggerAdapter {
	return StdLoggerAdapter{Logger: l}
}

// Level creates an Event context to write a new log entry.
func (s StdLoggerAdapter) Level(lvl Level) Event {
	ev := &StdEvent{
		level:  lvl,
		logger: s.Logger,
		module: s.ModuleName,
		fields: map[string]any{},
	}
	if s.ModuleName != "" {
		ev.WithField("module", s.ModuleName)
	}
	return ev
}

// Module allocates a Logger instance with a module field.
func (s StdLoggerAdapter) Module(name string) Logger {
	s.ModuleName = name
	return s
}

// Debug creates an Event context to write a new log entry with DebugLevel.
func (s StdLoggerAdapter) Debug() Event {
	return s.Level(DebugLevel)
}

// Info creates an Event context to write a new log entry with InfoLevel.
func (s StdLoggerAdapter) Info() Event {
	return s.Level(InfoLevel)
}

// Warn creates an Event context to write a new log entry with WarnLevel.
func (s StdLoggerAdapter) Warn() Event {
	return s.Level(WarnLevel)
}

// Trace creates an Event context to write a new log entry with TraceLevel.
func (s StdLoggerAdapter) Trace() Event {
	return s.Level(TraceLevel)
}

// Error creates an Event context to write a new log entry with ErrorLevel.
func (s StdLoggerAdapter) Error() Event {
	return s.Level(ErrorLevel)
}

// WithError creates an Event context to write a new log entry with ErrorLevel and appends an `error` field.
func (s StdLoggerAdapter) WithError(err error) Event {
	return s.Error().WithField("error", err.Error())
}
