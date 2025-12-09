package error_middleware

import (
	"log"
	"net/http"
)

// RecoveryMiddleware wraps an HTTP handler to recover from panics
func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// A panic happened! Write 500 response
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next(w, r) // Call the actual handler
	}
}

// ErrorLoggingMiddleware logs all errors
func ErrorLoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s, Path: %s, Headers: %v", r.Method, r.URL.Path, r.Header)
		next(w, r)
	}
}

// CombinedMiddleware combines recovery and logging
func CombinedMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	logging := ErrorLoggingMiddleware(next)
	recovery := RecoveryMiddleware(logging)
	return recovery
}
