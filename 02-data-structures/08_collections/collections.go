package collections

// TODO(human): Define Stack struct for LIFO (Last In First Out) data structure

// Push adds a value to the top of the stack
func (s *Stack) Push(value int) {
	// TODO(human): Add value to top of stack
}

// Pop removes and returns the top value (returns value, true if success)
func (s *Stack) Pop() (int, bool) {
	// TODO(human): Remove and return top value from stack
	return 0, false
}

// Peek returns the top value without removing it (returns value, true if success)
func (s *Stack) Peek() (int, bool) {
	// TODO(human): Return top value without removing it
	return 0, false
}

// IsEmpty returns true if stack has no elements
func (s *Stack) IsEmpty() bool {
	// TODO(human): Return true if stack is empty
	return false
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	// TODO(human): Return number of elements in stack
	return 0
}

// TODO(human): Define Queue struct for FIFO (First In First Out) data structure

// Enqueue adds a value to the back of the queue
func (q *Queue) Enqueue(value int) {
	// TODO(human): Add value to back of queue
}

// Dequeue removes and returns the front value (returns value, true if success)
func (q *Queue) Dequeue() (int, bool) {
	// TODO(human): Remove and return front value from queue
	return 0, false
}

// Front returns the front value without removing it (returns value, true if success)
func (q *Queue) Front() (int, bool) {
	// TODO(human): Return front value without removing it
	return 0, false
}

// IsEmpty returns true if queue has no elements
func (q *Queue) IsEmpty() bool {
	// TODO(human): Return true if queue is empty
	return false
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	// TODO(human): Return number of elements in queue
	return 0
}
