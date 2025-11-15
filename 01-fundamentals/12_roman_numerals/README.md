# Exercise 12: Roman Numerals

## Learning Goal
Master switch statements, string building, and pattern recognition by implementing Roman numeral conversion in both directions.

## Problem Description
Roman numerals use seven symbols to represent numbers:
- I = 1
- V = 5
- X = 10
- L = 50
- C = 100
- D = 500
- M = 1000

Numbers are formed by combining symbols and adding values. However, there are special subtraction cases:
- I before V or X makes 4 (IV) or 9 (IX)
- X before L or C makes 40 (XL) or 90 (XC)
- C before D or M makes 400 (CD) or 900 (CM)

Your task is to convert decimal numbers (1-3999) to Roman numerals and vice versa.

## Function Signatures
```go
func ToRoman(n int) string
func FromRoman(s string) int
```

## Examples

### ToRoman
```go
ToRoman(1)    // "I"
ToRoman(4)    // "IV"
ToRoman(9)    // "IX"
ToRoman(58)   // "LVIII"
ToRoman(1994) // "MCMXCIV"
ToRoman(3999) // "MMMCMXCIX"
```

### FromRoman
```go
FromRoman("I")       // 1
FromRoman("IV")      // 4
FromRoman("IX")      // 9
FromRoman("LVIII")   // 58
FromRoman("MCMXCIV") // 1994
FromRoman("MMMCMXCIX") // 3999
```

## Instructions

1. Implement `ToRoman`
2. Implement `FromRoman`
3. Run tests with `go test -v`

## Think About

1. Why is it important to process values from largest to smallest in ToRoman?
2. How does the "look ahead" or "look back" strategy help identify subtraction cases?
3. What are the edge cases you need to handle (e.g., minimum value 1, maximum value 3999)?
4. Could you implement this with a switch statement instead of a map? What are the trade-offs?

## What This Teaches

- **Switch statements and conditionals** - Pattern matching for conversion logic
- **String building** - Using strings.Builder for efficient concatenation
- **Algorithm design** - Greedy approach (largest values first) for ToRoman
- **State tracking** - Comparing current and previous values for FromRoman
- **Data structures** - Using maps, slices, or parallel arrays for lookups
- **Edge case handling** - Validating input ranges and handling boundaries
