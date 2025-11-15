# Number Classifier

## Learning Goal

Master switch statements in Go, including multiple cases, default cases, and comparing switch vs if-else chains. Learn when switch statements improve code readability and maintainability.

## Problem Description

Classification is a common programming task: taking an input and categorizing it into different groups. Go's `switch` statement provides a clean, readable way to handle multiple conditional branches.

You'll write several classification functions that categorize numbers based on different properties: parity (even/odd), sign (positive/negative/zero), magnitude (size ranges), and day names from numbers.

## Function Signatures

```go
func ClassifyParity(n int) string
func ClassifySign(n int) string
func ClassifyMagnitude(n int) string
func DayName(n int) string
```

## Examples

```go
// Parity classification
ClassifyParity(4)   // => "even"
ClassifyParity(7)   // => "odd"
ClassifyParity(0)   // => "even"
ClassifyParity(-3)  // => "odd"

// Sign classification
ClassifySign(42)    // => "positive"
ClassifySign(-15)   // => "negative"
ClassifySign(0)     // => "zero"

// Magnitude classification
ClassifyMagnitude(5)    // => "small"   (0-10)
ClassifyMagnitude(50)   // => "medium"  (11-100)
ClassifyMagnitude(500)  // => "large"   (>100)
ClassifyMagnitude(-5)   // => "small"   (use absolute value)

// Day name from number
DayName(1)  // => "Monday"
DayName(5)  // => "Friday"
DayName(7)  // => "Sunday"
DayName(0)  // => "Invalid day"
DayName(8)  // => "Invalid day"
```

## Instructions

1. Implement `ClassifyParity`
2. Implement `ClassifySign`
3. Implement `ClassifyMagnitude`
4. Implement `DayName`
5. Run tests with `go test -v`

## Think About

1. **When to use switch vs if-else?** Which function would be harder to read with if-else chains?
2. **Does case order matter?** In `ClassifyMagnitude`, what happens if you check `absN <= 100` before `absN <= 10`?
3. **Fallthrough behavior:** Go does NOT fall through cases by default (unlike C/Java). Why is this safer?
4. **Performance:** Are switch statements faster than if-else chains? Does it matter for these functions?

## What This Teaches

- **Switch statement syntax**: Basic switch, tagless switch, default cases
- **Boolean expressions in switches**: Switch on `true` with case conditions
- **Code readability**: When switch statements improve clarity over if-else
- **Pattern matching**: Categorizing inputs into discrete groups
- **Input validation**: Handling invalid inputs with default cases
