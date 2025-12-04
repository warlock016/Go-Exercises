# Exercise 04: Test Helpers

**Learning Goal:** Use t.Helper() to create reusable test helper functions and understand setup/teardown patterns.

---

## Problem Description

As tests grow, you'll notice repeated patterns - checking errors, comparing values, setting up test data. Test helpers extract this common code into reusable functions. The `t.Helper()` method ensures error messages point to the correct line in your test.

This exercise teaches:
- Creating test helper functions
- Using `t.Helper()` for correct error reporting
- Setup and teardown patterns
- Using `t.Cleanup()` for resource management

---

## Functions to Test

The `user.go` file provides:

```go
type User struct { ... }
func NewUser(name, email string) (*User, error)
func (u *User) Validate() error
func (u *User) UpdateEmail(email string) error
```

---

## Test Helpers

### Without Helper Functions (Repetitive)

```go
func TestNewUser(t *testing.T) {
    user, err := NewUser("Alice", "alice@example.com")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if user.Name != "Alice" {
        t.Errorf("Name = %q, want %q", user.Name, "Alice")
    }

    user2, err := NewUser("Bob", "bob@example.com")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if user2.Name != "Bob" {
        t.Errorf("Name = %q, want %q", user2.Name, "Bob")
    }
    // Lots of repetition!
}
```

### With Helper Functions (Clean)

```go
func assertNoError(t *testing.T, err error) {
    t.Helper()  // Mark this as a helper
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
}

func assertEqual(t *testing.T, got, want interface{}) {
    t.Helper()
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}

func TestNewUser(t *testing.T) {
    user, err := NewUser("Alice", "alice@example.com")
    assertNoError(t, err)
    assertEqual(t, user.Name, "Alice")

    user2, err := NewUser("Bob", "bob@example.com")
    assertNoError(t, err)
    assertEqual(t, user2.Name, "Bob")
    // Much cleaner!
}
```

---

## The t.Helper() Method

**Why use t.Helper()?**

```go
// Without t.Helper()
func checkError(t *testing.T, err error) {
    if err != nil {
        t.Fatal("error occurred")  // Line 3
    }
}

func TestSomething(t *testing.T) {
    err := SomeFunction()
    checkError(t, err)  // Line 9
}

// Output: helper.go:3: error occurred
// ❌ Points to helper function, not the actual test!
```

```go
// With t.Helper()
func checkError(t *testing.T, err error) {
    t.Helper()  // Mark as helper!
    if err != nil {
        t.Fatal("error occurred")
    }
}

func TestSomething(t *testing.T) {
    err := SomeFunction()
    checkError(t, err)  // Line 9
}

// Output: helper_test.go:9: error occurred
// ✅ Points to the test that called the helper!
```

**Rule:** Always call `t.Helper()` at the start of test helper functions.

---

## Setup and Teardown

### Using t.Cleanup()

```go
func TestWithCleanup(t *testing.T) {
    // Setup
    file, err := os.CreateTemp("", "test")
    if err != nil {
        t.Fatal(err)
    }

    // Register cleanup (runs even if test fails)
    t.Cleanup(func() {
        file.Close()
        os.Remove(file.Name())
    })

    // Test logic
    // ...

    // Cleanup happens automatically
}
```

### Helper Function for Setup

```go
func setupTestUser(t *testing.T) *User {
    t.Helper()

    user, err := NewUser("Test User", "test@example.com")
    if err != nil {
        t.Fatalf("setup failed: %v", err)
    }

    return user
}

func TestUserUpdate(t *testing.T) {
    user := setupTestUser(t)  // Clean setup

    err := user.UpdateEmail("new@example.com")
    if err != nil {
        t.Fatal(err)
    }
    // ...
}
```

---

## Your Task

Create test helpers and use them to test the User type. Build:

1. **Helper Functions:**
   - `assertNoError(t, err)` - Fail if error is not nil
   - `assertError(t, err)` - Fail if error is nil
   - `assertEqual(t, got, want)` - Compare any two values
   - `setupUser(t, name, email)` - Create a test user

2. **Tests Using Helpers:**
   - TestNewUser - test valid and invalid user creation
   - TestUserValidate - test validation logic
   - TestUserUpdateEmail - test email update

---

## Common Helper Patterns

### Error Checking Helpers

```go
func assertNoError(t *testing.T, err error) {
    t.Helper()
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
}

func assertError(t *testing.T, err error, msgContains string) {
    t.Helper()
    if err == nil {
        t.Fatal("expected error, got nil")
    }
    if !strings.Contains(err.Error(), msgContains) {
        t.Errorf("error %q does not contain %q", err.Error(), msgContains)
    }
}
```

### Comparison Helpers

