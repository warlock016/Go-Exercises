# Exercise 03: Buffered Channels

**Learning Goal:** Understand buffered vs unbuffered channels and when to use each

**Difficulty:** Tier 1 - Introduction
**Estimated Time:** 30-35 minutes

---

## Problem Description

Buffered channels have capacity to hold values without a receiver being ready. A send blocks only when the buffer is full; a receive blocks only when the buffer is empty.

This decouples sender and receiver timing, enabling patterns like producer-consumer where production can temporarily outpace consumption.

In this exercise, you'll explore buffer capacity, the `len()` and `cap()` functions on channels, and build a batching processor.

---

## Function Signatures

```go
// BufferDemo sends values to a buffered channel and returns them in order
// The buffer size determines how many sends can happen before blocking
func BufferDemo(values []int, bufferSize int) []int

// Producer sends all values to a buffered channel and returns the channel
// The channel is closed after all values are sent
func Producer(values []int, bufferSize int) <-chan int

// Consumer receives all values from a channel and returns them as a slice
func Consumer(ch <-chan int) []int

// BatchCollector collects values into batches of the specified size
// Returns a channel of slices, each containing up to batchSize values
func BatchCollector(in <-chan int, batchSize int) <-chan []int

// Semaphore uses a buffered channel to limit concurrent operations
// Returns a function that runs the given work with limited concurrency
func Semaphore(limit int) func(work func())
```

---

## Examples

### BufferDemo
```
Input:  BufferDemo([]int{1, 2, 3}, 2)
Output: []int{1, 2, 3}

// With buffer size 2, first two sends don't block
// Third send blocks until we start receiving
```

### Producer/Consumer
```
producer := Producer([]int{1, 2, 3, 4, 5}, 3)
result := Consumer(producer)
// result: []int{1, 2, 3, 4, 5}
```

### BatchCollector
```
in := Producer([]int{1, 2, 3, 4, 5, 6, 7}, 10)
batches := BatchCollector(in, 3)
// Produces: []int{1,2,3}, []int{4,5,6}, []int{7}
```

### Semaphore
```
sem := Semaphore(2)  // Max 2 concurrent operations

// These run with max 2 at a time:
for i := 0; i < 10; i++ {
    go sem(func() {
        // work that should be limited
    })
}
```

---

## Instructions

1. Implement `BufferDemo` - observe how buffer size affects blocking
2. Implement `Producer` - send values to buffered channel, close when done
3. Implement `Consumer` - collect all values from channel into slice
4. Implement `BatchCollector` - group values into fixed-size batches
5. Implement `Semaphore` - use buffered channel to limit concurrency
6. Run tests with `go test -v`

---

## Hints

### Basic
- Create buffered channel: `make(chan T, capacity)`
- `len(ch)` returns number of elements in buffer
- `cap(ch)` returns buffer capacity
- Buffered channel of size N allows N sends without blocking

### Intermediate
- Producer should close the channel after sending all values
- BatchCollector should emit partial batches at the end
- Semaphore pattern: buffered channel acts as a token pool
- Acquire token: `ch <- struct{}{}`; Release: `<-ch`

### Solution Pattern
```go
// Semaphore using buffered channel
func Semaphore(limit int) func(work func()) {
    tokens := make(chan struct{}, limit)
    return func(work func()) {
        tokens <- struct{}{}  // acquire
        defer func() { <-tokens }()  // release
        work()
    }
}
```

---

## Think About

1. Why might you choose a buffered channel over unbuffered?
2. What happens if the buffer size equals the number of values sent?
3. How does the semaphore pattern prevent resource exhaustion?
4. What's the difference between `len(ch)` on buffered vs unbuffered channels?

---

## What This Teaches

- **Buffered channels** - `make(chan T, n)` creates a channel with buffer capacity n
- **Non-blocking sends** - Sends don't block until buffer is full
- **len() and cap()** - Inspect current usage and total capacity
- **Producer-consumer** - Decoupling production and consumption rates
- **Batching pattern** - Collecting values into groups
- **Semaphore pattern** - Using buffered channel to limit concurrency
- **Buffer sizing** - Trade-offs between memory and throughput
