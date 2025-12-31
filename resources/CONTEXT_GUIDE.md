# Context Package Guide

A comprehensive guide to Go's `context` package for managing cancellation, timeouts, and request-scoped values.

---

## Why Context Exists

### The Problem

Imagine a web server that:
1. Receives HTTP request
2. Queries database
3. Calls external API
4. Processes results
5. Returns response

What happens if the client disconnects at step 2?

**Without context:**
- Database query continues (wasted work)
- API call still executes (wasted resources)
- Server processes abandoned data
- Response written to dead connection

**With context:**
- Each step checks: "Should I continue?"
- Cancellation propagates to all child operations
- Resources freed immediately
- Server moves to next request

```
┌────────────────────────────────────────────────────────────────┐
│  The Cascade Problem (Without Context)                         │
├────────────────────────────────────────────────────────────────┤
│                                                                │
│  Client ──► Handler ──► Database ──► API ──► Process           │
│    │                                                           │
│    ╳ disconnects                                               │
│                                                                │
│  But all operations continue running!                          │
│  ════════════════════════════════════                          │
│                                                                │
│  Handler:   Still waiting...                                   │
│  Database:  Still querying...                                  │
│  API:       Still calling...                                   │
│  Process:   Still processing...                                │
│                                                                │
│  Result: Wasted CPU, memory, network, database connections     │
│                                                                │
└────────────────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────────────────┐
│  The Solution (With Context)                                   │
├────────────────────────────────────────────────────────────────┤
│                                                                │
│  Client ──► Handler ──► Database ──► API ──► Process           │
│    │           │            │          │                       │
│    ╳ ─────────►╳───────────►╳─────────►╳  (cancellation)       │
│                                                                │
│  All operations stop immediately when client disconnects!      │
│                                                                │
└────────────────────────────────────────────────────────────────┘
```

---

## The Context Interface

```go
type Context interface {
    // Deadline returns the time when this context will be cancelled.
    // ok==false means no deadline is set.
    Deadline() (deadline time.Time, ok bool)

    // Done returns a channel that's closed when the context is cancelled.
    // Receiving from a closed channel returns immediately (zero value).
    Done() <-chan struct{}

    // Err returns nil if Done is not yet closed.
    // After Done is closed, returns why:
    //   - context.Canceled (explicit cancel)
    //   - context.DeadlineExceeded (timeout/deadline)
    Err() error

    // Value returns the value associated with key, or nil.
    Value(key any) any
}
```

### The Key Insight: Done() Channel

The `Done()` channel is the heart of context. It's a **broadcast mechanism**:

```go
// A closed channel returns immediately for ALL receivers
ch := make(chan struct{})
close(ch)

// All of these return immediately:
<-ch  // doesn't block
<-ch  // doesn't block
<-ch  // doesn't block (infinite receivers, zero memory)
```

This enables **one cancellation signal to notify many goroutines**.

---

## Creating Contexts

### Context Hierarchy

Contexts form a **tree**. Child contexts inherit from parents, but children can add:
- Deadlines (can only be sooner, never later)
- Cancel functions (independent cancellation)
- Values (key-value pairs)

```
┌────────────────────────────────────────────────────────────────┐
│  Context Tree Structure                                        │
├────────────────────────────────────────────────────────────────┤
│                                                                │
│                    context.Background()                        │
│                            │                                   │
│              ┌─────────────┼────────────┐                      │
│              ▼             ▼            ▼                      │
│         WithCancel    WithTimeout   WithValue                  │
│              │             │            │                      │
│         ┌────┴────┐        │       ┌────┴────┐                 │
│         ▼         ▼        ▼       ▼         ▼                 │
│    WithValue  WithDeadline │   WithCancel  WithTimeout         │
│                            │                                   │
│                       (child contexts)                         │
│                                                                │
│  Rule: Cancelling a parent cancels ALL descendants             │
│                                                                │
└────────────────────────────────────────────────────────────────┘
```

### Root Contexts

```go
// Background: The root context. Never cancelled, no deadline, no values.
// Use for: main(), init(), tests, top-level operations
ctx := context.Background()

// TODO: Identical to Background, but signals "I need to add context later"
// Use for: When you're not sure what context to use yet
ctx := context.TODO()
```

### Derived Contexts

#### WithCancel

