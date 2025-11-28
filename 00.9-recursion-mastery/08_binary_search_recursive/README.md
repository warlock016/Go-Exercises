# Exercise 08: Binary Search Recursive

**Tier:** 3 - Divide & Conquer
**Estimated Time:** 55 minutes
**Concepts:** Binary search, divide-and-conquer, logarithmic complexity, search bounds

---

## Learning Goal

Master divide-and-conquer with binary search - the fundamental algorithm that halves the problem at each step. Understand how recursion naturally expresses "divide the problem in half."

---

## Problem Description

Implement binary search and related algorithms recursively.

### 1. Binary Search

Find the index of a target value in a sorted slice. Return -1 if not found.

### 2. Find First Occurrence

In a sorted slice with duplicates, find the **first** (leftmost) occurrence of target.
- `[1, 2, 2, 2, 3]`, target=2 → index 1

### 3. Find Last Occurrence

In a sorted slice with duplicates, find the **last** (rightmost) occurrence of target.
- `[1, 2, 2, 2, 3]`, target=2 → index 3

### 4. Search Insert Position

Find the index where target would be inserted to maintain sorted order.
- `[1, 3, 5, 7]`, target=4 → index 2 (between 3 and 5)

---

## Function Signatures

```go
func BinarySearch(nums []int, target int) int

func FindFirst(nums []int, target int) int

func FindLast(nums []int, target int) int

func SearchInsertPosition(nums []int, target int) int
```

---

## Examples

### BinarySearch
```go
BinarySearch([]int{1, 2, 3, 4, 5}, 3)      // 2
BinarySearch([]int{1, 2, 3, 4, 5}, 5)      // 4
BinarySearch([]int{1, 2, 3, 4, 5}, 6)      // -1 (not found)
BinarySearch([]int{1, 3, 5, 7, 9}, 1)      // 0
BinarySearch([]int{}, 1)                   // -1 (empty)
```

### FindFirst
```go
FindFirst([]int{1, 2, 2, 2, 3}, 2)         // 1
FindFirst([]int{1, 1, 1, 1}, 1)            // 0
FindFirst([]int{1, 2, 3}, 2)               // 1 (single occurrence)
FindFirst([]int{1, 2, 3}, 4)               // -1 (not found)
```

### FindLast
```go
FindLast([]int{1, 2, 2, 2, 3}, 2)          // 3
FindLast([]int{1, 1, 1, 1}, 1)             // 3
FindLast([]int{1, 2, 3}, 2)                // 1 (single occurrence)
FindLast([]int{1, 2, 3}, 4)                // -1 (not found)
```

### SearchInsertPosition
```go
SearchInsertPosition([]int{1, 3, 5, 7}, 4) // 2 (insert between 3 and 5)
SearchInsertPosition([]int{1, 3, 5, 7}, 0) // 0 (insert at beginning)
SearchInsertPosition([]int{1, 3, 5, 7}, 8) // 4 (insert at end)
SearchInsertPosition([]int{1, 3, 5, 7}, 5) // 2 (already exists)
```

---

## Instructions

1. **Understand binary search**:
   ```
   Search for 7 in [1, 3, 5, 7, 9, 11, 13]:

   left=0, right=6, mid=3 → arr[3]=7 → FOUND!
   ```

   ```
   Search for 6 in [1, 3, 5, 7, 9, 11, 13]:

   left=0, right=6, mid=3 → arr[3]=7 > 6 → search left half
     left=0, right=2, mid=1 → arr[1]=3 < 6 → search right half
       left=2, right=2, mid=2 → arr[2]=5 < 6 → search right half
         left=3, right=2 → left > right → NOT FOUND
   ```

2. **Implement BinarySearch** with helper function:
   ```go
   func BinarySearch(nums []int, target int) int {
       return binarySearchHelper(nums, target, 0, len(nums)-1)
   }

   func binarySearchHelper(nums []int, target, left, right int) int {
       // Base case: search space exhausted
       // Recursive case: compare middle, recurse on half
   }
   ```

3. **Key decisions**:
   - If `arr[mid] == target` → found!
   - If `target < arr[mid]` → search left half (right = mid-1)
   - If `target > arr[mid]` → search right half (left = mid+1)

4. **Implement FindFirst**:
   - Even when you find target, keep searching left to find first occurrence
   - Track the best result found so far

5. **Implement FindLast**:
   - Even when you find target, keep searching right to find last occurrence

6. **Implement SearchInsertPosition**:
   - Similar to binary search, but return the position where target should be

7. **Run tests**: `go test -v`

---

## Hints

<details>
<summary><strong>Hint 1 - Basic (Binary Search Pattern)</strong></summary>

**Basic structure:**
```go
func binarySearchHelper(nums []int, target, left, right int) int {
    // Base case: search space is empty
    if left > right {
        return -1  // Not found
    }

    // Find middle
    mid := left + (right-left)/2  // Avoid overflow

    // Three cases:
    if nums[mid] == target {
        return mid  // Found!
    } else if target < nums[mid] {
        // Search left half
        return binarySearchHelper(nums, target, left, mid-1)
    } else {
        // Search right half
        return binarySearchHelper(nums, target, mid+1, right)
    }
}
```

**Why `left + (right-left)/2` instead of `(left+right)/2`?**
- Avoids integer overflow when left and right are large
- Mathematically equivalent but safer

</details>

<details>
<summary><strong>Hint 2 - Intermediate (Finding First/Last)</strong></summary>

