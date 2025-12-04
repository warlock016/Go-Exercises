# Exercise 03: Sentinel Errors

## 🎯 Learning Goal
Master the sentinel error pattern: package-level error variables that represent specific, well-known error conditions, and learn to check them using `errors.Is`.

## 📝 Problem Description

**Sentinel errors** are predeclared package-level error variables that represent specific conditions. They allow callers to distinguish between different error types without parsing error strings.

Examples from the standard library:
- `io.EOF` - End of file reached
- `sql.ErrNoRows` - Database query returned no rows
- `os.ErrNotExist` - File or directory doesn't exist

This pattern enables robust error handling where callers can make decisions based on error types, not brittle string matching.

## 🔧 Package Variables and Functions

Define these sentinel errors at package level:

```go
var (
    ErrNotFound      = errors.New("resource not found")
    ErrAlreadyExists = errors.New("resource already exists")
    ErrUnauthorized  = errors.New("unauthorized access")
    ErrInvalidInput  = errors.New("invalid input")
    ErrTimeout       = errors.New("operation timed out")
)
```

Implement these functions in `sentinel_errors.go`:

```go
// FindUser returns ErrNotFound if user doesn't exist
func FindUser(userID int, users map[int]string) (string, error)

// CreateUser returns ErrAlreadyExists if user exists, ErrInvalidInput if name is empty
func CreateUser(userID int, name string, users map[int]string) error

// DeleteUser returns ErrNotFound if user doesn't exist, ErrUnauthorized if userID is 0
func DeleteUser(userID int, users map[int]string) error

// IsNotFoundError checks if error is ErrNotFound using errors.Is
func IsNotFoundError(err error) bool

// IsTimeoutError checks if error is ErrTimeout
func IsTimeoutError(err error) bool

// HandleError demonstrates checking multiple sentinel errors and taking action
// Returns a user-friendly message based on the error type
func HandleError(err error) string
```

## 💡 Examples

```go
users := map[int]string{
    1: "Alice",
    2: "Bob",
}

// Find existing user
name, err := FindUser(1, users)
// name = "Alice", err = nil

// Find non-existent user
name, err = FindUser(99, users)
// name = "", err = ErrNotFound

// Check with errors.Is
if errors.Is(err, ErrNotFound) {
    fmt.Println("User not found")
}

// Create duplicate user
err = CreateUser(1, "Charlie", users)
// err = ErrAlreadyExists

// Create with invalid input
err = CreateUser(3, "", users)
// err = ErrInvalidInput

// Delete non-existent user
err = DeleteUser(99, users)
// err = ErrNotFound

// Unauthorized deletion
err = DeleteUser(0, users)
// err = ErrUnauthorized

// Error checking helpers
if IsNotFoundError(err) {
    // Handle not found case
}

// Handle error with friendly messages
msg := HandleError(ErrNotFound)
// msg = "The requested resource was not found"

msg = HandleError(ErrTimeout)
// msg = "The operation timed out, please try again"
```

## 📋 Instructions

1. **Define sentinel errors** at package level using `var` block
2. **FindUser:** Return ErrNotFound if userID not in map
3. **CreateUser:** Check if userID exists (return ErrAlreadyExists), check if name is empty (return ErrInvalidInput)
4. **DeleteUser:** Check if userID is 0 (return ErrUnauthorized), check if exists (return ErrNotFound)
5. **IsNotFoundError:** Use `errors.Is(err, ErrNotFound)`
6. **IsTimeoutError:** Use `errors.Is(err, ErrTimeout)`
7. **HandleError:** Check error type with errors.Is and return appropriate message for each sentinel error

## 🧪 Testing

Run tests with:
```bash
go test -v
```

Expected: ~25-30 test cases checking sentinel error returns and comparisons

## 🤔 Think About

1. **Why use sentinel errors?**
   - Allows programmatic error checking without string parsing
   - Creates a well-defined error API
   - Enables error handling decisions in calling code

2. **When to export sentinel errors?**
   - Export (ErrNotFound) when callers need to check the error
   - Keep private when errors are internal only

3. **Why errors.Is instead of ==?**
   - errors.Is works with wrapped errors
   - == only checks direct equality
   - errors.Is is more robust for error chains

