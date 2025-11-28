# Recursion: Variables, Scope, and Value Propagation

**Created:** 2025-11-23
**Purpose:** Understand how variables work in recursive functions, how values propagate through the call stack, and how to correctly combine results from recursive calls.

**Prerequisites:** Basic recursion understanding (see RECURSION_GUIDE.md)

---

## Table of Contents

1. [The Call Stack and Stack Frames](#the-call-stack-and-stack-frames)
2. [Variable Scope in Recursion](#variable-scope-in-recursion)
3. [Three Ways to Propagate Values](#three-ways-to-propagate-values)
4. [Common Patterns for Building Results](#common-patterns-for-building-results)
5. [Common Mistakes and How to Fix Them](#common-mistakes-and-how-to-fix-them)
6. [Visual Stack Traces](#visual-stack-traces)
7. [Practice Examples](#practice-examples)

---

## The Call Stack and Stack Frames

### What Happens When You Call a Function

Every function call creates a **stack frame** containing:
- Function parameters (copies of arguments)
- Local variables
- Return address (where to return after function completes)

These frames are **stacked** on top of each other.

### Visual Example: Factorial

```go
func Factorial(n int) int {
    if n == 0 {
        return 1
    }
    result := n * Factorial(n-1)
    return result
}

// Call: Factorial(3)
```

**The Call Stack Evolution:**

```
Step 1: Call Factorial(3)
┌─────────────────────┐
│ Factorial(3)        │
│ n = 3               │
│ result = ?          │
└─────────────────────┘

Step 2: Factorial(3) calls Factorial(2)
┌─────────────────────┐
│ Factorial(2)        │ ← Currently executing
│ n = 2               │
│ result = ?          │
├─────────────────────┤
│ Factorial(3)        │ ← Waiting
│ n = 3               │
│ result = ?          │
└─────────────────────┘

Step 3: Factorial(2) calls Factorial(1)
┌─────────────────────┐
│ Factorial(1)        │ ← Currently executing
│ n = 1               │
│ result = ?          │
├─────────────────────┤
│ Factorial(2)        │ ← Waiting
│ n = 2               │
│ result = ?          │
├─────────────────────┤
│ Factorial(3)        │ ← Waiting
│ n = 3               │
│ result = ?          │
└─────────────────────┘

Step 4: Factorial(1) calls Factorial(0)
┌─────────────────────┐
│ Factorial(0)        │ ← Currently executing
│ n = 0               │
│ return 1            │ ← Base case!
├─────────────────────┤
│ Factorial(1)        │ ← Waiting
│ n = 1               │
│ result = ?          │
├─────────────────────┤
│ Factorial(2)        │ ← Waiting
│ n = 2               │
│ result = ?          │
├─────────────────────┤
│ Factorial(3)        │ ← Waiting
│ n = 3               │
│ result = ?          │
└─────────────────────┘

Step 5: Factorial(0) returns 1 to Factorial(1)
┌─────────────────────┐
│ Factorial(1)        │ ← Resuming
│ n = 1               │
│ result = 1 * 1 = 1  │ ← Computed!
│ return 1            │
├─────────────────────┤
│ Factorial(2)        │ ← Still waiting
│ n = 2               │
│ result = ?          │
├─────────────────────┤
│ Factorial(3)        │ ← Still waiting
│ n = 3               │
│ result = ?          │
└─────────────────────┘

Step 6: Factorial(1) returns 1 to Factorial(2)
┌─────────────────────┐
│ Factorial(2)        │ ← Resuming
│ n = 2               │
│ result = 2 * 1 = 2  │ ← Computed!
│ return 2            │
├─────────────────────┤
│ Factorial(3)        │ ← Still waiting
│ n = 3               │
│ result = ?          │
└─────────────────────┘

Step 7: Factorial(2) returns 2 to Factorial(3)
┌─────────────────────┐
│ Factorial(3)        │ ← Resuming
│ n = 3               │
│ result = 3 * 2 = 6  │ ← Computed!
│ return 6            │
└─────────────────────┘

Final result: 6
```

**Key Insight:** Each call has its **own** `n` and `result` variables. They don't interfere with each other!

---

## Variable Scope in Recursion

### Rule 1: Each Recursive Call Gets Its Own Variables

```go
func Example(x int) int {
    localVar := x * 2  // Each call has its OWN localVar

    if x == 0 {
        return 0
    }

    result := Example(x - 1)
    // localVar is still x*2 here (unchanged by recursive call)

    return localVar + result
}
```

**Trace Example(2):**

```
Example(2):
  localVar = 4
  calls Example(1):
    localVar = 2  ← Different variable!
    calls Example(0):
      localVar = 0  ← Different variable!
      returns 0
    result = 0
    returns 2 + 0 = 2
  result = 2
  returns 4 + 2 = 6
```

### Rule 2: Parameters Are Copies

In Go, parameters are **passed by value** (copied), so modifying a parameter doesn't affect the caller:

```go
func Countdown(n int) {
    if n == 0 {
        return
    }
    fmt.Println(n)
    n = n - 1          // Modifies LOCAL copy
    Countdown(n)       // Passes the modified copy
}

// Each call has its own copy of n
```

### Rule 3: Slices Are References

**CRITICAL:** Slices are reference types! When you pass a slice, you're passing a reference to the underlying array:

```go
func ModifySlice(nums []int) {
    if len(nums) == 0 {
        return
    }
    nums[0] = 999  // ⚠️ Modifies the ORIGINAL array!
    ModifySlice(nums[1:])
}

func SafeSlice(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }
    // Create NEW slice with modified value
    result := make([]int, len(nums))
    result[0] = nums[0] * 2
    // Recurse and combine
    rest := SafeSlice(nums[1:])
    return append(result[:1], rest...)
}
```

---

## Three Ways to Propagate Values

### Pattern 1: Return Values (Most Common)

**Build result by COMBINING return values**

```go
func Sum(nums []int) int {
    if len(nums) == 0 {
        return 0  // Base case
    }

    // Combine: current element + recursive result
    return nums[0] + Sum(nums[1:])
}

// Trace: Sum([1, 2, 3])
// = 1 + Sum([2, 3])
// = 1 + (2 + Sum([3]))
// = 1 + (2 + (3 + Sum([])))
// = 1 + (2 + (3 + 0))
// = 1 + (2 + 3)
// = 1 + 5
// = 6
```

**Key:** Each level returns a value, and parent combines it with local work.

### Pattern 2: Accumulator (Tail Recursion)

**Pass accumulator DOWN, return it at base case**

```go
func SumAccum(nums []int) int {
    return sumHelper(nums, 0)  // Start with accumulator = 0
}

func sumHelper(nums []int, accum int) int {
    if len(nums) == 0 {
        return accum  // Base case: return accumulated result
    }

    // Pass updated accumulator DOWN to next call
    return sumHelper(nums[1:], accum + nums[0])
}

// Trace: SumAccum([1, 2, 3])
// = sumHelper([1, 2, 3], 0)
// = sumHelper([2, 3], 0+1=1)
// = sumHelper([3], 1+2=3)
// = sumHelper([], 3+3=6)
// = 6
```

**Key:** Build result as you go DOWN the stack, not on the way UP.

### Pattern 3: Shared State (Pointer/Reference)

**Pass a pointer to shared structure, modify it in place**

```go
func CollectEvens(nums []int) []int {
    result := []int{}  // Shared result
    collectHelper(nums, &result)
    return result
}

func collectHelper(nums []int, result *[]int) {
    if len(nums) == 0 {
        return
    }

    if nums[0]%2 == 0 {
        *result = append(*result, nums[0])  // Modify shared result
    }

    collectHelper(nums[1:], result)
}
```

**Key:** All calls modify the SAME result. Use with caution (less "pure").

---

## Common Patterns for Building Results

### Pattern A: Prepend Current + Recurse on Rest

```go
func ReverseString(s string) string {
    if len(s) <= 1 {
        return s
    }

    // Take last character + reverse the rest
    return string(s[len(s)-1]) + ReverseString(s[:len(s)-1])
}

// "hello" → "o" + ReverseString("hell")
//         → "o" + ("l" + ReverseString("hel"))
//         → "o" + ("l" + ("l" + ReverseString("he")))
//         → "o" + ("l" + ("l" + ("e" + ReverseString("h"))))
//         → "o" + ("l" + ("l" + ("e" + "h")))
//         → "olleh"
```

### Pattern B: Recurse First, Then Combine

```go
func Flatten(nested interface{}) []int {
    if num, ok := nested.(int); ok {
        return []int{num}  // Base case
    }

    if slice, ok := nested.([]interface{}); ok {
        result := []int{}
        for _, element := range slice {
            // Recurse first, then append
            result = append(result, Flatten(element)...)
        }
        return result
    }

    return []int{}
}
```

### Pattern C: Merge Two Recursive Results

```go
func MergeSort(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }

    mid := len(nums) / 2

    // Recurse on BOTH halves
    left := MergeSort(nums[:mid])
    right := MergeSort(nums[mid:])

    // Combine the two results
    return Merge(left, right)
}
```

---

## Common Mistakes and How to Fix Them

### Mistake 1: Not Using Return Value

```go
// ❌ WRONG
func Sum(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    Sum(nums[1:])  // ← Return value is IGNORED!

    return nums[0]  // Only returns first element
}

// ✅ CORRECT
func Sum(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    return nums[0] + Sum(nums[1:])  // ← USE the return value!
}
```

**Your Merge bug was exactly this!** You called `Merge(left, right)` but didn't use the result.

### Mistake 2: Modifying Local Variable Instead of Returning

```go
// ❌ WRONG
func Reverse(nums []int) []int {
    result := []int{}

    if len(nums) == 0 {
        return result
    }

    result = append(result, nums[len(nums)-1])
    Reverse(nums[:len(nums)-1])  // ← Doesn't affect result!

    return result  // Only has 1 element
}

// ✅ CORRECT
func Reverse(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }

    // Combine: last element + reversed rest
    rest := Reverse(nums[:len(nums)-1])
    return append(rest, nums[len(nums)-1])
}
```

### Mistake 3: Confusing Slice Slicing

```go
// ❌ WRONG - Creates new slice, doesn't modify original
func RemoveFirst(nums []int) {
    if len(nums) == 0 {
        return
    }
    nums = nums[1:]  // ← Only modifies LOCAL copy!
    // Original slice unchanged
}

// ✅ CORRECT - Return new slice
func RemoveFirst(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }
    return nums[1:]
}
```

### Mistake 4: Infinite Recursion (No Progress Toward Base Case)

```go
// ❌ WRONG
func BadSum(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    // BUG: Doesn't reduce problem size!
    return nums[0] + BadSum(nums)  // ← Should be nums[1:]
}

// ✅ CORRECT
func GoodSum(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    return nums[0] + GoodSum(nums[1:])  // ← Reduces size
}
```

---

## Visual Stack Traces

### Example: Merge Function (Your Bug)

**Your Code:**
```go
func Merge(left, right []int) []int {
    result := make([]int, 0, len(left)+len(right))

    if len(left) == 0 { return right }
    if len(right) == 0 { return left }

    leftElement := left[0]
    left = left[1:]
    rightElement := right[0]
    right = right[1:]

    if leftElement > rightElement {
        result = append(result, rightElement, leftElement)
    } else {
        result = append(result, leftElement, rightElement)
    }

    Merge(left, right)  // ❌ BUG: Return value not used!

    return result  // ← Only has 2 elements!
}
```

**Call Stack: Merge([1, 3], [2, 4])**

```
Step 1: Merge([1, 3], [2, 4])
┌──────────────────────────────┐
│ left = [1, 3]                │
│ right = [2, 4]               │
│ result = []                  │
│                              │
│ leftElement = 1              │
│ left = [3]  (sliced)         │
│ rightElement = 2             │
│ right = [4]  (sliced)        │
│                              │
│ 1 < 2, so append 1, 2        │
│ result = [1, 2]              │
│                              │
│ Calls Merge([3], [4])        │ ← Recursive call
│ ❌ But doesn't use result!   │
│                              │
│ Returns [1, 2]               │ ← WRONG! Missing 3, 4
└──────────────────────────────┘
```

**Correct Version:**

```go
func Merge(left, right []int) []int {
    if len(left) == 0 { return right }
    if len(right) == 0 { return left }

    if left[0] <= right[0] {
        // Take left[0], combine with merge of rest
        rest := Merge(left[1:], right)
        return append([]int{left[0]}, rest...)
    } else {
        // Take right[0], combine with merge of rest
        rest := Merge(left, right[1:])
        return append([]int{right[0]}, rest...)
    }
}
```

**Call Stack: Merge([1, 3], [2, 4]) - CORRECT**

```
Merge([1, 3], [2, 4])
  1 <= 2, take 1
  rest = Merge([3], [2, 4])
    2 <= 3, take 2
    rest = Merge([3], [4])
      3 <= 4, take 3
      rest = Merge([], [4])
        return [4]  ← Base case
      return [3] + [4] = [3, 4]
    return [2] + [3, 4] = [2, 3, 4]
  return [1] + [2, 3, 4] = [1, 2, 3, 4]  ✓
```

---

## Practice Examples

### Exercise 1: Fix the Bug

```go
// Bug: Doesn't collect all elements
func Collect(nums []int) []int {
    result := []int{}

    if len(nums) == 0 {
        return result
    }

    if nums[0]%2 == 0 {
        result = append(result, nums[0])
    }

    Collect(nums[1:])  // ← What's wrong?

    return result
}
```

<details>
<summary>Answer</summary>

The recursive call result is not used! Fix:

```go
func Collect(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }

    rest := Collect(nums[1:])  // ← Store result

    if nums[0]%2 == 0 {
        return append([]int{nums[0]}, rest...)
    }

    return rest
}
```

</details>

### Exercise 2: Trace the Execution

```go
func Mystery(n int, accum int) int {
    if n == 0 {
        return accum
    }
    return Mystery(n-1, accum*10+n)
}

// What does Mystery(123, 0) return?
```

<details>
<summary>Answer</summary>

```
Mystery(123, 0)
  → Mystery(122, 0*10+123=123)
    → Mystery(121, 123*10+122=1352)
      → Mystery(120, 1352*10+121=13641)
        → ... (continues)
```

This reverses the digits of a number using accumulator pattern.

</details>

### Exercise 3: Choose the Right Pattern

Which pattern should you use for:

1. **Reversing a slice** - ?
2. **Calculating factorial with tail recursion** - ?
3. **Finding all paths in a graph** - ?

<details>
<summary>Answers</summary>

1. **Return values** (build result by combining first + reversed rest)
2. **Accumulator** (tail recursive, pass product down)
3. **Shared state** (collect all paths in a shared slice)

</details>

---

## Key Takeaways

1. **Each recursive call has its own variables** (stack frame)
2. **Return values must be USED** - don't throw them away!
3. **Three patterns:**
   - Return values → combine results going UP the stack
   - Accumulator → build results going DOWN the stack
   - Shared state → all calls modify same structure
4. **Slices are tricky:** `nums[1:]` creates new slice view, doesn't modify original
5. **Always make progress:** Each recursive call must reduce problem size

---

## See Also

- [RECURSION_GUIDE.md](./RECURSION_GUIDE.md) - General recursion patterns
- [BFS_PATTERNS_GUIDE.md](./BFS_PATTERNS_GUIDE.md) - Iterative patterns for comparison
- Module 00.9 Exercises 05 (Helper Functions) and 06 (List Processing)

---

**Remember:** When stuck with recursion, draw the call stack! Visualize what each call's variables are, and trace how return values flow upward.
