package middleware_pattern

import "net/http"

// LoggingMiddleware logs the HTTP method and path
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return nil
}

// TimingMiddleware measures how long the handler takes
func TimingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return nil
}

// AuthMiddleware returns middleware that checks for an auth token
func AuthMiddleware(token string) func(http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return nil
}

// Chain applies multiple middleware to a handler
func Chain(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return nil
}
