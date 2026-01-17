# Exercise 02: Memory Allocation

## Learning Goal
Learn to identify and reduce memory allocations to improve performance.

## Problem Description

Memory allocations have a cost: allocation time plus eventual garbage collection. Reducing unnecessary allocations is one of the most effective Go optimizations.

Key concepts:
- `b.ReportAllocs()` shows allocations per operation
- Stack allocations are free; heap allocations cost
- Reusing buffers/slices reduces allocations
- Small objects can sometimes be stack-allocated

## Function Signatures

```go
// ConcatStrings joins strings (allocates on each concat)
func ConcatStrings(strs []string) string

// ConcatStringsOptimized joins strings efficiently
func ConcatStringsOptimized(strs []string) string

// ProcessItems creates new slices for each step (allocates)
func ProcessItems(items []int) []int

// ProcessItemsInPlace modifies slice in place (zero alloc)
func ProcessItemsInPlace(items []int) []int

// CreateUsers creates user objects (allocates structs)
func CreateUsers(names []string) []*User

// CreateUsersOptimized creates users with fewer allocations
func CreateUsersOptimized(names []string) []User
```

## Hints

### Basic
- strings.Builder preallocates buffer space
- Modifying in place avoids new allocations
- `[]User` vs `[]*User` matters for allocations

### Intermediate
- Use `make([]T, 0, capacity)` to preallocate
- strings.Builder.Grow() reserves capacity upfront
- Returning structs by value can avoid heap allocation

## What This Teaches
- Identifying allocation sources
- Preallocation strategies
- Value vs pointer semantics for performance
