# Exercise 03: Slice Recursion

**Tier:** 1 - Foundation
**Estimated Time:** 50 minutes
**Concepts:** Recursion on collections, slice operations, reducing problem size

---

## Learning Goal

Apply recursion to collections (slices) by progressively reducing the problem. Master the pattern: `process first element + recurse on rest`.

---

## Problem Description

Implement recursive functions that operate on slices using the pattern:
```
Result = f(first element) + RecursiveCall(remaining elements)
```

### 1. Sum of Slice

Calculate the sum of all integers in a slice.

### 2. Maximum in Slice

Find the maximum value in a non-empty slice.

### 3. Contains

Check if a slice contains a specific value.

### 4. Count Occurrences

Count how many times a value appears in a slice.

---

## Function Signatures

```go
func Sum(nums []int) int

func Max(nums []int) int

func Contains(nums []int, target int) bool

func CountOccurrences(nums []int, target int) int
```

---

## Examples

### Sum
```go
Sum([]int{})           // 0 (empty slice)
Sum([]int{5})          // 5
Sum([]int{1, 2, 3})    // 6
Sum([]int{10, -5, 3})  // 8
```

### Max
```go
Max([]int{5})             // 5
Max([]int{1, 9, 3})       // 9
Max([]int{-5, -2, -10})   // -2
Max([]int{100, 50, 75})   // 100
```

### Contains
```go
Contains([]int{}, 5)              // false
Contains([]int{1, 2, 3}, 2)       // true
Contains([]int{1, 2, 3}, 5)       // false
Contains([]int{10, 20, 30}, 20)   // true
```

### CountOccurrences
```go
CountOccurrences([]int{}, 5)                    // 0
CountOccurrences([]int{1, 2, 3}, 5)             // 0
CountOccurrences([]int{1, 2, 1, 3, 1}, 1)       // 3
CountOccurrences([]int{5, 5, 5, 5}, 5)          // 4
```

---

## Instructions

1. **Understand the slice recursion pattern**:
   ```go
   func ProcessSlice(nums []int) result {
       if len(nums) == 0 {
           return baseValue  // Base case: empty slice
       }
       first := nums[0]      // Process first element
       rest := nums[1:]      // Get remaining elements
       return combine(first, ProcessSlice(rest))  // Recursive case
   }
   ```

2. **Trace Sum([1, 2, 3]) by hand**:
   ```
   Sum([1, 2, 3])
     → 1 + Sum([2, 3])
           → 2 + Sum([3])
                 → 3 + Sum([])
                       → 0 (base case)
                 → 3 + 0 = 3
           → 2 + 3 = 5
     → 1 + 5 = 6
   ```

3. **Implement Sum**:
   - Base case: empty slice → 0
   - Recursive case: first + Sum(rest)

4. **Implement Max**:
   - Base case: single element → that element
   - Recursive case: max(first, Max(rest))
   - Hint: Use Go's `max()` built-in or write: `if a > b { return a } else { return b }`

5. **Implement Contains**:
   - Base case: empty slice → false
   - Check first element: if it matches, return true
   - Recursive case: Contains(rest)

6. **Implement CountOccurrences**:
   - Base case: empty slice → 0
   - If first matches: 1 + CountOccurrences(rest)
   - If first doesn't match: 0 + CountOccurrences(rest)

7. **Run tests**: `go test -v`

---

## Hints

<details>
<summary><strong>Hint 1 - Basic (Base Cases for Slices)</strong></summary>

**Common pattern for slice recursion:**
- Base case is usually `len(nums) == 0` or `len(nums) == 1`
- Empty slice is often the termination condition

**Sum:** empty slice → 0
**Max:** single element → that element (can't have max of empty slice)
**Contains:** empty slice → false (can't contain anything)
**CountOccurrences:** empty slice → 0

</details>

<details>
<summary><strong>Hint 2 - Intermediate (Making Progress)</strong></summary>

**The key pattern:**
```go
if len(nums) == 0 {
    return baseCase
}

first := nums[0]
rest := nums[1:]  // This is how you reduce the problem!

// Combine first with recursive result
return combine(first, RecursiveFunction(rest))
```

**Important:** `nums[1:]` creates a new slice with all elements except the first.
- This is how you make progress toward the base case
- Each recursive call has a smaller slice

</details>

<details>
<summary><strong>Hint 3 - Advanced (Combining Results)</strong></summary>

**How to combine first element with recursive result:**

**Sum:** Addition
```go
return first + Sum(rest)
```

**Max:** Take maximum
```go
maxOfRest := Max(rest)
if first > maxOfRest {
    return first
}
return maxOfRest
```

**Contains:** Logical OR
```go
if first == target {
    return true
}
return Contains(rest, target)
```

**CountOccurrences:** Conditional addition
```go
count := CountOccurrences(rest, target)
if first == target {
    return 1 + count
}
return count
```

</details>

<details>
<summary><strong>Hint 4 - Complete Solutions</strong></summary>

**Sum:**
```go
func Sum(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    return nums[0] + Sum(nums[1:])
}
```

**Max:**
```go
func Max(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    maxOfRest := Max(nums[1:])
    if nums[0] > maxOfRest {
        return nums[0]
    }
    return maxOfRest
}
```

**Contains:**
```go
func Contains(nums []int, target int) bool {
    if len(nums) == 0 {
        return false
    }
    if nums[0] == target {
        return true
    }
    return Contains(nums[1:], target)
}
```

**CountOccurrences:**
```go
func CountOccurrences(nums []int, target int) int {
    if len(nums) == 0 {
        return 0
    }
    count := CountOccurrences(nums[1:], target)
    if nums[0] == target {
        return 1 + count
    }
    return count
}
```

</details>

---

## Think About

1. **Slice copying performance:** Each `nums[1:]` creates a new slice. For a slice of length n, how many slice operations occur? Is this efficient?

2. **Iterative vs recursive:** Would a `for` loop be more efficient for these problems? When would you choose recursion over iteration?

3. **Stack depth:** For `Sum()` on a slice of 10,000 elements, how deep is the call stack? Could this cause a stack overflow?

4. **Alternative pattern:** Instead of `nums[1:]`, could you pass an index parameter? How would that change the implementation?
   ```go
   func SumHelper(nums []int, index int) int {
       if index >= len(nums) {
           return 0
       }
       return nums[index] + SumHelper(nums, index+1)
   }
   ```

5. **Tail recursion:** The helper pattern above is "tail recursive" (recursive call is the last operation). Why might this be better?

---

## What This Teaches

✅ **Collection recursion pattern** - `arr → f(arr[1:])`
✅ **Slice slicing syntax** - `nums[1:]` to get "rest"
✅ **Divide and conquer mindset** - Split into first + rest
✅ **Combining results** - Different strategies (add, max, OR, etc.)
✅ **Base case on empty** - `len(nums) == 0` pattern
✅ **Performance awareness** - Understanding recursion costs

**Pattern Recognition:** This `first + recurse(rest)` pattern appears in linked lists, trees, and many functional programming patterns.

---

**Next Exercise:** 04 - String Recursion (applying these patterns to strings)
