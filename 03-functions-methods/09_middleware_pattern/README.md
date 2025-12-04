# Exercise 09: Middleware Pattern

**Learning Goal:** Master chainable middleware for wrapping HTTP handlers

---

## 📝 Problem Description

Middleware wraps handlers to add cross-cutting concerns like logging, timing, authentication, etc. Middleware functions take a handler and return a new handler with added behavior.

Pattern: `middleware(handler) -> wrappedHandler`

---

## 🎯 Function Signatures

```go
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc

func TimingMiddleware(next http.HandlerFunc) http.HandlerFunc

func AuthMiddleware(token string) func(http.HandlerFunc) http.HandlerFunc

func Chain(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc
```

---

## 📖 Examples

```go
// Single middleware
handler := func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello"))
}
logged := LoggingMiddleware(handler)

// Chaining middleware
timed := TimingMiddleware(logged)

// Using Chain helper
final := Chain(handler,
    LoggingMiddleware,
    TimingMiddleware,
)

// Middleware with configuration
auth := AuthMiddleware("secret-token")
protected := auth(handler)
```

---

## 📋 Instructions

1. Implement `LoggingMiddleware` that logs method and path
2. Implement `TimingMiddleware` that measures handler duration
3. Implement `AuthMiddleware` factory that checks for auth token
4. Implement `Chain` to compose multiple middleware
5. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Basic Concept</summary>

Middleware wraps a handler:
```go
func Middleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Before handler
        next(w, r)  // Call wrapped handler
        // After handler
    }
}
```

</details>

<details>
<summary>Complete Solution</summary>

```go
package middleware_pattern

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
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

func AuthMiddleware(token string) func(http.HandlerFunc) http.HandlerFunc {
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

func Chain(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
```

</details>

---

## 🤔 Think About

1. Why apply middleware in reverse order in `Chain`?
2. How does middleware enable separation of concerns?
3. What are common use cases for middleware?
4. How would you pass data between middleware layers?

---

## 🎓 What This Teaches

- **Middleware pattern**: Wrapping handlers with additional behavior
- **Function composition**: Layering functionality
- **Cross-cutting concerns**: Logging, timing, auth applied uniformly
- **Decorator pattern**: Adding behavior without modifying original
- **Chain of responsibility**: Processing requests through multiple handlers

---

**Tier:** 3 - Integration
**Estimated Time:** 45-60 minutes
