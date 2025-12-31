package context_timeout

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// SlowOperation simulates slow work that respects context
func SlowOperation(ctx context.Context, duration time.Duration) error {
	// TODO(human): Implement

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(duration):
		return nil
	}
}

// SlowHandler handles requests with slow operations
func SlowHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement

	dur := r.URL.Query().Get("duration")
	t, err := time.ParseDuration(dur)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid duration param"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = SlowOperation(r.Context(), t)
	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		w.Write([]byte("op timed out"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}
}

// TimeoutMiddleware returns 504 if handler exceeds timeout
func TimeoutMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
	// TODO(human): Implement
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			done := make(chan struct{})
			go func() {
				h.ServeHTTP(w, r.WithContext(ctx))
				close(done)
			}()

			select {
			case <-done:
				// The handler already wrote the response, nothing to do! (reason for commenting out next line)
			case <-ctx.Done():
				w.WriteHeader(http.StatusGatewayTimeout)
			}
		})
	}
}

// RequestIDMiddleware adds request ID to context
func RequestIDMiddleware(next http.Handler) http.Handler {
	// TODO(human): Implement
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		ctx := context.WithValue(r.Context(), "requestID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
