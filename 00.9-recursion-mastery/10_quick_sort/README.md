# Exercise 10: Quick Sort

**Tier:** 3 - Divide & Conquer
**Estimated Time:** 70 minutes
**Concepts:** Quick sort, partitioning, in-place sorting, pivot selection

## Learning Goal

Implement quick sort - another O(n log n) divide-and-conquer algorithm that sorts in-place using partitioning.

## Algorithm

1. Choose a pivot element
2. Partition array: elements < pivot on left, > pivot on right
3. Recursively sort left and right partitions

## Function Signatures

```go
func QuickSort(nums []int) []int
func Partition(nums []int, low, high int) int
```

## Examples

```go
QuickSort([]int{5, 2, 8, 1, 9}) // [1, 2, 5, 8, 9]
QuickSort([]int{3, 3, 1, 2})    // [1, 2, 3, 3]
```

## Key Concepts

- **Average case:** O(n log n)
- **Worst case:** O(n²) if pivot is always min/max
- **In-place:** O(1) extra space (better than merge sort)
- **Not stable:** May change order of equal elements

**Next Exercise:** 11 - Subsets & Combinations
