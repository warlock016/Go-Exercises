# Module 09: Concurrency

**Focus Area:** Goroutines, Channels, Synchronization, and Production Patterns
**Prerequisites:** Module 01 (Fundamentals), Module 08 (HTTP APIs - context experience)
**Estimated Time:** 11-13 hours
**Exercises:** 15

---

## Overview

Concurrency is one of Go's defining features. This module takes you from goroutine basics to production-ready patterns like worker pools, pipelines, and graceful shutdown.

**What You'll Learn:**
- Goroutine lifecycle and the need for synchronization
- Channel-based communication between goroutines
- The `select` statement for multiplexing channels
- Mutex-based protection for shared state
- Production patterns: worker pools, fan-out/fan-in, pipelines
- Rate limiting and throttling
- Graceful shutdown with context integration

---

## Module Structure

### Tier 1: Introduction (Exercises 01-03)
Build foundational understanding of goroutines and channels.

| Exercise | Concept | Time |
|----------|---------|------|
| 01_goroutine_basics | `go` keyword, WaitGroup, race conditions | 25-30 min |
| 02_channel_fundamentals | send/receive, close, range, direction types | 30-35 min |
| 03_buffered_channels | buffered vs unbuffered, capacity, producer-consumer | 30-35 min |

### Tier 2: Application (Exercises 04-07)
Apply concurrency patterns to practical problems.

| Exercise | Concept | Time |
|----------|---------|------|
| 04_select_statement | select, time.After, default, channel merging | 35-40 min |
| 05_done_channel | cancellation pattern, broadcast via close | 35-40 min |
| 06_mutex_shared_state | sync.Mutex, RWMutex, critical sections | 35-40 min |
| 07_error_handling | error propagation, Result type, errgroup | 40-45 min |

### Tier 3: Integration (Exercises 08-12)
Combine patterns into production-ready solutions.

| Exercise | Concept | Time |
|----------|---------|------|
| 08_worker_pool | bounded concurrency, job queues | 45-50 min |
| 09_fan_out_fan_in | parallel distribution and collection | 45-50 min |
| 10_pipeline | composable processing stages | 45-50 min |
| 11_rate_limiting | ticker, token bucket, throttling | 50-55 min |
| 12_context_integration | context with concurrent patterns | 45-50 min |

### Tier 4: Mastery (Exercises 13-15)
Build production-grade concurrent systems.

| Exercise | Concept | Time |
|----------|---------|------|
| 13_graceful_shutdown | signals, draining, coordinated shutdown | 55-60 min |
| 14_semaphore_sync | semaphore, sync.Once, sync.Cond, barriers | 55-60 min |
| 15_concurrent_processor | capstone: complete processing system | 75-90 min |

---

## Key Concepts

### Goroutines
Lightweight threads managed by the Go runtime. Created with the `go` keyword:
```go
go func() {
    // runs concurrently
}()
```

### Channels
Typed conduits for communication between goroutines:
```go
ch := make(chan int)      // unbuffered
ch := make(chan int, 10)  // buffered with capacity 10
ch <- 42                  // send
val := <-ch               // receive
close(ch)                 // close (signals no more values)
```

### Select
Multiplexes channel operations:
```go
select {
case v := <-ch1:
    // received from ch1
case ch2 <- val:
    // sent to ch2
case <-time.After(timeout):
    // timeout
default:
    // non-blocking
}
```

### Synchronization Primitives
- `sync.WaitGroup` - Wait for goroutines to complete
- `sync.Mutex` - Exclusive access to shared state
- `sync.RWMutex` - Multiple readers, single writer
- `sync.Once` - One-time initialization
- `sync.Cond` - Condition variables

---

## Testing Commands

### Run All Module Tests
```bash
cd 09-concurrency
go test ./...
```

### Run Single Exercise
```bash
cd 09-concurrency/01_goroutine_basics
go test -v
```

### Detect Race Conditions
```bash
go test -race ./...
```

### Run Benchmarks
```bash
go test -bench=. -benchmem
```

---

## Resources

- `resources/CHANNELS_GUIDE.md` - Comprehensive channel patterns and examples
- `resources/CONTEXT_GUIDE.md` - Context for cancellation and timeouts

---

## Progression Path

```
01 Goroutine Basics
        │
02 Channel Fundamentals ──► 03 Buffered Channels
        │                          │
        └──────────┬───────────────┘
                   │
           04 Select Statement
                   │
     ┌─────────────┼─────────────┐
     │             │             │
05 Done        06 Mutex      07 Error
Channel        Shared       Handling
     │         State             │
     └─────────────┬─────────────┘
                   │
     ┌─────────────┼─────────────┬─────────────┐
     │             │             │             │
08 Worker      09 Fan-Out    10 Pipeline  11 Rate
   Pool           Fan-In                   Limiting
     │             │             │             │
     └─────────────┴─────────────┴─────────────┘
                         │
                 12 Context Integration
                         │
         ┌───────────────┼───────────────┐
         │               │               │
    13 Graceful     14 Semaphore    15 Capstone
      Shutdown        & Sync         Processor
```

---

## Anti-Patterns to Avoid

1. **Goroutine leaks** - Always ensure goroutines can exit
2. **Race conditions** - Use `-race` flag in tests
3. **Deadlocks** - Be careful with channel/mutex ordering
4. **Busy waiting** - Use channels or sync primitives instead
5. **Shared memory without protection** - Use mutex or channels

---

## Success Criteria

By completing this module, you should be able to:
- Spawn goroutines and coordinate their completion
- Use channels for safe communication between goroutines
- Implement common patterns: worker pool, pipeline, fan-out/fan-in
- Handle errors across goroutine boundaries
- Build systems with graceful shutdown
- Identify and fix race conditions
