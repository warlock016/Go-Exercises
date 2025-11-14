# Exercise 11: Regex Matcher (Finite Automaton)

**Difficulty:** Hard
**Time:** 60-70 minutes
**Concepts:** Finite state machines, pattern matching, regex fundamentals

## Learning Goal

Implement a simple regular expression matcher using finite automata. You'll learn how regex engines work under the hood by building state machines that recognize patterns like "a*" (zero or more a's), "a+b" (one or more a's then b), and "ab*c" (a, zero or more b's, c). This exercise bridges the gap between theoretical computer science and practical pattern matching.

## The Problem

Regular expressions are patterns that describe sets of strings. They're implemented using finite automata - state machines that transition based on input characters. In this exercise, you'll implement pattern matching for three core regex operators:

1. **`*` (Kleene star)**: Zero or more of the preceding character
   - Pattern "a*" matches: "", "a", "aa", "aaa", ...

2. **`+` (Plus)**: One or more of the preceding character
   - Pattern "a+" matches: "a", "aa", "aaa", ... (but NOT "")

3. **Literal characters**: Must match exactly
   - Pattern "abc" matches only: "abc"

You'll combine these to handle patterns like:
- "a*b" → zero or more a's, then one b
- "a+b" → one or more a's, then one b
- "ab*c" → one a, zero or more b's, then one c
- "a*b*c*" → any number of a's, b's, and c's in that order

## Your Task

Implement a single function that matches text against a pattern:

```go
func Match(pattern string, text string) bool
```

## Examples

### Basic Patterns

```go
// Kleene star (*)
Match("a*", "")         // → true  (zero a's)
Match("a*", "a")        // → true  (one a)
Match("a*", "aaa")      // → true  (three a's)
Match("a*", "b")        // → false (not an a)

// Plus (+)
Match("a+", "")         // → false (needs at least one a)
Match("a+", "a")        // → true  (one a)
Match("a+", "aaa")      // → true  (three a's)
Match("a+", "b")        // → false (not an a)

// Literals
Match("abc", "abc")     // → true
Match("abc", "ab")      // → false
Match("abc", "abcd")    // → false
```

### Combined Patterns

```go
// a* followed by b
Match("a*b", "b")       // → true  (0 a's, then b)
Match("a*b", "ab")      // → true  (1 a, then b)
Match("a*b", "aaab")    // → true  (3 a's, then b)
Match("a*b", "a")       // → false (no b at end)

// a+ followed by b
Match("a+b", "ab")      // → true  (1 a, then b)
Match("a+b", "aaab")    // → true  (3 a's, then b)
Match("a+b", "b")       // → false (needs at least one a)

// a, then b*, then c
Match("ab*c", "ac")     // → true  (0 b's)
Match("ab*c", "abc")    // → true  (1 b)
Match("ab*c", "abbbbc") // → true  (4 b's)
Match("ab*c", "aac")    // → false (no b)
Match("ab*c", "ab")     // → false (no c)
```

### Multiple Operators

```go
// Multiple Kleene stars
Match("a*b*c*", "")       // → true  (all zero)
Match("a*b*c*", "abc")    // → true
Match("a*b*c*", "aabbcc") // → true
Match("a*b*c*", "bbc")    // → true  (zero a's)
Match("a*b*c*", "cba")    // → false (wrong order)

// Complex patterns
Match("a+b*c", "abc")     // → true
Match("a+b*c", "aac")     // → false (no b allowed when using a+b*)
Match("a*b+c*", "bbc")    // → true  (0 a's, 2 b's, 1 c)
```

## Edge Cases

```go
// Empty strings
Match("", "")           // → true  (both empty)
Match("a*", "")         // → true  (star allows zero)
Match("a", "")          // → false (requires one a)

// Exact match only
Match("abc", "abc")     // → true
Match("abc", "abcd")    // → false (extra char)
Match("abc", "ab")      // → false (missing char)

// Pattern at start only
Match("a*", "aab")      // → false (b doesn't match)
Match("a*b", "aabc")    // → false (c doesn't match)
```

## Instructions

1. Open `regex_matcher.go`
2. Implement the `Match` function using state machine logic
3. For each operator in the pattern, determine the matching state
4. Process text character-by-character, tracking state transitions
5. Run `go test -v` to verify your solution
6. Ensure all edge cases pass

## Hints

### Hint 1: Parse the Pattern

