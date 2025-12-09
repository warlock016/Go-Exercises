# Exercise 05a: Closure Practice

**Difficulty:** Intermediate
**Time:** 45-60 minutes
**Prerequisites:** Exercise 05 (Closures basics)

---

## Learning Goal

Deepen your understanding of closures by implementing increasingly complex patterns that demonstrate state encapsulation, configuration capture, and real-world applications.

---

## Overview

This exercise provides additional closure practice beyond the basics. You'll implement:

1. **Memoizer** - Cache expensive function results
2. **Toggle** - Alternate between two values
3. **Once** - Execute a function exactly once
4. **Debouncer** - Delay execution until calls stop
5. **Sequence** - Generate custom sequences
6. **Retry** - Retry failed operations with backoff

---

## Exercises

### Exercise 1: Memoizer

Create a function that caches results of expensive computations.

```go
// Memoize wraps a function to cache its results
// Subsequent calls with the same input return cached result
func Memoize(fn func(int) int) func(int) int
```

**Example:**
```go
callCount := 0
expensive := func(n int) int {
    callCount++
    time.Sleep(100 * time.Millisecond)  // Simulate expensive work
    return n * n
}

cached := Memoize(expensive)
cached(5)  // Computes, returns 25, callCount = 1
cached(5)  // Returns 25 from cache, callCount still 1
cached(3)  // Computes, returns 9, callCount = 2
cached(5)  // Returns 25 from cache, callCount still 2
```

---

### Exercise 2: Toggle

Create a function that alternates between two values on each call.

```go
// Toggle returns a function that alternates between a and b
func Toggle[T any](a, b T) func() T
```

**Example:**
```go
onOff := Toggle("on", "off")
onOff()  // "on"
onOff()  // "off"
onOff()  // "on"
onOff()  // "off"

yesNo := Toggle(true, false)
yesNo()  // true
yesNo()  // false
```

---

### Exercise 3: Once

Create a function that ensures another function is called exactly once.

```go
// Once returns a function that calls fn only on the first invocation
// Subsequent calls return the first result without calling fn again
func Once[T any](fn func() T) func() T
```

**Example:**
```go
callCount := 0
init := func() string {
    callCount++
    return "initialized"
}

doOnce := Once(init)
doOnce()  // "initialized", callCount = 1
doOnce()  // "initialized", callCount still 1
doOnce()  // "initialized", callCount still 1
```

---

### Exercise 4: Debouncer

Create a function that delays execution until a pause in calls.

```go
// Debounce returns a function that delays calling fn until
// 'delay' duration passes without another call
// Returns a cancel function to stop pending execution
func Debounce(fn func(), delay time.Duration) (debounced func(), cancel func())
```

**Example:**
```go
executed := false
action := func() { executed = true }

debounced, cancel := Debounce(action, 100*time.Millisecond)

debounced()  // Start timer
debounced()  // Reset timer
debounced()  // Reset timer
// Wait 100ms...
// action() is called, executed = true
```

**Use case:** Search-as-you-type (wait for user to stop typing before querying).

---

### Exercise 5: Sequence

Create a generator that produces values according to a custom rule.

```go
// Sequence creates a generator starting at 'start'
// Each call applies 'next' function to get the following value
func Sequence(start int, next func(int) int) func() int
```

**Example:**
```go
// Counting: 1, 2, 3, 4...
counter := Sequence(1, func(n int) int { return n + 1 })
counter()  // 1
counter()  // 2
counter()  // 3

// Powers of 2: 1, 2, 4, 8, 16...
powers := Sequence(1, func(n int) int { return n * 2 })
powers()  // 1
powers()  // 2
powers()  // 4
powers()  // 8

// Fibonacci: 1, 1, 2, 3, 5... (hint: need to track two values)
// This requires a different signature - see bonus challenge
```

---

### Exercise 6: Retry (Challenge)

Create a function that retries failed operations with exponential backoff.

```go
// Retry wraps a function to retry on error up to maxAttempts times
// Uses exponential backoff: wait = initialDelay * 2^attempt
// Returns a function with the same signature that handles retries internally
func Retry(fn func() error, maxAttempts int, initialDelay time.Duration) func() error
```

**Example:**
```go
attempts := 0
flaky := func() error {
    attempts++
    if attempts < 3 {
        return errors.New("temporary failure")
    }
    return nil  // Success on 3rd attempt
}

reliable := Retry(flaky, 5, 10*time.Millisecond)
err := reliable()  // Retries automatically, returns nil after 3 attempts
// Delays: 10ms, 20ms between attempts
```

---

## Bonus Challenges

### Fibonacci Generator

The basic `Sequence` only tracks one value. Implement a Fibonacci generator that tracks two:

```go
func Fibonacci() func() int
// Returns: 0, 1, 1, 2, 3, 5, 8, 13...
```

### Memoize with Expiration

Extend `Memoize` to expire cached values after a duration:

```go
func MemoizeWithTTL(fn func(int) int, ttl time.Duration) func(int) int
```

### Thread-Safe Counter

Make the basic `Counter` safe for concurrent use:

```go
func SafeCounter() func() int
// Hint: Use sync.Mutex or sync/atomic
```

---

## Think About

1. Why does `Memoize` need a map while `Counter` only needs an int?
2. What happens if you call `Toggle` with the same value for both a and b?
3. Why is `Once` useful for initialization code?
4. How does `Debounce` prevent the "thundering herd" problem?
5. Could `Sequence` be implemented with generics for any type?
6. Why does `Retry` use exponential backoff instead of fixed delays?

---

## What This Teaches

- **State encapsulation**: Complex state hidden behind simple function interface
- **Configuration capture**: Settings "baked in" at creation time
- **Lazy evaluation**: Defer computation until needed (Once, Memoize)
- **Time-based logic**: Closures can capture time and manage delays
- **Generic closures**: Using type parameters with closures
- **Real-world patterns**: These patterns appear in production code constantly

---

## Hints

<details>
<summary>Memoize - Basic Approach</summary>

Use a `map[int]int` to store computed results. Check if key exists before computing.

</details>

<details>
<summary>Toggle - State Tracking</summary>

Use a boolean to track which value to return, flip it each call.

</details>

<details>
<summary>Once - Preventing Re-execution</summary>

Use a boolean `called` and store the result in a captured variable.

</details>

<details>
<summary>Debounce - Timer Management</summary>

Use `time.AfterFunc` to schedule execution. Cancel previous timer when new call comes in.

</details>

<details>
<summary>Sequence - Maintaining State</summary>

Store current value, return it, then update for next call.

</details>

<details>
<summary>Retry - Backoff Calculation</summary>

Use a loop with `time.Sleep(delay)` and `delay *= 2` after each failed attempt.

</details>
