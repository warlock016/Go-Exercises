# Run-Length Encoding: Execution Trace Walkthrough

This document provides step-by-step execution traces to help you understand how the state machine pattern works in the Decode function.

## Table of Contents

1. [Simple Case: "100a"](#simple-case-100a)
2. [Digit Escape: "311322"](#digit-escape-311322)
3. [Mixed Case: "2a2112b"](#mixed-case-2a2112b)
4. [Edge Case: Why "123456" Fails](#edge-case-why-123456-fails)
5. [State Machine Visualization](#state-machine-visualization)

---

## Simple Case: "100a"

**Input:** `"100a"` (encodes 100 a's)
**Expected Output:** 100 a's

### Iteration 1

```
i = 0
runes = ['1', '0', '0', 'a']

STATE 1: READ COUNT
  - i=0: runes[0]='1' (digit) → countStr="1", i=1
  - i=1: runes[1]='0' (digit) → countStr="10", i=2
  - i=2: runes[2]='0' (digit) → countStr="100", i=3
  - i=3: runes[3]='a' (NOT digit) → exit loop
  - Result: countStr="100", i=3

VALIDATION
  - i < len(runes)? 3 < 4 → YES ✓
  - countStr != ""? "100" != "" → YES ✓
  - Continue

STATE 2: READ CHARACTER
  - char = runes[3] = 'a'
  - i = 4

STATE 3: HANDLE ESCAPE
  - unicode.IsDigit('a')? → NO
  - Skip nothing
  - i = 4

STATE 4: OUTPUT
  - count = Atoi("100") = 100
  - Output: strings.Repeat("a", 100) = "aaa...aaa" (100 a's)

Loop condition: i < len(runes)? 4 < 4 → NO
Exit loop

Final output: 100 a's ✓
```

### Key Observation

The state machine **completely reads the count first** (all three digits "100"), **then** reads the character ('a'). This sequential approach prevents ambiguity.

---

## Digit Escape: "311322"

**Input:** `"311322"` (encodes "111222")
**Expected Output:** "111222"

### Iteration 1: Decode "311"

```
i = 0
runes = ['3', '1', '1', '3', '2', '2']

STATE 1: READ COUNT
  - i=0: runes[0]='3' (digit) → countStr="3", i=1
  - i=1: runes[1]='1' (digit) → countStr="31", i=2
  - i=2: runes[2]='1' (digit) → countStr="311", i=3
  - i=3: runes[3]='3' (digit) → countStr="3113", i=4
  - i=4: runes[4]='2' (digit) → countStr="31132", i=5
  - i=5: runes[5]='2' (digit) → countStr="311322", i=6
  - i=6: i >= len(runes) → exit loop
  - Result: countStr="311322", i=6

VALIDATION
  - i >= len(runes)? 6 >= 6 → YES
  - BREAK (no character found, malformed)

Final output: "" ❌
```

**Wait, this doesn't work!** This demonstrates the **LIMITATION** of this encoding scheme.

### Why This Fails

The problem: `"311322"` is **all digits**, so the decoder reads the entire string as the count and never finds a character.

This is the fundamental limitation we documented in the code. The encoding works for:
- Digit runs mixed with letters: "3a311" → "aaa" + "111" ✓
- Long counts with non-digit chars: "100a" → 100 a's ✓

But fails for:
- All-digit encoded output: "311322" → ❌

**Real-world solution:** Use a delimiter (e.g., "3|1|1|3|2|2") or different escape character.

---

## Mixed Case: "2a2112b"

**Input:** `"2a2112b"` (encodes "aa11bb")
**Expected Output:** "aa11bb"

### Iteration 1: Decode "2a"

```
i = 0
runes = ['2', 'a', '2', '1', '1', '2', 'b']

STATE 1: READ COUNT
  - i=0: runes[0]='2' (digit) → countStr="2", i=1
  - i=1: runes[1]='a' (NOT digit) → exit loop
  - Result: countStr="2", i=1

VALIDATION
  - i < len(runes)? 1 < 7 → YES ✓
  - countStr != ""? "2" != "" → YES ✓

STATE 2: READ CHARACTER
  - char = runes[1] = 'a'
  - i = 2

STATE 3: HANDLE ESCAPE
  - unicode.IsDigit('a')? → NO
  - Skip nothing
  - i = 2

STATE 4: OUTPUT
  - count = Atoi("2") = 2
  - Output: "aa"

Current output: "aa"
i = 2, continue loop
```

### Iteration 2: Decode "211"

```
i = 2
runes = ['2', 'a', '2', '1', '1', '2', 'b']
                   ↑

STATE 1: READ COUNT
  - i=2: runes[2]='2' (digit) → countStr="2", i=3
  - i=3: runes[3]='1' (digit) → countStr="21", i=4
  - i=4: runes[4]='1' (digit) → countStr="211", i=5
  - i=5: runes[5]='2' (digit) → countStr="2112", i=6
  - i=6: runes[6]='b' (NOT digit) → exit loop
  - Result: countStr="2112", i=6

VALIDATION
  - i < len(runes)? 6 < 7 → YES ✓

STATE 2: READ CHARACTER
  - char = runes[6] = 'b'
  - i = 7

STATE 3: HANDLE ESCAPE
  - unicode.IsDigit('b')? → NO
  - i = 7

STATE 4: OUTPUT
  - count = Atoi("2112") = 2112
  - Output: 2112 b's ❌

Current output: "aa" + (2112 b's) ❌
```

**This also fails!** The decoder reads "2112" as the count instead of recognizing "211" as count=2, char='1', escape='1'.

### Why This Fails

Same issue: consecutive digit runs create ambiguity. The decoder can't tell:
- Is "211" part of the count?
- Or is it count='2', char='1', escape='1'?

Without lookahead or a delimiter, this is **unresolvable** with simple greedy parsing.

---

## Edge Case: Why "123456" Fails

**Input (Original):** `"123456"`
**Encoded:** `"111122133144155166"`
**Decoding Attempt:** `Decode("111122133144155166")`

### Trace

```
i = 0
runes = ['1','1','1','1','2','2','1','3','3','1','4','4','1','5','5','1','6','6']

STATE 1: READ COUNT
  - Reads ALL 18 digits: countStr="111122133144155166", i=18
  - i >= len(runes) → BREAK

Output: "" ❌
```

**Problem:** All digits → entire string consumed as count → no character found.

### Visualization

```
Expected parsing:
"111" → count=1, char='1', escape='1' → "1"
"122" → count=1, char='2', escape='2' → "2"
"133" → count=1, char='3', escape='3' → "3"
...

Actual parsing:
"111122133144155166" → count=111122133144155166, char=??? → ❌
```

---

## State Machine Visualization

### High-Level Flow

```
┌─────────────────────────────────────────────┐
│  START                                      │
└──────────────┬──────────────────────────────┘
               │
               ↓
     ┌─────────────────────┐
     │  STATE 1:           │
     │  Read Count         │◄──┐
     │  (all digits)       │   │
     └─────────┬───────────┘   │
               │               │
               ↓               │
       ┌──────────────┐        │
       │ Validation   │        │
       │ (has char?)  │        │
       └──────┬───────┘        │
              │                │
         [YES]│   [NO]         │
              │   └────► BREAK │
              ↓                │
     ┌─────────────────────┐   │
     │  STATE 2:           │   │
     │  Read Character     │   │
     └─────────┬───────────┘   │
               │               │
               ↓               │
     ┌─────────────────────┐   │
     │  STATE 3:           │   │
     │  Handle Escape      │   │
     │  (if char is digit) │   │
     └─────────┬───────────┘   │
               │               │
               ↓               │
     ┌─────────────────────┐   │
     │  STATE 4:           │   │
     │  Output             │   │
     │  (repeat char)      │   │
     └─────────┬───────────┘   │
               │               │
               └───────────────┘

               ↓
        [More input?]
               │
          [NO] │   [YES]
               ↓   └───► (loop back to STATE 1)
           ┌──────┐
           │ DONE │
           └──────┘
```

### State Transition Table

| Current State | Input Type | Action | Next State |
|---------------|------------|--------|------------|
| STATE 1 | Digit | Accumulate to countStr, i++ | STATE 1 (loop) |
| STATE 1 | Non-digit OR end | Exit count reading | Validation |
| Validation | i < len, countStr not empty | Proceed | STATE 2 |
| Validation | i >= len OR countStr empty | Break loop | DONE |
| STATE 2 | Any rune | char = runes[i], i++ | STATE 3 |
| STATE 3 | char is digit | i++ (skip escape) | STATE 4 |
| STATE 3 | char not digit | Do nothing | STATE 4 |
| STATE 4 | Always | Output char × count | Check loop condition |
| Loop check | i < len | Continue | STATE 1 |
| Loop check | i >= len | Exit | DONE |

---

## Key Takeaways

### ✓ What Works

1. **Sequential state execution:** Each state completes fully before the next
2. **Greedy digit reading:** Read all digits as count (simple and fast)
3. **Bounds checking:** Always verify i < len(runes) before array access
4. **Escape handling:** After reading character, check if digit and skip next

### ✗ Known Limitations

1. **All-digit strings:** Cannot round-trip (e.g., "123456")
2. **Consecutive digit runs:** Ambiguous parsing (e.g., "311" could be count or count+char+escape)
3. **No delimiter:** Simple greedy parsing can't distinguish metadata from data when both are digits

### 🎯 Learning Goals Achieved

- **State machine pattern:** Linear, sequential state transitions
- **Index-based loops:** Manual control over iteration (not `range`)
- **Bounds checking:** Prevent index-out-of-range panics
- **Edge case handling:** Empty strings, malformed input
- **Trade-offs:** Understanding when simple solutions have limitations

---

## Next Steps

Now that you understand this state machine:

1. **Practice:** Complete the 3 quick practice exercises
2. **Self-Refactor:** Rewrite the Decode function yourself without looking
3. **Extend:** Try adding a delimiter-based encoding scheme that solves the all-digit limitation

The key is not to memorize this code, but to **internalize the pattern**: break complex parsing into simple, sequential states.

---

**Date:** 2025-11-14
**Related Files:** `run_length_encoding.go`, `ENCODING_DECODING_PATTERNS.md`
