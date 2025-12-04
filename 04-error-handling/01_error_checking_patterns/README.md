# Exercise 01: Error Checking Patterns

## 🎯 Learning Goal
Master the fundamental pattern of Go error handling: explicit checking with `if err != nil`, early returns to avoid nested conditionals, and proper error propagation up the call stack.

## 📝 Problem Description

Go's error handling is explicit and intentional. Unlike languages with try-catch blocks, Go treats errors as values that must be checked immediately after operations that can fail. This exercise teaches the idiomatic patterns you'll use thousands of times in Go code.

The key pattern is:
```go
result, err := operationThatCanFail()
if err != nil {
    // Handle or return the error
    return err
}
// Continue with result
```

## 🔧 Function Signatures

Implement these functions in `error_checking_patterns.go`:

```go
// Divide performs division and returns an error if divisor is zero
func Divide(a, b float64) (float64, error)

// ReadFileLength reads a file and returns its length in bytes
// Returns an error if the file doesn't exist or can't be read
func ReadFileLength(filename string) (int64, error)

// ParseAndDouble parses a string to int, doubles it, returns error if parsing fails
func ParseAndDouble(s string) (int, error)

// ChainOperations performs multiple operations that can fail
// It divides a by b, then multiplies result by c, then adds d
// Returns error from first operation that fails
func ChainOperations(a, b, c, d float64) (float64, error)

// ProcessFiles reads multiple files and returns total byte count
// Returns error immediately if any file fails to read
func ProcessFiles(filenames []string) (int64, error)

// SafeIndexAccess returns the element at index i, or error if out of bounds
func SafeIndexAccess(slice []int, i int) (int, error)
```

## 💡 Examples

```go
// Basic error checking
result, err := Divide(10, 2)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)  // 5.0

// Division by zero returns error
_, err = Divide(10, 0)
if err != nil {
    fmt.Println(err)  // "division by zero"
}

// File operations
length, err := ReadFileLength("testfile.txt")
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Printf("File size: %d bytes\n", length)

// Parse and process
doubled, err := ParseAndDouble("42")
if err != nil {
    return err
}
fmt.Println(doubled)  // 84

// Invalid input
_, err = ParseAndDouble("not a number")
// err: "invalid input: strconv.ParseInt: parsing \"not a number\": invalid syntax"

// Chain operations - stops at first error
result, err := ChainOperations(10, 2, 3, 5)  // (10/2)*3 + 5 = 20
result, err = ChainOperations(10, 0, 3, 5)   // Error: division by zero

// Multiple files
total, err := ProcessFiles([]string{"file1.txt", "file2.txt"})
if err != nil {
    fmt.Println("Failed to process files:", err)
}

// Safe array access
val, err := SafeIndexAccess([]int{1, 2, 3}, 1)  // 2, nil
val, err = SafeIndexAccess([]int{1, 2, 3}, 10) // 0, error
```

## 📋 Instructions

1. **Divide:** Check if b is zero before dividing. Return an error with a descriptive message if so.
2. **ReadFileLength:** Use `os.Stat(filename)` to get file info. Check the error, return it if present. Otherwise return `fileInfo.Size()`.
3. **ParseAndDouble:** Use `strconv.ParseInt(s, 10, 64)` to parse the string. Check error immediately. If successful, double the value and return.
4. **ChainOperations:** Call Divide(a, b), check error. Then multiply result by c. Then add d. Return error from any step that fails.
5. **ProcessFiles:** Loop through filenames, call ReadFileLength for each. Accumulate total. Return immediately if any file fails.
6. **SafeIndexAccess:** Check if i is within bounds (`i >= 0 && i < len(slice)`). Return error if not, otherwise return element.

## 🧪 Testing

Run tests with:
```bash
go test -v
```

Expected: ~30-35 test cases covering happy paths and error conditions

## 🤔 Think About

1. **Why check errors immediately?**
   - Prevents continuing with invalid state
   - Makes error handling explicit and visible
   - Avoids deep nesting of conditionals

2. **Why return errors instead of handling them?**
   - Lets callers decide how to handle errors
   - Maintains single responsibility principle
   - Enables error wrapping with context

3. **What's the "early return" pattern?**
   - Check error, return immediately if present
   - Avoids else blocks and reduces indentation
   - Makes the "happy path" visually clear

