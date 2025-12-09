# Module: Retry

## Purpose

Implement retry logic with exponential backoff for handling transient API failures. This module provides a generic retry mechanism that can be used by any HTTP client, automatically retrying failed requests that might succeed on a subsequent attempt.

---

## Package Location

```
client/retry.go
```

---

## Dependencies

| Package | Purpose |
|---------|---------|
| `errors` (internal) | Sentinel errors for classification |

---

## Public API

### Types

```go
// RetryConfig holds retry behavior configuration
type RetryConfig struct {
    MaxAttempts  int           // Maximum number of attempts (including first)
    InitialDelay time.Duration // Delay before first retry
    MaxDelay     time.Duration // Maximum delay between retries
    Multiplier   float64       // Delay multiplier for exponential backoff
}
```

### Functions

```go
// DefaultRetryConfig returns sensible defaults for HTTP API calls
func DefaultRetryConfig() RetryConfig

// RetryWithBackoff executes fn with retries for transient failures
// Returns the error from the last attempt if all retries fail
func RetryWithBackoff(fn func() error, config RetryConfig) error

// IsRetryable determines if an error should trigger a retry
func IsRetryable(err error) bool
```

---

## Behavior Specification

### DefaultRetryConfig()

**Input:** None

**Output:** `RetryConfig`

**Behavior:**
Return configuration with:
- `MaxAttempts: 3`
- `InitialDelay: 500 * time.Millisecond`
- `MaxDelay: 10 * time.Second`
- `Multiplier: 2.0`

---

### RetryWithBackoff(fn func() error, config RetryConfig)

**Input:**
- `fn` - Function to execute (returns error or nil)
- `config` - Retry configuration

**Output:** `error` - nil if any attempt succeeds, or last error after all attempts exhausted

**Behavior:**

```
RetryWithBackoff(fn, config)
   │
   ├──► for attempt := 1; attempt <= MaxAttempts; attempt++ {
   │         │
   │         ├──► err := fn()
   │         │
   │         ├──► if err == nil {
   │         │         return nil  // Success!
   │         │    }
   │         │
   │         ├──► if !IsRetryable(err) {
   │         │         return err  // Non-retryable, fail immediately
   │         │    }
   │         │
   │         ├──► if attempt == MaxAttempts {
   │         │         return err  // Last attempt, give up
   │         │    }
   │         │
   │         ├──► delay := CalculateDelay(attempt, config)
   │         │
   │         └──► time.Sleep(delay)
   │    }
   │
   └──► return lastErr
```

### Delay Calculation

```
delay = InitialDelay * (Multiplier ^ (attempt - 1))
delay = min(delay, MaxDelay)
```

**Example with defaults:**
| Attempt | Delay Calculation | Actual Delay |
|---------|-------------------|--------------|
| 1 | First attempt | 0 (no delay) |
| 2 | 500ms × 2^0 = 500ms | 500ms |
| 3 | 500ms × 2^1 = 1000ms | 1s |
| 4 | 500ms × 2^2 = 2000ms | 2s |
| 5 | 500ms × 2^3 = 4000ms | 4s |
| 6 | 500ms × 2^4 = 8000ms | 8s |
| 7 | 500ms × 2^5 = 16000ms | 10s (capped) |

---

### IsRetryable(err error)

**Input:** `error` - Any error

**Output:** `bool` - true if error is transient and worth retrying

**Behavior:**

Classify errors as retryable or not:

| Error | Retryable? | Reason |
|-------|------------|--------|
| `ErrTimeout` | Yes | Network might recover |
| `ErrRateLimited` | Yes | Rate limit will reset |
| `ErrNetwork` | Yes | Connection might be restored |
| `ErrNotFound` | No | Resource doesn't exist |
| `ErrUnauthorized` | No | Credentials are wrong |
| `ErrInvalidInput` | No | Request is malformed |
| HTTP 5xx | Yes | Server might recover |
| HTTP 4xx (except 429) | No | Client error |
| `context.Canceled` | No | User canceled |
| `context.DeadlineExceeded` | Yes | Timeout, might work later |

**Implementation:**
```go
func IsRetryable(err error) bool {
    if err == nil {
        return false
    }

    // Check sentinel errors
    if errors.Is(err, ErrTimeout) ||
       errors.Is(err, ErrRateLimited) ||
       errors.Is(err, ErrNetwork) ||
       errors.Is(err, context.DeadlineExceeded) {
        return true
    }

    // Check custom error types for status code
    var weatherErr *WeatherAPIError
    if errors.As(err, &weatherErr) {
        return weatherErr.StatusCode >= 500 || weatherErr.StatusCode == 429
    }

    var geocodeErr *GeocodeError
    if errors.As(err, &geocodeErr) {
        return geocodeErr.StatusCode >= 500 || geocodeErr.StatusCode == 429
    }

    // Non-retryable by default
    return false
}
```

---

## Edge Cases

