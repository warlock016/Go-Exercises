# Exercise 07: Validation Errors

## 🎯 Learning Goal
Learn to aggregate multiple validation errors into a single error type, allowing you to report all validation failures at once instead of stopping at the first error.

## 📝 Problem Description

When validating forms or user input, you often want to report ALL validation errors, not just the first one. This requires collecting errors and returning them together.

## 🔧 Type and Function Signatures

```go
// ValidationErrors holds multiple field validation errors
type ValidationErrors struct {
    Errors []FieldError
}

type FieldError struct {
    Field   string
    Message string
}

// Implement Error() interface
func (ve *ValidationErrors) Error() string

// Add adds a field error to the collection
func (ve *ValidationErrors) Add(field, message string)

// HasErrors returns true if there are any errors
func (ve *ValidationErrors) HasErrors() bool

// ValidateUser validates all user fields and returns all errors
func ValidateUser(name, email string, age int) error

// ValidateProduct validates product fields
func ValidateProduct(name string, price float64, quantity int) error
```

## 💡 Examples

```go
err := ValidateUser("", "bademail", 200)
if verrs, ok := err.(*ValidationErrors); ok {
    for _, ferr := range verrs.Errors {
        fmt.Printf("%s: %s\n", ferr.Field, ferr.Message)
    }
}
// Output:
// name: cannot be empty
// email: must contain @
// age: must be between 0 and 150
```

## 🎓 What This Teaches

- **Error aggregation** - Collecting multiple errors
- **User-friendly validation** - Reporting all issues at once
- **Custom error collections** - Building complex error structures

---

**Next Exercise:** `08_http_error_responses` - Mapping errors to HTTP status codes
