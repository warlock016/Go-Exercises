# Exercise 08: Divide Safely

**Concept:** Error handling, division, zero checks
**Difficulty:** Medium
**Estimated Time:** 5-8 minutes

## Problem

Write a function that divides two numbers safely, returning an error if dividing by zero.

## Function Signature

```go
func Divide(a, b float64) (float64, error)
```

## Examples

- `Divide(10, 2)` → `5.0, nil`
- `Divide(7, 2)` → `3.5, nil`
- `Divide(10, 0)` → `0, error` (error: "cannot divide by zero")
- `Divide(0, 5)` → `0.0, nil`

## Instructions

1. Open `divide.go`
2. Implement the `Divide` function
3. Run `go test -v` to check your solution

## Hints

- Check if `b == 0` before dividing
- Return `errors.New("cannot divide by zero")` for zero divisor
- Import "errors" package
- Return `a / b, nil` for valid divisions

## What This Tests

- Error handling pattern in Go
- Conditional logic
- Creating and returning errors
- Float64 operations
