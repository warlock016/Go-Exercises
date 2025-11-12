# Exercise 02: Type Conversions

**Concept:** Converting between different types safely
**Difficulty:** Easy
**Estimated Time:** 10-15 minutes

## Learning Goals
- Convert between numeric types (int, float64)
- Convert numbers to strings and vice versa
- Understand when conversions are necessary
- Handle conversion errors

## Problem

Implement functions that convert between different Go types.

## Tasks

### 1. IntToFloat
Convert an integer to a float64

### 2. FloatToInt
Convert a float64 to an integer (truncating decimal part)

### 3. StringToInt
Convert a string to an integer, returning an error if conversion fails

### 4. IntToString
Convert an integer to its string representation

## Function Signatures

```go
func IntToFloat(n int) float64
func FloatToInt(f float64) int
func StringToInt(s string) (int, error)
func IntToString(n int) string
```

## Examples

```go
IntToFloat(42)           // → 42.0
FloatToInt(3.14)         // → 3
FloatToInt(9.99)         // → 9
StringToInt("123")       // → 123, nil
StringToInt("abc")       // → 0, error
IntToString(42)          // → "42"
```

## Instructions

1. Open `convert.go`
2. Implement all functions
3. Run `go test -v`
4. Create `EXPLANATION.md`

## Hints

### Numeric Conversions
```go
// Int to float
i := 42
f := float64(i)

// Float to int (truncates)
f := 3.14
i := int(f)  // i = 3
```

### String Conversions
```go
import "strconv"

// String to int
num, err := strconv.Atoi("123")
if err != nil {
    // Handle error
}

// Int to string
str := strconv.Itoa(123)
```

## What This Teaches

- Explicit type conversions in Go
- Go doesn't allow implicit conversions
- Using the `strconv` package
- Error handling in conversions
- Integer truncation behavior
