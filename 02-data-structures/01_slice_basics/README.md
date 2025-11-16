# Exercise 01: Slice Basics

## 🎯 Learning Goal
Master the fundamentals of Go slices: creation methods, append operations, understanding length vs capacity, and basic iteration patterns.

## 📝 Problem Description

Slices are Go's most important data structure - dynamic, flexible arrays that grow automatically. Unlike arrays (which have fixed size), slices can change size during program execution.

In this exercise, you'll implement functions that demonstrate the core slice operations you'll use daily in Go programming.

## 🔧 Function Signatures

Implement these functions in `slice_basics.go`:

```go
// CreateSliceWithMake creates a slice of integers using make with specified length and capacity
func CreateSliceWithMake(length, capacity int) []int

// CreateSliceLiteral creates a slice containing the numbers 1, 2, 3, 4, 5
func CreateSliceLiteral() []int

// AppendToSlice appends the values to the slice and returns the result
func AppendToSlice(slice []int, values ...int) []int

// GetLengthAndCapacity returns both the length and capacity of the slice
func GetLengthAndCapacity(slice []int) (length, capacity int)

// IterateAndSum returns the sum of all elements in the slice
func IterateAndSum(slice []int) int

// IterateWithIndex returns a new slice where each element is doubled
func IterateWithIndex(slice []int) []int

// SliceFirst3 returns a new slice containing only the first 3 elements
// If the slice has fewer than 3 elements, return the entire slice
func SliceFirst3(slice []int) []int

// SliceFrom2ToEnd returns a new slice from index 2 to the end
// If the slice has fewer than 2 elements, return an empty slice
func SliceFrom2ToEnd(slice []int) []int
```

## 💡 Examples

```go
// Creating slices
s := CreateSliceWithMake(0, 10)  // len=0, cap=10, s = []

s2 := CreateSliceLiteral()        // [1, 2, 3, 4, 5]

// Appending
s3 := AppendToSlice([]int{1, 2}, 3, 4, 5)  // [1, 2, 3, 4, 5]

// Length and capacity
len, cap := GetLengthAndCapacity([]int{1, 2, 3})  // len=3, cap=3 (or more)

// Iteration
sum := IterateAndSum([]int{1, 2, 3, 4, 5})  // 15

doubled := IterateWithIndex([]int{1, 2, 3})  // [2, 4, 6]

// Slicing
first3 := SliceFirst3([]int{1, 2, 3, 4, 5})  // [1, 2, 3]

from2 := SliceFrom2ToEnd([]int{10, 20, 30, 40})  // [30, 40]
```

## 📋 Instructions

1. **CreateSliceWithMake:** Use the built-in `make()` function with the syntax `make([]int, length, capacity)`
2. **CreateSliceLiteral:** Use slice literal syntax: `[]int{...}`
3. **AppendToSlice:** Use the built-in `append()` function. Remember that `append()` returns a new slice.
4. **GetLengthAndCapacity:** Use the built-in `len()` and `cap()` functions
5. **IterateAndSum:** Use a `for` loop or `range` to iterate through the slice
6. **IterateWithIndex:** Create a new slice and use `range` with both index and value
7. **SliceFirst3:** Use slice expression `s[:3]`, but check length first
8. **SliceFrom2ToEnd:** Use slice expression `s[2:]`, but handle edge cases

## 🧪 Testing

Run tests with:
```bash
go test -v
```

Expected test count: ~25-30 tests across all functions

## 🤔 Think About

1. **What's the difference between length and capacity?**
   - Length is the number of elements currently in the slice
   - Capacity is the number of elements the slice can hold before needing to reallocate

2. **What happens when you append to a slice that's at capacity?**
   - Go allocates a new, larger backing array and copies the elements

3. **Why does `append()` return a slice?**
   - Because it might allocate a new backing array, changing the slice header

4. **What's the difference between `[]int{}` and `[]int(nil)`?**
   - `[]int{}` is an empty slice with a backing array (len=0, cap=0, not nil)
   - `[]int(nil)` is a nil slice (len=0, cap=0, is nil)

## 💡 Hints

<details>
<summary>Hint 1: Creating slices</summary>

There are three main ways to create a slice:
```go
// 1. Literal
s1 := []int{1, 2, 3}

// 2. Make with length and capacity
s2 := make([]int, 5, 10)  // len=5, cap=10, [0,0,0,0,0]

// 3. Make with just length (capacity = length)
s3 := make([]int, 5)  // len=5, cap=5, [0,0,0,0,0]
```
</details>

<details>
<summary>Hint 2: Appending to slices</summary>

The `append()` function is variadic and returns a new slice:
```go
s := []int{1, 2}
s = append(s, 3)        // s = [1, 2, 3]
s = append(s, 4, 5, 6)  // s = [1, 2, 3, 4, 5, 6]

// Append another slice (use ... to unpack)
s2 := []int{7, 8}
s = append(s, s2...)    // s = [1, 2, 3, 4, 5, 6, 7, 8]
```

**Important:** Always assign the result back, as `append()` might return a different slice!
</details>

<details>
<summary>Hint 3: Iteration patterns</summary>

Range gives you both index and value:
```go
s := []int{10, 20, 30}

// Both index and value
for i, val := range s {
    fmt.Printf("s[%d] = %d\n", i, val)
}

// Just values (ignore index with _)
for _, val := range s {
    sum += val
}

// Just indices
for i := range s {
    fmt.Println(i)
}
```
</details>

<details>
<summary>Hint 4: Slice expressions</summary>

Slicing syntax: `s[low:high]`
```go
s := []int{1, 2, 3, 4, 5}

s[1:4]   // [2, 3, 4]     (indices 1, 2, 3)
s[:3]    // [1, 2, 3]     (from beginning to index 3)
s[2:]    // [3, 4, 5]     (from index 2 to end)
s[:]     // [1, 2, 3, 4, 5]  (entire slice)
```

**Important:** Always check bounds to avoid panics!
```go
if len(s) >= 3 {
    first3 := s[:3]
}
```
</details>

<details>
<summary>Full Solution</summary>

```go
package slice_basics

func CreateSliceWithMake(length, capacity int) []int {
    return make([]int, length, capacity)
}

func CreateSliceLiteral() []int {
    return []int{1, 2, 3, 4, 5}
}

func AppendToSlice(slice []int, values ...int) []int {
    return append(slice, values...)
}

func GetLengthAndCapacity(slice []int) (length, capacity int) {
    return len(slice), cap(slice)
}

func IterateAndSum(slice []int) int {
    sum := 0
    for _, val := range slice {
        sum += val
    }
    return sum
}

func IterateWithIndex(slice []int) []int {
    result := make([]int, len(slice))
    for i, val := range slice {
        result[i] = val * 2
    }
    return result
}

func SliceFirst3(slice []int) []int {
    if len(slice) < 3 {
        return slice
    }
    return slice[:3]
}

func SliceFrom2ToEnd(slice []int) []int {
    if len(slice) < 2 {
        return []int{}
    }
    return slice[2:]
}
```
</details>

## 🎓 What This Teaches

- **Slice creation methods** - make() vs literals, when to use each
- **Dynamic arrays** - Slices grow automatically, unlike fixed arrays
- **Length vs Capacity** - Understanding the slice header's two size values
- **Append semantics** - Why append returns a value and when reallocation happens
- **Iteration patterns** - Using range effectively with slices
- **Slice expressions** - The `s[low:high]` syntax for creating sub-slices
- **Bounds checking** - Always validate indices before slicing

---

**Next Exercise:** `02_slice_operations` - Advanced slice operations, copying, and capacity growth
