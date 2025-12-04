# Exercise 02: Table-Driven Test Basics

**Learning Goal:** Master the idiomatic Go testing pattern - table-driven tests with struct slices.

---

## Problem Description

You have string utility functions that need comprehensive testing. Writing individual test functions for each case (like in Exercise 01) becomes repetitive. Table-driven tests solve this by separating **test data** from **test logic**.

This is THE standard pattern used throughout the Go standard library and professional Go codebases.

---

## Functions to Test

The `stringutils.go` file provides:

```go
func Reverse(s string) string
func IsPalindrome(s string) bool
func CountVowels(s string) int
```

---

## Table-Driven Test Pattern

Instead of writing many similar test functions:

```go
// ❌ Repetitive approach (don't do this)
func TestReverse1(t *testing.T) {
    result := Reverse("hello")
    if result != "olleh" {
        t.Errorf("...")
    }
}

func TestReverse2(t *testing.T) {
    result := Reverse("Go")
    if result != "oG" {
        t.Errorf("...")
    }
}
// ... many more similar functions
```

Use a test table:

```go
// ✅ Table-driven approach (idiomatic Go)
func TestReverse(t *testing.T) {
    tests := []struct {
        name  string
        input string
        want  string
    }{
        {"simple word", "hello", "olleh"},
        {"two letters", "Go", "oG"},
        {"empty string", "", ""},
        {"single char", "a", "a"},
    }

    for _, tt := range tests {
        got := Reverse(tt.input)
        if got != tt.want {
            t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.want)
        }
    }
}
```

---

## Your Task

Convert the tests in `stringutils_test.go` to use table-driven patterns. For each function:

1. Define a test table (slice of structs)
2. Each struct has: name, input(s), expected output
3. Loop over the table with `for _, tt := range tests`
4. Run the function and compare result
5. Include at least 5-6 test cases per function

---

## Test Table Structure

### Anatomy of a Test Table

```go
tests := []struct {
    name   string  // Description of this test case
    input  Type    // Input parameter(s)
    want   Type    // Expected result
    // Can add more fields as needed
}{
    {"description", inputValue, expectedValue},
    {"another case", inputValue, expectedValue},
}
```

### For Functions with Multiple Inputs

```go
tests := []struct {
    name string
    a    int
    b    int
    want int
}{
    {"positive numbers", 2, 3, 5},
    {"negative numbers", -2, -3, -5},
}
```

### For Functions Returning Errors

```go
tests := []struct {
    name    string
    input   string
    want    int
    wantErr bool
}{
    {"valid input", "123", 123, false},
    {"invalid input", "abc", 0, true},
}
```

---

## Examples

### Testing Reverse Function

```go
func TestReverse(t *testing.T) {
    tests := []struct {
        name  string
        input string
        want  string
    }{
        {"simple word", "hello", "olleh"},
        {"uppercase", "GOLANG", "GNALOG"},
        {"empty string", "", ""},
        {"single character", "x", "x"},
        {"palindrome", "racecar", "racecar"},
        {"with spaces", "hello world", "dlrow olleh"},
    }

    for _, tt := range tests {
        got := Reverse(tt.input)
        if got != tt.want {
            t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.want)
        }
    }
}
```

### Testing IsPalindrome Function

```go
func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        name  string
        input string
        want  bool
    }{
        {"simple palindrome", "racecar", true},
        {"not palindrome", "hello", false},
        {"empty string", "", true},
        {"single character", "a", true},
    }

    for _, tt := range tests {
        got := IsPalindrome(tt.input)
        if got != tt.want {
            t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
        }
    }
}
```

---

## Why Table-Driven Tests?

**Advantages:**
1. **Add cases easily** - Just add a line to the table
2. **Reduce duplication** - Test logic written once
3. **Readable** - Each case is clear and self-documenting
4. **Comprehensive** - Easy to add edge cases
5. **Standard** - What professional Go developers use

**When to use:**
- When testing the same function with different inputs
- When you have many similar test cases
- When you want comprehensive coverage

---

## Running Tests

