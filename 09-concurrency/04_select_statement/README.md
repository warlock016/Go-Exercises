# Exercise 04: Select Statement

**Learning Goal:** Use select to wait on multiple channels and implement timeouts

**Difficulty:** Tier 2 - Application
**Estimated Time:** 35-40 minutes

---

## Problem Description

The `select` statement is the control structure for channel operations. It waits on multiple channel operations simultaneously, executing whichever one is ready first. If multiple are ready, one is chosen randomly.

With `select`, you can implement timeouts, cancellation, non-blocking operations, and channel multiplexing.

---

## Function Signatures

```go
// FirstResponse returns the first value received from either channel
func FirstResponse(ch1, ch2 <-chan string) string

// WithTimeout receives from channel or returns error if timeout exceeded
func WithTimeout(ch <-chan string, timeout time.Duration) (string, error)

// TryReceive attempts non-blocking receive
// Returns (value, true) if value available, (zero, false) otherwise
func TryReceive(ch <-chan int) (int, bool)

// TrySend attempts non-blocking send
// Returns true if sent, false if channel is full/not ready
func TrySend(ch chan<- int, value int) bool

// Merge combines two channels into a single output channel
// Output channel closes when both input channels are closed
func Merge(ch1, ch2 <-chan int) <-chan int

// Multiplex receives from any of the input channels
// Returns each value with its source index
type IndexedValue struct {
    Index int
    Value int
}
func Multiplex(channels ...<-chan int) <-chan IndexedValue
```

---

## Examples

### FirstResponse
```go
ch1 := make(chan string, 1)
ch2 := make(chan string, 1)
ch2 <- "fast"
result := FirstResponse(ch1, ch2)
// result: "fast" (ch2 had a value ready first)
```

### WithTimeout
```go
ch := make(chan string)
result, err := WithTimeout(ch, 100*time.Millisecond)
// err: non-nil (timeout occurred)
// result: ""

ch2 := make(chan string, 1)
ch2 <- "hello"
result2, err2 := WithTimeout(ch2, 100*time.Millisecond)
// result2: "hello", err2: nil
```

### TryReceive / TrySend
```go
ch := make(chan int, 1)
v, ok := TryReceive(ch)  // ok: false (empty)
ch <- 42
v, ok = TryReceive(ch)   // v: 42, ok: true

full := make(chan int)  // unbuffered
ok := TrySend(full, 1)  // ok: false (no receiver)
```

### Merge
```go
ch1 := Producer([]int{1, 3, 5})
ch2 := Producer([]int{2, 4, 6})
merged := Merge(ch1, ch2)
// merged produces: 1, 2, 3, 4, 5, 6 (order may vary)
```

---

## Instructions

1. Implement `FirstResponse` using select with two channel cases
2. Implement `WithTimeout` using select with channel and time.After
3. Implement `TryReceive` using select with default case
4. Implement `TrySend` using select with default case
5. Implement `Merge` - spawn goroutine(s) to forward values from both channels
6. Implement `Multiplex` for arbitrary number of channels
7. Run tests with `go test -v`

---

## Hints

### Basic
- `select` syntax is like `switch` but for channel operations
- Each case is a channel send or receive
- `default` case runs if no other case is ready (non-blocking)
- `time.After(d)` returns a channel that receives after duration d

### Intermediate
- For `Merge`, track when each input closes using a counter or sync.WaitGroup
- For `Multiplex`, you might need a goroutine per channel
- Remember: receiving from a closed channel returns immediately (zero value)
- Check `ok` value when receiving to detect closed channels

### Solution Pattern
```go
select {
case v := <-ch1:
    return v
case v := <-ch2:
    return v
case <-time.After(timeout):
    return "", errors.New("timeout")
default:
    // non-blocking fallback
}
```

---

## Think About

1. What happens if both channels in FirstResponse have values ready?
2. Why does time.After return a channel instead of just blocking?
3. How would you implement a "try receive with timeout" (different from non-blocking)?
4. What's the memory cost of using time.After in a loop?

---

## What This Teaches

- **select statement** - Go's control structure for channel multiplexing
- **Timeout pattern** - Using time.After with select for deadlines
- **Non-blocking operations** - Using default case for try semantics
- **Channel merging** - Combining multiple channels into one
- **Random selection** - When multiple cases are ready, one is chosen randomly
- **Real-world application** - These patterns appear in network programming, UI events, etc.
