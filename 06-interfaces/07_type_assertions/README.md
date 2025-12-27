# Exercise 07: Type Assertions

**Concept:** Checking concrete types at runtime with type assertions
**Difficulty:** Medium
**Estimated Time:** 50 minutes

## Learning Goal

Learn to safely check and convert interface values to concrete types using the comma-ok idiom.

## The Problem

When you have an interface value, sometimes you need to know the concrete type:

```go
var s Shape = Circle{Radius: 5}

// Type assertion with comma-ok (safe)
if c, ok := s.(Circle); ok {
    fmt.Println("Radius:", c.Radius)
}

// Type assertion without check (panics if wrong type!)
c := s.(Circle)  // Dangerous!
```

## Your Task

Work with Shape interface and different shape types.

### Interface

**Shape** - geometric shapes
- `Area() float64`

### Types

1. **Circle** - `Radius float64`
2. **Rectangle** - `Width, Height float64`

### Functions

1. **Diameter(s Shape) float64** - returns diameter if Circle, otherwise 0
2. **GetRadius(s Shape) (float64, bool)** - safely gets radius from Circle

## Function Signatures

```go
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (c Circle) Area() float64
func (r Rectangle) Area() float64

func Diameter(s Shape) float64
func GetRadius(s Shape) (float64, bool)
```

## Examples

```go
circle := Circle{Radius: 5}
rect := Rectangle{Width: 3, Height: 4}

Diameter(circle)  // → 10.0
Diameter(rect)    // → 0.0

r, ok := GetRadius(circle)  // → (5.0, true)
r, ok := GetRadius(rect)    // → (0.0, false)
```

## What This Teaches

- Type assertions
- Comma-ok idiom
- Safe type checking
- Working with concrete types from interfaces

---

**Next up:** Exercise 08 - Type Switch
