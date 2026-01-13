# Concurrent Reasoning Guide

Practical techniques for reasoning about concurrent Go code with confidence.

---

## Why Concurrency is Hard

Concurrent code is fundamentally different from sequential code:

| Sequential | Concurrent |
|------------|------------|
| One thing happens at a time | Multiple things happen simultaneously |
| Predictable execution order | Non-deterministic execution order |
| Bugs are reproducible | Bugs may appear 1 in 10,000 runs |
| Local reasoning works | Must consider all goroutines |

**The core challenge:** You can't step through concurrent code line-by-line because multiple "lines" execute simultaneously.

---

## Mental Models That Simplify Reasoning

### 1. Ownership Thinking

For every piece of data, ask: **"Who owns this?"**

| Ownership Model | Pattern | Reasoning Complexity |
|-----------------|---------|---------------------|
| **Single owner** | One goroutine reads/writes | Simple - no coordination |
| **Transfer ownership** | Send via channel | Simple - sender gives up access |
| **Shared read-only** | Multiple readers, no writers | Simple - immutable data |
| **Shared read-write** | Mutex required | Complex - avoid if possible |

**Rule:** Prefer ownership transfer (channels) over shared state (mutexes).

```go
// Ownership transfer - data moves between goroutines
jobs <- job      // Producer gives up ownership
job := <-jobs    // Worker takes ownership

// vs Shared state - data accessed by multiple goroutines
mu.Lock()
shared.data = value  // Must remember to lock/unlock
mu.Unlock()
```

### 2. Sequential-First Design

Never start with concurrency. Instead:

1. **Write the sequential version first**
2. **Identify what can run in parallel**
3. **Add concurrency only where it provides value**

```
Sequential:   A → B → C → D → E
                  ↓
Analysis:     A must happen first
              B, C, D are independent
              E needs results from B, C, D
                  ↓
Concurrent:   A → [B, C, D in parallel] → E
```

### 3. Single Responsibility Per Goroutine

Each goroutine should do **one thing**:

```go
// ✅ Good: Single responsibility
func worker(jobs <-chan Job, results chan<- Result) {
    for job := range jobs {
        results <- process(job)
    }
}

// ❌ Bad: Multiple responsibilities
func worker(jobs <-chan Job, results chan<- Result, db *Database, metrics *Metrics) {
    for job := range jobs {
        result := process(job)
        db.Save(result)           // Database access
        metrics.Record(job.ID)    // Metrics
        log.Printf("done: %d", job.ID)  // Logging
        results <- result
    }
}
```

If a goroutine does multiple things, split it into a pipeline or separate goroutines.

### 4. Channel Direction as Documentation

Use directional channels in function signatures:

```go
// Clear ownership: reads from jobs, writes to results
func worker(jobs <-chan Job, results chan<- Result)

// Ambiguous: who closes what?
func worker(jobs chan Job, results chan Result)
```

The types document the data flow and prevent accidental misuse.

### 5. The Closer Owns the Channel

Establish a clear rule: **the sender (or creator) closes the channel**.

```go
// Producer creates and closes
func producer() <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
        close(ch)  // Producer closes
    }()
    return ch
}

// Consumer only reads
func consumer(ch <-chan int) {
    for v := range ch {
        process(v)
    }
    // Never closes ch - not the owner
}
```

---

## Practical Techniques

### 1. Draw the Data Flow

Before writing code, sketch the flow:

```
┌──────────┐     jobs      ┌──────────┐    results    ┌──────────┐
│ Producer │──────────────▶│ Workers  │──────────────▶│ Consumer │
└──────────┘               └──────────┘               └──────────┘
     │                          │                          │
   closes                   range loop                  range loop
   jobs ch                  exits when                  exits when
                           jobs closes                results closes
```

This makes ownership and lifecycle immediately visible.

### 2. Trace Through Scenarios

Pick specific scenarios and trace step-by-step:

```
Scenario: 2 workers, 3 jobs, job2 fails

Timeline:
─────────────────────────────────────────────────────────▶
Producer:  send(j1) send(j2) send(j3) close(jobs)
Worker A:          recv(j1)────process────send(r1)  recv(j3)────process────send(r3) exit
Worker B:                   recv(j2)────fail────send(err)                           exit
Cleanup:                                                                        wg.Wait() close(results) close(errs)
```

If you can't trace it clearly, the design is too complex.

### 3. The "What If" Checklist

For every concurrent design, systematically ask:

**Channel states:**
- [ ] What if the channel is nil?
- [ ] What if the channel is already closed?
- [ ] What if there are no receivers?

**Lifecycle:**
- [ ] What if context is cancelled mid-operation?
- [ ] What if a goroutine panics?
- [ ] What if we need to shut down gracefully?

**Scale:**
- [ ] What if there are 0 items?
- [ ] What if there is 1 item?
- [ ] What if there are 1 million items?

**Timing:**
- [ ] What if the consumer is slower than the producer?
- [ ] What if the producer is slower than the consumer?
- [ ] What if operations happen in a different order?

### 4. Write Down Invariants

State what must **always** be true:

```
Invariants for WorkerPool:
1. Only workers send to results channel
2. Only the cleanup goroutine closes results channel
3. Cleanup runs only after ALL workers exit (wg.Wait)
4. Each job produces exactly one result
5. Workers exit only when jobs channel closes OR context cancels
```

If your code can violate any invariant, you have a bug.

### 5. Describe in 2-3 Sentences

**The complexity test:** If you can't explain a concurrent design in 2-3 simple sentences, it's probably too complex.

```
StreamingPool: "Workers read from jobs, process each one, and send
results. When jobs closes, workers exit. When all workers exit,
results closes."

ProcessBatch: "A goroutine feeds jobs into a channel. Workers process
jobs concurrently. Results are collected into a slice after all
workers finish."
```

