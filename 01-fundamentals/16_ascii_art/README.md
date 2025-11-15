# Exercise 16: ASCII Art

## Learning Goal
Master complex string formatting by generating ASCII art patterns using loops, string building, and precise character placement.

## Problem Description
Create functions that generate ASCII art patterns. This exercise combines string manipulation, mathematical patterns, and character-based drawing to create visual output in the terminal.

## Function Signatures
```go
// DrawBox creates a rectangular box using the specified character
// Width and height must be at least 2
func DrawBox(width, height int, char rune) string

// DrawDiamond creates a diamond pattern with n rows for the top half
// n must be at least 1
func DrawDiamond(n int) string

// DrawChessboard creates an n×n chessboard with alternating filled/empty blocks
// n must be at least 1
func DrawChessboard(n int) string
```

## Examples

### DrawBox
```go
DrawBox(5, 3, '*')
// Returns:
// *****
// *   *
// *****

DrawBox(7, 4, '#')
// Returns:
// #######
// #     #
// #     #
// #######

DrawBox(3, 3, '+')
// Returns:
// +++
// + +
// +++
```

### DrawDiamond
```go
DrawDiamond(3)
// Returns:
//   *
//  ***
// *****
//  ***
//   *

DrawDiamond(1)
// Returns:
// *

DrawDiamond(4)
// Returns:
//    *
//   ***
//  *****
// *******
//  *****
//   ***
//    *
```

### DrawChessboard
```go
DrawChessboard(4)
// Returns:
// █░█░
// ░█░█
// █░█░
// ░█░█

DrawChessboard(3)
// Returns:
// █░█
// ░█░
// █░█

DrawChessboard(1)
// Returns:
// █
```

## Instructions

1. Implement `DrawBox`
2. Implement `DrawDiamond`
3. Implement `DrawChessboard`
4. Run tests with `go test -v`

## Think About

1. Why is `strings.Builder` more efficient than using `+=` for string concatenation in loops?
2. What mathematical relationship determines the number of stars in each row of the diamond?
3. How does the modulo operator help create the alternating chessboard pattern?
4. What would happen if you forgot the newline character at the end of each row?

## What This Teaches

- **String Building:** Using `strings.Builder` for efficient string construction in loops
- **Pattern Recognition:** Identifying mathematical patterns in visual output
- **Loop Design:** Nested loops for 2D patterns, single loops with calculations
- **Edge Case Handling:** Validating inputs and handling boundary conditions
- **Character Manipulation:** Working with runes and Unicode characters (█, ░)
- **Spatial Reasoning:** Translating visual patterns into code logic

## Success Criteria

- All three functions implemented correctly
- All tests pass (basic cases, edge cases, pattern validation)
- Efficient use of `strings.Builder` instead of string concatenation
- Exact character-for-character output match (including newlines)
- Edge cases handled (width/height < 2 for box, n < 1 for diamond/chessboard)
