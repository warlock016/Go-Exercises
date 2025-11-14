# Exercise 10: Backtracking Permutations

**Difficulty:** Hard
**Time:** 50-60 minutes
**Concepts:** Recursion, backtracking, string manipulation, combinatorics

## Learning Goal

Master backtracking - a powerful recursive technique for exploring all possible solutions by making choices, recursing, then undoing choices to try alternatives. This pattern appears in permutations, combinations, N-queens, sudoku solvers, and many constraint satisfaction problems.

## The Problem

Given a string, generate all possible permutations of its characters.

For `"abc"`:
```
["abc", "acb", "bac", "bca", "cab", "cba"]
```

A permutation is a rearrangement of all characters. For a string of length n, there are n! (n factorial) permutations.

## The Backtracking Pattern

Backtracking follows this template:

```
1. Choose an element
2. Explore (recurse) with that choice
3. Un-choose (backtrack) to try other options
```

For permutations:
```
permute("abc", start=0):
    Choose 'a' at position 0 → permute("abc", start=1)
        Choose 'b' at position 1 → permute("abc", start=2)
            Choose 'c' at position 2 → add "abc" to results
        Backtrack, swap back
        Choose 'c' at position 1 → permute("acb", start=2)
            Choose 'b' at position 2 → add "acb" to results
    Backtrack, swap back
    Choose 'b' at position 0 → permute("bac", start=1)
        ... and so on
```

## Function Signature

```go
func Permute(s string) []string
```

## Examples

**Example 1:**
```go
Permute("abc")
// → ["abc", "acb", "bac", "bca", "cab", "cba"]
// Total: 3! = 6 permutations
```

**Example 2:**
```go
Permute("ab")
// → ["ab", "ba"]
// Total: 2! = 2 permutations
```

**Example 3:**
```go
Permute("a")
// → ["a"]
// Total: 1! = 1 permutation
```

**Example 4:**
```go
Permute("")
// → [""]
// Total: 0! = 1 permutation (by convention)
```

**Example 5:**
```go
Permute("aab")
// → ["aab", "aba", "aab", "aba", "baa", "baa"]
// Note: Duplicate characters create duplicate permutations
// For this exercise, we include duplicates (advanced version would deduplicate)
```

## Algorithm: Swap-Based Backtracking

The most elegant approach uses swapping:

```go
func Permute(s string) []string {
    results := []string{}
    runes := []rune(s)  // Convert to rune slice for manipulation
    backtrack(runes, 0, &results)
    return results
}

func backtrack(runes []rune, start int, results *[]string) {
    // Base case: we've fixed all positions
    if start == len(runes) {
        *results = append(*results, string(runes))
        return
    }

    // Try each remaining character at position 'start'
    for i := start; i < len(runes); i++ {
        // Choose: swap runes[start] with runes[i]
        runes[start], runes[i] = runes[i], runes[start]

        // Explore: recurse with start+1
        backtrack(runes, start+1, results)

        // Un-choose: swap back for next iteration
        runes[start], runes[i] = runes[i], runes[start]
    }
}
```

## Visualization

For `"abc"`:

```
                    abc (start=0)
                   / | \
                  /  |  \
         (swap a,a)(a,b)(a,c)
           abc      bac    cba
          start=1  start=1 start=1
           / \      / \     / \
       (b,b)(b,c)(a,a)(a,c)(b,b)(b,a)
        abc  acb   bac  bca  cba  cab
       s=2  s=2   s=2  s=2  s=2  s=2
        ↓    ↓     ↓    ↓    ↓    ↓
       abc  acb   bac  bca  cba  cab  ← Results
```

## Instructions

1. **Convert string to rune slice:**
   - Strings are immutable in Go
   - Need mutable slice for in-place swapping
   - Use `[]rune(s)` for proper Unicode handling

2. **Create results slice:**
   - Will accumulate all permutations
   - Pass pointer to recursive function

3. **Implement backtrack helper:**
   - Parameters: runes slice, start index, results pointer
   - Base case: when start == len(runes), add permutation to results
   - Recursive case: for each position i from start to end:
     - Swap runes[start] with runes[i]
     - Recurse with start+1
     - Swap back (backtrack)

