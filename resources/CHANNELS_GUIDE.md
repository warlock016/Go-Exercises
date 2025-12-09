# Channels in Go: A Complete Guide

A comprehensive reference for understanding and using channels effectively.

---

## What Is a Channel?

A **channel** is a typed conduit for sending and receiving values between goroutines. Think of it as a **pipe** - data flows through it, not a container you query.

```go
ch := make(chan string)    // Create a channel

ch <- "hello"              // Send to channel
msg := <-ch                // Receive from channel
```

**Key insight:** Channels are for **communication**, not storage. They synchronize goroutines by blocking until both sender and receiver are ready (unless buffered).

---

## Creating Channels

### Unbuffered Channels

```go
ch := make(chan int)       // Unbuffered - blocks until receiver ready
```

Unbuffered channels **synchronize** sender and receiver:
- Send blocks until someone receives
- Receive blocks until someone sends

### Buffered Channels

```go
ch := make(chan int, 10)   // Buffered with capacity 10
```

Buffered channels allow **async** sending up to capacity:
- Send blocks only when buffer is full
- Receive blocks only when buffer is empty

```go
ch := make(chan string, 3)

ch <- "a"  // Doesn't block (buffer has space)
ch <- "b"  // Doesn't block
ch <- "c"  // Doesn't block
ch <- "d"  // BLOCKS - buffer full, waits for receiver
```

---

## Channel Operations

### Sending and Receiving

```go
ch <- value      // Send value to channel
value := <-ch    // Receive value from channel
```

### Receiving with Closed Check

```go
value, ok := <-ch

if ok {
    // Channel is open, value is valid
} else {
    // Channel is closed and empty, value is zero value
}
```

### Closing Channels

```go
close(ch)        // Signal "no more values will be sent"
```

**Important rules:**
- Only the **sender** should close a channel
- Closing is optional - channels are garbage collected when unreferenced
- Closing signals completion, not cleanup

---

## Channel Lifecycle

```
┌─────────────────────────────────────────────────────────────┐
│  make(chan string, 3)                                       │
│         │                                                   │
│         ▼                                                   │
│  ┌─────────────┐                                           │
│  │ OPEN, EMPTY │  ← ch <- "a" succeeds                     │
│  └─────────────┘                                           │
│         │                                                   │
│         ▼                                                   │
│  ┌──────────────┐                                          │
│  │ OPEN, 1 ITEM │  ← ch <- "b" succeeds                    │
│  └──────────────┘                                          │
│         │                                                   │
│         ▼                                                   │
│  ┌─────────────────┐                                       │
│  │ OPEN, FULL (3)  │  ← ch <- "d" BLOCKS (waits)           │
│  └─────────────────┘                                       │
│         │                                                   │
│    <-ch (consume)                                           │
│         │                                                   │
│         ▼                                                   │
│  ┌──────────────┐                                          │
│  │ OPEN, 2 ITEMS│  ← more operations...                    │
│  └──────────────┘                                          │
│         │                                                   │
│    close(ch)                                                │
│         │                                                   │
│         ▼                                                   │
│  ┌────────────────┐                                        │
│  │ CLOSED, 2 ITEMS│  ← can still receive remaining items   │
│  └────────────────┘                                        │
│         │                                                   │
│    <-ch, <-ch (drain)                                       │
│         │                                                   │
│         ▼                                                   │
│  ┌───────────────┐                                         │
│  │ CLOSED, EMPTY │  ← receives return ("", false)          │
│  └───────────────┘                                         │
└─────────────────────────────────────────────────────────────┘
```

---

## Operation Behavior Table

| Operation | Open + Empty | Open + Has Data | Closed + Empty | Closed + Has Data |
|-----------|--------------|-----------------|----------------|-------------------|
| `ch <- x` (send) | Blocks¹ | Succeeds/Blocks¹ | **PANIC** | **PANIC** |
| `<-ch` (receive) | Blocks | Returns value | Returns zero, false | Returns value, true |
| `close(ch)` | OK | OK | **PANIC** | OK |
| `len(ch)` | 0 | Count | 0 | Count |
| `cap(ch)` | Capacity | Capacity | Capacity | Capacity |

¹ For unbuffered channels, send always blocks until receiver ready. For buffered, blocks when full.

---

## Channel Direction Types

Go allows restricting channel direction in function signatures:

```go
chan T       // Bidirectional - can send and receive
chan<- T     // Send-only - can only send
<-chan T     // Receive-only - can only receive
```

### Why Use Directional Channels?

