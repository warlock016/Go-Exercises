# Exercise 03: Simple Calculator

**Concept:** Arithmetic operations, functions with multiple parameters
**Difficulty:** Easy
**Estimated Time:** 10-15 minutes

## Learning Goals
- Use arithmetic operators (+, -, *, /, %)
- Write functions with multiple parameters
- Understand integer vs float division
- Handle division by zero

## Problem

Implement basic calculator functions for arithmetic operations.

## Tasks

Implement the following calculator functions:

### 1. Add
Add two integers and return the result

### 2. Subtract
Subtract b from a and return the result

### 3. Multiply
Multiply two integers and return the result

### 4. Divide
Divide a by b, returning the quotient and remainder
Return an error if b is zero

### 5. Modulo
Return the remainder of a divided by b
Return an error if b is zero

## Function Signatures

```go
func Add(a, b int) int
func Subtract(a, b int) int
func Multiply(a, b int) int
func Divide(a, b int) (quotient, remainder int, err error)
func Modulo(a, b int) (int, error)
```

## Examples

```go
Add(5, 3)           // → 8
Subtract(10, 4)     // → 6
Multiply(6, 7)      // → 42
Divide(17, 5)       // → 3, 2, nil (17 = 5*3 + 2)
Divide(10, 0)       // → 0, 0, error
Modulo(17, 5)       // → 2, nil
Modulo(10, 0)       // → 0, error
```

## Instructions

1. Open `calc.go`
2. Implement all functions
3. Run `go test -v`
4. Create `EXPLANATION.md`

## Hints

### Basic Operators
```go
sum := a + b
diff := a - b
product := a * b
quotient := a / b       // Integer division (truncates)
remainder := a % b      // Modulo operator
```

### Division by Zero
```go
import "errors"

if b == 0 {
    return 0, errors.New("division by zero")
}
```

### Multiple Return Values
```go
func Divide(a, b int) (int, int, error) {
    // quotient, remainder, error
    return a / b, a % b, nil
}
```

## What This Teaches

- Arithmetic operators in Go
- Functions with multiple parameters
- Multiple return values (including named returns)
- Error handling for invalid inputs
- Integer division behavior
