# Exercise 01: Factorial & Power

**Tier:** 1 - Foundation
**Estimated Time:** 45 minutes
**Concepts:** Single recursion, base cases, recursive case identification

---

## Learning Goal

Master the fundamental structure of recursive functions by implementing classic mathematical recursions. Understand how to identify base cases and make progress toward them.

---

## Problem Description

Implement two classic recursive functions that follow the pattern: `f(n) → f(n-1)`.

### 1. Factorial

Calculate the factorial of a non-negative integer `n`:
- `factorial(n) = n × (n-1) × (n-2) × ... × 1`
- `factorial(0) = 1` (by mathematical definition)

**Example:**
- `factorial(5) = 5 × 4 × 3 × 2 × 1 = 120`

### 2. Power

Calculate `base^exponent` using only addition/multiplication (no `math.Pow`):
- `power(2, 3) = 2 × 2 × 2 = 8`
- `power(x, 0) = 1` for any x
- Assume exponent is non-negative

---

## Function Signatures

```go
func Factorial(n int) int

func Power(base, exponent int) int
```

---

## Examples

### Factorial
```go
Factorial(0)  // 1
Factorial(1)  // 1
Factorial(5)  // 120
Factorial(7)  // 5040
Factorial(10) // 3628800
```

### Power
```go
Power(2, 0)   // 1
Power(2, 1)   // 2
Power(2, 5)   // 32
Power(3, 4)   // 81
Power(5, 3)   // 125
Power(10, 2)  // 100
```

---

## Instructions

1. **Trace by hand first** - Work through `Factorial(4)` step by step on paper:
   - What is the base case?
   - What does each recursive call return?
   - Draw the call stack

2. **Implement Factorial**:
   - Start with the base case
   - Then write the recursive case
   - Test with small inputs first (0, 1, 2)

3. **Implement Power**:
   - Identify the base case (any number to the power of 0)
   - Recursive case: `base^exp = base × base^(exp-1)`
   - Test incrementally

4. **Run tests**: `go test -v`

5. **Trace with examples**: Mentally trace `Power(2, 3)`:
   ```
   Power(2, 3)
     → 2 × Power(2, 2)
           → 2 × Power(2, 1)
                 → 2 × Power(2, 0)
                       → 1 (base case!)
                 → 2 × 1 = 2
           → 2 × 2 = 4
     → 2 × 4 = 8
   ```

---

## Hints

<details>
<summary><strong>Hint 1 - Basic (Identify Base Cases)</strong></summary>

For **Factorial**:
- What is the smallest input where you know the answer immediately?
- `factorial(0) = 1` (mathematical definition)

For **Power**:
- What happens when the exponent is 0?
- `x^0 = 1` for any x

</details>

<details>
<summary><strong>Hint 2 - Intermediate (Recursive Cases)</strong></summary>

For **Factorial**:
- `factorial(n) = n × factorial(n-1)`
- Example: `factorial(5) = 5 × factorial(4)`

For **Power**:
- `base^exp = base × base^(exp-1)`
- Example: `2^5 = 2 × 2^4`

**Key insight**: Each recursive call gets closer to the base case.

</details>

<details>
<summary><strong>Hint 3 - Complete Solution</strong></summary>

**Factorial:**
```go
func Factorial(n int) int {
    if n == 0 {
        return 1  // Base case
    }
    return n * Factorial(n-1)  // Recursive case
}
```

**Power:**
```go
func Power(base, exponent int) int {
    if exponent == 0 {
        return 1  // Base case
    }
    return base * Power(base, exponent-1)  // Recursive case
}
```

**Why it works:**
- Each recursive call reduces the problem size by 1
- Eventually reaches base case (n=0 or exponent=0)
- Return values multiply back up the call stack

</details>

---

## Think About

1. **What happens if you forget the base case?** Try removing `if n == 0` and running `Factorial(3)`. What error do you get?

2. **Why must each recursive call make progress toward the base case?** What would happen with `Factorial(n-0)` instead of `Factorial(n-1)`?

3. **Can you implement these iteratively?** Which is clearer - the recursive or iterative version?

4. **Performance consideration:** For `Factorial(1000)`, how many stack frames are created? Is this a problem in Go?

---

## What This Teaches

✅ **Base case identification** - Recognizing when to stop recursing
✅ **Recursive case structure** - Making progress toward termination
✅ **Call stack mental model** - Understanding how values return
✅ **Single recursion pattern** - `f(n) → f(n-1)` template

This is the foundation for all recursive thinking. Master this pattern and more complex recursions become approachable.

---

**Next Exercise:** 02 - Fibonacci & Sequences (multiple recursion)
