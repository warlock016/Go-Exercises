package channel_fundamentals

// SendReceive demonstrates basic channel communication
// Creates a channel, sends value in a goroutine, returns the received value
func SendReceive(value int) int {
	// TODO(human): Implement
	ch := make(chan int)

	go func() {
		ch <- value
	}()

	return <-ch
}

// Generator returns a channel that will receive the numbers 1 to n, then close
func Generator(n int) <-chan int {
	// TODO(human): Implement
	ch := make(chan int)

	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

// Sum receives all values from the channel and returns their sum
func Sum(ch <-chan int) int {
	// TODO(human): Implement

	result := 0

	for val := range ch {
		result += val
	}

	return result
}

// Ping sends a message to the returned channel
func Ping(msg string) <-chan string {
	// TODO(human): Implement
	ch := make(chan string)
	go func() {
		ch <- msg
		close(ch)
	}()
	return ch
}

// PingPong bounces a counter between two goroutines n times
// Returns the final counter value
func PingPong(n int) int {
	// TODO(human): Implement

	// early return when n iterations == 0
	if n == 0 {
		return 0
	}

	// result channel is used to send data from "main" to "ping" goroutine (start), and finally to fetch the result after playing ping-pong between both goroutines
	result := make(chan int)
	// comm channel is used to move data between "ping" and "pong" goroutines, as long as counter has not reached 2*n
	comm := make(chan int)

	go func() {
		val := <-result
		val++
		comm <- val

		for range n {
			val = <-comm
			if val >= 2*n {
				result <- val
			} else {
				val++
				comm <- val
			}
		}
	}()

	go func() {
		for range n {
			val := <-comm
			val++
			comm <- val
		}
	}()

	result <- 0
	return <-result
}