**Type safety** - prevents accidental misuse:

```go
func producer(out chan<- string) {
    out <- "data"    // OK
    // <-out         // Compile error! Can't receive from send-only
}

func consumer(in <-chan string) {
    msg := <-in      // OK
    // in <- "data"  // Compile error! Can't send to receive-only
}

func main() {
    ch := make(chan string)

    go producer(ch)  // Implicitly converts to chan<-
    go consumer(ch)  // Implicitly converts to <-chan
}
```

### The EventBus Pattern

```go
func EventBus(bufferSize int) (
    publish func(string),           // Sends to channel
    subscribe func() <-chan string, // Returns receive-only view
    closeBus func(),
) {
    ch := make(chan string, bufferSize)

    subscribe = func() <-chan string {
        return ch  // Callers can only read, not write
    }
    // ...
}
```

---

## Common Patterns

### Pattern 1: Range Over Channel

Automatically stops when channel closes:

```go
ch := make(chan int)

go func() {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)  // Signal completion
}()

for value := range ch {
    fmt.Println(value)  // Prints 0, 1, 2, 3, 4
}
// Loop exits automatically when channel closes
```

### Pattern 2: Select Statement

Wait on multiple channels:

```go
select {
case msg := <-ch1:
    fmt.Println("from ch1:", msg)
case msg := <-ch2:
    fmt.Println("from ch2:", msg)
case <-time.After(1 * time.Second):
    fmt.Println("timeout!")
}
```

### Pattern 3: Non-Blocking Operations

```go
select {
case msg := <-ch:
    fmt.Println("received:", msg)
default:
    fmt.Println("no message available")
}
```

### Pattern 4: Done Channel (Cancellation)

```go
func worker(done <-chan struct{}, jobs <-chan int) {
    for {
        select {
        case <-done:
            return  // Exit when done is closed
        case job := <-jobs:
            process(job)
        }
    }
}

// Usage
done := make(chan struct{})
go worker(done, jobs)

// Later, to stop the worker:
close(done)
```

---

## Publish-Subscribe Pattern

The pub-sub pattern decouples message producers from consumers.

### Simple Implementation (Your EventBus)

```go
func EventBus(bufferSize int) (
    publish func(string),
    subscribe func() <-chan string,
    closeBus func(),
) {
    ch := make(chan string, bufferSize)
    var closed bool

    publish = func(s string) {
        if !closed {
            ch <- s
        }
    }

    subscribe = func() <-chan string {
        return ch
    }

    closeBus = func() {
        closed = true
        close(ch)
    }

    return publish, subscribe, closeBus
}
```

### Usage

```go
pub, sub, close := EventBus(100)

// Subscriber (could be in another goroutine)
ch := sub()
go func() {
    for msg := range ch {
        fmt.Println("Received:", msg)
    }
    fmt.Println("Channel closed")
}()

// Publisher
pub("hello")
pub("world")

// When done
close()
```

### Multi-Subscriber Pattern

Each subscriber gets their own channel:

```go
func EventBus() (
    publish func(string),
    subscribe func() <-chan string,
    closeBus func(),
) {
    var subscribers []chan string
    var closed bool
    var mu sync.Mutex

    publish = func(msg string) {
        mu.Lock()
        defer mu.Unlock()

        if closed {
            return
        }

        // Fan-out: send to all subscribers
        for _, ch := range subscribers {
            select {
            case ch <- msg:
            default:
                // Skip slow subscribers (non-blocking)
            }
        }
    }

    subscribe = func() <-chan string {
        mu.Lock()
        defer mu.Unlock()

        ch := make(chan string, 100)
        subscribers = append(subscribers, ch)
        return ch
    }

    closeBus = func() {
        mu.Lock()
        defer mu.Unlock()

        closed = true
        for _, ch := range subscribers {
            close(ch)
        }
    }

    return publish, subscribe, closeBus
}
```

---

## Channels + Closures

Channels are a natural fit with closures because:
1. Channels can be **captured state** shared across closures
2. Closures provide clean APIs for channel operations
3. Directional types in return values enforce correct usage

### The EventBus Structure

```go
func EventBus(bufferSize int) (...) {
    //  ┌─────────────────────────────────────────┐
    //  │  Shared State (captured by all closures) │
    //  ├─────────────────────────────────────────┤
    //  │  ch := make(chan string, bufferSize)    │
    //  │  closed := false                         │
    //  └─────────────────────────────────────────┘
    //           │           │           │
    //           ▼           ▼           ▼
    //      ┌────────┐  ┌─────────┐  ┌────────┐
    //      │publish │  │subscribe│  │closeBus│
    //      │        │  │         │  │        │
    //      │ch <- s │  │return ch│  │close(ch│
    //      └────────┘  └─────────┘  └────────┘
}
```

