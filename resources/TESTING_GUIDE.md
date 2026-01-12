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

## Testing Concurrent Code

Concurrent code introduces unique testing challenges: race conditions, deadlocks, and non-deterministic behavior. This section covers essential techniques for testing goroutines and channels.

### Race Detector

The `-race` flag detects data races at runtime:

```bash
# Run tests with race detector
go test -race ./...

# Run specific test with race detection
go test -race -run TestWorkerPool

# Run program with race detection
go run -race main.go
```

**What it catches:**
- Concurrent read/write to shared variables without synchronization
- Multiple goroutines accessing shared data unsafely

**What it doesn't catch:**
- Deadlocks
- Logic errors
- Channel misuse (nil channels, multiple receivers racing)

**Example output:**
```
WARNING: DATA RACE
Write at 0x00c000014088 by goroutine 7:
  main.main.func1()
      /path/to/file.go:15 +0x38

Previous read at 0x00c000014088 by goroutine 6:
  main.main.func1()
      /path/to/file.go:15 +0x38
```

### Detecting Flaky Tests

Flaky tests pass sometimes and fail sometimes — often due to race conditions:

```bash
# Run test multiple times to expose flakiness
go test -race -run TestPingPong -count=10

# Run without race flag to compare
go test -run TestPingPong -count=10
```

**Key insight:** A test that passes without `-race` but fails with it is a major red flag. The race detector changes scheduler behavior, exposing hidden bugs.

### Timeout Wrappers for Deadlock Detection

Wrap blocking operations to detect hangs:

```go
func TestWithTimeout(t *testing.T) {
    done := make(chan int)

    go func() {
        done <- FunctionThatMightDeadlock()
    }()

    select {
    case result := <-done:
        // Test assertions on result
        if result != expected {
            t.Errorf("got %d, want %d", result, expected)
        }
    case <-time.After(2 * time.Second):
        t.Fatal("test timed out - likely deadlock")
    }
}
```

### Testing with synctest (Go 1.25+)

The `testing/synctest` package provides fake time for testing time-dependent code:

```go
import (
    "testing"
    "testing/synctest"
    "time"
)

func TestTimeout(t *testing.T) {
    t.Run("times out after duration", func(t *testing.T) {
        synctest.Test(t, func(t *testing.T) {
            ch := Timeout(5 * time.Second)

            // Fast-forward time (instant, no real waiting)
            time.Sleep(5 * time.Second)

            select {
            case <-ch:
                // Success - channel closed
            default:
                t.Error("channel should be closed")
            }
        })
    })
}
```

**Important:** `t.Run` must be **outside** the `synctest.Test` bubble, not inside:

```go
// ✅ CORRECT
t.Run("subtest", func(t *testing.T) {
    synctest.Test(t, func(t *testing.T) {
        // test code
    })
})

// ❌ WRONG - causes panic
synctest.Test(t, func(t *testing.T) {
    t.Run("subtest", func(t *testing.T) {  // panic!
        // test code
    })
})
```

### Common Concurrency Testing Patterns

#### Testing Channel Producers

```go
func TestGenerator(t *testing.T) {
    done := make(chan struct{})
    ch := Generator(done)

    // Collect some values
    var got []int
    for i := 0; i < 5; i++ {
        got = append(got, <-ch)
    }

    // Signal stop
    close(done)

    want := []int{0, 1, 2, 3, 4}
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}
```

#### Testing Worker Pools

```go
func TestWorkerPool(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}
    got := WorkerPool(input)
    want := 30  // sum of doubled values

    if got != want {
        t.Errorf("WorkerPool(%v) = %d, want %d", input, got, want)
    }
}

// Also run with race detector:
// go test -race -run TestWorkerPool
```

#### Testing for Goroutine Leaks

```go
func TestNoGoroutineLeak(t *testing.T) {
    before := runtime.NumGoroutine()

    // Run your concurrent code
    done := make(chan struct{})
    ch := Generator(done)
    <-ch  // receive one value
    close(done)

    // Give goroutines time to clean up
    time.Sleep(100 * time.Millisecond)

    after := runtime.NumGoroutine()
    if after > before {
        t.Errorf("goroutine leak: %d before, %d after", before, after)
    }
}
```

### Debugging Failed Concurrent Tests

When a concurrent test fails:

1. **Add timeout wrapper** — Confirms if it's a deadlock
2. **Run with `-race`** — Catches data races
3. **Run with `-count=10`** — Exposes flaky behavior
4. **Use Delve** — `dlv test -- -test.run TestName`
   - `goroutines` command shows what each goroutine is blocked on
5. **Add strategic logging** — Track channel operations

```go
// Debug logging for channels
func debugSend(ch chan<- int, v int, name string) {
    log.Printf("[%s] about to send %d", name, v)
    ch <- v
    log.Printf("[%s] sent %d", name, v)
}
```

### Common Concurrent Test Bugs

| Symptom | Likely Cause | Fix |
|---------|--------------|-----|
| Test hangs | Deadlock, nil channel | Add timeout wrapper, check channel initialization |
| Flaky pass/fail | Race condition | Run with `-race -count=10` |
| Wrong values | Multiple receivers racing | Use dedicated channels per receiver |
| Goroutine leak | Missing close, blocked send | Check all goroutines exit, add done channels |

### Best Practices for Concurrent Tests

1. **Always use `-race` in CI** — Catches races before production
2. **Run tests multiple times** — `go test -count=10` exposes flakiness
3. **Use timeouts** — Don't let tests hang forever
4. **Test edge cases** — Empty input, single element, cancellation
5. **Verify cleanup** — Check goroutine count before/after
6. **Isolate concurrent logic** — Makes testing easier

---

## Resources

- [Go Testing Package](https://pkg.go.dev/testing)
- [Go Blog - Table Driven Tests](https://go.dev/blog/subtests)
- [httptest Package](https://pkg.go.dev/net/http/httptest)
- [Go by Example - Testing](https://gobyexample.com/testing)
- [Go Race Detector](https://go.dev/doc/articles/race_detector)
- [testing/synctest (Go 1.25)](https://pkg.go.dev/testing/synctest)
