# Exercise 05: Helper Functions

**Tier:** 2 - Patterns
**Estimated Time:** 55 minutes
**Concepts:** Accumulator pattern, wrapper + helper functions, tail recursion

---

## Learning Goal

Master the **accumulator pattern** - using helper functions to carry state through recursion. This pattern enables tail recursion and better performance.

---

## Problem Description

Implement functions using the **wrapper + helper** pattern:
- Public wrapper function provides clean API
- Private helper function carries accumulator(s) through recursion

### 1. Factorial with Accumulator

Implement factorial using tail recursion with an accumulator.

### 2. Reverse List with Accumulator

Reverse a slice by building the result in an accumulator.

### 3. Range Sum

Calculate sum from `start` to `end` using accumulator.
- `RangeSum(1, 5)` → 1+2+3+4+5 = 15

### 4. String to Integer (Simple atoi)

Convert a numeric string to integer recursively.
- `StringToInt("123")` → 123
- `StringToInt("-42")` → -42

---

## Function Signatures

```go
func FactorialAccum(n int) int

func ReverseList(nums []int) []int

func RangeSum(start, end int) int

func StringToInt(s string) int
```

---

## Examples

### FactorialAccum
```go
FactorialAccum(0)  // 1
FactorialAccum(5)  // 120
FactorialAccum(10) // 3628800
```

### ReverseList
```go
ReverseList([]int{})          // []
ReverseList([]int{1, 2, 3})   // [3, 2, 1]
ReverseList([]int{5})         // [5]
```

### RangeSum
```go
RangeSum(1, 5)    // 15 (1+2+3+4+5)
RangeSum(3, 7)    // 25 (3+4+5+6+7)
RangeSum(5, 5)    // 5
RangeSum(10, 1)   // 0 (invalid range)
```

### StringToInt
```go
StringToInt("0")      // 0
StringToInt("123")    // 123
StringToInt("4567")   // 4567
StringToInt("-42")    // -42
StringToInt("-100")   // -100
```

---

## Instructions

1. **Understand the accumulator pattern**:
   ```go
   // WITHOUT accumulator (not tail-recursive)
   func Factorial(n int) int {
       if n == 0 { return 1 }
       return n * Factorial(n-1)  // ❌ Multiplication AFTER recursive call
   }

   // WITH accumulator (tail-recursive)
   func FactorialAccum(n int) int {
       return factorialHelper(n, 1)  // Wrapper starts accumulator at 1
   }

   func factorialHelper(n, accum int) int {
       if n == 0 { return accum }
       return factorialHelper(n-1, n*accum)  // ✅ Recursive call is LAST operation
   }
   ```

2. **Implement FactorialAccum**:
   - Public function: `FactorialAccum(n)` calls helper
   - Helper: `factorialHelper(n, accum)` with accumulator
   - Base case: return accumulator
   - Recursive case: `helper(n-1, n*accum)`

3. **Implement ReverseList**:
   - Helper carries the reversed list
   - Pattern: move first element to accumulator
   - `reverseHelper([1,2,3], [])` → `reverseHelper([2,3], [1])` → `reverseHelper([3], [2,1])` → ...

4. **Implement RangeSum**:
   - Sum from start to end
   - Accumulator carries running sum
   - `rangeHelper(1, 5, 0)` → `rangeHelper(2, 5, 1)` → `rangeHelper(3, 5, 3)` → ...

5. **Implement StringToInt**:
   - Process digit by digit, accumulate result
   - Handle negative sign
   - `"123"` → '1' gives 1, then '2' gives 12, then '3' gives 123

6. **Run tests**: `go test -v`

7. **Compare performance**: Run benchmarks to see accumulator vs non-accumulator

---

## Hints

<details>
<summary><strong>Hint 1 - Basic (Accumulator Pattern Structure)</strong></summary>

**General pattern:**
```go
// Public API - clean interface
func Function(params) result {
    return helper(params, initialAccumulator)
}

// Private helper - carries state
func helper(params, accumulator) result {
    if baseCase {
        return accumulator  // ← Return accumulated result
    }

    // Update accumulator and recurse
    newAccumulator := update(accumulator, currentValue)
    return helper(reducedParams, newAccumulator)
}
```

**Key difference from regular recursion:**
- Regular: compute on the way UP the call stack
- Accumulator: compute on the way DOWN, return accumulated result at base case

</details>

