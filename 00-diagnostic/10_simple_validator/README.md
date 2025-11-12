# Exercise 10: Simple Validator

**Concept:** Interfaces (basic understanding)
**Difficulty:** Medium-Hard
**Estimated Time:** 10-15 minutes

## Problem

Create a `Validator` interface and implement it for a `EmailValidator` type.

## Requirements

1. Define a `Validator` interface with a method: `Validate(s string) bool`
2. Create an `EmailValidator` struct (can be empty)
3. Implement `Validate` method that checks if a string contains "@"

## Interface Signature

```go
type Validator interface {
    Validate(s string) bool
}

type EmailValidator struct {}

func (ev EmailValidator) Validate(s string) bool
```

## Example

```go
var v Validator = EmailValidator{}
v.Validate("user@example.com") // true
v.Validate("notanemail")       // false
```

## Instructions

1. Open `validator.go`
2. Define the `Validator` interface
3. Define the `EmailValidator` struct
4. Implement the `Validate` method
5. Run `go test -v` to check your solution

## Hints

- Interfaces define behavior (methods)
- Any type that implements all methods satisfies the interface
- Use `strings.Contains(s, "@")` to check for @ symbol
- Import "strings" package

## What This Tests

- Understanding of interfaces
- Interface implementation
- Method implementation
- Basic interface usage
