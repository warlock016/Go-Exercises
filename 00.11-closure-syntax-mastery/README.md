# Module 00.11: Closure Syntax Mastery

**Focus:** State management syntax in closures
**Exercises:** 10
**Time:** 2-3 hours
**Prerequisites:** Module 00.10 (Closure Practice)

---

## Why This Module?

You understand *what* closures do, but the *syntax* trips you up:
- "Where do I declare the state variable?"
- "How many `func` keywords do I need?"
- "How do I return multiple closures?"

This module drills the **syntax patterns** until they become automatic.

---

## Syntax Cheat Sheet

Study these 4 patterns before starting. Refer back when stuck.

### Pattern 1: Multiple Return Closures (Shared State)

```go
func Factory(initial int) (increment func(), getValue func() int) {
    //                     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    //                     Named return values - both are closures

    state := initial          // ← State declared BEFORE closures

    increment = func() {      // ← ASSIGN to named return, don't use :=
        state++
    }

    getValue = func() int {
        return state          // ← Both closures capture same 'state'
    }

    return                    // ← Named returns, can omit values
}
```

**Key points:**
- State goes BEFORE any closure definitions
- Use `=` to assign to named returns, not `:=`
- Both closures share the same captured variable

---

### Pattern 2: Accept Function, Return Function

```go
func Wrapper(fn func(int) int) func(int) int {
    //        ^^^^^^^^^^^^^^^ Accept a function
    //                         ^^^^^^^^^^^^^^^ Return a function

    callCount := 0            // ← Your state

    return func(x int) int {  // ← The wrapped version
        callCount++
        return fn(x)          // ← Call the captured function
    }
}
```

**Key points:**
- `fn` is captured just like any other variable
- Input and output function signatures can differ (see Pattern 6)

---

### Pattern 3: Three-Level Nesting (Middleware)

```go
func Middleware(config string) func(func(int) int) func(int) int {
//   ^^^^^^^^^^               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//   Level 1: Config          Level 2: Takes func, returns func

    return func(fn func(int) int) func(int) int {
    //          ^^^^^^^^^^^^^^^^ Level 2 parameter
    //                           ^^^^^^^^^^^^^^^ Level 2 returns Level 3

        return func(x int) int {
        //     ^^^^^^^^^^^^^^^^ Level 3: The actual wrapper
            fmt.Println(config)  // Uses Level 1 capture
            return fn(x)         // Uses Level 2 capture
        }
    }
}

// Usage:
// withLogging := Middleware("LOG")     // Returns Level 2
// wrapped := withLogging(myFunc)       // Returns Level 3
// result := wrapped(42)                // Executes
```

**Key points:**
- Read inside-out: Level 3 is what actually runs
- Each level can capture from outer levels
- Count the `func` keywords: 3 levels = 3 `func`s in signature

---

### Pattern 4: Self-Referential Type (Builder)

```go
type Builder func(string) Builder    // ← Type returns itself!

func NewBuilder() Builder {
    state := []string{}

    var builder Builder              // ← Declare BEFORE assigning
    builder = func(s string) Builder {
        state = append(state, s)
        return builder               // ← Return self for chaining
    }

    return builder
}

// Usage:
// b := NewBuilder()
// b("a")("b")("c")                  // Chained calls!
```

**Key points:**
- Can't use `:=` for self-reference (variable doesn't exist yet)
- Declare with `var`, then assign with `=`

---

## Exercises

| # | Name | Pattern | Difficulty |
|---|------|---------|------------|
| 1 | StatefulOperation | Multiple return closures | ⭐⭐ |
| 2 | UndoRedo | Triple return + tuple returns | ⭐⭐⭐ |
| 3 | Logger | Accept func + config | ⭐⭐ |
| 4 | Compose | Accept two funcs | ⭐⭐ |
| 5 | TimingMiddleware | Three-level nesting | ⭐⭐⭐⭐ |
| 6 | PartialApply | Arity transformation | ⭐⭐⭐ |
| 7 | Pipeline | Slice of functions | ⭐⭐⭐ |
| 8 | Builder | Self-referential type | ⭐⭐⭐⭐ |
| 9 | EventBus | Channel coordination | ⭐⭐⭐⭐ |
| 10 | MiddlewareChain | Capstone: all patterns | ⭐⭐⭐⭐⭐ |

---

## Exercise 1: StatefulOperation

Return two closures that share state: one adds to a value, one retrieves it.

```go
func StatefulOperation(initial int) (add func(int), getValue func() int)
```

**Example:**
```go
add, get := StatefulOperation(10)
get()      // 10
add(5)
get()      // 15
add(-3)
get()      // 12
```

---

## Exercise 2: UndoRedo

Return three closures for an undo/redo system. `do` applies a value, `undo` reverts, `redo` re-applies.

```go
func UndoRedo() (do func(int), undo func() (int, bool), redo func() (int, bool))
```

