# Exercise 08.5: Worker Pool Patterns

**Learning Goal:** Master worker pool patterns and understand when each applies—specifically, why the streaming pattern self-regulates while the batch pattern requires buffering.

**Difficulty:** Tier 3 - Integration
**Estimated Time:** 60-90 minutes

---

## Background: Why This Exercise Exists

Worker pools are a fundamental concurrency pattern, but there are **two distinct usage patterns** with different requirements:

| Pattern | Description | Buffer Requirement |
|---------|-------------|-------------------|
| **Streaming** | Producer and consumer run concurrently | Small (numWorkers) |
| **Batch** | Submit all jobs, then read all results | Large (≥ total jobs) |

Understanding **why** these differ is crucial for writing correct concurrent code.

---

## The Two Patterns Explained

### Pattern 1: Streaming (Self-Regulating)

```
Producer ──► [jobs channel] ──► Workers ──► [results channel] ──► Consumer
    │                              │                                 │
    └──────────── Backpressure ◄───┴────────── flows back ◄──────────┘
```

**How backpressure works:**
1. Results buffer fills up → Workers block trying to send
2. Workers blocked → Can't receive new jobs
3. Jobs buffer fills → Producer blocks on Submit()
4. **The system self-regulates!**

**This works because producer and consumer run CONCURRENTLY.** When consumer reads a result, it unblocks a worker, which can then accept a new job, which unblocks the producer.

### Pattern 2: Batch (Requires Buffering)

```
Producer ──► [jobs channel] ──► Workers ──► [results channel] ──► (waiting)
    │                              │                                 │
    │         SEQUENTIAL           │        NO CONSUMER YET          │
    └───── submit ALL first ───────┴──────── then read ALL ──────────┘
```

**Why this deadlocks with small buffers:**
1. Producer submits jobs until buffer full
2. Workers process, fill results buffer
3. Workers block (results full, nobody receiving)
4. Producer blocks (jobs full, workers stuck)
5. **Deadlock!** Consumer never starts because producer hasn't finished.

**Solution:** Buffer must hold ALL results, OR restructure to use streaming pattern.

---

## Function Signatures

### 1. StreamingPool (The Self-Regulating Pattern)

```go
// StreamingPool processes jobs from a channel using a fixed worker pool.
// Returns a results channel that closes when all jobs are processed.
//
// This pattern naturally self-regulates through backpressure:
// - When results buffer is full, workers block
// - Blocked workers can't accept new jobs
// - Producer naturally slows down
//
// Works correctly with small buffers (numWorkers is sufficient).
func StreamingPool(
    jobs <-chan Job,
    numWorkers int,
    process func(Job) Result,
) <-chan Result
```

**Usage:**
```go
// Producer runs concurrently
jobs := make(chan Job)
go func() {
    for i := 0; i < 1000; i++ {
        jobs <- Job{ID: i, Data: i * 10}
    }
    close(jobs)
}()

// Consumer runs concurrently (in main goroutine)
results := StreamingPool(jobs, 4, processFunc)
for r := range results {
    fmt.Println(r)
}
```

### 2. ProcessBatch (When You Need All Results)

```go
// ProcessBatch processes a slice of jobs and returns all results.
// This is a convenience function for when you have all jobs upfront.
//
// Note: This pattern requires internal buffering proportional to job count
// because results must be stored before the caller can read them.
func ProcessBatch(
    jobs []Job,
    numWorkers int,
    process func(Job) Result,
) []Result
```

**Usage:**
```go
jobs := []Job{{ID: 1}, {ID: 2}, {ID: 3}}
results := ProcessBatch(jobs, 2, processFunc)
// results contains all 3 results
```

### 3. PoolWithCancel (Graceful Shutdown)

```go
// PoolWithCancel processes jobs until context is cancelled.
// When cancelled:
// - No new jobs are accepted
// - In-flight jobs complete (graceful)
// - Results channel closes after all workers finish
func PoolWithCancel(
    ctx context.Context,
    jobs <-chan Job,
    numWorkers int,
    process func(Job) Result,
) <-chan Result
```

