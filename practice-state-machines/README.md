# State Machine Practice Exercises

**Purpose:** Quick, focused practice to solidify state machine patterns before refactoring the run-length encoding Decode function.

## Overview

These three exercises build progressively:

1. **Exercise A: Traffic Light** (15-20 min) - Basic state transitions
2. **Exercise B: Key-Value Parser** (30-40 min) - State-based parsing with delimiters
3. **Exercise C: CSV Parser** (30-40 min) - Nested states and escape handling

**Total Time:** ~90 minutes

## How to Use

### Recommended Workflow

1. **Read** the exercise markdown file
2. **Design** - Draw the state machine diagram on paper before coding
3. **Implement** - Write the code following the patterns shown
4. **Test** - Manually test with the provided examples
5. **Reflect** - Ask yourself: "What state am I in? How do I transition?"

### For Each Exercise

Create the `.go` file in this directory:

```bash
cd practice-state-machines

# Exercise A
nano traffic_light.go
# Test: go run traffic_light.go (add a main() to test)

# Exercise B
nano keyvalue_parser.go
# Test: go run keyvalue_parser.go

# Exercise C
nano csv_parser.go
# Test: go run csv_parser.go
```

## Key Patterns to Learn

### Pattern 1: Simple State Variable

```go
state := "STATE_A"

switch state {
case "STATE_A":
    // do something
    state = "STATE_B"
case "STATE_B":
    // do something else
    state = "STATE_C"
}
```

### Pattern 2: State + Buffer Accumulation

```go
buffer := ""
state := "READING_KEY"

for each character:
    if state == "READING_KEY":
        if char == delimiter:
            saveBuffer()
            state = "READING_VALUE"
        else:
            buffer += char
```

### Pattern 3: Boolean State (In/Out)

```go
inQuote := false

for each character:
    if inQuote:
        // Different parsing rules
    else:
        // Normal parsing rules
```

### Pattern 4: Manual Index Control

```go
for i := 0; i < len(input); i++ {
    char := input[i]

    // Check lookahead
    if i+1 < len(input) && input[i+1] == something {
        i++ // Skip next character
    }
}
```

## Common Mistakes to Avoid

### 1. Forgetting End-of-Input Data

```go
// ❌ WRONG - loses last item
for i := 0; i < len(s); i++ {
    if s[i] == delimiter {
        save(buffer)
        buffer = ""
    } else {
        buffer += string(s[i])
    }
}
return result

// ✓ RIGHT - saves last item
for i := 0; i < len(s); i++ {
    // ... same loop
}
if buffer != "" {
    save(buffer) // Don't forget!
}
return result
```

### 2. Missing Bounds Checks

```go
// ❌ WRONG - can panic
if input[i+1] == '"' {
    // ...
}

// ✓ RIGHT - bounds check first
if i+1 < len(input) && input[i+1] == '"' {
    // ...
}
```

### 3. Not Resetting State

```go
// ❌ WRONG - doesn't reset
if delimiter:
    save(buffer)
    state = "NEW_STATE"
    // buffer still has old data!

// ✓ RIGHT - reset everything
if delimiter:
    save(buffer)
    buffer = ""      // Reset!
    state = "NEW_STATE"
```

## After Completing All Three

### Self-Refactor Challenge

Return to the run-length encoding exercise:

1. Open `run_length_encoding.go`
2. **Delete** the Decode function (keep a copy commented out)
3. **Rewrite** Decode yourself using the state machine pattern
4. **Test** - verify `go test` passes
5. **Compare** your version to the reference

### Reflection Questions

1. How is the CSV parser's `inQuote` state similar to the RLE Decode escape handling?
2. Why do we use `for i := 0; i < len(); i++` instead of `for _, char := range`?
3. What would break if you forgot to save the last buffer after the loop?
4. How would you debug if a state machine isn't working? (Hint: print state at each step!)

## Connection to Full Module

These exercises are simplified versions of what you'll see in the full **Algorithms & State Machines** module:

- **Module Exercise 01:** Traffic Light (extended with timers)
- **Module Exercise 04:** CSV Parser (extended with multiple lines, validation)
- **Module Exercise 05:** Expression Evaluator (more complex state machine)

By practicing here first, you'll have a strong foundation for the full module!

## Need Help?

If you get stuck:

1. **Draw the states** on paper - visualize the flow
2. **Trace by hand** - what should happen for "a,b,c"?
3. **Print state** - add `fmt.Printf("State: %s, Char: %c\n", state, char)`
4. **Check bounds** - verify all array accesses are safe

Remember: State machines are about **sequential thinking**, not complex nested logic. One state at a time!

---

**Next Step:** Complete Exercise A, then B, then C. Then return to refactor the Decode function!
