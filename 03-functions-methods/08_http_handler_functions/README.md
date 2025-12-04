# Exercise 08: HTTP Handler Functions

**Learning Goal:** Master HTTP handlers and the `http.HandlerFunc` pattern

---

## 📝 Problem Description

Go's `net/http` package uses functions extensively for handling HTTP requests. You'll learn:
- The `http.Handler` interface
- The `http.HandlerFunc` adapter
- Writing handler functions
- Extracting request data
- Writing responses

---

## 🎯 Function Signatures

```go
func HelloHandler(w http.ResponseWriter, r *http.Request)

func EchoHandler(w http.ResponseWriter, r *http.Request)

func StatusHandler(code int) http.HandlerFunc

func JSONHandler(w http.ResponseWriter, r *http.Request)

func MethodFilter(method string, handler http.HandlerFunc) http.HandlerFunc
```

---

## 📖 Examples

```go
// HelloHandler - responds with "Hello, World!"
// GET / -> "Hello, World!"

// EchoHandler - echoes query parameter "message"
// GET /echo?message=test -> "test"

// StatusHandler - returns specific status code
// handler := StatusHandler(404)
// GET / -> 404 Not Found

// JSONHandler - returns JSON response
// GET / -> {"message": "Hello, JSON!", "timestamp": 123456789}

// MethodFilter - only allows specific HTTP method
// handler := MethodFilter("POST", someHandler)
// GET / -> 405 Method Not Allowed
// POST / -> calls someHandler
```

---

## 📋 Instructions

1. Implement `HelloHandler` that writes "Hello, World!"
2. Implement `EchoHandler` that echoes the "message" query parameter
3. Implement `StatusHandler` factory that returns a handler with a status code
4. Implement `JSONHandler` that responds with JSON
5. Implement `MethodFilter` that wraps handlers with method checking
6. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Basic Concept</summary>

HTTP handlers have the signature:
```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // Read request
    // Write response
    w.Write([]byte("response"))
}
```

Query parameters: `r.URL.Query().Get("key")`
Write response: `w.Write([]byte("text"))`
Set status: `w.WriteHeader(http.StatusOK)`

</details>

<details>
<summary>Complete Solution</summary>

```go
package http_handler_functions

import (
	"encoding/json"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	if message == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing message parameter"))
		return
	}
	w.Write([]byte(message))
}

func StatusHandler(code int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
	}
}

func JSONHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message":   "Hello, JSON!",
		"timestamp": time.Now().Unix(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func MethodFilter(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}
```

</details>

---

## 🤔 Think About

1. What's the difference between `http.Handler` and `http.HandlerFunc`?
2. Why does `http.HandlerFunc` exist as a type?
3. How does `MethodFilter` demonstrate function composition?
4. What order should you call `WriteHeader` and `Write`?

---

## 🎓 What This Teaches

- **HTTP handlers**: Standard Go HTTP handler pattern
- **Function adapters**: `http.HandlerFunc` type conversion
- **Handler factories**: Creating configured handlers
- **Middleware pattern**: Wrapping handlers with additional behavior
- **Request/Response**: Reading requests and writing responses

---

**Tier:** 2 - Application
**Estimated Time:** 40-50 minutes
