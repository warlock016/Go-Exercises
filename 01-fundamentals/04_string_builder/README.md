# Exercise 04: String Builder

## Learning Goal

Master the practical differences between string concatenation techniques (`+`, `fmt.Sprintf`, `strings.Builder`) and learn when to use each based on performance and readability requirements.

## Problem Description

You've learned about strings, runes, and `strings.Builder` in the String Mastery module. Now it's time to apply that knowledge in real-world scenarios. Different string concatenation techniques have different performance characteristics and readability trade-offs:

- **`+` operator**: Simple and readable for 2-3 strings, but creates new allocations each time
- **`fmt.Sprintf`**: Best for formatted output with type conversions, moderate performance
- **`strings.Builder`**: Most efficient for building strings in loops or with many parts

This exercise focuses on choosing the right tool for each situation.

## Function Signatures

```go
func Concat(parts ...string) string

func BuildGreeting(name, title string) string

func RepeatWithSeparator(s string, n int, sep string) string

func FormatTable(headers, values []string) string
```

## Examples

### Example 1: Concat
```go
Concat("Hello", " ", "World")
// Returns: "Hello World"

Concat("Go", "lang", "is", "awesome")
// Returns: "Golangisawesome"
```

### Example 2: BuildGreeting
```go
BuildGreeting("Smith", "Dr.")
// Returns: "Hello, Dr. Smith!"

BuildGreeting("Johnson", "")
// Returns: "Hello, Johnson!"

BuildGreeting("Alice", "Prof.")
// Returns: "Hello, Prof. Alice!"
```

### Example 3: RepeatWithSeparator
```go
RepeatWithSeparator("Go", 3, "-")
// Returns: "Go-Go-Go"

RepeatWithSeparator("*", 5, "")
// Returns: "*****"

RepeatWithSeparator("ha", 1, ",")
// Returns: "ha"
```

### Example 4: FormatTable
```go
headers := []string{"Name", "Age", "City"}
values := []string{"Alice", "30", "NYC"}
FormatTable(headers, values)
// Returns:
// "Name  | Age | City
//  Alice | 30  | NYC"

headers := []string{"ID", "Status"}
values := []string{"42", "Active"}
FormatTable(headers, values)
// Returns:
// "ID | Status
//  42 | Active"
```

### Example 5: Edge Cases
```go
Concat()
// Returns: ""

RepeatWithSeparator("x", 0, "-")
// Returns: ""

FormatTable([]string{}, []string{})
// Returns: ""
```

## Instructions

1. Implement `Concat` to join any number of strings
2. Implement `BuildGreeting` to create a greeting with optional title
3. Implement `RepeatWithSeparator` to repeat a string with a separator
4. Implement `FormatTable` to create a simple two-row table

## Think About

1. When would you choose `+` operator over `strings.Builder`?
2. Why is `strings.Builder` more efficient than repeated concatenation with `+`?
3. How does pre-allocating with `Grow()` improve performance?
4. In what scenarios is `fmt.Sprintf` the best choice despite being slower?

## What This Teaches

- **Performance awareness**: Understanding allocation costs and when they matter
- **Idiomatic Go**: Knowing when to use standard library helpers like `strings.Join`
- **Readability vs optimization**: Balancing code clarity with performance
- **Tool selection**: Matching the string-building technique to the use case
- **Practical string manipulation**: Real-world patterns like table formatting and repeated elements
