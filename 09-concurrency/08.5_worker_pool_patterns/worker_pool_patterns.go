package worker_pool_patterns

import (
	"context"
	"sync"
)

// Job represents a unit of work to be processed
type Job struct {
	ID   int
	Data int
}

// Result represents the output of processing a job
type Result struct {
	JobID  int
	Output int
	Err    error
}

// StreamingPool processes jobs from a channel using a fixed worker pool.
// Returns a results channel that closes when all jobs are processed.
// This pattern naturally self-regulates through backpressure.
func StreamingPool(jobs <-chan Job, numWorkers int, process func(Job) Result) <-chan Result {
	// TODO(human): Implement
	var wg sync.WaitGroup
	results := make(chan Result)

	for range numWorkers {
		wg.Go(func() {
			for job := range jobs {
				result := process(job)
				results <- result
			}
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	// close(results) // Stub: return closed channel so tests don't hang
	return results
}

// ProcessBatch processes a slice of jobs and returns all results.
// Internally uses StreamingPool with concurrent producer/consumer.
func ProcessBatch(jobs []Job, numWorkers int, process func(Job) Result) []Result {
	// TODO(human): Implement

	var wg sync.WaitGroup

	jobCh := make(chan Job)
	resultCh := make(chan Result)

	go func() {
		for _, job := range jobs {
			jobCh <- job
		}
		close(jobCh)
	}()

	for range numWorkers {
		wg.Go(func() {
			for job := range jobCh {
				resultCh <- process(job)
			}
		})
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	results := make([]Result, 0)

	for result := range resultCh {
		results = append(results, result)
	}

	return results
}

// PoolWithCancel processes jobs until context is cancelled.
// In-flight jobs complete gracefully before results channel closes.
func PoolWithCancel(ctx context.Context, jobs <-chan Job, numWorkers int, process func(Job) Result) <-chan Result {
	// TODO(human): Implement
	var wg sync.WaitGroup
	results := make(chan Result)

	for range numWorkers {
		wg.Go(func() {
			for {
				select {
				case job, ok := <-jobs:
					if !ok {
						return
					}
					results <- process(job)
				case <-ctx.Done():
					return
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

// PoolWithErrors processes jobs, returning separate channels for results and errors.
// Both channels close when all jobs are processed.
func PoolWithErrors(jobs <-chan Job, numWorkers int, process func(Job) (Result, error)) (<-chan Result, <-chan error) {
	// TODO(human): Implement
	results := make(chan Result)
	errs := make(chan error)

	var wg sync.WaitGroup

	for range numWorkers {
		wg.Go(func() {
			for job := range jobs {
				res, err := process(job)
				if err != nil {
					errs <- err
				} else {
					results <- res
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(results)
		close(errs)
	}()

	return results, errs
}