---

## Common Mistakes

### 1. Sending to Closed Channel (PANIC)

```go
ch := make(chan int)
close(ch)
ch <- 1  // PANIC: send on closed channel
```

**Fix:** Track closed state with a boolean flag:

```go
var closed bool

send := func(x int) {
    if !closed {
        ch <- x
    }
}
```

### 2. Closing Channel Multiple Times (PANIC)

```go
ch := make(chan int)
close(ch)
close(ch)  // PANIC: close of closed channel
```

**Fix:** Use `sync.Once` or boolean flag:

```go
var once sync.Once

closeBus := func() {
    once.Do(func() {
        close(ch)
    })
}
```

### 3. Forgetting to Close (Goroutine Leak)

```go
func process() <-chan int {
    ch := make(chan int)

    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
        // Forgot to close(ch)!
    }()

    return ch
}

// Consumer using range will block forever
for v := range process() {
    fmt.Println(v)
}
// Never reaches here!
```

### 4. Blocking Forever (Deadlock)

```go
ch := make(chan int)  // Unbuffered!

ch <- 1       // Blocks forever - no receiver
msg := <-ch   // Never reached

// Results in: fatal error: all goroutines are asleep - deadlock!
```

**Fix:** Use buffered channel or goroutine:

```go
ch := make(chan int, 1)  // Buffered
ch <- 1                   // OK - has space
msg := <-ch               // OK
```

### 5. Reading from Nil Channel (Blocks Forever)

```go
var ch chan int  // nil!

<-ch  // Blocks forever (not panic, just hangs)
```

---

## Channel vs Other Synchronization

| Mechanism | Use When |
|-----------|----------|
| **Channel** | Passing data between goroutines, signaling events |
| **sync.Mutex** | Protecting shared memory access |
| **sync.WaitGroup** | Waiting for multiple goroutines to complete |
| **sync.Once** | One-time initialization |
| **context.Context** | Cancellation, timeouts, request-scoped values |

### Go Proverb

> "Don't communicate by sharing memory; share memory by communicating."

Channels embody this philosophy - instead of goroutines accessing shared data with locks, they pass data through channels.

---

## Real-World Use Cases

### 1. Worker Pool

```go
func workerPool(numWorkers int, jobs <-chan Job) <-chan Result {
    results := make(chan Result)

    var wg sync.WaitGroup
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs {
                results <- process(job)
            }
        }()
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    return results
}
```

### 2. Rate Limiter

```go
func rateLimiter(rps int) <-chan struct{} {
    ch := make(chan struct{})

    go func() {
        ticker := time.NewTicker(time.Second / time.Duration(rps))
        defer ticker.Stop()

        for range ticker.C {
            ch <- struct{}{}
        }
    }()

    return ch
}

// Usage
limiter := rateLimiter(10)  // 10 requests per second

for request := range requests {
    <-limiter  // Wait for permission
    handle(request)
}
```

### 3. Timeout Pattern

```go
func fetchWithTimeout(url string, timeout time.Duration) (string, error) {
    result := make(chan string, 1)
    errCh := make(chan error, 1)

    go func() {
        data, err := fetch(url)
        if err != nil {
            errCh <- err
            return
        }
        result <- data
    }()

    select {
    case data := <-result:
        return data, nil
    case err := <-errCh:
        return "", err
    case <-time.After(timeout):
        return "", errors.New("timeout")
    }
}
```

---

## Summary

| Concept | Key Point |
|---------|-----------|
| Purpose | Communication between goroutines |
| Unbuffered | Synchronizes sender and receiver |
| Buffered | Allows async sending up to capacity |
| Close | Signals "no more sends", safe to receive remaining |
| Panic causes | Send to closed, close twice, close nil |
| `<-chan T` | Receive-only (for consumers) |
| `chan<- T` | Send-only (for producers) |
| Range | Iterates until channel closed |
| Select | Wait on multiple channels |

---

## Resources

- [Go Blog - Pipelines and Cancellation](https://go.dev/blog/pipelines)
- [Go Blog - Share Memory by Communicating](https://go.dev/blog/codelab-share)
- [Effective Go - Channels](https://go.dev/doc/effective_go#channels)
- [Go by Example - Channels](https://gobyexample.com/channels)
