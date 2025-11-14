# Module: Algorithms & State Machines

**Module Number:** 00.7 (Foundational - Custom Remediation)
**Focus:** Algorithmic thinking, state machines, common patterns
**Prerequisites:** String Mastery (00.5)
**Estimated Time:** 8-12 hours
**Difficulty:** Progressive (Easy → Medium → Hard)

---

## Why This Module Exists

During Exercise 10 (Run-Length Encoding) of String Mastery, you encountered challenges with:
- State machine design (sequential state processing)
- Parsing algorithms (when to read vs when to transition)
- Index management (manual control vs range loops)
- Bounds checking (preventing index out of range errors)

This module systematically builds the algorithmic foundation you need to solve complex problems confidently.

---

## Learning Objectives

By the end of this module, you will:

1. **Design state machines** from problem descriptions
2. **Implement parsing algorithms** for structured data
3. **Apply common algorithmic patterns** (two pointers, sliding window, frequency counting)
4. **Debug systematically** using step-by-step tracing
5. **Recognize patterns** across different problem domains
6. **Write efficient code** with proper time/space complexity

---

## Module Structure

### Tier 1: State Machine Fundamentals (Exercises 1-3)
**Focus:** Basic state transitions and tracking

- 01_traffic_light - Simple state cycling
- 02_string_tokenizer - State-based word extraction
- 03_parentheses_validator - Stack-based state machine

### Tier 2: Parsing Patterns (Exercises 4-6)
**Focus:** Structured data parsing

- 04_csv_parser - Quote-aware delimiter parsing
- 05_expression_evaluator - Infix arithmetic with precedence
- 06_json_lite_parser - Nested object/array parsing

### Tier 3: Core Algorithms (Exercises 7-10)
**Focus:** Common algorithmic patterns

- 07_two_pointers - Remove duplicates from sorted array
- 08_sliding_window - Maximum sum subarray
- 09_frequency_counter - Top K frequent elements
- 10_backtracking_permutations - Generate all permutations

### Tier 4: Advanced State Machines (Exercises 11-12)
**Focus:** Complex finite automata

- 11_regex_matcher - Simple pattern matching (a*, a+b, etc.)
- 12_protocol_parser - HTTP request parsing

---

## How to Use This Module

### For Each Exercise

1. **Read** the README.md - Understand the problem
2. **Design** - Draw state diagram or algorithm flow on paper
3. **Implement** - Write code in the .go file (look for TODO(human))
4. **Test** - Run `go test -v` to verify correctness
5. **Reflect** - Write EXPLANATION.md explaining your approach

### Learning Resources

Before starting, read these guides (in `/resources/`):

- **STATE_MACHINE_GUIDE.md** - Visual diagrams, patterns, examples
- **PARSING_TECHNIQUES.md** - Greedy, lookahead, recursive descent
- **DEBUGGING_ALGORITHMS.md** - Systematic troubleshooting
- **ALGORITHM_PATTERNS.md** - 15+ common patterns explained

### Practice Exercises

If you haven't already, complete the quick practice exercises in `/practice-state-machines/`:

- Exercise A: Traffic Light (15 min)
- Exercise B: Key-Value Parser (30 min)
- Exercise C: CSV Parser (30 min)

These warm-up exercises prepare you for this module.

---

## Difficulty Progression

| Exercise | Estimated Time | Concepts | Difficulty |
|----------|---------------|----------|------------|
| 01 | 20-30 min | Simple state transitions | Easy |
| 02 | 30-40 min | Character-by-character parsing | Easy-Medium |
| 03 | 30-40 min | Stack data structure | Medium |
| 04 | 40-50 min | Quote escape handling | Medium |
| 05 | 50-60 min | Recursive descent, precedence | Medium-Hard |
| 06 | 60-70 min | Nested structures, recursion | Hard |
| 07 | 30-40 min | Two-pointer technique | Medium |
| 08 | 30-40 min | Sliding window pattern | Medium |
| 09 | 40-50 min | Hash maps, sorting | Medium |
| 10 | 50-60 min | Recursion, backtracking | Hard |
| 11 | 60-70 min | Finite automata, regex | Hard |
| 12 | 60-70 min | Multi-state parsing | Hard |

**Total:** 8-10 hours

---

## Success Criteria

You've mastered this module when you can:

