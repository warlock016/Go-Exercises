package error_handling

import (
	"errors"
	"sync"
	"time"
)

// Result wraps a value and potential error
type Result[T any] struct {
	Value T
	Err   error
}

// ProcessWithErrors processes items concurrently, collecting all errors
func ProcessWithErrors(items []int, process func(int) error) []error {
	// TODO(human): Implement

	var wg sync.WaitGroup

	errs := make(chan error, len(items))
	output := make([]error, 0)

	for _, v := range items {
		wg.Go(func() {
			err := process(v)
			if err != nil {
				errs <- err
			}
		})
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	for err := range errs {
		output = append(output, err)
	}

	return output
}

// FirstError runs tasks concurrently, returns first error or nil
func FirstError(tasks []func() error) error {
	// TODO(human): Implement

	// early return if no tasks present
	if len(tasks) == 0 {
		return nil
	}

	// buffered channel for tracking maximum number of possible errors
	err := make(chan error, len(tasks))
	// var errResult error

	for _, t := range tasks {
		go func(func() error) {
			e := t()
			err <- e
		}(t)
	}

	count := 0
	for e := range err {
		if e != nil {
			return e
		}
		count++

		if count == len(tasks) {
			return nil
		}
	}

	return nil
}

// ProcessResults processes items and returns all results with errors
func ProcessResults[T any](items []int, process func(int) (T, error)) []Result[T] {
	// TODO(human): Implement

	type Order struct {
		Out Result[T]
		Idx int
	}

	var wg sync.WaitGroup

	result := make([]Result[T], len(items))
	res := make(chan Order)

	for i, v := range items {
		wg.Go(func() {
			t, err := process(v)
			out := Result[T]{
				Value: t,
				Err:   err,
			}

			ord := Order{
				Out: out,
				Idx: i,
			}
			res <- ord
		})
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		// result = append(result, r)
		result[r.Idx] = r.Out
	}

	return result
}

// RunWithTimeout runs task with timeout
func RunWithTimeout(task func() error, timeout time.Duration) error {
	// TODO(human): Implement
	ch := make(chan error)

	go func() {
		err := task()
		ch <- err
	}()

	select {
	case err := <-ch:
		return err
	case <-time.After(timeout):
		return errors.New("timeout")
	}
}

// ParallelFetch fetches from URLs concurrently
func ParallelFetch(urls []string, fetch func(string) (string, error)) (map[string]string, []error) {
	// TODO(human): Implement

	type outcome struct {
		Src string
		Res string
		Err error
	}

	var wg sync.WaitGroup

	resCh := make(chan outcome)

	res := make(map[string]string)
	errs := []error{}

	for _, s := range urls {
		wg.Go(func() {
			res, err := fetch(s)

			out := outcome{
				Src: s,
				Res: res,
				Err: err,
			}

			resCh <- out
		})
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	for o := range resCh {
		if o.Err != nil {
			errs = append(errs, o.Err)
		}
		res[o.Src] = o.Res
	}

	return res, errs
}