```go
// Creates a context that can be manually cancelled
ctx, cancel := context.WithCancel(parentCtx)
defer cancel()  // ALWAYS defer cancel to prevent leaks

// Later, to cancel:
cancel()  // closes Done() channel, Err() returns context.Canceled
```

#### WithTimeout

```go
// Creates a context that auto-cancels after duration
ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
defer cancel()  // Still required! Release resources early if done sooner

// After 5 seconds:
// - Done() channel closes
// - Err() returns context.DeadlineExceeded
```

#### WithDeadline

```go
// Creates a context that auto-cancels at specific time
deadline := time.Now().Add(30 * time.Second)
ctx, cancel := context.WithDeadline(parentCtx, deadline)
defer cancel()
```

#### WithValue

```go
// Attaches a key-value pair to context
type contextKey string
const requestIDKey contextKey = "requestID"

ctx := context.WithValue(parentCtx, requestIDKey, "abc-123")

// Retrieve later:
if id, ok := ctx.Value(requestIDKey).(string); ok {
    fmt.Println("Request ID:", id)
}
```

---

## How Cancellation Works Under the Hood

### Internal Structure (Simplified)

```go
// Conceptual implementation (actual stdlib is more complex)
type cancelCtx struct {
    Context                        // Parent context (embedded)
    mu       sync.Mutex            // Protects fields below
    done     chan struct{}         // Closed on first cancel call
    children map[*cancelCtx]struct{} // Child contexts to cancel
    err      error                 // Set when cancelled
}
```

### The Propagation Mechanism

```
┌────────────────────────────────────────────────────────────────┐
│  Cancellation Propagation                                      │
├────────────────────────────────────────────────────────────────┤
│                                                                │
│  1. Parent cancel() called                                     │
│                                                                │
│     parentCtx                                                  │
│         │                                                      │
│         ▼ close(done)                                          │
│     ┌───────┐                                                  │
│     │ done  │ ← channel closed                                 │
│     └───────┘                                                  │
│                                                                │
│  2. Parent iterates children, calls their cancel()             │
│                                                                │
│     parentCtx ──► childCtx1.cancel()                           │
│                   childCtx2.cancel()                           │
│                   childCtx3.cancel()                           │
│                                                                │
│  3. Each child closes its done channel and cancels its children│
│                                                                │
│     childCtx1 ──► grandchildCtx1.cancel()                      │
│                   grandchildCtx2.cancel()                      │
│                                                                │
│  Result: Entire subtree cancelled in one operation             │
│                                                                │
└────────────────────────────────────────────────────────────────┘
```

### Why defer cancel() Is Required

```go
ctx, cancel := context.WithCancel(parentCtx)
// If you don't call cancel():
// - cancelCtx remains in parent's children map
// - Goroutine watching parent's Done() never exits
// - Memory leak!

defer cancel()  // Always do this
```

Even if context times out naturally, calling `cancel()` releases resources earlier.

---

## Patterns for Using Context

### Pattern 1: Check Done() in Loops

```go
func processItems(ctx context.Context, items []Item) error {
    for _, item := range items {
        // Check if we should stop
        select {
        case <-ctx.Done():
            return ctx.Err()  // Return why we stopped
        default:
            // Continue processing
        }

        process(item)
    }
    return nil
}
```

### Pattern 2: Propagate to Blocking Operations

```go
func fetchData(ctx context.Context, url string) ([]byte, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }

    // http.Client respects context cancellation
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err  // Returns error if context cancelled
    }
    defer resp.Body.Close()

    return io.ReadAll(resp.Body)
}
```

### Pattern 3: Select with Done()

```go
func slowOperation(ctx context.Context, duration time.Duration) error {
    select {
    case <-time.After(duration):
        return nil  // Operation completed
    case <-ctx.Done():
        return ctx.Err()  // Cancelled before completion
    }
}
```

### Pattern 4: Database Operations

```go
func queryUser(ctx context.Context, db *sql.DB, id int) (*User, error) {
    // Context passed to database query
    row := db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", id)

    var user User
    if err := row.Scan(&user.ID, &user.Name); err != nil {
        return nil, err  // Includes context cancellation errors
    }
    return &user, nil
}
```

---

## Context in HTTP Handlers

### The Request Context

