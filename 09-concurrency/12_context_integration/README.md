# Exercise 12: Context Integration

**Learning Goal:** Combine context cancellation with concurrent patterns for robust production code

**Difficulty:** Tier 3 - Integration
**Estimated Time:** 50-55 minutes

---

## Problem Description

You've used context in HTTP handlers. Now it's time to integrate context deeply with concurrent patterns. Real production code needs:

- **Cancellation propagation** - When parent cancels, all children should stop
- **Timeout hierarchies** - Inner operations with tighter timeouts than outer
- **Graceful handling** - Clean shutdown when context is cancelled
- **Error context** - Knowing WHY something was cancelled

This exercise combines everything: goroutines, channels, select, and context.

---

## Function Signatures

```go
// WorkerWithContext processes items until context is cancelled
func WorkerWithContext(ctx context.Context, items <-chan int, process func(int) error) error

// FetchAll fetches from multiple sources with shared cancellation
func FetchAll(ctx context.Context, urls []string, fetch func(ctx context.Context, url string) (string, error)) (map[string]string, error)

// FetchRace returns result from first successful fetch, cancelling others
func FetchRace(ctx context.Context, urls []string, fetch func(ctx context.Context, url string) (string, error)) (string, error)

// TimeoutPipeline processes input with per-item timeout
func TimeoutPipeline[T, R any](ctx context.Context, input <-chan T, timeout time.Duration, process func(context.Context, T) (R, error)) <-chan Result[R]

type Result[T any] struct {
    Value T
    Err   error
}

// RetryWithBackoff retries operation with exponential backoff
func RetryWithBackoff(ctx context.Context, maxRetries int, initialDelay time.Duration, op func(context.Context) error) error

// Parallel runs functions concurrently, cancelling all on first error
func Parallel(ctx context.Context, fns ...func(context.Context) error) error
```

---

## Examples

### WorkerWithContext
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

items := make(chan int)
go func() {
    for i := 0; ; i++ {
        items <- i
        time.Sleep(100 * time.Millisecond)
    }
}()

err := WorkerWithContext(ctx, items, func(i int) error {
    fmt.Println(i)
    return nil
})
// Processes items for 2 seconds, then returns context.DeadlineExceeded
```

### FetchRace
```go
ctx := context.Background()
urls := []string{"http://fast.com", "http://slow.com", "http://medium.com"}

result, err := FetchRace(ctx, urls, httpFetch)
// Returns result from fastest responder
// Other fetches are cancelled
```

### RetryWithBackoff
```go
ctx := context.Background()
attempts := 0

err := RetryWithBackoff(ctx, 3, 100*time.Millisecond, func(ctx context.Context) error {
    attempts++
    if attempts < 3 {
        return errors.New("temporary failure")
    }
    return nil
})
// Retries with delays: 100ms, 200ms, 400ms
// Succeeds on 3rd attempt
```

---

## Instructions

1. Implement `WorkerWithContext` - process items respecting context cancellation
2. Implement `FetchAll` - parallel fetches with shared context
3. Implement `FetchRace` - first result wins, cancel others
4. Implement `TimeoutPipeline` - per-item timeouts in stream processing
5. Implement `RetryWithBackoff` - exponential backoff with context
6. Implement `Parallel` - run multiple functions, fail-fast with cancellation
7. Run tests with `go test -v -race`

---

## Hints

### Basic
- Always check `ctx.Done()` in select statements
- Use `context.WithCancel` to create child contexts you can cancel
- Return `ctx.Err()` to report why operation stopped
- `select { case <-ctx.Done(): ... }` for non-blocking cancel check

### Intermediate
- FetchRace: spawn goroutine per URL, first success cancels others
- TimeoutPipeline: use `context.WithTimeout` per item
- RetryWithBackoff: multiply delay by 2 each attempt
- Parallel: create child context, cancel on first error

### Solution Pattern
```go
func WorkerWithContext(ctx context.Context, items <-chan int, process func(int) error) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        case item, ok := <-items:
            if !ok {
                return nil
            }
            if err := process(item); err != nil {
                return err
            }
        }
    }
}
```

---

## Think About

1. What's the difference between context.Canceled and context.DeadlineExceeded?
2. Why cancel child contexts even if parent is already cancelled?
3. How would you add logging to see cancellation propagation?
4. What happens if process function ignores context and runs forever?

---

## What This Teaches

- **Context propagation** - Passing cancellation through call chain
- **Child contexts** - Creating local deadlines/cancellation
- **Race patterns** - First result wins with cleanup
- **Error handling** - Distinguishing cancellation from other errors
- **Retry patterns** - Backoff with context awareness
- **Production patterns** - Real-world concurrent API design