- [ ] Design a state machine diagram before coding
- [ ] Implement parsers for structured formats (CSV, JSON-like)
- [ ] Choose the right loop type (range vs index-based)
- [ ] Handle edge cases (empty input, single element, malformed data)
- [ ] Apply algorithmic patterns (two pointers, sliding window, etc.)
- [ ] Debug using systematic tracing (not random changes)
- [ ] Explain time/space complexity of your solutions

---

## Common Patterns You'll Learn

### 1. Sequential State Processing

```go
for i := 0; i < len(input); {
    // STATE 1: Read something
    while condition {
        accumulate()
        i++
    }

    // STATE 2: Process something
    process()
    i++

    // STATE 3: Output
    output()
}
```

### 2. Stack-Based Validation

```go
stack := []rune{}
for _, char := range input {
    if isOpening(char) {
        stack = append(stack, char)
    } else if isClosing(char) {
        if len(stack) == 0 {
            return false
        }
        stack = stack[:len(stack)-1]
    }
}
return len(stack) == 0
```

### 3. Two Pointers

```go
left, right := 0, 0
for right < len(arr) {
    if arr[right] != arr[left] {
        left++
        arr[left] = arr[right]
    }
    right++
}
return arr[:left+1]
```

### 4. Sliding Window

```go
windowSum := 0
maxSum := 0
for i := 0; i < len(arr); i++ {
    windowSum += arr[i]
    if i >= k-1 {
        maxSum = max(maxSum, windowSum)
        windowSum -= arr[i-k+1]
    }
}
```

### 5. Frequency Map

```go
freq := make(map[rune]int)
for _, char := range s {
    freq[char]++
}
// Find max frequency
maxFreq := 0
for _, count := range freq {
    if count > maxFreq {
        maxFreq = count
    }
}
```

---

## Testing Strategy

Each exercise has comprehensive tests:

```bash
# Run all tests in module
cd 00.7-algorithms-and-state-machines
go test ./...

# Run specific exercise
cd 01_traffic_light
go test -v

# Run with coverage
go test -cover

# Run specific test
go test -run=TestNext
```

---

## Tips for Success

### Before Coding

1. **Read the problem twice** - Ensure you understand
2. **Identify states/phases** - What are the distinct modes?
3. **Draw a diagram** - State machine or algorithm flow
4. **List edge cases** - Empty, single, malformed, etc.

### While Coding

1. **Write incrementally** - Implement one state at a time
2. **Test early** - Don't write all code before testing
3. **Add strategic prints** - Log state transitions
4. **Check bounds** - Always verify array access

### When Stuck

1. **Trace by hand** - Write state at each step on paper
2. **Check the guides** - STATE_MACHINE_GUIDE.md, DEBUGGING_ALGORITHMS.md
3. **Simplify** - Start with simpler input, build up
4. **Take a break** - Fresh perspective helps

---

## Connection to Real-World

These patterns appear in:

- **Compilers:** Lexers and parsers (Exercises 2, 5, 6, 11)
- **Network protocols:** HTTP, JSON, XML parsing (Exercise 12)
- **Data processing:** CSV, log file parsing (Exercise 4)
- **Algorithms:** Optimization problems (Exercises 7-10)
- **Text editors:** Syntax highlighting (state machines)
- **Game development:** AI state machines (Exercise 1)

---

## After This Module

You'll be ready for:

- **Testing Fundamentals** (00.6) - Writing robust tests
- **Core Concepts** (01-11) - Standard Go curriculum
- **Real projects** - Implementing parsers, algorithms, tools

---

## Acknowledgments

This module was created in response to struggles with run-length encoding (Exercise 00.5-10). Every exercise here addresses a specific challenge encountered in that exercise:

- **State machine confusion** → Exercises 1-3
- **Parsing ambiguity** → Exercises 4-6
- **Algorithmic thinking** → Exercises 7-10
- **Complex state management** → Exercises 11-12

---

## Quick Reference

**Key Files:**
- `/resources/STATE_MACHINE_GUIDE.md` - Theory and patterns
- `/resources/PARSING_TECHNIQUES.md` - Parsing strategies
- `/resources/DEBUGGING_ALGORITHMS.md` - Troubleshooting
- `/resources/ALGORITHM_PATTERNS.md` - Common patterns

**Practice First:**
- `/practice-state-machines/` - 3 warm-up exercises

**Module Exercises:**
- `01_traffic_light/` through `12_protocol_parser/`

---

**Let's build strong algorithmic foundations! 🚀**

**Recommended pace:** 2-3 exercises per day, complete module in 5-7 days.

**Remember:** Quality > Speed. Understanding the patterns matters more than finishing quickly.
