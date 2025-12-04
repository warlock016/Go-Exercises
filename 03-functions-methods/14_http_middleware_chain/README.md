# Exercise 14: HTTP Middleware Chain

**Learning Goal:** Build a production-ready middleware stack with recovery, logging, and CORS

---

## 📝 Problem Description

Production HTTP servers need multiple middleware layers for:
- Panic recovery
- Logging
- CORS headers
- Request timing
- Authentication

You'll build a complete middleware stack combining all previous patterns.

---

## 🎯 Function Signatures

```go
func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc

func CORSMiddleware(allowedOrigins []string) func(http.HandlerFunc) http.HandlerFunc

func RequestIDMiddleware(next http.HandlerFunc) http.HandlerFunc

func BuildMiddlewareStack(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc
```

---

## 📖 Examples

```go
handler := func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello"))
}

// Build complete stack
stack := BuildMiddlewareStack(
    handler,
    RecoveryMiddleware,
    LoggingMiddleware,
    TimingMiddleware,
    CORSMiddleware([]string{"*"}),
    RequestIDMiddleware,
)

// Stack handles panics, logs, times, adds CORS headers, and assigns request IDs
```

---

## 📋 Instructions

1. Implement `RecoveryMiddleware` to catch panics
2. Implement `CORSMiddleware` factory for CORS headers
3. Implement `RequestIDMiddleware` to add request IDs
4. Implement `BuildMiddlewareStack` to compose all middleware
5. Ensure middleware order is correct (recovery first!)
6. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Recovery Pattern</summary>

Use `defer` and `recover()`:
```go
defer func() {
    if err := recover(); err != nil {
        log.Printf("Panic: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
    }
}()
```

</details>

<details>
<summary>Complete Solution</summary>

```go
package http_middleware_chain

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type contextKey string

const requestIDKey contextKey = "requestID"

var requestIDCounter int

func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
			}
		}()
		next(w, r)
	}
}

func CORSMiddleware(allowedOrigins []string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			allowed := false

			for _, allowed Origin := range allowedOrigins {
				if allowedOrigin == "*" || allowedOrigin == origin {
					allowed = true
					break
				}
			}

			if allowed {
				if origin != "" {
					w.Header().Set("Access-Control-Allow-Origin", origin)
				} else if len(allowedOrigins) > 0 && allowedOrigins[0] == "*" {
					w.Header().Set("Access-Control-Allow-Origin", "*")
				}
			}

			next(w, r)
		}
	}
}

func RequestIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestIDCounter++
		requestID := fmt.Sprintf("req-%d", requestIDCounter)

		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		r = r.WithContext(ctx)

		w.Header().Set("X-Request-ID", requestID)

		next(w, r)
	}
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Context().Value(requestIDKey)
		log.Printf("[%v] %s %s", requestID, r.Method, r.URL.Path)
		next(w, r)
	}
}

func TimingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		duration := time.Since(start)
		log.Printf("Request took %v", duration)
	}
}

func BuildMiddlewareStack(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
```

</details>

---

## 🤔 Think About

1. Why should RecoveryMiddleware be first in the chain?
2. How does context propagate through middleware?
3. What are the performance implications of many middleware?
4. How would you add authentication middleware?

---

## 🎓 What This Teaches

- **Production patterns**: Real-world middleware stack
- **Panic recovery**: Preventing server crashes
- **CORS**: Cross-origin resource sharing
- **Request context**: Passing data through middleware
- **Middleware composition**: Building complex stacks

---

**Tier:** 4 - Mastery
**Estimated Time:** 50-60 minutes
