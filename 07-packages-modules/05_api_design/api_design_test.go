package apidesign

import (
	"strings"
	"testing"

	"github.com/warlock016/go-exercises/07-packages-modules/05_api_design/logger"
)

func TestLoggerDefaults(t *testing.T) {
	log := logger.NewLogger()

	if log == nil {
		t.Fatal("NewLogger() returned nil")
	}

	// Should have default level INFO
	if log.GetLevel() != "INFO" {
		t.Errorf("Default level = %q, want %q", log.GetLevel(), "INFO")
	}

	// Should log with default format
	result := log.Log("test message")
	if !strings.Contains(result, "INFO") || !strings.Contains(result, "test message") {
		t.Errorf("Log() = %q, should contain INFO and message", result)
	}
}

func TestLoggerWithLevel(t *testing.T) {
	tests := []struct {
		level   string
		message string
	}{
		{"DEBUG", "debug info"},
		{"INFO", "info message"},
		{"WARN", "warning"},
		{"ERROR", "error occurred"},
	}

	for _, tt := range tests {
		t.Run(tt.level, func(t *testing.T) {
			log := logger.NewLogger(logger.WithLevel(tt.level))

			if log.GetLevel() != tt.level {
				t.Errorf("GetLevel() = %q, want %q", log.GetLevel(), tt.level)
			}

			result := log.Log(tt.message)
			if !strings.Contains(result, tt.level) {
				t.Errorf("Log() = %q, should contain level %q", result, tt.level)
			}
			if !strings.Contains(result, tt.message) {
				t.Errorf("Log() = %q, should contain message %q", result, tt.message)
			}
		})
	}
}

func TestLoggerWithPrefix(t *testing.T) {
	tests := []struct {
		prefix  string
		message string
	}{
		{"myapp", "starting"},
		{"api", "request received"},
		{"db", "connected"},
	}

	for _, tt := range tests {
		t.Run(tt.prefix, func(t *testing.T) {
			log := logger.NewLogger(logger.WithPrefix(tt.prefix))

			result := log.Log(tt.message)
			if !strings.Contains(result, tt.prefix) {
				t.Errorf("Log() = %q, should contain prefix %q", result, tt.prefix)
			}
			if !strings.Contains(result, tt.message) {
				t.Errorf("Log() = %q, should contain message %q", result, tt.message)
			}
		})
	}
}

func TestLoggerWithMultipleOptions(t *testing.T) {
	tests := []struct {
		name    string
		level   string
		prefix  string
		message string
	}{
		{
			name:    "Debug with prefix",
			level:   "DEBUG",
			prefix:  "app",
			message: "test",
		},
		{
			name:    "Error with prefix",
			level:   "ERROR",
			prefix:  "service",
			message: "failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := logger.NewLogger(
				logger.WithLevel(tt.level),
				logger.WithPrefix(tt.prefix),
			)

			result := log.Log(tt.message)

			// Should contain all three parts
			if !strings.Contains(result, tt.level) {
				t.Errorf("Log() = %q, should contain level %q", result, tt.level)
			}
			if !strings.Contains(result, tt.prefix) {
				t.Errorf("Log() = %q, should contain prefix %q", result, tt.prefix)
			}
			if !strings.Contains(result, tt.message) {
				t.Errorf("Log() = %q, should contain message %q", result, tt.message)
			}
		})
	}
}

func TestLoggerOptionsOrder(t *testing.T) {
	// Options should work in any order
	log1 := logger.NewLogger(
		logger.WithLevel("DEBUG"),
		logger.WithPrefix("app"),
	)

	log2 := logger.NewLogger(
		logger.WithPrefix("app"),
		logger.WithLevel("DEBUG"),
	)

	msg := "test"
	result1 := log1.Log(msg)
	result2 := log2.Log(msg)

	if result1 != result2 {
		t.Errorf("Option order should not matter:\n  log1: %q\n  log2: %q", result1, result2)
	}
}

func TestLoggerEncapsulation(t *testing.T) {
	log := logger.NewLogger(logger.WithLevel("SECRET"))

	// We can only access level through GetLevel()
	level := log.GetLevel()
	if level != "SECRET" {
		t.Errorf("GetLevel() = %q, want %q", level, "SECRET")
	}

	// We cannot access log.level directly (it's unexported)
	// This would be a compile error: log.level

	t.Log("Logger correctly encapsulates internal state")
	t.Log("Fields can only be accessed through exported methods")
}

func TestCreateLogger(t *testing.T) {
	result := CreateLogger("WARN", "myapp")

	log, ok := result.(*logger.Logger)
	if !ok {
		t.Fatal("CreateLogger should return *logger.Logger")
	}

	if log.GetLevel() != "WARN" {
		t.Errorf("Created logger level = %q, want %q", log.GetLevel(), "WARN")
	}

	output := log.Log("test")
	if !strings.Contains(output, "WARN") || !strings.Contains(output, "myapp") {
		t.Errorf("Log() = %q, should contain WARN and myapp", output)
	}
}

func TestLogMessage(t *testing.T) {
	result := LogMessage("hello world")

	if result == "" {
		t.Error("LogMessage() returned empty string")
	}

	if !strings.Contains(result, "hello world") {
		t.Errorf("LogMessage() = %q, should contain message", result)
	}
}
