# Exercise 01: Goroutine Basics

**Learning Goal:** Understand goroutine lifecycle, the `go` keyword, and why synchronization is necessary

**Difficulty:** Tier 1 - Introduction
**Estimated Time:** 25-30 minutes

---

## Problem Description

Goroutines are lightweight threads managed by the Go runtime. They're incredibly cheap to create (a few KB of stack), enabling thousands to run concurrently. But this power comes with a challenge: the main goroutine doesn't wait for others to complete.

In this exercise, you'll experience the fundamental problem of unsynchronized goroutines and learn to use `sync.WaitGroup` to coordinate their completion.

---

## Function Signatures

```go
// PrintMessages prints each message in a separate goroutine
// WARNING: This function demonstrates the problem - messages may not print!
func PrintMessages(messages []string)

// PrintMessagesSync prints each message in a goroutine, waiting for all to complete
func PrintMessagesSync(messages []string)

// CounterUnsafe spawns n goroutines that each increment a counter
// Returns the final counter value (demonstrates race condition)
func CounterUnsafe(n int) int

// CounterSafe spawns n goroutines that safely increment a counter
// Returns the final counter value (should equal n)
func CounterSafe(n int) int
```

---

## Examples

### PrintMessages (Unsafe)
```
Input:  PrintMessages([]string{"hello", "world", "go"})
Output: (unpredictable - may print nothing, some, or all messages)

// This happens because main() doesn't wait for goroutines
```

### PrintMessagesSync
```
Input:  PrintMessagesSync([]string{"hello", "world", "go"})
Output: hello
        world
        go
        (order may vary, but all messages print)
```

### CounterUnsafe
```
Input:  CounterUnsafe(1000)
Output: (some number less than 1000 due to race condition)

// Multiple goroutines reading/writing counter simultaneously
// causes lost updates
```

### CounterSafe
```
Input:  CounterSafe(1000)
Output: 1000

// Proper synchronization ensures all increments are counted
```

---

## Instructions

1. Implement `PrintMessages` using `go func()` - observe that messages don't reliably print
2. Implement `PrintMessagesSync` using `sync.WaitGroup` to wait for all goroutines
3. Implement `CounterUnsafe` with a shared counter variable - observe the race condition
4. Implement `CounterSafe` using `sync.Mutex` to protect the counter
5. Run tests with `go test -v`
6. Run tests with `go test -race` to see the race detector catch `CounterUnsafe`

---

## Hints

### Basic
- The `go` keyword spawns a new goroutine: `go func() { ... }()`
- `sync.WaitGroup` has three methods: `Add(n)`, `Done()`, and `Wait()`
- Call `Add` before spawning, `Done` when goroutine finishes, `Wait` to block

### Intermediate
- For `PrintMessages`, the function returns before goroutines run - that's the point!
- For `PrintMessagesSync`, each goroutine should call `wg.Done()` (use `defer`)
- `sync.Mutex` protects critical sections with `Lock()` and `Unlock()`
- Always use `defer mu.Unlock()` right after `Lock()` to prevent deadlocks

### Solution Pattern
```go
var wg sync.WaitGroup
for _, item := range items {
    wg.Add(1)
    go func(val Type) {
        defer wg.Done()
        // work with val
    }(item)  // Pass item to avoid closure capture bug
}
wg.Wait()
```

---

## Think About

1. Why do we pass `item` as a parameter to the goroutine instead of capturing it directly?
2. What happens if you forget to call `wg.Done()`?
3. Why does `CounterUnsafe` return different values on different runs?
4. When would you use a mutex vs a channel for synchronization?

---

## What This Teaches

- **Goroutine spawning** - The `go` keyword creates concurrent execution
- **Main doesn't wait** - The program exits when main() returns, not when goroutines finish
- **sync.WaitGroup** - Coordinates completion of multiple goroutines
- **Race conditions** - Concurrent access to shared memory without protection causes bugs
- **sync.Mutex** - Provides mutual exclusion for safe shared state access
- **Race detector** - `go test -race` finds data races automatically
- **Closure variable capture** - Why loop variables need to be passed as parameters
