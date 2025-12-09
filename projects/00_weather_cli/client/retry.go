package client

import (
	"errors"
	"fmt"
	"time"

	appErrors "github.com/warlock016/weather_cli/errors"
)

type RetryConfig struct {
	MaxAttempts  int
	InitialDelay time.Duration
	MaxDelay     time.Duration
	Multiplier   float64
}

func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxAttempts:  4,
		InitialDelay: time.Millisecond * 200,
		MaxDelay:     time.Second * 3,
		Multiplier:   2,
	}
}

func RetryWithBackoff(fn func() error, config RetryConfig) error {

	delay := config.InitialDelay
	var err error = nil
	for i := range config.MaxAttempts {
		err = fn()
		if err == nil {
			return nil
		}
		if !IsRetriable(err) {
			return fmt.Errorf("non-retriable error %v", err)
		}
		if i == config.MaxAttempts-1 {
			return fmt.Errorf("reached max attempts: %v", err)
		}
		time.Sleep(delay)
		delay = time.Duration(config.Multiplier * float64(delay))
		delay = min(delay, config.MaxDelay)

	}
	return nil
}

func IsRetriable(err error) bool {
	if errors.Is(err, appErrors.ErrTimeout) || errors.Is(err, appErrors.ErrNetwork) || errors.Is(err, appErrors.ErrRateLimited) {
		return true
	}
	return false
}
