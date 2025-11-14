# Algorithmic Patterns Reference Guide

## Introduction

### Why Pattern Recognition Matters

Programming is fundamentally about recognizing patterns. When you encounter a problem, the challenge isn't usually the syntax of your language—it's identifying **which algorithmic pattern** applies. Once you recognize the pattern, implementation becomes straightforward.

This guide focuses on building your **mental models** for common algorithmic patterns. These are the fundamental building blocks that appear repeatedly across different problems, languages, and domains.

### How This Differs from Go Syntax Knowledge

- **Syntax knowledge**: "How do I iterate over a map in Go?"
- **Pattern recognition**: "This problem requires tracking the maximum value while iterating—I need the 'find max in collection' pattern"

You can know Go syntax perfectly but still struggle if you don't recognize which pattern solves your problem. This guide bridges that gap.

### How to Use This Guide

1. **Read through each pattern** to understand its mental model
2. **Study the "When to Use It" section** to recognize problem characteristics
3. **Understand the pseudocode** before looking at the Go implementation
4. **Try the practice problem** at the end of each pattern
5. **Test yourself** with the Pattern Recognition Practice section

Don't memorize code. Instead, internalize the **thinking process** behind each pattern.

---

## Core Patterns

### Pattern 1: Find Maximum (or Minimum) in a Collection

**When to Use It:**
- You need to find the "best", "largest", "smallest", or "most common" element
- You're comparing elements and tracking the winner
- Problem contains words like: "maximum", "minimum", "highest", "lowest", "most frequent"

**Mental Model:**

The key insight: **You don't need to see all items before deciding**. As you iterate, maintain a "current best" candidate. Compare each new item to your current best, and update if the new item wins.

**Generic Pseudocode:**
```
initialize best_candidate (often to first element or sentinel value)
initialize best_score (if tracking a metric)

for each item in collection:
    calculate score for current item
    if current_score > best_score:
        best_candidate = current item
        best_score = current_score

return best_candidate
```

**Go Implementation Example:**
```go
// Find the most common character in a string
func MostCommon(s string) rune {
    // First, build frequency map
    freq := make(map[rune]int)
    for _, char := range s {
        freq[char]++
    }

    // Find max in the map
    var mostCommon rune
    maxCount := 0

    for char, count := range freq {
        if count > maxCount {
            maxCount = count
            mostCommon = char
        }
    }

    return mostCommon
}
```

**Common Variations:**
- Find minimum: change `>` to `<`
- Find multiple items: track a slice of candidates
- Find last occurrence of max: change `>` to `>=`
- Early exit: return immediately when condition met (if searching)

**Practice Problem:**

Write a function that takes a slice of integers and returns the largest even number. If no even numbers exist, return -1.

---

### Pattern 2: Frequency Counting / Histogram

**When to Use It:**
- You need to count occurrences of items
- Problem asks "how many times", "frequency", "histogram", "distribution"
- You need to group items by category
- Building statistics or summaries

**Mental Model:**

Think of a **tally system**. Each time you see an item, make a mark next to it. In programming, a map serves as your tally sheet: keys are items, values are counts.

**Generic Pseudocode:**
```
create empty frequency_map

for each item in collection:
    if item exists in frequency_map:
        increment its count
    else:
        add item to frequency_map with count = 1

return frequency_map
```

**Go Implementation Example:**
```go
// Count frequency of words in a sentence
func WordFrequency(sentence string) map[string]int {
    freq := make(map[string]int)
    words := strings.Fields(sentence)

    for _, word := range words {
        freq[word]++  // Go's zero value makes this elegant
    }

    return freq
}
```

**Common Variations:**
- Count unique items: return `len(frequency_map)`
- Filter by threshold: only keep items with count > N
- Most/least common N items: combine with sorting or heap
- Group by category: use map[category][]items instead of counts

**Practice Problem:**

Write a function that takes a slice of strings and returns a map showing how many strings have each length. For example, `["hi", "go", "hello"]` would return `map[int]int{2: 2, 5: 1}`.

