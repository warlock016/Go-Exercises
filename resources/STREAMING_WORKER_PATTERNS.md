# Streaming Worker Patterns Guide

A reference for building long-running, production-grade worker pools in Go.

---

## Batch vs Streaming: Two Mental Models

### Batch Processing (Exercise Pattern)

```
Known work → Process all → Close channels → Exit
```

- Jobs channel **closes** when producer is done
- Workers exit when `for job := range jobs` ends
- Results channel closes after all workers finish
- **Lifecycle:** Start → Process → Cleanup → Exit

### Streaming Processing (Production Pattern)

```
Continuous work → Process indefinitely → Only stop on signal
```

- Jobs channel **stays open** indefinitely
- Workers run in `for { select { ... } }` loops
- Context cancellation drives shutdown, not channel closing
- **Lifecycle:** Start → Run Forever → Graceful Shutdown (rare)

---

## Pattern 1: Context-Driven Long-Running Worker

Workers that run indefinitely until explicitly cancelled:

```go
func LongRunningPool(ctx context.Context, jobs <-chan Job, numWorkers int, process func(Job) Result) <-chan Result {
    results := make(chan Result)
    var wg sync.WaitGroup

    for range numWorkers {
        wg.Go(func() {
            for {
                select {
                case <-ctx.Done():
                    // Shutdown signal received
                    return
                case job, ok := <-jobs:
                    if !ok {
                        // Producer closed (unusual in streaming)
                        return
                    }
                    // Safe send: also check ctx during send
                    select {
                    case results <- process(job):
                    case <-ctx.Done():
                        return
                    }
                }
            }
        })
    }

    // Cleanup goroutine
    go func() {
        wg.Wait()
        close(results)
    }()

    return results
}
```

**Key differences from batch:**
- `for { select { ... } }` instead of `for job := range jobs`
- Context checked at every blocking operation
- Channel closing is cleanup, not the shutdown signal

---

## Pattern 2: Supervisor with Worker Restart

Workers that automatically restart on failure:

```go
type Supervisor struct {
    ctx        context.Context
    cancel     context.CancelFunc
    jobs       <-chan Job
    results    chan Result
    numWorkers int
    process    func(Job) (Result, error)
    wg         sync.WaitGroup
}

func NewSupervisor(ctx context.Context, jobs <-chan Job, numWorkers int, process func(Job) (Result, error)) *Supervisor {
    ctx, cancel := context.WithCancel(ctx)
    return &Supervisor{
        ctx:        ctx,
        cancel:     cancel,
        jobs:       jobs,
        results:    make(chan Result),
        numWorkers: numWorkers,
        process:    process,
    }
}

func (s *Supervisor) Start() <-chan Result {
    for i := range s.numWorkers {
        s.startWorker(i)
    }
    return s.results
}

func (s *Supervisor) startWorker(id int) {
    s.wg.Go(func() {
        for {
            select {
            case <-s.ctx.Done():
                return
            case job, ok := <-s.jobs:
                if !ok {
                    return
                }
                s.processWithRecovery(id, job)
            }
        }
    })
}

func (s *Supervisor) processWithRecovery(workerID int, job Job) {
    defer func() {
        if r := recover(); r != nil {
            // Log panic, don't crash the worker
            log.Printf("worker %d recovered from panic: %v", workerID, r)
            // Worker continues to next job
        }
    }()

    result, err := s.process(job)
    if err != nil {
        // Log error but continue processing
        log.Printf("worker %d: job %d failed: %v", workerID, job.ID, err)
        return
    }

    select {
    case s.results <- result:
    case <-s.ctx.Done():
    }
}

func (s *Supervisor) Shutdown() {
    s.cancel()
    s.wg.Wait()
    close(s.results)
}
```

**Key features:**
- Workers recover from panics instead of crashing
- Errors are logged but don't stop processing
- Clean shutdown via context cancellation

---

## Pattern 3: Graceful Shutdown with Drain Period

Proper shutdown sequence for production systems:

```go
type GracefulPool struct {
    ctx         context.Context
    cancel      context.CancelFunc
    jobs        chan Job          // Owned by pool (can close)
    results     chan Result
    drainPeriod time.Duration
    wg          sync.WaitGroup
}

func (p *GracefulPool) Shutdown() error {
    // Step 1: Stop accepting new work
    close(p.jobs)

    // Step 2: Give workers time to finish in-flight work
    done := make(chan struct{})
    go func() {
        p.wg.Wait()
        close(done)
    }()

    select {
    case <-done:
        // All workers finished cleanly
        close(p.results)
        return nil
    case <-time.After(p.drainPeriod):
        // Drain period exceeded, force cancel
        p.cancel()
        <-done  // Wait for workers to exit
        close(p.results)
        return errors.New("shutdown: drain period exceeded, forced cancellation")
    }
}
```

