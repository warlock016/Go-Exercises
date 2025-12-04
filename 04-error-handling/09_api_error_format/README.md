# Exercise 09: API Error Format (RFC 7807)

## 🎯 Learning Goal
Implement RFC 7807 Problem Details format for consistent, standardized API error responses.

## 📝 Problem Description

RFC 7807 defines a standard format for HTTP API error responses. This exercise implements that format.

## 🔧 Type and Function Signatures

```go
// ProblemDetail represents RFC 7807 Problem Details
type ProblemDetail struct {
    Type     string `json:"type"`
    Title    string `json:"title"`
    Status   int    `json:"status"`
    Detail   string `json:"detail"`
    Instance string `json:"instance"`
}

// NewProblemDetail creates a problem detail from an error
func NewProblemDetail(err error, requestPath string) ProblemDetail

// ValidationProblemDetail creates a problem detail for validation errors
func ValidationProblemDetail(errors []string, requestPath string) ProblemDetail
```

## 🎓 What This Teaches

- **RFC 7807** - Industry standard for API errors
- **Consistent error format** - Predictable error structure
- **Error documentation** - Type URLs for error documentation

---

**Next Exercise:** `10_retry_with_backoff` - Retry logic with exponential backoff
