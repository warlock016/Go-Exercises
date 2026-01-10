package goroutine_basics

import (
	"fmt"
	"sync"
)

// PrintMessages prints each message in a separate goroutine
// This function demonstrates the problem: messages may not print
// because main doesn't wait for goroutines to complete
func PrintMessages(messages []string) {
	// TODO(human): Implement - spawn a goroutine for each message
	var wg sync.WaitGroup
	for _, message := range messages {
		wg.Go(func() {
			fmt.Println(message)
		})
	}
}

// PrintMessagesSync prints each message in a goroutine, waiting for all to complete
func PrintMessagesSync(messages []string) {
	// TODO(human): Implement - use sync.WaitGroup
	var wg sync.WaitGroup

	for _, message := range messages {
		wg.Go(func() {
			fmt.Println(message)
		})
	}
	wg.Wait()
}

// CounterUnsafe spawns n goroutines that each increment a shared counter
// Returns the final counter value (will be less than n due to race condition)
func CounterUnsafe(n int) int {
	// TODO(human): Implement - demonstrates race condition
	var wg sync.WaitGroup

	result := 0
	for range n {
		wg.Go(func() {
			result += 1
		})
	}

	wg.Wait()
	return result
}

// CounterSafe spawns n goroutines that safely increment a shared counter
// Returns the final counter value (should equal n)
func CounterSafe(n int) int {
	// TODO(human): Implement - use sync.Mutex for protection
	mu := sync.Mutex{}
	var wg sync.WaitGroup
	result := 0

	for range n {
		wg.Go(func() {
			mu.Lock()
			result += 1
			mu.Unlock()
		})
	}

	wg.Wait()
	return result
}