If your description requires "unless" or "except when" or "but if", simplify the design.

---

## Simplification Strategies

### Reduce Complexity

| Complex Pattern | Simpler Alternative |
|-----------------|---------------------|
| Shared mutable state | Ownership transfer via channels |
| Multiple mutexes | Single mutex or channel-based coordination |
| Bidirectional communication | Unidirectional pipeline |
| Dynamic goroutine count | Fixed worker pool |
| Custom synchronization | Standard `sync.WaitGroup`, `errgroup` |
| Unbounded concurrency | Bounded worker pool |

### Use Established Patterns

Don't invent—reuse patterns with known properties:

| Pattern | Use When |
|---------|----------|
| **Producer-Consumer** | One source feeds multiple processors |
| **Worker Pool** | Bounded parallelism for a set of jobs |
| **Fan-out/Fan-in** | Distribute work, collect results |
| **Pipeline** | Sequential stages with parallel execution within |
| **Pub-Sub** | Multiple consumers need same data |

Each pattern has documented behavior, failure modes, and testing strategies.

---

## Verification Tools

### Runtime Tools

| Tool | What It Catches | Command |
|------|-----------------|---------|
| Race detector | Data races | `go test -race` |
| Multiple runs | Flaky behavior | `go test -count=100` |
| Go vet | Common mistakes | `go vet ./...` |

### Debugging Tools

| Tool | Use For | Command |
|------|---------|---------|
| Delve | Inspecting goroutine state | `dlv test`, then `goroutines` |
| Stack traces | Finding blocked goroutines | `runtime.Stack()` |
| Goroutine count | Detecting leaks | `runtime.NumGoroutine()` |

### Goroutine Leak Detection Pattern

```go
func TestNoLeaks(t *testing.T) {
    before := runtime.NumGoroutine()

    // Run your concurrent code
    results := WorkerPool(jobs, 5, processFunc)
    for range results {
        // drain
    }

    // Give goroutines time to exit
    time.Sleep(100 * time.Millisecond)

    after := runtime.NumGoroutine()
    if after > before {
        t.Errorf("goroutine leak: %d → %d", before, after)
    }
}
```

---

## Common Pitfalls and Fixes

### Pitfall 1: Loop Variable Capture

```go
// ❌ Bug: All goroutines share same variable
for _, job := range jobs {
    go func() {
        process(job)  // job changes before goroutine runs!
    }()
}

// ✅ Fix: Pass as parameter
for _, job := range jobs {
    go func(j Job) {
        process(j)
    }(job)
}

// ✅ Fix (Go 1.22+): Loop variables are per-iteration
for _, job := range jobs {
    go func() {
        process(job)  // Safe in Go 1.22+
    }()
}
```

### Pitfall 2: Forgetting to Close

```go
// ❌ Bug: Consumer blocks forever
func produce() <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
        // Missing: close(ch)
    }()
    return ch
}

// Consumer
for v := range produce() {  // Never terminates!
    fmt.Println(v)
}
```

### Pitfall 3: Send on Closed Channel

```go
// ❌ Bug: Panic if channel closed while goroutine running
go func() {
    for {
        ch <- value  // Panic if ch closed!
    }
}()
close(ch)

// ✅ Fix: Use done channel or context
go func() {
    for {
        select {
        case <-done:
            return
        case ch <- value:
        }
    }
}()
close(done)  // Signal stop, don't close ch
```

### Pitfall 4: Nil Channel Operations

```go
// ❌ Bug: Blocks forever (no error, no panic)
var ch chan int  // nil
ch <- 1          // Blocks forever
<-ch             // Blocks forever

// ✅ Fix: Always initialize
ch := make(chan int)
```

---

## Decision Framework

When designing concurrent code, follow this sequence:

```
1. Can I avoid concurrency?
   └─ Yes → Do it sequentially
   └─ No → Continue

2. Can I use ownership transfer (channels)?
   └─ Yes → Use producer-consumer or pipeline
   └─ No → Continue

3. Can I make shared data read-only?
   └─ Yes → No synchronization needed
   └─ No → Continue

4. Do I need a mutex?
   └─ Use the simplest: sync.Mutex or sync.RWMutex
   └─ Keep critical sections small
   └─ Document what the mutex protects
```

---

## Quick Reference Card

### Safe Operations

| Operation | On Normal Channel | On Nil Channel | On Closed Channel |
|-----------|------------------|----------------|-------------------|
| Send | Blocks until received | **Blocks forever** | **Panic** |
| Receive | Blocks until sent | **Blocks forever** | Returns zero value |
| Close | Closes channel | **Panic** | **Panic** |

### Ownership Rules

1. Creator owns closing
2. Sender closes, never receiver
3. Transfer ownership via channel send
4. One closer per channel (use `sync.Once` if needed)

### The Golden Test

> "If you can't explain it in 2-3 sentences, simplify it."

---

## Related Resources

- [CHANNELS_GUIDE.md](CHANNELS_GUIDE.md) - Channel fundamentals
- [CONTEXT_GUIDE.md](CONTEXT_GUIDE.md) - Context and cancellation
- [MUTEX_GUIDE.md](MUTEX_GUIDE.md) - Mutex patterns
- [CONCURRENCY_DEBUGGING_GUIDE.md](CONCURRENCY_DEBUGGING_GUIDE.md) - Debugging techniques
- [STREAMING_WORKER_PATTERNS.md](STREAMING_WORKER_PATTERNS.md) - Production patterns

---

**Last Updated:** 2026-01-13
**Context:** Mental models and techniques for reasoning about concurrent Go code
