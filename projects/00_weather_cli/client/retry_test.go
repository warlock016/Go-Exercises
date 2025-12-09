package client_test

import (
	"testing"
	"time"

	"github.com/warlock016/weather_cli/client"
	"github.com/warlock016/weather_cli/errors"
)

func TestSucessFirst(t *testing.T) {

	attempts := 0
	err := client.RetryWithBackoff(func() error {
		attempts++
		return nil
	}, client.DefaultRetryConfig())

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if attempts != 1 {
		t.Errorf("expected 1 attempt, got %d", attempts)
	}
}

func TestSuccessAfterRetries(t *testing.T) {
	config := client.DefaultRetryConfig()
	attempts := 0
	err := client.RetryWithBackoff(func() error {
		attempts++
		if attempts < config.MaxAttempts {
			return errors.ErrNetwork
		}
		return nil
	}, config)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if attempts != config.MaxAttempts {
		t.Errorf("expected %d, got %d", config.MaxAttempts, attempts)
	}
}

func TestNonRetriable(t *testing.T) {
	config := client.DefaultRetryConfig()
	attempts := 0
	err := client.RetryWithBackoff(func() error {
		attempts++
		return errors.ErrInvalidInput
	}, config)

	if err == nil {
		t.Error("expected error, got nil")
	}

	if attempts != 1 {
		t.Errorf("unexpected retry count %d", attempts)
	}
}

func TestReachMaxAttempts(t *testing.T) {
	config := client.DefaultRetryConfig()
	attempts := 0
	err := client.RetryWithBackoff(func() error {
		attempts++
		return errors.ErrNetwork
	}, config)

	if err == nil {
		t.Errorf("expected error, got nil at %d attempt", attempts)
	}
	if attempts != config.MaxAttempts {
		t.Errorf("expected %d, got %d", config.MaxAttempts, attempts)
	}
}

func TestBackoffTiming(t *testing.T) {
	// t.Skip()
	config := client.DefaultRetryConfig()
	attempts := 0
	measurements := []time.Time{}
	durations := []time.Duration{}

	err := client.RetryWithBackoff(func() error {
		measurements = append(measurements, time.Now())
		attempts++
		if attempts < config.MaxAttempts {
			return errors.ErrNetwork
		} else {
			return nil
		}
	}, config)

	if err != nil {
		t.Errorf("unexpected error: got %+v, want nil", err)
	}

	if len(measurements) == 0 {
		t.Error("invalid measurements, empty slice")
	}

	for i := range measurements {
		if i < len(measurements)-1 {
			delta := measurements[i+1].Sub(measurements[i])
			durations = append(durations, delta)
		}
	}

	expectedDelay := config.InitialDelay
	tolerance := 0.1

	for i := 0; i < len(measurements)-1; i++ {
		actual := measurements[i+1].Sub(measurements[i])
		minExpected := time.Duration(float64(expectedDelay) * (1 - tolerance))
		maxExpected := time.Duration(float64(expectedDelay) * (1 + tolerance))

		if actual < minExpected || actual > maxExpected {
			t.Errorf("delay %d: got %v, want between %v and %v", i+1, actual, minExpected, maxExpected)
		}

		expectedDelay = min(time.Duration(float64(expectedDelay)*config.Multiplier), config.MaxDelay)
	}

}