---

### Pattern 3: Two-Pointer Technique

**When to Use It:**
- Working with sorted data
- Need to find pairs that meet a condition
- Comparing elements from different positions
- Detecting patterns like palindromes
- Problems involving "two elements that sum to X"

**Mental Model:**

Imagine **two fingers pointing at different positions** in your data. Move them according to what you find, often from opposite ends moving inward, or both from the same end moving at different speeds.

**Generic Pseudocode:**
```
initialize left pointer to start
initialize right pointer to end (or wherever appropriate)

while pointers haven't crossed:
    examine element(s) at pointer position(s)

    if condition met:
        record result
        move pointer(s)
    else if need to search higher:
        move left pointer forward
    else:
        move right pointer backward

return result
```

**Go Implementation Example:**
```go
// Check if string is a palindrome
func IsPalindrome(s string) bool {
    left := 0
    right := len(s) - 1

    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }

    return true
}

// Find pair in sorted slice that sums to target
func FindPairSum(nums []int, target int) (int, int, bool) {
    left := 0
    right := len(nums) - 1

    for left < right {
        sum := nums[left] + nums[right]

        if sum == target {
            return nums[left], nums[right], true
        } else if sum < target {
            left++  // need larger sum
        } else {
            right--  // need smaller sum
        }
    }

    return 0, 0, false
}
```

**Common Variations:**
- Fast/slow pointers: detect cycles (tortoise and hare)
- Same-direction pointers: sliding window (see next pattern)
- Multiple pointers: merging sorted arrays
- Reverse two-pointer: checking for specific patterns

**Practice Problem:**

Write a function that takes a slice of integers and removes duplicates in-place, returning the new length. The slice should be modified so all unique elements are at the beginning.

---

### Pattern 4: Sliding Window

**When to Use It:**
- Finding subarrays or substrings that meet criteria
- Problems mention "contiguous", "consecutive", "substring"
- Maximum/minimum sum of K consecutive elements
- Longest/shortest substring with property X
- Moving average or rolling calculations

**Mental Model:**

Imagine a **fixed or flexible window** sliding across your data. Instead of recalculating everything for each position, you efficiently update by removing the element leaving the window and adding the element entering it.

**Generic Pseudocode:**
```
// Fixed-size window
initialize window with first k elements
calculate initial result

for each position from k to end:
    remove leftmost element from window
    add new rightmost element to window
    update result
    track best result if needed

return result

// Variable-size window
left = 0
for right from 0 to end:
    add element at right to window

    while window violates condition:
        remove element at left from window
        move left forward

    update best result
```

**Go Implementation Example:**
```go
// Maximum sum of k consecutive elements
func MaxSumSubarray(nums []int, k int) int {
    if len(nums) < k {
        return 0
    }

    // Calculate first window
    windowSum := 0
    for i := 0; i < k; i++ {
        windowSum += nums[i]
    }
    maxSum := windowSum

    // Slide the window
    for i := k; i < len(nums); i++ {
        windowSum = windowSum - nums[i-k] + nums[i]
        if windowSum > maxSum {
            maxSum = windowSum
        }
    }

    return maxSum
}

// Longest substring without repeating characters
func LongestUniqueSubstring(s string) int {
    seen := make(map[rune]int)
    maxLen := 0
    left := 0

    for right, char := range s {
        if prevIndex, exists := seen[char]; exists && prevIndex >= left {
            left = prevIndex + 1
        }

        seen[char] = right
        currentLen := right - left + 1
        if currentLen > maxLen {
            maxLen = currentLen
        }
    }

    return maxLen
}
```

**Common Variations:**
- Fixed size: simpler, just slide one position at a time
- Variable size: expand until invalid, then contract
- Multiple conditions: track multiple metrics for window validity
- Minimum window: find smallest window that satisfies condition

**Practice Problem:**

