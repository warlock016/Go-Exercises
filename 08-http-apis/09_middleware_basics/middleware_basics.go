package middleware_basics

import (
	"net/http"
)

// Middleware is a function that wraps an http.Handler
type Middleware func(http.Handler) http.Handler

// LoggingMiddleware logs request method and path
func LoggingMiddleware(next http.Handler) http.Handler {
	// TODO(human): Implement
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
	})
}

// TimingMiddleware adds X-Response-Time header
func TimingMiddleware(next http.Handler) http.Handler {
	// TODO(human): Implement
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Response-Time", "500")
		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
	})
}

// HeaderMiddleware adds custom header
func HeaderMiddleware(next http.Handler) http.Handler {
	// TODO(human): Implement
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Custom-Header", "test-value")
		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
	})
}

// Chain combines multiple middleware functions
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	// TODO(human): Implement
	for _, ware := range middlewares {
		h = ware(h)
	}
	return h
}