**Basic Hint:** The pattern is a sequence of "tokens" where each token is either:
- A character followed by `*` (e.g., "a*")
- A character followed by `+` (e.g., "a+")
- A single character (e.g., "a")

**Intermediate Hint:** Loop through the pattern and identify tokens by looking ahead. If pattern[i+1] is `*` or `+`, create a token with the operator. Otherwise, it's a literal.

```go
type Token struct {
    char rune
    op   string // "", "*", or "+"
}
```

**Advanced Hint:** Parse the pattern into tokens first:
```go
func parsePattern(pattern string) []Token {
    tokens := []Token{}
    runes := []rune(pattern)
    for i := 0; i < len(runes); {
        char := runes[i]
        if i+1 < len(runes) && (runes[i+1] == '*' || runes[i+1] == '+') {
            tokens = append(tokens, Token{char, string(runes[i+1])})
            i += 2
        } else {
            tokens = append(tokens, Token{char, ""})
            i++
        }
    }
    return tokens
}
```

### Hint 2: State Machine for Each Token Type

**Basic Hint:** Each token type requires different matching logic:
- **Literal**: Must match exactly one character
- **`*`**: Can match zero or more times (stay in state or advance)
- **`+`**: Must match at least once, then can match more

**Intermediate Hint:** Use an index to track position in text. For each token:
- Literal: Check if text[textIdx] == token.char, advance textIdx
- `*`: While text[textIdx] == token.char, optionally advance textIdx
- `+`: Must match once, then acts like `*`

**Advanced Hint (State Machine Logic):**

For **Kleene star (`*`)**: The state "loops" - you can stay (match more) or leave (move to next token):
```go
// a* means: consume zero or more 'a's
for textIdx < len(text) && text[textIdx] == token.char {
    textIdx++  // Stay in state, consume another 'a'
}
// When done, advance to next token (transition out of state)
```

For **Plus (`+`)**: Must consume at least one, then acts like `*`:
```go
// a+ means: consume one or more 'a's
if textIdx >= len(text) || text[textIdx] != token.char {
    return false  // Didn't match required first char
}
textIdx++  // Consume required first char

// Now act like a*
for textIdx < len(textIdx) && text[textIdx] == token.char {
    textIdx++
}
```

For **Literal**: Must match exactly:
```go
if textIdx >= len(text) || text[textIdx] != token.char {
    return false
}
textIdx++
```

### Hint 3: The Core Loop

**Basic Hint:** Loop through tokens, applying the matching logic for each token type. Track your position in the text with `textIdx`. After processing all tokens, verify you consumed all text.

**Intermediate Hint:**
```go
func Match(pattern string, text string) bool {
    tokens := parsePattern(pattern)
    textRunes := []rune(text)
    textIdx := 0

    for _, token := range tokens {
        // Apply token matching logic here
        // Update textIdx based on token type
    }

    // After all tokens processed, check if we consumed all text
    return textIdx == len(textRunes)
}
```

**Advanced Hint:** The key insight is that `*` is **greedy** - it consumes as many characters as possible. However, in this simplified implementation, we can be greedy because we're only matching from the start of the string (not searching for the pattern within the text).

### Hint 4: Common Pitfalls

1. **Forgetting to check text exhaustion**: After processing all tokens, ensure `textIdx == len(text)`. Otherwise, "a*" would match "aaab" (wrong - extra 'b').

