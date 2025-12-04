# Testing Guide

A comprehensive reference for Go testing patterns and practices.

---

## Running Tests

### Basic Commands

```bash
# Run all tests in current directory
go test

# Run with verbose output
go test -v

# Run specific test by name
go test -run TestFunctionName

# Run tests matching pattern
go test -run "TestUser.*"

# Run tests in all subdirectories
go test ./...

# Run with coverage
go test -cover

# Generate coverage profile
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Useful Flags

| Flag | Purpose |
|------|---------|
| `-v` | Verbose output |
| `-run <regex>` | Run matching tests |
| `-count N` | Run tests N times |
| `-race` | Enable race detector |
| `-timeout 30s` | Set timeout |
| `-short` | Skip long tests |
| `-parallel N` | Max parallel tests |

---

## Test File Basics

### File Naming

```
mycode.go       → mycode_test.go
handler.go      → handler_test.go
```

### Test Function Signature

```go
import "testing"

func TestFunctionName(t *testing.T) {
    // Test code here
}

// Test names should describe what's being tested
func TestUserCreation(t *testing.T) {}
func TestUserValidation_EmptyEmail(t *testing.T) {}
func TestParseConfig_InvalidJSON(t *testing.T) {}
```

---

## testing.T Methods

### Reporting Failures

```go
func TestExample(t *testing.T) {
    // Log message and continue
    t.Error("something went wrong")

    // Log formatted message and continue
    t.Errorf("got %d, want %d", got, want)

    // Log message and stop test immediately
    t.Fatal("cannot continue")

    // Log formatted message and stop
    t.Fatalf("setup failed: %v", err)
}
```

### Logging

```go
func TestExample(t *testing.T) {
    t.Log("informational message")
    t.Logf("value is %d", value)
}
```

### Skipping Tests

```go
func TestLongRunning(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping in short mode")
    }
    // Long test...
}
```

---

## Table-Driven Tests

### Basic Pattern

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -1, -2, -3},
        {"mixed signs", -1, 5, 4},
        {"zeros", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Add(%d, %d) = %d, want %d",
                    tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

### With Error Cases

```go
func TestDivide(t *testing.T) {
    tests := []struct {
        name    string
        a, b    int
        want    int
        wantErr bool
    }{
        {"normal division", 10, 2, 5, false},
        {"integer division", 7, 2, 3, false},
        {"divide by zero", 1, 0, 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Divide(tt.a, tt.b)

            if (err != nil) != tt.wantErr {
                t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
                return
            }

            if got != tt.want {
                t.Errorf("Divide() = %d, want %d", got, tt.want)
            }
        })
    }
}
```

---

## Subtests

### Running Subtests

```go
func TestMath(t *testing.T) {
    t.Run("Addition", func(t *testing.T) {
        if Add(2, 3) != 5 {
            t.Error("2 + 3 should equal 5")
        }
    })

    t.Run("Subtraction", func(t *testing.T) {
        if Subtract(5, 3) != 2 {
            t.Error("5 - 3 should equal 2")
        }
    })
}

// Run specific subtest:
// go test -run "TestMath/Addition"
```

### Parallel Subtests

```go
func TestParallel(t *testing.T) {
    tests := []struct {
        name  string
        input string
    }{
        {"test1", "a"},
        {"test2", "b"},
        {"test3", "c"},
    }

    for _, tt := range tests {
        tt := tt  // Capture range variable
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()  // Run in parallel
            // Test code...
        })
    }
}
```

---

## Test Helpers

### Using t.Helper()

```go
func assertEqual(t *testing.T, got, want int) {
    t.Helper()  // Marks this as a helper function
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}

func TestWithHelper(t *testing.T) {
    assertEqual(t, Add(2, 3), 5)  // Error points here, not inside helper
}
```

### Setup and Teardown

```go
func setupTest(t *testing.T) func() {
    t.Log("Setting up test")
    // Setup code...

    return func() {
        t.Log("Tearing down test")
        // Cleanup code...
    }
}

func TestWithSetup(t *testing.T) {
    teardown := setupTest(t)
    defer teardown()

    // Test code...
}
```

---

## Testing Errors

### Check Error Exists

```go
func TestErrorReturned(t *testing.T) {
    _, err := ParseInt("not a number")
    if err == nil {
        t.Error("expected error for invalid input")
    }
}
```

### Check Specific Error

```go
func TestSpecificError(t *testing.T) {
    _, err := FindUser(-1)

    if !errors.Is(err, ErrInvalidID) {
        t.Errorf("got %v, want ErrInvalidID", err)
    }
}
```

### Check Error Type

```go
func TestErrorType(t *testing.T) {
    _, err := Validate(data)

    var validErr *ValidationError
    if !errors.As(err, &validErr) {
        t.Fatal("expected ValidationError")
    }

    if validErr.Field != "email" {
        t.Errorf("wrong field: got %s, want email", validErr.Field)
    }
}
```

---

## Benchmarks

### Basic Benchmark

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1, 2)
    }
}

// Run: go test -bench=.
// Output: BenchmarkAdd-8    1000000000    0.3 ns/op
```

