package done_channel

import "time"

// Worker processes items until done channel is closed
func Worker(items <-chan int, done <-chan struct{}) int {
	// TODO(human): Implement
	return 0
}

// CancellableLoop runs work function repeatedly until done is closed
func CancellableLoop(work func(), done <-chan struct{}) int {
	// TODO(human): Implement
	return 0
}

// Broadcaster sends value to all output channels, respecting done
func Broadcaster(value int, outputs []chan<- int, done <-chan struct{}) int {
	// TODO(human): Implement
	return 0
}

// Generator produces integers 0, 1, 2, ... until done is closed
func Generator(done <-chan struct{}) <-chan int {
	// TODO(human): Implement
	return nil
}

// Timeout returns a done channel that closes after duration
func Timeout(d time.Duration) <-chan struct{} {
	// TODO(human): Implement
	return nil
}

// MergeCancel combines done channels - closes when ANY input closes
func MergeCancel(done1, done2 <-chan struct{}) <-chan struct{} {
	// TODO(human): Implement
	return nil
}
