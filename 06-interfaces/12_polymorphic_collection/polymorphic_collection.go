package polymorphic_collection

// TODO(human): Define Prioritizable interface
type Prioritizable interface {
	Priority() int
}

// TODO(human): Define Task struct
type Task struct {
	priority    int
	description string
}

// TODO(human): Define Message struct
type Message struct {
	priority int
	text     string
}

// TODO(human): Define PriorityQueue struct
type PriorityQueue struct {
	items []Prioritizable
}

// TODO(human): Implement Priority() for Task
func (t Task) Priority() int {
	return t.priority
}

// TODO(human): Implement Priority() for Message
func (m Message) Priority() int {
	return m.priority
}

// NewTask creates a new Task
func NewTask(priority int, description string) Task {
	// TODO(human): Implement
	return Task{
		priority:    priority,
		description: description,
	}
}

// NewMessage creates a new Message
func NewMessage(priority int, text string) Message {
	// TODO(human): Implement
	return Message{
		priority: priority,
		text:     text,
	}
}

// Push adds an item to the queue
func (pq *PriorityQueue) Push(item Prioritizable) {
	// TODO(human): Implement
	pq.items = append(pq.items, item)
}

// Pop removes and returns the highest priority item
func (pq *PriorityQueue) Pop() (Prioritizable, bool) {
	// TODO(human): Implement
	if len(pq.items) == 0 {
		return nil, false
	}

	var maxPrio int
	var pos int
	for i, k := range pq.items {
		if k.Priority() > maxPrio {
			maxPrio = k.Priority()
			pos = i
		}
	}

	res := pq.items[pos]

	if pos == len(pq.items)-1 {
		pq.items = pq.items[:pos]
	} else {
		a := pq.items[:pos]
		a = append(a, pq.items[pos+1:]...)
		pq.items = a
	}
	return res, true
}

// Peek returns the highest priority item without removing it
func (pq *PriorityQueue) Peek() (Prioritizable, bool) {
	// TODO(human): Implement

	if len(pq.items) == 0 {
		return nil, false
	}

	var res Prioritizable
	var maxPrio int
	for _, k := range pq.items {
		if k.Priority() > maxPrio {
			maxPrio = k.Priority()
			res = k
		}
	}
	return res, true
}

// Len returns the number of items in the queue
func (pq *PriorityQueue) Len() int {
	// TODO(human): Implement
	return len(pq.items)
}
