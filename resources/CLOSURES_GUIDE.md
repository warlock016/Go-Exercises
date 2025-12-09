# Closures in Go: A Complete Guide

A comprehensive reference for understanding and using closures effectively.

---

## What Is a Closure?

A **closure** is a function that **captures** variables from its surrounding scope. The function "closes over" these variables, maintaining access to them even after the outer function returns.

```go
func outer() func() int {
    count := 0           // Variable in outer scope

    return func() int {  // Inner function (the closure)
        count++          // Captures 'count' from outer scope
        return count
    }
}

counter := outer()  // 'count' lives on, captured by the closure
counter()  // 1
counter()  // 2
counter()  // 3
```

**Key insight**: `count` is not destroyed when `outer()` returns because the returned function still references it. Go moves `count` to the heap automatically.

---

## Anatomy of a Closure

```go
func makeMultiplier(factor int) func(int) int {
    //  └─────────────┬─────────────┘
    //         Captured variable

    return func(x int) int {
        //      └──┬──┘
        //    Parameter (not captured)

        return x * factor
        //         └──┬──┘
        //      Uses captured variable
    }
}
```

| Component | Description |
|-----------|-------------|
| Outer function | Creates the environment, initializes captured variables |
| Captured variables | Variables from outer scope that the closure references |
| Inner function | The closure itself, returned or passed around |
| Parameters | Passed at call time, not captured |

---

## Variable Capture: Value vs Reference

### Closures Capture by Reference

```go
func demo() {
    x := 10

    f := func() {
        fmt.Println(x)  // Captures reference to x
    }

    x = 20  // Modify x after closure creation
    f()     // Prints 20, not 10!
}
```

### The Loop Variable Trap

```go
// WRONG - All closures share the same 'i'
funcs := []func(){}
for i := 0; i < 3; i++ {
    funcs = append(funcs, func() {
        fmt.Println(i)  // Captures reference to loop variable
    })
}
for _, f := range funcs {
    f()  // Prints: 3, 3, 3 (not 0, 1, 2!)
}
```

```go
// CORRECT - Capture current value
funcs := []func(){}
for i := 0; i < 3; i++ {
    i := i  // Shadow with new variable (Go idiom)
    funcs = append(funcs, func() {
        fmt.Println(i)  // Each closure has its own 'i'
    })
}
for _, f := range funcs {
    f()  // Prints: 0, 1, 2 ✓
}
```

**Note**: Go 1.22+ fixes this for `for` loops, but the pattern is still important to understand.

---

## Common Closure Patterns

### 1. Counter / Accumulator

Maintain running state across calls.

```go
func Counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := Counter()
c()  // 1
c()  // 2
```

### 2. Factory Functions

Create pre-configured functions.

```go
func Multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

double := Multiplier(2)
triple := Multiplier(3)
double(5)  // 10
triple(5)  // 15
```

### 3. Memoization / Caching

Cache expensive computation results.

```go
func Memoize(fn func(int) int) func(int) int {
    cache := make(map[int]int)

    return func(x int) int {
        if result, found := cache[x]; found {
            return result  // Return cached value
        }
        result := fn(x)    // Compute
        cache[x] = result  // Cache for next time
        return result
    }
}

slowFib := func(n int) int { /* expensive */ }
fastFib := Memoize(slowFib)
```

### 4. Rate Limiting

Track events over time windows.

```go
func RateLimiter(maxCalls int, window time.Duration) func() bool {
    calls := []time.Time{}

    return func() bool {
        now := time.Now()

        // Remove expired timestamps
        valid := []time.Time{}
        for _, t := range calls {
            if now.Sub(t) <= window {
                valid = append(valid, t)
            }
        }

        if len(valid) >= maxCalls {
            calls = valid
            return false  // Rate limited
        }

        calls = append(valid, now)
        return true  // Allowed
    }
}
```

### 5. Middleware / Decorators

Wrap functions with additional behavior.

```go
func WithLogging(fn func(int) int) func(int) int {
    return func(x int) int {
        log.Printf("Calling with %d", x)
        result := fn(x)
        log.Printf("Returned %d", result)
        return result
    }
}

add := func(x int) int { return x + 1 }
loggedAdd := WithLogging(add)
loggedAdd(5)  // Logs: Calling with 5, Returned 6
```

### 6. Deferred Configuration

Delay configuration until needed.

```go
func DatabaseConnector(dsn string) func() (*sql.DB, error) {
    var db *sql.DB
    var err error
    var once sync.Once

    return func() (*sql.DB, error) {
        once.Do(func() {
            db, err = sql.Open("postgres", dsn)
        })
        return db, err
    }
}

getDB := DatabaseConnector("postgres://...")
// Connection not established yet

db, err := getDB()  // Now it connects (once)
db, err = getDB()   // Returns same connection
```

