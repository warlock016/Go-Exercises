package goroutine_basics

// PrintMessages prints each message in a separate goroutine
// This function demonstrates the problem: messages may not print
// because main doesn't wait for goroutines to complete
func PrintMessages(messages []string) {
	// TODO(human): Implement - spawn a goroutine for each message
}

// PrintMessagesSync prints each message in a goroutine, waiting for all to complete
func PrintMessagesSync(messages []string) {
	// TODO(human): Implement - use sync.WaitGroup
}

// CounterUnsafe spawns n goroutines that each increment a shared counter
// Returns the final counter value (will be less than n due to race condition)
func CounterUnsafe(n int) int {
	// TODO(human): Implement - demonstrates race condition
	return 0
}

// CounterSafe spawns n goroutines that safely increment a shared counter
// Returns the final counter value (should equal n)
func CounterSafe(n int) int {
	// TODO(human): Implement - use sync.Mutex for protection
	return 0
}
