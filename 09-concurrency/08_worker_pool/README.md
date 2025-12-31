# Exercise 08: Worker Pool

**Learning Goal:** Build a bounded pool of workers that process jobs from a shared queue

**Difficulty:** Tier 3 - Integration
**Estimated Time:** 45-50 minutes

---

## Problem Description

A **worker pool** is one of the most common concurrency patterns. Instead of spawning unlimited goroutines, you create a fixed number of workers that pull jobs from a shared queue. This gives you:

- **Bounded resource usage** - Memory and CPU are predictable
- **Backpressure** - If workers are busy, job queue fills, senders block
- **Graceful degradation** - System stays responsive under load

The pattern: N workers, 1 job channel, 1 results channel. Workers loop forever receiving jobs until the job channel closes.

---

## Function Signatures

```go
// Job represents a unit of work
type Job struct {
    ID    int
    Input int
}

// Result represents the output of processing a job
type Result struct {
    JobID  int
    Output int
    Err    error
}

// WorkerPool manages a pool of workers
type WorkerPool struct {
    // Define fields
}

// NewWorkerPool creates a pool with the specified number of workers
func NewWorkerPool(numWorkers int, process func(Job) Result) *WorkerPool

// Submit adds a job to the pool (blocks if pool is full)
func (p *WorkerPool) Submit(job Job)

// Results returns the channel to receive results from
func (p *WorkerPool) Results() <-chan Result

// Shutdown signals workers to stop and waits for completion
func (p *WorkerPool) Shutdown()

// Simple function-based worker pool
func ProcessJobs(jobs []Job, numWorkers int, process func(Job) Result) []Result
```

---

## Examples

### ProcessJobs (Simple)
```go
jobs := []Job{
    {ID: 1, Input: 10},
    {ID: 2, Input: 20},
    {ID: 3, Input: 30},
}

results := ProcessJobs(jobs, 2, func(j Job) Result {
    return Result{JobID: j.ID, Output: j.Input * 2}
})
// results contain: [{1, 20, nil}, {2, 40, nil}, {3, 60, nil}]
// Order may vary due to concurrent processing
```

### WorkerPool (Streaming)
```go
pool := NewWorkerPool(3, func(j Job) Result {
    time.Sleep(100 * time.Millisecond) // Simulate work
    return Result{JobID: j.ID, Output: j.Input * 2}
})

// Submit jobs (can be done from multiple goroutines)
go func() {
    for i := 0; i < 10; i++ {
        pool.Submit(Job{ID: i, Input: i})
    }
    pool.Shutdown() // Signal no more jobs
}()

// Collect results
for result := range pool.Results() {
    fmt.Printf("Job %d: %d\n", result.JobID, result.Output)
}
```

---

## Instructions

1. Implement `ProcessJobs` - batch processing with worker pool
2. Define `WorkerPool` struct with necessary fields
3. Implement `NewWorkerPool` - spawn workers, return pool
4. Implement `Submit` - send job to workers
5. Implement `Results` - return read-only results channel
6. Implement `Shutdown` - close job channel, wait for workers
7. Run tests with `go test -v -race`

---

## Hints

### Basic
- Workers receive from job channel in a for-range loop
- Results channel collects output from all workers
- Use sync.WaitGroup to track when all workers are done
- Close results channel only after all workers exit

### Intermediate
- Job channel capacity controls backpressure (try 10-100)
- Results channel should be buffered to prevent blocking workers
- Shutdown: close jobs channel → workers exit → WaitGroup done → close results
- Consider: what if Submit is called after Shutdown?

### Solution Pattern
```go
func worker(jobs <-chan Job, results chan<- Result, process func(Job) Result, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        results <- process(job)
    }
}

func NewWorkerPool(numWorkers int, process func(Job) Result) *WorkerPool {
    jobs := make(chan Job, 100)
    results := make(chan Result, 100)
    var wg sync.WaitGroup

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go worker(jobs, results, process, &wg)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    return &WorkerPool{jobs: jobs, results: results}
}
```

---

## Think About

1. What happens if one worker panics? How could you recover?
2. How would you implement job priority (high priority jobs first)?
3. What's the optimal number of workers for CPU-bound vs I/O-bound tasks?
4. How could you add job timeout per-job?

---

## What This Teaches

- **Worker pool pattern** - Bounded concurrency for resource control
- **Channel-based job dispatch** - Fair distribution via blocking receive
- **Graceful shutdown** - Clean termination of worker goroutines
- **Backpressure** - Channel buffers as flow control
- **WaitGroup coordination** - Knowing when all workers are done
- **Production pattern** - Used in web servers, batch processors, etc.