**FindFirst pattern:**
```go
func findFirstHelper(nums []int, target, left, right int) int {
    if left > right {
        return -1
    }

    mid := left + (right-left)/2

    if nums[mid] == target {
        // Found, but check if there's an earlier occurrence
        if mid == 0 || nums[mid-1] != target {
            return mid  // This is the first!
        }
        // Keep searching left
        return findFirstHelper(nums, target, left, mid-1)
    } else if target < nums[mid] {
        return findFirstHelper(nums, target, left, mid-1)
    } else {
        return findFirstHelper(nums, target, mid+1, right)
    }
}
```

**FindLast pattern:** Mirror of FindFirst - search right when found.

</details>

<details>
<summary><strong>Hint 3 - Advanced (SearchInsertPosition)</strong></summary>

**Insert position logic:**
- If target found: return that index
- If target not found: return where it should be inserted

```go
func searchInsertHelper(nums []int, target, left, right int) int {
    if left > right {
        // Not found - left is the insert position
        return left
    }

    mid := left + (right-left)/2

    if nums[mid] == target {
        return mid
    } else if target < nums[mid] {
        return searchInsertHelper(nums, target, left, mid-1)
    } else {
        return searchInsertHelper(nums, target, mid+1, right)
    }
}
```

**Key insight:** When `left > right`, `left` points to where the element should be inserted!

</details>

<details>
<summary><strong>Hint 4 - Complete Solutions</strong></summary>

**BinarySearch:**
```go
func BinarySearch(nums []int, target int) int {
    return binarySearchHelper(nums, target, 0, len(nums)-1)
}

func binarySearchHelper(nums []int, target, left, right int) int {
    if left > right {
        return -1
    }

    mid := left + (right-left)/2

    if nums[mid] == target {
        return mid
    } else if target < nums[mid] {
        return binarySearchHelper(nums, target, left, mid-1)
    } else {
        return binarySearchHelper(nums, target, mid+1, right)
    }
}
```

**FindFirst:**
```go
func FindFirst(nums []int, target int) int {
    return findFirstHelper(nums, target, 0, len(nums)-1)
}

func findFirstHelper(nums []int, target, left, right int) int {
    if left > right {
        return -1
    }

    mid := left + (right-left)/2

    if nums[mid] == target {
        if mid == 0 || nums[mid-1] != target {
            return mid
        }
        return findFirstHelper(nums, target, left, mid-1)
    } else if target < nums[mid] {
        return findFirstHelper(nums, target, left, mid-1)
    } else {
        return findFirstHelper(nums, target, mid+1, right)
    }
}
```

**FindLast:**
```go
func FindLast(nums []int, target int) int {
    return findLastHelper(nums, target, 0, len(nums)-1)
}

func findLastHelper(nums []int, target, left, right int) int {
    if left > right {
        return -1
    }

    mid := left + (right-left)/2

    if nums[mid] == target {
        if mid == len(nums)-1 || nums[mid+1] != target {
            return mid
        }
        return findLastHelper(nums, target, mid+1, right)
    } else if target < nums[mid] {
        return findLastHelper(nums, target, left, mid-1)
    } else {
        return findLastHelper(nums, target, mid+1, right)
    }
}
```

Wait, we need to pass `nums` length info! Better approach:
```go
func findLastHelper(nums []int, target, left, right int) int {
    if left > right {
        return -1
    }

    mid := left + (right-left)/2

    if nums[mid] == target {
        if mid == right || nums[mid+1] != target {
            return mid
        }
        return findLastHelper(nums, target, mid+1, right)
    } else if target < nums[mid] {
        return findLastHelper(nums, target, left, mid-1)
    } else {
        return findLastHelper(nums, target, mid+1, right)
    }
}
```

**SearchInsertPosition:**
```go
func SearchInsertPosition(nums []int, target int) int {
    return searchInsertHelper(nums, target, 0, len(nums)-1)
}

func searchInsertHelper(nums []int, target, left, right int) int {
    if left > right {
        return left
    }

    mid := left + (right-left)/2

    if nums[mid] == target {
        return mid
    } else if target < nums[mid] {
        return searchInsertHelper(nums, target, left, mid-1)
    } else {
        return searchInsertHelper(nums, target, mid+1, right)
    }
}
```

</details>

---

## Think About

1. **Time complexity:** Binary search is O(log n). Why?
   - Each step halves the search space
   - log₂(n) steps to reduce n elements to 1

2. **Space complexity:** Recursive binary search uses O(log n) stack space. The iterative version uses O(1). Which is better?

3. **Why sorted?** Binary search only works on sorted data. Why is this requirement necessary?

4. **Overflow bug:** `mid = (left + right) / 2` can overflow for large arrays. Why does `left + (right-left)/2` fix this?

5. **Off-by-one errors:** Binary search is notorious for boundary bugs. Why is it so easy to get `mid-1` vs `mid+1` vs `mid` wrong?

---

## What This Teaches

✅ **Divide-and-conquer** - Halving the problem at each step
✅ **Logarithmic complexity** - Understanding O(log n)
✅ **Binary search pattern** - The foundation for many algorithms
✅ **Boundary handling** - left, right, mid management
✅ **Search variants** - First, last, insert position
✅ **Helper function with bounds** - Passing left/right indices

**Real-world applications:**
- Database indexing
- Dictionary lookups
- Version control (git bisect)
- Finding bugs in releases
- Any sorted data search

---

**Next Exercise:** 09 - Merge Sort (divide-recurse-combine pattern)
