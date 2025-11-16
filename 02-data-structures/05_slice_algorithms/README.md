# Exercise 05: Slice Algorithms

## 🎯 Learning Goal
Master functional-style operations on slices by implementing common algorithms like Map, Filter, Reduce, and more. Understand higher-order functions and predicate-based operations.

## 📝 Problem Description

While Go doesn't have built-in generics for collections (until Go 1.18+), you can still implement functional-style algorithms using functions as parameters. This exercise teaches you to work with higher-order functions, predicates, and transformations.

**Key Concepts:**
- Higher-order functions (functions that take functions as parameters)
- Predicate functions (functions that return bool)
- Transformation functions (mapping values)
- Reduction/aggregation patterns
- Set operations (uniqueness, membership)

## 🔧 Function Signatures

```go
// Map applies function to each element, returning new slice
func Map(slice []int, fn func(int) int) []int

// Filter keeps elements matching predicate, returning new slice
func Filter(slice []int, predicate func(int) bool) []int

// Reduce reduces slice to single value using accumulator function
func Reduce(slice []int, initial int, fn func(int, int) int) int

// Find returns first matching element and true, or zero value and false
func Find(slice []int, predicate func(int) bool) (int, bool)

// Contains checks if value exists in slice
func Contains(slice []int, value int) bool

// Unique returns slice with duplicates removed (maintains order)
func Unique(slice []int) []int

// Partition splits slice into two slices based on predicate
// First slice contains elements matching predicate, second contains rest
func Partition(slice []int, predicate func(int) bool) ([]int, []int)
```

## 💡 Examples

```go
// Map: double all values
nums := []int{1, 2, 3, 4}
doubled := Map(nums, func(x int) int { return x * 2 })
// [2, 4, 6, 8]

// Filter: keep only even numbers
nums := []int{1, 2, 3, 4, 5, 6}
evens := Filter(nums, func(x int) bool { return x % 2 == 0 })
// [2, 4, 6]

// Reduce: sum all elements
nums := []int{1, 2, 3, 4}
sum := Reduce(nums, 0, func(acc, x int) int { return acc + x })
// 10

// Find: first number > 5
nums := []int{1, 3, 7, 9}
val, found := Find(nums, func(x int) bool { return x > 5 })
// val=7, found=true

// Contains: check membership
nums := []int{1, 2, 3}
exists := Contains(nums, 2)  // true
exists = Contains(nums, 5)    // false

// Unique: remove duplicates
nums := []int{1, 2, 2, 3, 1, 4, 3}
unique := Unique(nums)
// [1, 2, 3, 4]

// Partition: split evens and odds
nums := []int{1, 2, 3, 4, 5, 6}
evens, odds := Partition(nums, func(x int) bool { return x % 2 == 0 })
// evens=[2, 4, 6], odds=[1, 3, 5]
```

## 📋 Instructions

1. **Map:** Create result slice, iterate with range, apply fn to each element
2. **Filter:** Create result slice, append elements where predicate returns true
3. **Reduce:** Start with initial value, iterate and accumulate using fn
4. **Find:** Iterate with range, return first element where predicate is true
5. **Contains:** Iterate and check equality (or use Find with predicate)
6. **Unique:** Use a map to track seen values, maintain insertion order
7. **Partition:** Create two result slices, append to first if predicate true, else second

## 🧪 Testing

```bash
go test -v
```

Expected test count: ~35+ tests

## 🤔 Think About

1. **Why pass functions as parameters?**
   - Enables code reuse - one Map function works for any transformation!

2. **What's a predicate function?**
   - A function that returns bool, used for testing conditions

3. **How does Reduce work?**
   - Starts with initial value, combines it with each element using accumulator function

4. **Why doesn't Find return just the value?**
   - Need to distinguish "found zero" from "not found" - hence the bool return

5. **Why use a map for Unique?**
   - Map lookups are O(1), making membership checks fast

## 💡 Hints

<details>
<summary>Hint 1: Map implementation</summary>

```go
func Map(slice []int, fn func(int) int) []int {
    result := make([]int, len(slice))
    for i, val := range slice {
        result[i] = fn(val)  // Apply transformation
    }
    return result
}
```
</details>

<details>
<summary>Hint 2: Filter implementation</summary>

```go
func Filter(slice []int, predicate func(int) bool) []int {
    result := []int{}
    for _, val := range slice {
        if predicate(val) {
            result = append(result, val)
        }
    }
    return result
}
```
</details>

<details>
<summary>Hint 3: Reduce implementation</summary>

```go
func Reduce(slice []int, initial int, fn func(int, int) int) int {
    acc := initial
    for _, val := range slice {
        acc = fn(acc, val)  // Accumulate
    }
    return acc
}
```
</details>

<details>
<summary>Hint 4: Unique using map</summary>

```go
func Unique(slice []int) []int {
    seen := make(map[int]bool)
    result := []int{}
    for _, val := range slice {
        if !seen[val] {
            seen[val] = true
            result = append(result, val)
        }
    }
    return result
}
```
</details>

<details>
<summary>Complete Solution</summary>

```go
func Map(slice []int, fn func(int) int) []int {
    result := make([]int, len(slice))
    for i, val := range slice {
        result[i] = fn(val)
    }
    return result
}

func Filter(slice []int, predicate func(int) bool) []int {
    result := []int{}
    for _, val := range slice {
        if predicate(val) {
            result = append(result, val)
        }
    }
    return result
}

func Reduce(slice []int, initial int, fn func(int, int) int) int {
    acc := initial
    for _, val := range slice {
        acc = fn(acc, val)
    }
    return acc
}

func Find(slice []int, predicate func(int) bool) (int, bool) {
    for _, val := range slice {
        if predicate(val) {
            return val, true
        }
    }
    return 0, false
}

func Contains(slice []int, value int) bool {
    for _, val := range slice {
        if val == value {
            return true
        }
    }
    return false
}

func Unique(slice []int) []int {
    seen := make(map[int]bool)
    result := []int{}
    for _, val := range slice {
        if !seen[val] {
            seen[val] = true
            result = append(result, val)
        }
    }
    return result
}

func Partition(slice []int, predicate func(int) bool) ([]int, []int) {
    matching := []int{}
    nonMatching := []int{}
    for _, val := range slice {
        if predicate(val) {
            matching = append(matching, val)
        } else {
            nonMatching = append(nonMatching, val)
        }
    }
    return matching, nonMatching
}
```
</details>

## 🎓 What This Teaches

- **Higher-order functions** - Functions as first-class values
- **Functional patterns** - Map/Filter/Reduce paradigm from functional programming
- **Predicates** - Boolean-returning functions for conditions
- **Accumulation** - Reducing collections to single values
- **Set operations** - Uniqueness and membership testing
- **Multiple return values** - Using (value, bool) pattern for optional results
- **Map-based algorithms** - Using hash tables for efficient lookups

---

**Next Exercise:** `06_map_patterns` - Common patterns for working with Go maps
