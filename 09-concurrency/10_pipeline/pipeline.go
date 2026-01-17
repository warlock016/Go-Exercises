package pipeline

import "sync"

// Stage is a function that processes input channel and returns output channel
type Stage[I, O any] func(input <-chan I) <-chan O

// Generator creates a channel that emits values from a slice
func Generator[T any](values ...T) <-chan T {
	// TODO(human): Implement
	result := make(chan T)

	go func() {
		defer close(result)
		for _, v := range values {
			result <- v
		}
	}()

	return result
}

// Filter keeps only values that pass the predicate
func Filter[T any](predicate func(T) bool) Stage[T, T] {
	// TODO(human): Implement
	return func(input <-chan T) <-chan T {

		result := make(chan T)
		go func() {
			defer close(result)
			for v := range input {
				if predicate(v) {
					result <- v
				}
			}
		}()
		return result
	}
}

// Map transforms each value
func Map[I, O any](transform func(I) O) Stage[I, O] {
	// TODO(human): Implement
	return func(input <-chan I) <-chan O {

		result := make(chan O)
		go func() {
			defer close(result)
			for v := range input {
				result <- transform(v)
			}
		}()
		return result
	}
}

// Take emits only the first n values
func Take[T any](n int) Stage[T, T] {
	// TODO(human): Implement
	return func(input <-chan T) <-chan T {

		result := make(chan T)
		go func() {
			defer close(result)
			for range n {
				v, ok := <-input
				if !ok {
					break
				}
				result <- v
			}
		}()

		return result
	}
}

// Skip discards the first n values
func Skip[T any](n int) Stage[T, T] {
	// TODO(human): Implement
	return func(input <-chan T) <-chan T {
		result := make(chan T)

		go func() {
			defer close(result)
			for range n {
				<-input
			}

			for {
				v, ok := <-input
				if !ok {
					return
				}
				result <- v
			}
		}()

		return result
	}
}

// Reduce collects all values into a single result
func Reduce[T, R any](initial R, reducer func(R, T) R) func(<-chan T) R {
	// TODO(human): Implement
	return func(c <-chan T) R {
		acc := initial
		for v := range c {
			acc = reducer(acc, v)
		}

		return acc
	}
}

// Pipeline chains multiple stages together
func Pipeline[T any](input <-chan T, stages ...Stage[T, T]) <-chan T {
	// TODO(human): Implement

	result := input

	for _, stage := range stages {
		result = stage(result)
	}

	return result
}

// PipelineAsync runs each stage with n workers
func PipelineAsync[T any](input <-chan T, n int, stages ...Stage[T, T]) <-chan T {
	// TODO(human): Implement (bonus)
	result := make(chan T)
	var wg sync.WaitGroup

	for range n {
		wg.Go(func() {
			out := Pipeline(input, stages...)
			for v := range out {
				result <- v
			}
		})
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}
