package collections

// TODO(human): Define Stack struct for LIFO (Last In First Out) data structure

type Stack struct {
	values []int
}

// Push adds a value to the top of the stack
func (s *Stack) Push(value int) {
	// TODO(human): Add value to top of stack
	s.values = append(s.values, value)
}

// Pop removes and returns the top value (returns value, true if success)
func (s *Stack) Pop() (int, bool) {
	// TODO(human): Remove and return top value from stack
	var removed int

	if len(s.values) <= 0 {
		return removed, false
	} else {
		removed = s.values[len(s.values)-1]
		s.values = s.values[:len(s.values)-1]
	}

	return removed, true
}

// Peek returns the top value without removing it (returns value, true if success)
func (s *Stack) Peek() (int, bool) {
	// TODO(human): Return top value without removing it

	if len(s.values) <= 0 {
		return 0, false
	} else {
		return s.values[len(s.values)-1], true
	}

	// return 0, false
}

// IsEmpty returns true if stack has no elements
func (s *Stack) IsEmpty() bool {
	// TODO(human): Return true if stack is empty
	return len(s.values) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	// TODO(human): Return number of elements in stack
	return len(s.values)
}

// TODO(human): Define Queue struct for FIFO (First In First Out) data structure

type Queue struct {
	values []int
}

// Enqueue adds a value to the back of the queue
func (q *Queue) Enqueue(value int) {
	q.values = append(q.values, value)
	// TODO(human): Add value to back of queue
}

// Dequeue removes and returns the front value (returns value, true if success)
func (q *Queue) Dequeue() (int, bool) {
	// TODO(human): Remove and return front value from queue
	var removed int

	if len(q.values) == 0 {
		return removed, false
	}

	removed = q.values[0]
	// fmt.Printf("Enqueued: %v / length: %d\n", q.values, len(q.values))
	q.values = q.values[1:] // narrow the slice's window starting from the second element until the end
	// fmt.Printf("Dequeued: %v / length: %d\n", q.values, len(q.values))

	return removed, true
}

// Front returns the front value without removing it (returns value, true if success)
func (q *Queue) Front() (int, bool) {
	// TODO(human): Return front value without removing it
	if len(q.values) == 0 {
		return 0, false
	}
	return q.values[0], true
}

// IsEmpty returns true if queue has no elements
func (q *Queue) IsEmpty() bool {
	// TODO(human): Return true if queue is empty
	return len(q.values) == 0
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	// TODO(human): Return number of elements in queue
	return len(q.values)
}