**Example:**
```go
do, undo, redo := UndoRedo()
do(10)
do(20)
do(30)
val, ok := undo()   // 20, true (reverted from 30)
val, ok = undo()    // 10, true
val, ok = redo()    // 20, true
val, ok = undo()    // 10, true
val, ok = undo()    // 0, false (nothing to undo)
```

---

## Exercise 3: Logger

Wrap a function with logging that prints before and after each call.

```go
func Logger(prefix string, fn func(int) int) func(int) int
```

**Example:**
```go
double := func(x int) int { return x * 2 }
logged := Logger("CALC", double)
logged(5)
// Prints: [CALC] input: 5
// Prints: [CALC] output: 10
// Returns: 10
```

---

## Exercise 4: Compose

Return a function that applies g first, then f: `Compose(f, g)(x)` = `f(g(x))`

```go
func Compose(f, g func(int) int) func(int) int
```

**Example:**
```go
addOne := func(x int) int { return x + 1 }
double := func(x int) int { return x * 2 }

addThenDouble := Compose(double, addOne)  // double(addOne(x))
addThenDouble(5)  // 12: (5+1)*2

doubleThenAdd := Compose(addOne, double)  // addOne(double(x))
doubleThenAdd(5)  // 11: (5*2)+1
```

---

## Exercise 5: TimingMiddleware

Create middleware that logs execution time for slow functions.

```go
func TimingMiddleware(threshold time.Duration) func(func(int) int) func(int) int
```

**Example:**
```go
slowWarning := TimingMiddleware(100 * time.Millisecond)

slowFunc := func(x int) int {
    time.Sleep(150 * time.Millisecond)
    return x * 2
}

wrapped := slowWarning(slowFunc)
wrapped(5)
// Prints: "SLOW: took 150ms (threshold: 100ms)"
// Returns: 10
```

---

## Exercise 6: PartialApply

Pre-fill the first argument of a two-argument function.

```go
func PartialApply(fn func(int, int) int, first int) func(int) int
```

**Example:**
```go
add := func(a, b int) int { return a + b }
add10 := PartialApply(add, 10)
add10(5)   // 15
add10(20)  // 30

multiply := func(a, b int) int { return a * b }
triple := PartialApply(multiply, 3)
triple(7)  // 21
```

---

## Exercise 7: Pipeline

Build a processing pipeline by adding transformation functions.

```go
func Pipeline() (add func(func(int) int), execute func(int) int)
```

**Example:**
```go
add, exec := Pipeline()
add(func(x int) int { return x + 1 })
add(func(x int) int { return x * 2 })
add(func(x int) int { return x - 3 })

exec(5)  // ((5+1)*2)-3 = 9
exec(10) // ((10+1)*2)-3 = 19
```

---

## Exercise 8: Builder

Create a fluent builder using self-referential function type.

```go
type OptionBuilder func(string) OptionBuilder

func NewOptionBuilder() (builder OptionBuilder, getOptions func() []string)
```

**Example:**
```go
builder, getOpts := NewOptionBuilder()
builder("verbose")("debug")("color")

opts := getOpts()  // ["verbose", "debug", "color"]
```

---

## Exercise 9: EventBus

Create a simple pub/sub system with channels.

```go
func EventBus(bufferSize int) (publish func(string), subscribe func() <-chan string, close func())
```

**Example:**
```go
pub, sub, close := EventBus(10)

ch := sub()  // Get receive-only channel

pub("hello")
pub("world")

msg1 := <-ch  // "hello"
msg2 := <-ch  // "world"

close()  // Closes channel, subsequent reads return zero value
```

---

## Exercise 10: MiddlewareChain (Capstone)

Chain multiple middleware functions together.

```go
type Handler func(string) string
type Middleware func(Handler) Handler

func Chain(middlewares ...Middleware) Middleware
```

**Example:**
```go
upper := func(next Handler) Handler {
    return func(s string) string {
        return next(strings.ToUpper(s))
    }
}

exclaim := func(next Handler) Handler {
    return func(s string) string {
        return next(s) + "!"
    }
}

base := func(s string) string { return s }

chained := Chain(upper, exclaim)
handler := chained(base)
handler("hello")  // "HELLO!"
```

---

## Think About

1. In Pattern 1, why can't you use `:=` for named return values?
2. In Pattern 3, which variables are visible at each nesting level?
3. In Pattern 4, why must you declare `var builder Builder` before assigning?
4. When would you choose multiple return closures vs a struct with methods?
5. How does the Pipeline pattern relate to function composition?

---

## What This Teaches

After completing this module, you'll instinctively know:
- **Where state goes**: Before any closure definitions
- **How to count `func` keywords**: One per nesting level
- **When to use `var` vs `:=`**: `var` for self-reference or named returns
- **How closures share state**: Same variable, different closures
