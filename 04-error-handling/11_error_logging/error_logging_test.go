package error_logging

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"
)

func TestLogError(t *testing.T) {
	err := errors.New("test error")
	context := map[string]interface{}{
		"user_id": 123,
		"action":  "create",
	}

	entry := LogError(err, context)

	if entry.Error != err {
		t.Errorf("LogError() Error = %v, want %v", entry.Error, err)
	}
	if entry.Level != ERROR {
		t.Errorf("LogError() Level = %v, want ERROR", entry.Level)
	}
}

func TestLogError_MessageFromError(t *testing.T) {
	err := errors.New("database connection failed")
	entry := LogError(err, nil)

	// Message should reflect the error, not be hardcoded
	if entry.Message == "" {
		t.Error("LogError() Message should not be empty")
	}
	if entry.Message == "debug message" {
		t.Error("LogError() Message appears to be hardcoded - should derive from error")
	}
}

func TestLogError_TimestampSet(t *testing.T) {
	before := time.Now()
	entry := LogError(errors.New("test"), nil)
	after := time.Now()

	if entry.Timestamp.IsZero() {
		t.Error("LogError() Timestamp should be set")
	}
	if entry.Timestamp.Before(before) || entry.Timestamp.After(after) {
		t.Errorf("LogError() Timestamp = %v, should be between %v and %v",
			entry.Timestamp, before, after)
	}
}

func TestLogError_ContextPreserved(t *testing.T) {
	context := map[string]interface{}{
		"user_id":    123,
		"request_id": "abc-123",
		"endpoint":   "/api/users",
	}

	entry := LogError(errors.New("test"), context)

	if entry.Context == nil {
		t.Fatal("LogError() Context should not be nil")
	}
	if entry.Context["user_id"] != 123 {
		t.Errorf("LogError() Context[user_id] = %v, want 123", entry.Context["user_id"])
	}
	if entry.Context["request_id"] != "abc-123" {
		t.Errorf("LogError() Context[request_id] = %v, want abc-123", entry.Context["request_id"])
	}
}

func TestLogError_NilContext(t *testing.T) {
	entry := LogError(errors.New("test"), nil)

	// Should handle nil context gracefully
	if entry.Level != ERROR {
		t.Errorf("LogError() with nil context: Level = %v, want ERROR", entry.Level)
	}
}

func TestFormatLogEntry(t *testing.T) {
	entry := LogEntry{
		Level:   ERROR,
		Message: "test message",
		Error:   errors.New("test error"),
	}

	formatted := FormatLogEntry(entry)
	if formatted == "" {
		t.Error("FormatLogEntry() returned empty string")
	}
}

func TestFormatLogEntry_ValidJSON(t *testing.T) {
	entry := LogEntry{
		Level:     ERROR,
		Message:   "something went wrong",
		Error:     errors.New("connection timeout"),
		Context:   map[string]any{"user_id": 42},
		Timestamp: time.Now(),
	}

	formatted := FormatLogEntry(entry)

	// Must be valid JSON
	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(formatted), &parsed); err != nil {
		t.Errorf("FormatLogEntry() produced invalid JSON: %v\nOutput: %s", err, formatted)
	}
}

func TestFormatLogEntry_ContainsErrorText(t *testing.T) {
	entry := LogEntry{
		Level:   ERROR,
		Message: "operation failed",
		Error:   errors.New("disk full"),
	}

	formatted := FormatLogEntry(entry)

	// The error text should appear somewhere in the output
	// Note: error interface doesn't JSON marshal - implementation needs to handle this
	if !strings.Contains(formatted, "disk full") {
		t.Errorf("FormatLogEntry() should contain error text 'disk full'\nGot: %s", formatted)
	}
}

func TestFormatLogEntry_ContainsMessage(t *testing.T) {
	entry := LogEntry{
		Level:   WARN,
		Message: "rate limit exceeded",
	}

	formatted := FormatLogEntry(entry)

	if !strings.Contains(formatted, "rate limit exceeded") {
		t.Errorf("FormatLogEntry() should contain message\nGot: %s", formatted)
	}
}

func TestFormatLogEntry_ContainsContext(t *testing.T) {
	entry := LogEntry{
		Level:   INFO,
		Message: "user action",
		Context: map[string]any{
			"user_id": 999,
			"action":  "login",
		},
	}

	formatted := FormatLogEntry(entry)

	if !strings.Contains(formatted, "999") {
		t.Errorf("FormatLogEntry() should contain context value 999\nGot: %s", formatted)
	}
	if !strings.Contains(formatted, "login") {
		t.Errorf("FormatLogEntry() should contain context value 'login'\nGot: %s", formatted)
	}
}

func TestFormatLogEntry_NilError(t *testing.T) {
	entry := LogEntry{
		Level:   INFO,
		Message: "operation completed",
		Error:   nil,
	}

	formatted := FormatLogEntry(entry)

	// Should not panic and should produce valid JSON
	if formatted == "" {
		t.Error("FormatLogEntry() with nil error returned empty string")
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(formatted), &parsed); err != nil {
		t.Errorf("FormatLogEntry() with nil error produced invalid JSON: %v", err)
	}
}

func TestLogLevel_Values(t *testing.T) {
	// Verify log level ordering (DEBUG < INFO < WARN < ERROR)
	if DEBUG >= INFO {
		t.Error("DEBUG should be less than INFO")
	}
	if INFO >= WARN {
		t.Error("INFO should be less than WARN")
	}
	if WARN >= ERROR {
		t.Error("WARN should be less than ERROR")
	}
}