<details>
<summary><strong>Hint 2 - Intermediate (Each Function's Accumulator)</strong></summary>

**FactorialAccum:**
- Accumulator: running product
- Initial value: 1
- Update: `accum * n`

**ReverseList:**
- Accumulator: reversed slice built so far
- Initial value: `[]int{}`
- Update: prepend current element to accumulator

**RangeSum:**
- Accumulator: running sum
- Initial value: 0
- Update: `accum + current`

**StringToInt:**
- Accumulator: integer built so far
- Initial value: 0
- Update: `accum*10 + digitValue`

</details>

<details>
<summary><strong>Hint 3 - Advanced (Tail Recursion Insight)</strong></summary>

**What is tail recursion?**

A function is tail-recursive if the recursive call is the **very last operation**:

```go
// NOT tail-recursive
func Factorial(n int) int {
    if n == 0 { return 1 }
    return n * Factorial(n-1)  // ❌ Multiplication after call
}

// Tail-recursive
func factorialHelper(n, accum int) int {
    if n == 0 { return accum }
    return factorialHelper(n-1, n*accum)  // ✅ Call is last
}
```

**Why does this matter?**
- Tail-recursive functions can be optimized to use O(1) space (no stack growth)
- Some languages/compilers automatically convert tail recursion to loops
- Go doesn't do this optimization, but the pattern is still cleaner and faster

**Trace the difference:**

Regular factorial(5):
```
factorial(5)
  → 5 * factorial(4)
         → 4 * factorial(3)
                → 3 * factorial(2)
                       → 2 * factorial(1)
                              → 1 * factorial(0)
                                     → 1
                              ← 1
                       ← 2
                ← 6
         ← 24
  ← 120
```

Accumulator factorial(5):
```
factorialHelper(5, 1)
  → factorialHelper(4, 5)
       → factorialHelper(3, 20)
            → factorialHelper(2, 60)
                 → factorialHelper(1, 120)
                      → factorialHelper(0, 120)
                           → 120  (done!)
```

</details>

<details>
<summary><strong>Hint 4 - Complete Solutions</strong></summary>

**FactorialAccum:**
```go
func FactorialAccum(n int) int {
    return factorialHelper(n, 1)
}

func factorialHelper(n, accum int) int {
    if n == 0 {
        return accum
    }
    return factorialHelper(n-1, n*accum)
}
```

**ReverseList:**
```go
func ReverseList(nums []int) []int {
    return reverseHelper(nums, []int{})
}

func reverseHelper(nums, accum []int) []int {
    if len(nums) == 0 {
        return accum
    }
    // Prepend first element to accumulator
    newAccum := append([]int{nums[0]}, accum...)
    return reverseHelper(nums[1:], newAccum)
}
```

**RangeSum:**
```go
func RangeSum(start, end int) int {
    if start > end {
        return 0
    }
    return rangeSumHelper(start, end, 0)
}

func rangeSumHelper(current, end, accum int) int {
    if current > end {
        return accum
    }
    return rangeSumHelper(current+1, end, accum+current)
}
```

**StringToInt:**
```go
func StringToInt(s string) int {
    if len(s) == 0 {
        return 0
    }

    negative := false
    if s[0] == '-' {
        negative = true
        s = s[1:]
    }

    result := stringToIntHelper(s, 0)

    if negative {
        return -result
    }
    return result
}

func stringToIntHelper(s string, accum int) int {
    if len(s) == 0 {
        return accum
    }

    digit := int(s[0] - '0')
    return stringToIntHelper(s[1:], accum*10+digit)
}
```

</details>

---

## Think About

1. **Stack growth:** How does accumulator recursion affect stack usage compared to regular recursion?

2. **Tail call optimization:** Go doesn't optimize tail calls. Would other languages (like Scheme or Scala) make this faster?

3. **When to use accumulators?**
   - When you need to build a result incrementally
   - When you want tail recursion
   - When regular recursion is inefficient

4. **Wrapper vs direct:** Why use a wrapper function instead of making users pass the initial accumulator?

5. **Multiple accumulators:** Could you use two accumulators? When would that be useful?

---

## What This Teaches

✅ **Accumulator pattern** - Carrying state through recursion
✅ **Wrapper + helper** - Clean public API with internal helper
✅ **Tail recursion** - Last operation is recursive call
✅ **Building results incrementally** - Constructing answer as you recurse
✅ **Performance improvement** - Avoiding repeated work
✅ **Design pattern** - Separating API from implementation

**This is a professional pattern!** You'll see this in real codebases for parsers, compilers, and functional programming.

---

**Next Exercise:** 06 - List Processing (nested recursion on complex structures)
