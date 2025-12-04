# Exercise 05: Testing Errors

**Learning Goal:** Learn comprehensive error testing patterns - checking error types, messages, and error wrapping.

---

## Problem Description

Errors are first-class citizens in Go. Testing error conditions is as important as testing success cases. You need to verify:
- That errors are returned when expected
- Error messages are helpful
- Error types are correct (when using custom errors)
- Errors are properly wrapped with context

---

## Functions to Test

```go
func ParseInt(s string) (int, error)
func Divide(a, b float64) (float64, error)
func ReadFile(path string) (string, error)
```

---

## Error Testing Patterns

### Pattern 1: Check Error Exists

```go
_, err := FunctionThatErrors()
if err == nil {
    t.Error("expected error, got nil")
}
```

### Pattern 2: Check Error Message

```go
_, err := Divide(10, 0)
if err == nil {
    t.Fatal("expected error")
}
if !strings.Contains(err.Error(), "division by zero") {
    t.Errorf("unexpected error message: %v", err)
}
```

### Pattern 3: Check Error Type (Custom Errors)

```go
var validationErr *ValidationError
if !errors.As(err, &validationErr) {
    t.Errorf("expected ValidationError, got %T", err)
}
```

### Pattern 4: Check Sentinel Error

```go
if !errors.Is(err, ErrNotFound) {
    t.Errorf("expected ErrNotFound, got %v", err)
}
```

---

## Your Task

Write comprehensive tests for functions that return errors. Test both success and failure cases.

---

## Key Concepts

### errors.Is() - Check Sentinel Errors

```go
var ErrNotFound = errors.New("not found")

func FindUser(id int) error {
    if id < 0 {
        return ErrNotFound
    }
    return nil
}

// Test
err := FindUser(-1)
if !errors.Is(err, ErrNotFound) {
    t.Error("expected ErrNotFound")
}
```

### errors.As() - Check Error Type

```go
type ValidationError struct {
    Field string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed: %s", e.Field)
}

// Test
var validErr *ValidationError
if !errors.As(err, &validErr) {
    t.Error("expected ValidationError")
}
```

---

## Instructions

1. Implement tests for all functions
2. Test both success and error cases
3. Check error messages contain expected text
4. Use table-driven tests with subtests
5. Create a helper for error message checking

---

## Hints

<details>
<summary>Hint 1: Error Message Helper</summary>

```go
func assertErrorContains(t *testing.T, err error, substr string) {
    t.Helper()
    if err == nil {
        t.Fatal("expected error, got nil")
    }
    if !strings.Contains(err.Error(), substr) {
        t.Errorf("error %q does not contain %q", err.Error(), substr)
    }
}
```
</details>

<details>
<summary>Hint 2: Testing Both Success and Errors</summary>

```go
tests := []struct {
    name    string
    input   string
    want    int
    wantErr bool
    errMsg  string
}{
    {"valid", "123", 123, false, ""},
    {"invalid", "abc", 0, true, "invalid syntax"},
}
```
</details>

---

## What This Teaches

- Error testing patterns
- Checking error messages
- Custom error types
- errors.Is() and errors.As()
- Comprehensive error coverage

---

**Next Exercise:** `06_benchmarks` - Performance testing with testing.B
