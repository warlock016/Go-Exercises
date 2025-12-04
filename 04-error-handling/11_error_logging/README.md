# Exercise 11: Error Logging

## 🎯 Learning Goal
Implement structured error logging with context preservation, log levels, and traceability.

## 📝 Problem Description

Production systems need to log errors with enough context to debug issues. This exercise teaches structured logging patterns.

## 🔧 Function Signatures

```go
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
    Level      LogLevel
    Message    string
    Error      error
    Context    map[string]interface{}
    Timestamp  time.Time
}

// LogError logs an error with context
func LogError(err error, context map[string]interface{}) LogEntry

// FormatLogEntry formats a log entry as JSON
func FormatLogEntry(entry LogEntry) string
```

## 🎓 What This Teaches

- **Structured logging** - Key-value pairs vs unstructured strings
- **Log levels** - DEBUG, INFO, WARN, ERROR
- **Context preservation** - Adding metadata to logs

---

**Next Exercise:** `12_repository_errors` - Data layer error handling
