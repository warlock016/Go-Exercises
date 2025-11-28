# Exercise 07: Integer Algorithms

**Tier:** 2 - Patterns
**Estimated Time:** 50 minutes
**Concepts:** Mathematical recursion, GCD, digit manipulation, number properties

---

## Learning Goal

Apply recursion to classic mathematical algorithms. Understand how recursion naturally expresses mathematical recurrence relations.

---

## Problem Description

Implement mathematical algorithms using recursion.

### 1. Greatest Common Divisor (GCD)

Find the GCD of two positive integers using Euclid's algorithm:
- `gcd(a, b) = gcd(b, a mod b)` when b ≠ 0
- `gcd(a, 0) = a`

Example: `gcd(48, 18) = 6`

### 2. Sum of Digits

Calculate the sum of all digits in a number.
- `sumDigits(1234) = 1 + 2 + 3 + 4 = 10`

### 3. Count Digits

Count how many digits are in a number.
- `countDigits(12345) = 5`

### 4. Is Power of Two

Check if a number is a power of 2 using recursion.
- Powers of 2: 1, 2, 4, 8, 16, 32, 64...

---

## Function Signatures

```go
func GCD(a, b int) int

func SumDigits(n int) int

func CountDigits(n int) int

func IsPowerOfTwo(n int) bool
```

---

## Examples

### GCD
```go
GCD(48, 18)   // 6
GCD(100, 35)  // 5
GCD(7, 3)     // 1 (coprime)
GCD(12, 8)    // 4
GCD(17, 17)   // 17 (same number)
```

### SumDigits
```go
SumDigits(0)      // 0
SumDigits(5)      // 5
SumDigits(123)    // 6 (1+2+3)
SumDigits(999)    // 27 (9+9+9)
SumDigits(1234)   // 10
```

### CountDigits
```go
CountDigits(0)      // 1
CountDigits(5)      // 1
CountDigits(42)     // 2
CountDigits(999)    // 3
CountDigits(12345)  // 5
```

### IsPowerOfTwo
```go
IsPowerOfTwo(1)    // true (2^0)
IsPowerOfTwo(2)    // true (2^1)
IsPowerOfTwo(8)    // true (2^3)
IsPowerOfTwo(16)   // true (2^4)
IsPowerOfTwo(3)    // false
IsPowerOfTwo(10)   // false
IsPowerOfTwo(0)    // false
```

---

## Instructions

1. **Understand Euclid's GCD algorithm**:
   ```
   gcd(48, 18):
     → gcd(18, 48 % 18) = gcd(18, 12)
       → gcd(12, 18 % 12) = gcd(12, 6)
         → gcd(6, 12 % 6) = gcd(6, 0)
           → 6 (base case: b = 0)
   ```

2. **Implement GCD**:
   - Base case: `b == 0` → return `a`
   - Recursive case: `GCD(b, a % b)`

3. **Digit manipulation pattern**:
   ```go
   lastDigit := n % 10
   remainingDigits := n / 10
   ```

4. **Implement SumDigits**:
   - Base case: `n == 0` → return 0
   - Recursive: `(n % 10) + SumDigits(n / 10)`
   - Example: 123 → 3 + SumDigits(12)

5. **Implement CountDigits**:
   - Base case: `n < 10` → return 1 (single digit)
   - Recursive: `1 + CountDigits(n / 10)`

6. **Implement IsPowerOfTwo**:
   - Base cases:
     - `n == 1` → true (2^0)
     - `n == 0 or n is odd` → false
   - Recursive: `IsPowerOfTwo(n / 2)`

7. **Run tests**: `go test -v`

---

## Hints

<details>
<summary><strong>Hint 1 - Basic (Understanding the Algorithms)</strong></summary>