**Usage:**
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

results := PoolWithCancel(ctx, jobs, 4, processFunc)
for r := range results {
    fmt.Println(r)
}
// Loop exits when context times out and workers finish
```

### 4. PoolWithErrors (Error Separation)

```go
// PoolWithErrors processes jobs, separating successful results from errors.
// Returns two channels: one for results, one for errors.
// Both channels close when all jobs are processed.
func PoolWithErrors(
    jobs <-chan Job,
    numWorkers int,
    process func(Job) (Result, error),
) (results <-chan Result, errors <-chan error)
```

**Usage:**
```go
results, errs := PoolWithErrors(jobs, 4, processFunc)

// Must drain both channels (can use select or separate goroutines)
go func() {
    for err := range errs {
        log.Println("Error:", err)
    }
}()
for r := range results {
    fmt.Println(r)
}
```

---

## Type Definitions

```go
type Job struct {
    ID   int
    Data int
}

type Result struct {
    JobID  int
    Output int
    Err    error  // Only used by ProcessBatch; PoolWithErrors uses separate channel
}
```

---

## Instructions

1. **Start with `StreamingPool`** - This is the fundamental pattern. Focus on:
   - Workers that loop on `for job := range jobs`
   - Proper channel closing (sender closes)
   - WaitGroup to know when all workers finish

2. **Implement `ProcessBatch`** using `StreamingPool` internally:
   - Create a jobs channel, send all jobs, close it
   - Run producer in goroutine so consumer can run concurrently
   - Collect results into slice

3. **Add `PoolWithCancel`** - Extend streaming pattern:
   - Check `ctx.Done()` in worker loop
   - Use `select` to receive from jobs OR context

4. **Implement `PoolWithErrors`** - Separate success/error paths:
   - Two output channels (results and errors)
   - Single WaitGroup tracking all workers
   - Close both channels when done

---

## Hints

### Basic: Worker Structure
A worker is a goroutine that loops until input closes:
```go
go func() {
    for job := range jobs {
        result := process(job)
        results <- result
    }
    // Exits when jobs channel closes
}()
```

### Intermediate: Closing Results Channel
The results channel should close when ALL workers finish:
```go
go func() {
    wg.Wait()
    close(results)
}()
```

### Advanced: Context Cancellation
To respect context in workers:
```go
for {
    select {
    case <-ctx.Done():
        return  // Stop accepting new jobs
    case job, ok := <-jobs:
        if !ok {
            return  // Jobs channel closed
        }
        // Process job...
    }
}
```

---

## Think About

1. **Why does `StreamingPool` work with buffer size 1?**
   - Trace through what happens when buffer is full
   - How does the consumer unblock the system?

2. **Why can't you "fix" batch pattern with backpressure?**
   - What would need to change in the caller's code?

3. **What happens to in-flight jobs when context cancels?**
   - Should they complete or be abandoned?
   - How does your implementation handle this?

4. **Why separate error channels in `PoolWithErrors`?**
   - What's the alternative (errors in Result struct)?
   - When is each approach better?

---

## What This Teaches

- **Backpressure**: Self-regulating concurrent systems
- **Channel semantics**: Who closes, when, and why
- **Pattern selection**: Choosing the right pattern for the use case
- **Graceful shutdown**: Context-based cancellation
- **Error handling**: Separating success and failure paths in concurrent code

---

## Common Mistakes to Avoid

1. **Closing from receiver**: Only the SENDER should close a channel
2. **Double close**: Use `sync.Once` if multiple goroutines might close
3. **Forgetting to close**: Results in infinite `range` loop (deadlock)
4. **Blocking in main**: Producer must run in goroutine for streaming pattern
5. **Ignoring context**: Check `ctx.Done()` regularly in long-running workers
