# Exercise 08: Type Switch

**Concept:** Handling multiple types with switch v.(type)
**Difficulty:** Medium-Hard
**Estimated Time:** 50 minutes

## Learning Goal

Learn the type switch pattern for handling multiple concrete types efficiently.

## The Problem

Type assertions are great for one type, but what if you need to handle many?

```go
// Tedious with multiple type assertions
if s, ok := v.(string); ok {
    // handle string
} else if i, ok := v.(int); ok {
    // handle int
} else if f, ok := v.(float64); ok {
    // handle float64
}

// Better with type switch
switch v := v.(type) {
case string:
    // v is string
case int:
    // v is int
case float64:
    // v is float64
default:
    // unknown type
}
```

## Your Task

Create functions that work with any type using type switches.

### Functions

1. **Stringify(v any) string** - converts any value to string
   - string → return as-is
   - int → convert to string
   - float64 → format with 2 decimals
   - bool → "true" or "false"
   - default → "unknown type"

2. **TypeName(v any) string** - returns the type name
   - Returns "string", "int", "float64", "bool", or "other"

3. **Add(a, b any) (any, bool)** - adds two values if same type
   - Both int → return int sum
   - Both float64 → return float64 sum
   - Both string → return concatenation
   - Different types → return (nil, false)

## Examples

```go
Stringify("hello")    // → "hello"
Stringify(42)         // → "42"
Stringify(3.14)       // → "3.14"
Stringify(true)       // → "true"

TypeName("hello")     // → "string"
TypeName(42)          // → "int"

Add(1, 2)             // → (3, true)
Add(1.5, 2.5)         // → (4.0, true)
Add("hi", " there")   // → ("hi there", true)
Add(1, "hello")       // → (nil, false)
```

## What This Teaches

- Type switch pattern
- Working with any/interface{}
- Handling heterogeneous data
- Type-safe operations on dynamic types

---

**Next up:** Exercise 09 - Mock for Testing
