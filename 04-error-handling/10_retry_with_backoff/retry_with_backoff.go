package retry_with_backoff

import "time"

// RetryConfig holds retry configuration
type RetryConfig struct {
	MaxAttempts  int
	InitialDelay time.Duration
	MaxDelay     time.Duration
}

// RetryWithBackoff retries a function with exponential backoff
func RetryWithBackoff(fn func() error, config RetryConfig) error {
	// TODO(human): Implement
	return nil
}

// IsRetryable determines if an error should be retried
func IsRetryable(err error) bool {
	// TODO(human): Implement
	return false
}

// CalculateBackoff calculates the delay for the next retry
func CalculateBackoff(attempt int, initialDelay time.Duration) time.Duration {
	// TODO(human): Implement
	return 0
}
