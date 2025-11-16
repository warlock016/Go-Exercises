# Exercise 12: Custom Types

## 🎯 Learning Goal
Learn to define custom types with methods, understanding how Go's type system enables domain-specific abstractions. Preview methods (Module 03) while reinforcing data structure concepts.

## 📝 Problem Description

Custom types in Go let you create domain-specific abstractions with their own methods. This makes code more:

- **Expressive** - `temp.ToFahrenheit()` vs `celsiusToFahrenheit(temp)`
- **Type-safe** - Can't accidentally add `Distance` to `Temperature`
- **Self-documenting** - Methods show what operations are valid
- **Encapsulated** - Hide implementation details

Common patterns:
- **Type aliases for primitives** - `type Temperature float64` (adds methods to numbers)
- **Struct-based types** - `type StringSet map[string]struct{}` (custom collections)
- **Unit conversions** - Temperature (C/F), Distance (m/km), Time (s/ms)
- **Set implementations** - Using maps with empty struct values

This exercise bridges Module 02 (data structures) and Module 03 (methods), showing how custom types combine both concepts.

## 🔧 Type Definitions and Function Signatures

Define these types and methods in `custom_types.go`:

### Temperature Type

```go
// Temperature represents a temperature in Celsius
type Temperature float64

// ToFahrenheit converts Celsius to Fahrenheit
func (t Temperature) ToFahrenheit() float64

// ToCelsius returns the temperature in Celsius (identity function)
func (t Temperature) ToCelsius() float64

// Add adds two temperatures
func (t Temperature) Add(other Temperature) Temperature

// IsFrezing returns true if temperature is at or below 0°C
func (t Temperature) IsFreezing() bool
```

### Distance Type

```go
// Distance represents a distance in meters
type Distance int

// ToKilometers converts meters to kilometers
func (d Distance) ToKilometers() float64

// ToMeters returns the distance in meters (identity function)
func (d Distance) ToMeters() int

// Add adds two distances
func (d Distance) Add(other Distance) Distance

// IsMarathon returns true if distance is at least 42195 meters (marathon distance)
func (d Distance) IsMarathon() bool
```

### StringSet Type

```go
// StringSet represents a set of unique strings
// Implemented as map[string]struct{} for memory efficiency
type StringSet map[string]struct{}

// NewStringSet creates a new empty StringSet
func NewStringSet() StringSet

// Add adds a string to the set
func (s StringSet) Add(item string)

// Remove removes a string from the set
func (s StringSet) Remove(item string)

// Contains returns true if the string is in the set
func (s StringSet) Contains(item string) bool

// Size returns the number of elements in the set
func (s StringSet) Size() int

// ToSlice returns all elements as a slice
func (s StringSet) ToSlice() []string

// Union returns a new set containing all elements from both sets
func (s StringSet) Union(other StringSet) StringSet

// Intersection returns a new set containing only elements in both sets
func (s StringSet) Intersection(other StringSet) StringSet
```

## 💡 Examples

```go
// Temperature
temp := Temperature(20.0)
f := temp.ToFahrenheit()  // 68.0
c := temp.ToCelsius()      // 20.0
sum := temp.Add(Temperature(5.0))  // 25.0
freezing := temp.IsFreezing()  // false

// Distance
dist := Distance(5000)
km := dist.ToKilometers()  // 5.0
m := dist.ToMeters()       // 5000
total := dist.Add(Distance(1000))  // 6000
isMarathon := dist.IsMarathon()  // false

// StringSet
set := NewStringSet()
set.Add("apple")
set.Add("banana")
set.Add("apple")  // Duplicate, no effect

contains := set.Contains("apple")  // true
size := set.Size()  // 2

set.Remove("banana")
slice := set.ToSlice()  // ["apple"]

// Set operations
set1 := NewStringSet()
set1.Add("a")
set1.Add("b")

set2 := NewStringSet()
set2.Add("b")
set2.Add("c")

union := set1.Union(set2)  // {"a", "b", "c"}
inter := set1.Intersection(set2)  // {"b"}
```

## 📋 Instructions

1. **Temperature:**
   - ToFahrenheit: `(celsius * 9/5) + 32`
   - ToCelsius: return `float64(t)`
   - Add: return `t + other`
   - IsFreezing: return `t <= 0`

