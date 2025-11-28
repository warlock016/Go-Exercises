# Exercise 09: Merge Sort

**Tier:** 3 - Divide & Conquer
**Estimated Time:** 70 minutes
**Concepts:** Merge sort, divide-recurse-combine, O(n log n) sorting, stable sorting

---

## Learning Goal

Master the classic divide-and-conquer sorting algorithm. Understand the "divide ’ sort recursively ’ merge" pattern that achieves O(n log n) performance.

---

## Problem Description

Implement merge sort - a recursive sorting algorithm that:
1. **Divides** the slice into two halves
2. **Recursively sorts** each half
3. **Merges** the two sorted halves

### Functions to Implement

1. **MergeSort** - Main sorting function
2. **Merge** - Combine two sorted slices into one sorted slice

---

## Function Signatures

```go
func MergeSort(nums []int) []int

func Merge(left, right []int) []int
```

---

## Examples

```go
MergeSort([]int{})                    // []
MergeSort([]int{5})                   // [5]
MergeSort([]int{3, 1, 2})             // [1, 2, 3]
MergeSort([]int{5, 2, 8, 1, 9})       // [1, 2, 5, 8, 9]
MergeSort([]int{3, 3, 1, 2})          // [1, 2, 3, 3]
```

```go
Merge([]int{1, 3, 5}, []int{2, 4, 6}) // [1, 2, 3, 4, 5, 6]
Merge([]int{1}, []int{2})             // [1, 2]
Merge([]int{}, []int{1, 2})           // [1, 2]
```

---

## Instructions

1. **Understand the algorithm**:
   ```
   MergeSort([5, 2, 8, 1, 9]):
     Divide: [5, 2, 8] and [1, 9]
       MergeSort([5, 2, 8]):
         Divide: [5, 2] and [8]
           MergeSort([5, 2]):
             Divide: [5] and [2]
             ’ [5], [2] (base case)
             Merge: [2, 5]
           MergeSort([8]):
             ’ [8] (base case)
         Merge: [2, 5, 8]
       MergeSort([1, 9]):
         Divide: [1] and [9]
         ’ [1], [9] (base case)
         Merge: [1, 9]
     Merge: [1, 2, 5, 8, 9]
   ```

2. **Implement Merge first** (easier):
   - Two pointers, one for each sorted slice
   - Compare elements, append smaller one
   - Append remaining elements from either slice

3. **Implement MergeSort**:
   - Base case: 0 or 1 elements ’ already sorted
   - Divide: find middle, split into left and right
   - Recurse: MergeSort(left), MergeSort(right)
   - Combine: Merge(sorted left, sorted right)

4. **Run tests**: `go test -v`

---

## Key Concepts

**Time Complexity:** O(n log n)
- log n levels (halving each time)
- n work at each level (merging)

**Space Complexity:** O(n)
- Need temporary slices for merging

**Stability:** Merge sort is stable (preserves order of equal elements)

---

## What This Teaches

 **Divide-recurse-combine** pattern
 **Merge operation** - combining sorted sequences
 **O(n log n) sorting** - optimal comparison-based sort
 **Recursion depth** - log n levels
 **Stable sorting** - important for certain applications

---

**Next Exercise:** 10 - Quick Sort
