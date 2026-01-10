package buffered_channels

// BufferDemo sends values to a buffered channel and returns them in order
func BufferDemo(values []int, bufferSize int) []int {
	// TODO(human): Implement

	ch := make(chan int, bufferSize)

	go func() {
		for _, v := range values {
			ch <- v
		}
		close(ch)
	}()

	result := []int{}

	for o := range ch {
		result = append(result, o)
	}

	return result
}

// Producer sends all values to a buffered channel and returns the channel
func Producer(values []int, bufferSize int) <-chan int {
	// TODO(human): Implement

	ch := make(chan int, bufferSize)

	// if len(values) == 0 {
	// 	close(ch)
	// 	return ch
	// }

	go func() {
		for _, v := range values {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

// Consumer receives all values from a channel and returns them as a slice
func Consumer(ch <-chan int) []int {
	// TODO(human): Implement

	result := []int{}

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		result = append(result, v)
	}

	return result
}

// BatchCollector collects values into batches of the specified size
func BatchCollector(in <-chan int, batchSize int) <-chan []int {
	// TODO(human): Implement

	out := make(chan []int)
	slice := []int{}

	go func() {
		for val := range in {

			slice = append(slice, val)

			if len(slice) == batchSize {
				out <- slice
				slice = []int{}
			}

		}
		if len(slice) > 0 {
			out <- slice
		}
		close(out)
	}()

	return out
}

// Semaphore uses a buffered channel to limit concurrent operations
func Semaphore(limit int) func(work func()) {
	// TODO(human): Implement
	tokens := make(chan struct{}, limit)
	return func(work func()) {
		tokens <- struct{}{}
		defer func() {
			<-tokens
		}()
		work()
	}
}