Write a function that finds the longest substring with at most K distinct characters. For example, in "aabbcc" with K=2, the answer is 4 ("aabb" or "bbcc").

---

### Pattern 5: State Tracking During Iteration

**When to Use It:**
- Processing sequential data where previous elements affect current decisions
- Parsing or validation problems
- State machines or mode-based processing
- Problems with context-dependent behavior
- Tracking "have I seen X yet?" conditions

**Mental Model:**

As you iterate, you're not just looking at each element in isolation—you're **maintaining context** about what you've encountered. Think of it like reading a book: each page makes sense because you remember what came before.

**Generic Pseudocode:**
```
initialize state variables

for each element:
    examine current element
    consider current state

    decide action based on element + state
    update state for next iteration
    accumulate results if needed

return final result based on state
```

**Go Implementation Example:**
```go
// Validate that parentheses are balanced
func IsBalanced(s string) bool {
    openCount := 0

    for _, char := range s {
        if char == '(' {
            openCount++
        } else if char == ')' {
            openCount--
            if openCount < 0 {
                return false  // closed before opened
            }
        }
    }

    return openCount == 0
}

// Find longest consecutive sequence of 1s
func LongestOnesSequence(nums []int) int {
    maxLen := 0
    currentLen := 0

    for _, num := range nums {
        if num == 1 {
            currentLen++
            if currentLen > maxLen {
                maxLen = currentLen
            }
        } else {
            currentLen = 0  // reset state
        }
    }

    return maxLen
}
```

**Common Variations:**
- Boolean flags: track "have seen X", "inside quote", etc.
- Counter states: track nesting depth, streak length
- Previous element tracking: compare with last seen value
- Multiple state variables: complex state machines
- Stack-based state: for nested structures

**Practice Problem:**

Write a function that determines if a string has alternating vowels and consonants. For example, "banana" follows this pattern but "apple" does not. Consider only letters, ignoring other characters.

---

### Pattern 6: Accumulator Pattern

**When to Use It:**
- Building up a result incrementally
- Calculating running totals, products, or aggregations
- Combining or transforming multiple items into one result
- Problems asking for "sum", "product", "concatenation", "total"

**Mental Model:**

Think of an **empty bucket** that you fill as you go. Each iteration adds something to your accumulator. The final accumulator value is your answer.

**Generic Pseudocode:**
```
initialize accumulator to appropriate starting value
    (0 for sum, 1 for product, empty for concatenation)

for each element:
    combine element with accumulator
    (add, multiply, append, etc.)

return accumulator
```

**Go Implementation Example:**
```go
// Sum all elements in a slice
func Sum(nums []int) int {
    total := 0  // accumulator starts at 0

    for _, num := range nums {
        total += num
    }

    return total
}

// Product of all non-zero elements
func ProductNonZero(nums []int) int {
    product := 1  // accumulator starts at 1 for multiplication

    for _, num := range nums {
        if num != 0 {
            product *= num
        }
    }

    return product
}

// Concatenate strings with separator
func JoinStrings(words []string, sep string) string {
    if len(words) == 0 {
        return ""
    }

    result := words[0]

    for i := 1; i < len(words); i++ {
        result += sep + words[i]
    }

    return result
}
```

**Common Variations:**
- Conditional accumulation: only add if condition met
- Multiple accumulators: track several aggregations simultaneously
- Nested accumulation: accumulate within accumulate
- Transform while accumulating: apply function before adding

**Practice Problem:**

Write a function that calculates the sum of squares of all even numbers in a slice of integers.

---

### Pattern 7: Filter and Transform

**When to Use It:**
- Creating a new collection based on an existing one
- Problems ask to "extract", "select", "filter", "convert", "map"
- You need items that match certain criteria
- You need to modify each item in a collection
- Building a subset or modified version of data

**Mental Model:**

You're a **factory assembly line** with two stations: the filter station (which items pass through?) and the transform station (how do we modify them?). Items may go through just filtering, just transformation, or both.

