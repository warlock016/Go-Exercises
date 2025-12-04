# Exercise 13: Testing Middleware (🌐 HTTP)

**Learning Goal:** Test HTTP middleware that wraps handlers.

---

## Middleware to Test

```go
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc
```

---

## Testing Pattern

```go
func TestLoggingMiddleware(t *testing.T) {
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })

    wrapped := LoggingMiddleware(handler)

    req := httptest.NewRequest("GET", "/", nil)
    rec := httptest.NewRecorder()

    wrapped(rec, req)

    // Test that handler was called
    if rec.Body.String() != "OK" {
        t.Error("handler not called")
    }

    // Test that middleware did its job (logged, added headers, etc)
}
```

---

**Next Exercise:** `14_test_weather_client` - Comprehensive Weather CLI testing
