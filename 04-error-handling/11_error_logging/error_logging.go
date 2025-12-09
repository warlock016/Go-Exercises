package error_logging

import (
	"encoding/json"
	"strconv"
	"time"
)

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
	Level     LogLevel       `json:"level"`
	Message   string         `json:"message"`
	Error     error          `json:"error"`
	Context   map[string]any `json:"context"`
	Timestamp time.Time      `json:"timestamp"`
}

// LogError logs an error with context
func LogError(err error, context map[string]any) LogEntry {
	// TODO(human): Implement
	newEntry := LogEntry{
		Level:     ERROR,
		Message:   err.Error(),
		Error:     err,
		Context:   context,
		Timestamp: time.Now(),
	}
	return newEntry
}

// FormatLogEntry formats a log entry as JSON
func FormatLogEntry(entry LogEntry) string {
	// TODO(human): Implement
	type FormatLogEntry struct {
		Level     string         `json:"level"`
		Message   string         `json:"message"`
		Error     string         `json:"error"`
		Context   map[string]any `json:"context"`
		Timestamp string         `json:"timestamp"`
	}

	stringified := FormatLogEntry{
		Message:   entry.Message,
		Timestamp: strconv.FormatInt(entry.Timestamp.Unix(), 10),
		Context:   entry.Context,
	}

	if entry.Error != nil {
		stringified.Error = entry.Error.Error()
	}

	switch entry.Level {
	case DEBUG:
		stringified.Level = "DEBUG"
	case INFO:
		stringified.Level = "INFO"
	case WARN:
		stringified.Level = "WARN"
	case ERROR:
		stringified.Level = "ERROR"
	default:
		stringified.Level = "UNKNOWN"
	}

	conv, err := json.Marshal(stringified)

	if err != nil {
		return ""
	}

	return string(conv)
}
