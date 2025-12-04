# Exercise 04: Custom Error Types

## 🎯 Learning Goal
Learn to create custom error types by implementing the `error` interface on structs, enabling errors that carry additional context and data beyond simple strings.

## 📝 Problem Description

Sometimes sentinel errors and simple strings aren't enough. Custom error types let you attach structured data to errors (like field names, values, error codes, timestamps) while still implementing the standard `error` interface.

Any type that implements `Error() string` satisfies the error interface and can be used as an error.

## 🔧 Type Definitions and Functions

Define these custom error types:

```go
// ValidationError represents a validation failure with field details
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

// Implement error interface
func (e *ValidationError) Error() string

// HTTPError represents an HTTP error with status code
type HTTPError struct {
    StatusCode int
    Message    string
    URL        string
}

// Implement error interface
func (e *HTTPError) Error() string

// TimeoutError represents a timeout with duration info
type TimeoutError struct {
    Operation string
    Duration  time.Duration
}

// Implement error interface
func (e *TimeoutError) Error() string
```

Implement these functions:

```go
// ValidateUserInput checks user struct fields and returns ValidationError if invalid
func ValidateUserInput(name, email string, age int) error

// FetchResource simulates an HTTP request, returns HTTPError for bad status codes
func FetchResource(url string, statusCode int) (string, error)

// PerformOperation simulates an operation with timeout
func PerformOperation(name string, timeoutSeconds int) error

// ExtractValidationError attempts to extract ValidationError from error
func ExtractValidationError(err error) (*ValidationError, bool)
```

## 💡 Examples

```go
// Validation error with field details
err := ValidateUserInput("", "test@example.com", 25)
// Returns: &ValidationError{Field: "name", Value: "", Message: "cannot be empty"}

if verr, ok := ExtractValidationError(err); ok {
    fmt.Printf("Field %s failed: %s\n", verr.Field, verr.Message)
}

// HTTP error with status code
data, err := FetchResource("https://api.example.com/users", 404)
// Returns: &HTTPError{StatusCode: 404, Message: "Not Found", URL: "https://api.example.com/users"}

if herr, ok := err.(*HTTPError); ok {
    fmt.Printf("HTTP %d: %s\n", herr.StatusCode, herr.Message)
}

// Timeout error with duration
err = PerformOperation("database-query", 30)
// Returns: &TimeoutError{Operation: "database-query", Duration: 30 * time.Second}

if terr, ok := err.(*TimeoutError); ok {
    fmt.Printf("%s timed out after %v\n", terr.Operation, terr.Duration)
}
```

## 📋 Instructions

1. **Define ValidationError struct** with Field, Value, Message fields
2. **Implement Error() method** on *ValidationError that formats the error message nicely
3. **Define HTTPError struct** with StatusCode, Message, URL
4. **Implement Error() method** on *HTTPError
5. **Define TimeoutError struct** with Operation and Duration (time.Duration)
6. **Implement Error() method** on *TimeoutError
7. **ValidateUserInput:** Check name (not empty), email (contains @), age (0-150). Return appropriate ValidationError for first failing field.
8. **FetchResource:** Return HTTPError if statusCode >= 400, otherwise return "success"
9. **PerformOperation:** Return TimeoutError if timeoutSeconds > 30
10. **ExtractValidationError:** Use type assertion to extract ValidationError from error

## 🧪 Testing

```bash
go test -v
```

Expected: ~25-30 test cases

## 🤔 Think About

1. **Why pointer receivers for Error()?**
   - Allows modifying the error after creation
   - Consistent with how errors are typically used

2. **When to use custom types vs sentinel errors?**
   - Custom types: Need additional data (field names, codes, values)
   - Sentinel errors: Simple, well-known conditions

3. **How do custom types work with errors.Is/As?**
   - errors.As can extract custom types from wrapped errors
   - errors.Is works if you implement Is() method

## 💡 Hints

<details>
<summary>Hint 1: Implementing the error interface</summary>

Any type with an `Error() string` method is an error:
```go
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s (got %v)", e.Field, e.Message, e.Value)
}
```

Use pointer receiver: `(e *ValidationError)` not `(e ValidationError)`
</details>

<details>
<summary>Hint 2: Creating and returning custom errors</summary>

Return a pointer to your error struct:
```go
func ValidateUserInput(name, email string, age int) error {
    if name == "" {
        return &ValidationError{
            Field:   "name",
            Value:   name,
            Message: "cannot be empty",
        }
    }
    // ... more checks
    return nil
}
```
</details>

<details>
<summary>Hint 3: Type assertions for custom errors</summary>

Use type assertion or errors.As:
```go
func ExtractValidationError(err error) (*ValidationError, bool) {
    verr, ok := err.(*ValidationError)
    return verr, ok
}
```

Or with errors.As:
```go
func ExtractValidationError(err error) (*ValidationError, bool) {
    var verr *ValidationError
    if errors.As(err, &verr) {
        return verr, true
    }
    return nil, false
}
```
</details>

<details>
<summary>Full Solution</summary>

```go
package custom_error_types

import (
    "errors"
    "fmt"
    "strings"
    "time"
)

type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s (got %v)", e.Field, e.Message, e.Value)
}

type HTTPError struct {
    StatusCode int
    Message    string
    URL        string
}

func (e *HTTPError) Error() string {
    return fmt.Sprintf("HTTP %d: %s (URL: %s)", e.StatusCode, e.Message, e.URL)
}

type TimeoutError struct {
    Operation string
    Duration  time.Duration
}

func (e *TimeoutError) Error() string {
    return fmt.Sprintf("operation '%s' timed out after %v", e.Operation, e.Duration)
}

func ValidateUserInput(name, email string, age int) error {
    if name == "" {
        return &ValidationError{Field: "name", Value: name, Message: "cannot be empty"}
    }
    if !strings.Contains(email, "@") {
        return &ValidationError{Field: "email", Value: email, Message: "must contain @"}
    }
    if age < 0 || age > 150 {
        return &ValidationError{Field: "age", Value: age, Message: "must be between 0 and 150"}
    }
    return nil
}

func FetchResource(url string, statusCode int) (string, error) {
    if statusCode >= 400 {
        return "", &HTTPError{
            StatusCode: statusCode,
            Message:    "Request failed",
            URL:        url,
        }
    }
    return "success", nil
}

func PerformOperation(name string, timeoutSeconds int) error {
    if timeoutSeconds > 30 {
        return &TimeoutError{
            Operation: name,
            Duration:  time.Duration(timeoutSeconds) * time.Second,
        }
    }
    return nil
}

func ExtractValidationError(err error) (*ValidationError, bool) {
    var verr *ValidationError
    if errors.As(err, &verr) {
        return verr, true
    }
    return nil, false
}
```
</details>

## 🎓 What This Teaches

- **error interface** - Any type with `Error() string` is an error
- **Custom error types** - Structs that carry additional context
- **Pointer receivers** - Why `(e *ValidationError)` vs `(e ValidationError)`
- **Type assertions** - Extracting custom types with `err.(*CustomType)`
- **errors.As** - Extracting custom types from error chains
- **Structured error data** - Field names, codes, URLs, durations
- **Error formatting** - Creating helpful Error() string representations

---

**Next Exercise:** `05_error_wrapping` - Using %w to create error chains