**Generic Pseudocode:**
```
// Filter only
create empty result collection

for each element:
    if element matches criteria:
        add element to result

return result

// Transform only
create empty result collection

for each element:
    transformed = apply transformation to element
    add transformed to result

return result

// Filter then transform
create empty result collection

for each element:
    if element matches criteria:
        transformed = apply transformation to element
        add transformed to result

return result
```

**Go Implementation Example:**
```go
// Filter: keep only even numbers
func FilterEven(nums []int) []int {
    result := []int{}

    for _, num := range nums {
        if num%2 == 0 {
            result = append(result, num)
        }
    }

    return result
}

// Transform: double all numbers
func DoubleAll(nums []int) []int {
    result := make([]int, len(nums))

    for i, num := range nums {
        result[i] = num * 2
    }

    return result
}

// Filter and transform: extract lengths of long words
func LongWordLengths(words []string, minLen int) []int {
    result := []int{}

    for _, word := range words {
        if len(word) >= minLen {
            result = append(result, len(word))
        }
    }

    return result
}
```

**Common Variations:**
- In-place filtering: modify original slice
- Multiple filters: combine conditions with AND/OR
- Chained transformations: apply multiple transforms
- Filter with index: keep elements at certain positions
- Transform based on condition: different transforms for different items

**Practice Problem:**

Write a function that takes a slice of strings and returns a slice containing only the palindromes, converted to uppercase.

---

### Pattern 8: String Building

**When to Use It:**
- Constructing strings iteratively
- Concatenating many pieces together
- Building formatted output
- Problems involving "construct", "build", "format", "generate"
- Any time you're repeatedly adding to a string

**Mental Model:**

Strings in Go are immutable—concatenating creates new strings each time. Think of **strings.Builder as a growing buffer** that efficiently accumulates string pieces without constant reallocation.

**Generic Pseudocode:**
```
create builder/buffer

for each piece to add:
    calculate or determine next piece
    append piece to builder

convert builder to final string
return result
```

**Go Implementation Example:**
```go
// Build a string efficiently
func BuildString(words []string, sep string) string {
    var builder strings.Builder

    for i, word := range words {
        builder.WriteString(word)
        if i < len(words)-1 {
            builder.WriteString(sep)
        }
    }

    return builder.String()
}

// Reverse a string
func ReverseString(s string) string {
    runes := []rune(s)  // handle multi-byte characters correctly

    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    return string(runes)
}

// Repeat pattern N times
func RepeatPattern(pattern string, n int) string {
    var builder strings.Builder
    builder.Grow(len(pattern) * n)  // pre-allocate

    for i := 0; i < n; i++ {
        builder.WriteString(pattern)
    }

    return builder.String()
}
```

**Common Variations:**
- With transformation: modify pieces before adding
- Conditional building: only add certain pieces
- Formatted building: use fmt.Sprintf for complex formatting
- Character-by-character: build from individual characters
- Pre-allocation: use Grow() when final size is known

**Practice Problem:**

Write a function that takes a sentence and returns it with every other word reversed. For example, "hello world today" becomes "hello dlrow today".

---

### Pattern 9: Linear Search Patterns

**When to Use It:**
- Finding if an element exists
- Finding first/last occurrence of something
- Searching unsorted data
- Problems asking "does X exist?", "find first X", "locate Y"
- Validation or existence checks

**Mental Model:**

Think of **looking through a deck of cards** one at a time. You examine each card until you find what you're looking for. Sometimes you need to check every card; sometimes you can stop early.

**Generic Pseudocode:**
```
// Simple existence
for each element:
    if element matches criteria:
        return true (or return element, or return index)

return false (or not found indicator)

// Find all occurrences
create results collection

for each element with index:
    if element matches criteria:
        add element/index to results

return results

// Find first matching
for each element:
    if element matches criteria:
        return element immediately

return not-found indicator
```

