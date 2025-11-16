# Exercise 04: Struct Basics

## 🎯 Learning Goal
Master the fundamentals of Go structs: defining custom types, struct instantiation patterns, field access, zero values, struct comparison, and basic methods.

## 📝 Problem Description

Structs are Go's way of creating custom data types that group related fields together. They're the foundation of object-oriented programming in Go (though Go doesn't have classes).

**Key Concepts:**
- Struct type definition with `type`
- Multiple instantiation patterns (literal, named fields, positional)
- Zero values for structs and their fields
- Field access with dot notation
- Struct comparison (when fields are comparable)
- Value semantics - structs are copied when passed to functions

## 🔧 Type Definitions

Define these struct types in `struct_basics.go`:

```go
// Person represents a person with a name and age
type Person struct {
	Name string
	Age  int
}

// Point represents a 2D coordinate
type Point struct {
	X int
	Y int
}

// Book represents a book with metadata
type Book struct {
	Title     string
	Author    string
	Pages     int
	Available bool
}
```

## 🔧 Function Signatures

```go
// CreatePerson creates and returns a Person with the given name and age
func CreatePerson(name string, age int) Person

// CreateZeroPerson returns a zero-valued Person (empty name, age 0)
func CreateZeroPerson() Person

// UpdatePersonAge returns a new Person with the age updated
// Note: This returns a copy, the original is not modified
func UpdatePersonAge(p Person, newAge int) Person

// ComparePersons returns true if both persons have the same name and age
func ComparePersons(p1, p2 Person) bool

// CreatePoint creates and returns a Point with the given x and y coordinates
func CreatePoint(x, y int) Point

// DistanceFromOrigin calculates the Euclidean distance from the origin (0, 0)
// Formula: sqrt(x^2 + y^2)
func DistanceFromOrigin(p Point) float64

// CreateBook creates and returns a Book with Available set to true by default
func CreateBook(title, author string, pages int) Book

// IsBookAvailable returns the availability status of the book
func IsBookAvailable(b Book) bool
```

## 💡 Examples

```go
// Creating structs
p1 := CreatePerson("Alice", 30)  // Person{Name: "Alice", Age: 30}

p2 := CreateZeroPerson()  // Person{Name: "", Age: 0}

// Updating (returns a copy!)
p3 := UpdatePersonAge(p1, 31)  // p1 is unchanged, p3 has age 31

// Comparing
same := ComparePersons(p1, p3)  // false (different ages)

// Points
pt := CreatePoint(3, 4)
dist := DistanceFromOrigin(pt)  // 5.0

// Books
book := CreateBook("1984", "George Orwell", 328)
// book.Available is true by default
available := IsBookAvailable(book)  // true
```

## 📋 Instructions

1. **Define the three struct types** at the package level (before the functions)
2. **CreatePerson:** Use struct literal with named fields
3. **CreateZeroPerson:** Return `Person{}` (zero value)
4. **UpdatePersonAge:** Create a new Person, copy fields, update age
5. **ComparePersons:** Use `==` operator (works when all fields are comparable)
6. **CreatePoint:** Use struct literal
7. **DistanceFromOrigin:** Use `math.Sqrt` and `math.Pow` from the math package
8. **CreateBook:** Set `Available: true` in the literal
9. **IsBookAvailable:** Return the `Available` field

## 🧪 Testing

```bash
go test -v
```

Expected test count: ~25 tests across all functions

## 🤔 Think About

1. **What's the difference between struct literals with and without field names?**
   - Named: `Person{Name: "Alice", Age: 30}` - Order independent, safer
   - Positional: `Person{"Alice", 30}` - Order matters, fragile to field reordering

2. **What happens when you pass a struct to a function?**
   - Go copies the entire struct (value semantics)
   - Changes inside the function don't affect the original
   - Use pointers (`*Person`) if you want to modify the original

3. **When can you compare structs with `==`?**
   - When all fields are comparable (numbers, strings, bools, pointers, comparable structs)
   - NOT if the struct contains slices, maps, or functions

4. **What are zero values for struct fields?**
   - `int` → 0
   - `string` → ""
   - `bool` → false
   - Pointers, slices, maps → `nil`

## 💡 Hints

<details>
<summary>Hint 1: Defining struct types</summary>

```go
// Define at package level, before functions
type Person struct {
    Name string
    Age  int
}

// Now you can use Person as a type
func CreatePerson(name string, age int) Person {
    return Person{Name: name, Age: age}
}
```
</details>

<details>
<summary>Hint 2: Struct instantiation patterns</summary>

```go
// Named fields (recommended - clear and order-independent)
p1 := Person{Name: "Alice", Age: 30}

// Positional (fragile - must match field order)
p2 := Person{"Bob", 25}

// Partial initialization (unspecified fields get zero values)
p3 := Person{Name: "Charlie"}  // Age is 0

// Zero value
p4 := Person{}  // Name is "", Age is 0

// Using new (returns pointer)
p5 := new(Person)  // &Person{Name: "", Age: 0}
```
</details>

<details>
<summary>Hint 3: Struct comparison</summary>

```go
p1 := Person{Name: "Alice", Age: 30}
p2 := Person{Name: "Alice", Age: 30}
p3 := Person{Name: "Bob", Age: 30}

p1 == p2  // true (all fields match)
p1 == p3  // false (names differ)

// Can also use in if statements
if p1 == p2 {
    fmt.Println("Same person!")
}
```
</details>

<details>
<summary>Hint 4: Updating structs (value semantics)</summary>

```go
func UpdatePersonAge(p Person, newAge int) Person {
    // Create a copy and modify it
    p.Age = newAge  // This modifies the copy, not the original!
    return p
}

// Or create a new struct
func UpdatePersonAge(p Person, newAge int) Person {
    return Person{
        Name: p.Name,
        Age:  newAge,
    }
}
```
</details>

<details>
<summary>Hint 5: Using math package</summary>

```go
import "math"

func DistanceFromOrigin(p Point) float64 {
    // sqrt(x^2 + y^2)
    return math.Sqrt(math.Pow(float64(p.X), 2) + math.Pow(float64(p.Y), 2))

    // Or simpler:
    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}
```

Don't forget to `import "math"` at the top of the file!
</details>

<details>
<summary>Full Solution</summary>

```go
package struct_basics

import "math"

// Person represents a person with a name and age
type Person struct {
	Name string
	Age  int
}

// Point represents a 2D coordinate
type Point struct {
	X int
	Y int
}

// Book represents a book with metadata
type Book struct {
	Title     string
	Author    string
	Pages     int
	Available bool
}

func CreatePerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

func CreateZeroPerson() Person {
	return Person{}
}

func UpdatePersonAge(p Person, newAge int) Person {
	p.Age = newAge
	return p
}

func ComparePersons(p1, p2 Person) bool {
	return p1 == p2
}

func CreatePoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func DistanceFromOrigin(p Point) float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func CreateBook(title, author string, pages int) Book {
	return Book{
		Title:     title,
		Author:    author,
		Pages:     pages,
		Available: true,
	}
}

func IsBookAvailable(b Book) bool {
	return b.Available
}
```
</details>

## 🎓 What This Teaches

- **Type definitions** - Creating custom types with `type`
- **Struct literals** - Named fields vs positional initialization
- **Zero values** - Default values for uninitialized struct fields
- **Field access** - Dot notation for reading and writing fields
- **Value semantics** - Structs are copied, not referenced (unlike slices/maps)
- **Struct comparison** - Using `==` when all fields are comparable
- **Package imports** - Using the math package for calculations
- **Design patterns** - Returning modified copies instead of mutating originals

---

**Next Exercise:** `05_slice_algorithms` - Common algorithmic patterns with slices