### With Memory Allocation

```go
func BenchmarkConcat(b *testing.B) {
    b.ReportAllocs()  // Report memory allocations

    for i := 0; i < b.N; i++ {
        Concat("hello", "world")
    }
}

// Output: BenchmarkConcat-8    10000000    150 ns/op    32 B/op    1 allocs/op
```

### Benchmark with Setup

```go
func BenchmarkSort(b *testing.B) {
    data := generateData(10000)

    b.ResetTimer()  // Don't count setup time

    for i := 0; i < b.N; i++ {
        sorted := make([]int, len(data))
        copy(sorted, data)
        sort.Ints(sorted)
    }
}
```

---

## Test Fixtures

### testdata Directory

```
mypackage/
├── mycode.go
├── mycode_test.go
└── testdata/
    ├── valid_input.json
    ├── invalid_input.json
    └── expected_output.json
```

### Loading Test Data

```go
func TestParseConfig(t *testing.T) {
    // testdata is ignored by go build
    data, err := os.ReadFile("testdata/valid_config.json")
    if err != nil {
        t.Fatalf("failed to read test data: %v", err)
    }

    config, err := ParseConfig(data)
    if err != nil {
        t.Fatalf("ParseConfig failed: %v", err)
    }

    // Assertions...
}
```

### Golden Files

```go
var update = flag.Bool("update", false, "update golden files")

func TestOutput(t *testing.T) {
    got := Generate()

    golden := filepath.Join("testdata", t.Name()+".golden")

    if *update {
        os.WriteFile(golden, got, 0644)
    }

    want, _ := os.ReadFile(golden)

    if !bytes.Equal(got, want) {
        t.Errorf("output mismatch")
    }
}

// Update golden files: go test -update
```

---

## HTTP Testing

### Testing Handlers with httptest

```go
import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandler(t *testing.T) {
    // Create request
    req := httptest.NewRequest("GET", "/users/123", nil)

    // Create response recorder
    w := httptest.NewRecorder()

    // Call handler
    UserHandler(w, req)

    // Check response
    resp := w.Result()

    if resp.StatusCode != http.StatusOK {
        t.Errorf("got status %d, want %d", resp.StatusCode, http.StatusOK)
    }

    body, _ := io.ReadAll(resp.Body)
    // Check body...
}
```

### Testing with POST Body

```go
func TestCreateUser(t *testing.T) {
    body := strings.NewReader(`{"name": "Alice", "email": "alice@example.com"}`)
    req := httptest.NewRequest("POST", "/users", body)
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    CreateUserHandler(w, req)

    if w.Code != http.StatusCreated {
        t.Errorf("got %d, want %d", w.Code, http.StatusCreated)
    }
}
```

### Mocking HTTP Server

```go
func TestAPIClient(t *testing.T) {
    // Create mock server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"id": 1, "name": "Test"}`))
    }))
    defer server.Close()

    // Use mock server URL
    client := NewClient(server.URL)
    result, err := client.GetUser(1)

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    // Check result...
}
```

---

## Build Tags for Test Types

### Integration Tests

```go
//go:build integration

package mypackage

func TestDatabaseConnection(t *testing.T) {
    // This only runs with: go test -tags=integration
}
```

### Running Tagged Tests

```bash
# Run unit tests only (default)
go test ./...

# Run integration tests
go test -tags=integration ./...

# Run all tests
go test -tags="integration,e2e" ./...
```

---

## Test Coverage

### Generate Coverage Report

```bash
# Show coverage percentage
go test -cover

# Generate coverage profile
go test -coverprofile=coverage.out

# View in browser
go tool cover -html=coverage.out

# View in terminal
go tool cover -func=coverage.out
```

### Coverage Modes

```bash
# Count: how many times each statement runs
go test -covermode=count -coverprofile=coverage.out

# Atomic: like count but safe for concurrent tests
go test -covermode=atomic -coverprofile=coverage.out
```

---

## Best Practices

### Do

1. **Use table-driven tests** - They're clear and easy to extend
2. **Test behavior, not implementation** - Focus on inputs and outputs
3. **Use meaningful test names** - `TestUser_InvalidEmail_ReturnsError`
4. **Keep tests independent** - No shared state between tests
5. **Test edge cases** - Empty, nil, boundary values
6. **Use t.Helper()** - For cleaner error messages

### Don't

1. **Don't test private functions** - Test through public API
2. **Don't over-mock** - Keep tests close to real behavior
3. **Don't ignore flaky tests** - Fix them or remove them
4. **Don't write tests after the fact** - TDD is more effective
5. **Don't aim for 100% coverage** - Diminishing returns

---

## Resources

- [Go Testing Package](https://pkg.go.dev/testing)
- [Go Blog - Table Driven Tests](https://go.dev/blog/subtests)
- [httptest Package](https://pkg.go.dev/net/http/httptest)
- [Go by Example - Testing](https://gobyexample.com/testing)
