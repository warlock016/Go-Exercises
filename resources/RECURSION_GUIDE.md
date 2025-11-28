# Recursion: A Comprehensive Guide

**Author:** Learning Journey Documentation
**Date:** 2025-11-23
**Context:** Recursion Mastery Module (00.9)

---

## Table of Contents

1. [What is Recursion?](#what-is-recursion)
2. [The Recursive Mindset](#the-recursive-mindset)
3. [Anatomy of a Recursive Function](#anatomy-of-a-recursive-function)
4. [The Call Stack](#the-call-stack)
5. [Common Recursive Patterns](#common-recursive-patterns)
6. [Recursion vs Iteration](#recursion-vs-iteration)
7. [Debugging Recursive Functions](#debugging-recursive-functions)
8. [Optimization Techniques](#optimization-techniques)
9. [When to Use Recursion](#when-to-use-recursion)
10. [Common Pitfalls](#common-pitfalls)

---

## What is Recursion?

**Simple Definition:** A function that calls itself.

**Better Definition:** A problem-solving technique where you solve a problem by solving smaller instances of the same problem.

**Best Definition:** Recursion is **self-reference with progress toward a base case**.

###

 The Three Essential Components

Every recursive solution has exactly three parts:

1. **Base Case** - When to stop (termination condition)
2. **Recursive Case** - How to break down the problem
3. **Progress** - Each call must get closer to the base case

**Example: Factorial**
```go
func Factorial(n int) int {
    // 1. BASE CASE: When to stop
    if n == 0 {
        return 1
    }

    // 2. RECURSIVE CASE: Break down the problem
    // 3. PROGRESS: n-1 is smaller than n
    return n * Factorial(n-1)
}
```

---

## The Recursive Mindset

### **The Leap of Faith**

The hardest part of learning recursion is **trusting that it works**.

**Normal thinking:**
"I need to know exactly what happens at each step"

**Recursive thinking:**
"I assume the function works for smaller inputs, then use that assumption"

**Example: Reversing a string**

**Normal approach (iterative):**
```
"hello" → start from end → build "olleh"
```

**Recursive approach:**
```
reverse("hello") = reverse("ello") + "h"
                 = (reverse("llo") + "e") + "h"
                 = ((reverse("lo") + "l") + "e") + "h"
                 = (((reverse("o") + "l") + "l") + "e") + "h"
                 = ((("o" + "l") + "l") + "e") + "h"
                 = "olleh"
```

**The leap:** When writing `reverse("hello")`, you don't think about ALL the recursive calls. You just think:
1. What's the simplest case? (single character → return it)
2. How do I use the solution to a smaller problem? (reverse rest + first char)

---

## Anatomy of a Recursive Function

### **Template for Any Recursive Function**

```go
func RecursiveFunction(problem) result {
    // 1. BASE CASE(S) - Handle simplest case(s)
    if problemIsSimple(problem) {
        return directSolution(problem)
    }

    // 2. RECURSIVE CASE - Break problem into smaller pieces
    smallerProblem := reduceProblemSize(problem)
    smallerSolution := RecursiveFunction(smallerProblem)

    // 3. COMBINE - Use smaller solution to solve original
    return combineSolutions(smallerSolution, problem)
}
```

### **Example: Sum of Array**

```go
func Sum(arr []int) int {
    // BASE CASE: Empty array
    if len(arr) == 0 {
        return 0
    }

    // RECURSIVE CASE: First element + sum of rest
    return arr[0] + Sum(arr[1:])
}
```

**Why this works:**
- `Sum([5])` = 5 + Sum([]) = 5 + 0 = 5 ✓
- `Sum([3,5])` = 3 + Sum([5]) = 3 + 5 = 8 ✓
- `Sum([1,3,5])` = 1 + Sum([3,5]) = 1 + 8 = 9 ✓

---

## The Call Stack

### **What Actually Happens**

When a function calls itself, each call gets its own **stack frame** with its own parameters and variables.

**Example: Factorial(3)**

```
Call Stack Visualization:

Step 1: Call Factorial(3)
┌────────────────┐
│ Factorial(3)   │
│ n = 3          │
│ waiting...     │
└────────────────┘

Step 2: Factorial(3) calls Factorial(2)
┌────────────────┐
│ Factorial(2)   │  ← Current
│ n = 2          │
│ waiting...     │
├────────────────┤
│ Factorial(3)   │
│ n = 3          │
│ waiting...     │
└────────────────┘

Step 3: Factorial(2) calls Factorial(1)
┌────────────────┐
│ Factorial(1)   │  ← Current
│ n = 1          │
│ waiting...     │
├────────────────┤
│ Factorial(2)   │
│ n = 2          │
│ waiting...     │
├────────────────┤
│ Factorial(3)   │
│ n = 3          │
│ waiting...     │
└────────────────┘

Step 4: Factorial(1) calls Factorial(0)
┌────────────────┐
│ Factorial(0)   │  ← Current
│ n = 0          │
│ return 1       │  ← BASE CASE!
├────────────────┤
│ Factorial(1)   │
│ n = 1          │
│ waiting...     │
├────────────────┤
│ Factorial(2)   │
│ n = 2          │
│ waiting...     │
├────────────────┤
│ Factorial(3)   │
│ n = 3          │
│ waiting...     │
└────────────────┘

Step 5: Factorial(0) returns 1
┌────────────────┐
│ Factorial(1)   │  ← Current
│ n = 1          │
│ return 1 * 1   │  = 1
├────────────────┤
│ Factorial(2)   │
│ n = 2          │
│ waiting...     │
├────────────────┤
│ Factorial(3)   │
│ n = 3          │
│ waiting...     │
└────────────────┘

Step 6: Factorial(1) returns 1
┌────────────────┐
│ Factorial(2)   │  ← Current
│ n = 2          │
│ return 2 * 1   │  = 2
├────────────────┤
│ Factorial(3)   │
│ n = 3          │
│ waiting...     │
└────────────────┘

Step 7: Factorial(2) returns 2
┌────────────────┐
│ Factorial(3)   │  ← Current
│ n = 3          │
│ return 3 * 2   │  = 6
└────────────────┘

Step 8: Factorial(3) returns 6
(Stack empty)
Final result: 6
```

**Key Insight:** Recursive calls go "down" until hitting base case, then return values go "up" the stack.

---

## Common Recursive Patterns

### **Pattern 1: Single Recursion**

One recursive call per function execution.

**Structure:**
```go
func f(n) {
    if baseCase { return simple }
    return combine(n, f(n-1))  // One recursive call
}
```

**Examples:** Factorial, power, sum of digits

**Complexity:** Usually O(n) time, O(n) space (call stack)

---

### **Pattern 2: Multiple Recursion**

Multiple recursive calls per execution.

**Structure:**
```go
func f(n) {
    if baseCase { return simple }
    return combine(f(n-1), f(n-2))  // Two recursive calls
}
```

**Examples:** Fibonacci, binary tree traversal

**Complexity:** Can be exponential! O(2^n) for naive Fibonacci

**Call tree for Fib(5):**
```
                    fib(5)
                   /      \
              fib(4)      fib(3)
             /     \       /    \
        fib(3)   fib(2) fib(2) fib(1)
        /   \     /  \    /  \
    fib(2) fib(1) ...  ...  ...
    /  \
fib(1) fib(0)
```

---

### **Pattern 3: Helper Function (Accumulator)**

Use a helper function to carry state through recursive calls.

**Structure:**
```go
func f(input) result {
    return helper(input, initialState)
}

func helper(input, state) result {
    if baseCase { return state }
    return helper(smallerInput, updatedState)
}
```

**Example: Reverse List**
```go
func Reverse(list []int) []int {
    return reverseHelper(list, []int{})
}

func reverseHelper(list, acc []int) []int {
    if len(list) == 0 {
        return acc
    }
    return reverseHelper(list[1:], append([]int{list[0]}, acc...))
}
```

**Why useful:** Builds up result as you go, avoids post-processing

---

### **Pattern 4: Divide and Conquer**

Split problem in half, solve both halves, combine results.

**Structure:**
```go
func divideAndConquer(arr) {
    if len(arr) <= 1 { return arr }

    mid := len(arr) / 2
    left := divideAndConquer(arr[:mid])
    right := divideAndConquer(arr[mid:])

    return combine(left, right)
}
```

**Examples:** Merge sort, quick sort, binary search

**Complexity:** Often O(n log n) time

---

### **Pattern 5: Backtracking**

Try options recursively, undo if they don't work.

**Structure:**
```go
func backtrack(state, choices) {
    if isGoal(state) {
        recordSolution(state)
        return
    }

    for _, choice := range choices {
        makeChoice(state, choice)      // Choose
        backtrack(state, choices)       // Explore
        undoChoice(state, choice)       // Unchoose
    }
}
```

**Examples:** N-Queens, Sudoku, all permutations, FindAllPathsDFS

**You've already mastered this pattern in Graph Theory!**

---

## Recursion vs Iteration

### **When Each is Better**

| Factor | Recursion | Iteration |
|--------|-----------|-----------|
| **Readability** | Better for tree/graph problems | Better for simple loops |
| **Performance** | Slower (function call overhead) | Faster (no overhead) |
| **Memory** | O(depth) stack space | O(1) usually |
| **Natural fit** | Trees, graphs, divide-and-conquer | Arrays, counting, accumulation |
| **Risk** | Stack overflow on deep recursion | Infinite loops |

### **Same Problem, Both Approaches**

**Iterative Factorial:**
```go
func FactorialIter(n int) int {
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}
```

**Recursive Factorial:**
```go
func FactorialRec(n int) int {
    if n == 0 { return 1 }
    return n * FactorialRec(n-1)
}
```

**Pros of Iteration:** Faster, no stack overflow risk
**Pros of Recursion:** More elegant, matches mathematical definition

**The Rule:** Use recursion when it makes code clearer, not just because you can.

---

## Debugging Recursive Functions

### **Strategy 1: Add Print Statements**

```go
func Factorial(n int) int {
    fmt.Printf("Factorial(%d) called\n", n)

    if n == 0 {
        fmt.Printf("Factorial(%d) = 1 (base case)\n", n)
        return 1
    }

    result := n * Factorial(n-1)
    fmt.Printf("Factorial(%d) = %d\n", n, result)
    return result
}
```

**Output for Factorial(3):**
```
Factorial(3) called
Factorial(2) called
Factorial(1) called
Factorial(0) called
Factorial(0) = 1 (base case)
Factorial(1) = 1
Factorial(2) = 2
Factorial(3) = 6
```

---

### **Strategy 2: Trace on Paper**

**Example: Sum([1,2,3])**

```
Call 1: Sum([1,2,3])
  → 1 + Sum([2,3])

  Call 2: Sum([2,3])
    → 2 + Sum([3])

    Call 3: Sum([3])
      → 3 + Sum([])

      Call 4: Sum([])
        → return 0 (base case)

      ← Sum([3]) returns 3 + 0 = 3

    ← Sum([2,3]) returns 2 + 3 = 5

  ← Sum([1,2,3]) returns 1 + 5 = 6

Final: 6
```

---

### **Strategy 3: Check Base Cases First**

**Common bugs:**
- Missing base case → stack overflow
- Wrong base case → incorrect result
- Not progressing toward base case → infinite recursion

**Debug checklist:**
1. ✅ Does base case exist?
2. ✅ Does base case return correct value?
3. ✅ Do all paths reach base case?
4. ✅ Does recursive call use smaller input?

---

## Optimization Techniques

### **Technique 1: Memoization**

Cache results to avoid recomputation.

**Problem: Naive Fibonacci is O(2^n)**
```go
func Fib(n int) int {
    if n <= 1 { return n }
    return Fib(n-1) + Fib(n-2)  // Recomputes same values many times!
}
```

**Solution: Memoization makes it O(n)**
```go
func FibMemo(n int) int {
    memo := make(map[int]int)
    return fibHelper(n, memo)
}

func fibHelper(n int, memo map[int]int) int {
    if n <= 1 { return n }

    if val, exists := memo[n]; exists {
        return val  // Already computed!
    }

    memo[n] = fibHelper(n-1, memo) + fibHelper(n-2, memo)
    return memo[n]
}
```

---

### **Technique 2: Tail Recursion**

When the recursive call is the LAST operation (no work after it returns).

**Not tail-recursive:**
```go
func Factorial(n int) int {
    if n == 0 { return 1 }
    return n * Factorial(n-1)  // Multiplication AFTER recursive call
}
```

**Tail-recursive (with accumulator):**
```go
func FactorialTail(n int) int {
    return factHelper(n, 1)
}

func factHelper(n, acc int) int {
    if n == 0 { return acc }
    return factHelper(n-1, n*acc)  // Nothing after recursive call
}
```

**Why it matters:** Some languages optimize tail recursion to avoid stack growth. **Go does NOT** (as of 2025), but it's still good to know the pattern.

---

### **Technique 3: Convert to Iteration**

When recursion depth is too deep, convert to iteration.

**Recursive (risky for large n):**
```go
func Sum(arr []int) int {
    if len(arr) == 0 { return 0 }
    return arr[0] + Sum(arr[1:])
}
```

**Iterative (safe):**
```go
func Sum(arr []int) int {
    result := 0
    for _, val := range arr {
        result += val
    }
    return result
}
```

---

## When to Use Recursion

### ✅ **USE Recursion When:**

1. **Problem has recursive structure**
   - Trees (each subtree is a tree)
   - Graphs (explore neighbors recursively)
   - Nested data structures

2. **Divide-and-conquer applies**
   - Merge sort, quick sort
   - Binary search
   - Problems that split naturally

3. **Backtracking is needed**
   - Generate all permutations/combinations
   - Solve puzzles (N-Queens, Sudoku)
   - Find all paths in graph

4. **Mathematical definition is recursive**
   - Factorial: n! = n × (n-1)!
   - Fibonacci: F(n) = F(n-1) + F(n-2)
   - GCD: gcd(a,b) = gcd(b, a mod b)

---

### ❌ **AVOID Recursion When:**

1. **Simple iteration is clearer**
   ```go
   // Don't do this recursively:
   func PrintNumbers(n int) {
       if n == 0 { return }
       fmt.Println(n)
       PrintNumbers(n-1)
   }

   // Use a loop:
   for i := n; i > 0; i-- {
       fmt.Println(i)
   }
   ```

2. **Stack depth could be very deep**
   - Processing millions of items
   - User input determines depth
   - Risk of stack overflow

3. **Performance is critical**
   - Recursion has overhead (function calls)
   - Iteration is faster for simple cases

4. **Go doesn't optimize tail calls**
   - Tail recursion won't save stack space in Go
   - Use iteration instead

---

## Common Pitfalls

### ❌ Pitfall 1: Missing Base Case

```go
func Bad(n int) int {
    return n + Bad(n-1)  // ❌ No base case → stack overflow
}
```

**Fix:** Always have a termination condition
```go
func Good(n int) int {
    if n == 0 { return 0 }  // ✅ Base case
    return n + Good(n-1)
}
```

---

### ❌ Pitfall 2: Multiple Base Cases Needed

```go
func Fib(n int) int {
    if n == 0 { return 0 }
    return Fib(n-1) + Fib(n-2)  // ❌ Fib(-1) called when n=1!
}
```

**Fix:** Handle ALL base cases
```go
func Fib(n int) int {
    if n <= 1 { return n }  // ✅ Covers 0 and 1
    return Fib(n-1) + Fib(n-2)
}
```

---

### ❌ Pitfall 3: Not Making Progress

```go
func Bad(arr []int) int {
    if len(arr) == 0 { return 0 }
    return arr[0] + Bad(arr)  // ❌ Same array! No progress!
}
```

**Fix:** Reduce problem size each call
```go
func Good(arr []int) int {
    if len(arr) == 0 { return 0 }
    return arr[0] + Good(arr[1:])  // ✅ Smaller array
}
```

---

### ❌ Pitfall 4: Forgetting to Return Recursive Result

```go
func Bad(n int) int {
    if n == 0 { return 1 }
    n * Bad(n-1)  // ❌ Missing return!
}
```

**Fix:** Always return the recursive call result
```go
func Good(n int) int {
    if n == 0 { return 1 }
    return n * Good(n-1)  // ✅ Return statement
}
```

---

### ❌ Pitfall 5: Modifying Shared State

```go
var result int  // ❌ Global state

func Bad(arr []int) {
    if len(arr) == 0 { return }
    result += arr[0]  // ❌ Side effect!
    Bad(arr[1:])
}
```

**Fix:** Pass state as parameters or return values
```go
func Good(arr []int) int {
    if len(arr) == 0 { return 0 }
    return arr[0] + Good(arr[1:])  // ✅ Functional style
}
```

---

## Mental Models

### **Model 1: The Domino Effect**

Think of recursion like dominos:
- If you can knock over domino N by knocking over domino N-1...
- And you can knock over domino 1 (base case)...
- Then you can knock over ALL dominos!

---

### **Model 2: The Delegation Pattern**

**Manager thinking:**
"I'll handle my part, then delegate the rest to someone else (who uses the same strategy)"

```go
func CountFiles(dir string) int {
    myFiles := countFilesInThisDir(dir)  // My part
    subdirFiles := 0
    for _, subdir := range getSubdirs(dir) {
        subdirFiles += CountFiles(subdir)  // Delegate to "subordinate"
    }
    return myFiles + subdirFiles
}
```

---

### **Model 3: Mathematical Induction**

If you've studied induction:
- **Base case** = Prove P(0)
- **Recursive case** = Prove P(n) using P(n-1)
- **Conclusion** = P(n) true for all n

Recursion IS induction in code form!

---

## Quick Reference

### **Recursion Checklist**

When writing a recursive function:

- [ ] Identified base case(s)
- [ ] Base case returns correct value
- [ ] Recursive case makes problem smaller
- [ ] All paths eventually reach base case
- [ ] Returning recursive call result
- [ ] No shared mutable state
- [ ] Tested with small inputs first

---

### **Common Recursive Algorithms**

| Algorithm | Time | Space | Pattern |
|-----------|------|-------|---------|
| Factorial | O(n) | O(n) | Single recursion |
| Fibonacci (naive) | O(2^n) | O(n) | Multiple recursion |
| Fibonacci (memo) | O(n) | O(n) | Memoization |
| Binary Search | O(log n) | O(log n) | Divide-and-conquer |
| Merge Sort | O(n log n) | O(n) | Divide-and-conquer |
| Quick Sort | O(n log n) avg | O(log n) | Divide-and-conquer |
| DFS (graph) | O(V+E) | O(V) | Backtracking |
| Permutations | O(n!) | O(n) | Backtracking |

---

## Summary

**Recursion in one sentence:**
*A function solves a problem by solving a smaller version of the same problem, until reaching a trivial case.*

**The key insight:**
*Trust the recursion. If it works for smaller inputs, it works for all inputs.*

**When to use it:**
*When the problem has natural recursive structure (trees, divide-and-conquer, backtracking).*

**When to avoid it:**
*When simple iteration is clearer or stack depth is risky.*

---

**Additional Resources:**
- Module 00.9 - Recursion Mastery (13 exercises practicing these concepts)
- Module 00.8 Exercise 02 - Your FindAllPathsDFS implementation (backtracking example)

---

**Last Updated:** 2025-11-23
**Related Modules:** 00.9-recursion-mastery
**Key Quote:** *"To iterate is human, to recurse divine."* - L. Peter Deutsch