**GCD (Euclid's algorithm):**
- Repeatedly replace the larger number with the remainder
- When remainder is 0, the other number is the GCD
- Example: gcd(24, 9) → gcd(9, 6) → gcd(6, 3) → gcd(3, 0) → 3

**Digit extraction:**
- `n % 10` gives last digit
- `n / 10` removes last digit (integer division)
- Example: 1234 % 10 = 4, 1234 / 10 = 123

**Power of 2 property:**
- Every power of 2 is divisible by 2 (except 1)
- Keep dividing by 2; if you reach 1, it's a power of 2
- If you hit an odd number (other than 1), it's not

</details>

<details>
<summary><strong>Hint 2 - Intermediate (Base Cases)</strong></summary>

**GCD:**
```go
if b == 0 {
    return a
}
```

**SumDigits:**
```go
if n == 0 {
    return 0
}
// Or: n < 10 returns n for single digit
```

**CountDigits:**
```go
if n < 10 {
    return 1
}
```

**IsPowerOfTwo:**
```go
if n == 1 {
    return true  // 2^0 = 1
}
if n <= 0 || n%2 != 0 {
    return false  // Not a power of 2
}
```

</details>

<details>
<summary><strong>Hint 3 - Advanced (Handling Edge Cases)</strong></summary>

**GCD edge cases:**
- What if a < b? Euclid's algorithm still works! First step swaps them: `gcd(18, 48)` → `gcd(48, 18)`
- What if one is 0? Return the other number
- What about negative numbers? Take absolute values first

**SumDigits with negatives:**
```go
if n < 0 {
    n = -n  // Take absolute value
}
```

**CountDigits with 0:**
- 0 has 1 digit, not 0 digits
- Special case: `if n == 0 { return 1 }`

**IsPowerOfTwo:**
- 0 is not a power of 2
- 1 is a power of 2 (2^0)
- Negative numbers are not powers of 2

</details>

<details>
<summary><strong>Hint 4 - Complete Solutions</strong></summary>

**GCD:**
```go
func GCD(a, b int) int {
    if b == 0 {
        return a
    }
    return GCD(b, a%b)
}
```

**SumDigits:**
```go
func SumDigits(n int) int {
    if n < 0 {
        n = -n  // Handle negative
    }
    if n == 0 {
        return 0
    }
    return (n % 10) + SumDigits(n/10)
}
```

**CountDigits:**
```go
func CountDigits(n int) int {
    if n < 0 {
        n = -n  // Handle negative
    }
    if n < 10 {
        return 1
    }
    return 1 + CountDigits(n/10)
}
```

**IsPowerOfTwo:**
```go
func IsPowerOfTwo(n int) bool {
    if n == 1 {
        return true
    }
    if n <= 0 || n%2 != 0 {
        return false
    }
    return IsPowerOfTwo(n / 2)
}
```

</details>

---

## Think About

1. **Euclid's genius:** Why does the GCD algorithm work? (It preserves common divisors through each step)

2. **Iterative alternatives:** All of these can be done with loops. When is recursion clearer?
   ```go
   // Iterative GCD
   for b != 0 {
       a, b = b, a%b
   }
   return a
   ```

3. **Tail recursion:** Notice that GCD is tail-recursive! Could a compiler optimize it to a loop?

4. **Digit manipulation:** Why does `n / 10` remove the last digit in base 10? Would this work in binary (n / 2)?

5. **Power of 2 - bit manipulation:** There's a clever one-line solution using bitwise AND:
   ```go
   return n > 0 && (n & (n-1)) == 0
   ```
   Can you understand why this works?

6. **Mathematical induction:** These algorithms prove themselves through induction - each step reduces to a smaller, solved problem.

---

## What This Teaches

✅ **Mathematical recursion** - Expressing math formulas as code
✅ **Euclid's algorithm** - Classic recursive algorithm (2300 years old!)
✅ **Digit manipulation** - Using `% 10` and `/ 10` pattern
✅ **Number properties** - Powers, divisibility, digit sums
✅ **Tail recursion in practice** - GCD is naturally tail-recursive
✅ **Recurrence relations** - Understanding how problems reduce

**Historical note:** Euclid's GCD algorithm (300 BCE) is one of the oldest algorithms still in common use!

---

**Tier 2 Complete!** You've mastered common recursive patterns. Next: Tier 3 - Divide & Conquer algorithms.

**Next Exercise:** 08 - Binary Search Recursive
