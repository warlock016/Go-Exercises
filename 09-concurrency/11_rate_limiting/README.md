# Exercise 11: Rate Limiting

**Learning Goal:** Control the rate of operations using tickers and token bucket algorithms

**Difficulty:** Tier 3 - Integration
**Estimated Time:** 45-50 minutes

---

## Problem Description

**Rate limiting** controls how often an operation can occur. This is critical for:
- API calls (respecting external rate limits)
- Resource protection (preventing overload)
- Fair usage (ensuring equitable access)

Go provides `time.Ticker` for regular intervals, but real-world rate limiting often needs:
- **Token bucket** - Allows bursts up to bucket size, then rate-limits
- **Sliding window** - Limits operations in a rolling time window
- **Concurrent-safe** - Multiple goroutines sharing a limiter

---

## Function Signatures

```go
// RateLimiter limits the rate of operations
type RateLimiter struct {
    // Define fields
}

// NewRateLimiter creates a limiter that allows 'rate' operations per second
func NewRateLimiter(rate int) *RateLimiter

// Wait blocks until an operation is allowed
func (r *RateLimiter) Wait()

// TryAcquire returns true if operation is allowed (non-blocking)
func (r *RateLimiter) TryAcquire() bool

// TokenBucket implements token bucket rate limiting
type TokenBucket struct {
    // Define fields
}

// NewTokenBucket creates a bucket with capacity and refill rate
func NewTokenBucket(capacity int, refillRate time.Duration) *TokenBucket

// Take removes a token, blocking if none available
func (tb *TokenBucket) Take()

// TryTake removes a token if available (non-blocking)
func (tb *TokenBucket) TryTake() bool

// Tokens returns current token count
func (tb *TokenBucket) Tokens() int

// Throttle wraps a function with rate limiting
func Throttle[T any](fn func(T) error, limiter *RateLimiter) func(T) error

// BatchRateLimiter processes items in rate-limited batches
func BatchRateLimiter[T any](items []T, batchSize int, interval time.Duration, process func([]T) error) error
```

---

## Examples

### RateLimiter
```go
limiter := NewRateLimiter(10) // 10 per second

for i := 0; i < 30; i++ {
    limiter.Wait()
    // This will spread 30 operations over ~3 seconds
    doSomething()
}
```

### TokenBucket
```go
bucket := NewTokenBucket(5, 100*time.Millisecond)

// Can immediately do 5 operations (burst)
for i := 0; i < 5; i++ {
    bucket.Take() // Instant
    doSomething()
}

// 6th operation waits for refill
bucket.Take() // Waits ~100ms
```

### Throttle
```go
limiter := NewRateLimiter(5)
throttledCall := Throttle(apiCall, limiter)

// All calls respect rate limit
for _, item := range items {
    err := throttledCall(item)
    // ...
}
```

### BatchRateLimiter
```go
items := []string{"a", "b", "c", "d", "e", "f", "g"}

err := BatchRateLimiter(items, 2, 500*time.Millisecond, func(batch []string) error {
    return sendBatch(batch)
})
// Processes: [a,b] ... wait 500ms ... [c,d] ... wait 500ms ... [e,f] ... wait 500ms ... [g]
```

---

## Instructions

1. Implement `RateLimiter` using time.Ticker
2. Implement `Wait()` that blocks until next tick
3. Implement `TryAcquire()` for non-blocking attempt
4. Implement `TokenBucket` with capacity and refill
5. Implement `Take()` and `TryTake()` for token bucket
6. Implement `Throttle` wrapper function
7. Implement `BatchRateLimiter` for batch processing
8. Run tests with `go test -v`

---

## Hints

### Basic
- `time.NewTicker(duration)` creates a ticker that fires every duration
- `<-ticker.C` blocks until next tick
- For rate N per second, interval = time.Second / N
- Remember to stop tickers when done: `ticker.Stop()`

### Intermediate
- RateLimiter needs to handle first call (shouldn't wait initially)
- TokenBucket: track tokens and last refill time
- Use mutex to protect shared state in concurrent access
- For TryAcquire, use select with default case on ticker channel

### Solution Pattern
```go
type RateLimiter struct {
    ticker   *time.Ticker
    mu       sync.Mutex
    lastTick time.Time
}

func (r *RateLimiter) Wait() {
    <-r.ticker.C
}

func (r *RateLimiter) TryAcquire() bool {
    select {
    case <-r.ticker.C:
        return true
    default:
        return false
    }
}
```

---

## Think About

1. What's the difference between rate limiting and throttling?
2. Why would you choose token bucket over fixed interval?
3. How would you implement distributed rate limiting across multiple servers?
4. What happens if the rate limiter ticker stops but goroutines are still waiting?

---

## What This Teaches

- **time.Ticker** - Regular interval timing in Go
- **Token bucket algorithm** - Allowing bursts while limiting overall rate
- **Concurrent-safe limiting** - Protecting shared rate limit state
- **Non-blocking tries** - Using select default for immediate feedback
- **Higher-order functions** - Wrapping functions with rate limiting
- **Resource management** - Stopping tickers to prevent leaks
