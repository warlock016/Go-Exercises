package error_logging

import "time"

// LogLevel represents logging severity
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// LogEntry represents a structured log entry
type LogEntry struct {
	Level     LogLevel
	Message   string
	Error     error
	Context   map[string]interface{}
	Timestamp time.Time
}

// LogError logs an error with context
func LogError(err error, context map[string]interface{}) LogEntry {
	// TODO(human): Implement
	return LogEntry{}
}

// FormatLogEntry formats a log entry as JSON
func FormatLogEntry(entry LogEntry) string {
	// TODO(human): Implement
	return ""
}
