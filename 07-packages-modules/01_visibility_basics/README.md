# Exercise 01: Visibility Basics

**Concept:** Understanding exported vs unexported identifiers
**Difficulty:** Easy
**Estimated Time:** 20 minutes

## 🎯 Learning Goal

Master Go's visibility rules: uppercase = exported (public), lowercase = unexported (private).

## The Problem

Go's visibility is simple but strict:
- **Exported (uppercase):** Can be accessed from other packages
- **Unexported (lowercase):** Only accessible within the same package

This applies to:
- Types: `User` vs `user`
- Functions: `NewUser()` vs `validate()`
- Struct fields: `Name` vs `password`
- Constants: `MaxSize` vs `defaultSize`

## Your Task

Create a user management system that demonstrates proper visibility:

1. **Exported type:** `User` struct with exported `Name` field and unexported `password` field
2. **Unexported type:** `config` struct (package-internal configuration)
3. **Exported constructor:** `NewUser()` to create users properly
4. **Unexported helper:** `validate()` to validate user data (internal use only)

## Function Signatures

```go
// Exported type - visible to other packages
type User struct {
    Name     string  // Exported field
    password string  // Unexported field (encapsulated)
}

// Unexported type - only visible in this package
type config struct {
    minPasswordLength int
    maxNameLength     int
}

// Exported constructor
func NewUser(name, password string) (*User, error)

// Unexported validation helper
func validate(name, password string) error

// Exported method to check password
func (u *User) CheckPassword(password string) bool
```

## Examples

```go
// In another package using this:
u, err := visibility.NewUser("Alice", "secret123")
if err != nil {
    log.Fatal(err)
}

fmt.Println(u.Name)           // OK - Name is exported
// fmt.Println(u.password)    // COMPILE ERROR - password unexported

ok := u.CheckPassword("secret123")  // OK - method is exported
fmt.Println(ok)               // true

// c := visibility.config{}   // COMPILE ERROR - config unexported
// visibility.validate(...)   // COMPILE ERROR - validate unexported
```

## Instructions

1. Open `visibility_basics.go`
2. Define the `User` and `config` types with proper visibility
3. Implement `NewUser()` constructor
4. Implement `validate()` helper (unexported)
5. Implement `CheckPassword()` method
6. Run `go test -v`
7. Read test failures carefully - they verify visibility rules

## Hints

### Basic - Understanding Visibility
- Uppercase first letter = exported (public)
- Lowercase first letter = unexported (private)
- This applies to: types, functions, methods, fields, constants

### Intermediate - Constructor Pattern
```go
func NewUser(name, password string) (*User, error) {
    // 1. Validate inputs using validate()
    // 2. Create and return User
    // 3. Return error if validation fails
}
```

### Advanced - Complete Implementation
```go
var cfg = config{
    minPasswordLength: 8,
    maxNameLength:     50,
}

func validate(name, password string) error {
    if len(name) == 0 {
        return errors.New("name cannot be empty")
    }
    if len(name) > cfg.maxNameLength {
        return errors.New("name too long")
    }
    if len(password) < cfg.minPasswordLength {
        return errors.New("password too short")
    }
    return nil
}

func NewUser(name, password string) (*User, error) {
    if err := validate(name, password); err != nil {
        return nil, err
    }
    return &User{
        Name:     name,
        password: password,
    }, nil
}

func (u *User) CheckPassword(password string) bool {
    return u.password == password
}
```

## 🧠 Think About

1. Why can't external packages access `u.password` directly?
2. What would happen if `validate()` was exported?
3. Why use a constructor (`NewUser`) instead of letting users create `User{}` directly?
4. How does this compare to `public`/`private` keywords in other languages?

## What This Teaches

- Go's simple but effective visibility system
- Encapsulation through unexported fields
- Constructor pattern for controlled initialization
- When to export vs hide implementation details
- Package-level vs type-level access control

## After Completing

Write in your `EXPLANATION.md`:
- The difference between exported and unexported
- Why `password` should be unexported
- The role of the `NewUser()` constructor
- How this prevents misuse of the User type
