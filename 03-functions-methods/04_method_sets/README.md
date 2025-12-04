# Exercise 04: Method Sets

**Learning Goal:** Master attaching methods to custom types beyond structs

---

## 📝 Problem Description

Methods aren't just for structs! You can attach methods to any named type. This is powerful for:
- Adding behavior to simple types (int, string, etc.)
- Creating type-safe wrappers
- Implementing domain-specific operations

You'll create several custom types with methods:
1. `Temperature` - Celsius/Fahrenheit conversions
2. `StringList` - String slice with utility methods
3. `Bytes` - Byte count formatter (KB, MB, GB)

---

## 🎯 Type Definitions & Method Signatures

```go
type Temperature float64

func (t Temperature) Celsius() float64
func (t Temperature) Fahrenheit() float64
func (t Temperature) Kelvin() float64

type StringList []string

func (s StringList) Join(separator string) string
func (s StringList) Contains(item string) bool
func (s StringList) Filter(predicate func(string) bool) StringList
func (s *StringList) Append(items ...string)

type Bytes int64

func (b Bytes) String() string
func (b Bytes) Kilobytes() float64
func (b Bytes) Megabytes() float64
func (b Bytes) Gigabytes() float64
```

---

## 📖 Examples

```go
// Temperature examples (stored as Celsius)
t := Temperature(25)
t.Celsius()      // 25.0
t.Fahrenheit()   // 77.0
t.Kelvin()       // 298.15

t2 := Temperature(0)
t2.Fahrenheit()  // 32.0

// StringList examples
list := StringList{"apple", "banana", "cherry"}
list.Join(", ")                    // "apple, banana, cherry"
list.Contains("banana")            // true
list.Contains("grape")             // false

filtered := list.Filter(func(s string) bool {
    return len(s) > 5
})                                 // StringList{"banana", "cherry"}

list.Append("date", "elderberry")  // modifies list

// Bytes examples
b := Bytes(1024)
b.String()        // "1.00 KB"
b.Kilobytes()     // 1.0
b.Megabytes()     // 0.0009765625

b2 := Bytes(5242880)
b2.String()       // "5.00 MB"
b2.Megabytes()    // 5.0
```

---

## 📋 Instructions

1. Define `Temperature` as `float64` (stores Celsius)
   - `Celsius()` returns value as-is
   - `Fahrenheit()` converts: `C * 9/5 + 32`
   - `Kelvin()` converts: `C + 273.15`

2. Define `StringList` as `[]string`
   - `Join()` combines with separator
   - `Contains()` checks for item presence
   - `Filter()` returns new list with matching items
   - `Append()` modifies list in place

3. Define `Bytes` as `int64`
   - `String()` formats with appropriate unit (B/KB/MB/GB)
   - `Kilobytes()` divides by 1024
   - `Megabytes()` divides by 1024²
   - `Gigabytes()` divides by 1024³

4. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Basic Concept</summary>

You can define methods on any named type:
```go
type MyInt int

func (m MyInt) Double() int {
    return int(m) * 2
}
```

The type must be defined in the same package as the methods.

</details>

<details>
<summary>Intermediate Hint</summary>

For `Temperature`:
- Store the value in Celsius
- Use conversion formulas for other units

For `StringList`:
- Use `strings.Join()` for `Join()`
- Iterate to check `Contains()`
- Create new slice for `Filter()`
- Use `append()` for `Append()` (pointer receiver!)

For `Bytes`:
- Use `fmt.Sprintf("%.2f KB", ...)` for formatting
- Choose unit based on magnitude (< 1024 = B, < 1024² = KB, etc.)

</details>

<details>
<summary>Complete Solution</summary>

```go
package method_sets

import (
	"fmt"
	"strings"
)

type Temperature float64

func (t Temperature) Celsius() float64 {
	return float64(t)
}

func (t Temperature) Fahrenheit() float64 {
	return float64(t)*9/5 + 32
}

func (t Temperature) Kelvin() float64 {
	return float64(t) + 273.15
}

type StringList []string

func (s StringList) Join(separator string) string {
	return strings.Join(s, separator)
}

func (s StringList) Contains(item string) bool {
	for _, str := range s {
		if str == item {
			return true
		}
	}
	return false
}

func (s StringList) Filter(predicate func(string) bool) StringList {
	result := StringList{}
	for _, str := range s {
		if predicate(str) {
			result = append(result, str)
		}
	}
	return result
}

func (s *StringList) Append(items ...string) {
	*s = append(*s, items...)
}

type Bytes int64

func (b Bytes) String() string {
	value := float64(b)
	switch {
	case b >= 1024*1024*1024:
		return fmt.Sprintf("%.2f GB", value/(1024*1024*1024))
	case b >= 1024*1024:
		return fmt.Sprintf("%.2f MB", value/(1024*1024))
	case b >= 1024:
		return fmt.Sprintf("%.2f KB", value/1024)
	default:
		return fmt.Sprintf("%d B", b)
	}
}

func (b Bytes) Kilobytes() float64 {
	return float64(b) / 1024
}

func (b Bytes) Megabytes() float64 {
	return float64(b) / (1024 * 1024)
}

func (b Bytes) Gigabytes() float64 {
	return float64(b) / (1024 * 1024 * 1024)
}
```

</details>

---

## 🤔 Think About

1. Why can't you add methods to built-in types like `int` directly?
2. What's the difference between `type StringList []string` and just using `[]string`?
3. When would you use a custom type wrapper vs just writing functions?
4. Why does `Append()` need a pointer receiver while `Filter()` doesn't?

---

## 🎓 What This Teaches

- **Named types**: Creating type aliases with `type Name UnderlyingType`
- **Methods on primitives**: Attaching behavior to simple types
- **Type safety**: Custom types prevent mixing incompatible values
- **Domain modeling**: Types represent concepts, not just data
- **Stringer interface**: `String()` method for custom formatting
- **Functional patterns**: Higher-order methods like `Filter()`

**Design patterns:**
- **Wrapper types**: Adding behavior to existing types
- **Fluent interfaces**: Methods that enable chaining
- **Type-driven design**: Using types to enforce constraints

---

**Tier:** 1 - Foundation
**Estimated Time:** 30-40 minutes