**Shutdown sequence:**
1. Close jobs channel (stop new work)
2. Wait for in-flight work (drain period)
3. Force cancel if timeout exceeded
4. Close results channel (cleanup)

---

## Pattern 4: Health Monitoring

Track worker health for observability:

```go
type HealthyPool struct {
    workers     []*workerState
    mu          sync.RWMutex
    healthCheck chan struct{}
}

type workerState struct {
    id           int
    lastActive   time.Time
    jobsHandled  int64
    errors       int64
    currentJob   *Job  // nil if idle
}

func (p *HealthyPool) Health() map[string]any {
    p.mu.RLock()
    defer p.mu.RUnlock()

    activeWorkers := 0
    totalJobs := int64(0)
    totalErrors := int64(0)

    for _, w := range p.workers {
        if w.currentJob != nil {
            activeWorkers++
        }
        totalJobs += w.jobsHandled
        totalErrors += w.errors
    }

    return map[string]any{
        "workers_total":   len(p.workers),
        "workers_active":  activeWorkers,
        "workers_idle":    len(p.workers) - activeWorkers,
        "jobs_processed":  totalJobs,
        "errors_total":    totalErrors,
        "error_rate":      float64(totalErrors) / float64(max(totalJobs, 1)),
    }
}
```

**Expose via HTTP:**
```go
http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(pool.Health())
})
```

---

## Production Architecture Example

```
┌─────────────────────────────────────────────────────────────────┐
│                        Main Process                              │
│                                                                  │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐       │
│  │   Producer   │───▶│  jobs chan   │◀───│  Supervisor  │       │
│  │ (Kafka/HTTP) │    │ (never closes│    │  (restarts   │       │
│  └──────────────┘    │  in normal   │    │   workers)   │       │
│                      │  operation)  │    └──────────────┘       │
│                      └──────────────┘           │                │
│                                                 ▼                │
│                      ┌──────────────┐    ┌──────────────┐       │
│                      │ results chan │◀───│   Workers    │       │
│                      │ (stays open) │    │ (long-lived) │       │
│                      └──────────────┘    └──────────────┘       │
│                             │                                    │
│                             ▼                                    │
│                      ┌──────────────┐                           │
│                      │   Consumer   │                           │
│                      │  (DB/API)    │                           │
│                      └──────────────┘                           │
│                                                                  │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │                    Shutdown Handler                       │   │
│  │  SIGTERM → cancel context → drain period → force exit    │   │
│  └──────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
```

---

## Signal Handling for Graceful Shutdown

```go
func main() {
    ctx, cancel := context.WithCancel(context.Background())

    // Setup signal handling
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

    // Start your pool
    pool := NewSupervisor(ctx, jobs, 10, processFunc)
    results := pool.Start()

    // Consumer goroutine
    go func() {
        for result := range results {
            saveToDatabase(result)
        }
    }()

    // Wait for shutdown signal
    <-sigCh
    log.Println("Shutdown signal received, draining...")

    // Graceful shutdown
    shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer shutdownCancel()

    if err := pool.Shutdown(); err != nil {
        log.Printf("Shutdown warning: %v", err)
    }

    log.Println("Shutdown complete")
}
```

---

## Key Takeaways

| Concept | Batch (Exercises) | Streaming (Production) |
|---------|------------------|------------------------|
| **Shutdown trigger** | Channel close | Context cancellation |
| **Worker lifetime** | Short (process batch) | Long (hours/days) |
| **Error handling** | Return error, exit | Log error, continue |
| **Panic handling** | Crashes worker | Recover, log, continue |
| **Channel closing** | Work completion signal | Cleanup after shutdown |
| **Health monitoring** | Not needed | Essential |

---

## When to Use Each Pattern

| Scenario | Pattern |
|----------|---------|
| Processing a file/batch | Batch (exercise pattern) |
| Kafka/NATS consumer | Long-running + Supervisor |
| HTTP request handler | Short-lived (per-request) |
| Scheduled job processor | Long-running + Health |
| Real-time data pipeline | Long-running + Graceful shutdown |

---

## Related Resources

- [CHANNELS_GUIDE.md](CHANNELS_GUIDE.md) - Channel fundamentals
- [CONTEXT_GUIDE.md](CONTEXT_GUIDE.md) - Context patterns
- [CONCURRENCY_DEBUGGING_GUIDE.md](CONCURRENCY_DEBUGGING_GUIDE.md) - Debugging techniques

---

**Last Updated:** 2026-01-13
**Context:** Production streaming patterns for Phase 2 Telemetry Ingestion Platform
