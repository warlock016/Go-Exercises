# Exercise 07: Person Struct

**Concept:** Structs, methods, receivers
**Difficulty:** Medium
**Estimated Time:** 7-10 minutes

## Problem

Create a `Person` struct and implement a method that returns a greeting.

## Requirements

1. Define a `Person` struct with fields: `Name` (string) and `Age` (int)
2. Implement a method `Greet()` that returns a string like: "Hello, my name is Alice and I am 25 years old."

## Function/Method Signatures

```go
type Person struct {
    // Define fields here
}

func (p Person) Greet() string
```

## Example

```go
p := Person{Name: "Alice", Age: 25}
p.Greet() // Returns: "Hello, my name is Alice and I am 25 years old."
```

## Instructions

1. Open `person.go`
2. Define the `Person` struct
3. Implement the `Greet()` method
4. Run `go test -v` to check your solution

## Hints

- Struct fields should be exported (start with uppercase)
- Use `fmt.Sprintf()` to format the string
- The method receiver is `(p Person)`

## What This Tests

- Defining structs
- Exported vs unexported fields
- Methods with receivers
- String formatting
