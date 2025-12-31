# Exercise 13: Graceful Shutdown

**Learning Goal:** Implement production-grade shutdown handling for concurrent systems

**Difficulty:** Tier 4 - Mastery
**Estimated Time:** 55-60 minutes

---

## Problem Description

Production applications must shut down gracefully. A "graceful shutdown" means:
1. Stop accepting new work
2. Complete in-flight operations (with timeout)
3. Release resources (close connections, flush buffers)
4. Exit cleanly

This requires coordination between:
- OS signals (SIGTERM, SIGINT)
- Multiple goroutines
- Timeouts for stuck operations
- Proper cleanup sequence

---

## Function Signatures

```go
// Server represents a long-running service
type Server struct {
    // Define fields
}

// NewServer creates a new server
func NewServer(handler func(request int) error) *Server

// Start begins accepting requests on the channel
func (s *Server) Start(requests <-chan int)

// Shutdown gracefully stops the server with timeout
func (s *Server) Shutdown(ctx context.Context) error

// SignalHandler sets up OS signal handling and returns shutdown channel
func SignalHandler(signals ...os.Signal) <-chan struct{}

// GracefulShutdown coordinates shutdown of multiple components
func GracefulShutdown(ctx context.Context, components ...Shutdownable) error

// Shutdownable is implemented by components that can shutdown
type Shutdownable interface {
    Shutdown(ctx context.Context) error
}

// Drainer drains a channel with timeout
func Drainer[T any](ctx context.Context, ch <-chan T, process func(T) error) error

// Coordinator manages graceful shutdown of multiple workers
type Coordinator struct {
    // Define fields
}

func NewCoordinator() *Coordinator
func (c *Coordinator) Add(name string, worker func(ctx context.Context) error)
func (c *Coordinator) Run(ctx context.Context) error
func (c *Coordinator) Shutdown(ctx context.Context) error
```

---

## Examples

### Server Shutdown
```go
server := NewServer(func(req int) error {
    time.Sleep(100 * time.Millisecond)
    return nil
})

requests := make(chan int)
server.Start(requests)

// Send some requests
requests <- 1
requests <- 2

// Graceful shutdown with 5 second timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

err := server.Shutdown(ctx)
// Completes in-flight requests, then returns
```

### Signal Handler
```go
shutdown := SignalHandler(syscall.SIGTERM, syscall.SIGINT)

// Start services...

<-shutdown // Blocks until signal received
fmt.Println("Shutting down...")

// Cleanup...
```

### Coordinator
```go
coord := NewCoordinator()

coord.Add("http-server", func(ctx context.Context) error {
    return http.Serve(ctx, ":8080", handler)
})

coord.Add("worker", func(ctx context.Context) error {
    return processJobs(ctx)
})

ctx, cancel := context.WithCancel(context.Background())
// Handle signals...
go func() {
    <-signalCh
    cancel()
}()

coord.Run(ctx) // Runs until context cancelled
```

---

## Instructions

1. Implement `Server` with Start and Shutdown methods
2. Implement `SignalHandler` using os/signal package
3. Implement `GracefulShutdown` for multiple components
4. Implement `Drainer` to drain channels with timeout
5. Implement `Coordinator` for multi-worker management
6. Handle edge cases: stuck operations, multiple signals, cleanup order
7. Run tests with `go test -v`

---

## Hints

### Basic
- `signal.Notify(ch, signals...)` registers signal handlers
- Server shutdown: stop accepting new work, wait for in-flight
- Use sync.WaitGroup to track in-flight operations
- Close done channel to signal workers to stop

### Intermediate
- Server needs: requests channel, done channel, WaitGroup, mutex for state
- Shutdown sequence: set shutdown flag → stop accepting → wait for WaitGroup → timeout
- Coordinator: map of workers, each with cancel function
- Handle shutdown timeout: return early with error if context expires

### Solution Pattern
```go
func (s *Server) Shutdown(ctx context.Context) error {
    s.mu.Lock()
    if s.shutdown {
        s.mu.Unlock()
        return nil
    }
    s.shutdown = true
    close(s.done)
    s.mu.Unlock()

    // Wait for in-flight with timeout
    done := make(chan struct{})
    go func() {
        s.wg.Wait()
        close(done)
    }()

    select {
    case <-done:
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}
```

---

## Think About

1. What's the difference between SIGTERM and SIGKILL?
2. Why is shutdown order important (e.g., HTTP server before database)?
3. How do you handle goroutines that ignore cancellation?
4. What should happen if shutdown times out?

---

## What This Teaches

- **Signal handling** - Responding to OS signals in Go
- **Graceful degradation** - Proper shutdown sequence
- **Timeout coordination** - Bounded waits for slow operations
- **Resource cleanup** - Closing connections, flushing buffers
- **Production patterns** - Real-world service lifecycle management
- **Error propagation** - Distinguishing timeout from success
