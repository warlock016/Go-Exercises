# Exercise 13: Error Middleware

## 🎯 Learning Goal
Implement HTTP middleware for panic recovery, error logging, and converting panics to proper HTTP error responses.

## 📝 Problem Description

Production HTTP servers need to recover from panics gracefully and log errors properly. This exercise implements middleware patterns for robust error handling.

## 🔧 Function Signatures

```go
// RecoveryMiddleware wraps an HTTP handler to recover from panics
func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc

// ErrorLoggingMiddleware logs all errors
func ErrorLoggingMiddleware(next http.HandlerFunc) http.HandlerFunc

// CombinedMiddleware combines recovery and logging
func CombinedMiddleware(next http.HandlerFunc) http.HandlerFunc
```

## 🎓 What This Teaches

- **Panic recovery** - defer/recover in HTTP handlers
- **Middleware pattern** - Wrapping handlers
- **Production error handling** - Graceful degradation

---

**Next Exercise:** `14_weather_api_errors` - Apply error handling to Weather CLI