2. **Not handling empty text with `*`**: "a*" should match "" (zero a's is valid).

3. **Confusing `+` with `*`**: Plus requires **at least one** match, star allows **zero or more**.

4. **Off-by-one errors**: When checking `i+1` in pattern parsing, ensure `i+1 < len(runes)`.

5. **Not converting to runes**: Use `[]rune(pattern)` and `[]rune(text)` for Unicode correctness.

## Think About

1. **Why is `*` called a "Kleene star"?**
   - Named after mathematician Stephen Kleene, it represents the closure of a set under concatenation (zero or more repetitions).

2. **How does this relate to regex engines like in Go's `regexp` package?**
   - Real regex engines use NFAs (Non-deterministic Finite Automata) or DFAs (Deterministic Finite Automata). Your implementation is a simplified DFA.

3. **What would it take to add `.` (match any character)?**
   - In the literal case, instead of checking equality, you'd check if the token is `.` and skip the equality check.

4. **Why does this only match from the start of the string?**
   - Full regex engines search for patterns anywhere in the text. This simplified version requires the pattern to match the entire string from start to finish.

5. **How would you add `?` (zero or one) operator?**
   - Similar to `*`, but limit the loop to at most one iteration: `if textIdx < len(text) && text[textIdx] == char { textIdx++ }`.

6. **What's the time complexity?**
   - O(n) where n is the length of the text, since we process each character once. Pattern parsing is O(m) where m is pattern length.

## What This Teaches

- **Finite automata** - How state machines recognize patterns
- **Regex internals** - How regex engines work under the hood
- **Pattern matching** - Core concept in compilers, parsers, search engines
- **Greedy algorithms** - `*` and `+` consume as much as possible
- **Tokenization** - Parsing pattern into meaningful units
- **State transitions** - Moving between matching states
- **Lookahead** - Checking next character without consuming it (in parsing)

## Common Mistakes to Avoid

1. **Not parsing the pattern first**
   ```go
   // WRONG - trying to match character-by-character without structure
   for i, char := range pattern {
       // This doesn't handle operators like * and +
   }

   // RIGHT - parse into tokens first
   tokens := parsePattern(pattern)
   for _, token := range tokens {
       // Now handle each token type
   }
   ```

2. **Forgetting the final check**
   ```go
   // BUG - doesn't verify all text was consumed
   func Match(pattern, text string) bool {
       // ... process tokens ...
       return true  // Always returns true if tokens processed!
   }

   // FIX
   return textIdx == len(textRunes)
   ```

3. **Not handling `+` correctly**
   ```go
   // BUG - treats + like *
   case "+":
       for textIdx < len(text) && text[textIdx] == token.char {
           textIdx++
       }
       // Missing: check that at least one was matched!

   // FIX - must consume at least one
   if textIdx >= len(text) || text[textIdx] != token.char {
       return false
   }
   textIdx++
   for textIdx < len(text) && text[textIdx] == token.char {
       textIdx++
   }
   ```

4. **Bounds checking in pattern parsing**
   ```go
   // BUG - can panic
   for i := 0; i < len(pattern); i++ {
       if pattern[i+1] == '*' {  // Panic if i is last index!

   // FIX
   for i := 0; i < len(pattern); i++ {
       if i+1 < len(pattern) && pattern[i+1] == '*' {
   ```

5. **Using bytes instead of runes**
   ```go
   // WRONG - breaks on Unicode
   for i := 0; i < len(pattern); i++ {
       char := pattern[i]  // byte, not rune!
   }

   // RIGHT
   runes := []rune(pattern)
   for i := 0; i < len(runes); i++ {
       char := runes[i]  // rune
   }
   ```

## Challenge Extensions (Optional)

After completing the basic version, try these:

1. **Add `.` (any character) operator**
   ```go
   Match("a.c", "abc")  // → true
   Match("a.c", "axc")  // → true
   Match("a.c", "ac")   // → false (. must match exactly one char)
   ```

2. **Add `?` (zero or one) operator**
   ```go
   Match("ab?c", "ac")   // → true
   Match("ab?c", "abc")  // → true
   Match("ab?c", "abbc") // → false
   ```

3. **Add character classes `[abc]`**
   ```go
   Match("[abc]", "a")  // → true
   Match("[abc]", "b")  // → true
   Match("[abc]", "d")  // → false
   ```

4. **Add `|` (alternation) operator**
   ```go
   Match("a|b", "a")  // → true
   Match("a|b", "b")  // → true
   ```

5. **Support searching within text (not just exact match)**
   ```go
   Contains("hello world", "wo")  // → true
   ```

## Real-World Applications

- **Compilers**: Lexical analysis (tokenization)
- **Text editors**: Find/replace, syntax highlighting
- **Validation**: Email, phone numbers, URLs
- **Log processing**: Extracting structured data from logs
- **Search engines**: Query pattern matching
- **Network security**: Intrusion detection systems (pattern matching in packets)
- **Bioinformatics**: DNA sequence matching

## After Completing

You now understand:
- How finite automata recognize patterns
- How regex engines work at a fundamental level
- The difference between `*` (zero or more) and `+` (one or more)
- How to implement state machines for pattern matching
- The concept of greedy matching
- Why regular expressions are "regular" (recognized by finite automata)

This is the foundation of computational theory that powers tools like `grep`, `sed`, `awk`, and every regex engine in modern programming languages!

---

**Ready to implement?** Open `regex_matcher.go` and look for the `TODO(human)` markers!
