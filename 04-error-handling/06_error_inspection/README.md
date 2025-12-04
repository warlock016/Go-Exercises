# Exercise 06: Error Inspection

## 🎯 Learning Goal
Learn to inspect error chains using `errors.Is` for sentinel error checking and `errors.As` for extracting custom error types, even when errors are wrapped multiple times.

## 📝 Problem Description

When errors are wrapped, you need special functions to inspect them:
- `errors.Is(err, target)` - Checks if any error in the chain matches target
- `errors.As(err, &target)` - Extracts a specific error type from the chain

These work with wrapped errors created using `%w`.

## 🔧 Function Signatures

```go
// CheckErrorType inspects an error and returns what type it is
func CheckErrorType(err error) string

// ExtractHTTPStatus extracts the HTTP status code from a wrapped HTTPError
func ExtractHTTPStatus(err error) (int, bool)

// IsTemporaryError checks if error is a temporary/transient error
func IsTemporaryError(err error) bool

// HandleBasedOnType performs different actions based on error type
func HandleBasedOnType(err error) string
```

## 💡 Examples

```go
err := fmt.Errorf("database operation failed: %w", sql.ErrNoRows)

// errors.Is checks the entire chain
if errors.Is(err, sql.ErrNoRows) {
    fmt.Println("No rows found")  // This works!
}

// errors.As extracts custom types
var httpErr *HTTPError
if errors.As(err, &httpErr) {
    fmt.Printf("Status: %d\n", httpErr.StatusCode)
}
```

## 🎓 What This Teaches

- **errors.Is** - Checking sentinel errors in wrapped chains
- **errors.As** - Extracting custom error types
- **Error chain traversal** - How Go walks the error chain
- **Type-based error handling** - Making decisions based on error types

---

**Next Exercise:** `07_validation_errors` - Aggregating multiple validation errors