Every `*http.Request` carries a context:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()  // Get the request's context

    // This context is cancelled when:
    // 1. Client disconnects
    // 2. Request handler returns
    // 3. Server shuts down
}
```

### Creating Derived Contexts in Handlers

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Add timeout to request context
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()

    // Pass to downstream operations
    result, err := fetchData(ctx, "https://api.example.com/data")
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            http.Error(w, "request timeout", http.StatusGatewayTimeout)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(result)
}
```

### Middleware Pattern

```go
func TimeoutMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx, cancel := context.WithTimeout(r.Context(), timeout)
            defer cancel()

            // Create new request with timeout context
            r = r.WithContext(ctx)

            // Pass to next handler
            next.ServeHTTP(w, r)
        })
    }
}
```

---

## Context Values: Use Sparingly

### When to Use Context Values

✅ **Good uses:**
- Request ID (for logging/tracing)
- Authentication info (user ID, permissions)
- Request-scoped data that crosses API boundaries

❌ **Bad uses:**
- Function parameters (pass explicitly instead)
- Optional configuration (use functional options)
- Any data that has a clear owner

### The Type-Safe Key Pattern

```go
// WRONG: String keys can collide
ctx = context.WithValue(ctx, "userID", 123)  // Another package might use "userID"!

// RIGHT: Use unexported type for key
type contextKey string

const (
    userIDKey    contextKey = "userID"
    requestIDKey contextKey = "requestID"
)

// Even better: Completely unexported type
type ctxKey struct{}
var userIDKey ctxKey

// Usage
ctx = context.WithValue(ctx, userIDKey, 123)
```

### Retrieving Values Safely

```go
func GetUserID(ctx context.Context) (int, bool) {
    id, ok := ctx.Value(userIDKey).(int)
    return id, ok
}

// Usage
if userID, ok := GetUserID(ctx); ok {
    // Use userID
}
```

---

## Common Pitfalls

### Pitfall 1: Forgetting defer cancel()

```go
// WRONG: Memory leak
ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
// ... use ctx ...
// Forgot to call cancel()!

// RIGHT
ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
defer cancel()  // Always!
```

### Pitfall 2: Storing Context in Struct

```go
// WRONG: Context should be passed explicitly
type Server struct {
    ctx context.Context  // Don't do this!
}

// RIGHT: Pass context to methods
func (s *Server) HandleRequest(ctx context.Context, req Request) error {
    // Use ctx here
}
```

### Pitfall 3: Nil Context

```go
// WRONG: Will panic
var ctx context.Context
ctx.Done()  // panic: nil pointer

// RIGHT: Use Background or TODO
ctx := context.Background()
```

### Pitfall 4: Wrong Deadline Direction

```go
parentCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)

// This does NOT extend the deadline!
childCtx, _ := context.WithTimeout(parentCtx, 10*time.Second)
// childCtx still expires when parentCtx does (5 seconds)
```

### Pitfall 5: Checking Error Before Done

```go
// WRONG order
if ctx.Err() != nil {
    return ctx.Err()
}
// Done() might close between check and here!

// RIGHT: Check via select
select {
case <-ctx.Done():
    return ctx.Err()
default:
    // Continue
}
```

---

## The Done() Channel: Deep Dive

### Why chan struct{}?

```go
Done() <-chan struct{}
```

- `struct{}` is zero-size (no memory allocation)
- Channel is for **signaling**, not passing data
- Receive-only (`<-chan`) prevents accidental sends

### The Closed Channel Trick

```go
// Once closed, receiving always succeeds immediately
ch := make(chan struct{})
close(ch)

// These never block:
<-ch        // Returns zero value immediately
<-ch        // Same
val := <-ch // val == struct{}{}

// Select case triggers immediately:
select {
case <-ch:
    fmt.Println("triggered!")  // Always runs
}
```

### Non-Blocking Check Pattern

```go
select {
case <-ctx.Done():
    // Context cancelled
    return ctx.Err()
default:
    // Context still active, continue
}
```

---

## Complete Example: HTTP Handler with Timeout

