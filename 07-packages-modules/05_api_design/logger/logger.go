package logger

import "fmt"

// Logger handles formatted logging
type Logger struct {
	// TODO(human): Add unexported fields for level and prefix
	level  string
	prefix string
}

// Option configures a Logger
type Option func(*Logger)

// NewLogger creates a new logger with options
func NewLogger(opts ...Option) *Logger {
	// TODO(human): Implement with defaults and option application

	result := &Logger{
		level:  "INFO",
		prefix: "",
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

// WithLevel sets the log level
func WithLevel(level string) Option {
	// TODO(human): Implement
	return func(l *Logger) {
		l.level = level
	}
}

// WithPrefix sets the log prefix
func WithPrefix(prefix string) Option {
	// TODO(human): Implement
	return func(l *Logger) {
		l.prefix = prefix
	}
}

// Log formats and returns a log message
func (l *Logger) Log(message string) string {
	// TODO(human): Implement formatting with level and prefix
	if l.prefix != "" {
		return fmt.Sprintf("%s [%s] %s", l.prefix, l.level, message)
	}
	return fmt.Sprintf("[%s] %s", l.level, message)
}

// GetLevel returns the logger's level
func (l *Logger) GetLevel() string {
	// TODO(human): Implement
	return l.level
}