4. **Sentinel errors vs custom types?**
   - Sentinel errors: Simple, well-known conditions
   - Custom types: When you need additional data/context

## 💡 Hints

<details>
<summary>Hint 1: Defining sentinel errors</summary>

Define at package level (outside functions):
```go
package sentinel_errors

import "errors"

var (
    ErrNotFound      = errors.New("resource not found")
    ErrAlreadyExists = errors.New("resource already exists")
    ErrUnauthorized  = errors.New("unauthorized access")
    ErrInvalidInput  = errors.New("invalid input")
    ErrTimeout       = errors.New("operation timed out")
)
```

Place this at the top of your file, after package declaration and imports.
</details>

<details>
<summary>Hint 2: Returning sentinel errors</summary>

Return the package variable directly:
```go
func FindUser(userID int, users map[int]string) (string, error) {
    name, exists := users[userID]
    if !exists {
        return "", ErrNotFound  // Return the sentinel
    }
    return name, nil
}
```
</details>

<details>
<summary>Hint 3: Using errors.Is</summary>

Import "errors" and use errors.Is:
```go
func IsNotFoundError(err error) bool {
    return errors.Is(err, ErrNotFound)
}
```

This works even if err is wrapped with additional context.
</details>

<details>
<summary>Hint 4: Multiple condition checking</summary>

Check multiple conditions, return appropriate sentinel:
```go
func CreateUser(userID int, name string, users map[int]string) error {
    if name == "" {
        return ErrInvalidInput
    }
    if _, exists := users[userID]; exists {
        return ErrAlreadyExists
    }
    users[userID] = name
    return nil
}
```
</details>

<details>
<summary>Hint 5: Handling multiple error types</summary>

Use errors.Is in if-else chain:
```go
func HandleError(err error) string {
    if errors.Is(err, ErrNotFound) {
        return "The requested resource was not found"
    }
    if errors.Is(err, ErrAlreadyExists) {
        return "That resource already exists"
    }
    // ... check other errors
    return "An unknown error occurred"
}
```
</details>

<details>
<summary>Full Solution</summary>

```go
package sentinel_errors

import "errors"

var (
	ErrNotFound      = errors.New("resource not found")
	ErrAlreadyExists = errors.New("resource already exists")
	ErrUnauthorized  = errors.New("unauthorized access")
	ErrInvalidInput  = errors.New("invalid input")
	ErrTimeout       = errors.New("operation timed out")
)

func FindUser(userID int, users map[int]string) (string, error) {
	name, exists := users[userID]
	if !exists {
		return "", ErrNotFound
	}
	return name, nil
}

func CreateUser(userID int, name string, users map[int]string) error {
	if name == "" {
		return ErrInvalidInput
	}
	if _, exists := users[userID]; exists {
		return ErrAlreadyExists
	}
	users[userID] = name
	return nil
}

func DeleteUser(userID int, users map[int]string) error {
	if userID == 0 {
		return ErrUnauthorized
	}
	if _, exists := users[userID]; !exists {
		return ErrNotFound
	}
	delete(users, userID)
	return nil
}

func IsNotFoundError(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func IsTimeoutError(err error) bool {
	return errors.Is(err, ErrTimeout)
}

func HandleError(err error) string {
	if errors.Is(err, ErrNotFound) {
		return "The requested resource was not found"
	}
	if errors.Is(err, ErrAlreadyExists) {
		return "That resource already exists"
	}
	if errors.Is(err, ErrUnauthorized) {
		return "You are not authorized to perform this action"
	}
	if errors.Is(err, ErrInvalidInput) {
		return "The input provided is invalid"
	}
	if errors.Is(err, ErrTimeout) {
		return "The operation timed out, please try again"
	}
	return "An unknown error occurred"
}
```
</details>

## 🎓 What This Teaches

- **Sentinel error pattern** - Package-level error variables for well-known conditions
- **errors.Is** - Robust error checking that works with wrapped errors
- **Error API design** - Creating a consistent, programmatic error interface
- **Exported vs unexported** - When to make errors public
- **Error discrimination** - Allowing callers to make decisions based on error type
- **Standard library patterns** - Following conventions from io.EOF, sql.ErrNoRows, etc.
- **Map operations** - Using existence checking with map lookups

---

**Next Exercise:** `04_custom_error_types` - Implementing the error interface on structs
