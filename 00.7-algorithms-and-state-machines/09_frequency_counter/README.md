# Exercise 09: Top K Frequent Elements

**Difficulty:** Medium
**Time:** 40-50 minutes
**Concepts:** Hash maps, frequency counting, sorting, custom comparators

## Learning Goal

Master the frequency counter pattern - using a map to count occurrences, then processing those counts to extract meaningful information. This is a fundamental technique in data analysis, statistics, and algorithm optimization.

## The Problem

Given an integer array `nums` and an integer `k`, return the `k` most frequent elements. The answer can be returned in any order.

For example, with `nums = [1, 1, 1, 2, 2, 3]` and `k = 2`:
- Frequency: {1: 3, 2: 2, 3: 1}
- Top 2 frequent: [1, 2] (or [2, 1])

## Real-World Applications

- **Analytics:** Find top N products sold
- **Text processing:** Most common words in a document
- **Network monitoring:** Most active IP addresses
- **Recommendation systems:** Trending items
- **Debugging:** Most frequent errors in logs

## Function Signature

```go
func TopKFrequent(nums []int, k int) []int
```

## Examples

**Example 1:**
```go
nums := []int{1, 1, 1, 2, 2, 3}
k := 2
TopKFrequent(nums, k)  // → [1, 2] (or [2, 1])
// Frequencies: 1 appears 3 times, 2 appears 2 times, 3 appears 1 time
```

**Example 2:**
```go
nums := []int{1}
k := 1
TopKFrequent(nums, k)  // → [1]
```

**Example 3:**
```go
nums := []int{4, 1, -1, 2, -1, 2, 3}
k := 2
TopKFrequent(nums, k)  // → [-1, 2] (or [2, -1])
// Frequencies: -1 appears 2 times, 2 appears 2 times (tied for top)
```

**Example 4:**
```go
nums := []int{5, 5, 5, 5, 5}
k := 1
TopKFrequent(nums, k)  // → [5]
```

**Example 5:**
```go
nums := []int{1, 2, 3, 4, 5}
k := 3
TopKFrequent(nums, k)  // → [1, 2, 3] (or any 3 elements - all have frequency 1)
```

## Algorithm Overview

**Step 1: Build frequency map**
```go
freq := make(map[int]int)
for _, num := range nums {
    freq[num]++
}
// Result: {1: 3, 2: 2, 3: 1}
```

**Step 2: Convert map to slice of pairs**
```go
type pair struct {
    num   int
    count int
}

pairs := make([]pair, 0, len(freq))
for num, count := range freq {
    pairs = append(pairs, pair{num, count})
}
// Result: [{1, 3}, {2, 2}, {3, 1}]
```

**Step 3: Sort by frequency (descending)**
```go
sort.Slice(pairs, func(i, j int) bool {
    return pairs[i].count > pairs[j].count
})
// Result: [{1, 3}, {2, 2}, {3, 1}] (already sorted)
```

**Step 4: Extract top k elements**
```go
result := make([]int, k)
for i := 0; i < k; i++ {
    result[i] = pairs[i].num
}
// Result: [1, 2]
```

## Instructions

1. **Handle edge cases:**
   - If `nums` is empty, return empty slice
   - If `k` is 0, return empty slice
   - If `k >= len(unique elements)`, return all unique elements

2. **Build frequency map:**
   - Create `map[int]int` to count occurrences
   - Iterate through `nums` and increment counts

3. **Create sortable structure:**
   - Define a struct to hold (number, frequency) pairs
   - Convert map entries to slice of these pairs

4. **Sort by frequency:**
   - Use `sort.Slice()` with custom comparator
   - Sort descending (highest frequency first)

5. **Extract top k:**
   - Take first k elements from sorted slice
   - Return just the numbers (not frequencies)

## Hints

<details>
<summary>Hint 1: Frequency map (Click to reveal)</summary>

```go
freq := make(map[int]int)
for _, num := range nums {
    freq[num]++
}

// Example: [1, 1, 1, 2, 2, 3]
// freq = {1: 3, 2: 2, 3: 1}
```
</details>

<details>
<summary>Hint 2: Pair struct for sorting (Click to reveal)</summary>

```go
type pair struct {
    num   int  // The number
    count int  // How many times it appears
}

// Convert map to slice
pairs := make([]pair, 0, len(freq))
for num, count := range freq {
    pairs = append(pairs, pair{num, count})
}
```
</details>

<details>
<summary>Hint 3: Sorting (Click to reveal)</summary>

```go
import "sort"

// Sort pairs by count (descending)
sort.Slice(pairs, func(i, j int) bool {
    return pairs[i].count > pairs[j].count  // > for descending
})

// Now pairs[0] has highest frequency
```
</details>

<details>
<summary>Hint 4: Extracting top k (Click to reveal)</summary>

```go
// Handle k larger than available unique elements
if k > len(pairs) {
    k = len(pairs)
}

result := make([]int, k)
for i := 0; i < k; i++ {
    result[i] = pairs[i].num
}
return result
```
</details>

<details>
<summary>Hint 5: Complete structure (Click to reveal)</summary>

```go
import "sort"

func TopKFrequent(nums []int, k int) []int {
    // 1. Edge cases
    if len(nums) == 0 || k == 0 {
        return []int{}
    }

    // 2. Build frequency map
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
    }

    // 3. Convert to sortable slice
    type pair struct {
        num   int
        count int
    }
    pairs := make([]pair, 0, len(freq))
    for num, count := range freq {
        pairs = append(pairs, pair{num, count})
    }

    // 4. Sort by frequency (descending)
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].count > pairs[j].count
    })

    // 5. Extract top k
    if k > len(pairs) {
        k = len(pairs)
    }
    result := make([]int, k)
    for i := 0; i < k; i++ {
        result[i] = pairs[i].num
    }

    return result
}
```
</details>

## Think About

1. **What's the time complexity?**
   - Building map: O(n)
   - Converting to slice: O(u) where u = unique elements
   - Sorting: O(u log u)
   - Extracting k: O(k)
   - **Total: O(n + u log u)**

2. **What's the space complexity?**
   - Frequency map: O(u)
   - Pairs slice: O(u)
   - Result: O(k)
   - **Total: O(u)**

3. **Can you optimize further?**
   - Use a heap (priority queue) to get O(n + u log k) time
   - Bucket sort if frequencies are bounded: O(n) time
   - Trade-offs between simplicity and performance

4. **What if there are ties in frequency?**
   - Current solution: arbitrary order among ties
   - Could add secondary sort (by value, alphabetically, etc.)

5. **Why not just sort the original array?**
   - Sorting gives you most/least values, not most frequent
   - Need to count occurrences first

## What This Teaches

- **Frequency counting:** Using maps to aggregate data
- **Map iteration:** Converting maps to slices for sorting
- **Custom sorting:** Using `sort.Slice()` with comparator functions
- **Struct design:** Creating helper types for intermediate data
- **Algorithm composition:** Combining multiple techniques (map + sort + slice)
- **Edge case handling:** k bounds, empty inputs

Ready to implement? Open `frequency_counter.go` and look for `TODO(human)` markers!