4. **Should every function return an error?**
   - Only if it can fail in a way the caller should handle
   - Constructor functions often don't return errors
   - Pure functions without I/O rarely need errors

## 💡 Hints

<details>
<summary>Hint 1: Basic error pattern</summary>

The fundamental Go error pattern:
```go
result, err := operation()
if err != nil {
    return defaultValue, err  // Or handle it
}
// Use result here
```

**Key points:**
- Check error immediately after the call
- Return early if error present
- Continue with result only if no error
</details>

<details>
<summary>Hint 2: Creating errors</summary>

Use `errors.New()` or `fmt.Errorf()` to create errors:
```go
import "errors"

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

Import `errors` package at the top of your file.
</details>

<details>
<summary>Hint 3: File operations</summary>

Use `os.Stat()` to get file information:
```go
import "os"

func ReadFileLength(filename string) (int64, error) {
    info, err := os.Stat(filename)
    if err != nil {
        return 0, err  // Propagate the error
    }
    return info.Size(), nil
}
```

The error from `os.Stat` already contains useful information, so just return it.
</details>

<details>
<summary>Hint 4: String to int parsing</summary>

Use `strconv.ParseInt`:
```go
import "strconv"

func ParseAndDouble(s string) (int, error) {
    n, err := strconv.ParseInt(s, 10, 64)
    if err != nil {
        return 0, fmt.Errorf("invalid input: %v", err)
    }
    return int(n * 2), nil
}
```

Add context to parsing errors so users know what failed.
</details>

<details>
<summary>Hint 5: Chaining operations</summary>

Chain operations by checking each step:
```go
func ChainOperations(a, b, c, d float64) (float64, error) {
    // Step 1: Divide
    result, err := Divide(a, b)
    if err != nil {
        return 0, err
    }

    // Step 2: Multiply
    result = result * c

    // Step 3: Add
    result = result + d

    return result, nil
}
```

Stop at the first error and return it immediately.
</details>

<details>
<summary>Hint 6: Loop with error checking</summary>

Process multiple items, fail fast on first error:
```go
func ProcessFiles(filenames []string) (int64, error) {
    var total int64
    for _, filename := range filenames {
        length, err := ReadFileLength(filename)
        if err != nil {
            return 0, err  // Stop on first error
        }
        total += length
    }
    return total, nil
}
```

Don't continue processing if one item fails.
</details>

<details>
<summary>Hint 7: Bounds checking</summary>

Check array bounds before accessing:
```go
func SafeIndexAccess(slice []int, i int) (int, error) {
    if i < 0 || i >= len(slice) {
        return 0, fmt.Errorf("index %d out of bounds (len=%d)", i, len(slice))
    }
    return slice[i], nil
}
```

Provide helpful error messages with the invalid index and slice length.
</details>

<details>
<summary>Full Solution</summary>

```go
package error_checking_patterns

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func ReadFileLength(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func ParseAndDouble(s string) (int, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %v", err)
	}
	return int(n * 2), nil
}

func ChainOperations(a, b, c, d float64) (float64, error) {
	result, err := Divide(a, b)
	if err != nil {
		return 0, err
	}
	result = result * c
	result = result + d
	return result, nil
}

func ProcessFiles(filenames []string) (int64, error) {
	var total int64
	for _, filename := range filenames {
		length, err := ReadFileLength(filename)
		if err != nil {
			return 0, err
		}
		total += length
	}
	return total, nil
}

func SafeIndexAccess(slice []int, i int) (int, error) {
	if i < 0 || i >= len(slice) {
		return 0, fmt.Errorf("index %d out of bounds (len=%d)", i, len(slice))
	}
	return slice[i], nil
}
```
</details>

## 🎓 What This Teaches

- **Explicit error checking** - The `if err != nil` idiom you'll use constantly
- **Early returns** - Return errors immediately to avoid nested conditionals
- **Error propagation** - Passing errors up the call stack to let callers decide
- **Fail-fast principle** - Stop processing when an error occurs
- **Zero values for errors** - Return zero/empty values along with errors
- **Error creation** - Using errors.New and fmt.Errorf to create errors
- **Descriptive error messages** - Making errors actionable with context

---

**Next Exercise:** `02_error_creation` - Creating meaningful errors with errors.New and fmt.Errorf
