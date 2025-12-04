# Exercise 11: Test Coverage

**Learning Goal:** Use go test -cover to measure coverage, generate coverage profiles, and identify untested code paths.

---

## Problem Description

Test coverage shows which code is executed by tests. Go provides built-in coverage tools to help you identify untested code paths and improve test quality.

---

## Functions to Test

```go
func Classify(n int) string  // Has multiple branches
func ProcessData(data []int) (int, error)  // Has error paths
```

---

## Coverage Commands

```bash
# Basic coverage
go test -cover

# Coverage with profile
go test -coverprofile=coverage.out

# View coverage in browser
go tool cover -html=coverage.out

# Coverage by function
go tool cover -func=coverage.out
```

---

## Your Task

1. Write tests for the functions
2. Run `go test -cover` - observe initial coverage
3. Generate coverage profile: `go test -coverprofile=coverage.out`
4. View in browser: `go tool cover -html=coverage.out`
5. Add tests to cover untested paths
6. Achieve >90% coverage

---

**Next Exercise:** `12_integration_vs_unit` - Unit vs integration testing strategies
