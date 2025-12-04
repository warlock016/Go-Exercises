# Exercise 01: Variadic Functions

**Learning Goal:** Master variadic parameters to write functions accepting variable numbers of arguments

---

## 📝 Problem Description

Variadic functions accept zero or more arguments of a specified type using the `...` syntax. This is useful for functions like `fmt.Println()` that need flexibility in the number of parameters.

You'll implement three variadic functions:
1. `Sum` - Add any number of integers
2. `Max` - Find the maximum value from multiple integers
3. `Concat` - Join any number of strings with a separator

---

## 🎯 Function Signatures

```go
func Sum(nums ...int) int

func Max(nums ...int) (int, error)

func Concat(separator string, parts ...string) string
```

---

## 📖 Examples

```go
// Sum examples
Sum()           // 0
Sum(5)          // 5
Sum(1, 2, 3)    // 6
Sum(10, -5, 3)  // 8

// Max examples
Max(5)                  // 5, nil
Max(1, 9, 3, 7)         // 9, nil
Max(-5, -10, -2)        // -2, nil
Max()                   // 0, error("no values provided")

// Concat examples
Concat(", ")                      // ""
Concat(", ", "apple")             // "apple"
Concat(", ", "a", "b", "c")       // "a, b, c"
Concat(" - ", "Go", "is", "fun")  // "Go - is - fun"
Concat("/", "usr", "local", "bin") // "usr/local/bin"
```

---

## 📋 Instructions

1. Implement `Sum` that returns the sum of all arguments (return 0 if empty)
2. Implement `Max` that returns the maximum value or an error if no values provided
3. Implement `Concat` that joins strings with the separator (separator is NOT variadic)
4. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Basic Concept</summary>

Variadic parameters use `...Type` and are treated as a slice inside the function:
```go
func Example(values ...int) {
    // values is a []int slice
    for _, v := range values {
        // process each value
    }
}
```

</details>

<details>
<summary>Intermediate Hint</summary>

For `Max`:
- Check if the slice is empty first (length zero)
- Return an error if no values
- Initialize max to the first element
- Iterate through remaining elements

For `Concat`:
- Only `parts` is variadic, not `separator`
- Check the length of `parts`
- Build the result string incrementally

</details>

<details>
<summary>Complete Solution</summary>

```go
package variadic_functions

import (
	"errors"
	"strings"
)

func Sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Max(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("no values provided")
	}

	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max, nil
}

func Concat(separator string, parts ...string) string {
	if len(parts) == 0 {
		return ""
	}

	var result strings.Builder
	result.WriteString(parts[0])

	for _, part := range parts[1:] {
		result.WriteString(separator)
		result.WriteString(part)
	}

	return result.String()
}
```

</details>

---

## 🤔 Think About

1. Why is `separator` in `Concat` not variadic?
2. What happens when you call `Sum()` with no arguments?
3. How would you pass a slice to a variadic function? (Hint: `Sum(slice...`)
4. When would you use variadic parameters vs accepting a slice parameter?

---

## 🎓 What This Teaches

- **Variadic syntax**: Using `...Type` for flexible parameter lists
- **Slice behavior**: Variadic parameters become slices inside the function
- **Error handling**: Returning errors when input is invalid (empty slice)
- **Parameter order**: Variadic parameters must be last
- **Empty cases**: Handling zero arguments gracefully
- **Standard library patterns**: How `fmt.Println()`, `append()`, etc. work

---

**Tier:** 1 - Foundation
**Estimated Time:** 20-30 minutes
