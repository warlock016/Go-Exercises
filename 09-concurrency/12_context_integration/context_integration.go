package context_integration

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Result wraps a value and potential error
type Result[T any] struct {
	Value T
	Err   error
}

// WorkerWithContext processes items until context is cancelled
func WorkerWithContext(ctx context.Context, items <-chan int, process func(int) error) error {
	// TODO(human): Implement

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case v, ok := <-items:
			if !ok {
				return nil
			}
			if err := process(v); err != nil {
				return err
			}
		}
	}
}

// FetchAll fetches from multiple sources with shared cancellation
func FetchAll(ctx context.Context, urls []string, fetch func(ctx context.Context, url string) (string, error)) (map[string]string, error) {
	// TODO(human): Implement
	var wg sync.WaitGroup
	var mu sync.Mutex
	var err error
	result := make(map[string]string)

	for _, s := range urls {
		wg.Go(func() {
			res, e := fetch(ctx, s)
			mu.Lock()
			defer mu.Unlock()
			if e != nil {
				err = e
				return
			}
			result[s] = res
		})
	}

	wg.Wait()

	return result, err
}

// FetchRace returns result from first successful fetch
func FetchRace(ctx context.Context, urls []string, fetch func(ctx context.Context, url string) (string, error)) (string, error) {
	// TODO(human): Implement

	if len(urls) == 0 {
		return "", errors.New("missing urls")
	}

	type fetchResult struct {
		result string
		err    error
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	result := make(chan fetchResult, len(urls))

	for _, s := range urls {
		go func(s string) {
			res, err := fetch(ctx, s)
			if err != nil {
				result <- fetchResult{
					result: "",
					err:    err,
				}
			} else {
				result <- fetchResult{
					result: res,
					err:    nil,
				}
			}
		}(s)
	}

	select {
	case got := <-result:
		return got.result, got.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// TimeoutPipeline processes input with per-item timeout
func TimeoutPipeline[T, R any](ctx context.Context, input <-chan T, timeout time.Duration, process func(context.Context, T) (R, error)) <-chan Result[R] {
	// TODO(human): Implement
	result := make(chan Result[R])

	go func() {
		defer close(result)

		for {
			select {
			case <-ctx.Done():
				return
			case in, ok := <-input:
				if !ok {
					return
				}

				itemCtx, cancel := context.WithTimeout(ctx, timeout)
				r, err := process(itemCtx, in)
				cancel()

				select {
				case result <- Result[R]{Value: r, Err: err}:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return result
}

// RetryWithBackoff retries operation with exponential backoff
func RetryWithBackoff(ctx context.Context, maxRetries int, initialDelay time.Duration, op func(context.Context) error) error {
	// TODO(human): Implement

	for i := range maxRetries {

		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			if err := op(ctx); err == nil {
				return nil
			} else if i+1 == maxRetries {
				return err
			}

			timeout := initialDelay * time.Duration(1<<i)
			timer := time.NewTimer(timeout)
			select {
			case <-timer.C:
			case <-ctx.Done():
				timer.Stop()
				return ctx.Err()
			}
		}
	}

	return nil
}

// Parallel runs functions concurrently, cancelling all on first error
func Parallel(ctx context.Context, fns ...func(context.Context) error) error {
	// TODO(human): Implement

	if len(fns) == 0 {
		// return errors.New("no functions to execute")
		return nil
	}

	err := make(chan error, len(fns))
	fnCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, fn := range fns {
		go func(f func(context.Context) error) {
			err <- f(fnCtx)
		}(fn)
	}

	var firstErr error
	for range fns {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case e := <-err:
			if e != nil && firstErr == nil {
				cancel()
				firstErr = e
				// return e // if we return here, race conditions arise!
			}
		}
	}

	return firstErr
}
