# Exercise 08: HTTP Error Responses

## 🎯 Learning Goal
Learn to map domain errors to HTTP status codes and create JSON error responses for REST APIs.

## 📝 Problem Description

Web APIs need to translate internal errors to HTTP status codes and JSON responses. This exercise teaches the mapping between error types and HTTP semantics.

## 🔧 Function Signatures

```go
// ErrorToStatusCode maps errors to HTTP status codes
func ErrorToStatusCode(err error) int

// ErrorResponse creates a JSON-serializable error response
type ErrorResponse struct {
    Status  int    `json:"status"`
    Message string `json:"message"`
    Code    string `json:"code"`
}

// NewErrorResponse creates an ErrorResponse from an error
func NewErrorResponse(err error) ErrorResponse

// WriteErrorResponse simulates writing JSON error response
func WriteErrorResponse(err error) string
```

## 💡 Examples

```go
status := ErrorToStatusCode(ErrNotFound)    // 404
status = ErrorToStatusCode(ErrUnauthorized) // 401
status = ErrorToStatusCode(ErrInvalidInput) // 400

resp := NewErrorResponse(ErrNotFound)
// {Status: 404, Message: "Resource not found", Code: "NOT_FOUND"}
```

## 🎓 What This Teaches

- **HTTP status codes** - 4xx for client errors, 5xx for server errors
- **Error mapping** - Translating domain errors to HTTP semantics
- **JSON error responses** - Structured error API responses

---

**Next Exercise:** `09_api_error_format` - RFC 7807 Problem Details format
