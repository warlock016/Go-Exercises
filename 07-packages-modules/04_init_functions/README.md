# Exercise 04: Init Functions

**Concept:** Package initialization with init()
**Difficulty:** Medium
**Estimated Time:** 30 minutes

## 🎯 Learning Goal

Understand how init() functions work for package initialization, registration patterns, and execution order.

## The Problem

Go's `init()` function is special:
- Runs automatically before `main()`
- Executed once per package
- Can have multiple `init()` in a package (run in declaration order)
- Cannot be called manually
- Used for: setup, validation, registration

Common use cases:
- Registering handlers/plugins
- Validating configuration
- Initializing package-level variables
- Database driver registration

## Your Task

Create a handler registry system using init() functions:

1. Package-level `handlers` map (unexported)
2. `Register()` function to add handlers
3. `GetHandler()` function to retrieve handlers
4. `init()` function to register default handlers
5. `IsInitialized()` function to check if initialization completed

## Function Signatures

```go
type Handler func(string) string

var handlers map[string]Handler  // Package-level, initialized in init()

func init()

func Register(name string, handler Handler)

func GetHandler(name string) (Handler, bool)

func IsInitialized() bool
```

## Examples

```go
// Package automatically initializes handlers
import "path/to/initfuncs"

// Default handlers already registered via init()
handler, ok := initfuncs.GetHandler("uppercase")
if ok {
    result := handler("hello")  // "HELLO"
}

// Register custom handler
initfuncs.Register("reverse", func(s string) string {
    // Reverse logic
    return reversed
})

handler, ok = initfuncs.GetHandler("reverse")
if ok {
    result := handler("hello")  // "olleh"
}

// Check if initialization happened
if initfuncs.IsInitialized() {
    // Safe to use
}
```

## Instructions

1. Open `init_functions.go`
2. Define the `Handler` type and `handlers` map
3. Implement `init()` to initialize the map and register defaults
4. Implement `Register()` to add new handlers
5. Implement `GetHandler()` to retrieve handlers
6. Implement `IsInitialized()` to verify initialization
7. Register at least 2 default handlers in init()
8. Run `go test -v`

## Hints

### Basic - Init Function Basics
```go
var initialized bool

func init() {
    // Runs automatically before main()
    // Cannot be called manually
    // Perfect for setup and registration
    initialized = true
}
```

### Intermediate - Registry Pattern
```go
type Handler func(string) string

var handlers = make(map[string]Handler)
var initialized bool

func init() {
    // Register default handlers
    Register("uppercase", func(s string) string {
        return strings.ToUpper(s)
    })

    Register("lowercase", func(s string) string {
        return strings.ToLower(s)
    })

    initialized = true
}

func Register(name string, handler Handler) {
    handlers[name] = handler
}

func GetHandler(name string) (Handler, bool) {
    h, ok := handlers[name]
    return h, ok
}

func IsInitialized() bool {
    return initialized
}
```

### Advanced - Multiple Init Functions
```go
// You can have multiple init() functions in the same package
// They execute in declaration order

var config Config

func init() {
    // First init: Load config
    config = loadConfig()
}

func init() {
    // Second init: Validate config (runs after first)
    if err := validateConfig(config); err != nil {
        panic(err)  // Fatal if config invalid
    }
}

func init() {
    // Third init: Register handlers based on config
    registerHandlers(config)
}
```

### Init Execution Order
```go
// Package A
var x = initX()       // 1. Package-level variable initialization
func initX() int {    //    (in dependency order)
    return 1
}
func init() {         // 2. init() functions
    // setup             (in declaration order)
}

// Package B imports A
import "A"            // A's init runs first
func init() {         // Then B's init runs
    // B's setup
}
```

## 🧠 Think About

1. Why can't you call init() manually?
2. What happens if multiple packages all have init() functions?
3. When would you use init() vs a constructor function?
4. What's the danger of doing too much in init()?
5. How do database drivers use init() for registration?

## What This Teaches

- init() function purpose and behavior
- Package initialization order
- Registry pattern implementation
- When to use init() vs explicit initialization
- Package-level variable initialization
- Side effects and automatic setup

## Real-World Example

Database drivers use this pattern:

```go
// In your application
import _ "github.com/lib/pq"  // PostgreSQL driver

// In github.com/lib/pq package:
func init() {
    sql.Register("postgres", &Driver{})
}

// Now you can use:
db, err := sql.Open("postgres", connStr)
```

## After Completing

Write in your `EXPLANATION.md`:
- What init() does and when it runs
- Why the registry pattern uses init()
- Execution order of multiple init() functions
- When to use init() vs explicit initialization
- A real-world example where init() is useful
