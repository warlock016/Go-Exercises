# Exercise 08: Sliding Window Maximum Sum

**Difficulty:** Medium
**Time:** 30-40 minutes
**Concepts:** Sliding window technique, subarray processing, optimization

## Learning Goal

Master the sliding window pattern - a powerful technique for efficiently processing contiguous subarrays. Instead of recalculating sums from scratch for each window position, you maintain a running sum by adding the new element and removing the old element.

## The Problem

Given an array of integers and a window size `k`, find the maximum sum of any contiguous subarray of length `k`.

For example, with `nums = [2, 1, 5, 1, 3, 2]` and `k = 3`:
- Window [2, 1, 5] → sum = 8
- Window [1, 5, 1] → sum = 7
- Window [5, 1, 3] → sum = 9 ← **maximum**
- Window [1, 3, 2] → sum = 6

Return `9`.

## Why Sliding Window?

**Naive approach (inefficient):**
```go
// Recalculate sum for each window - O(n*k) time
for i := 0; i <= len(nums)-k; i++ {
    sum := 0
    for j := i; j < i+k; j++ {
        sum += nums[j]
    }
    maxSum = max(maxSum, sum)
}
```

**Sliding window (efficient):**
```go
// Maintain running sum - O(n) time
sum := sum of first k elements
maxSum := sum
for i := k; i < len(nums); i++ {
    sum = sum + nums[i] - nums[i-k]  // Slide the window
    maxSum = max(maxSum, sum)
}
```

## Function Signature

```go
func MaxSumSubarray(nums []int, k int) int
```

## Examples

**Example 1:**
```go
nums := []int{2, 1, 5, 1, 3, 2}
k := 3
MaxSumSubarray(nums, k)  // → 9
// Windows: [2,1,5]=8, [1,5,1]=7, [5,1,3]=9, [1,3,2]=6
```

**Example 2:**
```go
nums := []int{-1, -2, -3, -4}
k := 2
MaxSumSubarray(nums, k)  // → -3
// Windows: [-1,-2]=-3, [-2,-3]=-5, [-3,-4]=-7
// Maximum is -3 (least negative)
```

**Example 3:**
```go
nums := []int{10, 20, 30, 40, 50}
k := 1
MaxSumSubarray(nums, k)  // → 50
// Each element is its own window
```

**Example 4:**
```go
nums := []int{5, 4, 3, 2, 1}
k := 5
MaxSumSubarray(nums, k)  // → 15
// Only one window: [5,4,3,2,1]=15
```

**Example 5:**
```go
nums := []int{100}
k := 1
MaxSumSubarray(nums, k)  // → 100
```

## Instructions

1. **Handle edge cases:**
   - If `k` is 0 or `nums` is empty, return 0
   - If `k` equals `len(nums)`, return sum of entire array
   - If `k` is 1, return the maximum element

2. **Build the first window:**
   - Calculate the sum of the first `k` elements
   - This is your initial `maxSum`

3. **Slide the window:**
   - For each position from `k` to `len(nums)-1`:
     - Add the new element entering the window (nums[i])
     - Subtract the element leaving the window (nums[i-k])
     - Update `maxSum` if current window sum is larger

4. **Return the maximum sum found**

## Hints

<details>
<summary>Hint 1: Edge cases (Click to reveal)</summary>

```go
if k <= 0 || len(nums) == 0 || k > len(nums) {
    // Handle invalid input
    if len(nums) == 0 || k <= 0 {
        return 0
    }
}
```
</details>

<details>
<summary>Hint 2: First window setup (Click to reveal)</summary>

```go
// Calculate sum of first window
windowSum := 0
for i := 0; i < k; i++ {
    windowSum += nums[i]
}
maxSum := windowSum  // Initialize maximum
```
</details>

<details>
<summary>Hint 3: Sliding the window (Click to reveal)</summary>

```go
// Slide window from position k to end
for i := k; i < len(nums); i++ {
    // Add new element, remove old element
    windowSum = windowSum + nums[i] - nums[i-k]

    // Update maximum
    if windowSum > maxSum {
        maxSum = windowSum
    }
}
```
</details>

<details>
<summary>Hint 4: Complete structure (Click to reveal)</summary>

```go
func MaxSumSubarray(nums []int, k int) int {
    // 1. Edge cases
    if len(nums) == 0 || k <= 0 || k > len(nums) {
        return 0
    }

    // 2. Calculate first window
    windowSum := 0
    for i := 0; i < k; i++ {
        windowSum += nums[i]
    }
    maxSum := windowSum

    // 3. Slide window
    for i := k; i < len(nums); i++ {
        windowSum = windowSum + nums[i] - nums[i-k]
        if windowSum > maxSum {
            maxSum = windowSum
        }
    }

    // 4. Return maximum
    return maxSum
}
```
</details>

## Think About

1. **Why is this O(n) instead of O(n*k)?**
   - Each element is added once and removed once
   - No nested loops recalculating sums

2. **What if k > len(nums)?**
   - Invalid input - return 0 or sum of entire array?
   - Design decision - be consistent

3. **Can you use this pattern for other problems?**
   - Minimum sum subarray
   - Average of subarrays
   - Products instead of sums

4. **What breaks if you forget `- nums[i-k]`?**
   - You'd be calculating cumulative sum, not window sum
   - All windows would include all previous elements

## What This Teaches

- **Sliding window pattern:** Efficient subarray processing
- **Optimization:** Avoiding redundant recalculation
- **Running state:** Maintaining sum across iterations
- **Index arithmetic:** `i-k` to find element leaving window
- **Time complexity:** O(n*k) → O(n) improvement

Ready to implement? Open `sliding_window.go` and look for `TODO(human)` markers!
