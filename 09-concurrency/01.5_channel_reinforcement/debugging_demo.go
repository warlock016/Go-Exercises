package channel_reinforcement

import (
	"sync"
)

// =============================================================================
// DEBUGGING DEMO - Contains intentional bugs for learning
// =============================================================================
//
// This file contains 3 buggy functions for practicing concurrency debugging.
// Each function has 1-3 bugs. Your task:
//
// 1. Run the tests: go test -v -run TestBuggy
// 2. Use debugging tools to find the bugs
// 3. DO NOT fix the bugs in this file - just identify them
//
// Suggested debugging approaches:
// - go test -race -run TestBuggy
// - dlv test -- -test.run TestBuggy
// - Add timeout wrappers
// - Check goroutine states with dlv's `goroutines` command
//
// =============================================================================

// BuggyWorkerPool attempts to process values through a worker pool
// Expected: Launch 3 workers, each doubles values, return sum of all results
// Bugs: 2

// fixed!
func BuggyWorkerPool(values []int) int {
	if len(values) == 0 {
		return 0
	}

	jobs := make(chan int)
	results := make(chan int)

	// Launch workers
	var wg sync.WaitGroup
	numWorkers := 3

	for range numWorkers {
		wg.Go(func() {
			// defer statement automatically handled by wg.Go(), check its code for additional context
			// defer wg.Done()
			for job := range jobs {
				results <- 2 * job
			}
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Send jobs
	go func() {
		for _, v := range values {
			jobs <- v
		}
		close(jobs)
	}()

	// Collect results
	sum := 0
	for {
		v, ok := <-results
		if !ok {
			break
		}
		sum += v
	}

	return sum
}

// BuggyPingPong bounces a value between two goroutines n times
// Expected: Value starts at 0, each bounce adds 1, return final value (should equal n)
// Bugs: 2

// fixed!
func BuggyPingPong(n int) int {
	if n == 0 {
		return 0
	}

	ping := make(chan int)
	pong := make(chan int)
	result := make(chan int)

	// Ping goroutine
	go func() {
		for i := range n {
			val := <-ping
			val++
			if i == n-1 {
				result <- val
				close(result)
			} else {
				pong <- val
			}
		}
	}()

	// Pong goroutine
	go func() {
		for range n - 1 {
			val := <-pong
			ping <- val
		}
	}()

	// Start the game
	ping <- 0

	// Get final result
	return <-result
}

// BuggyFanOut distributes work to multiple goroutines and collects results
// Expected: Split values among workers, each worker sums its portion, return total
// Bugs: 3

// fixed!
func BuggyFanOut(values []int, numWorkers int) int {
	if len(values) == 0 || numWorkers == 0 {
		return 0
	}

	var wg sync.WaitGroup
	results := make(chan int)
	chunkSize := len(values) / numWorkers

	// fmt.Printf("Workers: %d, work: %v, chunk size: %d\n", numWorkers, values, chunkSize)

	// Launch workers
	for i := range numWorkers {
		wg.Go(func() {
			start := i * chunkSize
			end := start + chunkSize
			if i == numWorkers-1 {
				end = len(values)
			}

			sum := 0
			for j := start; j < end; j++ {
				sum += values[j] // Process chunk
			}
			results <- sum
		})
		// go func(workerID int) {
		// 	start := workerID * chunkSize
		// 	end := start + chunkSize
		// 	if workerID == numWorkers-1 {
		// 		end = len(values)
		// 	}

		// 	sum := 0
		// 	for j := start; j < end; j++ {
		// 		sum += values[i] // Process chunk
		// 	}
		// 	results <- sum
		// }(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	total := 0
	for {
		v, ok := <-results
		if !ok {
			break
		}
		total += v
	}

	return total
}
