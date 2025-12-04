package error_logging

import (
	"errors"
	"testing"
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
