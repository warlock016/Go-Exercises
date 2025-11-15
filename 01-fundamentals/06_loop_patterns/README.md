# Loop Patterns

## Learning Goal

Master Go's versatile `for` loop in all its forms: traditional three-component, while-style, infinite with break, and range-based iteration. Understand when to use each variant and how to build slices dynamically.

## Problem Description

Go has only one looping construct: the `for` loop. However, this single loop type can take multiple forms, making it incredibly versatile. In this exercise, you'll practice four essential loop patterns:

1. **Traditional for loop** - `for i := 0; i < n; i++` (initialization; condition; post)
2. **While-style loop** - `for condition` (just a condition, like while in other languages)
3. **Infinite loop with break** - `for { ... break }` (explicit exit control)
4. **Range-based loop** - `for i, v := range slice` (iterate over collections)

You'll build functions that construct slices using different loop patterns, helping you internalize when each form is most appropriate.

## Function Signatures

```go
// CountUp returns a slice containing [1, 2, 3, ..., n]
// Use traditional for loop: for i := start; i <= n; i++
func CountUp(n int) []int

// CountDown returns a slice containing [n, n-1, n-2, ..., 1]
// Use while-style for loop: for condition { ... }
func CountDown(n int) []int

// SumSlice returns the sum of all numbers in the slice
// Use range-based for loop: for _, v := range nums
func SumSlice(nums []int) int

// FirstNEvens returns the first n even positive numbers [2, 4, 6, ...]
// Use infinite for loop with break: for { ... if count == n { break } }
func FirstNEvens(n int) []int
```

## Examples

### CountUp
```go
CountUp(5)    // [1, 2, 3, 4, 5]
CountUp(1)    // [1]
CountUp(0)    // []
CountUp(-3)   // []
```

### CountDown
```go
CountDown(5)   // [5, 4, 3, 2, 1]
CountDown(1)   // [1]
CountDown(0)   // []
CountDown(-2)  // []
```

### SumSlice
```go
SumSlice([]int{1, 2, 3, 4, 5})    // 15
SumSlice([]int{10, 20, 30})       // 60
SumSlice([]int{-5, 5})            // 0
SumSlice([]int{})                 // 0
```

### FirstNEvens
```go
FirstNEvens(5)   // [2, 4, 6, 8, 10]
FirstNEvens(3)   // [2, 4, 6]
FirstNEvens(1)   // [2]
FirstNEvens(0)   // []
```

## Instructions

1. Implement `CountUp`
2. Implement `CountDown`
3. Implement `SumSlice`
4. Implement `FirstNEvens`
5. Run tests with `go test -v`

## Think About

1. **When would you choose traditional for over while-style?**
   - Traditional: When you know iteration count or need index/counter
   - While-style: When you loop until a condition changes (unknown iterations)

2. **Why does Go only have `for` instead of `while` and `for`?**
   - Simplicity: One construct does everything
   - The for loop is flexible enough to handle all looping scenarios
   - Reduces language complexity without reducing power

3. **When is an infinite loop + break pattern useful?**
   - When exit condition is in the middle of loop logic
   - When you need to process at least once before checking
   - When the loop exit logic is complex or has multiple conditions

4. **What's the advantage of range over traditional indexing?**
   - Cleaner syntax, less error-prone (no off-by-one errors)
   - Automatically handles slice length
   - Can ignore index or value with `_` when not needed

## What This Teaches

- **Loop versatility**: Go's single `for` construct handles all iteration patterns
- **Idiomatic Go**: Using range for collections is more Go-like than index-based iteration
- **Slice building**: Dynamic slice construction with `append`
- **Pattern recognition**: Knowing which loop form fits which problem
- **Edge case handling**: Empty results for invalid inputs (n <= 0)
- **Control flow**: Using break to exit loops explicitly
