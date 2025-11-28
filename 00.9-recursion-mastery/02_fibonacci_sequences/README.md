# Exercise 02: Fibonacci & Sequences

**Tier:** 1 - Foundation
**Estimated Time:** 60 minutes
**Concepts:** Multiple recursion, exponential growth, optimization awareness

---

## Learning Goal

Understand multiple recursive calls and their performance implications. Experience the difference between naive recursion and optimized approaches.

---

## Problem Description

Implement functions that make **multiple recursive calls** - a function that calls itself more than once.

### 1. Fibonacci Sequence

The Fibonacci sequence is defined as:
- `fib(0) = 0`
- `fib(1) = 1`
- `fib(n) = fib(n-1) + fib(n-2)` for n ≥ 2

Sequence: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89...

### 2. Tribonacci Sequence

Similar to Fibonacci but with three previous terms:
- `trib(0) = 0`
- `trib(1) = 1`
- `trib(2) = 1`
- `trib(n) = trib(n-1) + trib(n-2) + trib(n-3)` for n ≥ 3

Sequence: 0, 1, 1, 2, 4, 7, 13, 24, 44, 81...

### 3. Count Ways to Climb Stairs

You can climb 1 or 2 steps at a time. How many distinct ways can you climb `n` steps?

Example: For n=3, there are 3 ways:
- 1 + 1 + 1
- 1 + 2
- 2 + 1

**Hint:** This is secretly Fibonacci!

---

## Function Signatures

```go
func Fibonacci(n int) int

func Tribonacci(n int) int

func ClimbStairs(n int) int
```

---

## Examples

### Fibonacci
```go
Fibonacci(0)  // 0
Fibonacci(1)  // 1
Fibonacci(2)  // 1
Fibonacci(5)  // 5
Fibonacci(10) // 55
Fibonacci(15) // 610
```

### Tribonacci
```go
Tribonacci(0)  // 0
Tribonacci(1)  // 1
Tribonacci(2)  // 1
Tribonacci(4)  // 4
Tribonacci(10) // 149
```

### ClimbStairs
```go
ClimbStairs(1) // 1 (only one way: 1)
ClimbStairs(2) // 2 (two ways: 1+1, 2)
ClimbStairs(3) // 3 (three ways: 1+1+1, 1+2, 2+1)
ClimbStairs(5) // 8
```

---

## Instructions

1. **Trace Fibonacci by hand** - Draw the recursion tree for `Fibonacci(5)`:
   ```
           fib(5)
          /      \
      fib(4)    fib(3)
      /   \      /   \
   fib(3) fib(2) ...
   ```
   - Notice how `fib(3)` is calculated multiple times!
   - Count total function calls

2. **Implement Fibonacci**:
   - Two base cases: n=0 and n=1
   - Recursive case: sum of two previous terms
   - Test with small values first

3. **Experience the slowdown**:
   - Try `Fibonacci(35)` - notice how slow it is?
   - Try `Fibonacci(40)` - it takes even longer!
   - **Why?** Draw the tree to see duplicate work

4. **Implement Tribonacci**:
   - Three base cases this time
   - Three recursive calls

5. **Implement ClimbStairs**:
   - Think: to reach step `n`, you could come from step `n-1` (1-step) or `n-2` (2-step)
   - Total ways = ways(n-1) + ways(n-2)
   - What does this remind you of?

6. **Run tests**: `go test -v`

7. **Run benchmarks**: `go test -bench=. -benchmem`
   - Compare Fibonacci vs Tribonacci performance

---

## Hints

<details>
<summary><strong>Hint 1 - Basic (Base Cases)</strong></summary>

**Fibonacci:**
- Need TWO base cases: `n == 0` returns 0, `n == 1` returns 1
- Why both? Because recursive case needs `fib(n-2)`, which could go negative without proper base cases

**Tribonacci:**
- Need THREE base cases: 0, 1, and 2

**ClimbStairs:**
- Base cases: `n == 1` returns 1, `n == 2` returns 2
- Or think of it as: `n == 0` returns 1 (one way to climb 0 steps: do nothing)

</details>

<details>
<summary><strong>Hint 2 - Intermediate (Recursive Cases)</strong></summary>

**Fibonacci pattern:**
```go
if n <= 1 {
    return n  // Handles both 0 and 1!
}
return Fibonacci(n-1) + Fibonacci(n-2)
```

**Tribonacci pattern:**
```go
return Tribonacci(n-1) + Tribonacci(n-2) + Tribonacci(n-3)
```

**ClimbStairs pattern:**
- To reach step n, you either:
  - Came from step n-1 (one 1-step move)
  - Came from step n-2 (one 2-step move)
- Total = ways to reach (n-1) + ways to reach (n-2)

</details>

<details>
<summary><strong>Hint 3 - Performance Insight</strong></summary>

**Why is Fibonacci slow?**

For `Fibonacci(5)`, the recursion tree looks like:
```
                    fib(5)
                   /      \
              fib(4)      fib(3)
             /     \      /     \
         fib(3)  fib(2) fib(2) fib(1)
         /   \    /  \   /  \
     fib(2) fib(1) ...
```

Notice:
- `fib(3)` is calculated **twice**
- `fib(2)` is calculated **three times**
- Total calls for `fib(5)`: **15 calls** (exponential growth!)

**Time Complexity:** O(2^n) - doubles with each increase in n
**Space Complexity:** O(n) - maximum call stack depth

This is why `Fibonacci(40)` takes seconds while `Fibonacci(50)` is impractical.

**Optimization preview:** Exercise 12 (Memoization) will make this O(n)!

</details>

<details>
<summary><strong>Hint 4 - Complete Solutions</strong></summary>

**Fibonacci:**
```go
func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}
```

**Tribonacci:**
```go
func Tribonacci(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 || n == 2 {
        return 1
    }
    return Tribonacci(n-1) + Tribonacci(n-2) + Tribonacci(n-3)
}
```

**ClimbStairs:**
```go
func ClimbStairs(n int) int {
    if n <= 2 {
        return n
    }
    return ClimbStairs(n-1) + ClimbStairs(n-2)
}
```

</details>

---

## Think About

1. **Why is Fibonacci O(2^n)?** Draw the recursion tree for n=6 and count the total function calls.

2. **What work is being duplicated?** How many times is `fib(3)` computed when calculating `fib(6)`?

3. **ClimbStairs revelation:** Why does counting stair-climbing ways produce the Fibonacci sequence? What's the connection?

4. **Could you optimize this?** What if you stored previous results in a map? (Hint: That's memoization - Exercise 12!)

5. **Iterative alternative:** Can you implement Fibonacci iteratively with a loop? Would it be faster?

---

## What This Teaches

✅ **Multiple recursion** - Functions calling themselves multiple times
✅ **Exponential time complexity** - Understanding performance implications
✅ **Recursion trees** - Visualizing execution flow
✅ **Duplicate computation** - Recognizing inefficiency (sets up memoization)
✅ **Pattern recognition** - Seeing Fibonacci in unexpected places (stairs!)

This exercise shows both the elegance and the cost of naive recursion. You'll optimize this pattern later!

---

**Performance Warning:** Don't try `Fibonacci(50)` with this naive implementation - it will take minutes! Memoization (Exercise 12) will make it instant.

**Next Exercise:** 03 - Slice Recursion (operating on collections)
