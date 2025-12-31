package channel_fundamentals

// SendReceive demonstrates basic channel communication
// Creates a channel, sends value in a goroutine, returns the received value
func SendReceive(value int) int {
	// TODO(human): Implement
	return 0
}

// Generator returns a channel that will receive the numbers 1 to n, then close
func Generator(n int) <-chan int {
	// TODO(human): Implement
	return nil
}

// Sum receives all values from the channel and returns their sum
func Sum(ch <-chan int) int {
	// TODO(human): Implement
	return 0
}

// Ping sends a message to the returned channel
func Ping(msg string) <-chan string {
	// TODO(human): Implement
	return nil
}

// PingPong bounces a counter between two goroutines n times
// Returns the final counter value
func PingPong(n int) int {
	// TODO(human): Implement
	return 0
}
