# Number Pyramid

## Learning Goal

Master nested loops and string formatting by building number patterns. This exercise strengthens loop control, string concatenation, and understanding how to construct 2D patterns with code.

## Problem Description

You need to create three functions that generate different number patterns:

1. **Number Triangle** - Each row contains consecutive numbers from 1 up to the row number
2. **Reverse Pyramid** - Starts with all numbers and removes one each row
3. **Multiplication Table** - Classic n×n multiplication grid

These patterns require nested loops (loop within a loop) and careful attention to formatting, including newlines and spacing.

## Function Signatures

```go
func NumberTriangle(n int) string

func ReversePyramid(n int) string

func MultiplicationTable(n int) string
```

## Examples

### NumberTriangle(5)
```
1
12
123
1234
12345
```

### NumberTriangle(3)
```
1
12
123
```

### ReversePyramid(5)
```
12345
1234
123
12
1
```

### ReversePyramid(4)
```
1234
123
12
1
```

### MultiplicationTable(3)
```
1 2 3
2 4 6
3 6 9
```

### MultiplicationTable(5)
```
1 2 3 4 5
2 4 6 8 10
3 6 9 12 15
4 8 12 16 20
5 10 15 20 25
```

## Instructions

1. Implement `NumberTriangle`
2. Implement `ReversePyramid`
3. Implement `MultiplicationTable`
4. Run tests with `go test -v`

## Think About

1. Why do we use nested loops for these patterns? What does each loop control?

2. What happens if you forget the newline character at the end of each row?

3. For the multiplication table, why is it important to NOT add a trailing space after the last number in each row?

4. How would you modify NumberTriangle to create a right-aligned triangle instead?

5. What's the difference in performance between using `strings.Builder` vs string concatenation with `+=`? Does it matter for small values of n?

## What This Teaches

- **Nested loop control** - Understanding how to use loops within loops
- **Loop variable relationships** - Inner loop limit depends on outer loop variable
- **String building patterns** - Efficiently constructing multi-line strings
- **Format precision** - Exact spacing and newline placement
- **Pattern recognition** - Seeing mathematical relationships in visual patterns
- **Edge case handling** - What happens with n=0, n=1?
