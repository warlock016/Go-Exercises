# Concurrency Debugging Guide

A practical guide to debugging goroutines, channels, and concurrent Go code.

---

## Table of Contents

1. [Why Concurrent Code is Hard to Debug](#why-concurrent-code-is-hard-to-debug)
2. [Common Concurrency Bugs](#common-concurrency-bugs)
3. [Debugging Tools](#debugging-tools)
4. [Debugging Patterns](#debugging-patterns)
5. [Worked Example](#worked-example)
6. [Quick Reference](#quick-reference)

---

## Why Concurrent Code is Hard to Debug

### The Heisenberg Problem

```go
// Adding print statements CHANGES the timing!
for i := range 3 {
    go func() {
        fmt.Println(i)  // This synchronization changes behavior
    }()
}
```

**Why print debugging fails:**
- `fmt.Printf` has internal locks that serialize output
- I/O operations change goroutine scheduling
- The bug might disappear when you add prints (and reappear when you remove them!)

### Non-Determinism

The same code can:
- Work 99 times, fail on the 100th
- Pass on your machine, fail in CI
- Work with small inputs, deadlock with large ones

### Silent Failures

| Operation | On nil channel | On closed channel |
|-----------|---------------|-------------------|
| Send `ch <- v` | **Blocks forever** | **Panic** |
| Receive `<-ch` | **Blocks forever** | Returns zero value |
| Close `close(ch)` | **Panic** | **Panic** |

Nil channels block silently—no error, no panic, just hangs.

---

## Common Concurrency Bugs

### 1. Nil Channel (Silent Deadlock)

```go
// BUG: channels slice contains nil channels
channels := make([]chan int, 5)  // All nil!
channels[0] <- 42                 // Blocks forever, no panic

// FIX: Initialize each channel
for i := range channels {
    channels[i] = make(chan int)
}
```

**Symptoms:** Program hangs, no error message
**Detection:** Delve `goroutines` command shows goroutine blocked on nil channel

### 2. Forgotten Close (Receiver Blocks Forever)

```go
// BUG: Channel never closed
func generate(n int) <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; i < n; i++ {
            ch <- i
        }
        // Missing: close(ch)
    }()
    return ch
}

// Consumer blocks forever after receiving all values
for v := range generate(5) {  // Never terminates!
    fmt.Println(v)
}
```

**Symptoms:** `for range` loop never exits
**Detection:** Count goroutines before/after; receiver goroutine count doesn't decrease

### 3. Double Close (Panic)

```go
// BUG: Closing already-closed channel
close(ch)
close(ch)  // panic: close of closed channel
```

**Symptoms:** Runtime panic
**Prevention:** Only one goroutine should own closing; use `sync.Once` if needed

### 4. Send on Closed Channel (Panic)

```go
// BUG: Sending after close
close(ch)
ch <- 42  // panic: send on closed channel
```

**Symptoms:** Runtime panic
**Prevention:** Sender should close; coordinate with done channels or WaitGroups

### 5. Loop Variable Capture (Race Condition)

```go
// BUG: All goroutines share same variable
for _, v := range values {
    go func() {
        result <- v  // v changes before goroutine runs!
    }()
}

// FIX: Pass as parameter
for _, v := range values {
    go func(val int) {
        result <- val
    }(v)
}
```

**Symptoms:** Wrong values, non-deterministic results
**Detection:** `go test -race` catches this

### 6. Unbuffered Channel Deadlock

```go
// BUG: Send blocks because no receiver yet
ch := make(chan int)
ch <- 42           // Blocks here!
result := <-ch     // Never reached

// FIX: Use goroutine or buffered channel
ch := make(chan int)
go func() { ch <- 42 }()
result := <-ch
```

**Symptoms:** Program hangs on send
**Detection:** Delve shows goroutine blocked on channel send

### 7. Wrong Channel Count in Relay/Pipeline

```go
// BUG: Using input size instead of stage count
func Pipeline(n int) int {
    comm := make([]chan int, 5)
    for i := range n {  // Wrong! Should be range 5 or range comm
        comm[i] = make(chan int)
    }
    // When n < 5, some channels remain nil
}
```

**Symptoms:** Works for some inputs, deadlocks for others
**Detection:** Test with various input sizes

---

## Debugging Tools

### 1. Race Detector (`-race`)

**What it detects:** Data races (concurrent read/write without synchronization)

```bash
# Run tests with race detection
go test -race ./...

# Run program with race detection
go run -race main.go
```

**Output example:**
```
WARNING: DATA RACE
Write at 0x00c000014088 by goroutine 7:
  main.main.func1()
      /path/to/file.go:15 +0x38

Previous read at 0x00c000014088 by goroutine 6:
  main.main.func1()
      /path/to/file.go:15 +0x38
```

**Limitations:**
- Only detects races that actually occur during execution
- Slows down execution ~10x
- Doesn't detect deadlocks

### 2. Delve Debugger (`dlv`)

**Installation:**
```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

**Basic usage:**
```bash
# Debug a test
dlv test ./path/to/package -- -test.run TestName

# Debug a program
dlv debug ./main.go
```

**Key commands for concurrency:**

| Command | Description |
|---------|-------------|
| `goroutines` | List all goroutines with their state |
| `goroutine <id>` | Switch to specific goroutine |
| `goroutine <id> bt` | Show backtrace for goroutine |
| `on <breakpoint> goroutine` | Show goroutine info at breakpoint |

**Example session:**
```
(dlv) break Pipeline
(dlv) continue
(dlv) goroutines
  Goroutine 1 - User: ./main.go:42 main.main (0x10a3b40) [running]
  Goroutine 6 - User: ./pipeline.go:15 main.generate.func1 (0x10a4020) [chan send]
  Goroutine 7 - User: ./pipeline.go:25 main.filter.func1 (0x10a4120) [chan receive (nil chan)]
                                                          ^^^^^^^^^^^^^^^^^^^^^^^^
                                                          This tells you it's blocked on nil channel!
```

### 3. Execution Tracer (`go tool trace`)

**Capture trace:**
```bash
go test -trace=trace.out ./...
go tool trace trace.out
```

**What it shows:**
- Goroutine creation/destruction timeline
- Blocking events (channel ops, syscalls, GC)
- Which goroutine ran on which processor
- Scheduling latency

**Best for:** Understanding timing, finding bottlenecks, visualizing goroutine lifecycle

### 4. Runtime Diagnostics

**Goroutine count:**
```go
import "runtime"

fmt.Println("Goroutines:", runtime.NumGoroutine())
```

**Stack traces of all goroutines:**
```go
import "runtime/debug"

debug.PrintStack()  // Current goroutine only

// All goroutines:
buf := make([]byte, 1<<16)
n := runtime.Stack(buf, true)  // true = all goroutines
fmt.Printf("%s", buf[:n])
```

**Detecting goroutine leaks:**
```go
func TestNoLeak(t *testing.T) {
    before := runtime.NumGoroutine()

    // ... run your concurrent code ...

    // Give goroutines time to finish
    time.Sleep(100 * time.Millisecond)

    after := runtime.NumGoroutine()
    if after > before {
        t.Errorf("Goroutine leak: %d before, %d after", before, after)
    }
}
```

---

## Debugging Patterns

### Pattern 1: Timeout Wrapper

Wrap blocking operations to detect hangs:

```go
func debugReceive[T any](ch <-chan T, name string, timeout time.Duration) (T, bool) {
    select {
    case v, ok := <-ch:
        return v, ok
    case <-time.After(timeout):
        var zero T
        log.Fatalf("TIMEOUT: receive from %s blocked for %v", name, timeout)
        return zero, false
    }
}

func debugSend[T any](ch chan<- T, v T, name string, timeout time.Duration) {
    select {
    case ch <- v:
        return
    case <-time.After(timeout):
        log.Fatalf("TIMEOUT: send to %s blocked for %v", name, timeout)
    }
}
```

**Usage:**
```go
// Instead of: val := <-comm[0]
val, _ := debugReceive(comm[0], "comm[0]", 2*time.Second)

// Instead of: comm[1] <- val
debugSend(comm[1], val, "comm[1]", 2*time.Second)
```

### Pattern 2: Channel Inspector

Wrap channels with logging:

```go
type debugChan[T any] struct {
    ch   chan T
    name string
}

func newDebugChan[T any](name string) *debugChan[T] {
    return &debugChan[T]{
        ch:   make(chan T),
        name: name,
    }
}

func (d *debugChan[T]) Send(v T) {
    log.Printf("[%s] sending: %v", d.name, v)
    d.ch <- v
    log.Printf("[%s] sent: %v", d.name, v)
}

func (d *debugChan[T]) Receive() (T, bool) {
    log.Printf("[%s] waiting to receive", d.name)
    v, ok := <-d.ch
    log.Printf("[%s] received: %v (ok=%v)", d.name, v, ok)
    return v, ok
}
```

### Pattern 3: Goroutine Counter

Track goroutine lifecycle:

```go
var goroutineCount atomic.Int32

func trackedGo(name string, f func()) {
    goroutineCount.Add(1)
    log.Printf("Starting goroutine %s (total: %d)", name, goroutineCount.Load())
    go func() {
        defer func() {
            goroutineCount.Add(-1)
            log.Printf("Ending goroutine %s (total: %d)", name, goroutineCount.Load())
        }()
        f()
    }()
}

// Usage
trackedGo("generator", func() {
    for i := 0; i < n; i++ {
        ch <- i
    }
    close(ch)
})
```

### Pattern 4: State Machine Logging

For pipelines, log stage transitions:

```go
func logStage(stage, action string, value any) {
    log.Printf("[%-12s] %-8s %v", stage, action, value)
}

// Usage in pipeline stages
go func() {
    for val := range input {
        logStage("FilterEven", "received", val)
        if val%2 == 0 {
            logStage("FilterEven", "passing", val)
            output <- val
        } else {
            logStage("FilterEven", "dropping", val)
        }
    }
    logStage("FilterEven", "closing", nil)
    close(output)
}()
```

---

## Worked Example

### The Bug

Here's a buggy concurrent sum function:

```go
// BuggySum attempts to sum values concurrently
func BuggySum(values []int) int {
    if len(values) == 0 {
        return 0
    }

    results := make([]chan int, len(values))

    // Launch workers
    for i, v := range values {
        go func(idx int) {
            results[idx] <- v * 2  // Double each value
        }(i)
    }

    // Collect results
    sum := 0
    for i := range values {
        sum += <-results[i]
    }
    return sum
}
```

**Can you spot the bugs?** (There are 3!)

### Debugging Session

**Step 1: Run with race detector**
```bash
go test -race -run TestBuggySum
```

Output:
```
WARNING: DATA RACE
Write at 0x... by goroutine 8:
  main.BuggySum.func1()
Read at 0x... by goroutine 7:
  main.BuggySum.func1()
```

**Bug #1 found:** Loop variable `v` captured by closure

**Step 2: Run test (likely hangs)**
```bash
go test -timeout 5s -run TestBuggySum
```

Output:
```
panic: test timed out after 5s
```

**Step 3: Use Delve**
```
(dlv) goroutines
  Goroutine 6 - User: ./buggy.go:12 main.BuggySum.func1 (0x...) [chan send (nil chan)]
```

**Bug #2 found:** Sending to nil channel (channels not initialized)

**Step 4: After fixing initialization, still wrong results**

Add debug logging or check values—discover that `v` is always the last value.

**Bug #3 confirmed:** Loop variable capture

### The Fix

```go
func FixedSum(values []int) int {
    if len(values) == 0 {
        return 0
    }

    results := make([]chan int, len(values))

    // FIX #1: Initialize each channel
    for i := range results {
        results[i] = make(chan int)
    }

    // Launch workers
    for i, v := range values {
        // FIX #2: Pass v as parameter to avoid closure capture
        go func(idx int, val int) {
            results[idx] <- val * 2
        }(i, v)
    }

    // Collect results
    sum := 0
    for i := range values {
        sum += <-results[i]
    }
    return sum
}
```

---

## Quick Reference

### Debugging Decision Tree

```
Program hangs?
├── Use `dlv` → `goroutines` command
│   ├── "chan send (nil chan)" → Nil channel, initialize it
│   ├── "chan send" → Receiver missing or blocked
│   ├── "chan receive" → Sender missing, not closing, or blocked
│   └── "select (no cases)" → Empty select{} (intentional block?)
│
├── Check goroutine count with runtime.NumGoroutine()
│   └── Count keeps growing → Goroutine leak
│
└── Add timeout wrappers to find which operation blocks

Wrong results?
├── Run `go test -race`
│   └── Data race found → Fix synchronization
│
├── Add logging to trace values through pipeline
│   └── Values wrong → Check loop variable capture
│
└── Test with various inputs
    └── Some work, some fail → Check edge cases (0, 1, nil)

Panic?
├── "send on closed channel" → Sender doesn't know channel closed
├── "close of closed channel" → Multiple closers, use sync.Once
└── "close of nil channel" → Channel not initialized
```

### Command Cheat Sheet

```bash
# Race detection
go test -race ./...
go run -race main.go

# Delve debugging
dlv test ./pkg -- -test.run TestName
dlv debug ./main.go

# Inside Delve
goroutines                 # List all goroutines
goroutine 5                # Switch to goroutine 5
goroutine 5 bt             # Backtrace for goroutine 5
break pkg.Function         # Set breakpoint
continue                   # Run until breakpoint
next                       # Step over
step                       # Step into

# Execution trace
go test -trace=trace.out ./...
go tool trace trace.out
```

### Mental Checklist

Before running concurrent code:

- [ ] All channels in slices initialized?
- [ ] Loop variables passed as parameters to goroutines?
- [ ] Each channel has exactly one closer?
- [ ] Sender closes, not receiver?
- [ ] Number of sends matches number of receives (or channel closed)?
- [ ] Edge cases handled (n=0, empty slice)?

---

## Related Resources

- [CHANNELS_GUIDE.md](CHANNELS_GUIDE.md) - Channel fundamentals
- [MAKE_AND_NESTED_STRUCTURES.md](MAKE_AND_NESTED_STRUCTURES.md) - Nil channel behavior
- [DELVE_DEBUGGER_GUIDE.md](DELVE_DEBUGGER_GUIDE.md) - Delve basics

---

**Last Updated:** January 5, 2026
**Module:** 09 Concurrency - Debugging
**Context:** Debugging goroutines, channels, deadlocks, and race conditions