2. **Distance:**
   - ToKilometers: `float64(d) / 1000`
   - ToMeters: return `int(d)`
   - Add: return `d + other`
   - IsMarathon: return `d >= 42195`

3. **StringSet:**
   - NewStringSet: `return make(StringSet)` or `return make(map[string]struct{})`
   - Add: `s[item] = struct{}{}`
   - Remove: `delete(s, item)`
   - Contains: use comma-ok idiom `_, exists := s[item]`
   - Size: return `len(s)`
   - ToSlice: iterate through map, collect keys
   - Union: create new set, add all from both sets
   - Intersection: create new set, add only common elements

## 🧪 Testing

Run tests with:
```bash
go test -v
```

Expected test count: ~40-45 tests across all types

## 🤔 Think About

1. **Why use `Temperature` instead of just `float64`?**
   - Type safety: can't accidentally mix temperatures with other floats
   - Methods: can call `temp.ToFahrenheit()` directly
   - Intent: makes code self-documenting

2. **What is `struct{}`? Why use it for sets?**
   - Empty struct uses 0 bytes (smallest possible type)
   - `map[string]struct{}` only stores keys, values are ignored
   - More memory-efficient than `map[string]bool`

3. **What's the difference between `(t Temperature)` and `(t *Temperature)` methods?**
   - Value receiver: operates on a copy (can't modify original)
   - Pointer receiver: operates on original (can modify, more efficient for large structs)

4. **Why return new sets for Union/Intersection instead of modifying existing ones?**
   - Functional style: doesn't modify inputs (immutable operations)
   - Safer: caller still has original sets
   - More flexible: can chain operations

## 💡 Hints

<details>
<summary>Hint 1: Method syntax</summary>

Methods are functions with a receiver:

```go
// Regular function
func ToFahrenheit(t Temperature) float64 {
    return (float64(t) * 9/5) + 32
}

// Method (preferred)
func (t Temperature) ToFahrenheit() float64 {
    return (float64(t) * 9/5) + 32
}

// Called as:
temp := Temperature(20.0)
f := temp.ToFahrenheit()  // Method style
```
</details>

<details>
<summary>Hint 2: Empty struct for sets</summary>

Why `struct{}` is better than `bool`:

```go
// Less efficient (1 byte per value)
type StringSet map[string]bool

set["key"] = true

// More efficient (0 bytes per value)
type StringSet map[string]struct{}

set["key"] = struct{}{}  // Empty struct literal
```

The empty struct uses no memory!
</details>

<details>
<summary>Hint 3: Set operations</summary>

Union and intersection patterns:

```go
// Union: everything from both sets
func (s StringSet) Union(other StringSet) StringSet {
    result := NewStringSet()
    for key := range s {
        result.Add(key)
    }
    for key := range other {
        result.Add(key)
    }
    return result
}

// Intersection: only common elements
func (s StringSet) Intersection(other StringSet) StringSet {
    result := NewStringSet()
    for key := range s {
        if other.Contains(key) {
            result.Add(key)
        }
    }
    return result
}
```
</details>

<details>
<summary>Hint 4: Converting set to slice</summary>

Collect map keys into a slice:

```go
func (s StringSet) ToSlice() []string {
    result := make([]string, 0, len(s))
    for key := range s {
        result = append(result, key)
    }
    return result
}
```

**Note:** Map iteration order is random, so slice order is non-deterministic.
</details>

<details>
<summary>Full Solution</summary>

See solution in hints above. Key patterns:
- Methods attach to types with receiver syntax `(t Type)`
- Empty struct `struct{}{}` is zero-size, perfect for sets
- Set operations create new sets (functional style)
- Type conversions enable domain-specific abstractions
</details>

## 🎓 What This Teaches

- **Custom type definitions** - Creating domain-specific types from primitives
- **Method syntax** - Attaching functions to types with receivers
- **Type safety** - Preventing mixing of incompatible types
- **Set implementation** - Using `map[string]struct{}` for efficient sets
- **Empty struct optimization** - Zero-byte values for memory efficiency
- **Functional operations** - Immutable set operations (union, intersection)
- **Method chaining** - Designing types that work well together
- **Domain modeling** - Using types to represent real-world concepts
- **Preview of Module 03** - Methods, receivers, and object-oriented patterns

---

**Next Exercise:** `13_slice_internals` - Deep dive into slice backing arrays and capacity
