package logger

// Logger handles formatted logging
type Logger struct {
	// TODO(human): Add unexported fields for level and prefix
}

// Option configures a Logger
type Option func(*Logger)

// NewLogger creates a new logger with options
func NewLogger(opts ...Option) *Logger {
	// TODO(human): Implement with defaults and option application
	return nil
}

// WithLevel sets the log level
func WithLevel(level string) Option {
	// TODO(human): Implement
	return nil
}

// WithPrefix sets the log prefix
func WithPrefix(prefix string) Option {
	// TODO(human): Implement
	return nil
}

// Log formats and returns a log message
func (l *Logger) Log(message string) string {
	// TODO(human): Implement formatting with level and prefix
	return ""
}

// GetLevel returns the logger's level
func (l *Logger) GetLevel() string {
	// TODO(human): Implement
	return ""
}
