# Module 07: Packages and Modules

**Prerequisites:** Completed modules 01-06
**Estimated Time:** 4-6 hours
**Exercises:** 8 progressive exercises

## 🎯 Learning Objectives

By the end of this module, you will:
- Understand Go's package visibility rules (exported vs unexported)
- Organize code across multiple files and packages
- Use import paths and aliases correctly
- Master package initialization with init()
- Design clean, maintainable package APIs
- Avoid import cycles through proper architecture
- Use internal/ packages for encapsulation
- Structure modules for real-world applications

## 📚 Key Concepts

### Package Visibility

Go has simple visibility rules:
- **Exported:** Identifiers starting with uppercase letter (e.g., `User`, `NewLogger`)
- **Unexported:** Identifiers starting with lowercase letter (e.g., `config`, `validate`)

### Package Organization

```go
// Multiple files can share the same package
// calculator/add.go
package calculator

func Add(a, b int) int { return a + b }

// calculator/multiply.go
package calculator  // Same package!

func Multiply(a, b int) int { return a * b }
```

### Import Paths

```go
import (
    "fmt"                    // Standard library
    "encoding/json"          // Standard library nested

    "github.com/user/repo"   // External module

    myapp "mycompany.com/app" // Aliasing
    _ "github.com/lib/pq"     // Blank import (side effects only)
)
```

### Package Initialization

```go
var registry = make(map[string]Handler)

func init() {
    // Runs automatically before main()
    // Use for registration, validation, setup
    register("default", defaultHandler)
}
```

## 📋 Module Structure

### Tier 1: Introduction (Exercises 01-03)
Understanding the basics of package organization and visibility.

### Tier 2: Application (Exercises 04-06)
Applying package concepts to real-world patterns and avoiding common pitfalls.

### Tier 3: Integration (Exercises 07-08)
Advanced package architecture and module design.

## 🎓 Exercise List

| # | Name | Concept | Difficulty | Est. Time |
|---|------|---------|------------|-----------|
| 01 | Visibility Basics | Exported vs unexported | Easy | 20 min |
| 02 | Package Organization | Multi-file packages | Easy | 25 min |
| 03 | Import Paths | Import conventions | Easy | 20 min |
| 04 | Init Functions | Package initialization | Medium | 30 min |
| 05 | API Design | Clean package APIs | Medium | 35 min |
| 06 | Avoiding Cycles | Cycle prevention | Medium | 35 min |
| 07 | Internal Packages | internal/ visibility | Medium | 30 min |
| 08 | Module Structure | Architecture design | Hard | 40 min |

## 🚀 How to Use This Module

### Working Through Exercises

1. **Read each exercise README carefully**
2. **Understand the package structure** - some exercises have multiple packages
3. **Pay attention to visibility rules** - exported vs unexported is crucial
4. **Run tests frequently** to verify your understanding
5. **Write EXPLANATION.md** explaining your design decisions

### Testing Commands

```bash
# Run all module tests
cd 07-packages-modules
go test ./...

# Run specific exercise
cd 01_visibility_basics
go test -v

# Check if code compiles (important for visibility exercises)
go build ./...
```

### Success Criteria

To complete this module:
- [ ] Complete all 8 exercises (tests passing)
- [ ] Understand exported vs unexported identifiers
- [ ] Can organize code across multiple files/packages
- [ ] Know how to avoid import cycles
- [ ] Can design clean package APIs with constructors
- [ ] Understand init() function execution order
- [ ] Can explain when to use internal/ packages

## 💡 Key Patterns You'll Learn

### Constructor Pattern
```go
// Exported type, unexported fields
type Config struct {
    host string
    port int
}

// Constructor function (convention: New + TypeName)
func NewConfig(host string, port int) *Config {
    return &Config{host: host, port: port}
}
```

### Functional Options Pattern
```go
type Option func(*Logger)

func WithLevel(level string) Option {
    return func(l *Logger) {
        l.level = level
    }
}

func NewLogger(opts ...Option) *Logger {
    l := &Logger{level: "info"}
    for _, opt := range opts {
        opt(l)
    }
    return l
}
```

### Registry Pattern
```go
var handlers = make(map[string]Handler)

func Register(name string, h Handler) {
    handlers[name] = h
}

func Get(name string) (Handler, bool) {
    h, ok := handlers[name]
    return h, ok
}
```

### Avoiding Import Cycles

**Problem:**
```
package A imports package B
package B imports package A
→ ERROR: import cycle
```

**Solutions:**
1. Extract shared types to a third package
2. Use interfaces to reverse dependencies
3. Merge packages if they're too coupled

## 🧠 Think About

1. Why does Go use capitalization for visibility instead of keywords like `public`/`private`?
2. When should you split code into multiple packages vs keeping it in one?
3. How do init() functions help with plugin architectures?
4. Why are import cycles not allowed in Go?
5. When should you use the internal/ directory pattern?

## 📖 Recommended Reading

1. **Effective Go - Packages:** https://go.dev/doc/effective_go#names
2. **Go Blog - Package Names:** https://go.dev/blog/package-names
3. **Go Modules Reference:** https://go.dev/ref/mod
4. **Standard Library Structure:** Look at how `encoding/json`, `net/http` are organized

## ⚡ Quick Reference

### Visibility Rules
- `User` → Exported (visible outside package)
- `user` → Unexported (package-private)
- `NewUser()` → Exported function (constructor)
- `validate()` → Unexported function (helper)

### Import Rules
- Package name = last element of import path (usually)
- `import "fmt"` → use as `fmt.Println()`
- `import f "fmt"` → use as `f.Println()`
- `import _ "pkg"` → run init() only, don't use

### Init Function
- Runs automatically before main()
- Can have multiple init() per package
- Executes in declaration order
- Use for: setup, registration, validation

## 🎯 After This Module

You will be able to:
- Structure Go projects with multiple packages
- Design APIs that are easy to use and hard to misuse
- Understand package boundaries and dependencies
- Write maintainable, well-organized Go code
- Navigate and understand real-world Go projects

---

**Ready to master Go packages and modules?** Let's build well-structured Go applications!

Start with **Exercise 01: Visibility Basics** 🚀