### 7. Iterator / Generator

Produce values on demand.

```go
func Fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        result := a
        a, b = b, a+b
        return result
    }
}

fib := Fibonacci()
fib()  // 0
fib()  // 1
fib()  // 1
fib()  // 2
fib()  // 3
```

### 8. Partial Application

Pre-fill some function arguments.

```go
func Add(a, b, c int) int {
    return a + b + c
}

func PartialAdd(a int) func(int, int) int {
    return func(b, c int) int {
        return Add(a, b, c)
    }
}

add10 := PartialAdd(10)
add10(5, 3)  // 18 (10 + 5 + 3)
```

---

## Closures in HTTP Handlers

A very common real-world use case:

```go
// Without closure - need global variables or complex setup
var db *sql.DB  // Global - bad!

func UserHandler(w http.ResponseWriter, r *http.Request) {
    user := db.Query(...)  // Uses global
}

// With closure - dependencies are explicit and encapsulated
func MakeUserHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        user := db.Query(...)  // Uses captured db
    }
}

// Usage
db := connectToDatabase()
http.HandleFunc("/user", MakeUserHandler(db))
```

### Middleware Chain with Closures

```go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s took %v", r.Method, r.URL.Path, time.Since(start))
    })
}

func AuthMiddleware(secret string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            token := r.Header.Get("Authorization")
            if !validateToken(token, secret) {
                http.Error(w, "Unauthorized", 401)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
```

---

## Closures vs Structs with Methods

Both can encapsulate state. When to use which?

### Use Closures When:
- Single function with captured state
- Simple state (1-3 variables)
- Short-lived, throwaway functions
- Functional programming style (map, filter, callbacks)

```go
multiplier := func(factor int) func(int) int {
    return func(x int) int { return x * factor }
}
```

### Use Structs When:
- Multiple methods sharing state
- Complex state (many fields)
- Need to serialize/deserialize state
- State needs to be inspected or modified externally

```go
type Multiplier struct {
    Factor int
}

func (m Multiplier) Multiply(x int) int {
    return x * m.Factor
}

func (m *Multiplier) SetFactor(f int) {
    m.Factor = f
}
```

### Hybrid Approach

```go
type Counter struct {
    increment func() int
    decrement func() int
    value     func() int
}

func NewCounter() Counter {
    count := 0
    return Counter{
        increment: func() int { count++; return count },
        decrement: func() int { count--; return count },
        value:     func() int { return count },
    }
}
```

---

## Memory and Lifetime

### Escape Analysis

Go's compiler determines when captured variables must move to the heap:

```go
func example() func() int {
    x := 42  // Compiler sees x is captured by returned function
             // x "escapes" to heap (not stack)
    return func() int {
        return x
    }
}
```

You can see escape analysis with:
```bash
go build -gcflags="-m" main.go
# Output: "moved to heap: x"
```

### Garbage Collection

Closures keep captured variables alive:

```go
func leaky() func() {
    hugeData := make([]byte, 1<<30)  // 1GB

    return func() {
        _ = hugeData[0]  // Closure captures hugeData
    }
}

leak := leaky()  // 1GB stays in memory as long as 'leak' exists
leak = nil       // Now hugeData can be garbage collected
```

**Best practice**: Only capture what you need:

```go
func better() func() byte {
    hugeData := make([]byte, 1<<30)
    firstByte := hugeData[0]  // Extract only what's needed

    return func() byte {
        return firstByte  // Only captures single byte
    }
}
```

---

## Error Handling in Closure Factories

### Return Error from Factory

```go
func MakeProcessor(config Config) (func([]byte) error, error) {
    if config.MaxSize <= 0 {
        return nil, errors.New("MaxSize must be positive")
    }

    return func(data []byte) error {
        if len(data) > config.MaxSize {
            return errors.New("data too large")
        }
        // process...
        return nil
    }, nil
}

proc, err := MakeProcessor(config)
if err != nil {
    log.Fatal(err)
}
```

### Panic for Programmer Errors

```go
func MustMakeProcessor(config Config) func([]byte) error {
    if config.MaxSize <= 0 {
        panic("MustMakeProcessor: MaxSize must be positive")
    }
    // ...
}
```

### Reinitializing Closures

You don't reinitialize - you replace:

```go
var processor func([]byte) error

// Initial setup
processor, _ = MakeProcessor(config1)

// Later, need different config
processor, _ = MakeProcessor(config2)  // Old closure is GC'd
```

---

## Testing Closures

### Test the Factory

