# Exercise 02: Multiple Returns

**Learning Goal:** Master Go's multiple return value pattern, especially the `(result, error)` idiom

---

## 📝 Problem Description

Unlike many languages, Go allows functions to return multiple values. The most common pattern is returning `(result, error)`, which is the idiomatic way to handle errors in Go.

You'll implement several functions demonstrating multiple return patterns:
1. `Divide` - Division with error handling for divide-by-zero
2. `ParseName` - Extract first and last name from a full name string
3. `Stats` - Calculate min, max, and average from a slice of numbers
4. `FindIndex` - Locate an element in a slice, returning `(index, found)`

---

## 🎯 Function Signatures

```go
func Divide(a, b float64) (float64, error)

func ParseName(fullName string) (firstName string, lastName string, err error)

func Stats(numbers []int) (min int, max int, avg float64, err error)

func FindIndex(slice []int, target int) (int, bool)
```

---

## 📖 Examples

```go
// Divide examples
Divide(10, 2)     // 5.0, nil
Divide(7, 3)      // 2.333..., nil
Divide(5, 0)      // 0.0, error("division by zero")
Divide(0, 5)      // 0.0, nil

// ParseName examples
ParseName("John Doe")              // "John", "Doe", nil
ParseName("Mary Jane Watson")      // "Mary Jane", "Watson", nil
ParseName("Alice")                 // "", "", error("invalid name format")
ParseName("")                      // "", "", error("empty name")

// Stats examples
Stats([]int{5, 2, 9, 1, 7})       // 1, 9, 4.8, nil
Stats([]int{-5, -10, -2})         // -10, -2, -5.666..., nil
Stats([]int{42})                  // 42, 42, 42.0, nil
Stats([]int{})                    // 0, 0, 0.0, error("empty slice")

// FindIndex examples
FindIndex([]int{1, 2, 3}, 2)      // 1, true
FindIndex([]int{10, 20, 30}, 40)  // 0, false
FindIndex([]int{5}, 5)            // 0, true
FindIndex([]int{}, 1)             // 0, false
```

---

## 📋 Instructions

1. Implement `Divide` to return quotient or error for divide-by-zero
2. Implement `ParseName` to extract first and last names (last space splits the name)
3. Implement `Stats` to calculate min, max, average (or error if slice empty)
4. Implement `FindIndex` to find element position using the `(value, ok)` idiom
5. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Basic Concept</summary>

Multiple return values are declared in the function signature:
```go
func Example() (int, string, error) {
    return 42, "hello", nil
}

// Calling:
num, str, err := Example()
if err != nil {
    // handle error
}
```

Named return values create variables automatically:
```go
func Example() (result int, err error) {
    result = 42
    err = nil
    return // naked return uses named values
}
```

</details>

<details>
<summary>Intermediate Hint</summary>

For `ParseName`:
- Use `strings.LastIndex()` to find the last space
- Split at that position
- Handle edge cases: empty string, no space

For `Stats`:
- Return error immediately if slice is empty
- Initialize min/max to first element
- Calculate sum while finding min/max
- Average = sum / count (convert to float64)

For `FindIndex`:
- Iterate with index
- Return `(index, true)` when found
- Return `(0, false)` if not found

</details>

<details>
<summary>Complete Solution</summary>

```go
package multiple_returns

import (
	"errors"
	"strings"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func ParseName(fullName string) (firstName string, lastName string, err error) {
	fullName = strings.TrimSpace(fullName)
	if fullName == "" {
		return "", "", errors.New("empty name")
	}

	lastSpaceIndex := strings.LastIndex(fullName, " ")
	if lastSpaceIndex == -1 {
		return "", "", errors.New("invalid name format")
	}

	firstName = fullName[:lastSpaceIndex]
	lastName = fullName[lastSpaceIndex+1:]
	return firstName, lastName, nil
}

func Stats(numbers []int) (min int, max int, avg float64, err error) {
	if len(numbers) == 0 {
		return 0, 0, 0, errors.New("empty slice")
	}

	min = numbers[0]
	max = numbers[0]
	sum := 0

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += num
	}

	avg = float64(sum) / float64(len(numbers))
	return min, max, avg, nil
}

func FindIndex(slice []int, target int) (int, bool) {
	for i, val := range slice {
		if val == target {
			return i, true
		}
	}
	return 0, false
}
```

</details>

---

## 🤔 Think About

1. Why does Go use `(result, error)` instead of exceptions?
2. What's the benefit of named return values? Any downsides?
3. When should you use `(value, ok)` vs `(value, error)`?
4. How does the caller know to check the error first?

---

## 🎓 What This Teaches

- **Multiple returns**: Go's unique feature for returning multiple values
- **Error handling**: The idiomatic `(result, error)` pattern
- **Named returns**: Pre-declaring return variables in the signature
- **Boolean flags**: The `(value, ok)` idiom for existence checks
- **Error first**: Convention of checking errors before using results
- **Explicit errors**: No hidden control flow like exceptions

---

**Tier:** 1 - Foundation
**Estimated Time:** 30-40 minutes
