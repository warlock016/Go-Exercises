# Exercise 05: Grade Evaluator

**Difficulty:** Easy-Medium
**Concepts:** If/else conditionals, else-if chains, comparison operators, range checking
**Estimated Time:** 20-30 minutes

## Learning Goals

- Use if/else statements to make decisions based on conditions
- Chain multiple conditions with else-if
- Use comparison operators (<, <=, >, >=, ==, !=)
- Check if values fall within ranges
- Return different values based on conditions

## Problem Description

You're building a grading system for a school. Teachers need to convert numeric scores into letter grades, determine if students are passing, and calculate class averages.

Your task is to implement several functions that evaluate grades using different grading schemes:
- Basic letter grading (A, B, C, D, F)
- Pass/fail determination
- Detailed grading with plus/minus modifiers (A+, A-, B+, etc.)
- Class average calculation with letter grade

## Function Signatures

```go
func LetterGrade(score int) string

func IsPassing(score int) bool

func GradeWithPlus(score int) string

func ClassAverage(scores []int) (average float64, grade string)
```

## Examples

### LetterGrade
```go
LetterGrade(95)  // "A"
LetterGrade(87)  // "B"
LetterGrade(73)  // "C"
LetterGrade(65)  // "D"
LetterGrade(42)  // "F"
```

### IsPassing
```go
IsPassing(85)  // true
IsPassing(60)  // true (60 is the minimum passing score)
IsPassing(59)  // false
IsPassing(0)   // false
```

### GradeWithPlus
```go
GradeWithPlus(98)  // "A+"
GradeWithPlus(92)  // "A"
GradeWithPlus(90)  // "A-"
GradeWithPlus(88)  // "B+"
GradeWithPlus(82)  // "B"
GradeWithPlus(80)  // "B-"
GradeWithPlus(55)  // "F" (no F+ or F-)
```

**Grading Scale with Plus/Minus:**
- **A+:** 97-100
- **A:** 93-96
- **A-:** 90-92
- **B+:** 87-89
- **B:** 83-86
- **B-:** 80-82
- **C+:** 77-79
- **C:** 73-76
- **C-:** 70-72
- **D+:** 67-69
- **D:** 63-66
- **D-:** 60-62
- **F:** 0-59

### ClassAverage
```go
ClassAverage([]int{90, 85, 78, 92, 88})  // (86.6, "B")
ClassAverage([]int{100, 95, 90})         // (95.0, "A")
ClassAverage([]int{60, 65, 70})          // (65.0, "D")
ClassAverage([]int{})                    // (0.0, "F")
```

## Instructions

1. Implement `LetterGrade(score int) string`
2. Implement `IsPassing(score int) bool`
3. Implement `GradeWithPlus(score int) string`
4. Implement `ClassAverage(scores []int) (average float64, grade string)`

## Think About

1. Why do we start with the highest grade when using if/else-if chains?
2. What would happen if you checked `score >= 60` before checking `score >= 90`?
3. Why does `IsPassing` work with just a single comparison while `LetterGrade` needs multiple?
4. In `ClassAverage`, why do we convert to `float64` when dividing?

## What This Teaches

- **Conditional logic:** Making decisions based on data values
- **Range checking:** Determining if values fall within specific ranges
- **If/else-if chains:** Handling multiple mutually exclusive conditions
- **Order of conditions:** Why the sequence of checks matters
- **Type conversion:** Converting between int and float64 for accurate calculations
- **Named return values:** Returning multiple values with clear names
- **Reusing functions:** Calling one function from another to avoid duplication

## Running Tests

```bash
go test -v
```

## Success Criteria

- All tests pass
- Code uses if/else-if chains appropriately
- Conditions are ordered correctly (highest to lowest)
- No redundant comparisons (like `score >= 90 && score <= 100` in else-if chains)
- ClassAverage handles empty slice correctly