```go
package main

import (
    "context"
    "encoding/json"
    "errors"
    "net/http"
    "time"
)

// SlowOperation simulates work that respects context
func SlowOperation(ctx context.Context, duration time.Duration) error {
    select {
    case <-time.After(duration):
        return nil  // Completed successfully
    case <-ctx.Done():
        return ctx.Err()  // Cancelled
    }
}

// SlowHandler handles requests with simulated slow work
func SlowHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    // Simulate 2 seconds of work
    if err := SlowOperation(ctx, 2*time.Second); err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            http.Error(w, "operation timed out", http.StatusGatewayTimeout)
            return
        }
        if errors.Is(err, context.Canceled) {
            // Client disconnected, no point responding
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "completed"})
}

// TimeoutMiddleware wraps handlers with a timeout
func TimeoutMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx, cancel := context.WithTimeout(r.Context(), timeout)
            defer cancel()

            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

func main() {
    mux := http.NewServeMux()

    // Wrap SlowHandler with 1-second timeout
    slowHandler := TimeoutMiddleware(1*time.Second)(http.HandlerFunc(SlowHandler))
    mux.Handle("/slow", slowHandler)

    http.ListenAndServe(":8080", mux)
}
```

---

## Context Flow Visualization

```
┌────────────────────────────────────────────────────────────────┐
│  HTTP Request Context Flow                                     │
├────────────────────────────────────────────────────────────────┤
│                                                                │
│  1. Request arrives                                            │
│     ┌─────────────────────────────────────────────────┐        │
│     │  http.Server creates context.Background()       │        │
│     │  Attaches to *http.Request via WithCancel       │        │
│     └─────────────────────────────────────────────────┘        │
│                         │                                      │
│                         ▼                                      │
│  2. Middleware adds timeout                                    │
│     ┌─────────────────────────────────────────────────┐        │
│     │  ctx, cancel := WithTimeout(r.Context(), 5s)    │        │
│     │  r = r.WithContext(ctx)                         │        │
│     └─────────────────────────────────────────────────┘        │
│                         │                                      │
│                         ▼                                      │
│  3. Handler uses context                                       │
│     ┌─────────────────────────────────────────────────┐        │
│     │  ctx := r.Context()                             │        │
│     │  result, err := fetchData(ctx, url)             │        │
│     └─────────────────────────────────────────────────┘        │
│                         │                                      │
│                         ▼                                      │
│  4. Downstream propagates context                              │
│     ┌─────────────────────────────────────────────────┐        │
│     │  req := http.NewRequestWithContext(ctx, ...)    │        │
│     │  db.QueryContext(ctx, ...)                      │        │
│     └─────────────────────────────────────────────────┘        │
│                         │                                      │
│                         ▼                                      │
│  5. If timeout/cancel occurs                                   │
│     ┌─────────────────────────────────────────────────┐        │
│     │  ctx.Done() closes                              │        │
│     │  All operations checking ctx.Done() return      │        │
│     │  ctx.Err() returns reason (Canceled/Deadline)   │        │
│     └─────────────────────────────────────────────────┘        │
│                                                                │
└────────────────────────────────────────────────────────────────┘
```

---

## Quick Reference

| Function | Purpose | Returns |
|----------|---------|---------|
| `context.Background()` | Root context, never cancelled | `Context` |
| `context.TODO()` | Placeholder when unsure | `Context` |
| `context.WithCancel(parent)` | Manual cancellation | `Context, CancelFunc` |
| `context.WithTimeout(parent, d)` | Auto-cancel after duration | `Context, CancelFunc` |
| `context.WithDeadline(parent, t)` | Auto-cancel at time | `Context, CancelFunc` |
| `context.WithValue(parent, k, v)` | Attach key-value | `Context` |

| Method | Purpose |
|--------|---------|
| `ctx.Done()` | Channel closed on cancel |
| `ctx.Err()` | `nil`, `Canceled`, or `DeadlineExceeded` |
| `ctx.Deadline()` | When context expires (if set) |
| `ctx.Value(key)` | Retrieve attached value |

| Error | Meaning |
|-------|---------|
| `context.Canceled` | Explicit `cancel()` called |
| `context.DeadlineExceeded` | Timeout or deadline reached |

---

## Resource Leaks: What Happens When You Forget

### Memory Leak vs Resource Leak

Go has garbage collection, so traditional "memory leaks" (forgetting to `free()`) don't exist. What we call "leaks" in Go are actually:

1. **Resource leaks** - OS resources (file descriptors, sockets) not returned
2. **Goroutine leaks** - Goroutines blocked forever, holding memory
3. **Reference leaks** - Objects reachable from live code, so GC can't collect them

### What Happens When You Forget defer cancel()

```go
ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
// Forgot: defer cancel()
```

**Under the hood:**

