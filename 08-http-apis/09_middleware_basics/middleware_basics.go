package middleware_basics

import "net/http"

// Middleware is a function that wraps an http.Handler
type Middleware func(http.Handler) http.Handler

// LoggingMiddleware logs request method and path
func LoggingMiddleware(next http.Handler) http.Handler {
	// TODO(human): Implement
	return nil
}

// TimingMiddleware adds X-Response-Time header
func TimingMiddleware(next http.Handler) http.Handler {
	// TODO(human): Implement
	return nil
}

// HeaderMiddleware adds custom header
func HeaderMiddleware(next http.Handler) http.Handler {
	// TODO(human): Implement
	return nil
}

// Chain combines multiple middleware functions
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	// TODO(human): Implement
	return nil
}
