package polymorphic_collection

// TODO(human): Define Prioritizable interface

// TODO(human): Define Task struct

// TODO(human): Define Message struct

// TODO(human): Define PriorityQueue struct

// TODO(human): Implement Priority() for Task

// TODO(human): Implement Priority() for Message

// NewTask creates a new Task
func NewTask(priority int, description string) Task {
	// TODO(human): Implement
	return Task{}
}

// NewMessage creates a new Message
func NewMessage(priority int, text string) Message {
	// TODO(human): Implement
	return Message{}
}

// Push adds an item to the queue
func (pq *PriorityQueue) Push(item Prioritizable) {
	// TODO(human): Implement
}

// Pop removes and returns the highest priority item
func (pq *PriorityQueue) Pop() (Prioritizable, bool) {
	// TODO(human): Implement
	return nil, false
}

// Peek returns the highest priority item without removing it
func (pq *PriorityQueue) Peek() (Prioritizable, bool) {
	// TODO(human): Implement
	return nil, false
}

// Len returns the number of items in the queue
func (pq *PriorityQueue) Len() int {
	// TODO(human): Implement
	return 0
}
