# Exercise 03: Subtests

**Learning Goal:** Use t.Run() to organize table-driven tests hierarchically and enable selective test running.

---

## Problem Description

Table-driven tests are great, but when a test fails, it can be hard to identify which specific case failed. Subtests solve this by giving each test case its own named scope and clear output.

With subtests, you can:
- See exactly which test case failed
- Run specific test cases individually
- Organize tests hierarchically
- Run tests in parallel safely

---

## Functions to Test

The `math.go` file provides:

```go
func Abs(x int) int
func Max(a, b int) int
func Clamp(value, min, max int) int
```

---

## Subtests with t.Run()

### Without Subtests (Exercise 02 style)

```go
func TestAbs(t *testing.T) {
    tests := []struct {
        name  string
        input int
        want  int
    }{
        {"positive", 5, 5},
        {"negative", -5, 5},
        {"zero", 0, 0},
    }

    for _, tt := range tests {
        got := Abs(tt.input)
        if got != tt.want {
            t.Errorf("Abs(%d) = %d, want %d", tt.input, got, tt.want)
            // Output: "Abs(-5) = -5, want 5"
            // Which test case was this? Need to look at the value
        }
    }
}
```

### With Subtests (Better!)

```go
func TestAbs(t *testing.T) {
    tests := []struct {
        name  string
        input int
        want  int
    }{
        {"positive", 5, 5},
        {"negative", -5, 5},
        {"zero", 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Abs(tt.input)
            if got != tt.want {
                t.Errorf("Abs(%d) = %d, want %d", tt.input, got, tt.want)
                // Output: "TestAbs/negative: Abs(-5) = -5, want 5"
                // Clear which case failed!
            }
        })
    }
}
```

---

## Benefits of Subtests

### 1. Clear Failure Messages

```bash
# Without subtests
go test -v
--- FAIL: TestAbs (0.00s)
    math_test.go:15: Abs(-5) = -5, want 5

# With subtests
go test -v
--- FAIL: TestAbs (0.00s)
    --- FAIL: TestAbs/negative (0.00s)
        math_test.go:15: Abs(-5) = -5, want 5
```

### 2. Run Specific Test Cases

```bash
# Run all Abs tests
go test -v -run TestAbs

# Run only the "negative" case
go test -v -run TestAbs/negative

# Run multiple specific cases with regex
go test -v -run TestAbs/positive|negative
```

### 3. Better Organization

Subtests create a hierarchy:
```
TestMath
├── TestMath/add
├── TestMath/subtract
└── TestMath/multiply
```

---

## Your Task

Convert the table-driven tests to use `t.Run()` for each test case. For each function:

1. Keep the table-driven structure
2. Wrap the test logic in `t.Run(tt.name, func(t *testing.T) { ... })`
3. Test selective running with `-run` flag
4. Observe the improved output with `-v` flag

---

## Examples

### Testing Max Function with Subtests

```go
func TestMax(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"a greater", 5, 3, 5},
        {"b greater", 3, 5, 5},
        {"equal", 4, 4, 4},
        {"negative numbers", -3, -5, -3},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Max(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Max(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

### Running Specific Subtests

```bash
# Run all Max tests
go test -v -run TestMax

# Run only "a greater" case
go test -v -run TestMax/a_greater

# Note: spaces in names become underscores in -run
# "a greater" → TestMax/a_greater
```

---

## Running Tests Selectively

```bash
# Run all tests in the package
go test -v

# Run specific test function
go test -v -run TestAbs

# Run specific subtest
go test -v -run TestAbs/negative

# Run multiple subtests with regex
go test -v -run TestMax/greater

# This matches both "a greater" and "b greater"
```

---

## Instructions

1. Open `math_test.go`
2. Follow TODO(human) markers to add `t.Run()` to each test
3. Create test tables with descriptive names
4. Run `go test -v` to see the hierarchical output
5. Practice running specific subtests with `-run`

**Experiment:**
- Run `go test -v` and observe the output structure
- Run `go test -v -run TestAbs/positive`
- Run `go test -v -run TestClamp/below`
- See how test names appear in the output

---

## Hints

<details>
<summary>Hint 1: The t.Run() Pattern</summary>

```go
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // Test logic here
        // Use tt.input, tt.want, etc.
    })
}
```

The key is wrapping the test logic in `t.Run()` with `tt.name`.
</details>

<details>
<summary>Hint 2: Descriptive Test Names</summary>

Good subtest names are:
- Clear: "negative number", not "test2"
- Concise: "a greater than b", not "when_the_first_parameter_is_greater"
- Lowercase with spaces: "edge case", not "EdgeCase"

Examples:
- ✅ "positive number"
- ✅ "zero"
- ✅ "negative number"
- ❌ "Test1"
- ❌ "test_positive_number"
</details>

<details>
<summary>Hint 3: Test Case Ideas</summary>

For **Abs**:
- Positive numbers
- Negative numbers
- Zero
- Large numbers

For **Max**:
- a > b
- b > a
- a == b
- Negative numbers
- Zero

For **Clamp**:
- Value within range
- Value below min
- Value above max
- Value equals min/max
- Min equals max
</details>

<details>
<summary>Full Solution</summary>

```go
package math

import "testing"

func TestAbs(t *testing.T) {
    tests := []struct {
        name  string
        input int
        want  int
    }{
        {"positive", 5, 5},
        {"negative", -5, 5},
        {"zero", 0, 0},
        {"large positive", 1000, 1000},
        {"large negative", -1000, 1000},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Abs(tt.input)
            if got != tt.want {
                t.Errorf("Abs(%d) = %d, want %d", tt.input, got, tt.want)
            }
        })
    }
}

func TestMax(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"a greater", 5, 3, 5},
        {"b greater", 3, 5, 5},
        {"equal", 4, 4, 4},
        {"negative numbers", -3, -5, -3},
        {"both zero", 0, 0, 0},
        {"one zero", 5, 0, 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Max(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Max(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}

func TestClamp(t *testing.T) {
    tests := []struct {
        name       string
        value      int
        min, max   int
        want       int
    }{
        {"within range", 5, 0, 10, 5},
        {"below minimum", -5, 0, 10, 0},
        {"above maximum", 15, 0, 10, 10},
        {"equals minimum", 0, 0, 10, 0},
        {"equals maximum", 10, 0, 10, 10},
        {"min equals max", 5, 5, 5, 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Clamp(tt.value, tt.min, tt.max)
            if got != tt.want {
                t.Errorf("Clamp(%d, %d, %d) = %d, want %d",
                    tt.value, tt.min, tt.max, got, tt.want)
            }
        })
    }
}
```
</details>

---

## What This Teaches

- **Subtests** - Organizing tests hierarchically with t.Run()
- **Test naming** - Clear, descriptive test case names
- **Selective running** - How to run specific test cases
- **Better debugging** - Clear failure messages showing exact case
- **Professional practice** - Standard pattern in Go codebases

---

## Think About

1. What happens if two test cases have the same name?
2. Why are spaces allowed in subtest names but not in test function names?
3. How would you organize subtests for a function with many edge cases?
4. When might you NOT want to use subtests?

---

**Next Exercise:** `04_test_helpers` - Reduce duplication with helper functions
