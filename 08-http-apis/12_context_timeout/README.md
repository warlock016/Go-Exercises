# Exercise 12: Context Timeout

**Learning Goal:** Use context for request cancellation and timeouts

**Difficulty:** Tier 3 - Integration
**Estimated Time:** 40-45 minutes

## Problem Description

Build handlers that respect context cancellation and timeouts. You'll learn to use context.Context to handle request timeouts gracefully and avoid doing unnecessary work when clients disconnect.

Context is crucial for production servers - it prevents wasted resources on abandoned requests.

## Function Signatures

```go
// SlowOperation simulates slow work that respects context cancellation
// Returns error if context is cancelled before duration elapses
func SlowOperation(ctx context.Context, duration time.Duration) error

// SlowHandler uses SlowOperation and returns 500 if context cancelled
func SlowHandler(w http.ResponseWriter, r *http.Request)

// TimeoutMiddleware wraps handler with timeout, returns 504 if exceeded
func TimeoutMiddleware(timeout time.Duration) func(http.Handler) http.Handler

// RequestIDMiddleware adds request ID to context
func RequestIDMiddleware(next http.Handler) http.Handler
```

## Examples

### SlowHandler
```
Request:  GET /slow?duration=100ms
Response: 200 OK (after 100ms)
Body:     {"status":"completed"}

Request:  GET /slow?duration=5s (with 1s timeout middleware)
Response: 504 Gateway Timeout
Body:     request timeout
```

### RequestIDMiddleware
```
Adds request ID to context, accessible via:
requestID := r.Context().Value("requestID").(string)
```

## What This Teaches

- **Context cancellation**: Stopping work when client disconnects
- **Request timeouts**: Preventing slow requests from blocking server
- **Context values**: Passing request-scoped data
- **504 Gateway Timeout**: Proper status for timeout errors
- **Graceful shutdown**: Respecting cancellation signals
- **Production patterns**: Essential for reliable services