**Go Implementation Example:**
```go
// Check if element exists
func Contains(nums []int, target int) bool {
    for _, num := range nums {
        if num == target {
            return true
        }
    }
    return false
}

// Find first index where condition is true
func FindFirst(nums []int, predicate func(int) bool) int {
    for i, num := range nums {
        if predicate(num) {
            return i
        }
    }
    return -1  // not found
}

// Find all indices where element appears
func FindAll(nums []int, target int) []int {
    indices := []int{}

    for i, num := range nums {
        if num == target {
            indices = append(indices, i)
        }
    }

    return indices
}

// Binary search pattern (for sorted data)
func BinarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1
}
```

**Common Variations:**
- Search with early exit: return as soon as found
- Search with counting: how many matches?
- Search in nested structures: loop within loop
- Binary search: O(log n) for sorted data
- Search with position: return index instead of value

**Practice Problem:**

Write a function that finds the first non-repeating character in a string. Return the character itself, or a zero value if all characters repeat.

---

### Pattern 10: Running Calculations

**When to Use It:**
- Computing values that depend on previous computations
- Moving averages, running sums, cumulative products
- Problems asking for "running total", "moving average", "prefix sum"
- Optimization problems where you need intermediate results
- Any calculation where each step depends on the previous

**Mental Model:**

Imagine **keeping a running tally** as you walk through data. At each step, you update your running value based on the new element and possibly the previous running value. You might track the entire history or just the current value.

**Generic Pseudocode:**
```
// Store all running values
create results array
initialize running_value

for each element:
    update running_value using element
    append running_value to results

return results

// Use just current running value
initialize running_value

for each element:
    update running_value using element
    possibly update best/worst running_value

return final running_value or best/worst
```

**Go Implementation Example:**
```go
// Prefix sum array (running sum at each position)
func PrefixSum(nums []int) []int {
    prefix := make([]int, len(nums))
    running := 0

    for i, num := range nums {
        running += num
        prefix[i] = running
    }

    return prefix
}

// Running average
func RunningAverage(nums []float64) []float64 {
    result := make([]float64, len(nums))
    sum := 0.0

    for i, num := range nums {
        sum += num
        result[i] = sum / float64(i+1)
    }

    return result
}

// Maximum subarray sum (Kadane's algorithm)
func MaxSubarraySum(nums []int) int {
    maxSoFar := nums[0]
    maxEndingHere := nums[0]

    for i := 1; i < len(nums); i++ {
        maxEndingHere = max(nums[i], maxEndingHere+nums[i])
        maxSoFar = max(maxSoFar, maxEndingHere)
    }

    return maxSoFar
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

**Common Variations:**
- Prefix sum: enables O(1) range sum queries
- Running min/max: track extremes seen so far
- Moving window average: only last K elements
- Cumulative product: multiply instead of add
- Running comparison: track how current compares to previous

**Practice Problem:**

Write a function that returns a slice where each element is the product of all elements before it in the input slice. The first element should be 1. For example, `[2, 3, 4]` becomes `[1, 2, 6]` (because 1, 1*2, 1*2*3).

---

## Pattern Recognition Practice

For each problem below, identify which pattern(s) would be most helpful. Don't write code—just think about the approach.

### Problem 1
Write a function that finds the longest word in a sentence.

### Problem 2
Given a slice of integers, determine if any two adjacent numbers sum to a specific target value.

### Problem 3
Count how many times each unique number appears in a slice, then return only the numbers that appear more than once.

### Problem 4
Find the smallest window in a string that contains all vowels (a, e, i, o, u) at least once.

### Problem 5
Determine if a string of brackets (including (), [], {}) is properly balanced and nested.

### Problem 6
Calculate the sum of all prime numbers in a slice of integers.

### Problem 7
Given a slice of strings, return a new slice containing only strings that start and end with the same character.

### Problem 8
Find the longest increasing subsequence in a slice of integers (elements don't need to be consecutive, but must maintain order).

### Problem 9
Reverse the words in a sentence but keep the sentence structure intact.

### Problem 10
Given a sorted slice of integers, find if there exists a triplet (three numbers) that sums to zero.

---

## Pattern Recognition Answers

**Problem 1:** Find Maximum in Collection (Pattern 1)
- Track the longest word seen so far as you iterate

**Problem 2:** State Tracking During Iteration (Pattern 5)
- Remember the previous number as you iterate through current numbers

**Problem 3:** Frequency Counting (Pattern 2) + Filter and Transform (Pattern 7)
- First build frequency map, then filter entries with count > 1

**Problem 4:** Sliding Window (Pattern 4) + State Tracking (Pattern 5)
- Variable-size window that tracks which vowels are present

**Problem 5:** State Tracking During Iteration (Pattern 5)
- Track nesting depth and bracket type using stack or counter states

**Problem 6:** Accumulator Pattern (Pattern 6) + Filter (Pattern 7)
- Filter for primes, then accumulate their sum

**Problem 7:** Filter and Transform (Pattern 7)
- Filter based on first/last character condition

**Problem 8:** Running Calculations (Pattern 10) + Find Maximum (Pattern 1)
- Dynamic programming approach: track longest subsequence ending at each position

**Problem 9:** String Building (Pattern 8) + Filter/Transform (Pattern 7)
- Split into words, reverse each word, rebuild sentence

**Problem 10:** Two-Pointer Technique (Pattern 3) with nested iteration
- Fix one element, use two-pointers on remaining sorted elements

---

## Go-Specific Tips

### Leveraging `range` for Patterns

Go's `range` keyword is your friend for most iteration patterns:

```go
// Value only (when index doesn't matter)
for _, value := range slice {
    // use value
}