```bash
# Run all tests
go test -v

# See how table-driven tests appear in output
go test -v

# Run specific test function
go test -v -run TestReverse
```

---

## Instructions

1. Open `stringutils_test.go`
2. Follow TODO(human) markers to create test tables
3. For each function, create a table with 5-6 test cases
4. Think about edge cases: empty strings, single characters, special cases
5. Run `go test -v` to verify all cases pass

**Test Cases to Include:**
- Normal cases (typical inputs)
- Edge cases (empty, single character)
- Special cases (spaces, punctuation for relevant functions)
- Boundary conditions

---

## Hints

<details>
<summary>Hint 1: Table Structure</summary>

The basic pattern is always:
1. Create slice of anonymous structs
2. Each struct has name, inputs, expected output
3. Loop with `for _, tt := range tests`
4. Call function with `tt.input`
5. Compare `got` with `tt.want`
</details>

<details>
<summary>Hint 2: Test Case Ideas</summary>

For **Reverse**:
- Normal words
- Empty string
- Single character
- Palindromes
- Strings with spaces

For **IsPalindrome**:
- True palindromes: "racecar", "noon"
- False palindromes: "hello", "golang"
- Edge cases: "", "a"
- Consider case sensitivity

For **CountVowels**:
- Strings with vowels: "hello" → 2
- No vowels: "xyz" → 0
- All vowels: "aeiou" → 5
- Empty string → 0
- Mixed case
</details>

<details>
<summary>Hint 3: The Loop Pattern</summary>

```go
for _, tt := range tests {
    got := FunctionToTest(tt.input)
    if got != tt.want {
        t.Errorf("FunctionToTest(%q) = %v, want %v",
            tt.input, got, tt.want)
    }
}
```

The variable name `tt` is idiomatic Go for "table test" or "this test".
</details>

<details>
<summary>Full Solution</summary>

```go
package stringutils

import "testing"

func TestReverse(t *testing.T) {
    tests := []struct {
        name  string
        input string
        want  string
    }{
        {"simple word", "hello", "olleh"},
        {"uppercase", "GOLANG", "GNALOG"},
        {"empty string", "", ""},
        {"single character", "x", "x"},
        {"palindrome", "racecar", "racecar"},
        {"with spaces", "hello world", "dlrow olleh"},
    }

    for _, tt := range tests {
        got := Reverse(tt.input)
        if got != tt.want {
            t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.want)
        }
    }
}

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        name  string
        input string
        want  bool
    }{
        {"simple palindrome", "racecar", true},
        {"not palindrome", "hello", false},
        {"empty string", "", true},
        {"single character", "a", true},
        {"two same letters", "aa", true},
        {"two different letters", "ab", false},
        {"longer palindrome", "noon", true},
    }

    for _, tt := range tests {
        got := IsPalindrome(tt.input)
        if got != tt.want {
            t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
        }
    }
}

func TestCountVowels(t *testing.T) {
    tests := []struct {
        name  string
        input string
        want  int
    }{
        {"simple word", "hello", 2},
        {"no vowels", "xyz", 0},
        {"all vowels", "aeiou", 5},
        {"empty string", "", 0},
        {"uppercase vowels", "AEIOU", 5},
        {"mixed case", "Hello World", 3},
        {"single vowel", "a", 1},
    }

    for _, tt := range tests {
        got := CountVowels(tt.input)
        if got != tt.want {
            t.Errorf("CountVowels(%q) = %d, want %d", tt.input, got, tt.want)
        }
    }
}
```
</details>

---

## What This Teaches

- **Table-driven testing** - The Go way to test comprehensively
- **Test organization** - Separating data from logic
- **Comprehensive coverage** - Easy to add many test cases
- **Readability** - Self-documenting test cases
- **Professional patterns** - What you'll see in real Go codebases

---

## Think About

1. What happens if you forget the `name` field in the test struct?
2. How would you add a test case that expects an error?
3. Why is `tt` used as the loop variable name?
4. How many test cases is "enough" for good coverage?
5. Compare this to Exercise 01 - which approach do you prefer and why?

---

**Next Exercise:** `03_subtests` - Organize table-driven tests with t.Run()
