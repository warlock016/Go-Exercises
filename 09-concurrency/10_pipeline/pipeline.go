package pipeline

// Stage is a function that processes input channel and returns output channel
type Stage[I, O any] func(input <-chan I) <-chan O

// Generator creates a channel that emits values from a slice
func Generator[T any](values ...T) <-chan T {
	// TODO(human): Implement
	return nil
}

// Filter keeps only values that pass the predicate
func Filter[T any](predicate func(T) bool) Stage[T, T] {
	// TODO(human): Implement
	return nil
}

// Map transforms each value
func Map[I, O any](transform func(I) O) Stage[I, O] {
	// TODO(human): Implement
	return nil
}

// Take emits only the first n values
func Take[T any](n int) Stage[T, T] {
	// TODO(human): Implement
	return nil
}

// Skip discards the first n values
func Skip[T any](n int) Stage[T, T] {
	// TODO(human): Implement
	return nil
}

// Reduce collects all values into a single result
func Reduce[T, R any](initial R, reducer func(R, T) R) func(<-chan T) R {
	// TODO(human): Implement
	return nil
}

// Pipeline chains multiple stages together
func Pipeline[T any](input <-chan T, stages ...Stage[T, T]) <-chan T {
	// TODO(human): Implement
	return nil
}

// PipelineAsync runs each stage with n workers
func PipelineAsync[T any](input <-chan T, n int, stages ...Stage[T, T]) <-chan T {
	// TODO(human): Implement (bonus)
	return nil
}
