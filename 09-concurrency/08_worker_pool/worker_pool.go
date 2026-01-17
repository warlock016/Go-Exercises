package worker_pool

import (
	"sync"
)

// Job represents a unit of work
type Job struct {
	ID    int
	Input int
}

// Result represents the output of processing a job
type Result struct {
	JobID  int
	Output int
	Err    error
}

// WorkerPool manages a pool of workers
type WorkerPool struct {
	// TODO(human): Define fields
	numWorkers int
	process    func(Job) Result
	jobs       chan Job
	results    chan Result
	wg         sync.WaitGroup
}

// NewWorkerPool creates a pool with the specified number of workers
func NewWorkerPool(numWorkers int, process func(Job) Result) *WorkerPool {
	// TODO(human): Implement
	p := WorkerPool{
		numWorkers: numWorkers,
		process:    process,
		jobs:       make(chan Job, numWorkers*25),
		results:    make(chan Result, numWorkers*25),
	}
	p.Start()
	return &p
}

func (p *WorkerPool) Start() {
	for range p.numWorkers {
		p.wg.Go(func() {
			// fmt.Printf("goroutine %d launched\n", i)
			for job := range p.jobs {
				p.results <- p.process(job)
				// fmt.Printf("goroutine %d pushed processed job %v to results channel\n", i, job)
			}
		})
	}
}

// Submit adds a job to the pool
func (p *WorkerPool) Submit(job Job) {
	// TODO(human): Implement
	p.jobs <- job
}

// Results returns the channel to receive results from
func (p *WorkerPool) Results() <-chan Result {
	// TODO(human): Implement
	return p.results
}

// Shutdown signals workers to stop and waits for completion
func (p *WorkerPool) Shutdown() {
	// TODO(human): Implement
	close(p.jobs)
	// fmt.Println("closed jobs channel, waiting to close results channel")

	go func() {
		p.wg.Wait()
		close(p.results)
		// fmt.Println("closed results channel!")
	}()
}

// ProcessJobs processes all jobs with a worker pool and returns results
func ProcessJobs(jobs []Job, numWorkers int, process func(Job) Result) []Result {
	// TODO(human): Implement
	pool := NewWorkerPool(numWorkers, process)
	results := make([]Result, 0)

	go func() {
		for _, job := range jobs {
			pool.Submit(job)
		}
		pool.Shutdown()
	}()

	for res := range pool.results {
		results = append(results, res)
	}

	return results
}
