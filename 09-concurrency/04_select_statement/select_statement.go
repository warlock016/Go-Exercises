package select_statement

import "time"

// IndexedValue represents a value with its source channel index
type IndexedValue struct {
	Index int
	Value int
}

// FirstResponse returns the first value received from either channel
func FirstResponse(ch1, ch2 <-chan string) string {
	// TODO(human): Implement
	return ""
}

// WithTimeout receives from channel or returns error if timeout exceeded
func WithTimeout(ch <-chan string, timeout time.Duration) (string, error) {
	// TODO(human): Implement
	return "", nil
}

// TryReceive attempts non-blocking receive
func TryReceive(ch <-chan int) (int, bool) {
	// TODO(human): Implement
	return 0, false
}

// TrySend attempts non-blocking send
func TrySend(ch chan<- int, value int) bool {
	// TODO(human): Implement
	return false
}

// Merge combines two channels into a single output channel
func Merge(ch1, ch2 <-chan int) <-chan int {
	// TODO(human): Implement
	return nil
}

// Multiplex receives from any of the input channels
func Multiplex(channels ...<-chan int) <-chan IndexedValue {
	// TODO(human): Implement
	return nil
}