// Index and value (when position matters)
for i, value := range slice {
    // use both i and value
}

// Map iteration (order not guaranteed)
for key, value := range myMap {
    // use key and value
}

// String iteration (rune by rune)
for i, char := range str {
    // i is byte index, char is rune
}
```

### When to Use Maps vs Slices

**Use maps when:**
- You need O(1) lookup by key
- You're doing frequency counting
- Keys are not sequential integers
- You need to check existence quickly

**Use slices when:**
- You need ordered data
- You're accumulating results
- You need sequential access
- Memory efficiency is critical for dense integer keys

```go
// Map for frequency counting
freq := make(map[string]int)
freq[word]++  // Clean and efficient

// Slice for accumulating results
results := []int{}
results = append(results, value)
```

### String Building Best Practices

```go
// DON'T: Repeated concatenation (creates many temporary strings)
result := ""
for _, s := range strings {
    result += s  // Inefficient!
}

// DO: Use strings.Builder
var builder strings.Builder
for _, s := range strings {
    builder.WriteString(s)
}
result := builder.String()

// DO: Pre-allocate if you know the size
var builder strings.Builder
builder.Grow(estimatedSize)  // One allocation

// DO: Use strings.Join for simple cases
result := strings.Join(strings, separator)
```

### Using the Standard Library Effectively

**Key packages for common patterns:**

```go
import (
    "sort"      // Sorting, binary search
    "strings"   // String manipulation
    "strconv"   // String/number conversion
    "math"      // Math operations, constants
)

// Sorting enables better algorithms
sort.Ints(nums)  // Now you can use binary search

// Built-in string functions
strings.Contains(s, substr)
strings.HasPrefix(s, prefix)
strings.Fields(s)  // Split on whitespace
strings.Count(s, substr)

// Sort package has useful helpers
sort.Search(len(data), func(i int) bool {
    return data[i] >= target
})
```

### Zero Values Are Your Friend

```go
// Map: zero value for missing key
counts := make(map[string]int)
counts["key"]++  // Works even if key doesn't exist (0 + 1)

// Boolean: zero value is false
seen := make(map[string]bool)
if !seen[item] {  // Clean check
    seen[item] = true
}

// Slice: zero value is nil, but can append
var results []int  // nil slice
results = append(results, 1)  // Works fine
```

### Idiomatic Go Patterns

```go
// Check map existence
if value, exists := myMap[key]; exists {
    // key was found, value is valid
}

