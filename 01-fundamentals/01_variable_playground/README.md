# Exercise 01: Variable Playground

**Concept:** Variable declarations, basic types, initialization
**Difficulty:** Easy
**Estimated Time:** 10-15 minutes

## Learning Goals
- Understand different variable declaration styles in Go
- Work with basic types: int, float64, string, bool
- Practice type inference with `:=`
- Learn the difference between `var` and `:=`

## Problem

Implement functions that demonstrate different ways to declare and initialize variables in Go.

## Tasks

Implement the following functions:

### 1. DeclareInteger
Declare and return an integer variable with value 42

### 2. DeclareFloat
Declare and return a float64 variable with value 3.14

### 3. DeclareString
Declare and return a string variable with value "Hello, Go!"

### 4. DeclareBoolean
Declare and return a boolean variable with value true

### 5. DeclareMultiple
Declare and return three variables: name (string), age (int), height (float64)
Values: "Alice", 25, 5.6

## Instructions

1. Open `variables.go`
2. Implement all five functions
3. Run `go test -v` to verify your solutions
4. Create `EXPLANATION.md` explaining your implementations

## Hints

There are multiple ways to declare variables in Go:
```go
// Method 1: var with type
var x int = 42

// Method 2: var with type inference
var y = 42

// Method 3: short declaration (most common)
z := 42

// Method 4: declaration without initialization
var w int  // defaults to 0
w = 42
```

For multiple variables:
```go
// Method 1: separate declarations
name := "Alice"
age := 25

// Method 2: multiple declaration
var name, age = "Alice", 25

// Method 3: short form
name, age := "Alice", 25
```

## What This Teaches

- Variable declaration syntax
- Type inference with `:=`
- When to use `var` vs `:=`
- Go's basic types
- Multiple return values

## Example Output

```go
i := DeclareInteger()      // i = 42
f := DeclareFloat()        // f = 3.14
s := DeclareString()       // s = "Hello, Go!"
b := DeclareBoolean()      // b = true
name, age, height := DeclareMultiple()  // "Alice", 25, 5.6
```