```
┌────────────────────────────────────────────────────────────────┐
│  What WithTimeout Creates Internally                           │
├────────────────────────────────────────────────────────────────┤
│                                                                │
│  parentCtx                                                     │
│      │                                                         │
│      ├── children map: {&childCtx: struct{}{}}  ◄── REFERENCE  │
│      │                                                         │
│      ▼                                                         │
│  childCtx (your ctx)                                           │
│      │                                                         │
│      ├── timer goroutine watching time.After()  ◄── GOROUTINE  │
│      ├── done channel (not closed)                             │
│      └── propagation goroutine watching parent  ◄── GOROUTINE  │
│                                                                │
│  If you don't call cancel():                                   │
│  • childCtx stays in parent's children map (can't be GC'd)     │
│  • Timer goroutine runs until timeout (wasted CPU)             │
│  • Propagation goroutine blocks on parent.Done() forever       │
│                                                                │
└────────────────────────────────────────────────────────────────┘
```

**The leak:** Parent holds reference to child → child can't be garbage collected → goroutines keep running.

### Why defer cancel() Even If Timeout Works?

```go
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()  // Why? Timeout will cancel automatically!

result := doFastOperation(ctx)  // Completes in 100ms
```

**Answer:** Early cleanup!

```
┌────────────────────────────────────────────────────────────────┐
│  Without defer cancel()                                        │
├────────────────────────────────────────────────────────────────┤
│                                                                │
│  t=0ms:    Context created, timer started for 5 seconds        │
│  t=100ms:  Operation completes, function returns               │
│  t=100ms-5000ms: Timer goroutine still running! (wasted)       │
│  t=5000ms: Timer fires, context cancelled (nobody cares)       │
│                                                                │
│  For 4.9 seconds, resources held unnecessarily                 │
│                                                                │
├────────────────────────────────────────────────────────────────┤
│  With defer cancel()                                           │
├────────────────────────────────────────────────────────────────┤
│                                                                │
│  t=0ms:    Context created, timer started                      │
│  t=100ms:  Operation completes                                 │
│  t=100ms:  defer cancel() runs → timer stopped, cleanup done   │
│                                                                │
│  Resources freed immediately!                                  │
│                                                                │
└────────────────────────────────────────────────────────────────┘
```

### Other Common Resource Leaks

#### HTTP Response Body

```go
resp, err := http.Get(url)
// Forgot: defer resp.Body.Close()
```

```
┌────────────────────────────────────────────────────────────────┐
│  OS File Descriptor Table (per process)                        │
├────────────────────────────────────────────────────────────────┤
│                                                                │
│  ┌────┬────────────────────────────────────────┐               │
│  │ FD │ Resource                               │               │
│  ├────┼────────────────────────────────────────┤               │
│  │  0 │ stdin                                  │               │
│  │  1 │ stdout                                 │               │
│  │  2 │ stderr                                 │               │
│  │  3 │ TCP to api.example.com:443 (LEAKED!)   │               │
│  │  4 │ TCP to api.example.com:443 (LEAKED!)   │               │
│  │  5 │ TCP to api.example.com:443 (LEAKED!)   │               │
│  │ .. │ ...                                    │               │
│  │1024│ LIMIT REACHED - "too many open files"  │               │
│  └────┴────────────────────────────────────────┘               │
│                                                                │
│  Go's GC can't help: OS resources are outside Go's memory!     │
│                                                                │
└────────────────────────────────────────────────────────────────┘
```

#### Database Connections

```go
rows, _ := db.Query("SELECT * FROM users")
// Forgot: defer rows.Close()
```

Database connection stays "in use" → connection pool exhausted → new queries block or fail.

---

## How to Know If Something Needs Closing

### Rule 1: Check for io.Closer Interface

```go
type Closer interface {
    Close() error
}
```

If a type has a `Close()` method, you probably need to call it:

```go
// These all implement io.Closer:
*os.File
*http.Response.Body  // (io.ReadCloser)
*sql.Rows
*sql.DB
net.Conn
*gzip.Reader
```

### Rule 2: Functions Returning Cancel Functions

```go
// If a function returns a cancel/cleanup function, call it!
ctx, cancel := context.WithCancel(...)    // defer cancel()
ctx, cancel := context.WithTimeout(...)   // defer cancel()
ctx, cancel := context.WithDeadline(...)  // defer cancel()
timer := time.AfterFunc(...)              // defer timer.Stop()
ticker := time.NewTicker(...)             // defer ticker.Stop()
```

