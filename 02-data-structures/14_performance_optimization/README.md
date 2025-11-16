# Exercise 14: Performance Optimization

## 🎯 Learning Goal
Master performance optimization techniques in Go: benchmarking, understanding allocation costs, preallocating slices, and using efficient string concatenation with strings.Builder. Learn to write both naive and optimized versions to see the performance difference.

## 📝 Problem Description

Performance matters. In production code, inefficient slice operations and string concatenation can cause significant slowdowns and excessive memory allocations. Go provides tools to measure performance (benchmarks) and techniques to optimize common operations.

In this exercise, you'll implement pairs of functions - naive versions and optimized versions - for three common performance bottlenecks:
1. **Slice append without preallocation** - Causes repeated reallocations
2. **String concatenation with +** - Creates many intermediate strings
3. **Filtering slices without capacity hints** - Wastes memory allocations

You'll write benchmark tests to measure the performance difference (expect 2-10x improvements!).

## 🔧 Function Signatures

Implement these functions in `performance_optimization.go`:

```go
// AppendNaive appends n integers (0 to n-1) to a slice without preallocation
func AppendNaive(n int) []int

// AppendOptimized appends n integers (0 to n-1) with proper preallocation
func AppendOptimized(n int) []int

// ConcatStringsNaive concatenates n copies of str using the + operator
func ConcatStringsNaive(str string, n int) string

// ConcatStringsOptimized concatenates n copies of str using strings.Builder
func ConcatStringsOptimized(str string, n int) string

// FilterNaive filters a slice, keeping only even numbers (no capacity hint)
func FilterNaive(numbers []int) []int

// FilterOptimized filters a slice, keeping only even numbers (with capacity hint)
func FilterOptimized(numbers []int) []int
```

## 💡 Examples

```go
// Append - Both produce same result, different performance
naive := AppendNaive(5)      // [0, 1, 2, 3, 4] - Multiple reallocations
optimized := AppendOptimized(5) // [0, 1, 2, 3, 4] - Single allocation

// String concatenation - Both produce same result
naive := ConcatStringsNaive("Go", 3)      // "GoGoGo" - 3 allocations
optimized := ConcatStringsOptimized("Go", 3) // "GoGoGo" - 1 allocation

// Filtering - Both produce same result
nums := []int{1, 2, 3, 4, 5, 6}
naive := FilterNaive(nums)      // [2, 4, 6] - May reallocate
optimized := FilterOptimized(nums) // [2, 4, 6] - Preallocated capacity
```

## 📋 Instructions

1. **AppendNaive:**
   - Create an empty slice: `result := []int{}`
   - Use a for loop from 0 to n-1
   - Append each number: `result = append(result, i)`
   - This causes multiple reallocations as the slice grows

2. **AppendOptimized:**
   - Preallocate with exact capacity: `result := make([]int, 0, n)`
   - Use the same for loop
   - Append each number - no reallocations needed!

3. **ConcatStringsNaive:**
   - Start with empty string: `result := ""`
   - Use a for loop from 0 to n-1
   - Concatenate: `result += str`
   - Each += creates a new string (very inefficient!)

4. **ConcatStringsOptimized:**
   - Import "strings" package
   - Create a Builder: `var builder strings.Builder`
   - Optionally hint capacity: `builder.Grow(len(str) * n)`
   - In loop: `builder.WriteString(str)`
   - Return: `builder.String()`

5. **FilterNaive:**
   - Create empty slice: `result := []int{}`
   - Range over input
   - If number is even (n%2 == 0), append to result

6. **FilterOptimized:**
   - Estimate capacity (worst case: all even): `result := make([]int, 0, len(numbers))`
   - Same filtering logic as naive version
   - Preallocated capacity prevents reallocations

## 🧪 Testing

Run tests with:
```bash
go test -v
```

Run benchmarks (this is the key part!):
```bash
go test -bench=. -benchmem
```

The benchmark output shows:
- **ns/op** - nanoseconds per operation (lower is better)
- **B/op** - bytes allocated per operation (lower is better)
- **allocs/op** - number of allocations per operation (lower is better)

Expected improvements:
- Append: 5-10x faster, 50% fewer allocations
- String concat: 10-100x faster for large n, 90% fewer allocations
- Filter: 2-3x faster, fewer allocations

## 🤔 Think About

1. **Why does preallocation help?**
   - Without preallocation, append() must allocate new arrays and copy data when capacity is exceeded
   - With preallocation, we allocate once with the exact size needed

