package rate_limiting

import (
	"sync"
	"time"
)

// RateLimiter limits the rate of operations
type RateLimiter struct {
	// TODO(human): Define fields
	ticker *time.Ticker
	mu     sync.Mutex
	first  bool
}

// NewRateLimiter creates a limiter that allows 'rate' operations per second
func NewRateLimiter(rate int) *RateLimiter {
	// TODO(human): Implement
	return &RateLimiter{
		ticker: time.NewTicker(time.Second / time.Duration(rate)),
		first:  true,
	}
}

// Wait blocks until an operation is allowed
func (r *RateLimiter) Wait() {
	// TODO(human): Implement
	r.mu.Lock()

	if r.first {
		r.first = false
		r.mu.Unlock()
		return
	}
	r.mu.Unlock()
	<-r.ticker.C
}

// TryAcquire returns true if operation is allowed (non-blocking)
func (r *RateLimiter) TryAcquire() bool {
	// TODO(human): Implement
	r.mu.Lock()
	if r.first {
		r.first = false
		r.mu.Unlock()
		return true
	}
	r.mu.Unlock()
	select {
	case <-r.ticker.C:
		return true
	default:
		return false
	}
}

// Stop releases resources
func (r *RateLimiter) Stop() {
	// TODO(human): Implement
	r.ticker.Stop()
}

// TokenBucket implements token bucket rate limiting
type TokenBucket struct {
	// TODO(human): Define fields
	bucket chan int
	stop   chan struct{}
	ticker *time.Ticker
}

// NewTokenBucket creates a bucket with capacity and refill rate
func NewTokenBucket(capacity int, refillRate time.Duration) *TokenBucket {
	// TODO(human): Implement
	bucket := make(chan int, capacity)
	stop := make(chan struct{})
	ticker := time.NewTicker(refillRate)

	for range capacity {
		bucket <- 1
	}

	go func() {
		for {
			select {
			case <-stop:
				return
			case <-ticker.C:
				select {
				case bucket <- 1: // token added
				default: // bucket full, skip
				}
			}
		}
	}()

	return &TokenBucket{
		bucket: bucket,
		stop:   stop,
		ticker: ticker,
	}
}

func (tb *TokenBucket) Stop() {
	close(tb.stop)
	tb.ticker.Stop()
}

// Take removes a token, blocking if none available
func (tb *TokenBucket) Take() {
	// TODO(human): Implement
	<-tb.bucket
}

// TryTake removes a token if available (non-blocking)
func (tb *TokenBucket) TryTake() bool {
	// TODO(human): Implement
	select {
	case <-tb.bucket:
		return true
	default:
		return false
	}
}

// Tokens returns current token count
func (tb *TokenBucket) Tokens() int {
	// TODO(human): Implement
	return len(tb.bucket)
}

// Throttle wraps a function with rate limiting
func Throttle[T any](fn func(T) error, limiter *RateLimiter) func(T) error {
	// TODO(human): Implement
	return func(t T) error {
		limiter.Wait()
		return fn(t)
	}
}

// BatchRateLimiter processes items in rate-limited batches
func BatchRateLimiter[T any](items []T, batchSize int, interval time.Duration, process func([]T) error) error {
	// TODO(human): Implement
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for i := 0; i < len(items); i += batchSize {

		end := min(i+batchSize, len(items))
		batch := items[i:end]

		if err := process(batch); err != nil {
			return err
		}

		if end != len(items) {
			<-ticker.C
		}
	}

	return nil
}
