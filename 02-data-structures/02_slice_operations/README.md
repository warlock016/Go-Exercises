# Exercise 02: Slice Operations

## 🎯 Learning Goal
Master advanced slice operations including copying, capacity management, understanding backing arrays, and slice growth patterns.

## 📝 Problem Description

While Exercise 01 covered slice basics, this exercise dives deeper into how slices work internally. You'll learn about the backing array, how capacity grows, and the critical differences between copying slices correctly vs incorrectly.

**Key Concepts:**
- Copying slices with `copy()` (not just assignment!)
- Understanding when slices share backing arrays
- Capacity growth patterns
- Full slice expressions `s[low:high:max]`

## 🔧 Function Signatures

```go
// CopySlice creates a true copy of the slice (not shared backing array)
func CopySlice(src []int) []int

// AppendWithCapacityCheck appends value to slice only if there's capacity
// Returns the slice and a bool indicating if append happened
func AppendWithCapacityCheck(slice []int, value int) ([]int, bool)

// ExtendSlice extends the slice to double its current length
// New elements should be zero-initialized
func ExtendSlice(slice []int) []int

// TruncateSlice returns a new slice with only the first half of elements
func TruncateSlice(slice []int) []int

// ReverseSlice reverses the slice in place and returns it
func ReverseSlice(slice []int) []int

// FilterEven returns a new slice containing only even numbers
func FilterEven(slice []int) []int

// InsertAt inserts value at the specified index (shifting elements right)
// If index is out of bounds, append to the end
func InsertAt(slice []int, index, value int) []int
```

## 💡 Examples

```go
// Copying
original := []int{1, 2, 3}
copied := CopySlice(original)
// copied is independent, changes don't affect original

// Capacity check
s := make([]int, 3, 5)  // len=3, cap=5
s2, ok := AppendWithCapacityCheck(s, 4)  // ok=true (had capacity)
s3, ok := AppendWithCapacityCheck(s2, 5) // ok=true (still has capacity)
s4, ok := AppendWithCapacityCheck(s3, 6) // ok=false (at capacity)

// Extending
s := []int{1, 2, 3}
extended := ExtendSlice(s)  // [1, 2, 3, 0, 0, 0]

// Reversing
s := []int{1, 2, 3, 4, 5}
ReverseSlice(s)  // [5, 4, 3, 2, 1]

// Filtering
s := []int{1, 2, 3, 4, 5, 6}
evens := FilterEven(s)  // [2, 4, 6]

// Inserting
s := []int{1, 2, 4, 5}
InsertAt(s, 2, 3)  // [1, 2, 3, 4, 5]
```

## 📋 Instructions

1. **CopySlice:** Create a new slice with `make()` and use `copy()` built-in
2. **AppendWithCapacityCheck:** Check `len(slice) < cap(slice)` before appending
3. **ExtendSlice:** Use `make()` with double the length, then copy original elements
4. **TruncateSlice:** Use slice expression to return first half
5. **ReverseSlice:** Swap elements from both ends moving inward
6. **FilterEven:** Create result slice and append even numbers (use `num % 2 == 0`)
7. **InsertAt:** Use append with three-way slice expressions

## 🧪 Testing

```bash
go test -v
```

Expected test count: ~30+ tests

## 🤔 Think About

1. **Why do we need the `copy()` function?**
   - Assignment `s2 := s1` just copies the slice header, both share the same backing array!

2. **How does Go grow slice capacity?**
   - Typically doubles capacity up to 1024 elements, then grows by 25%

3. **What's a slice header?**
   - A struct with three fields: pointer to backing array, length, capacity

4. **When do slices share backing arrays?**
   - When created with slice expressions: `s2 := s1[:]`

## 💡 Hints

<details>
<summary>Hint 1: Copying slices correctly</summary>

```go
// WRONG: Just copies the slice header, shares backing array
s1 := []int{1, 2, 3}
s2 := s1  // s1 and s2 share the same backing array!
s2[0] = 99  // s1[0] is now also 99!

// CORRECT: Use copy() to create independent slice
s1 := []int{1, 2, 3}
s2 := make([]int, len(s1))
copy(s2, s1)  // s2 is a true copy
s2[0] = 99    // s1[0] is still 1
```
</details>

<details>
<summary>Hint 2: Checking capacity before append</summary>

```go
func AppendWithCapacityCheck(slice []int, value int) ([]int, bool) {
    if len(slice) < cap(slice) {
        // Has room, can append without reallocation
        return append(slice, value), true
    }
    // At capacity, would need to reallocate
    return slice, false
}
```
</details>

<details>
<summary>Hint 3: Reversing in place</summary>

```go
// Two-pointer technique
func ReverseSlice(slice []int) []int {
    for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
        slice[i], slice[j] = slice[j], slice[i]  // Swap
    }
    return slice
}
```
</details>

<details>
<summary>Hint 4: Inserting at index</summary>

```go
// Method: append the slice to itself with inserted value
func InsertAt(slice []int, index, value int) []int {
    if index >= len(slice) {
        return append(slice, value)  // Out of bounds, append to end
    }

    // Make room for one more element
    slice = append(slice, 0)

    // Shift elements right
    copy(slice[index+1:], slice[index:])

    // Insert the value
    slice[index] = value

    return slice
}
```
</details>

## 🎓 What This Teaches

- **The copy() function** - Creating independent slices
- **Slice header semantics** - Understanding the pointer/length/capacity struct
- **Backing array sharing** - When slices are independent vs shared
- **Capacity growth** - How Go manages memory allocation for slices
- **In-place modification** - Reversing without allocating new memory
- **Filtering patterns** - Creating new slices with subset of elements
- **Insertion complexity** - O(n) operation requiring element shifting

---

**Next Exercise:** `03_map_fundamentals` - Hash tables and key-value operations
