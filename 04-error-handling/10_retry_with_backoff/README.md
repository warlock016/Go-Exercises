# Exercise 10: Retry with Exponential Backoff

## 🎯 Learning Goal
Implement retry logic with exponential backoff for handling transient errors in distributed systems.

## 📝 Problem Description

Network operations can fail temporarily. This exercise implements retry logic that waits progressively longer between attempts.

## 🔧 Function Signatures

```go
// RetryConfig holds retry configuration
type RetryConfig struct {
    MaxAttempts int
    InitialDelay time.Duration
    MaxDelay     time.Duration
}

// RetryWithBackoff retries a function with exponential backoff
func RetryWithBackoff(fn func() error, config RetryConfig) error

// IsRetryable determines if an error should be retried
func IsRetryable(err error) bool

// CalculateBackoff calculates the delay for the next retry
func CalculateBackoff(attempt int, initialDelay time.Duration) time.Duration
```

## 🎓 What This Teaches

- **Exponential backoff** - delay *= 2 pattern
- **Retry strategies** - When and how to retry
- **Transient vs permanent errors** - Error classification

---

**Next Exercise:** `11_error_logging` - Structured error logging