```go
func TestCounter(t *testing.T) {
    counter := Counter()

    if got := counter(); got != 1 {
        t.Errorf("first call = %d, want 1", got)
    }
    if got := counter(); got != 2 {
        t.Errorf("second call = %d, want 2", got)
    }
}
```

### Test Independence

```go
func TestCounterIndependence(t *testing.T) {
    c1 := Counter()
    c2 := Counter()

    c1()
    c1()
    c1()  // c1 is at 3

    if got := c2(); got != 1 {
        t.Errorf("c2 first call = %d, want 1 (should be independent)", got)
    }
}
```

### Test Edge Cases

```go
func TestRateLimiter(t *testing.T) {
    limiter := RateLimiter(2, 100*time.Millisecond)

    // Should allow first 2 calls
    if !limiter() { t.Error("call 1 should be allowed") }
    if !limiter() { t.Error("call 2 should be allowed") }

    // Should block 3rd call
    if limiter() { t.Error("call 3 should be blocked") }

    // After window expires, should allow again
    time.Sleep(150 * time.Millisecond)
    if !limiter() { t.Error("call after window should be allowed") }
}
```

---

## Common Mistakes

### 1. Capturing Loop Variables (Pre-Go 1.22)
```go
// Wrong
for i := 0; i < 3; i++ {
    go func() { fmt.Println(i) }()  // All print 3
}

// Right
for i := 0; i < 3; i++ {
    i := i
    go func() { fmt.Println(i) }()  // Prints 0, 1, 2
}
```

### 2. Nil Function Values
```go
var fn func() int  // nil

fn()  // PANIC: nil pointer dereference

// Always check or initialize
if fn != nil {
    fn()
}
```

### 3. Unintended Sharing
```go
func wrong() (func(), func()) {
    count := 0
    inc := func() { count++ }
    dec := func() { count-- }
    return inc, dec  // Both share same count!
}

// If you want independent state, create separate closures
```

### 4. Capturing Large Objects
```go
// Bad - captures entire large slice
func bad(data []int) func() int {
    return func() int {
        return data[0]
    }
}

// Good - capture only what's needed
func good(data []int) func() int {
    first := data[0]
    return func() int {
        return first
    }
}
```

---

## Summary

| Concept | Key Point |
|---------|-----------|
| Definition | Function that captures variables from enclosing scope |
| Capture | By reference, not value |
| Lifetime | Captured variables live as long as the closure |
| Loop trap | Shadow loop variables: `i := i` |
| Factories | Validate config upfront, return error or panic |
| Reinitialize | Create new closure, assign to same variable |
| vs Structs | Closures for single function, structs for multiple methods |
| HTTP | Common for handlers, middleware with dependencies |

---

## Syntax Quick Reference

When syntax trips you up, consult these patterns:

### Where Does State Go?

```go
func Factory() (a func(), b func()) {
    state := 0                // ← HERE: Before any closure definitions

    a = func() { state++ }    // ← Use = not := for named returns
    b = func() { return state }

    return
}
```

### How Many `func` Keywords?

Count the nesting levels:

```go
// Level 1: Simple factory
func Counter() func() int { ... }
//          ^^^^^^^^^^^^^ 1 func in return type

// Level 2: Middleware (accepts func, returns func)
func Wrapper(fn func(int) int) func(int) int { ... }
//              ^^^^^^^^^^^^^^ ^^^^^^^^^^^^^^ 2 funcs in signature

// Level 3: Configurable middleware
func Middleware(cfg T) func(func(int) int) func(int) int { ... }
//                     ^^^^ ^^^^^^^^^^^^^^ ^^^^^^^^^^^^^^ 3 funcs
```

### Multiple Return Closures

```go
func TwoClosures() (first func(), second func() int) {
    shared := 0

    first = func() {           // Assign with =
        shared++
    }

    second = func() int {      // Both capture 'shared'
        return shared
    }

    return                     // Named returns, can omit values
}
```

### Self-Referential Types

```go
type Builder func(string) Builder

func NewBuilder() Builder {
    var b Builder              // ← Declare before assign
    b = func(s string) Builder {
        return b               // ← Return self for chaining
    }
    return b
}
```

### Accept Function + Return Function

```go
func Transform(fn func(A) B) func(A) B {
    return func(x A) B {
        // Pre-processing
        result := fn(x)        // ← Call captured function
        // Post-processing
        return result
    }
}
```

---

## Resources

- [Go Blog - First Class Functions](https://go.dev/blog/functions-codewalk)
- [Go by Example - Closures](https://gobyexample.com/closures)
- [Effective Go - Functions](https://go.dev/doc/effective_go#functions)
