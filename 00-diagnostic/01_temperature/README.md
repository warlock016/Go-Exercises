# Exercise 01: Temperature Converter

**Concept:** Variables, types, functions, basic arithmetic
**Difficulty:** Easy
**Estimated Time:** 3-5 minutes

## Problem

Write a function that converts temperature from Celsius to Fahrenheit.

The formula is: `F = C * 9/5 + 32`

## Function Signature

```go
func CelsiusToFahrenheit(celsius float64) float64
```

## Examples

- `CelsiusToFahrenheit(0)` → `32.0`
- `CelsiusToFahrenheit(100)` → `212.0`
- `CelsiusToFahrenheit(-40)` → `-40.0`
- `CelsiusToFahrenheit(37)` → `98.6`

## Instructions

1. Open `temperature.go`
2. Implement the `CelsiusToFahrenheit` function
3. Run `go test -v` to check your solution

## Hints

- Remember to use `float64` for decimal numbers
- The formula needs parentheses in the right places
- You can test with `go test -v` in this directory

## What This Tests

- Basic function syntax
- Working with float64 type
- Simple arithmetic operations
- Return statements
