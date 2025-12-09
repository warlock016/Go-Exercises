package retry_with_backoff

import (
	"errors"
	"fmt"
	"time"
)

// RetryConfig holds retry configuration
type RetryConfig struct {
	MaxAttempts  int
	InitialDelay time.Duration
	MaxDelay     time.Duration
}

// RetryWithBackoff retries a function with exponential backoff
func RetryWithBackoff(fn func() error, config RetryConfig) error {
	// TODO(human): Implement
	var lastErr error
	for attempt := 0; attempt < config.MaxAttempts; attempt++ {
		lastErr = fn()
		if lastErr == nil {
			return nil
		}

		if attempt < config.MaxAttempts-1 {
			delay := CalculateBackoff(attempt+1, config.InitialDelay)
			// if delay > config.MaxDelay {
			// 	delay = config.MaxDelay
			// }
			delay = min(delay, config.MaxDelay)
			time.Sleep(delay)
		}
	}

	return fmt.Errorf("timeout")
}

var ErrTemporary = errors.New("temporary error")
var ErrTimeout = errors.New("timeout")

// IsRetryable determines if an error should be retried
func IsRetryable(err error) bool {
	// TODO(human): Implement
	if errors.Is(err, ErrTemporary) {
		return true
	}
	if errors.Is(err, ErrTimeout) {
		return true
	}
	return false
}

// CalculateBackoff calculates the delay for the next retry
func CalculateBackoff(attempt int, initialDelay time.Duration) time.Duration {
	// TODO(human): Implement
	for range attempt {
		initialDelay *= 2
	}
	return initialDelay
}