2. **Why is string + operator slow?**
   - Strings are immutable in Go
   - Each += creates a new string, copying all previous content
   - For n concatenations, this is O(n²) time complexity!

3. **When should you preallocate?**
   - When you know (or can estimate) the final size
   - For loops that append many items
   - When performance profiling shows allocation hotspots

4. **What's the trade-off with capacity hints?**
   - Slight memory waste if estimate is too high
   - Massive performance gain if estimate is accurate
   - Usually worth it for performance-critical code

## 💡 Hints

<details>
<summary>Hint 1: Preallocation pattern</summary>

The key pattern for slice preallocation:
```go
// Naive - starts with capacity 0, grows dynamically
result := []int{}
for i := 0; i < n; i++ {
    result = append(result, i)  // May reallocate multiple times
}

// Optimized - preallocate exact capacity
result := make([]int, 0, n)  // len=0, cap=n
for i := 0; i < n; i++ {
    result = append(result, i)  // Never reallocates
}
```

Both produce the same result, but optimized is much faster!
</details>

<details>
<summary>Hint 2: strings.Builder usage</summary>

```go
import "strings"

// Naive string concatenation
func naive(str string, n int) string {
    result := ""
    for i := 0; i < n; i++ {
        result += str  // Creates new string each time!
    }
    return result
}

// Optimized with Builder
func optimized(str string, n int) string {
    var builder strings.Builder
    builder.Grow(len(str) * n)  // Optional but helpful capacity hint
    for i := 0; i < n; i++ {
        builder.WriteString(str)  // Efficient append
    }
    return builder.String()
}
```
</details>

<details>
<summary>Hint 3: Filter capacity estimation</summary>

For filtering, we don't know the exact final size, but we can estimate:

```go
// Naive - no capacity hint
result := []int{}
for _, num := range numbers {
    if num%2 == 0 {
        result = append(result, num)
    }
}

// Optimized - worst-case capacity (all elements pass filter)
result := make([]int, 0, len(numbers))
for _, num := range numbers {
    if num%2 == 0 {
        result = append(result, num)
    }
}
```

The optimized version may allocate slightly more than needed, but avoids reallocations.
</details>

<details>
<summary>Hint 4: Writing benchmarks</summary>

Benchmarks use the testing package:

```go
func BenchmarkAppendNaive(b *testing.B) {
    for i := 0; i < b.N; i++ {
        AppendNaive(1000)
    }
}

func BenchmarkAppendOptimized(b *testing.B) {
    for i := 0; i < b.N; i++ {
        AppendOptimized(1000)
    }
}
```

The testing framework runs the loop many times (b.N) to get accurate measurements.
Run with: `go test -bench=.`
</details>

<details>
<summary>Full Solution</summary>

```go
package performance_optimization

import "strings"

func AppendNaive(n int) []int {
    result := []int{}
    for i := 0; i < n; i++ {
        result = append(result, i)
    }
    return result
}

func AppendOptimized(n int) []int {
    result := make([]int, 0, n)
    for i := 0; i < n; i++ {
        result = append(result, i)
    }
    return result
}

func ConcatStringsNaive(str string, n int) string {
    result := ""
    for i := 0; i < n; i++ {
        result += str
    }
    return result
}

func ConcatStringsOptimized(str string, n int) string {
    var builder strings.Builder
    builder.Grow(len(str) * n)
    for i := 0; i < n; i++ {
        builder.WriteString(str)
    }
    return builder.String()
}

func FilterNaive(numbers []int) []int {
    result := []int{}
    for _, num := range numbers {
        if num%2 == 0 {
            result = append(result, num)
        }
    }
    return result
}

func FilterOptimized(numbers []int) []int {
    result := make([]int, 0, len(numbers))
    for _, num := range numbers {
        if num%2 == 0 {
            result = append(result, num)
        }
    }
    return result
}
```
</details>

## 🎓 What This Teaches

- **Benchmarking** - Using Go's testing.B to measure performance scientifically
- **Allocation costs** - Understanding memory allocations are expensive operations
- **Preallocation** - Using make([]T, 0, capacity) to avoid reallocations
- **strings.Builder** - The correct way to build strings efficiently
- **Capacity estimation** - When to preallocate and how to estimate sizes
- **Performance profiling** - Reading benchmark output (ns/op, B/op, allocs/op)
- **Optimization trade-offs** - Balancing code complexity with performance gains
- **Production patterns** - Real techniques used in high-performance Go code

---

**Next Exercise:** `15_generic_collections` - Building reusable data structures with interfaces or generics
