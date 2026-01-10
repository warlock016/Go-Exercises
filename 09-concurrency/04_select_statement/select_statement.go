package select_statement

import (
	"errors"
	"time"
)

// IndexedValue represents a value with its source channel index
type IndexedValue struct {
	Index int
	Value int
}

// FirstResponse returns the first value received from either channel
func FirstResponse(ch1, ch2 <-chan string) string {
	// TODO(human): Implement

	select {
	case v := <-ch1:
		return v
	case v := <-ch2:
		return v
	}
}

// WithTimeout receives from channel or returns error if timeout exceeded
func WithTimeout(ch <-chan string, timeout time.Duration) (string, error) {
	// TODO(human): Implement

	select {
	case v := <-ch:
		return v, nil
	case <-time.After(timeout):
		return "", errors.New("timeout")
	}
	// return "", nil
}

// TryReceive attempts non-blocking receive
func TryReceive(ch <-chan int) (int, bool) {
	// TODO(human): Implement
	select {
	case v := <-ch:
		return v, true
	default:
		return 0, false
	}
	// return 0, false
}

// TrySend attempts non-blocking send
func TrySend(ch chan<- int, value int) bool {
	// TODO(human): Implement
	select {
	case ch <- value:
		return true
	default:
		return false
	}
	// return false
}

// Merge combines two channels into a single output channel
func Merge(ch1, ch2 <-chan int) <-chan int {
	// TODO(human): Implement
	out := make(chan int)

	go func() {
		defer close(out)

		for ch1 != nil || ch2 != nil {
			select {
			case v, ok := <-ch1:
				if !ok {
					// set ch1 to nil so that the for loop condition is satisfied during the next iteration
					ch1 = nil
					continue
				}
				out <- v

			case v, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				out <- v
			}
		}
	}()

	return out
}

// Multiplex receives from any of the input channels
func Multiplex(channels ...<-chan int) <-chan IndexedValue {
	// TODO(human): Implement
	out := make(chan IndexedValue)

	go func() {
		defer close(out)

		for i, ch := range channels {
			for ch != nil {
				v, ok := <-ch
				if !ok {
					ch = nil
					continue
				}
				res := IndexedValue{
					Index: i,
					Value: v,
				}
				out <- res
			}
		}
	}()

	return out
}
