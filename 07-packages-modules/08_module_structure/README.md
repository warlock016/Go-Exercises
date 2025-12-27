# Exercise 08: Module Structure

**Concept:** Designing well-structured Go modules and package architecture
**Difficulty:** Hard
**Estimated Time:** 40 minutes

## 🎯 Learning Goal

Learn to design maintainable Go module structures by understanding package organization principles, coupling, and cohesion.

## The Problem

As projects grow, poor structure leads to:
- **Import cycles** (circular dependencies)
- **High coupling** (packages depend on too many others)
- **Low cohesion** (unrelated code in the same package)
- **Unclear boundaries** (what belongs where?)
- **Hard to test** (can't isolate components)

Good structure requires:
- **Clear layers** (models → repository → service → handlers)
- **Low coupling** (few dependencies)
- **High cohesion** (related code together)
- **No cycles** (directed acyclic graph)
- **Testability** (easy to mock and test)

## Your Task

Implement functions that analyze and validate module structure:

1. **PackageInfo** - Represents a package with its dependencies
2. **ModuleDesign** - Represents a module's architecture
3. **ValidateDesign()** - Checks for structural problems
4. **DetectCycles()** - Finds import cycles
5. **CalculateCoupling()** - Measures package coupling
6. **SuggestStructure()** - Recommends package organization

## Function Signatures

```go
type PackageInfo struct {
    Name         string
    Imports      []string
    ExportsTypes bool
    IsInternal   bool
}

type ModuleDesign struct {
    Packages []PackageInfo
}

func ValidateDesign(design ModuleDesign) []string

func DetectCycles(design ModuleDesign) [][]string

func CalculateCoupling(pkg PackageInfo) int

func SuggestStructure(pkgCount int, hasAPI, hasDB bool) []string
```

## Examples

```go
// Define packages
models := PackageInfo{
    Name:         "models",
    Imports:      []string{},
    ExportsTypes: true,
    IsInternal:   false,
}

repo := PackageInfo{
    Name:         "repository",
    Imports:      []string{"models"},
    ExportsTypes: false,
    IsInternal:   true,
}

service := PackageInfo{
    Name:         "service",
    Imports:      []string{"models", "repository"},
    ExportsTypes: false,
    IsInternal:   true,
}

design := ModuleDesign{
    Packages: []PackageInfo{models, repo, service},
}

// Validate structure
problems := ValidateDesign(design)
// Returns: [] (no problems - good layering)

// Check for cycles
cycles := DetectCycles(design)
// Returns: [] (no cycles)

// Calculate coupling
coupling := CalculateCoupling(service)
// Returns: 2 (depends on 2 packages)

// Suggest structure
structure := SuggestStructure(5, true, true)
// Returns: ["models/", "internal/database/", "internal/service/", "api/", "cmd/"]
```

## Instructions

1. Open `module_structure.go`
2. Define `PackageInfo` and `ModuleDesign` types
3. Implement `ValidateDesign()` - find structural anti-patterns
4. Implement `DetectCycles()` - find circular dependencies
5. Implement `CalculateCoupling()` - count dependencies
6. Implement `SuggestStructure()` - recommend package layout
7. Run `go test -v`

## Hints

### Basic - Package Information
```go
type PackageInfo struct {
    Name         string   // Package name
    Imports      []string // Packages it depends on
    ExportsTypes bool     // Does it export types?
    IsInternal   bool     // Is it in internal/?
}

type ModuleDesign struct {
    Packages []PackageInfo
}
```

### Intermediate - Validation Rules
```go
func ValidateDesign(design ModuleDesign) []string {
    var problems []string

    for _, pkg := range design.Packages {
        // Rule 1: Packages should not have too many dependencies
        if len(pkg.Imports) > 5 {
            problems = append(problems,
                fmt.Sprintf("%s has too many dependencies (%d)",
                    pkg.Name, len(pkg.Imports)))
        }

        // Rule 2: Check for cycles
        // (use DetectCycles for this)

        // Rule 3: Internal packages should not import public packages
        if pkg.IsInternal {
            for _, imp := range pkg.Imports {
                // Check if importing non-internal package
                impPkg := findPackage(design, imp)
                if impPkg != nil && !impPkg.IsInternal && !impPkg.ExportsTypes {
                    problems = append(problems,
                        fmt.Sprintf("internal/%s imports non-internal %s",
                            pkg.Name, imp))
                }
            }
        }
    }

    return problems
}
```

### Advanced - Cycle Detection
```go
func DetectCycles(design ModuleDesign) [][]string {
    var cycles [][]string

    // Build adjacency map
    graph := make(map[string][]string)
    for _, pkg := range design.Packages {
        graph[pkg.Name] = pkg.Imports
    }

    // Track visited and recursion stack
    visited := make(map[string]bool)
    recStack := make(map[string]bool)
    path := []string{}

    var dfs func(string) bool
    dfs = func(pkg string) bool {
        visited[pkg] = true
        recStack[pkg] = true
        path = append(path, pkg)

        for _, dep := range graph[pkg] {
            if !visited[dep] {
                if dfs(dep) {
                    return true
                }
            } else if recStack[dep] {
                // Found cycle - extract it from path
                cycleStart := 0
                for i, p := range path {
                    if p == dep {
                        cycleStart = i
                        break
                    }
                }
                cycle := append([]string{}, path[cycleStart:]...)
                cycle = append(cycle, dep) // Close the cycle
                cycles = append(cycles, cycle)
                return true
            }
        }

        path = path[:len(path)-1]
        recStack[pkg] = false
        return false
    }

    for pkg := range graph {
        if !visited[pkg] {
            dfs(pkg)
        }
    }

    return cycles
}
```

### Coupling Calculation
```go
func CalculateCoupling(pkg PackageInfo) int {
    // Simple metric: number of imports
    return len(pkg.Imports)

    // Advanced metric could consider:
    // - Direct dependencies (imports)
    // - Indirect dependencies (transitive)
    // - Number of packages that depend on this package
}
```

### Structure Suggestion
```go
func SuggestStructure(pkgCount int, hasAPI, hasDB bool) []string {
    structure := []string{}

    // Always have models/types
    structure = append(structure, "models/")

    // Database layer if needed
    if hasDB {
        structure = append(structure, "internal/database/")
        structure = append(structure, "internal/repository/")
    }

    // Business logic
    structure = append(structure, "internal/service/")

    // API layer if needed
    if hasAPI {
        structure = append(structure, "api/")
        structure = append(structure, "api/handlers/")
    }

    // Command
    structure = append(structure, "cmd/")

    return structure
}
```

## 🧠 Think About

1. What makes a package "well-designed"?
2. How do you decide where to put a new piece of code?
3. What's the relationship between coupling and testability?
4. When should you split a package into two?
5. What are the characteristics of a good module structure?

## What This Teaches

- Package organization principles
- Coupling and cohesion metrics
- Detecting architectural problems
- Layered architecture patterns
- Best practices for module structure
- When to use internal/ packages

## Common Module Structures

### Simple API Server
```
myapp/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   ├── database/
│   ├── handlers/
│   └── service/
├── models/
└── go.mod
```

### Library
```
mylib/
├── mylib.go          # Main API
├── options.go        # Configuration
├── internal/
│   ├── parser/
│   └── validator/
└── go.mod
```

### Complex Application
```
myapp/
├── cmd/
│   ├── server/
│   └── cli/
├── pkg/               # Public APIs
│   └── client/
├── internal/
│   ├── api/
│   ├── auth/
│   ├── database/
│   ├── models/
│   └── service/
├── api/               # API definitions (OpenAPI, etc.)
├── docs/
└── go.mod
```

## Design Principles

### 1. Dependency Rule
Dependencies point inward (toward stable abstractions):
```
handlers → service → repository → models
```

### 2. Single Responsibility
Each package should have one reason to change:
- `models/` - data structures change
- `database/` - schema changes
- `api/` - API contract changes

### 3. Interface Segregation
Use interfaces to decouple:
```go
// service doesn't depend on concrete repository
type Repository interface {
    Get(id string) (Model, error)
}
```

### 4. Low Coupling, High Cohesion
- **Coupling**: Minimize dependencies between packages
- **Cohesion**: Related code stays together

## After Completing

Write in your `EXPLANATION.md`:
- The key principles of module structure
- How to identify high coupling
- The role of internal/ in large projects
- A complete module structure for a web application
- How you would refactor a poorly-structured project
- Why layered architecture prevents cycles
