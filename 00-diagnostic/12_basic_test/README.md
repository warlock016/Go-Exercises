# Exercise 12: Basic Testing

**Concept:** Writing tests, table-driven tests
**Difficulty:** Easy
**Estimated Time:** 5-8 minutes

## Problem

Write a test for the provided `Add` function using table-driven testing.

## Given Function

```go
func Add(a, b int) int {
    return a + b
}
```

## Your Task

Write a test function `TestAdd` that tests the `Add` function with multiple test cases using the table-driven test pattern.

## Test Cases to Include

- `Add(2, 3)` should return `5`
- `Add(0, 0)` should return `0`
- `Add(-1, 1)` should return `0`
- `Add(10, -5)` should return `5`

## Instructions

1. Open `math_test.go`
2. Implement the `TestAdd` function using table-driven tests
3. Run `go test -v` to check your solution

## Hints

- Use the pattern shown in other test files
- Create a slice of test structs with fields: name, a, b, want
- Loop through tests with `range`
- Use `t.Run()` for subtests
- Use `t.Errorf()` to report failures

## What This Tests

- Ability to write tests
- Understanding of table-driven test pattern
- Using testing.T methods
- Test organization
