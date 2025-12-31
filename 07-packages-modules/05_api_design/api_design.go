package apidesign

import "github.com/warlock016/go-exercises/07-packages-modules/05_api_design/logger"

// CreateLogger creates a logger with the given options
func CreateLogger(level, prefix string) any {
	// TODO(human): Import and use logger package
	// Return logger.NewLogger with appropriate options
	lvl := logger.WithLevel(level)
	opt := logger.WithPrefix(prefix)
	return logger.NewLogger(lvl, opt)
}

// LogMessage logs a message using a default logger
func LogMessage(message string) string {
	// TODO(human): Create logger and log the message

	logger := logger.NewLogger()
	return logger.Log(message)
}
