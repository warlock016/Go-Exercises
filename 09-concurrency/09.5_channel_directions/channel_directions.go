package channel_directions

import "sync"

// =============================================================================
// Channel Directions Practice Exercise
// =============================================================================
//
// This exercise focuses on understanding and correctly using channel direction
// types in function signatures:
//
//   - chan T      : Bidirectional (send and receive)
//   - <-chan T    : Receive-only (can only receive and range)
//   - chan<- T    : Send-only (can only send)
//
// Key concepts to practice:
//   - Channels are ALWAYS created as bidirectional with make(chan T)
//   - Direction types are used in function SIGNATURES to restrict access
//   - The sender/creator closes the channel, never the receiver
//   - Returning <-chan T prevents callers from closing your channel
//   - Accepting <-chan T means you can only read (someone else closes)
//
// =============================================================================

// =============================================================================
// Function 1: Numbers (Generator Pattern)
// =============================================================================
//
// Goal: Create a generator that emits integers from 1 to n, then closes.
//
// Behavior:
//   - Takes an integer n as input
//   - Returns a channel that will emit values 1, 2, 3, ..., n
//   - Channel should close after emitting all values
//   - Values should be emitted in a separate goroutine (non-blocking return)
//
// Think about:
//   - What channel direction should the RETURN type be?
//   - Who owns closing the channel?
//   - What happens if n <= 0?
//
// Example usage:
//   ch := Numbers(5)
//   for v := range ch {
//       fmt.Println(v)  // Prints: 1, 2, 3, 4, 5
//   }

// TODO(human): Implement Numbers - design the signature and implementation
func Generator(n int) <-chan int {

	// should I return nil or rather declare the channel first, close it and then return a closed channel?
	// the closed channel return pattern seems more robust due to it's cascading effect (?)
	result := make(chan int)

	if n <= 0 {
		close(result)
		return result
	}

	go func() {
		for i := range n {
			result <- (i + 1)
		}
		close(result)
	}()

	return result
}

// =============================================================================
// Function 2: Square (Pipeline Stage Pattern)
// =============================================================================
//
// Goal: Create a pipeline stage that squares each incoming value.
//
// Behavior:
//   - Takes an input channel of integers
//   - Returns an output channel that emits the square of each input value
//   - Output channel should close when input channel closes
//   - Processing should happen in a separate goroutine
//
// Think about:
//   - What channel direction should the INPUT parameter be?
//   - What channel direction should the RETURN type be?
//   - Who closes what?
//
// Example usage:
//   nums := Numbers(3)       // 1, 2, 3
//   squared := Square(nums)  // 1, 4, 9
//   for v := range squared {
//       fmt.Println(v)
//   }

// TODO(human): Implement Square - design the signature and implementation
func Square(input <-chan int) <-chan int {
	result := make(chan int)

	go func() {
		for v := range input {
			result <- (v * v)
		}
		close(result)
	}()

	return result
}

// =============================================================================
// Function 3: Filter (Pipeline Stage with Predicate)
// =============================================================================
//
// Goal: Create a pipeline stage that only passes values matching a predicate.
//
// Behavior:
//   - Takes an input channel of integers and a predicate function
//   - Returns an output channel that only emits values where predicate returns true
//   - Output channel should close when input channel closes
//
// Think about:
//   - Same channel direction considerations as Square
//   - What is the signature of the predicate function?
//
// Example usage:
//   nums := Numbers(10)
//   evens := Filter(nums, func(n int) bool { return n%2 == 0 })
//   // evens emits: 2, 4, 6, 8, 10

// TODO(human): Implement Filter - design the signature and implementation
func Filter(input <-chan int, fn func(int) bool) <-chan int {

	result := make(chan int)

	go func() {
		for v := range input {
			if fn(v) {
				result <- v
			}
		}
		close(result)
	}()
	return result
}

// =============================================================================
// Function 4: Sum (Consumer/Sink Pattern)
// =============================================================================
//
// Goal: Consume all values from a channel and return their sum.
//
// Behavior:
//   - Takes an input channel of integers
//   - Reads all values until channel closes
//   - Returns the sum of all values
//   - This function BLOCKS until the input channel closes
//
// Think about:
//   - What channel direction should the INPUT parameter be?
//   - This is a consumer - should it close anything?
//   - What should it return if the channel is empty?
//
// Example usage:
//   nums := Numbers(5)       // 1, 2, 3, 4, 5
//   total := Sum(nums)       // 15

// TODO(human): Implement Sum - design the signature and implementation
func Sum(in <-chan int) int {

	result := 0
	for v := range in {
		result += v
	}
	return result
}

// =============================================================================
// Function 5: Merge (Fan-In Pattern)
// =============================================================================
//
// Goal: Merge multiple input channels into a single output channel.
//
// Behavior:
//   - Takes multiple input channels (variadic)
//   - Returns a single output channel that emits all values from all inputs
//   - Output channel closes when ALL input channels have closed
//   - Order of output values is non-deterministic
//
// Think about:
//   - What channel direction for inputs? (hint: variadic of channels)
//   - What channel direction for output?
//   - How do you know when ALL inputs are done?
//   - What if zero channels are passed?
//
// Example usage:
//   ch1 := Numbers(3)  // 1, 2, 3
//   ch2 := Numbers(3)  // 1, 2, 3
//   merged := Merge(ch1, ch2)
//   // merged emits all 6 values (order varies)

// TODO(human): Implement Merge - design the signature and implementation
func Merge(input ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	result := make(chan int)

	for _, ch := range input {
		wg.Go(func() {
			for v := range ch {
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

// =============================================================================
// Function 6: Tee (Broadcast Pattern)
// =============================================================================
//
// Goal: Split one input channel into two output channels (broadcast).
//
// Behavior:
//   - Takes one input channel
//   - Returns TWO output channels
//   - Each value from input is sent to BOTH output channels
//   - Both output channels close when input closes
//
// Think about:
//   - What happens if one output consumer is slower than the other?
//   - Who owns closing the output channels?
//   - This is different from fan-out (which distributes, not broadcasts)
//
// Example usage:
//   nums := Numbers(3)
//   ch1, ch2 := Tee(nums)
//   // Both ch1 and ch2 receive: 1, 2, 3

// TODO(human): Implement Tee - design the signature and implementation
func Tee(in <-chan int) (<-chan int, <-chan int) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for v := range in {
			ch1 <- v
			ch2 <- v
		}
		close(ch1)
		close(ch2)
	}()
	return ch1, ch2
}

// =============================================================================
// Putting It All Together: Pipeline
// =============================================================================
//
// Once you've implemented the above functions, you can compose them:
//
//   nums := Numbers(10)                           // 1-10
//   squared := Square(nums)                       // 1, 4, 9, 16, 25, 36, 49, 64, 81, 100
//   evens := Filter(squared, isEven)              // 4, 16, 36, 64, 100
//   total := Sum(evens)                           // 220
//
// Or with Tee and Merge:
//
//   nums := Numbers(5)
//   a, b := Tee(nums)
//   doubled := Square(a)        // Process one stream
//   filtered := Filter(b, isOdd) // Process the other differently
//   combined := Merge(doubled, filtered)
//
// =============================================================================
