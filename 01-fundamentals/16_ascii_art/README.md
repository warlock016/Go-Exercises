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

// Row Count conjecture: (2x - 1) -> 2 - 1 == 0 (Approved)
// Max Space Count conjecture: x - diff == (2x - 1) - x == x - 1 (Approved)
// Max Symbol Count conjecture: 2x - 1 (Approved)

DrawDiamond(1) 
// 1 row // dif 0 // 2 - 1 == 1
// max space // 1 - 1 == 0
// max symbol // 1 - 1 == 0
// Returns:
// *      1 symbol, 0 spaces

DrawDiamond(2) 
// 3 rows // diff 1 // 4 - 1 == 3
// max space // 2 - 1 == 1
// max symbol // 4 - 1 == 3

// Returns:
//   *      1 symbol, 1 space
//  ***     3 symbols, 0 space
//   *      1 symbol, 1 space

DrawDiamond(3) 
// 5 rows // diff 2 // 6  - 1 == 5
// max space // 3 - 1 == 2
// max symbol // 6 - 1 

// Returns:
//   *      1 symbol, 2 space
//  ***     3 symbol, 1 space
// *****    5 symbol, 0 space
//  ***     3 symbol, 1 space
//   *      1 symbol, 2 space

DrawDiamond(4) 
// 7 rows // diff 3 // 8 - 1 == 7
// max space // 4 - 1 == 3
// max symbol // 8 - 1 == 7

// Returns:
//    *     1 symbol, 3 space
//   ***    3 symbol, 2 space
//  *****   5 symbol, 1 space
// *******  7 symbol, 0 space
//  *****   5 symbol, 1 space
//   ***    3 symbol, 2 space
//    *     1 symbol, 3 space

DrawDiamond(5) 
// 9 rows // diff 4 // 10 - 1 == 9

// Returns:
//     *        1 symbol, 4 space
//    ***       3 symbol, 3 space
//   *****      5 symbol, 2 space
//  *******     7 symbol, 1 space
// *********    9 symbol, 0 space
//  *******     7 symbol, 1 space
//   *****      5 symbol, 2 space
//    ***       3 symbol, 3 space
//     *        1 symbol, 4 space

DrawDiamond(6) 
// 11 rows // diff 5 // 12 - 1 == 11

// Returns:
//      *       1  symbol, 5 spaces
//     ***      3  symbol, 4 spaces
//    *****     5  symbol, 3 spaces
//   *******    7  symbol, 2 spaces
//  *********   9  symbol, 1 spaces
// ***********  11 symbol, 0 spaces
//  *********   9  symbol, 1 spaces
//   *******    7  symbol, 2 spaces
//    *****     5  symbol, 3 spaces
//     ***      3  symbol, 4 spaces
//      *       1  symbol, 5 spaces

DrawDiamond(7) 
// 13 rows // diff 6 // 14 - 1 == 13

// Returns:
//0        *
//1       ***
//2      *****
//3     *******
//4    *********
//5   ***********
//6  *************
//7   ***********
//8    *********
//9     *******
//10     *****
//11      ***
//12       *

DrawDiamond(8) 
// 15 rows // diff 7 // 16 - 1 == 15
// max space // 8 - 1 == 7
// max symbol // 16 - 1 == 15

// Returns:
//        *
//       ***
//      *****
//     *******
//    *********
//   ***********
//  *************
// ***************
//  *************
//   ***********
//    *********
//     *******
//      *****
//       ***
//        *
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
