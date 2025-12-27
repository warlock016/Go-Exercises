# Exercise 07: Internal Packages

**Concept:** Using internal/ directory for package encapsulation
**Difficulty:** Medium
**Estimated Time:** 30 minutes

## 🎯 Learning Goal

Understand Go's special `internal/` directory and how it enforces package visibility at the module level.

## The Problem

Sometimes you want packages that are:
- Exported (uppercase names) for use within your module
- NOT accessible to external modules/packages
- More restricted than regular unexported identifiers

**Solution:** The `internal/` directory pattern.

## How internal/ Works

Go has a special rule for packages inside an `internal/` directory:

```
mymodule/
├── myapp/
│   ├── main.go                  ✓ Can import mymodule/myapp/internal/auth
│   └── internal/
│       └── auth/
│           └── auth.go
└── external/
    └── main.go                  ✗ CANNOT import mymodule/myapp/internal/auth
```

**Rule:** A package inside `internal/` can only be imported by:
- Packages in the same subtree (parent and siblings)
- NOT by packages outside the `internal/` ancestor

## Visibility Comparison

```go
// Regular unexported (lowercase)
type config struct {}  // Only visible in same package

// Regular exported (uppercase)
type Config struct {}  // Visible to all importers

// Internal package exported (uppercase in internal/)
// internal/auth/auth.go
type Token struct {}   // Visible to parent module only
```

## Your Task

This is a **conceptual exercise**. You'll implement functions that explain and validate internal package rules.

### Functions to Implement

1. `CanImport(importer, target string) bool` - Determines if import is allowed
2. `ExplainInternalRule() string` - Explains the internal/ rule
3. `ValidateStructure(paths []string) []string` - Finds invalid imports

## Function Signatures

```go
func CanImport(importer, target string) bool

func ExplainInternalRule() string

func ValidateStructure(paths []string) []string
```

## Examples

```go
// Can myapp import myapp/internal/auth?
CanImport("mymodule/myapp", "mymodule/myapp/internal/auth")
// true - myapp is parent of internal/

// Can external import myapp/internal/auth?
CanImport("mymodule/external", "mymodule/myapp/internal/auth")
// false - external is not in myapp/ subtree

// Can sibling import?
CanImport("mymodule/myapp/handler", "mymodule/myapp/internal/auth")
// true - handler is in myapp/ subtree

rule := ExplainInternalRule()
// Returns explanation of internal/ visibility

paths := []string{
    "mymodule/myapp imports mymodule/myapp/internal/auth",
    "mymodule/external imports mymodule/myapp/internal/auth",
}
invalid := ValidateStructure(paths)
// Returns: ["mymodule/external imports mymodule/myapp/internal/auth"]
```

## Instructions

1. Open `internal_packages.go`
2. Implement `CanImport()` - check if importer can access target
3. Implement `ExplainInternalRule()` - return clear explanation
4. Implement `ValidateStructure()` - find violations
5. Run `go test -v`

## Hints

### Basic - Understanding internal/
```go
// An internal package path looks like:
// "mymodule/myapp/internal/auth"
//                ^^^^^^^^
// The "internal" directory is special

// Can import if:
// 1. Target doesn't contain "internal" → always allowed
// 2. Target contains "internal" → only if importer is in parent tree
```

### Intermediate - Checking Ancestry
```go
import "strings"

func CanImport(importer, target string) bool {
    // If target doesn't have internal, always allowed
    if !strings.Contains(target, "/internal/") {
        return true
    }

    // Find the internal/ ancestor
    // "mymodule/myapp/internal/auth" → ancestor is "mymodule/myapp"
    parts := strings.Split(target, "/internal/")
    ancestor := parts[0]

    // Importer must be in ancestor's subtree
    return strings.HasPrefix(importer, ancestor)
}
```

### Advanced - Complete Implementation
```go
package internal

import "strings"

func CanImport(importer, target string) bool {
    if !strings.Contains(target, "/internal/") {
        return true
    }

    parts := strings.Split(target, "/internal/")
    ancestor := parts[0]

    return importer == ancestor || strings.HasPrefix(importer, ancestor+"/")
}

func ExplainInternalRule() string {
    return "Packages inside 'internal/' can only be imported by packages " +
           "in the same subtree. This enforces module-level encapsulation."
}

func ValidateStructure(paths []string) []string {
    var violations []string

    for _, path := range paths {
        // Parse "importer imports target"
        parts := strings.Split(path, " imports ")
        if len(parts) != 2 {
            continue
        }

        importer := strings.TrimSpace(parts[0])
        target := strings.TrimSpace(parts[1])

        if !CanImport(importer, target) {
            violations = append(violations, path)
        }
    }

    return violations
}
```

## 🧠 Think About

1. Why does Go have the internal/ pattern?
2. How is this different from unexported (lowercase) identifiers?
3. When would you use internal/ vs a regular package?
4. Can you have multiple internal/ directories in a project?
5. What happens if you try to import an internal package from outside?

## What This Teaches

- Go's internal/ directory special rule
- Module-level encapsulation
- When to use internal/ packages
- Package visibility beyond exported/unexported
- How to structure large projects with internal APIs

## Real-World Usage

### Standard Library
The Go standard library uses internal/ extensively:
```
net/http/internal/     ← Used by net/http, not by you
encoding/json/internal/
crypto/internal/
```

### Your Projects
```
myapp/
├── cmd/
│   └── server/
│       └── main.go           ← Can use internal/
├── internal/                 ← Private to myapp/
│   ├── auth/
│   ├── database/
│   └── config/
└── pkg/                      ← Public API
    └── client/
```

### Benefits
- **Encapsulation:** Keep implementation details private
- **Refactoring:** Change internal packages without breaking external users
- **API clarity:** Public vs private is explicit in structure
- **Versioning:** Internal changes don't require semver bumps

## After Completing

Write in your `EXPLANATION.md`:
- What the internal/ directory does
- How it differs from unexported identifiers
- The rule for who can import internal packages
- A real project structure using internal/
- When you would use internal/ in your own code
