# Exercise 02: Package Organization

**Concept:** Multi-file packages and code organization
**Difficulty:** Easy
**Estimated Time:** 25 minutes

## 🎯 Learning Goal

Understand that a package can span multiple files, and how to organize related functionality across files.

## The Problem

As code grows, putting everything in one file becomes unwieldy. Go allows you to split a package across multiple files. Key insights:

1. **All files in a directory = same package** (usually)
2. **Package members are shared** across all files
3. **No need to import** between files in the same package
4. **Unexported identifiers are accessible** across files in the same package

## Your Task

Create a `calculator` sub-package organized across three files:

### File Structure
```
02_package_organization/
├── README.md
├── package_organization.go      # Main file (imports calculator)
├── package_organization_test.go
└── calculator/                   # Sub-package
    ├── calculator.go             # Basic operations
    ├── operations.go             # Advanced operations
    └── validation.go             # Internal helpers
```

### calculator/calculator.go
- `Add(a, b int) int` - exported
- `Subtract(a, b int) int` - exported

### calculator/operations.go
- `Multiply(a, b int) int` - exported
- `Divide(a, b int) (int, error)` - exported

### calculator/validation.go
- `validate(a, b int, op string) error` - unexported helper
- Used by Divide to check for division by zero

## Function Signatures

```go
// calculator/calculator.go
package calculator

func Add(a, b int) int
func Subtract(a, b int) int

// calculator/operations.go
package calculator  // Same package!

func Multiply(a, b int) int
func Divide(a, b int) (int, error)

// calculator/validation.go
package calculator  // Same package!

func validate(a, b int, op string) error  // unexported
```

## Examples

```go
import "path/to/calculator"

result := calculator.Add(5, 3)        // 8
result = calculator.Multiply(4, 7)    // 28

quotient, err := calculator.Divide(10, 2)
if err != nil {
    // Handle error
}
// quotient = 5

_, err = calculator.Divide(10, 0)
// err != nil (division by zero)

// calculator.validate(...) // COMPILE ERROR - unexported
```

## Instructions

1. Create the `calculator/` directory
2. Create `calculator.go` with Add and Subtract
3. Create `operations.go` with Multiply and Divide
4. Create `validation.go` with validate helper
5. Have Divide use validate to check for division by zero
6. Update `package_organization.go` to import and use calculator
7. Run `go test -v`

## Hints

### Basic - Multi-file Package Structure
- All three files start with `package calculator`
- They can use each other's functions directly (even unexported ones)
- External packages import once: `import "path/to/calculator"`

### Intermediate - Organization Pattern
```go
// calculator/calculator.go
package calculator

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}

// calculator/operations.go
package calculator  // Same package name!

func Multiply(a, b int) int {
    return a * b
}

func Divide(a, b int) (int, error) {
    // Call validate() - it's in the same package
    if err := validate(a, b, "divide"); err != nil {
        return 0, err
    }
    return a / b, nil
}

// calculator/validation.go
package calculator  // Same package name!

import "errors"

func validate(a, b int, op string) error {
    if op == "divide" && b == 0 {
        return errors.New("division by zero")
    }
    return nil
}
```

### Advanced - Using the Package
```go
// package_organization.go
package organization

import "path/to/07-packages-modules/02_package_organization/calculator"

func Calculate(a, b int, operation string) (int, error) {
    switch operation {
    case "add":
        return calculator.Add(a, b), nil
    case "subtract":
        return calculator.Subtract(a, b), nil
    case "multiply":
        return calculator.Multiply(a, b), nil
    case "divide":
        return calculator.Divide(a, b)
    default:
        return 0, errors.New("unknown operation")
    }
}
```

## 🧠 Think About

1. Why split code across multiple files instead of one large file?
2. How does Go know which files belong to the same package?
3. Can two packages exist in the same directory?
4. Why can `Divide()` call unexported `validate()` from a different file?
5. How would you decide what goes in which file?

## What This Teaches

- Multi-file package organization
- Package members are shared across all files
- Unexported functions are accessible within the package
- How to structure a package for readability
- Import paths point to directories (packages), not files

## After Completing

Write in your `EXPLANATION.md`:
- How Go determines package boundaries
- Why validate() can be called from operations.go
- Benefits of splitting code across files
- How you decided what to put in each file