4. **Return results**

## Hints

<details>
<summary>Hint 1: Function structure (Click to reveal)</summary>

```go
func Permute(s string) []string {
    results := []string{}
    runes := []rune(s)
    backtrack(runes, 0, &results)
    return results
}
```
</details>

<details>
<summary>Hint 2: Helper function signature (Click to reveal)</summary>

```go
func backtrack(runes []rune, start int, results *[]string) {
    // start: current position we're fixing
    // runes: working permutation (modified in-place)
    // results: pointer to accumulate all permutations
}
```
</details>

<details>
<summary>Hint 3: Base case (Click to reveal)</summary>

```go
// When we've fixed all positions, we have a complete permutation
if start == len(runes) {
    *results = append(*results, string(runes))
    return
}
```

Why `start == len(runes)` not `start > len(runes)`?
- We're fixing positions 0, 1, 2, ..., n-1
- When start reaches n, we're done
</details>

<details>
<summary>Hint 4: Recursive case with swapping (Click to reveal)</summary>

```go
for i := start; i < len(runes); i++ {
    // Choose: put runes[i] at position 'start'
    runes[start], runes[i] = runes[i], runes[start]

    // Explore: fix next position
    backtrack(runes, start+1, results)

    // Un-choose: restore original order for next iteration
    runes[start], runes[i] = runes[i], runes[start]
}
```

Why swap back?
- The for loop tries different characters at position 'start'
- Must restore original state for next iteration
- This is the "backtracking" step
</details>

<details>
<summary>Hint 5: Complete implementation (Click to reveal)</summary>

```go
func Permute(s string) []string {
    results := []string{}
    runes := []rune(s)
    backtrack(runes, 0, &results)
    return results
}

func backtrack(runes []rune, start int, results *[]string) {
    // Base case: fixed all positions
    if start == len(runes) {
        *results = append(*results, string(runes))
        return
    }

    // Try each character at position 'start'
    for i := start; i < len(runes); i++ {
        // Swap to choose runes[i] for position 'start'
        runes[start], runes[i] = runes[i], runes[start]

        // Recurse to fix remaining positions
        backtrack(runes, start+1, results)

        // Swap back to restore state
        runes[start], runes[i] = runes[i], runes[start]
    }
}
```
</details>

## Think About

1. **Why use runes instead of bytes?**
   - Proper Unicode support
   - "café" has 4 runes but 5 bytes (é is 2 bytes)
   - Permutations should work with any characters

2. **What's the time complexity?**
   - O(n × n!) - n! permutations, each takes O(n) to construct
   - Cannot do better - must generate all permutations

3. **What's the space complexity?**
   - O(n) for recursion depth
   - O(n × n!) for results storage
   - Recursion stack never exceeds n levels

4. **Why does swapping work?**
   - At each level, we "fix" one position
   - Swapping brings different characters to that position
   - By the time we reach the end, we've tried all arrangements

5. **How would you remove duplicates?**
   - Use a set (map[string]bool) to track seen permutations
   - Or skip duplicates during recursion (more efficient)
   - For "aab", you'd get 3 unique permutations instead of 6

6. **Could you do this iteratively?**
   - Yes, but much more complex
   - Would need to track state manually
   - Recursion is the natural fit for backtracking

## What This Teaches

- **Backtracking pattern:** Choose, explore, un-choose
- **Recursion with state:** Passing indices to track progress
- **In-place modification:** Swapping for efficiency
- **Factorial growth:** Understanding exponential algorithms
- **Problem decomposition:** Break into "fix position k" subproblems
- **Rune manipulation:** Proper string handling in Go

## Real-World Applications

- **Optimization:** Traveling salesman (try all routes)
- **Puzzles:** Sudoku, N-queens, crossword solvers
- **Testing:** Generate all input combinations
- **Cryptography:** Key space exploration
- **Scheduling:** All possible arrangements
- **Genetic algorithms:** Initial population generation

Ready to implement? Open `backtracking_permutations.go` and look for `TODO(human)` markers!
