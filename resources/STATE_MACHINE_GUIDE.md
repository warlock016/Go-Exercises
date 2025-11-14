# State Machine Guide: From Theory to Practice

## Table of Contents

1. [What is a State Machine?](#what-is-a-state-machine)
2. [Why State Machines Matter](#why-state-machines-matter)
3. [Visual Representations](#visual-representations)
4. [Common State Machine Patterns](#common-state-machine-patterns)
5. [Implementation Strategies in Go](#implementation-strategies-in-go)
6. [State Machine Design Process](#state-machine-design-process)
7. [Real-World Examples](#real-world-examples)
8. [Common Pitfalls](#common-pitfalls)
9. [Advanced Patterns](#advanced-patterns)

---

## What is a State Machine?

A **state machine** (or finite state machine, FSM) is a computational model that:
- Exists in exactly **one state** at any time
- **Transitions** between states based on inputs or events
- Performs **actions** when entering states or during transitions

### Formal Definition

A state machine consists of:
- **States** (S): A finite set of possible conditions
- **Initial State** (s₀): Where the machine starts
- **Transitions** (δ): Rules for moving between states
- **Inputs** (Σ): Events or data that trigger transitions
- **Actions** (optional): Operations performed during transitions

### Simple Example: Door

```
States: {OPEN, CLOSED, LOCKED}
Initial: CLOSED
Transitions:
  - CLOSED + "open" → OPEN
  - OPEN + "close" → CLOSED
  - CLOSED + "lock" → LOCKED
  - LOCKED + "unlock" → CLOSED
```

---

## Why State Machines Matter

### Problem: Complex Nested Logic

**Without state machines:**
```go
func processCharacter(char rune, inQuote bool, inEscape bool, buffer string) {
    if inQuote {
        if inEscape {
            if char == 'n' {
                // ...
            } else if char == 't' {
                // ...
            } else if ...
        } else {
            if char == '\\' {
                // ...
            } else if char == '"' {
                // ...
            } else {
                // ...
            }
        }
    } else {
        if char == '"' {
            // ...
        } else if char == ',' {
            // ...
        } else {
            // ...
        }
    }
}
```

**With state machine:**
```go
func processCharacter(char rune, state string, buffer string) string {
    switch state {
    case "NORMAL":
        return handleNormal(char, buffer)
    case "IN_QUOTE":
        return handleQuote(char, buffer)
    case "IN_ESCAPE":
        return handleEscape(char, buffer)
    }
}
```

### Benefits

1. **Clarity:** One state = one behavior set
2. **Testability:** Test each state independently
3. **Maintainability:** Add states without touching existing code
4. **Debuggability:** "What state am I in?" is easy to answer
5. **Documentation:** State diagram IS the spec

---

## Visual Representations

### State Diagram

Circle = State, Arrow = Transition, Label = Input/Event

```
     ┌─────────┐  open   ┌─────────┐
     │ CLOSED  │────────>│  OPEN   │
     └────┬────┘         └────┬────┘
          │                   │
     lock │                   │ close
          ↓                   ↓
     ┌─────────┐         ┌─────────┐
     │ LOCKED  │<────────│ CLOSED  │
     └─────────┘  unlock └─────────┘
```

### State Transition Table

| Current State | Input | Next State | Action |
|---------------|-------|------------|--------|
| CLOSED | open | OPEN | - |
| OPEN | close | CLOSED | - |
| CLOSED | lock | LOCKED | Engage bolt |
| LOCKED | unlock | CLOSED | Disengage bolt |

### State Tree (Hierarchical)

```
ROOT
├── IDLE
├── RUNNING
│   ├── ACCELERATING
│   ├── CRUISING
│   └── DECELERATING
└── STOPPED
```

---

## Common State Machine Patterns

### Pattern 1: Enum-Style State Variable

**Use when:** Few states (2-5), simple transitions

```go
type State int

const (
    RED State = iota
    YELLOW
    GREEN
)

type TrafficLight struct {
    current State
}

func (t *TrafficLight) Next() {
    switch t.current {
    case RED:
        t.current = GREEN
    case GREEN:
        t.current = YELLOW
    case YELLOW:
        t.current = RED
    }
}
```

**Pros:** Type-safe, fast, simple
**Cons:** Not extensible, hard to add state data

---

### Pattern 2: String-Based State

**Use when:** States are dynamic or numerous

```go
type Parser struct {
    state string
}

func (p *Parser) Parse(char rune) {
    switch p.state {
    case "READING_KEY":
        if char == '=' {
            p.state = "READING_VALUE"
        }
    case "READING_VALUE":
        if char == '&' {
            p.state = "READING_KEY"
        }
    }
}
```

**Pros:** Flexible, easy to debug (print state name)
**Cons:** No type safety, string comparisons

---

### Pattern 3: Boolean Flags

**Use when:** Only 2 states

```go
type CSVParser struct {
    inQuote bool
}

func (p *CSVParser) Parse(char rune) {
    if p.inQuote {
        // Behavior A
    } else {
        // Behavior B
    }
}
```

**Pros:** Minimal overhead
**Cons:** Only works for binary states, doesn't scale

---

### Pattern 4: State Objects (Interface-Based)

**Use when:** Each state has complex behavior

```go
type State interface {
    Handle(input rune) State
}

type NormalState struct{}
func (s NormalState) Handle(input rune) State {
    if input == '"' {
        return QuotedState{}
    }
    return s
}

type QuotedState struct{}
func (s QuotedState) Handle(input rune) State {
    if input == '"' {
        return NormalState{}
    }
    return s
}

type Parser struct {
    state State
}

func (p *Parser) Process(input rune) {
    p.state = p.state.Handle(input)
}
```

**Pros:** Very extensible, each state is independent
**Cons:** More boilerplate, overkill for simple cases

---

### Pattern 5: Table-Driven State Machine

**Use when:** Transitions are data, not code

```go
type Transition struct {
    From  string
    Input string
    To    string
    Action func()
}

var transitions = []Transition{
    {"IDLE", "start", "RUNNING", startEngine},
    {"RUNNING", "stop", "IDLE", stopEngine},
}

func (m *Machine) Process(input string) {
    for _, t := range transitions {
        if t.From == m.state && t.Input == input {
            if t.Action != nil {
                t.Action()
            }
            m.state = t.To
            return
        }
    }
}
```

**Pros:** Transitions are data (can load from config)
**Cons:** Runtime overhead, harder to validate

---

## Implementation Strategies in Go

### Strategy 1: Sequential State Processing

**Best for:** Parsers, decoders, text processing

```go
func Decode(input string) string {
    runes := []rune(input)
    output := ""

    for i := 0; i < len(runes); {
        // STATE 1: Read count
        count := ""
        for i < len(runes) && isDigit(runes[i]) {
            count += string(runes[i])
            i++
        }

        // STATE 2: Read character
        char := runes[i]
        i++

        // STATE 3: Handle escape
        if isDigit(char) && i < len(runes) {
            i++ // Skip duplicate
        }

        // STATE 4: Output
        output += repeat(char, count)
    }

    return output
}
```

**Key:** Each state completes before next begins (sequential, not nested)

---

### Strategy 2: Event-Driven State Machine

**Best for:** UI, game logic, network protocols

```go
type EventType int

const (
    CLICK EventType = iota
    HOVER
    DRAG
)

type UIElement struct {
    state string
}

func (u *UIElement) HandleEvent(event EventType) {
    switch u.state {
    case "IDLE":
        if event == CLICK {
            u.state = "ACTIVE"
        }
    case "ACTIVE":
        if event == CLICK {
            u.state = "IDLE"
        }
    }
}
```

---

### Strategy 3: Character-at-a-Time Processing

**Best for:** Streaming, real-time input

```go
type Tokenizer struct {
    state string
    buffer string
}

func (t *Tokenizer) Feed(char rune) *Token {
    switch t.state {
    case "START":
        if isLetter(char) {
            t.buffer = string(char)
            t.state = "IN_WORD"
        } else if isDigit(char) {
            t.buffer = string(char)
            t.state = "IN_NUMBER"
        }

    case "IN_WORD":
        if isLetter(char) {
            t.buffer += string(char)
        } else {
            token := &Token{Type: WORD, Value: t.buffer}
            t.buffer = ""
            t.state = "START"
            return token
        }
    }
    return nil
}
```

---

## State Machine Design Process

### Step 1: Identify States

Ask: "What distinct modes or phases does my system have?"

**Example: Email Validator**
- Reading username
- Reading domain
- Reading TLD (top-level domain)

### Step 2: Define Transitions

Ask: "What causes the system to move between states?"

**Example:**
- '@' triggers: username → domain
- '.' triggers: domain → TLD

### Step 3: Determine Actions

Ask: "What should happen when entering/exiting states?"

**Example:**
- On '@': save username, reset buffer
- On '.': save domain, reset buffer
- On end: validate TLD

### Step 4: Handle Edge Cases

Ask: "What happens if...?"
- Input ends mid-state?
- Invalid transitions?
- No valid next state?

### Step 5: Draw the Diagram

**Before coding**, sketch on paper:

```
   ┌──────────────┐  @  ┌──────────────┐  .  ┌──────────────┐
   │  USERNAME    │────>│   DOMAIN     │────>│     TLD      │
   └──────────────┘     └──────────────┘     └──────────────┘
        │ start                                      │ end
        ↓                                            ↓
     [EMPTY]                                      [VALID]
```

### Step 6: Implement Incrementally

1. Write skeleton (states, no logic)
2. Add one transition
3. Test that transition
4. Add next transition
5. Repeat

---

## Real-World Examples

### Example 1: TCP Connection

```
States: CLOSED, LISTEN, SYN_SENT, SYN_RECEIVED, ESTABLISHED, FIN_WAIT, CLOSE_WAIT

         CLOSED
           │
      listen│     connect
           ↓           ↓
        LISTEN ───> SYN_SENT
           │           │
    SYN    │           │ SYN+ACK
           ↓           ↓
    SYN_RECEIVED ─> ESTABLISHED
                       │
                  close│
                       ↓
                  FIN_WAIT
                       │
                   ACK │
                       ↓
                   CLOSED
```

### Example 2: Regex Matcher

```go
// Match: a*b (zero or more a's, then b)

States: START, MATCHING_A, SUCCESS, FAIL

func Match(input string) bool {
    state := "START"

    for _, char := range input {
        switch state {
        case "START":
            if char == 'a' {
                state = "MATCHING_A"
            } else if char == 'b' {
                state = "SUCCESS"
            } else {
                state = "FAIL"
            }

        case "MATCHING_A":
            if char == 'a' {
                // Stay in MATCHING_A
            } else if char == 'b' {
                state = "SUCCESS"
            } else {
                state = "FAIL"
            }

        case "SUCCESS", "FAIL":
            state = "FAIL" // Anything after success/fail = fail
        }
    }

    return state == "SUCCESS"
}
```

### Example 3: JSON Parser (Simplified)

```
States: START, IN_STRING, IN_NUMBER, IN_OBJECT, IN_ARRAY

Transitions:
  START + '{' → IN_OBJECT
  START + '[' → IN_ARRAY
  START + '"' → IN_STRING
  START + digit → IN_NUMBER

  IN_STRING + '"' (non-escaped) → END_STRING
  IN_NUMBER + non-digit → END_NUMBER

  IN_OBJECT + '}' → END_OBJECT
  IN_ARRAY + ']' → END_ARRAY
```

---

## Common Pitfalls

### Pitfall 1: Forgetting to Update State

```go
// ❌ WRONG - state never changes
switch state {
case "A":
    // Do work, but forget: state = "B"
}

// ✓ RIGHT
switch state {
case "A":
    // Do work
    state = "B" // Transition!
}
```

### Pitfall 2: Missing End-of-Input Handling

```go
// ❌ WRONG
for i := 0; i < len(input); i++ {
    if input[i] == ',' {
        save(buffer)
        buffer = ""
    } else {
        buffer += string(input[i])
    }
}
// buffer still has last item!

// ✓ RIGHT
for i := 0; i < len(input); i++ {
    // ... same
}
if buffer != "" {
    save(buffer) // Save last item
}
```

### Pitfall 3: Ambiguous State Definitions

```go
// ❌ WRONG - what's the difference?
state := "processing"
state := "handling"

// ✓ RIGHT - clear, distinct names
state := "READING_HEADER"
state := "READING_BODY"
```

### Pitfall 4: No Error State

```go
// ❌ WRONG - no way to handle invalid input
switch state {
case "A":
    if validInput {
        state = "B"
    }
    // What if invalid? State stays "A" forever?
}

// ✓ RIGHT - explicit error state
switch state {
case "A":
    if validInput {
        state = "B"
    } else {
        state = "ERROR"
    }
case "ERROR":
    // Stop processing or log error
    return
}
```

### Pitfall 5: State Explosion

```go
// ❌ WRONG - too many states
// "READING_USERNAME_FIRST_CHAR"
// "READING_USERNAME_MIDDLE_CHAR"
// "READING_USERNAME_LAST_CHAR"

// ✓ RIGHT - combine similar states
// "READING_USERNAME" (track position separately if needed)
```

---

## Advanced Patterns

### Hierarchical State Machines

States can contain sub-states:

```
VEHICLE
├── OFF
└── ON
    ├── PARKED
    ├── DRIVING
    │   ├── ACCELERATING
    │   ├── CRUISING
    │   └── BRAKING
    └── IDLING
```

**Implementation:**
```go
type Vehicle struct {
    power  string // "OFF" or "ON"
    motion string // "PARKED", "DRIVING", "IDLING"
    driving string // "ACCELERATING", "CRUISING", "BRAKING"
}

func (v *Vehicle) Process(event string) {
    if v.power == "OFF" {
        if event == "start" {
            v.power = "ON"
            v.motion = "IDLING"
        }
        return
    }

    // Only process motion if power is ON
    switch v.motion {
    case "PARKED":
        if event == "drive" {
            v.motion = "DRIVING"
            v.driving = "ACCELERATING"
        }
    case "DRIVING":
        // Handle driving sub-states
    }
}
```

### State Stack (Pushdown Automaton)

Used for nested structures (parentheses, HTML tags, function calls):

```go
type StateStack struct {
    stack []string
}

func (s *StateStack) Push(state string) {
    s.stack = append(s.stack, state)
}

func (s *StateStack) Pop() string {
    if len(s.stack) == 0 {
        return ""
    }
    state := s.stack[len(s.stack)-1]
    s.stack = s.stack[:len(s.stack)-1]
    return state
}

func (s *StateStack) Current() string {
    if len(s.stack) == 0 {
        return "EMPTY"
    }
    return s.stack[len(s.stack)-1]
}

// Example: Matching nested parentheses
func MatchParens(input string) bool {
    stack := &StateStack{}

    for _, char := range input {
        if char == '(' {
            stack.Push("OPEN")
        } else if char == ')' {
            if stack.Pop() == "" {
                return false // No matching open
            }
        }
    }

    return len(stack.stack) == 0 // All matched
}
```

### Timed State Machines

States with timeouts:

```go
type TimedState struct {
    name      string
    enteredAt time.Time
    timeout   time.Duration
}

func (s *TimedState) IsExpired() bool {
    return time.Since(s.enteredAt) > s.timeout
}

type TrafficLight struct {
    state TimedState
}

func (t *TrafficLight) Update() {
    if t.state.IsExpired() {
        switch t.state.name {
        case "RED":
            t.state = TimedState{"GREEN", time.Now(), 30 * time.Second}
        case "GREEN":
            t.state = TimedState{"YELLOW", time.Now(), 5 * time.Second}
        case "YELLOW":
            t.state = TimedState{"RED", time.Now(), 30 * time.Second}
        }
    }
}
```

---

## Key Takeaways

1. **State machines make complex behavior simple** by breaking it into discrete modes
2. **Draw before coding** - visual diagrams prevent logic errors
3. **One state = one behavior set** - no nested ifs trying to handle multiple states
4. **Sequential > Nested** - complete each state fully before moving to next
5. **Test each state independently** - easier debugging
6. **Handle edge cases explicitly** - EOF, error states, invalid transitions
7. **Choose the right pattern** - enum for simple, string for flexible, interface for extensible

---

## Further Reading

- **"Crafting Interpreters"** by Robert Nystrom - Excellent state machine examples in lexers
- **"The Mythical Man-Month"** by Fred Brooks - State diagrams in system design
- **UML State Diagrams** - Standard notation for state machines
- **Automata Theory** - Theoretical foundation (DFA, NFA, pushdown automata)

---

## Next Steps

1. Read `PARSING_TECHNIQUES.md` for parsing-specific patterns
2. Read `DEBUGGING_ALGORITHMS.md` for troubleshooting strategies
3. Practice with the 3 quick exercises
4. Apply to run-length encoding Decode refactor
5. Complete the full Algorithms & State Machines module

**Remember:** State machines are everywhere. Once you recognize the pattern, you'll see it in parsers, protocols, UI, games, workflows, and more!

---

**Version:** 1.0
**Last Updated:** 2025-11-14
**Related:** PARSING_TECHNIQUES.md, DEBUGGING_ALGORITHMS.md, ENCODING_DECODING_PATTERNS.md
