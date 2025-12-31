# Exercise 09: Fan-Out/Fan-In

**Learning Goal:** Distribute work to multiple goroutines and merge results back into a single stream

**Difficulty:** Tier 3 - Integration
**Estimated Time:** 45-50 minutes

---

## Problem Description

**Fan-out** distributes work from one channel to multiple goroutines. **Fan-in** collects results from multiple channels into one. Together, they enable parallel processing pipelines.

This pattern is perfect when you have:
- A stream of items to process
- Independent processing (items don't depend on each other)
- A need to merge results into a single output

Think: splitting a deck of cards among players, then collecting all played cards into one pile.

---

## Function Signatures

```go
// FanOut distributes items from input to n output channels
// Returns slice of output channels
func FanOut[T any](input <-chan T, n int) []<-chan T

// FanIn merges multiple input channels into a single output channel
func FanIn[T any](inputs ...<-chan T) <-chan T

// ParallelMap applies fn to each item using n workers
func ParallelMap[T, R any](items []T, n int, fn func(T) R) []R

// ParallelMapStream streams results as they complete
func ParallelMapStream[T, R any](input <-chan T, n int, fn func(T) R) <-chan R

// ProcessWithFanOut processes input with fan-out/fan-in pattern
func ProcessWithFanOut[T, R any](input <-chan T, n int, process func(T) R) <-chan R
```

---

## Examples

### FanOut
```go
input := make(chan int, 3)
input <- 1
input <- 2
input <- 3
close(input)

outputs := FanOut(input, 2)
// len(outputs) == 2
// Values 1, 2, 3 are distributed between outputs[0] and outputs[1]
```

### FanIn
```go
ch1 := make(chan int, 2)
ch2 := make(chan int, 2)
ch1 <- 1
ch1 <- 3
ch2 <- 2
ch2 <- 4
close(ch1)
close(ch2)

merged := FanIn(ch1, ch2)
// merged receives: 1, 2, 3, 4 (order may vary)
```

### ParallelMap
```go
items := []int{1, 2, 3, 4, 5}
results := ParallelMap(items, 3, func(x int) int {
    time.Sleep(100 * time.Millisecond) // Simulate work
    return x * 2
})
// results: [2, 4, 6, 8, 10] (order preserved)
// Total time: ~200ms (not 500ms)
```

### ProcessWithFanOut
```go
input := Generator(10) // produces 0-9
output := ProcessWithFanOut(input, 4, func(x int) string {
    return fmt.Sprintf("processed-%d", x)
})

for result := range output {
    fmt.Println(result) // "processed-0", "processed-1", etc.
}
```

---

## Instructions

1. Implement `FanOut` - spawn goroutine to distribute input values round-robin or on-demand
2. Implement `FanIn` - spawn goroutines to forward from each input to shared output
3. Implement `ParallelMap` - process slice items in parallel, preserve order
4. Implement `ParallelMapStream` - streaming version (order not preserved)
5. Implement `ProcessWithFanOut` - combine FanOut + process + FanIn
6. Run tests with `go test -v -race`

---

## Hints

### Basic
- FanOut: one goroutine reads input, sends to outputs in round-robin or whoever is ready
- FanIn: one goroutine per input channel, all send to same output
- Use sync.WaitGroup to know when to close output channels
- Remember to close output channels when done

### Intermediate
- For FanOut, `select` with multiple cases can send to first available output
- For FanIn, simpler approach: spawn goroutine per input, each forwards to output
- ParallelMap with order: use indexed results with sync
- ParallelMapStream: workers read from shared input, write to shared output

### Solution Pattern
```go
func FanIn[T any](inputs ...<-chan T) <-chan T {
    out := make(chan T)
    var wg sync.WaitGroup

    for _, ch := range inputs {
        wg.Add(1)
        go func(c <-chan T) {
            defer wg.Done()
            for v := range c {
                out <- v
            }
        }(ch)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}
```

---

## Think About

1. What's the difference between FanOut with round-robin vs on-demand distribution?
2. How does ParallelMap preserve order while processing in parallel?
3. When would you choose streaming (ParallelMapStream) vs batch (ParallelMap)?
4. What happens if one worker panics in ProcessWithFanOut?

---

## What This Teaches

- **Fan-out pattern** - Work distribution to multiple processors
- **Fan-in pattern** - Result merging from multiple sources
- **Parallel processing** - Concurrent transformation of data
- **Order preservation** - Techniques for maintaining order in parallel
- **Pipeline composition** - Building larger patterns from smaller ones
- **Generics in Go** - Type-safe concurrent patterns