// Multi-value swap (useful for two-pointers)
left, right = right, left

// Defer for cleanup
func ProcessFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()  // Ensures cleanup

    // process file
}

// Blank identifier for unused values
for i := range slice {  // Don't need value
    // just need index
}
```

---

## Resources

### Go Documentation

**Official Documentation:**
- [Go Tour](https://go.dev/tour/) - Interactive introduction
- [Effective Go](https://go.dev/doc/effective_go) - Idiomatic patterns
- [pkg.go.dev](https://pkg.go.dev/) - Package documentation

**Navigating pkg.go.dev:**
1. Search for package name (e.g., "strings")
2. Look at "Index" for all functions
3. Click functions to see examples
4. Check "Examples" section for real usage

**Key standard library packages:**
- `strings`: String manipulation
- `sort`: Sorting and searching
- `strconv`: String conversions
- `fmt`: Formatting and printing
- `math`: Mathematical functions
- `regexp`: Regular expressions

### Problem-Solving Approach Flowchart

```
Start: Read problem carefully
  |
  v
Identify what you're looking for
(max/min? existence? transformation? accumulation?)
  |
  v
Ask: "What pattern fits?"
  |
  +-- Need maximum/minimum? --> Pattern 1: Find Max/Min
  |
  +-- Need to count things? --> Pattern 2: Frequency Counting
  |
  +-- Working with pairs/comparisons? --> Pattern 3: Two-Pointer
  |
  +-- Looking at subarrays/substrings? --> Pattern 4: Sliding Window
  |
  +-- Need context from previous elements? --> Pattern 5: State Tracking
  |
  +-- Building up a total/sum/product? --> Pattern 6: Accumulator
  |
  +-- Selecting/modifying items? --> Pattern 7: Filter/Transform
  |
  +-- Building a string? --> Pattern 8: String Building
  |
  +-- Checking existence/finding? --> Pattern 9: Search
  |
  +-- Need values at each step? --> Pattern 10: Running Calculations
  |
  v
Write pseudocode first (don't code yet!)
  |
  v
Test pseudocode with small example
  |
  v
Implement in Go
  |
  v
Test with edge cases:
  - Empty input
  - Single element
  - All same values
  - Extreme values
  |
  v
Done!
```

### Recommended Practice Approach

1. **Read the problem twice** - Understand what's being asked
2. **Identify the pattern** - What's the core algorithmic approach?
3. **Write pseudocode** - Language-agnostic logic
4. **Consider edge cases** - What could go wrong?
5. **Implement in Go** - Now focus on syntax
6. **Test thoroughly** - Don't trust your first attempt

### Additional Learning Resources

**Books:**
- "The Go Programming Language" by Donovan & Kernighan
- "Learning Go" by Jon Bodner

**Practice Sites:**
- [Exercism Go Track](https://exercism.org/tracks/go) - Guided practice
- [LeetCode](https://leetcode.com/) - Algorithm practice (try Easy problems)
- [Go by Example](https://gobyexample.com/) - Code examples

**Community:**
- [Go Forum](https://forum.golangbridge.org/)
- [r/golang](https://reddit.com/r/golang)
- [Gophers Slack](https://gophers.slack.com/)

---

## Final Thoughts

Pattern recognition is a **skill that develops with practice**. Don't expect to master it immediately. Each time you solve a problem:

1. Identify which pattern you used
2. Reflect on how you recognized it
3. Note variations you encountered
4. Add to your mental library

Over time, you'll start recognizing patterns automatically. When you see "find the most common", you'll immediately think "that's frequency counting followed by find maximum". When you see "longest substring with condition", you'll think "sliding window".

**The goal isn't to memorize code—it's to recognize the underlying pattern.**

Keep this guide handy. Refer back when you're stuck. And remember: every expert programmer started exactly where you are now.

Happy coding!
