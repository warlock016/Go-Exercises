package fan_out_fan_in

import "sync"

// FanOut distributes items from input to n output channels
func FanOut[T any](input <-chan T, n int) []<-chan T {
	// TODO(human): Implement
	result := make([]<-chan T, n)
	for i := range n {
		ch := make(chan T)
		result[i] = ch

		go func() {
			for in := range input {
				ch <- in
			}
			close(ch)
		}()
	}

	return result
}

// FanIn merges multiple input channels into a single output channel
func FanIn[T any](inputs ...<-chan T) <-chan T {
	// TODO(human): Implement

	var wg sync.WaitGroup
	result := make(chan T)

	for _, input := range inputs {
		wg.Go(func() {
			for in := range input {
				result <- in
			}
		})
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

// ParallelMap applies fn to each item using n workers, preserving order
func ParallelMap[T, R any](items []T, n int, fn func(T) R) []R {
	// TODO(human): Implement

	if len(items) == 0 || n == 0 || fn == nil {
		return []R{}
	}

	// wrapper struct to track items and index
	type Item struct {
		Idx  int
		Item T
	}

	// wrapper struct to track result and index
	type Result struct {
		Idx    int
		Result R
	}

	// WaitGroup for synchronization
	var wg sync.WaitGroup

	jobCh := make(chan Item)
	resultCh := make(chan Result)

	results := make([]R, len(items))

	go func() {
		for i, t := range items {
			job := Item{
				Idx:  i,
				Item: t,
			}
			jobCh <- job
		}
		close(jobCh)
	}()

	for range n {
		wg.Go(func() {
			for item := range jobCh {
				res := fn(item.Item)
				out := Result{
					Idx:    item.Idx,
					Result: res,
				}
				resultCh <- out
			}
		})
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for res := range resultCh {
		results[res.Idx] = res.Result
	}

	return results
}

// ParallelMapStream streams results as they complete (order not preserved)
func ParallelMapStream[T, R any](input <-chan T, n int, fn func(T) R) <-chan R {
	// TODO(human): Implement
	var wg sync.WaitGroup
	results := make(chan R)

	for range n {
		wg.Go(func() {
			for in := range input {
				res := fn(in)
				results <- res
			}
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

// ProcessWithFanOut processes input with fan-out/fan-in pattern
func ProcessWithFanOut[T, R any](input <-chan T, n int, process func(T) R) <-chan R {
	// TODO(human): Implement

	var wg sync.WaitGroup

	pipes := FanOut(input, n)
	result := make(chan R)

	for i := range n {
		wg.Go(func() {
			for ch := range pipes[i] {
				res := process(ch)
				result <- res
			}
		})
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}
