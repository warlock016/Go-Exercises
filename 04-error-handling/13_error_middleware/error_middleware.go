package error_middleware

import "net/http"

// RecoveryMiddleware wraps an HTTP handler to recover from panics
func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return nil
}

// ErrorLoggingMiddleware logs all errors
func ErrorLoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return nil
}

// CombinedMiddleware combines recovery and logging
func CombinedMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return nil
}
