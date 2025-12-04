# Exercise 01: First Test

**Learning Goal:** Learn the basics of Go testing - test functions, testing.T methods, and running tests.

---

## Problem Description

You have a simple calculator with basic arithmetic functions. Your task is to write comprehensive tests for these functions to understand how Go's testing framework works.

This exercise introduces:
- Test function naming convention: `func TestXxx(t *testing.T)`
- Error reporting methods: `t.Error()`, `t.Errorf()`, `t.Fatal()`, `t.Fatalf()`
- Running tests with `go test` and `go test -v`
- The difference between failing and fatal errors

---

## Functions to Test

The `calculator.go` file provides these functions:

```go
func Add(a, b int) int
func Subtract(a, b int) int
func Multiply(a, b int) int
func Divide(a, b int) (int, error)  // Returns error if b == 0
func IsEven(n int) bool
```

---

## Your Task

Write comprehensive tests in `calculator_test.go` for all functions. For each function:

1. Test normal cases (typical inputs)
2. Test edge cases (zero, negative numbers)
3. Test error conditions (for Divide)
4. Use descriptive test names
5. Use appropriate error reporting methods

---

## Examples

### Testing Add Function
```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d, want 5", result)
    }
}
```

### Testing Divide with Error
```go
func TestDivideByZero(t *testing.T) {
    _, err := Divide(10, 0)
    if err == nil {
        t.Error("Divide(10, 0) expected error, got nil")
    }
}
```

---

## Test Methods Reference

### t.Error() and t.Errorf()
- Reports test failure but **continues** running the test
- Use for multiple independent checks in one test
- `t.Errorf()` supports formatting like `fmt.Printf()`

```go
// Example: Multiple checks in one test
func TestMultiple(t *testing.T) {
    if Add(1, 1) != 2 {
        t.Error("1+1 failed")
    }
    if Add(2, 2) != 4 {
        t.Error("2+2 failed")  // This still runs even if first check failed
    }
}
```

### t.Fatal() and t.Fatalf()
- Reports test failure and **stops** the test immediately
- Use when later checks depend on earlier ones
- Use when continuing after failure would cause panic

```go
// Example: Stop on critical failure
func TestWithFatal(t *testing.T) {
    result, err := Divide(10, 2)
    if err != nil {
        t.Fatalf("Divide(10, 2) unexpected error: %v", err)
        // Test stops here if error occurs
    }
    // This only runs if no error above
    if result != 5 {
        t.Errorf("Divide(10, 2) = %d, want 5", result)
    }
}
```

**When to use which?**
- Use `t.Error()` for independent checks
- Use `t.Fatal()` when test can't continue meaningfully

---

## Running Tests

```bash
# Run all tests
go test

# Run with verbose output (shows all test names)
go test -v

# Run specific test
go test -v -run TestAdd

# Run tests matching pattern
go test -v -run TestDivide
```

---

## Instructions

1. Open `calculator_test.go`
2. Follow the TODO(human) markers to implement tests
3. Write at least 2-3 test cases per function
4. Run `go test -v` frequently to see your progress
5. Make sure all tests pass

**Test Checklist:**
- [ ] TestAdd - test positive numbers, negative numbers, zero
- [ ] TestSubtract - test various cases including negative results
- [ ] TestMultiply - test positive, negative, zero
- [ ] TestDivide - test normal division
- [ ] TestDivideByZero - test error condition
- [ ] TestIsEven - test even numbers, odd numbers, zero, negative

---

## Hints

<details>
<summary>Hint 1: Basic Test Structure</summary>

Every test function:
1. Calls the function being tested
2. Compares result to expected value
3. Reports failure if they don't match

```go
func TestSomething(t *testing.T) {
    got := FunctionToTest(input)
    want := expectedValue
    if got != want {
        t.Errorf("FunctionToTest(%v) = %v, want %v", input, got, want)
    }
}
```
</details>

<details>
<summary>Hint 2: Testing Errors</summary>

When testing functions that return errors:

```go
// Test error is returned
result, err := FunctionThatMightError(badInput)
if err == nil {
    t.Error("expected error, got nil")
}

// Test no error is returned
result, err := FunctionThatMightError(goodInput)
if err != nil {
    t.Fatalf("unexpected error: %v", err)
}
```
</details>

<details>
<summary>Hint 3: Test Names</summary>

Good test names describe what they test:
- ✅ `TestAddPositiveNumbers`
- ✅ `TestDivideByZero`
- ✅ `TestIsEvenWithNegative`
- ❌ `TestAdd1`
- ❌ `TestCase2`
</details>

<details>
<summary>Full Solution</summary>

```go
package calculator

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d, want 5", result)
    }
}

func TestAddNegative(t *testing.T) {
    result := Add(-2, -3)
    if result != -5 {
        t.Errorf("Add(-2, -3) = %d, want -5", result)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(5, 3)
    if result != 2 {
        t.Errorf("Subtract(5, 3) = %d, want 2", result)
    }
}

func TestMultiply(t *testing.T) {
    result := Multiply(4, 5)
    if result != 20 {
        t.Errorf("Multiply(4, 5) = %d, want 20", result)
    }
}

func TestMultiplyByZero(t *testing.T) {
    result := Multiply(5, 0)
    if result != 0 {
        t.Errorf("Multiply(5, 0) = %d, want 0", result)
    }
}

func TestDivide(t *testing.T) {
    result, err := Divide(10, 2)
    if err != nil {
        t.Fatalf("Divide(10, 2) unexpected error: %v", err)
    }
    if result != 5 {
        t.Errorf("Divide(10, 2) = %d, want 5", result)
    }
}

func TestDivideByZero(t *testing.T) {
    _, err := Divide(10, 0)
    if err == nil {
        t.Error("Divide(10, 0) expected error, got nil")
    }
}

func TestIsEven(t *testing.T) {
    if !IsEven(4) {
        t.Error("IsEven(4) = false, want true")
    }
    if IsEven(5) {
        t.Error("IsEven(5) = true, want false")
    }
    if !IsEven(0) {
        t.Error("IsEven(0) = false, want true")
    }
    if IsEven(-3) {
        t.Error("IsEven(-3) = true, want false")
    }
}
```
</details>

---

## What This Teaches

- **Test function signature** - Must be `func TestXxx(t *testing.T)`
- **Error reporting** - Difference between `t.Error()` and `t.Fatal()`
- **Test organization** - One test function per scenario or group of related checks
- **Running tests** - How to use `go test` command
- **Test-driven thinking** - How to think about what to test

---

## Think About

1. What happens if you name a test function `test_add()` (lowercase)? Try it!
2. Why might you use `t.Fatal()` instead of `t.Error()`?
3. How would you test a function that takes many parameters?
4. What makes a good error message in `t.Errorf()`?

---

**Next Exercise:** `02_table_driven_basics` - Learn the idiomatic Go testing pattern
