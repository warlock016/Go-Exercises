package interface_composition

import (
	"testing"
)

func TestFileLoggerLog(t *testing.T) {
	logger := &FileLogger{Filename: "test.log"}

	logger.Log("First message")
	logger.Log("Second message")
	logger.Log("Third message")

	logs := logger.GetLogs()
	if len(logs) != 3 {
		t.Errorf("GetLogs() len = %d, want 3", len(logs))
	}

	want := []string{"First message", "Second message", "Third message"}
	for i, msg := range want {
		if logs[i] != msg {
			t.Errorf("GetLogs()[%d] = %q, want %q", i, logs[i], msg)
		}
	}
}

func TestFileLoggerClose(t *testing.T) {
	logger := &FileLogger{Filename: "test.log"}

	logger.Log("Before close")
	err := logger.Close()

	if err != nil {
		t.Errorf("Close() error = %v, want nil", err)
	}

	// Verify closed
	if !logger.closed {
		t.Error("Logger should be marked as closed")
	}
}

func TestFileLoggerImplementsLogger(t *testing.T) {
	var _ Logger = &FileLogger{}
	t.Log("✓ FileLogger implements Logger interface")
}

func TestFileLoggerImplementsCloser(t *testing.T) {
	var _ Closer = &FileLogger{}
	t.Log("✓ FileLogger implements Closer interface")
}

func TestFileLoggerImplementsLogCloser(t *testing.T) {
	var _ LogCloser = &FileLogger{}
	t.Log("✓ FileLogger implements LogCloser interface")
}

func TestInterfaceComposition(t *testing.T) {
	logger := &FileLogger{Filename: "app.log"}

	// Can assign to Logger interface
	var l Logger = logger
	l.Log("As Logger")

	// Can assign to Closer interface
	var c Closer = logger
	_ = c

	// Can assign to LogCloser interface
	var lc LogCloser = logger
	lc.Log("As LogCloser")
	lc.Close()

	logs := logger.GetLogs()
	if len(logs) != 2 {
		t.Errorf("Should have 2 log messages, got %d", len(logs))
	}
}

func TestMultipleLoggers(t *testing.T) {
	logger1 := &FileLogger{Filename: "app1.log"}
	logger2 := &FileLogger{Filename: "app2.log"}

	logger1.Log("App1 message")
	logger2.Log("App2 message")

	if len(logger1.GetLogs()) != 1 || logger1.GetLogs()[0] != "App1 message" {
		t.Error("Logger1 has incorrect logs")
	}

	if len(logger2.GetLogs()) != 1 || logger2.GetLogs()[0] != "App2 message" {
		t.Error("Logger2 has incorrect logs")
	}
}
