# Exercise 13: Slice Internals

## 🎯 Learning Goal
Understand how slices work internally: backing arrays, capacity management, slice sharing, and the dangers of append. Master the full slice expression and deep vs shallow copying.

## 📝 Problem Description

Slices are Go's most powerful data structure, but their internal behavior can be surprising. Understanding slice internals is crucial because:

- **Append can cause subtle bugs** - Modifying one slice might affect another
- **Capacity affects performance** - Improper preallocation causes repeated reallocations
- **Backing arrays are shared** - Multiple slices can point to the same underlying array
- **Memory leaks can occur** - Holding references to large arrays via small slices

**Key Concepts:**

A slice is a descriptor with three fields:
```
type slice struct {
    ptr *array  // Pointer to backing array
    len int     // Number of elements
    cap int     // Capacity of backing array
}
```

When you slice `s[low:high]`:
- Creates a new slice header (pointer, len, cap)
- Both slices share the same backing array
- Modifying one affects the other (until append reallocates)

This exercise teaches you to visualize and control slice behavior at a deep level.

## 🔧 Function Signatures

Implement these functions in `slice_internals.go`:

```go
// SharesBackingArray returns true if s1 and s2 share the same backing array
// Hint: Compare the address of their first elements
func SharesBackingArray(s1, s2 []int) bool

// DemonstrateCapacityGrowth returns a slice showing how capacity grows
// Start with empty slice, append numbers 1-n, record capacity after each append
// Return slice of capacities: [1, 2, 4, 4, 8, 8, 8, 8, 16, ...]
func DemonstrateCapacityGrowth(n int) []int

// SafeAppend creates a new slice that doesn't share backing array with original
// This prevents append from affecting the original slice
func SafeAppend(original []int, value int) []int

// DeepCopy creates a completely independent copy of the slice
// No shared backing array
func DeepCopy(s []int) []int

// FullSliceExpression demonstrates s[low:high:max] syntax
// Returns slice with restricted capacity to prevent append from affecting original
// Given s and indices low, high, max, return s[low:high:max]
func FullSliceExpression(s []int, low, high, max int) []int

// TrimSlice removes elements from index i onwards but maintains same backing array
// Returns s[:i] but uses full slice expression to prevent capacity issues
func TrimSlice(s []int, i int) []int

// ExtendSlice tries to extend the slice by appending values
// Returns (newSlice, wasReallocated)
// wasReallocated is true if append caused reallocation (new backing array)
func ExtendSlice(s []int, values ...int) ([]int, bool)

// CountSliceAllocations appends n elements and returns number of reallocations
// Each reallocation happens when len == cap before append
func CountSliceAllocations(n int) int

// PreallocatedSlice creates a slice with specified capacity to avoid reallocations
// Should be able to append n elements without reallocating
func PreallocatedSlice(n int) []int
```

## 💡 Examples

```go
// Sharing backing arrays
s1 := []int{1, 2, 3, 4, 5}
s2 := s1[1:4]  // [2, 3, 4]
shared := SharesBackingArray(s1, s2)  // true

s3 := make([]int, 3)
copy(s3, s1)
shared = SharesBackingArray(s1, s3)  // false (different arrays)

// Capacity growth
capacities := DemonstrateCapacityGrowth(10)
// [1, 2, 4, 4, 8, 8, 8, 8, 16, 16] (growth pattern varies by Go version)

// Safe append
original := []int{1, 2, 3}
safe := SafeAppend(original, 4)
// safe = [1, 2, 3, 4], original unchanged

// Deep copy
s1 := []int{1, 2, 3}
s2 := DeepCopy(s1)
s2[0] = 99
// s1 = [1, 2, 3] (unchanged), s2 = [99, 2, 3]

// Full slice expression
s := []int{0, 1, 2, 3, 4, 5}
sub := FullSliceExpression(s, 1, 4, 4)  // s[1:4:4]
// sub = [1, 2, 3], len=3, cap=3 (not 5)

// Extend and detect reallocation
s := make([]int, 2, 5)  // len=2, cap=5
extended, reallocated := ExtendSlice(s, 1, 2, 3)  // Append 3 elements
// extended = [0, 0, 1, 2, 3], reallocated = false (cap was sufficient)

extended2, reallocated2 := ExtendSlice(extended, 6, 7, 8)
// reallocated2 = true (exceeded capacity)
```

## 📋 Instructions