```go
func assertEqual(t *testing.T, got, want interface{}) {
    t.Helper()
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}

func assertStringContains(t *testing.T, s, substr string) {
    t.Helper()
    if !strings.Contains(s, substr) {
        t.Errorf("string %q does not contain %q", s, substr)
    }
}
```

### Setup Helpers

```go
func setupUser(t *testing.T, name, email string) *User {
    t.Helper()
    user, err := NewUser(name, email)
    assertNoError(t, err)
    return user
}
```

---

## Instructions

1. Open `user_test.go`
2. Create the helper functions at the top of the file
3. Use the helpers in your tests (follow TODO markers)
4. Run tests and observe how error messages point to test lines, not helper lines
5. Intentionally break a test to see t.Helper() in action

**Experiment:**
- Remove `t.Helper()` from a helper and break a test - see where error points
- Add `t.Helper()` back and see the difference

---

## Hints

<details>
<summary>Hint 1: Basic Helper Structure</summary>

```go
func helperName(t *testing.T, params...) returnType {
    t.Helper()  // ALWAYS first line

    // Helper logic

    // Use t.Fatal/Fatalf for critical failures
    // Use t.Error/Errorf for non-critical failures

    return result
}
```
</details>

<details>
<summary>Hint 2: When to Use Helpers</summary>

Create helpers when you:
- Check errors repeatedly
- Make the same comparison often
- Set up test data
- Need cleanup logic
- Want more readable tests

Don't create helpers for:
- One-off checks
- Overly generic operations
- Things that are already clear
</details>

<details>
<summary>Hint 3: Setup Helper Pattern</summary>

```go
func setupUser(t *testing.T, name, email string) *User {
    t.Helper()

    user, err := NewUser(name, email)
    if err != nil {
        t.Fatalf("setup failed: %v", err)
    }

    return user
}

// Usage
func TestSomething(t *testing.T) {
    user := setupUser(t, "Alice", "alice@example.com")
    // Test with user...
}
```
</details>

<details>
<summary>Full Solution</summary>

```go
package user

import (
    "strings"
    "testing"
)

// Helper functions

func assertNoError(t *testing.T, err error) {
    t.Helper()
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
}

func assertError(t *testing.T, err error) {
    t.Helper()
    if err == nil {
        t.Fatal("expected error, got nil")
    }
}

func assertErrorContains(t *testing.T, err error, substr string) {
    t.Helper()
    if err == nil {
        t.Fatal("expected error, got nil")
    }
    if !strings.Contains(err.Error(), substr) {
        t.Errorf("error %q does not contain %q", err.Error(), substr)
    }
}

func assertEqual(t *testing.T, got, want interface{}) {
    t.Helper()
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}

func setupUser(t *testing.T, name, email string) *User {
    t.Helper()
    user, err := NewUser(name, email)
    assertNoError(t, err)
    return user
}

// Tests

func TestNewUser(t *testing.T) {
    t.Run("valid user", func(t *testing.T) {
        user, err := NewUser("Alice", "alice@example.com")
        assertNoError(t, err)
        assertEqual(t, user.Name, "Alice")
        assertEqual(t, user.Email, "alice@example.com")
    })

    t.Run("empty name", func(t *testing.T) {
        _, err := NewUser("", "alice@example.com")
        assertErrorContains(t, err, "name")
    })

    t.Run("invalid email", func(t *testing.T) {
        _, err := NewUser("Alice", "invalid")
        assertErrorContains(t, err, "email")
    })
}

func TestUserValidate(t *testing.T) {
    t.Run("valid user", func(t *testing.T) {
        user := setupUser(t, "Alice", "alice@example.com")
        err := user.Validate()
        assertNoError(t, err)
    })
}

func TestUserUpdateEmail(t *testing.T) {
    t.Run("valid email", func(t *testing.T) {
        user := setupUser(t, "Alice", "alice@example.com")
        err := user.UpdateEmail("newalice@example.com")
        assertNoError(t, err)
        assertEqual(t, user.Email, "newalice@example.com")
    })

    t.Run("invalid email", func(t *testing.T) {
        user := setupUser(t, "Alice", "alice@example.com")
        err := user.UpdateEmail("invalid")
        assertError(t, err)
        assertEqual(t, user.Email, "alice@example.com") // Unchanged
    })
}
```
</details>

---

## What This Teaches

- **Test helpers** - Extract common test code
- **t.Helper()** - Correct error line reporting
- **Setup/teardown** - Organizing test dependencies
- **t.Cleanup()** - Automatic resource cleanup
- **Readable tests** - Helpers make tests clearer

---

## Think About

1. Why is t.Helper() necessary? What problem does it solve?
2. When would you use t.Cleanup() instead of defer?
3. Should test helpers ever return errors, or always call t.Fatal()?
4. How many levels of helper functions is too many?

---

**Next Exercise:** `05_testing_errors` - Comprehensive error testing patterns
