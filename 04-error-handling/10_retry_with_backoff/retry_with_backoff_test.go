package retry_with_backoff

import (
	"errors"
	"testing"
	"time"
)

func TestCalculateBackoff(t *testing.T) {
	tests := []struct {
		name         string
		attempt      int
		initialDelay time.Duration
		wantMin      time.Duration
	}{
		{"first retry", 1, 100 * time.Millisecond, 100 * time.Millisecond},
		{"second retry", 2, 100 * time.Millisecond, 200 * time.Millisecond},
		{"third retry", 3, 100 * time.Millisecond, 400 * time.Millisecond},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateBackoff(tt.attempt, tt.initialDelay)
			if got < tt.wantMin {
				t.Errorf("CalculateBackoff() = %v, want at least %v", got, tt.wantMin)
			}
		})
	}
}

func TestRetryWithBackoff(t *testing.T) {
	attempts := 0
	fn := func() error {
		attempts++
		if attempts < 3 {
			return errors.New("temporary error")
		}
		return nil
	}

	config := RetryConfig{
		MaxAttempts:  5,
		InitialDelay: 1 * time.Millisecond,
		MaxDelay:     100 * time.Millisecond,
	}

	err := RetryWithBackoff(fn, config)
	if err != nil {
		t.Errorf("RetryWithBackoff() failed: %v", err)
	}
}
