# Exercise 03: Error Interface

**Concept:** Custom error types implementing the error interface
**Difficulty:** Easy-Medium
**Estimated Time:** 40 minutes

## Learning Goal

Understand that `error` is just an interface, and learn to create custom error types that provide more context than simple strings. This is essential for building robust, informative error handling in Go applications.

## The Problem

Go's built-in `errors.New()` creates simple error strings:

```go
err := errors.New("validation failed")
fmt.Println(err)  // → "validation failed"
```

But what if you need more information? Custom error types let you include additional context:

```go
err := &ValidationError{
    Field:   "username",
    Problem: "too short",
    MinLen:  3,
}
fmt.Println(err)  // → "validation error: username is too short (minimum 3 characters)"
```

## The error Interface

The error interface is incredibly simple:

```go
type error interface {
    Error() string
}
```

ANY type with an `Error() string` method automatically satisfies the error interface. This means you can create custom error types with whatever fields you need!

## Your Task

Create two custom error types and functions that return them.

### Part 1: ValidationError

Create a type for validation failures:

**Fields:**
- `Field string` - which field failed validation
- `Problem string` - what went wrong
- `MinLen int` - minimum required length (for length validation)

**Error() output:** "validation error: {Field} is {Problem} (minimum {MinLen} characters)"

**Function:** `Validate(username string) error`
- Returns nil if username is 3+ characters
- Returns ValidationError if too short

### Part 2: NetworkError

Create a type for network failures:

**Fields:**
- `URL string` - which URL failed
- `StatusCode int` - HTTP status code
- `Message string` - error description

**Error() output:** "network error: {Message} (URL: {URL}, status: {StatusCode})"

**Function:** `FetchData(url string, simulateFailure bool) (string, error)`
- If simulateFailure is true: return NetworkError
- If simulateFailure is false: return ("data from " + url, nil)

## Function Signatures

```go
type ValidationError struct {
    Field   string
    Problem string
    MinLen  int
}

type NetworkError struct {
    URL        string
    StatusCode int
    Message    string
}

func (e *ValidationError) Error() string
func (e *NetworkError) Error() string

func Validate(username string) error
func FetchData(url string, simulateFailure bool) (string, error)
```

## Examples

```go
// ValidationError
err := Validate("ab")
if err != nil {
    fmt.Println(err)
    // → "validation error: username is too short (minimum 3 characters)"
}

// NetworkError
data, err := FetchData("https://api.example.com", true)
if err != nil {
    fmt.Println(err)
    // → "network error: request failed (URL: https://api.example.com, status: 404)"
}

// Success cases
err = Validate("alice")  // → nil
data, err = FetchData("https://api.example.com", false)  // → ("data from https://api.example.com", nil)
```

## Instructions

1. Open `error_interface.go`
2. Define ValidationError and NetworkError structs
3. Implement Error() method for each (use pointer receiver)
4. Implement Validate function
5. Implement FetchData function
6. Run `go test -v`

## Hints

<details>
<summary>Hint 1: Error() Method</summary>

Use fmt.Sprintf to format the error message:

```go
func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error: %s is %s (minimum %d characters)",
        e.Field, e.Problem, e.MinLen)
}
```

Note: Use pointer receiver (*ValidationError) for Error() method.
</details>

<details>
<summary>Hint 2: Returning Custom Errors</summary>

Create and return the error type:

```go
func Validate(username string) error {
    if len(username) < 3 {
        return &ValidationError{
            Field:   "username",
            Problem: "too short",
            MinLen:  3,
        }
    }
    return nil
}
```
</details>

<details>
<summary>Hint 3: Pointer vs Value Receivers</summary>

For error types, use pointer receivers for Error():

```go
// GOOD - pointer receiver
func (e *ValidationError) Error() string { ... }

// Also works, but pointer is idiomatic for errors
func (e ValidationError) Error() string { ... }
```

Why? Errors are often compared by pointer, and you might want to add methods that modify the error later.
</details>

<details>
<summary>Hint 4: NetworkError Format</summary>

```go
func (e *NetworkError) Error() string {
    return fmt.Sprintf("network error: %s (URL: %s, status: %d)",
        e.Message, e.URL, e.StatusCode)
}
```
</details>

<details>
<summary>Hint 5: FetchData Implementation</summary>

```go
func FetchData(url string, simulateFailure bool) (string, error) {
    if simulateFailure {
        return "", &NetworkError{
            URL:        url,
            StatusCode: 404,
            Message:    "request failed",
        }
    }
    return "data from " + url, nil
}
```
</details>

## Think About

1. **Why use pointer receiver for Error()?**
   - Idiomatic for error types
   - Allows pointer comparison: `if err == specificErr`
   - Consistent with how errors are typically handled

2. **When should you create custom error types?**
   - When you need to programmatically inspect errors
   - When you want to provide structured error information
   - When callers might handle different errors differently

3. **How can you check for specific error types?**
   ```go
   err := Validate("ab")
   if verr, ok := err.(*ValidationError); ok {
       fmt.Println("Field:", verr.Field)
       fmt.Println("Problem:", verr.Problem)
   }
   ```

4. **What about errors.Is() and errors.As()?**
   - Go 1.13+ added these for better error handling
   - They work great with custom error types
   - You'll use them in later exercises!

## What This Teaches

- **error is just an interface** - any type with Error() string satisfies it
- **Custom error types** - adding context to errors
- **Pointer receivers** - why they're idiomatic for errors
- **Structured error information** - fields that can be accessed programmatically
- **Error formatting** - making errors informative and readable

## Common Mistakes to Avoid

1. **Returning nil struct instead of nil interface:**
   ```go
   // WRONG - returns (*ValidationError)(nil), not nil
   func Validate(s string) error {
       var err *ValidationError
       if len(s) < 3 {
           err = &ValidationError{...}
       }
       return err  // BUG: err != nil even when not set!
   }

   // RIGHT - explicitly return nil
   func Validate(s string) error {
       if len(s) < 3 {
           return &ValidationError{...}
       }
       return nil  // Explicitly return nil interface
   }
   ```

2. **Forgetting to use pointer when creating error:**
   ```go
   // WRONG - if Error() has pointer receiver
   return ValidationError{...}

   // RIGHT
   return &ValidationError{...}
   ```

3. **Not implementing Error() method:**
   ```go
   // Forgot Error() method - won't satisfy error interface!
   type MyError struct {
       Message string
   }
   // Missing: func (e *MyError) Error() string { ... }
   ```

4. **Inconsistent error messages:**
   ```go
   // BAD - different formats
   "Error: username invalid"
   "username is invalid"
   "Invalid username"

   // GOOD - consistent format
   "validation error: username is too short"
   "validation error: email is invalid"
   ```

## Challenge Extensions (Optional)

After completing the basic version, try these:

1. **Add type assertion in tests:**
   ```go
   err := Validate("ab")
   if verr, ok := err.(*ValidationError); ok {
       fmt.Println("Failed field:", verr.Field)
   }
   ```

2. **Create a FileError type:**
   ```go
   type FileError struct {
       Path      string
       Operation string  // "read", "write", "delete"
       Err       error   // Underlying error
   }
   ```

3. **Implement error wrapping:**
   ```go
   func (e *FileError) Unwrap() error {
       return e.Err
   }
   ```

## After Completing

You now understand:
- error is just an interface
- How to create custom error types
- Why pointer receivers are used for errors
- How to provide structured error information
- The foundation for advanced error handling

Custom error types are essential for production Go code. You'll see them everywhere!

---

**Next up:** Exercise 04 - Reader Basics (implementing io.Reader)
