package http_middleware_chain

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type contextKey string

const requestIDKey contextKey = "requestID"

// RecoveryMiddleware recovers from panics
func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil { // recover is available only inside of defered functions
				log.Printf("Panic recovered: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal server error"))
			}
		}()
		next(w, r)
	}
}

// CORSMiddleware adds CORS headers for allowed origins
func CORSMiddleware(allowedOrigins []string) func(http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			allowed := false

			for _, allowedOrigin := range allowedOrigins {
				if allowedOrigin == "*" || allowedOrigin == origin {
					allowed = true
					break
				}
			}

			if allowed {
				if len(allowedOrigins) > 0 && allowedOrigins[0] == "*" {
					w.Header().Set("Access-Control-Allow-Origin", "*")
				} else if origin != "" {
					w.Header().Set("Access-Control-Allow-Origin", origin)
				}

			}
			next(w, r)
		}
	}
}

// RequestIDMiddleware assigns a unique ID to each request
func RequestIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	counter := 0
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := fmt.Sprintf("req-%d", counter)
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		r = r.WithContext(ctx)
		w.Header().Set("X-Request-ID", requestID)
		counter++
		next(w, r)
	}
}

// LoggingMiddleware logs request details
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Context().Value(requestIDKey)
		log.Printf("[%v] %s %s", requestID, r.Method, r.URL.Path)
		next(w, r)
	}
}

// TimingMiddleware measures request duration
func TimingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		duration := time.Since(start)
		log.Printf("request took %v", duration)
	}
}

// BuildMiddlewareStack composes multiple middleware
func BuildMiddlewareStack(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return func(w http.ResponseWriter, r *http.Request) {
		for i := len(middlewares) - 1; i >= 0; i-- {
			handler = middlewares[i](handler)
		}
		handler(w, r)
	}
}
