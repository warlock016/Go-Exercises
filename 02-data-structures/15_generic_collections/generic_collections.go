package generic_collections

// TODO(human): Define Stack type for Last-In-First-Out collection
type Stack[T any] struct {
	Elements []T
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T] {
	// TODO(human): Implement
	return &Stack[T]{
		Elements: make([]T, 0),
	}
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	// TODO(human): Implement
	s.Elements = append(s.Elements, value)
}

// Pop removes and returns the top element (returns zero value and false if empty)
func (s *Stack[T]) Pop() (T, bool) {
	// TODO(human): Implement
	// 1. check if []Elements is empty
	// 1a. if empty, return 0, false
	// 1b. else copy it
	// 2. pop the stack []Elements slice
	// 3. return the element, true

	if len(s.Elements) == 0 {
		return *new(T), false
	}

	result := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]

	return result, true
}

// Peek returns the top element without removing it (returns zero value and false if empty)
func (s *Stack[T]) Peek() (T, bool) {
	// TODO(human): Implement
	if len(s.Elements) == 0 {
		return *new(T), false
	}
	return s.Elements[len(s.Elements)-1], true
}

// IsEmpty returns true if the stack has no elements
func (s *Stack[T]) IsEmpty() bool {
	// TODO(human): Implement
	if len(s.Elements) == 0 {
		return true
	}
	return false
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	// TODO(human): Implement
	return len(s.Elements)
}

// TODO(human): Define Queue type for First-In-First-Out collection
type Queue[T any] struct {
	Elements []T
}

// NewQueue creates a new empty queue
func NewQueue[T any]() *Queue[T] {
	// TODO(human): Implement
	return &Queue[T]{}
}

// Enqueue adds an element to the back of the queue
func (q *Queue[T]) Enqueue(value T) {
	// TODO(human): Implement
	q.Elements = append(q.Elements, value)
}

// Dequeue removes and returns the front element (returns zero value and false if empty)
func (q *Queue[T]) Dequeue() (T, bool) {
	// TODO(human): Implement

	if len(q.Elements) == 0 {
		return *new(T), false
	}

	removed := q.Elements[0]
	q.Elements = q.Elements[1:]

	return removed, true
}

// Front returns the front element without removing it (returns zero value and false if empty)
func (q *Queue[T]) Front() (T, bool) {
	// TODO(human): Implement
	if len(q.Elements) == 0 {
		return *new(T), false
	}
	return q.Elements[0], true
}

// IsEmpty returns true if the queue has no elements
func (q *Queue[T]) IsEmpty() bool {
	// TODO(human): Implement
	return len(q.Elements) == 0
}

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() int {
	// TODO(human): Implement
	return len(q.Elements)
}

// TODO(human): Define Set type for collection of unique elements (requires comparable constraint)
type Set[T comparable] struct {
	Elements map[T]struct{}
}

// NewSet creates a new empty set
func NewSet[T comparable]() *Set[T] {
	// TODO(human): Implement
	return &Set[T]{
		Elements: make(map[T]struct{}),
	}
}

// Add adds an element to the set (idempotent - adding same element twice has no effect)
func (s *Set[T]) Add(value T) {
	// TODO(human): Implement

	s.Elements[value] = struct{}{}
}

// Remove removes an element from the set
func (s *Set[T]) Remove(value T) {
	// TODO(human): Implement

	copy := s.Elements
	delete(copy, value)
	s.Elements = copy
}

// Contains checks if an element exists in the set
func (s *Set[T]) Contains(value T) bool {
	// TODO(human): Implement
	if _, ok := s.Elements[value]; ok {
		return true
	}
	return false
}

// Size returns the number of elements in the set
func (s *Set[T]) Size() int {
	// TODO(human): Implement
	return len(s.Elements)
}

// ToSlice returns all elements as a slice (order not guaranteed)
func (s *Set[T]) ToSlice() []T {
	// TODO(human): Implement

	result := make([]T, 0, len(s.Elements))

	for i := range s.Elements {
		result = append(result, i)
	}
	return result
}

// TODO(human): Define Pair struct as a tuple holding two values of potentially different types
type Pair[T1, T2 any] struct {
	First  T1
	Second T2
}

// NewPair creates a new pair with the given values
func NewPair[T1, T2 any](first T1, second T2) Pair[T1, T2] {
	return Pair[T1, T2]{
		First:  first,
		Second: second,
	}
}
