# Exercise 06: Error Responses

**Learning Goal:** Build structured error responses with consistent format

**Difficulty:** Tier 2 - Application
**Estimated Time:** 30-35 minutes

## Problem Description

Create reusable error response structures and helper functions that return consistent JSON error messages. You'll learn to build an error response system that makes debugging easier for API consumers.

## Type Definitions

```go
type APIError struct {
    Code    string                 `json:"code"`
    Message string                 `json:"message"`
    Details map[string]string      `json:"details,omitempty"`
}
```

## Function Signatures

```go
// WriteError writes a JSON error response with given status code
func WriteError(w http.ResponseWriter, status int, code, message string)

// WriteErrorWithDetails writes a JSON error with additional details map
func WriteErrorWithDetails(w http.ResponseWriter, status int, code, message string, details map[string]string)

// ValidationHandler demonstrates various error responses
// Returns 400 for missing "email" query param
// Returns 422 for invalid "age" (must be >= 0 and <= 150)
// Returns 200 with {"status":"valid"} if all valid
func ValidationHandler(w http.ResponseWriter, r *http.Request)
```

## Examples

### ValidationHandler - Success
```
Request:  GET /validate?email=alice@example.com&age=25
Response: 200 OK
Body:     {"status":"valid"}
```

### ValidationHandler - Missing Email
```
Request:  GET /validate?age=25
Response: 400 Bad Request
Body:     {
            "code": "MISSING_PARAM",
            "message": "email parameter is required"
          }
```

### ValidationHandler - Invalid Age
```
Request:  GET /validate?email=alice@example.com&age=200
Response: 422 Unprocessable Entity
Body:     {
            "code": "VALIDATION_ERROR",
            "message": "age must be between 0 and 150",
            "details": {
              "field": "age",
              "value": "200"
            }
          }
```

## Instructions

1. Implement `WriteError` to write JSON error response
2. Implement `WriteErrorWithDetails` to include details map
3. Implement `ValidationHandler` with proper error codes
4. Run tests with `go test -v`

## Hints

**Basic:**
- Create APIError struct and encode as JSON
- Set Content-Type to "application/json"
- Use `w.WriteHeader(status)` before encoding

**Intermediate:**
- `omitempty` in JSON tag means omit if nil/empty
- Convert age string to int with `strconv.Atoi()`
- Use descriptive error codes like "MISSING_PARAM", "VALIDATION_ERROR"

## What This Teaches

- **Structured errors**: Consistent JSON error format
- **Error codes**: Machine-readable error identification
- **Error details**: Additional context for debugging
- **Validation patterns**: 400 vs 422 status codes
- **Reusable helpers**: DRY error handling functions