1. **SharesBackingArray:** Use `&s1[0] == &s2[0]` to compare addresses (handle empty slices)
2. **DemonstrateCapacityGrowth:** Start with `var s []int`, append 1-n, record `cap(s)` after each append
3. **SafeAppend:** Create new slice with `make`, copy original, then append
4. **DeepCopy:** Use `copy(dest, src)` with newly allocated slice
5. **FullSliceExpression:** Return `s[low:high:max]` (validates indices)
6. **TrimSlice:** Return `s[:i:i]` to set cap=len, preventing append issues
7. **ExtendSlice:** Save old capacity, append values, check if capacity changed
8. **CountSliceAllocations:** Track when `len == cap`, increment counter before append
9. **PreallocatedSlice:** Use `make([]int, 0, n)` for zero-length, n-capacity slice

## 🧪 Testing

Run tests with:
```bash
go test -v
```

Expected test count: ~35-40 tests across all functions

## 🤔 Think About

1. **Why can append affect other slices?**
   - Multiple slices can share the same backing array
   - If there's capacity, append modifies the shared array
   - After reallocation, slices become independent

2. **What is the full slice expression `s[low:high:max]`?**
   - Limits capacity: `cap = max - low`
   - Prevents append from affecting elements beyond max
   - Syntax: `s[1:4:4]` means indices 1-3, cap=3

3. **How does capacity grow?**
   - Pre-Go 1.18: doubles when cap < 1024, then grows by 25%
   - Go 1.18+: smoother growth curve
   - Exact formula is implementation-dependent

4. **When should you preallocate slices?**
   - When you know the size in advance
   - Avoids repeated allocations (O(log n) instead of O(n))
   - Use `make([]T, 0, n)` for append, `make([]T, n)` for indexing

## 💡 Hints

<details>
<summary>Hint 1: Checking backing array sharing</summary>

Compare addresses of first elements:

```go
func SharesBackingArray(s1, s2 []int) bool {
    // Handle empty slices
    if len(s1) == 0 || len(s2) == 0 {
        return false
    }
    return &s1[0] == &s2[0]
}
```

**Caveat:** This only checks if the first element is the same. True sharing means any overlap.
</details>

<details>
<summary>Hint 2: Recording capacity growth</summary>

Track capacity after each append:

```go
func DemonstrateCapacityGrowth(n int) []int {
    var s []int
    capacities := make([]int, n)

    for i := 0; i < n; i++ {
        s = append(s, i+1)
        capacities[i] = cap(s)
    }

    return capacities
}
```
</details>

<details>
<summary>Hint 3: Safe append pattern</summary>

Create independent slice before appending:

```go
func SafeAppend(original []int, value int) []int {
    // Create new slice with enough capacity
    result := make([]int, len(original), len(original)+1)
    copy(result, original)
    return append(result, value)
}
```
</details>

<details>
<summary>Hint 4: Full slice expression</summary>

The third index controls capacity:

```go
s := []int{0, 1, 2, 3, 4, 5}

sub1 := s[1:4]      // [1,2,3], cap=5 (from index 1 to end)
sub2 := s[1:4:4]    // [1,2,3], cap=3 (from index 1 to index 4)

// Now append behaves differently:
sub1 = append(sub1, 99)  // Modifies s[4] (shared array)
sub2 = append(sub2, 99)  // Allocates new array (at capacity)
```
</details>

<details>
<summary>Hint 5: Detecting reallocation</summary>

Check if capacity increased:

```go
func ExtendSlice(s []int, values ...int) ([]int, bool) {
    oldCap := cap(s)
    s = append(s, values...)
    return s, cap(s) != oldCap
}
```

**More robust:** Compare addresses with `&s[0]` before/after (if len > 0).
</details>

<details>
<summary>Full Solution</summary>

See hints above for key patterns. Important concepts:

- Backing array is shared between slices from same origin
- Append reallocates when `len == cap`
- Full slice expression `s[low:high:max]` limits capacity
- Deep copy uses `copy()` with new allocation
- Preallocation with `make([]T, 0, n)` prevents repeated allocations
</details>

## 🎓 What This Teaches

- **Slice internals** - Understanding the three-field slice header (ptr, len, cap)
- **Backing array sharing** - How multiple slices can reference the same array
- **Append mechanics** - When reallocation occurs and why
- **Capacity growth** - How Go automatically grows slice capacity
- **Full slice expression** - Using `s[low:high:max]` to control capacity
- **Deep vs shallow copy** - When to use `copy()` vs simple assignment
- **Preallocation** - Optimizing performance by setting capacity upfront
- **Common pitfalls** - Avoiding bugs from unexpected slice sharing
- **Memory management** - Understanding when slices hold references to large arrays

---

**Next Module:** After completing all Module 02 exercises, proceed to Module 03: Functions & Methods

**Key Takeaway:** Slices are powerful but require understanding their internal representation to use safely and efficiently. Always consider whether slices share backing arrays when using append or modifying elements.
