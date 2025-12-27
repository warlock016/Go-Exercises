# Exercise 09: Middleware Basics

**Learning Goal:** Write middleware to wrap handlers with cross-cutting functionality

**Difficulty:** Tier 3 - Integration
**Estimated Time:** 35-40 minutes

## Problem Description

Build middleware functions that wrap http.Handler to add logging, timing, and header modification. Middleware is a powerful pattern for adding functionality to all handlers without modifying each one.

## Type Definition

```go
type Middleware func(http.Handler) http.Handler
```

## Function Signatures

```go
// LoggingMiddleware logs request method and path
func LoggingMiddleware(next http.Handler) http.Handler

// TimingMiddleware adds X-Response-Time header with duration in milliseconds
func TimingMiddleware(next http.Handler) http.Handler

// HeaderMiddleware adds custom header X-Custom-Header: test-value
func HeaderMiddleware(next http.Handler) http.Handler

// Chain combines multiple middleware functions
func Chain(h http.Handler, middlewares ...Middleware) http.Handler
```

## What This Teaches

- **Middleware pattern**: Wrapping handlers with additional behavior
- **Function closures**: Returning http.HandlerFunc from functions
- **Handler composition**: Chaining multiple middleware
- **Cross-cutting concerns**: Logging, timing, headers without repeating code