### Rule 3: Resource Acquisition Pattern

If you "open", "create", "dial", "connect", or "acquire" something:

```go
// Pattern: Open/Create/Dial → defer Close
file, err := os.Open(...)           // defer file.Close()
file, err := os.Create(...)         // defer file.Close()
conn, err := net.Dial(...)          // defer conn.Close()
listener, err := net.Listen(...)    // defer listener.Close()
resp, err := http.Get(...)          // defer resp.Body.Close()
db, err := sql.Open(...)            // defer db.Close() (at app shutdown)
```

### Decision Tree

```
┌────────────────────────────────────────────────────────────────┐
│  Does This Need Closing?                                       │
├────────────────────────────────────────────────────────────────┤
│                                                                │
│  Got value from Open/Create/Dial/Get/Connect?                  │
│      │                                                         │
│      ├── YES → Check for Close() method → defer Close()        │
│      │                                                         │
│      └── NO                                                    │
│            │                                                   │
│            ▼                                                   │
│  Function returned (value, cancelFunc)?                        │
│      │                                                         │
│      ├── YES → defer cancelFunc()                              │
│      │                                                         │
│      └── NO                                                    │
│            │                                                   │
│            ▼                                                   │
│  Type implements io.Closer?                                    │
│      │                                                         │
│      ├── YES → Probably needs Close() at some point            │
│      │                                                         │
│      └── NO → Probably safe (GC handles it)                    │
│                                                                │
└────────────────────────────────────────────────────────────────┘
```

---

## Testing for Leaks

### Method 1: Goroutine Count (Quick Check)

```go
func TestNoGoroutineLeak(t *testing.T) {
    before := runtime.NumGoroutine()

    // Run your code
    doSomething()

    // Give goroutines time to exit
    time.Sleep(100 * time.Millisecond)

    after := runtime.NumGoroutine()
    if after > before {
        t.Errorf("goroutine leak: before=%d, after=%d", before, after)
    }
}
```

### Method 2: Memory Stats (More Detailed)

```go
func TestNoMemoryLeak(t *testing.T) {
    var before, after runtime.MemStats

    runtime.GC()  // Clean slate
    runtime.ReadMemStats(&before)

    // Run your code many times
    for i := 0; i < 1000; i++ {
        doSomething()
    }

    runtime.GC()  // Force collection
    runtime.ReadMemStats(&after)

    // Check heap growth
    heapGrowth := after.HeapAlloc - before.HeapAlloc
    if heapGrowth > 1024*1024 {  // 1MB threshold
        t.Errorf("possible memory leak: heap grew by %d bytes", heapGrowth)
    }
}
```

### Method 3: Race Detector

```bash
go test -race ./...
```

Detects data races that often accompany goroutine leaks.

### Method 4: pprof (Production Debugging)

```go
import _ "net/http/pprof"

func main() {
    go http.ListenAndServe("localhost:6060", nil)
    // ... your app ...
}
```

Then inspect:
```bash
# Goroutine dump
go tool pprof http://localhost:6060/debug/pprof/goroutine

# Heap profile
go tool pprof http://localhost:6060/debug/pprof/heap
```

### Method 5: goleak (Best for Tests)

```go
import "go.uber.org/goleak"

func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m)  // Fails if any goroutines leaked
}
```

---

## The Golden Rule

```go
// If you see this pattern:
thing, cleanup := CreateThing()

// IMMEDIATELY write:
defer cleanup()

// Then continue with your code
```

This becomes muscle memory. The moment you write `ctx, cancel :=`, your fingers should automatically type `defer cancel()` on the next line.

---

## Best Practices Summary

1. **Pass context as first parameter**: `func DoSomething(ctx context.Context, ...)`
2. **Always defer cancel()**: Even if timeout will trigger anyway
3. **Don't store context in structs**: Pass explicitly through call chain
4. **Use typed keys for values**: Prevents collisions between packages
5. **Check Done() in long operations**: Loops, I/O, network calls
6. **Propagate to all blocking calls**: HTTP, database, file I/O
7. **Use Background() at top level**: main(), init(), tests
8. **Return ctx.Err() when cancelled**: Tells caller why operation stopped
9. **Don't pass nil context**: Use Background() or TODO() instead
10. **Child deadlines can't exceed parent**: Timeouts only get shorter
