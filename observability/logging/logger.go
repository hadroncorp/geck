package logging

import "context"

// Level is a piece of information telling how important a given log message is.
type Level uint8

const (
	_ Level = iota
	// DebugLevel DEBUG log level.
	DebugLevel
	// InfoLevel INFO log level.
	InfoLevel
	// WarnLevel WARN log level.
	WarnLevel
	// TraceLevel TRACE log level.
	TraceLevel
	// ErrorLevel ERROR log level.
	ErrorLevel
)

// An Event represents a logging context instance. Use it to write a new log entry.
type Event interface {
	// WithField appends a field to the context.
	WithField(field string, val any) Event
	// Write writes a new log entry into the Logger instance (most probably will write to an underlying io.Writer instance).
	Write(msg string)
	// WriteWithCtx writes a new log entry into the Logger instance (most probably will write to an underlying io.Writer instance).
	//
	// Uses context.Context to retrieve (and possibly append) useful information like trace identifiers.
	WriteWithCtx(ctx context.Context, msg string)
}

// A Logger is used to log messages for a specific system or application component.
type Logger interface {
	// Level creates an Event context to write a new log entry.
	Level(lvl Level) Event
	// Module allocates a Logger instance with a module field.
	Module(name string) Logger
	// Debug creates an Event context to write a new log entry with DebugLevel.
	Debug() Event
	// Info creates an Event context to write a new log entry with InfoLevel.
	Info() Event
	// Warn creates an Event context to write a new log entry with WarnLevel.
	Warn() Event
	// Trace creates an Event context to write a new log entry with TraceLevel.
	Trace() Event
	// Error creates an Event context to write a new log entry with ErrorLevel.
	Error() Event
	// WithError creates an Event context to write a new log entry with ErrorLevel and appends an `error` field.
	WithError(err error) Event
}
