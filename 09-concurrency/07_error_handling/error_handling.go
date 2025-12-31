package error_handling

import "time"

// Result wraps a value and potential error
type Result[T any] struct {
	Value T
	Err   error
}

// ProcessWithErrors processes items concurrently, collecting all errors
func ProcessWithErrors(items []int, process func(int) error) []error {
	// TODO(human): Implement
	return nil
}

// FirstError runs tasks concurrently, returns first error or nil
func FirstError(tasks []func() error) error {
	// TODO(human): Implement
	return nil
}

// ProcessResults processes items and returns all results with errors
func ProcessResults[T any](items []int, process func(int) (T, error)) []Result[T] {
	// TODO(human): Implement
	return nil
}

// RunWithTimeout runs task with timeout
func RunWithTimeout(task func() error, timeout time.Duration) error {
	// TODO(human): Implement
	return nil
}

// ParallelFetch fetches from URLs concurrently
func ParallelFetch(urls []string, fetch func(string) (string, error)) (map[string]string, []error) {
	// TODO(human): Implement
	return nil, nil
}
