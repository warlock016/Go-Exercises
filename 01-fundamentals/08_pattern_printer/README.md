# Pattern Printer

## Learning Goal

Master nested loops by generating ASCII art patterns with precise string formatting.

## Problem Description

Pattern printing is a classic programming exercise that develops understanding of nested loops, string concatenation, and spatial reasoning. You'll create functions that generate various ASCII art patterns using asterisks and numbers.

Each function returns a complete string with newlines, making it easy to test and print.

## Function Signatures

```go
func PrintSquare(n int) string
func PrintTriangle(n int) string
func PrintPyramid(n int) string
func PrintNumberSquare(n int) string
```

## Examples

### Example 1: Square Pattern
```go
PrintSquare(3)
// Returns:
// "***\n***\n***\n"
// Displays as:
// ***
// ***
// ***
```

### Example 2: Right Triangle
```go
PrintTriangle(4)
// Returns:
// "*\n**\n***\n****\n"
// Displays as:
// *
// **
// ***
// ****
```

### Example 3: Centered Pyramid
```go
PrintPyramid(3)
// Returns:
// "  *\n ***\n*****\n"
// Displays as:
//   *
//  ***
// *****
```

### Example 4: Number Square
```go
PrintNumberSquare(3)
// Returns:
// "111\n222\n333\n"
// Displays as:
// 111
// 222
// 333
```

### Example 5: Edge Cases
```go
PrintSquare(1)    // "*\n"
PrintTriangle(1)  // "*\n"
PrintPyramid(1)   // "*\n"
PrintSquare(0)    // ""
```

## Instructions

1. Implement `PrintSquare`
2. Implement `PrintTriangle`
3. Implement `PrintPyramid`
4. Implement `PrintNumberSquare`
5. Run tests with `go test -v`

## Think About

1. Why is `strings.Builder` more efficient than using `+=` for string concatenation?
2. How would you modify `PrintPyramid` to create an inverted pyramid?
3. What happens if you swap the order of the nested loops?
4. How would you create a diamond pattern (pyramid + inverted pyramid)?

## What This Teaches

- **Nested Loop Control**: Understanding how outer and inner loops interact
- **String Building**: Efficient string concatenation with strings.Builder
- **Mathematical Patterns**: Translating visual patterns into loop logic
- **Spatial Reasoning**: Calculating spaces and characters for alignment
- **Edge Case Handling**: Dealing with zero or negative inputs
- **Test-Driven Output**: Returning exact strings that can be tested programmatically
