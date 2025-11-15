# Leap Year

## Learning Goal

Master complex boolean logic with multiple conditional rules by implementing leap year determination.

## Problem Description

Determine whether a given year is a leap year. Leap years have an extra day (February 29th) and follow specific rules based on divisibility.

## Leap Year Rules

A year is a leap year if:
1. It is divisible by 4
2. **EXCEPT** if it is divisible by 100, then it is NOT a leap year
3. **EXCEPT** if it is divisible by 400, then it IS a leap year

Think of it as three cascading rules where later rules override earlier ones.

## Function Signature

```go
func IsLeapYear(year int) bool
```

## Examples

```go
IsLeapYear(2000) // true  - divisible by 400
IsLeapYear(1900) // false - divisible by 100 but not 400
IsLeapYear(2004) // true  - divisible by 4 but not 100
IsLeapYear(2001) // false - not divisible by 4
IsLeapYear(1600) // true  - divisible by 400
IsLeapYear(1700) // false - divisible by 100 but not 400
IsLeapYear(2100) // false - divisible by 100 but not 400
```

## Instructions

1. Implement `IsLeapYear`
2. Run tests with `go test -v`

## Think About

1. Why do we need the exception for years divisible by 100?
2. What happens if you check conditions in a different order?
3. Can you express all three rules as a single boolean expression?
4. How would you explain this to someone who doesn't understand modulo arithmetic?

## What This Teaches

- Complex conditional logic with multiple rules
- Boolean operator precedence (`&&` vs `||`)
- The importance of condition ordering in if-else chains
- Translating real-world rules into code logic
- Compound boolean expressions vs explicit if-else statements
