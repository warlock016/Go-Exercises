# Exercise 07: Two Pointers

**Difficulty:** Medium
**Time:** 30-40 minutes
**Concepts:** Two pointer technique, in-place array manipulation, algorithm optimization

## Learning Goal

Learn the two-pointer technique, a fundamental algorithmic pattern for efficiently processing arrays. This pattern is widely used in coding interviews and real-world scenarios where you need to process arrays with O(1) extra space.

## The Problem

Implement a function that removes duplicates from a sorted array in-place, returning a new array containing only unique elements.

**Important:** The array is already sorted, which means duplicates are adjacent.

```
Input:  [1, 1, 2, 2, 3]
Output: [1, 2, 3]

Input:  [1, 2, 3, 4]
Output: [1, 2, 3, 4]  (no duplicates)

Input:  [1, 1, 1, 1]
Output: [1]  (all duplicates)
```

## Function Signatures

```go
func RemoveDuplicates(nums []int) []int
```

## Examples

```go
result := RemoveDuplicates([]int{1, 1, 2, 2, 3})
// result: [1, 2, 3]

result = RemoveDuplicates([]int{1, 2, 3, 4})
// result: [1, 2, 3, 4]

result = RemoveDuplicates([]int{1, 1, 1, 1})
// result: [1]

result = RemoveDuplicates([]int{})
// result: []

result = RemoveDuplicates([]int{5})
// result: [5]

result = RemoveDuplicates([]int{1, 1, 2, 3, 3, 3, 4, 4, 5})
// result: [1, 2, 3, 4, 5]
```

## Instructions

1. Use two pointers: a "slow" pointer and a "fast" pointer
2. The slow pointer tracks where to write the next unique element
3. The fast pointer scans through the array
4. When you find a new unique element, write it at the slow pointer position and advance slow
5. Return the slice from the beginning to the slow pointer position

## Hints

### Basic Hint: Two Pointer Concept

Think of two pointers moving through the array:
- **Fast pointer:** Scans every element
- **Slow pointer:** Points to where the next unique element should be written

Example walkthrough for `[1, 1, 2, 2, 3]`:
```
Initial: slow=0, fast=0
         [1, 1, 2, 2, 3]
          ^

Step 1: fast=1, nums[fast]=1 == nums[slow]=1 (duplicate, skip)
        [1, 1, 2, 2, 3]
         ^  ^

Step 2: fast=2, nums[fast]=2 != nums[slow]=1 (unique!)
        slow++, nums[slow]=2
        [1, 2, 2, 2, 3]
            ^  ^

Continue this pattern...
```

### Intermediate Hint: Algorithm Structure

```go
func RemoveDuplicates(nums []int) []int {
    if len(nums) == 0 {
        return nums
    }

    slow := 0

    for fast := 1; fast < len(nums); fast++ {
        if nums[fast] != nums[slow] {
            // Found a unique element
            slow++
            nums[slow] = nums[fast]
        }
    }

    return nums[:slow+1]
}
```

### Advanced Hint: Why This Works

The key insight: **In a sorted array, duplicates are adjacent**.

- We keep the first occurrence at `nums[slow]`
- We scan ahead with `fast`
- When `nums[fast] != nums[slow]`, we found a new unique value
- We move `slow` forward and copy the new value there
- All elements from `0` to `slow` are unique

Time complexity: O(n) - single pass through array
Space complexity: O(1) - only two integer pointers

### Challenge: Generalize the Pattern

Once you solve this, you've learned a pattern that applies to many problems:
- Remove duplicates (this exercise)
- Move zeros to end
- Partition array
- Merge sorted arrays

All use variants of the two-pointer technique!

## Think About

1. Why does this algorithm only work for sorted arrays?
2. What would change if the array wasn't sorted?
3. How would you modify this to keep at most 2 duplicates instead of removing all duplicates?
4. Could you solve this with just one pointer? Why or why not?

## What This Teaches

- **Two pointer technique:** A fundamental algorithmic pattern
- **In-place algorithms:** Modifying data without extra space
- **Array manipulation:** Understanding slice operations
- **Optimization thinking:** Achieving O(n) time with O(1) space
- **Pattern recognition:** Identifying when two pointers are useful

Ready to implement? Open `two_pointers.go` and look for `TODO(human)` markers!
