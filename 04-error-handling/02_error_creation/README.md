# Exercise 02: Error Creation

## 🎯 Learning Goal
Learn to create meaningful, actionable errors using `errors.New` for static messages and `fmt.Errorf` for dynamic error messages with context and formatting.

## 📝 Problem Description

Good errors tell you exactly what went wrong and why. Go provides two primary ways to create errors:

1. **`errors.New("static message")`** - For fixed error messages
2. **`fmt.Errorf("format %s", args...)`** - For dynamic messages with values

This exercise teaches you to create errors that help users understand and fix problems, not just report failures.

## 🔧 Function Signatures

Implement these functions in `error_creation.go`:

```go
// ValidateAge returns an error if age is negative or unrealistic
func ValidateAge(age int) error

// ValidateEmail returns an error if email doesn't contain @
func ValidateEmail(email string) error

// WithdrawMoney returns error if amount is negative or exceeds balance
func WithdrawMoney(balance, amount float64) (float64, error)

// FormatUserError creates a descriptive error message for user operations
// operation: "create", "update", "delete", etc.
// username: the username being operated on
// reason: why the operation failed
func FormatUserError(operation, username, reason string) error

// ValidatePassword returns detailed error describing password requirements
// Requirements: min 8 chars, at least 1 digit, at least 1 uppercase
func ValidatePassword(password string) error

// ParseConfig returns error with filename and line number if parsing fails
func ParseConfig(filename string, lineNumber int, content string) error
```

## 💡 Examples

```go
// Simple validation
err := ValidateAge(-5)
// Error: "age cannot be negative: -5"

err = ValidateAge(200)
// Error: "age must be between 0 and 150, got 200"

// Email validation
err = ValidateEmail("invalidemail")
// Error: "invalid email: must contain @"

// Financial operations
newBalance, err := WithdrawMoney(100, 150)
// Error: "insufficient funds: balance 100.00, withdrawal 150.00"

newBalance, err = WithdrawMoney(100, -50)
// Error: "withdrawal amount must be positive, got -50.00"

// Formatted errors with context
err = FormatUserError("delete", "john_doe", "user not found")
// Error: "failed to delete user 'john_doe': user not found"

// Detailed validation
err = ValidatePassword("pass")
// Error: "password must be at least 8 characters, got 4"

err = ValidatePassword("password")
// Error: "password must contain at least 1 digit"

// Config parsing errors
err = ParseConfig("app.yaml", 42, "invalid: [value")
// Error: "failed to parse app.yaml at line 42: invalid: [value"
```

## 📋 Instructions

1. **ValidateAge:** Check if age < 0 (return error with the negative value) or age > 150 (return error with max allowed and actual value)
2. **ValidateEmail:** Check if email contains "@". Use `strings.Contains()`
3. **WithdrawMoney:** Check if amount <= 0 (format with %.2f) or amount > balance (show both balance and amount in error)
4. **FormatUserError:** Use fmt.Errorf to create a consistent message: "failed to {operation} user '{username}': {reason}"
5. **ValidatePassword:** Check length >= 8, contains digit (loop + unicode.IsDigit), contains uppercase (unicode.IsUpper). Return first failing requirement.
6. **ParseConfig:** Use fmt.Errorf to include filename, line number, and content in the error message

## 🧪 Testing

Run tests with:
```bash
go test -v
```

Expected: ~30-40 test cases covering various error scenarios

## 🤔 Think About

1. **What makes a good error message?**
   - Specific about what went wrong
   - Includes relevant values/context
   - Actionable (user knows how to fix it)
   - Consistent format across the application

2. **When to use errors.New vs fmt.Errorf?**
   - errors.New: Static messages that never change
   - fmt.Errorf: Dynamic messages with runtime values

3. **Should errors include implementation details?**
   - Generally no - focus on what the user needs to know
   - Exception: debugging errors for developers

4. **How detailed should errors be?**
   - Balance between helpful and verbose
   - Include context needed to understand and fix the issue
   - Don't expose security-sensitive information

## 💡 Hints

<details>
<summary>Hint 1: Using errors.New</summary>

For static error messages:
```go
import "errors"

if someCondition {
    return errors.New("simple error message")
}
```

Best for messages that don't need runtime values.
</details>

<details>
<summary>Hint 2: Using fmt.Errorf</summary>

For dynamic error messages with values:
```go
import "fmt"

if age < 0 {
    return fmt.Errorf("age cannot be negative: %d", age)
}
```

Supports all fmt format verbs: %d (int), %s (string), %f (float), %.2f (float with 2 decimals), %v (any)
</details>

<details>
<summary>Hint 3: Multiple conditions</summary>

Check multiple conditions, return specific errors:
```go
func ValidateAge(age int) error {
    if age < 0 {
        return fmt.Errorf("age cannot be negative: %d", age)
    }
    if age > 150 {
        return fmt.Errorf("age must be between 0 and 150, got %d", age)
    }
    return nil  // No error
}
```
</details>

<details>
<summary>Hint 4: String validation</summary>

Use strings package for checks:
```go
import "strings"

func ValidateEmail(email string) error {
    if !strings.Contains(email, "@") {
        return errors.New("invalid email: must contain @")
    }
    return nil
}
```
</details>

<details>
<summary>Hint 5: Checking for digits and uppercase</summary>

Use unicode package:
```go
import "unicode"

func hasDigit(s string) bool {
    for _, ch := range s {
        if unicode.IsDigit(ch) {
            return true
        }
    }
    return false
}

func hasUpper(s string) bool {
    for _, ch := range s {
        if unicode.IsUpper(ch) {
            return true
        }
    }
    return false
}
```
</details>

<details>
<summary>Full Solution</summary>

```go
package error_creation

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func ValidateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("age cannot be negative: %d", age)
	}
	if age > 150 {
		return fmt.Errorf("age must be between 0 and 150, got %d", age)
	}
	return nil
}

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return errors.New("invalid email: must contain @")
	}
	return nil
}

func WithdrawMoney(balance, amount float64) (float64, error) {
	if amount <= 0 {
		return balance, fmt.Errorf("withdrawal amount must be positive, got %.2f", amount)
	}
	if amount > balance {
		return balance, fmt.Errorf("insufficient funds: balance %.2f, withdrawal %.2f", balance, amount)
	}
	return balance - amount, nil
}

func FormatUserError(operation, username, reason string) error {
	return fmt.Errorf("failed to %s user '%s': %s", operation, username, reason)
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters, got %d", len(password))
	}

	hasDigit := false
	for _, ch := range password {
		if unicode.IsDigit(ch) {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		return errors.New("password must contain at least 1 digit")
	}

	hasUpper := false
	for _, ch := range password {
		if unicode.IsUpper(ch) {
			hasUpper = true
			break
		}
	}
	if !hasUpper {
		return errors.New("password must contain at least 1 uppercase letter")
	}

	return nil
}

func ParseConfig(filename string, lineNumber int, content string) error {
	return fmt.Errorf("failed to parse %s at line %d: %s", filename, lineNumber, content)
}
```
</details>

## 🎓 What This Teaches

- **errors.New** - Creating simple, static error messages
- **fmt.Errorf** - Creating formatted error messages with runtime values
- **Format verbs** - %d for ints, %s for strings, %.2f for floats with precision
- **Error context** - Including relevant information in error messages
- **Validation patterns** - Checking preconditions and returning descriptive errors
- **User-friendly errors** - Writing errors that help users understand what went wrong
- **Consistency** - Using consistent error message formats across functions

---

**Next Exercise:** `03_sentinel_errors` - Package-level error variables and errors.Is
