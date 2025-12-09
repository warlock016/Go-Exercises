package middleware_pattern

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the HTTP method and path
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next(w, r)
	}
}

// TimingMiddleware measures how long the handler takes
func TimingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		duration := time.Since(start)
		log.Printf("Request took %v", duration)
	}
}

// AuthMiddleware returns middleware that checks for an auth token
func AuthMiddleware(token string) func(http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth != "Bearer "+token {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next(w, r)
		}
	}
}

// Chain applies multiple middleware to a handler
func Chain(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement

	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
