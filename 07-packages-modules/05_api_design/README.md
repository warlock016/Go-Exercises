# Exercise 05: API Design

**Concept:** Clean package APIs with functional options pattern
**Difficulty:** Medium
**Estimated Time:** 35 minutes

## 🎯 Learning Goal

Learn to design clean, flexible package APIs using the functional options pattern and proper encapsulation.

## The Problem

When designing a package API, you want:
- **Encapsulation:** Hide implementation details
- **Flexibility:** Easy to add new options later
- **Usability:** Clear, hard to misuse
- **Backward compatibility:** Adding options doesn't break existing code

The **functional options pattern** solves this elegantly.

## Your Task

Create a `logger` sub-package with a clean API:

### File Structure
```
05_api_design/
├── README.md
├── api_design.go
├── api_design_test.go
└── logger/
    └── logger.go
```

### logger Package API

1. **Logger type** (exported) with unexported fields:
   - `level string`
   - `prefix string`

2. **Option type** for configuration:
   - `type Option func(*Logger)`

3. **Constructor with options:**
   - `NewLogger(opts ...Option) *Logger`

4. **Option functions:**
   - `WithLevel(level string) Option`
   - `WithPrefix(prefix string) Option`

5. **Methods:**
   - `Log(message string) string` - returns formatted log message
   - `GetLevel() string` - exported getter

## Function Signatures

```go
// logger/logger.go
package logger

type Logger struct {
    // Unexported fields
}

type Option func(*Logger)

func NewLogger(opts ...Option) *Logger

func WithLevel(level string) Option

func WithPrefix(prefix string) Option

func (l *Logger) Log(message string) string

func (l *Logger) GetLevel() string
```

## Examples

```go
import "path/to/logger"

// Default logger
log := logger.NewLogger()
log.Log("Hello")  // "[INFO] Hello"

// With custom level
log = logger.NewLogger(logger.WithLevel("DEBUG"))
log.Log("Hello")  // "[DEBUG] Hello"

// With multiple options
log = logger.NewLogger(
    logger.WithLevel("ERROR"),
    logger.WithPrefix("myapp"),
)
log.Log("Failed")  // "myapp [ERROR] Failed"

// Options can be applied in any order
log = logger.NewLogger(
    logger.WithPrefix("api"),
    logger.WithLevel("WARN"),
)

// Easy to add new options later without breaking existing code
log = logger.NewLogger(
    logger.WithLevel("INFO"),
    // future: logger.WithOutput(os.Stderr),
    // future: logger.WithTimestamp(true),
)
```

## Instructions

1. Create `logger/logger.go`
2. Define `Logger` type with unexported `level` and `prefix` fields
3. Define `Option` type as `func(*Logger)`
4. Implement `NewLogger(opts ...Option) *Logger` with defaults
5. Implement `WithLevel(level string) Option`
6. Implement `WithPrefix(prefix string) Option`
7. Implement `Log(message string) string` method
8. Implement `GetLevel() string` getter
9. Update `api_design.go` to use the logger
10. Run `go test -v`

## Hints

### Basic - Functional Options Pattern
```go
type Option func(*Logger)

func WithLevel(level string) Option {
    return func(l *Logger) {
        l.level = level
    }
}

// Used as:
opt := WithLevel("DEBUG")
opt(myLogger)  // Modifies myLogger.level
```

### Intermediate - Constructor with Options
```go
func NewLogger(opts ...Option) *Logger {
    // 1. Create logger with defaults
    l := &Logger{
        level:  "INFO",
        prefix: "",
    }

    // 2. Apply all options
    for _, opt := range opts {
        opt(l)
    }

    // 3. Return configured logger
    return l
}
```

### Advanced - Complete Implementation
```go
package logger

import "fmt"

type Logger struct {
    level  string
    prefix string
}

type Option func(*Logger)

func NewLogger(opts ...Option) *Logger {
    l := &Logger{
        level:  "INFO",
        prefix: "",
    }

    for _, opt := range opts {
        opt(l)
    }

    return l
}

func WithLevel(level string) Option {
    return func(l *Logger) {
        l.level = level
    }
}

func WithPrefix(prefix string) Option {
    return func(l *Logger) {
        l.prefix = prefix
    }
}

func (l *Logger) Log(message string) string {
    if l.prefix != "" {
        return fmt.Sprintf("%s [%s] %s", l.prefix, l.level, message)
    }
    return fmt.Sprintf("[%s] %s", l.level, message)
}

func (l *Logger) GetLevel() string {
    return l.level
}
```

## 🧠 Think About

1. Why use functional options instead of a config struct?
2. How does this pattern support backward compatibility?
3. Why are Logger fields unexported?
4. What if you want to add a new option (like WithOutput) later?
5. How would you make some options mutually exclusive?

## What This Teaches

- Functional options pattern for flexible APIs
- Encapsulation with unexported fields
- Variadic functions (`opts ...Option`)
- Designing APIs that are hard to misuse
- Backward compatibility in API design
- Constructor pattern with configuration

## Benefits of This Pattern

1. **Backward compatible:** Adding options doesn't break existing code
2. **Optional parameters:** Caller chooses what to configure
3. **Clear defaults:** Constructor sets sensible defaults
4. **Composable:** Options can be created and passed around
5. **Type-safe:** Compiler enforces correct usage

## Real-World Examples

This pattern is used extensively in Go:
- `grpc.Dial(target, opts ...DialOption)`
- `http.Server` configuration
- Database connection pools
- Many third-party libraries

## After Completing

Write in your `EXPLANATION.md`:
- How the functional options pattern works
- Why it's better than a config struct with many fields
- The role of the Option type
- How this supports adding features without breaking changes
- A diagram showing the function call chain
