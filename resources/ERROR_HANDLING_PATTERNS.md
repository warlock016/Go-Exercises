# Error Handling Patterns Guide

A comprehensive reference for Go error handling idioms and patterns.

---

## The error Interface

```go
// The built-in error interface
type error interface {
    Error() string
}

// Any type with an Error() string method satisfies error
```

---

## Creating Errors

### errors.New

```go
import "errors"

var ErrNotFound = errors.New("not found")

func findUser(id int) (*User, error) {
    if id <= 0 {
        return nil, errors.New("invalid user id")
    }
    // ...
}
```

### fmt.Errorf

```go
import "fmt"

func openFile(path string) error {
    if path == "" {
        return fmt.Errorf("path cannot be empty")
    }
    // With formatting
    return fmt.Errorf("file %q not found", path)
}
```

---

## Error Checking Patterns

### Basic Pattern

```go
result, err := doSomething()
if err != nil {
    return err  // Propagate error
}
// Use result
```

### Early Return

```go
func processData(data []byte) error {
    if len(data) == 0 {
        return errors.New("empty data")
    }

    parsed, err := parse(data)
    if err != nil {
        return err
    }

    if err := validate(parsed); err != nil {
        return err
    }

    return save(parsed)
}
```

### Inline Error Declaration

```go
if err := doSomething(); err != nil {
    return err
}
```

---

## Sentinel Errors

### Definition

```go
// Package-level error variables
var (
    ErrNotFound     = errors.New("not found")
    ErrUnauthorized = errors.New("unauthorized")
    ErrInvalidInput = errors.New("invalid input")
)
```

### Checking Sentinel Errors

```go
// Pre-Go 1.13: direct comparison (fragile)
if err == ErrNotFound {
    // handle not found
}

// Go 1.13+: errors.Is (handles wrapped errors)
if errors.Is(err, ErrNotFound) {
    // handle not found
}
```

---

## Custom Error Types

### Basic Custom Error

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

// Usage
return &ValidationError{Field: "email", Message: "invalid format"}
```

### Error with Context

```go
type QueryError struct {
    Query string
    Err   error
}

func (e *QueryError) Error() string {
    return fmt.Sprintf("query %q failed: %v", e.Query, e.Err)
}

func (e *QueryError) Unwrap() error {
    return e.Err  // Allows errors.Is and errors.As to work
}
```

---

## Error Wrapping (Go 1.13+)

### Wrapping with %w

```go
func readConfig(path string) error {
    data, err := os.ReadFile(path)
    if err != nil {
        // Wrap the error with context
        return fmt.Errorf("reading config %s: %w", path, err)
    }
    // ...
}

// Creates an error chain:
// "reading config /etc/app.conf: open /etc/app.conf: permission denied"
```

### Unwrapping Errors

```go
// errors.Unwrap - get the wrapped error
wrapped := fmt.Errorf("outer: %w", innerErr)
inner := errors.Unwrap(wrapped)  // Returns innerErr

// errors.Is - check if any error in chain matches
if errors.Is(err, os.ErrNotExist) {
    // File doesn't exist (even if wrapped multiple times)
}

// errors.As - extract specific error type from chain
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println("Path:", pathErr.Path)
}
```

---

## Multi-Error Aggregation

### Validation Errors

```go
type ValidationErrors struct {
    Errors []error
}

func (v *ValidationErrors) Error() string {
    if len(v.Errors) == 0 {
        return "no errors"
    }

    var msgs []string
    for _, err := range v.Errors {
        msgs = append(msgs, err.Error())
    }
    return strings.Join(msgs, "; ")
}

func (v *ValidationErrors) Add(err error) {
    v.Errors = append(v.Errors, err)
}

func (v *ValidationErrors) HasErrors() bool {
    return len(v.Errors) > 0
}

