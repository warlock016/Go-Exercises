# Exercise 15: Concurrent Processor (Capstone)

**Learning Goal:** Build a production-grade concurrent processing system using all patterns learned

**Difficulty:** Tier 4 - Mastery (Capstone)
**Estimated Time:** 75-90 minutes

---

## Problem Description

This capstone exercise combines everything: goroutines, channels, select, context, mutexes, worker pools, pipelines, rate limiting, and graceful shutdown.

You'll build a **Concurrent Job Processor** that:
1. Accepts jobs from multiple sources
2. Processes them with bounded concurrency
3. Handles failures with retry and backoff
4. Collects metrics
5. Shuts down gracefully

This mirrors real-world systems like task queues, stream processors, and batch job runners.

---

## Function Signatures

```go
// Job represents a unit of work
type Job struct {
    ID       string
    Payload  interface{}
    Priority int  // Higher = more urgent
    Retries  int
    MaxRetry int
}

// JobResult contains the result of processing
type JobResult struct {
    JobID    string
    Success  bool
    Error    error
    Duration time.Duration
}

// Metrics tracks processor statistics
type Metrics struct {
    Submitted  int64
    Completed  int64
    Failed     int64
    Retried    int64
    InFlight   int64
    AvgLatency time.Duration
}

// ProcessorConfig configures the processor
type ProcessorConfig struct {
    Workers       int
    QueueSize     int
    RateLimit     int              // Jobs per second, 0 = unlimited
    RetryDelay    time.Duration
    MaxRetryDelay time.Duration
}

// Processor is the main concurrent job processor
type Processor struct {
    // Define fields
}

func NewProcessor(config ProcessorConfig, handler func(context.Context, Job) error) *Processor

// Submit adds a job to the processor
func (p *Processor) Submit(job Job) error

// SubmitBatch adds multiple jobs
func (p *Processor) SubmitBatch(jobs []Job) error

// Results returns channel of completed job results
func (p *Processor) Results() <-chan JobResult

// Metrics returns current metrics
func (p *Processor) Metrics() Metrics

// Start begins processing
func (p *Processor) Start(ctx context.Context) error

// Shutdown gracefully stops the processor
func (p *Processor) Shutdown(ctx context.Context) error

// Wait blocks until all submitted jobs are processed
func (p *Processor) Wait() error
```

---

## Examples

### Basic Usage
```go
config := ProcessorConfig{
    Workers:    5,
    QueueSize:  100,
    RateLimit:  10,  // 10 jobs/second
    RetryDelay: 100 * time.Millisecond,
}

processor := NewProcessor(config, func(ctx context.Context, job Job) error {
    // Process the job
    return doWork(job.Payload)
})

ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go processor.Start(ctx)

// Submit jobs
processor.Submit(Job{ID: "1", Payload: data1})
processor.Submit(Job{ID: "2", Payload: data2})

// Collect results
go func() {
    for result := range processor.Results() {
        if result.Success {
            fmt.Printf("Job %s completed in %v\n", result.JobID, result.Duration)
        } else {
            fmt.Printf("Job %s failed: %v\n", result.JobID, result.Error)
        }
    }
}()

// Graceful shutdown
processor.Shutdown(context.Background())
```

### With Metrics
```go
// Periodically check metrics
go func() {
    ticker := time.NewTicker(time.Second)
    for range ticker.C {
        m := processor.Metrics()
        fmt.Printf("In-flight: %d, Completed: %d, Failed: %d\n",
            m.InFlight, m.Completed, m.Failed)
    }
}()
```

---

## Instructions

1. Design the `Processor` struct with all required fields
2. Implement `NewProcessor` - initialize workers, channels, rate limiter
3. Implement `Submit` and `SubmitBatch` - add jobs to queue
4. Implement `Start` - spawn workers, begin processing
5. Implement worker logic - process jobs, handle retries with backoff
6. Implement `Results` - stream results as they complete
7. Implement `Metrics` - thread-safe metrics collection
8. Implement `Shutdown` - stop accepting jobs, drain queue, wait for workers
9. Implement `Wait` - block until queue is empty
10. Run tests with `go test -v -race`

---

## Requirements

### Core Functionality
- [ ] Workers pull from shared job queue
- [ ] Rate limiting controls job processing speed
- [ ] Failed jobs retry with exponential backoff
- [ ] Jobs exceeding max retries go to results as failed
- [ ] Metrics update in real-time

### Concurrency Safety
- [ ] All public methods are thread-safe
- [ ] No data races (passes `go test -race`)
- [ ] Proper channel closing sequence

### Graceful Shutdown
- [ ] Stop accepting new jobs
- [ ] Complete in-flight jobs (with timeout)
- [ ] Close results channel after last result
- [ ] Return error if shutdown times out

### Edge Cases
- [ ] Submit after shutdown returns error
- [ ] Empty queue shutdown is fast
- [ ] Handler panic doesn't crash processor

---

## Hints

### Architecture
```
                    ┌─────────────┐
   Submit() ───────>│  Job Queue  │
                    │ (buffered)  │
                    └──────┬──────┘
                           │
           ┌───────────────┼───────────────┐
           │               │               │
           ▼               ▼               ▼
      ┌─────────┐    ┌─────────┐    ┌─────────┐
      │ Worker 1│    │ Worker 2│    │ Worker N│
      └────┬────┘    └────┬────┘    └────┬────┘
           │               │               │
           └───────────────┼───────────────┘
                           │
                    ┌──────▼──────┐
                    │   Results   │──────> Results()
                    │  (channel)  │
                    └─────────────┘
```

### Key Components
- Job queue: buffered channel
- Results: buffered channel
- Rate limiter: time.Ticker or token bucket
- Metrics: atomic counters or mutex-protected struct
- Shutdown: done channel + WaitGroup

### Solution Patterns
```go
// Worker loop
for {
    select {
    case <-ctx.Done():
        return
    case job, ok := <-p.jobs:
        if !ok {
            return
        }
        p.processWithRetry(ctx, job)
    }
}

// Retry with backoff
delay := p.config.RetryDelay
for attempt := 0; attempt <= job.MaxRetry; attempt++ {
    err := p.handler(ctx, job)
    if err == nil {
        return nil
    }
    if attempt < job.MaxRetry {
        select {
        case <-ctx.Done():
            return ctx.Err()
        case <-time.After(delay):
            delay = min(delay*2, p.config.MaxRetryDelay)
        }
    }
}
```

---

## Think About

1. How would you add job priority (process high-priority first)?
2. How would you handle a "poison" job that always fails?
3. How would you add observability (logging, tracing)?
4. How would you scale this across multiple machines?

---

## What This Teaches

This capstone integrates ALL concurrency patterns:
- **Goroutines & Channels** - Worker pool architecture
- **Select** - Non-blocking checks, timeouts
- **Context** - Cancellation propagation
- **Mutex** - Protecting shared state (metrics)
- **WaitGroup** - Coordinating worker completion
- **Rate Limiting** - Controlling throughput
- **Retry & Backoff** - Handling transient failures
- **Graceful Shutdown** - Production-grade lifecycle

Completing this means you've mastered Go concurrency!
