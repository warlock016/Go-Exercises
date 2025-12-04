# Exercise 08: HTTP Handler Tests (🌐 HTTP)

**Learning Goal:** Test HTTP handlers using httptest.ResponseRecorder without starting a real server.

---

## Problem Description

Testing HTTP handlers doesn't require a real server. The `net/http/httptest` package provides `ResponseRecorder` to capture handler responses for testing.

---

## Handlers to Test

```go
func HelloHandler(w http.ResponseWriter, r *http.Request)
func JSONHandler(w http.ResponseWriter, r *http.Request)
func EchoHandler(w http.ResponseWriter, r *http.Request)
```

---

## Testing Pattern

```go
func TestHelloHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/hello", nil)
    rec := httptest.NewRecorder()

    HelloHandler(rec, req)

    if rec.Code != http.StatusOK {
        t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
    }

    if rec.Body.String() != "Hello, World!" {
        t.Errorf("body = %q, want %q", rec.Body.String(), "Hello, World!")
    }
}
```

---

## Your Task

Write tests for all handlers checking:
- Status code
- Response body
- Headers (Content-Type)
- Different request methods

---

**Next Exercise:** `09_testing_json_apis` - Testing JSON APIs
