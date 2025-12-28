package interface_composition

import (
	"errors"
)

// TODO(human): Define Logger interface
type Logger interface {
	Log(message string)
}

// TODO(human): Define Closer interface
type Closer interface {
	Close() error
}

// TODO(human): Define LogCloser interface (compose Logger and Closer)
type LogCloser interface {
	Logger
	Closer
}

// TODO(human): Define FileLogger struct
type FileLogger struct {
	Filename string
	logs     []string
	closed   bool
}

// TODO(human): Implement Log() method for FileLogger
func (f *FileLogger) Log(message string) {
	f.logs = append(f.logs, message)
}

// TODO(human): Implement Close() method for FileLogger
func (f *FileLogger) Close() error {
	if f.closed {
		return errors.New("file is already closed")
	} else {
		f.closed = true
		return nil
	}
}

// GetLogs returns all logged messages
func (f *FileLogger) GetLogs() []string {
	// TODO(human): Implement
	if f.logs != nil {
		return f.logs
	}
	return []string{}
}
