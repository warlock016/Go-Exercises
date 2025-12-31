# Exercise 05: Done Channel Pattern

**Learning Goal:** Implement cancellation and shutdown signaling using the done channel idiom

**Difficulty:** Tier 2 - Application
**Estimated Time:** 30-35 minutes

---

## Problem Description

The done channel pattern is Go's idiomatic way to signal cancellation or completion to goroutines. Unlike context (which we'll cover later), the done channel is a raw primitive you build yourself—understanding it deeply helps you understand how context.Done() works internally.

The key insight: closing a channel broadcasts to ALL receivers simultaneously. A receive on a closed channel returns immediately with the zero value.

---

## Function Signatures

```go
// Worker processes items until done channel is closed
// Returns the count of items processed
func Worker(items <-chan int, done <-chan struct{}) int

// CancellableLoop runs work function repeatedly until done is closed
// Returns how many times work was executed
func CancellableLoop(work func(), done <-chan struct{}) int

// Broadcaster sends value to all output channels, respecting done
// Returns number of channels successfully sent to
func Broadcaster(value int, outputs []chan<- int, done <-chan struct{}) int

// Generator produces integers 0, 1, 2, ... until done is closed
func Generator(done <-chan struct{}) <-chan int

// Timeout returns a done channel that closes after duration
func Timeout(d time.Duration) <-chan struct{}

// Merge combines multiple done channels into one
// Result closes when ANY input closes
func MergeCancel(done1, done2 <-chan struct{}) <-chan struct{}
```

---

## Examples

### Worker
```go
items := make(chan int, 3)
items <- 1
items <- 2
items <- 3

done := make(chan struct{})
close(done) // Signal cancellation immediately

count := Worker(items, done)
// count: 0 (stopped before processing due to done)
```

### Generator with Timeout
```go
done := Timeout(100 * time.Millisecond)
gen := Generator(done)

var values []int
for v := range gen {
    values = append(values, v)
}
// values contains 0, 1, 2, ... until timeout closed done
```

### MergeCancel
```go
done1 := make(chan struct{})
done2 := make(chan struct{})

merged := MergeCancel(done1, done2)

close(done2) // Close either one
<-merged     // merged is now closed too
```

---

## Instructions

1. Implement `Worker` that processes items but checks done channel on each iteration
2. Implement `CancellableLoop` using select to check done between work calls
3. Implement `Broadcaster` that sends to multiple channels with cancellation support
4. Implement `Generator` that produces incrementing integers until cancelled
5. Implement `Timeout` that creates a done channel closing after specified duration
6. Implement `MergeCancel` that closes when either input closes
7. Run tests with `go test -v`

---

## Hints

### Basic
- `chan struct{}` uses zero memory per element—it's purely for signaling
- Closing a channel causes all blocked receives to return immediately
- `select` with `default` for non-blocking checks
- A receive from closed channel: `<-done` returns `struct{}{}` and `ok=false`

### Intermediate
- For Worker: use select inside a for loop to check both channels
- For Generator: spawn a goroutine that sends until done, then closes output
- For MergeCancel: spawn a goroutine with select on both done channels
- Remember: you cannot detect if a channel is closed without receiving from it

### Solution Pattern
```go
for {
    select {
    case <-done:
        return count
    case item, ok := <-items:
        if !ok {
            return count // items channel closed
        }
        // process item
        count++
    }
}
```

---

## Think About

1. Why use `chan struct{}` instead of `chan bool` for done channels?
2. What happens if you forget to close the done channel?
3. How does this pattern relate to context.Context?
4. Why is closing preferred over sending a value for cancellation?

---

## What This Teaches

- **Done channel idiom** - The foundation of cancellation in Go
- **Broadcast via close** - One close, many receivers notified
- **Select with done** - Checking for cancellation in loops
- **Cleanup coordination** - Signaling goroutines to exit cleanly
- **Context foundation** - Understanding what context.Done() does internally
