# Exercise 03: Import Paths

**Concept:** Import conventions and patterns
**Difficulty:** Easy
**Estimated Time:** 20 minutes

## 🎯 Learning Goal

Master Go's import patterns: standard imports, aliasing, blank imports, and import organization.

## The Problem

Go has several import patterns:

1. **Standard import:** `import "fmt"`
2. **Import aliasing:** `import f "fmt"` (use as `f.Println()`)
3. **Dot import:** `import . "fmt"` (use as `Println()` - rarely used)
4. **Blank import:** `import _ "package"` (run init() only, for side effects)
5. **Grouped imports:** Organize stdlib, external, and internal packages

## Your Task

Create functions demonstrating different import patterns:

1. `StandardImport()` - Uses standard library packages normally
2. `AliasedImport()` - Uses aliased imports for clarity
3. `GroupedImports()` - Demonstrates proper import organization
4. `BlankImportExample()` - Explains when to use blank imports

## Function Signatures

```go
func StandardImport(data map[string]int) string
func AliasedImport(text string) (string, error)
func GroupedImports(n int) []string
func BlankImportExample() string
```

## Examples

```go
// Standard import
import "fmt"
fmt.Println("Hello")

// Aliased import (useful for long package names)
import str "strings"
str.ToUpper("hello")  // instead of strings.ToUpper

// Blank import (side effects only)
import _ "github.com/lib/pq"  // Registers PostgreSQL driver
// Can't call any pq functions, but init() ran

// Grouped imports (conventional style)
import (
    "fmt"           // Standard library
    "strings"

    "github.com/pkg/errors"  // External

    "myapp/internal/db"      // Internal
)
```

## Instructions

1. Open `import_paths.go`
2. Implement `StandardImport()` using encoding/json to marshal the map
3. Implement `AliasedImport()` using aliased imports for strings and encoding/base64
4. Implement `GroupedImports()` demonstrating import organization
5. Implement `BlankImportExample()` returning a description
6. Run `go test -v`

## Hints

### Basic - Import Styles
```go
// Standard
import "fmt"
import "strings"

// Grouped (preferred)
import (
    "fmt"
    "strings"
)

// Aliased
import (
    f "fmt"
    str "strings"
)

// Use as:
f.Println("Hello")
str.ToUpper("hello")
```

### Intermediate - When to Alias
```go
// When package names conflict
import (
    "crypto/rand"
    mathrand "math/rand"  // Alias to avoid conflict
)

// When package name doesn't match import path
import (
    jwt "github.com/golang-jwt/jwt/v5"
)

// For clarity with long names
import (
    pb "mycompany.com/api/protobuf/v1/users"
)
```

### Advanced - Complete Example
```go
package imports

import (
    "encoding/base64"
    "encoding/json"
    "fmt"
    str "strings"  // Aliased for demonstration
)

func StandardImport(data map[string]int) string {
    bytes, err := json.Marshal(data)
    if err != nil {
        return ""
    }
    return string(bytes)
}

func AliasedImport(text string) (string, error) {
    // Using aliased strings package
    upper := str.ToUpper(text)
    // Using base64 (not aliased)
    encoded := base64.StdEncoding.EncodeToString([]byte(upper))
    return encoded, nil
}

func GroupedImports(n int) []string {
    // This function demonstrates proper import grouping
    // See the imports at the top of the file
    result := make([]string, n)
    for i := 0; i < n; i++ {
        result[i] = fmt.Sprintf("item-%d", i)
    }
    return result
}

func BlankImportExample() string {
    return "Blank imports are used for side effects, like registering database drivers: import _ \"github.com/lib/pq\""
}
```

### Import Organization Best Practices
```go
import (
    // Group 1: Standard library
    "encoding/json"
    "fmt"
    "strings"

    // Group 2: External packages (leave blank line)
    "github.com/pkg/errors"
    "github.com/gorilla/mux"

    // Group 3: Internal packages (leave blank line)
    "mycompany.com/myapp/internal/db"
    "mycompany.com/myapp/pkg/models"
)
```

## 🧠 Think About

1. Why would you alias an import instead of using the default name?
2. When is a blank import (`import _`) appropriate?
3. Why organize imports into groups (stdlib, external, internal)?
4. What happens if you import a package but don't use it?
5. Can you have two imports with the same package name without aliasing?

## What This Teaches

- Different import syntaxes and their use cases
- Import aliasing for clarity and conflict resolution
- Blank imports for side effects (init() functions)
- Conventional import organization
- How package names relate to import paths

## After Completing

Write in your `EXPLANATION.md`:
- The three main import patterns you learned
- When to use import aliasing
- What blank imports are for
- How to organize imports conventionally
