# Debugging Algorithms: A Systematic Approach

## Table of Contents

1. [The Debugging Mindset](#the-debugging-mindset)
2. [Before You Debug](#before-you-debug)
3. [Debugging Techniques](#debugging-techniques)
4. [Common Bug Patterns](#common-bug-patterns)
5. [Step-by-Step Tracing](#step-by-step-tracing)
6. [Debugging Tools in Go](#debugging-tools-in-go)
7. [Real-World Debugging Examples](#real-world-debugging-examples)
8. [Prevention > Cure](#prevention--cure)

---

## The Debugging Mindset

### What Debugging Is NOT

- ❌ Random code changes hoping something works
- ❌ Adding print statements everywhere without a plan
- ❌ Blaming the compiler/language/libraries
- ❌ Starting over from scratch

### What Debugging IS

- ✅ **Scientific method:** Hypothesis → Test → Refine
- ✅ **Detective work:** Follow evidence to root cause
- ✅ **Systematic elimination:** Remove possibilities one by one
- ✅ **Understanding:** Learn WHY the bug exists

### The Debugging Loop

```
1. Observe unexpected behavior
2. Form hypothesis about cause
3. Design experiment to test hypothesis
4. Run experiment (add logging, use debugger)
5. Analyze results
6. If hypothesis wrong → go to step 2
7. If hypothesis right → fix and verify
```

---

## Before You Debug

### Step 1: Reproduce the Bug

**Can't fix what you can't reproduce.**

```go
// ❌ "It sometimes crashes" - hard to debug
// ✅ "It crashes when input is empty string" - easy to debug

func TestReproduceBug(t *testing.T) {
    // Minimal test case that triggers the bug
    result := BuggyFunction("")
    // Should not panic
}
```

**Tips:**
- Find the **minimal input** that triggers the bug
- Document the **exact steps** to reproduce
- Note any **environmental factors** (OS, Go version, etc.)

---

### Step 2: Understand Expected Behavior

**What SHOULD happen vs what DOES happen?**

```
Input: "100a"
Expected: 100 a's
Actual: Empty string

Gap: Why is output empty?
```

Write down:
1. Input
2. Expected output
3. Actual output
4. The difference

---

### Step 3: Isolate the Problem

**Is it the whole function or just one part?**

```go
func Decode(input string) string {
    // Part 1: Parse input
    count, char := parse(input)

    // Part 2: Generate output
    return repeat(char, count)
}

// Test each part separately
func TestParse(t *testing.T) {
    count, char := parse("100a")
    assert(count == 100)  // ✓ or ✗?
    assert(char == 'a')   // ✓ or ✗?
}

func TestRepeat(t *testing.T) {
    result := repeat('a', 100)
    assert(len(result) == 100)  // ✓ or ✗?
}
```

---

## Debugging Techniques

### Technique 1: Printf Debugging

**Strategic print statements.**

```go
func Decode(s string) string {
    fmt.Printf("=== Decode START ===\n")
    fmt.Printf("Input: %q (len=%d)\n", s, len(s))

    runes := []rune(s)

    for i := 0; i < len(runes); {
        fmt.Printf("\n--- Iteration: i=%d ---\n", i)

        // STATE 1: Read count
        countStr := ""
        for i < len(runes) && unicode.IsDigit(runes[i]) {
            fmt.Printf("  Reading digit: %c\n", runes[i])
            countStr += string(runes[i])
            i++
        }
        fmt.Printf("  Count string: %q, i=%d\n", countStr, i)

        if i >= len(runes) {
            fmt.Printf("  BREAK: i=%d >= len=%d\n", i, len(runes))
            break
        }

        // STATE 2: Read character
        char := runes[i]
        fmt.Printf("  Character: %c\n", char)
        i++

        // ... rest
    }

    fmt.Printf("=== Decode END ===\n")
    return output.String()
}
```

**Tips:**
- Use **headers** (=== START ===) to segment output
- Print **state** at each step
- Print **values** of key variables
- Print **control flow** (which branch executed)

---

### Technique 2: Binary Search Debugging

**Comment out half the code, see if bug remains.**

```go
func BuggyFunction(input string) string {
    // Part A: Parse input
    parsed := parse(input)

    // Part B: Transform
    // transformed := transform(parsed)

    // Part C: Output
    // return format(transformed)

    return format(parsed)  // Skip Part B
}

// If bug gone → bug is in Part B
// If bug remains → bug is in Part A or C
```

**Repeat** until you narrow down to single line.

---

### Technique 3: Rubber Duck Debugging

**Explain code line-by-line to someone (or something).**

```
"Okay, so first I read all the digits..."
"Wait, ALL the digits? Even when there's a character mixed in?"
"Oh! That's the bug - I'm reading '3', '1', '1', '3', '2', '2' all as one number!"
```

**Why it works:** Forcing verbalization exposes assumptions.

---

### Technique 4: Diff Against Working Code

**Compare your code to reference implementation.**

```diff
  func Decode(s string) string {
      for i := 0; i < len(runes); {
          countStr := ""
-         for unicode.IsDigit(runes[i]) {  // ❌ Your code
+         for i < len(runes) && unicode.IsDigit(runes[i]) {  // ✓ Reference
              countStr += string(runes[i])
              i++
          }
      }
  }
```

**Find:** Missing bounds check!

---

### Technique 5: Reverse Engineering

**Start from the output and work backwards.**

```
Actual output: ""
Why empty?

→ Output is built by appending to strings.Builder
→ Check: are we ever calling output.WriteString()?
→ Print before WriteString: never reached
→ Why? Loop exits early
→ Check loop condition: i >= len(runes) immediately
→ Why? countStr consumed all runes
→ Found root cause!
```

---

## Common Bug Patterns

### Pattern 1: Off-By-One Errors

```go
// ❌ BUG: Misses last element
for i := 0; i < len(arr)-1; i++ {
    process(arr[i])
}

// ✓ FIX
for i := 0; i < len(arr); i++ {
    process(arr[i])
}
```

**How to catch:**
- Test with **single element** (array of length 1)
- Test with **two elements** (minimal non-trivial case)
- Print **loop bounds**

---

### Pattern 2: Uninitialized Variables

```go
// ❌ BUG: count never set
var count int
if condition {
    count = getValue()
}
// count might be 0!

// ✓ FIX: Initialize
count := 0
if condition {
    count = getValue()
}
```

**How to catch:**
- Print variable values **immediately after declaration**
- Use **named return values** in Go (auto-initialized)

---

### Pattern 3: Wrong Loop Type

```go
// ❌ BUG: Can't skip ahead with range
for _, r := range slice {
    if r == delimiter && slice[i+1] == delimiter {
        // Can't access i+1, no 'i' variable!
    }
}

// ✓ FIX: Use index loop
for i := 0; i < len(slice); i++ {
    if slice[i] == delimiter && i+1 < len(slice) && slice[i+1] == delimiter {
        i++ // Can skip
    }
}
```

**How to catch:**
- If you need to **skip elements**, use index loop
- If you need **lookahead**, use index loop

---

### Pattern 4: Missing Bounds Check

```go
// ❌ BUG: Panic on short input
if input[i+1] == something {  // Crash if i+1 >= len

// ✓ FIX: Check bounds first
if i+1 < len(input) && input[i+1] == something {
```

**How to catch:**
- Test with **empty input**
- Test with **single element**
- **Always** check bounds before array access

---

### Pattern 5: Not Handling End-of-Input

```go
// ❌ BUG: Loses last item
for i := 0; i < len(input); i++ {
    if input[i] == delimiter {
        save(buffer)
        buffer = ""
    } else {
        buffer += string(input[i])
    }
}
// buffer still has last item!

// ✓ FIX: Save after loop
for i := 0; i < len(input); i++ {
    // ...
}
if buffer != "" {
    save(buffer)
}
```

**How to catch:**
- Trace **what happens after loop ends**
- Check **buffer contents** after loop

---

### Pattern 6: Wrong Data Type

```go
// ❌ BUG: Byte vs Rune
s := "café"
for i := 0; i < len(s); i++ {
    char := s[i]  // byte, not rune!
    // char is 'Ã' at position 3, not 'é'
}

// ✓ FIX: Use runes for Unicode
for _, r := range s {  // r is rune
    char := r
}
```

**How to catch:**
- Test with **non-ASCII input** (emojis, accents)
- Use `[]rune()` conversion when needed

---

### Pattern 7: Forgetting to Reset State

```go
// ❌ BUG: State carries over
buffer := ""
for ... {
    if delimiter {
        save(buffer)
        state = "NEW_STATE"
        // buffer still has old data!
    }
}

// ✓ FIX: Reset everything
if delimiter {
    save(buffer)
    buffer = ""  // Reset!
    state = "NEW_STATE"
}
```

---

## Step-by-Step Tracing

### Manual Trace Template

```
Function: Decode("100a")

Initial state:
  input = "100a"
  runes = ['1', '0', '0', 'a']
  output = ""
  i = 0

Iteration 1:
  STATE 1: Read count
    i=0: runes[0]='1', digit? YES → countStr="1", i=1
    i=1: runes[1]='0', digit? YES → countStr="10", i=2
    i=2: runes[2]='0', digit? YES → countStr="100", i=3
    i=3: runes[3]='a', digit? NO → exit loop
    Result: countStr="100", i=3

  Validation:
    i < len(runes)? 3 < 4 → YES ✓

  STATE 2: Read character
    char = runes[3] = 'a'
    i = 4

  STATE 3: Handle escape
    IsDigit('a')? NO
    Skip nothing

  STATE 4: Output
    count = Atoi("100") = 100
    output += strings.Repeat("a", 100)

  Loop check:
    i < len(runes)? 4 < 4 → NO
    Exit loop

Final state:
  output = "aaa..." (100 a's) ✓
```

### Trace Checklist

- [ ] List initial values
- [ ] Track variable changes at each step
- [ ] Note control flow decisions (if/else taken)
- [ ] Mark loop iterations
- [ ] Verify expectations at each step
- [ ] Compare final state to expected

---

## Debugging Tools in Go

### Tool 1: `go test -v`

**Verbose output shows which tests fail.**

```bash
go test -v
=== RUN   TestDecode
=== RUN   TestDecode/Empty_string
=== RUN   TestDecode/Simple_case
    decode_test.go:25: Decode("100a") = "", want "aaa...aaa" (100 a's)
--- FAIL: TestDecode (0.00s)
    --- PASS: TestDecode/Empty_string (0.00s)
    --- FAIL: TestDecode/Simple_case (0.00s)
```

---

### Tool 2: `go test -run=<pattern>`

**Run specific test.**

```bash
go test -run="TestDecode/Simple_case"
```

---

### Tool 3: `t.Logf()` in Tests

**Print from test without failing.**

```go
func TestDebug(t *testing.T) {
    result := Decode("100a")

    t.Logf("Result: %q", result)
    t.Logf("Length: %d", len(result))

    if result != expected {
        t.Errorf("Expected %q, got %q", expected, result)
    }
}
```

---

### Tool 4: Debugger (Delve)

**Step through code line by line.**

```bash
# Install
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug test
dlv test
(dlv) break Decode
(dlv) continue
(dlv) print i
(dlv) print countStr
(dlv) next  # Step to next line
(dlv) step  # Step into function
```

---

### Tool 5: `go vet` and `golint`

**Catch common mistakes.**

```bash
go vet ./...
golangci-lint run
```

---

## Real-World Debugging Examples

### Example 1: Infinite Loop

**Symptom:** Program hangs.

**Debug process:**
```go
func buggy(input string) {
    i := 0
    for i < len(input) {
        if input[i] == 'a' {
            // Forgot to increment i!
        }
    }
}

// Add print
for i < len(input) {
    fmt.Printf("i=%d\n", i)  // Prints "i=0" forever
    if input[i] == 'a' {
        // Aha! i never changes
    }
}

// Fix
for i < len(input) {
    if input[i] == 'a' {
        // ...
    }
    i++  // Add increment
}
```

---

### Example 2: Panic: Index Out of Range

**Symptom:** Runtime panic.

```
panic: runtime error: index out of range [5] with length 4
```

**Debug process:**

```go
// Add bounds checking prints
func buggy(input []int) {
    for i := 0; i <= len(input); i++ {  // ❌ Should be <, not <=
        fmt.Printf("Accessing input[%d], len=%d\n", i, len(input))
        process(input[i])
    }
}

// Output shows:
// Accessing input[0], len=4
// Accessing input[1], len=4
// Accessing input[2], len=4
// Accessing input[3], len=4
// Accessing input[4], len=4  ← Aha! i=4 >= len=4

// Fix
for i := 0; i < len(input); i++ {  // ✓
```

---

### Example 3: Wrong Output

**Symptom:** Decode("3a") returns "aaaaaa" (6 a's) instead of "aaa" (3 a's).

**Debug process:**

```go
func Decode(s string) string {
    fmt.Printf("Input: %q\n", s)

    // ... parsing ...

    count, _ := strconv.Atoi(countStr)
    fmt.Printf("Count: %d\n", count)  // Prints: Count: 3 ✓

    result := strings.Repeat(string(char), count)
    fmt.Printf("Result: %q (len=%d)\n", result, len(result))
    // Prints: Result: "aaaaaa" (len=6) ← Wait, len is 6 but count is 3?

    // Check char
    fmt.Printf("Char: %c (as string: %q)\n", char, string(char))
    // Prints: Char: a (as string: "aa") ← AHA! char is TWO a's!

    // Found bug: char includes the escape character
}
```

---

## Prevention > Cure

### Write Tests First

```go
// Write the test BEFORE implementing
func TestDecode(t *testing.T) {
    tests := []struct {
        input string
        want  string
    }{
        {"", ""},           // Edge: empty
        {"1a", "a"},        // Edge: single
        {"100a", strings.Repeat("a", 100)},  // Normal
        {"311", "111"},     // Digit escape
    }

    for _, tt := range tests {
        got := Decode(tt.input)
        if got != tt.want {
            t.Errorf("Decode(%q) = %q, want %q", tt.input, got, tt.want)
        }
    }
}
```

---

### Validate Assumptions

```go
func Decode(s string) string {
    // Assumption: Input is well-formed
    // Validate: Check this assumption
    if !isWellFormed(s) {
        return "" // Or panic/error
    }

    // Assumption: Count fits in int
    if len(countStr) > 10 {  // Max int is ~10 digits
        // Handle overflow
    }
}
```

---

### Add Assertions

```go
func Decode(s string) string {
    // ...

    // Assertion: i should never exceed length
    if i > len(runes) {
        panic(fmt.Sprintf("BUG: i=%d > len=%d", i, len(runes)))
    }

    // Assertion: count should be positive
    if count < 0 {
        panic(fmt.Sprintf("BUG: negative count: %d", count))
    }
}
```

---

### Document Invariants

```go
// Invariant: After reading count, i points to character
// Invariant: countStr is never empty when we reach character reading
// Invariant: output only grows (never shrinks)

func Decode(s string) string {
    // ... implementation that maintains these invariants
}
```

---

## Key Takeaways

1. **Debugging is a skill** - gets better with practice
2. **Reproduce first** - can't fix what you can't reproduce
3. **Use scientific method** - hypothesis, test, refine
4. **Strategic prints** - not random - state, values, control flow
5. **Isolate the problem** - test each part separately
6. **Trace by hand** - pen and paper, step by step
7. **Prevention beats cure** - tests, assertions, validation
8. **Learn from bugs** - each bug teaches a pattern to avoid

---

## Debugging Checklist

When stuck:

- [ ] Can I reproduce the bug consistently?
- [ ] Do I know the expected vs actual behavior?
- [ ] Have I added strategic print statements?
- [ ] Have I traced the code by hand?
- [ ] Have I tested with edge cases (empty, single element)?
- [ ] Have I checked bounds on all array accesses?
- [ ] Have I verified loop conditions (<= vs <)?
- [ ] Have I checked variable initialization?
- [ ] Have I verified state resets correctly?
- [ ] Have I rubber-duck explained the code?

---

## Further Reading

- **"The Pragmatic Programmer"** - Debugging chapter
- **"Debugging: The 9 Indispensable Rules"** by David J. Agans
- **"Why Programs Fail"** by Andreas Zeller
- **Go blog:** [Debugging with Delve](https://go.dev/blog/using-go-modules)

---

## Next Steps

1. Practice with the 3 quick exercises
2. When stuck, use this guide systematically
3. Keep a "bug journal" - record bugs you find and how you fixed them
4. Review your journal weekly - spot patterns

**Remember:** Every bug is a learning opportunity. The more bugs you fix, the faster you'll recognize patterns!

---

**Version:** 1.0
**Last Updated:** 2025-11-14
**Related:** STATE_MACHINE_GUIDE.md, PARSING_TECHNIQUES.md, ALGORITHM_PATTERNS.md
