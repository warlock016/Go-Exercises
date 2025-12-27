# Exercise 06: Avoiding Cycles

**Concept:** Understanding and preventing import cycles
**Difficulty:** Medium
**Estimated Time:** 35 minutes

## 🎯 Learning Goal

Understand why import cycles are forbidden in Go and learn strategies to avoid them through proper architecture.

## The Problem

Go does not allow import cycles:

```
package A imports package B
package B imports package A
→ COMPILE ERROR: import cycle not allowed
```

This is intentional! Import cycles indicate:
- Poor separation of concerns
- Unclear dependencies
- Tightly coupled code
- Potential circular logic

## Common Scenarios

### Scenario 1: Mutual Dependencies
```
user package needs order types
order package needs user types
```

### Scenario 2: Shared Types
```
package A needs type from B
package B needs type from A
```

### Scenario 3: Helper Functions
```
package A has helper used by B
package B has helper used by A
```

## Solutions

### Solution 1: Extract Shared Types
Create a third package for shared types:
```
types package: User, Order
user package: imports types
order package: imports types
```

### Solution 2: Use Interfaces (Dependency Inversion)
```
package A: defines interface needed from B
package B: implements A's interface
```

### Solution 3: Merge Packages
If packages are too coupled, they might belong together.

## Your Task

Implement functions that demonstrate understanding of import cycles:

1. **AnalyzeDependency()** - Describes a dependency relationship
2. **DetectCycle()** - Detects if dependencies form a cycle
3. **SuggestFix()** - Suggests how to break the cycle

You'll work with conceptual dependency data.

## Function Signatures

```go
type Dependency struct {
    From string
    To   string
}

func AnalyzeDependency(from, to string) Dependency

func DetectCycle(deps []Dependency) bool

func SuggestFix(cycle []string) string
```

## Examples

```go
// Simple dependencies (no cycle)
deps := []Dependency{
    {"A", "B"},
    {"B", "C"},
    {"C", "D"},
}
DetectCycle(deps)  // false

// Cycle detected
deps = []Dependency{
    {"A", "B"},
    {"B", "C"},
    {"C", "A"},  // Cycle: A -> B -> C -> A
}
DetectCycle(deps)  // true

// Suggest fix
cycle := []string{"user", "order", "user"}
fix := SuggestFix(cycle)
// Returns suggestion like:
// "Extract shared types to a separate package (e.g., 'types')"
```

## Instructions

1. Open `avoiding_cycles.go`
2. Define the `Dependency` struct
3. Implement `AnalyzeDependency()`
4. Implement `DetectCycle()` - check for cycles in dependency graph
5. Implement `SuggestFix()` - provide architectural advice
6. Run `go test -v`

## Hints

### Basic - Understanding Cycles
```go
type Dependency struct {
    From string  // Package that imports
    To   string  // Package being imported
}

// A cycle exists if you can follow imports back to the starting package
// A -> B -> C -> A (cycle)
// A -> B -> C (no cycle)
```

### Intermediate - Detecting Cycles
```go
func DetectCycle(deps []Dependency) bool {
    // Build adjacency map: package -> list of packages it imports
    graph := make(map[string][]string)
    for _, dep := range deps {
        graph[dep.From] = append(graph[dep.From], dep.To)
    }

    // Track visited packages and current path
    visited := make(map[string]bool)
    recStack := make(map[string]bool)

    // Check each package
    for pkg := range graph {
        if hasCycle(pkg, graph, visited, recStack) {
            return true
        }
    }
    return false
}

func hasCycle(pkg string, graph map[string][]string, visited, recStack map[string]bool) bool {
    if recStack[pkg] {
        return true  // Found cycle!
    }
    if visited[pkg] {
        return false
    }

    visited[pkg] = true
    recStack[pkg] = true

    for _, dep := range graph[pkg] {
        if hasCycle(dep, graph, visited, recStack) {
            return true
        }
    }

    recStack[pkg] = false
    return false
}
```

### Advanced - Suggesting Fixes
```go
func SuggestFix(cycle []string) string {
    if len(cycle) < 2 {
        return "No cycle detected"
    }

    // Different suggestions based on cycle characteristics
    if len(cycle) == 2 {
        return "Consider merging these packages or extracting shared types to a third package"
    }

    if len(cycle) <= 4 {
        return "Extract shared types to a separate package (e.g., 'types' or 'models')"
    }

    return "Refactor to use interfaces (dependency inversion principle)"
}
```

## 🧠 Think About

1. Why does Go forbid import cycles when other languages allow them?
2. What does an import cycle reveal about your architecture?
3. Which solution (extract types, interfaces, merge) is best when?
4. How do import cycles relate to the dependency inversion principle?
5. Can you have a cycle involving 3+ packages?

## What This Teaches

- Why import cycles are forbidden
- How to detect dependency cycles
- Architectural patterns to avoid cycles
- Dependency inversion principle
- When to extract shared packages
- When to merge over-separated packages

## Real-World Strategies

### Strategy 1: Layered Architecture
```
models/      (types only, no imports)
repository/  (imports models)
service/     (imports models, repository)
handlers/    (imports service)
```

### Strategy 2: Dependency Inversion
```go
// Instead of service importing repository concretely:
package service
import "myapp/repository"
func NewService() *Service {
    return &Service{repo: repository.New()}  // Tight coupling!
}

// Use interfaces:
package service
type Repository interface {
    Get(id string) (Item, error)
}
func NewService(repo Repository) *Service {  // Loose coupling!
    return &Service{repo: repo}
}
```

### Strategy 3: Internal Packages
```
myapp/
  internal/
    types/     (shared types)
    db/        (imports types)
  api/         (imports internal/types, internal/db)
```

## After Completing

Write in your `EXPLANATION.md`:
- What an import cycle is and why it's forbidden
- The three strategies to avoid cycles
- When to use each strategy
- A real example from your code or a project
- How dependency inversion helps break cycles