// Usage
func validateUser(u User) error {
    errs := &ValidationErrors{}

    if u.Name == "" {
        errs.Add(errors.New("name is required"))
    }
    if u.Email == "" {
        errs.Add(errors.New("email is required"))
    }

    if errs.HasErrors() {
        return errs
    }
    return nil
}
```

---

## HTTP Error Responses

### Error to Status Code Mapping

```go
func errorToStatusCode(err error) int {
    switch {
    case errors.Is(err, ErrNotFound):
        return http.StatusNotFound  // 404
    case errors.Is(err, ErrUnauthorized):
        return http.StatusUnauthorized  // 401
    case errors.Is(err, ErrInvalidInput):
        return http.StatusBadRequest  // 400
    default:
        return http.StatusInternalServerError  // 500
    }
}
```

### JSON Error Response

```go
type APIError struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Details any    `json:"details,omitempty"`
}

func writeError(w http.ResponseWriter, status int, err error) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)

    apiErr := APIError{
        Code:    http.StatusText(status),
        Message: err.Error(),
    }
    json.NewEncoder(w).Encode(apiErr)
}
```

### RFC 7807 Problem Details

```go
type ProblemDetails struct {
    Type     string `json:"type"`
    Title    string `json:"title"`
    Status   int    `json:"status"`
    Detail   string `json:"detail,omitempty"`
    Instance string `json:"instance,omitempty"`
}

func writeProblem(w http.ResponseWriter, status int, title, detail string) {
    w.Header().Set("Content-Type", "application/problem+json")
    w.WriteHeader(status)

    problem := ProblemDetails{
        Type:   "about:blank",
        Title:  title,
        Status: status,
        Detail: detail,
    }
    json.NewEncoder(w).Encode(problem)
}
```

---

## Retry with Backoff

```go
func retry(attempts int, sleep time.Duration, fn func() error) error {
    var err error
    for i := 0; i < attempts; i++ {
        err = fn()
        if err == nil {
            return nil
        }

        // Check if error is retryable
        if !isRetryable(err) {
            return err
        }

        time.Sleep(sleep)
        sleep *= 2  // Exponential backoff
    }
    return fmt.Errorf("after %d attempts: %w", attempts, err)
}

func isRetryable(err error) bool {
    // Network timeouts, temporary failures, etc.
    var netErr net.Error
    if errors.As(err, &netErr) && netErr.Temporary() {
        return true
    }
    return false
}
```

---

## Panic Recovery Middleware

```go
func recoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("panic recovered: %v\n%s", err, debug.Stack())

                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(map[string]string{
                    "error": "internal server error",
                })
            }
        }()

        next.ServeHTTP(w, r)
    })
}
```

---

## Error Logging

```go
type ErrorContext struct {
    Err       error
    Operation string
    UserID    string
    RequestID string
}

func logError(ctx ErrorContext) {
    log.Printf(
        "[ERROR] op=%s user=%s req=%s err=%v",
        ctx.Operation,
        ctx.UserID,
        ctx.RequestID,
        ctx.Err,
    )
}

// Structured logging (with slog in Go 1.21+)
import "log/slog"

slog.Error("operation failed",
    "operation", "fetch_user",
    "user_id", userID,
    "error", err,
)
```

---

## Best Practices

### Do

1. **Always check errors** - Don't ignore returned errors
2. **Add context when wrapping** - Use `fmt.Errorf("context: %w", err)`
3. **Use sentinel errors for expected conditions** - `ErrNotFound`
4. **Use custom types for rich errors** - When you need fields
5. **Return errors, don't panic** - Except for truly unrecoverable states
6. **Log at the top, not throughout** - Log once, at the handler level

### Don't

1. **Don't use `errors.New` in return statements** - Create sentinels
2. **Don't compare error strings** - Use `errors.Is` or `errors.As`
3. **Don't wrap errors without context** - `return fmt.Errorf("%w", err)` is pointless
4. **Don't panic in library code** - Return errors instead
5. **Don't swallow errors** - At minimum, log them

---

## Error Handling Decision Tree

```
Is it a programming error? → panic (rare)
                  ↓ No
Is it expected? → sentinel error (ErrNotFound)
       ↓ No
Need extra info? → custom error type
       ↓ No
Just propagate → fmt.Errorf("context: %w", err)
```

---

## Resources

- [Go Blog - Error Handling](https://go.dev/blog/error-handling-and-go)
- [Go Blog - Errors are Values](https://go.dev/blog/errors-are-values)
- [Go Blog - Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- [Effective Go - Errors](https://go.dev/doc/effective_go#errors)
