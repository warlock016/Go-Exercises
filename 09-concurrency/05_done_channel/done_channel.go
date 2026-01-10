package done_channel

import (
	"time"
)

// Worker processes items until done channel is closed
func Worker(items <-chan int, done <-chan struct{}) int {
	// TODO(human): Implement

	/*
		Select's fairness problem:
		When multiple cases are ready, Go picks randomly. This is usually fine, but when cancellation needs priority, check it explicitly first.
		This pattern is common in production code where graceful shutdown needs to be responsive.
	*/

	count := 0
	for {
		// priority check for done
		select {
		case <-done:
			return count
		default:
		}

		// then continue with normal processing
		select {
		case _, ok := <-items:
			if !ok {
				return count
			}
			count++
		case <-done:
			return count
		}
	}
}

// CancellableLoop runs work function repeatedly until done is closed
func CancellableLoop(work func(), done <-chan struct{}) int {
	// TODO(human): Implement
	count := 0

	for {
		select {
		case <-done:
			return count
		default:
			work()
			count++
		}
	}
}

// Broadcaster sends value to all output channels, respecting done
func Broadcaster(value int, outputs []chan<- int, done <-chan struct{}) int {
	// TODO(human): Implement
	count := 0
	for _, ch := range outputs {
		select {
		case <-done:
			// if sender closes, then receiver loses control and cannot reuse the channels at a later point in time.
			// for _, ch := range outputs {
			// 	close(ch)
			// }
			return count
		case ch <- value:
			count++
		}
	}
	return count
}

// Generator produces integers 0, 1, 2, ... until done is closed
func Generator(done <-chan struct{}) <-chan int {
	// TODO(human): Implement
	out := make(chan int)
	// count := 0

	go func() {
		defer close(out)
		for i := 0; ; i++ {
			select {
			case <-done:
				return
			case out <- i:
			}
		}
	}()

	return out
}

// Timeout returns a done channel that closes after duration
func Timeout(d time.Duration) <-chan struct{} {
	// TODO(human): Implement

	signal := make(chan struct{})

	go func() {
		time.Sleep(d)
		// blocks if no receiver!
		// signal <- struct{}{}
		close(signal)
	}()

	return signal
}

// MergeCancel combines done channels - closes when ANY input closes
func MergeCancel(done1, done2 <-chan struct{}) <-chan struct{} {
	// TODO(human): Implement

	out := make(chan struct{})

	go func() {
		defer close(out)

		// compared to Merge in select_statement, we only need to wait for a single signal from either done1 or done 2, whatever comes first.
		// select blocks until either signal is sent via the channel.
		// As soon as a signal arrives, the select statement unblocks and the out channel is closed via the defer statement.
		// since struct channels are "data-less", we do not need to read their data, a close channel statement suffices.
		select {
		case <-done1:
		case <-done2:
		}
	}()
	return out
}
