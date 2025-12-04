# Exercise 05: Error Wrapping

## 🎯 Learning Goal
Master error wrapping with `fmt.Errorf` and the `%w` verb to build error chains that preserve context as errors propagate up the call stack.

## 📝 Problem Description

Error wrapping (introduced in Go 1.13) lets you add context to errors without losing the original error. The `%w` verb creates an error chain that can be inspected with `errors.Is` and `errors.Unwrap`.

```go
// Without wrapping - loses original error
return fmt.Errorf("failed to process user: %v", err)

// With wrapping - preserves original error
return fmt.Errorf("failed to process user: %w", err)
```

## 🔧 Function Signatures

```go
// OpenAndReadFile opens a file and reads it, wrapping errors at each step
func OpenAndReadFile(filename string) (string, error)

// ProcessUser validates and saves user, wrapping errors with context
func ProcessUser(name, email string) error

// FetchAndParse fetches URL and parses JSON, wrapping errors
func FetchAndParse(url string) (map[string]interface{}, error)

// UnwrapOnce unwraps one level of error wrapping
func UnwrapOnce(err error) error

// UnwrapAll repeatedly unwraps until reaching the root error
func UnwrapAll(err error) error
```

## 💡 Examples

```go
// Error wrapping preserves the chain
_, err := OpenAndReadFile("missing.txt")
// Error chain: "failed to read file: open missing.txt: no such file or directory"

if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File doesn't exist")  // Works because error is wrapped
}

// Multiple layers of wrapping
err = ProcessUser("", "test@example.com")
// "failed to process user: validation failed: name cannot be empty"

// Unwrap one level
unwrapped := UnwrapOnce(err)
// Returns: "validation failed: name cannot be empty"

// Unwrap to root
root := UnwrapAll(err)
// Returns: "name cannot be empty"
```

## 📋 Instructions

1. **OpenAndReadFile:** Use os.Open, wrap error with "failed to open %s". Use io.ReadAll, wrap error with "failed to read %s".
2. **ProcessUser:** Validate inputs, wrap validation errors with "validation failed". Call a save function, wrap errors with "failed to save user".
3. **FetchAndParse:** Simulate HTTP fetch and JSON parsing, wrap errors at each step
4. **UnwrapOnce:** Use errors.Unwrap to unwrap one level
5. **UnwrapAll:** Loop with errors.Unwrap until nil

## 🤔 Think About

1. **%w vs %v?** - %w creates wrappable errors, %v converts to string only
2. **When to wrap?** - Add context at each layer, but don't duplicate information
3. **Performance?** - Wrapping has minimal overhead

## 💡 Hints

<details>
<summary>Hint: Error wrapping with %w</summary>

```go
func OpenAndReadFile(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", fmt.Errorf("failed to open %s: %w", filename, err)
    }
    defer f.Close()

    data, err := io.ReadAll(f)
    if err != nil {
        return "", fmt.Errorf("failed to read %s: %w", filename, err)
    }
    return string(data), nil
}
```
</details>

## 🎓 What This Teaches

- **Error wrapping** - Using %w to preserve error chains
- **errors.Unwrap** - Extracting wrapped errors
- **Context preservation** - Adding information without losing the original
- **errors.Is compatibility** - Wrapped errors work with Is/As

---

**Next Exercise:** `06_error_inspection` - Using errors.Is and errors.As