| Scenario | Behavior |
|----------|----------|
| `fn` succeeds on first try | Return nil immediately, no retries |
| `fn` fails with non-retryable error | Return error immediately, no retries |
| `fn` fails 2 times, succeeds on 3rd | Return nil after 2 retries |
| `fn` fails all attempts | Return last error |
| `MaxAttempts = 1` | No retries, just one attempt |
| `MaxAttempts = 0` | Return nil (no attempts made) |
| Delay exceeds `MaxDelay` | Cap at `MaxDelay` |

---

## Useful Packages

| Package | Why Useful |
|---------|------------|
| `time` | `time.Sleep()`, `time.Duration` |
| `errors` | `errors.Is()`, `errors.As()` for classification |
| `math` | `math.Pow()` for exponential calculation |
| `context` | `context.DeadlineExceeded` check |

---

## Test Cases

| Scenario | Setup | Expected |
|----------|-------|----------|
| Success first try | fn returns nil | nil, fn called once |
| Success after retry | fn fails once, then succeeds | nil, fn called twice |
| Non-retryable error | fn returns ErrNotFound | ErrNotFound, fn called once |
| Retryable exhausted | fn always returns ErrTimeout | ErrTimeout after MaxAttempts |
| Delay calculation | Check delay increases | 500ms, 1s, 2s, 4s... |
| MaxDelay cap | Many retries | Delay never exceeds MaxDelay |
| Rate limited | fn returns ErrRateLimited | Retried |
| HTTP 500 error | WeatherAPIError with 500 | Retried |
| HTTP 404 error | WeatherAPIError with 404 | Not retried |

### Testing Retry Logic

```go
func TestRetryWithBackoff_SuccessAfterRetry(t *testing.T) {
    attempts := 0
    fn := func() error {
        attempts++
        if attempts < 3 {
            return errors.ErrTimeout // Retryable
        }
        return nil // Success on 3rd attempt
    }

    err := RetryWithBackoff(fn, RetryConfig{
        MaxAttempts:  5,
        InitialDelay: 1 * time.Millisecond, // Fast for tests
        MaxDelay:     10 * time.Millisecond,
        Multiplier:   2.0,
    })

    if err != nil {
        t.Errorf("expected success, got %v", err)
    }
    if attempts != 3 {
        t.Errorf("expected 3 attempts, got %d", attempts)
    }
}

func TestRetryWithBackoff_NonRetryable(t *testing.T) {
    attempts := 0
    fn := func() error {
        attempts++
        return errors.ErrNotFound // Non-retryable
    }

    err := RetryWithBackoff(fn, DefaultRetryConfig())

    if !errors.Is(err, errors.ErrNotFound) {
        t.Errorf("expected ErrNotFound, got %v", err)
    }
    if attempts != 1 {
        t.Errorf("expected 1 attempt (no retry), got %d", attempts)
    }
}
```

---

## Implementation Hints

1. **No retry on first attempt:** Don't sleep before the first attempt:
   ```go
   for attempt := 1; attempt <= config.MaxAttempts; attempt++ {
       err := fn()
       if err == nil {
           return nil
       }
       // Only sleep if we're going to retry
       if attempt < config.MaxAttempts && IsRetryable(err) {
           time.Sleep(calculateDelay(attempt, config))
       }
   }
   ```

2. **Exponential calculation:** Use `math.Pow()`:
   ```go
   func calculateDelay(attempt int, config RetryConfig) time.Duration {
       multiplier := math.Pow(config.Multiplier, float64(attempt-1))
       delay := time.Duration(float64(config.InitialDelay) * multiplier)
       if delay > config.MaxDelay {
           delay = config.MaxDelay
       }
       return delay
   }
   ```

3. **Track last error:** Keep the last error to return it:
   ```go
   var lastErr error
   for attempt := 1; attempt <= config.MaxAttempts; attempt++ {
       lastErr = fn()
       // ...
   }
   return lastErr
   ```

4. **Type-safe error checking:** Use `errors.As()` to check custom error types:
   ```go
   var apiErr *WeatherAPIError
   if errors.As(err, &apiErr) {
       // Now apiErr is the unwrapped error
       if apiErr.StatusCode >= 500 {
           return true // Retryable
       }
   }
   ```

5. **Testing with fast delays:** In tests, use very short delays to avoid slow tests:
   ```go
   config := RetryConfig{
       InitialDelay: 1 * time.Millisecond,
       MaxDelay:     5 * time.Millisecond,
       // ...
   }
   ```

6. **Optional: Add jitter:** For production systems, add random jitter to prevent thundering herd:
   ```go
   // Optional enhancement (not required for this exercise)
   jitter := time.Duration(rand.Float64() * float64(delay) * 0.1)
   delay = delay + jitter
   ```

---

## Usage Example

```go
// In weather client
func (c *WeatherClient) FetchWeather(ctx context.Context, req WeatherRequest) (*types.WeatherData, error) {
    var result *types.WeatherData

    err := RetryWithBackoff(func() error {
        resp, err := c.doRequest(ctx, req)
        if err != nil {
            return err // Will be classified by IsRetryable
        }
        result = resp
        return nil
    }, DefaultRetryConfig())

    return result, err
}
```
