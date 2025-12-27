package context_timeout

import (
	"context"
	"net/http"
	"time"
)

// SlowOperation simulates slow work that respects context
func SlowOperation(ctx context.Context, duration time.Duration) error {
	// TODO(human): Implement
	return nil
}

// SlowHandler handles requests with slow operations
func SlowHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
}

// TimeoutMiddleware returns 504 if handler exceeds timeout
func TimeoutMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
	// TODO(human): Implement
	return nil
}

// RequestIDMiddleware adds request ID to context
func RequestIDMiddleware(next http.Handler) http.Handler {
	// TODO(human): Implement
	return nil
}
