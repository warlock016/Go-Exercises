# Exercise 03: Value vs Pointer Receivers

**Learning Goal:** Understand the critical difference between value and pointer receivers in methods

---

## 📝 Problem Description

Methods can have either value receivers `(t Type)` or pointer receivers `(t *Type)`. This choice affects:
- Whether the method can modify the receiver
- Performance (copying vs referencing)
- Which method set the type belongs to

You'll implement a `Counter` type with methods demonstrating both receiver types, and a `Rectangle` type showing when each is appropriate.

---

## 🎯 Type Definitions & Method Signatures

```go
type Counter struct {
    count int
}

// Value receiver - doesn't modify original
func (c Counter) Value() int

// Pointer receiver - modifies original
func (c *Counter) Increment()

// Pointer receiver - modifies original
func (c *Counter) Add(n int)

// Pointer receiver - modifies original
func (c *Counter) Reset()

type Rectangle struct {
    Width  float64
    Height float64
}

// Value receiver - read-only operation
func (r Rectangle) Area() float64

// Value receiver - read-only operation
func (r Rectangle) Perimeter() float64

// Pointer receiver - modifies original
func (r *Rectangle) Scale(factor float64)

// Pointer receiver - modifies original
func (r *Rectangle) SetDimensions(width, height float64)
```

---

## 📖 Examples

```go
// Counter examples
c := Counter{count: 0}
c.Value()              // 0
c.Increment()          // modifies c, count becomes 1
c.Value()              // 1
c.Add(5)               // modifies c, count becomes 6
c.Value()              // 6
c.Reset()              // modifies c, count becomes 0
c.Value()              // 0

// Rectangle examples
r := Rectangle{Width: 10, Height: 5}
r.Area()               // 50.0
r.Perimeter()          // 30.0
r.Scale(2)             // modifies r, width=20, height=10
r.Area()               // 200.0
r.SetDimensions(8, 4)  // modifies r
r.Area()               // 32.0
```

---

## 📋 Instructions

1. Define `Counter` struct with `count int` field
2. Implement `Value()` with value receiver (read-only)
3. Implement `Increment()`, `Add()`, `Reset()` with pointer receivers (modify state)
4. Define `Rectangle` struct with `Width` and `Height` fields
5. Implement `Area()` and `Perimeter()` with value receivers (calculations only)
6. Implement `Scale()` and `SetDimensions()` with pointer receivers (modify state)
7. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Basic Concept</summary>

**Value receiver** - operates on a copy:
```go
func (t Type) Method() {
    // t is a copy, changes don't affect original
}
```

**Pointer receiver** - operates on the original:
```go
func (t *Type) Method() {
    // t is a pointer, changes affect original
}
```

**Rule of thumb:**
- Use pointer receivers when you need to modify the receiver
- Use value receivers for read-only operations on small types
- Use pointer receivers for large structs (avoid copying)

</details>

<details>
<summary>Intermediate Hint</summary>

For `Counter`:
- `Value()` just returns `c.count` (no modification)
- `Increment()` does `c.count++`
- `Add(n)` does `c.count += n`
- `Reset()` does `c.count = 0`

For `Rectangle`:
- `Area()` returns `r.Width * r.Height`
- `Perimeter()` returns `2 * (r.Width + r.Height)`
- `Scale(factor)` multiplies both dimensions by factor
- `SetDimensions()` assigns new width and height

</details>

<details>
<summary>Complete Solution</summary>

```go
package value_vs_pointer_receivers

type Counter struct {
	count int
}

// Value receiver - read-only
func (c Counter) Value() int {
	return c.count
}

// Pointer receiver - modifies
func (c *Counter) Increment() {
	c.count++
}

// Pointer receiver - modifies
func (c *Counter) Add(n int) {
	c.count += n
}

// Pointer receiver - modifies
func (c *Counter) Reset() {
	c.count = 0
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Value receiver - calculation only
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Value receiver - calculation only
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Pointer receiver - modifies
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Pointer receiver - modifies
func (r *Rectangle) SetDimensions(width, height float64) {
	r.Width = width
	r.Height = height
}
```

</details>

---

## 🤔 Think About

1. What happens if you use a value receiver for `Increment()`?
2. Why can you call pointer receiver methods on values (e.g., `counter.Increment()` instead of `(&counter).Increment()`)?
3. When would you use value receivers for large structs?
4. What are the implications for method sets with interfaces?

---

## 🎓 What This Teaches

- **Receiver types**: Value `(t Type)` vs pointer `(t *Type)`
- **Mutability**: Pointer receivers can modify, value receivers cannot
- **Performance**: Copying vs referencing considerations
- **Design decisions**: When to choose each receiver type
- **Go syntactic sugar**: Automatic referencing/dereferencing
- **Method sets**: How receiver type affects interface satisfaction

**Guidelines for choosing receivers:**
1. **Must use pointer** if method modifies receiver
2. **Use pointer** for large structs (avoid copying)
3. **Use pointer** if any method has pointer receiver (consistency)
4. **Use value** for small, immutable types
5. **Use value** for read-only operations on simple types

---

**Tier:** 1 - Foundation
**Estimated Time:** 30-40 minutes
