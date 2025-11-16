# Exercise 15: Generic Collections

## 🎯 Learning Goal
Master building reusable data structures using Go generics (Go 1.18+). Learn to implement type-safe collections (Stack, Queue, Set) that work with any type, understanding generic type parameters, type constraints, and when generics provide better solutions than interface{}.

## 📝 Problem Description

Before Go 1.18, building reusable data structures required using `interface{}` and type assertions, which sacrificed type safety. With generics, you can build collections that work with any type while maintaining compile-time type checking.

In this exercise, you'll implement four fundamental generic data structures:
1. **Stack** - Last-In-First-Out (LIFO) collection
2. **Queue** - First-In-First-Out (FIFO) collection
3. **Set** - Unique elements only (no duplicates)
4. **Pair** - Tuple holding two values of potentially different types

These are the building blocks for more complex algorithms and data structures.

## 🔧 Function Signatures

Implement these types and methods in `generic_collections.go`:

```go
// Stack - Last-In-First-Out collection
type Stack[T any] struct {
	// TODO(human): Add a slice field to store elements
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T]

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T)

// Pop removes and returns the top element (returns zero value and false if empty)
func (s *Stack[T]) Pop() (T, bool)

// Peek returns the top element without removing it (returns zero value and false if empty)
func (s *Stack[T]) Peek() (T, bool)

// IsEmpty returns true if the stack has no elements
func (s *Stack[T]) IsEmpty() bool

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int

// Queue - First-In-First-Out collection
type Queue[T any] struct {
	// TODO(human): Add a slice field to store elements
}

// NewQueue creates a new empty queue
func NewQueue[T any]() *Queue[T]

// Enqueue adds an element to the back of the queue
func (q *Queue[T]) Enqueue(value T)

// Dequeue removes and returns the front element (returns zero value and false if empty)
func (q *Queue[T]) Dequeue() (T, bool)

// Front returns the front element without removing it (returns zero value and false if empty)
func (q *Queue[T]) Front() (T, bool)

// IsEmpty returns true if the queue has no elements
func (q *Queue[T]) IsEmpty() bool

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() int

// Set - Collection of unique elements (requires comparable constraint)
type Set[T comparable] struct {
	// TODO(human): Add a map field to store elements (map[T]bool or map[T]struct{})
}

// NewSet creates a new empty set
func NewSet[T comparable]() *Set[T]

// Add adds an element to the set (idempotent - adding same element twice has no effect)
func (s *Set[T]) Add(value T)

// Remove removes an element from the set
func (s *Set[T]) Remove(value T)

// Contains checks if an element exists in the set
func (s *Set[T]) Contains(value T) bool

// Size returns the number of elements in the set
func (s *Set[T]) Size() int

// ToSlice returns all elements as a slice (order not guaranteed)
func (s *Set[T]) ToSlice() []T

// Pair - Tuple holding two values of potentially different types
type Pair[T1, T2 any] struct {
	First  T1
	Second T2
}

// NewPair creates a new pair with the given values
func NewPair[T1, T2 any](first T1, second T2) Pair[T1, T2]
```

## 💡 Examples

```go
// Stack of integers
stack := NewStack[int]()
stack.Push(1)
stack.Push(2)
stack.Push(3)
val, ok := stack.Pop()  // val=3, ok=true (LIFO)
val, ok = stack.Pop()   // val=2, ok=true

// Stack of strings
strStack := NewStack[string]()
strStack.Push("hello")
strStack.Push("world")
top, _ := strStack.Peek()  // top="world", stack unchanged

// Queue of integers
queue := NewQueue[int]()
queue.Enqueue(1)
queue.Enqueue(2)
queue.Enqueue(3)
val, ok := queue.Dequeue()  // val=1, ok=true (FIFO)
val, ok = queue.Dequeue()   // val=2, ok=true

// Set of strings
set := NewSet[string]()
set.Add("apple")
set.Add("banana")
set.Add("apple")  // No effect - already exists
set.Contains("apple")   // true
set.Size()              // 2 (duplicates ignored)

// Pair examples
pair1 := NewPair(42, "answer")           // Pair[int, string]
pair2 := NewPair("name", "Alice")        // Pair[string, string]
pair3 := NewPair(1.5, true)              // Pair[float64, bool]
```

## 📋 Instructions

1. **Stack Implementation:**
   - Use a slice to store elements: `items []T`
   - Push: append to the end
   - Pop: remove from the end (check if empty first!)
   - Peek: return last element without removing
   - Return zero value and false when empty

2. **Queue Implementation:**
   - Use a slice to store elements: `items []T`
   - Enqueue: append to the end
   - Dequeue: remove from the front (use slice reslicing: `items[1:]`)
   - Front: return first element without removing
   - Note: This simple implementation is O(n) for Dequeue (acceptable for learning)

3. **Set Implementation:**
   - Use a map to store elements: `map[T]bool` or `map[T]struct{}`
   - `map[T]struct{}` is more memory-efficient (empty struct uses 0 bytes)
   - Add: set map key to true
   - Contains: check if key exists in map
   - ToSlice: iterate over map keys

4. **Pair Implementation:**
   - Simple struct with two fields
   - NewPair is just a constructor function

## 🧪 Testing

Run tests with:
```bash
go test -v
```

Expected test count: ~40-50 tests across all types

## 🤔 Think About

1. **What does `any` mean in generics?**
   - `any` is an alias for `interface{}` - means "any type"
   - Used when we don't need any specific operations on the type

2. **What does `comparable` mean?**
   - A built-in constraint meaning the type can be compared with == and !=
   - Required for map keys and set elements
   - Includes: numbers, strings, booleans, pointers, arrays/structs of comparable types
   - Excludes: slices, maps, functions

3. **Why use generics instead of interface{}?**
   - Type safety: compile-time errors instead of runtime panics
   - No type assertions needed
   - Better IDE support and autocomplete
   - Self-documenting code

4. **When should you use generics?**
   - Data structures (containers, collections)
   - Algorithms that work on many types (sorting, filtering)
   - Utility functions (min, max, clamp)
   - NOT for everything - keep it simple when possible

## 💡 Hints

<details>
<summary>Hint 1: Generic type syntax</summary>

Generic types use square brackets for type parameters:

```go
// Single type parameter
type Stack[T any] struct {
    items []T
}

// Multiple type parameters
type Pair[T1, T2 any] struct {
    First  T1
    Second T2
}

// Type parameter with constraint
type Set[T comparable] struct {
    elements map[T]bool
}
```

Methods on generic types:
```go
func (s *Stack[T]) Push(value T) {
    s.items = append(s.items, value)
}
```
</details>

<details>
<summary>Hint 2: Zero values and the comma-ok pattern</summary>

When a function might fail, return both the value and a boolean:

```go
func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T  // Zero value of type T
        return zero, false
    }

    lastIndex := len(s.items) - 1
    value := s.items[lastIndex]
    s.items = s.items[:lastIndex]  // Remove last element
    return value, true
}
```

This pattern is common in Go (like map lookups: `val, ok := m[key]`)
</details>

<details>
<summary>Hint 3: Set implementation with map</summary>

Using `map[T]struct{}` for sets:

```go
type Set[T comparable] struct {
    elements map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
    return &Set[T]{
        elements: make(map[T]struct{}),
    }
}

func (s *Set[T]) Add(value T) {
    s.elements[value] = struct{}{}  // Empty struct as value
}

func (s *Set[T]) Contains(value T) bool {
    _, exists := s.elements[value]
    return exists
}
```

Why `struct{}`? It uses 0 bytes of memory - we only care about the keys!
</details>

<details>
<summary>Hint 4: Queue dequeue operation</summary>

Simple queue using slice reslicing:

```go
func (q *Queue[T]) Dequeue() (T, bool) {
    if len(q.items) == 0 {
        var zero T
        return zero, false
    }

    value := q.items[0]
    q.items = q.items[1:]  // Remove first element
    return value, true
}
```

Note: This creates garbage. Production queues use circular buffers or linked lists.
</details>

<details>
<summary>Full Solution</summary>

```go
package generic_collections

// Stack - Last-In-First-Out collection
type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: []T{}}
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Queue - First-In-First-Out collection
type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: []T{}}
}

func (q *Queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	value := q.items[0]
	q.items = q.items[1:]
	return value, true
}

func (q *Queue[T]) Front() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	return q.items[0], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

// Set - Collection of unique elements
type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

func (s *Set[T]) Add(value T) {
	s.elements[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.elements, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, exists := s.elements[value]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s.elements))
	for elem := range s.elements {
		result = append(result, elem)
	}
	return result
}

// Pair - Tuple holding two values
type Pair[T1, T2 any] struct {
	First  T1
	Second T2
}

func NewPair[T1, T2 any](first T1, second T2) Pair[T1, T2] {
	return Pair[T1, T2]{First: first, Second: second}
}
```
</details>

## 🎓 What This Teaches

- **Go generics** - Type parameters with `[T any]` and `[T comparable]` syntax
- **Type safety** - Compile-time guarantees without interface{} and type assertions
- **Data structures** - Stack (LIFO), Queue (FIFO), Set (unique elements)
- **Generic constraints** - Understanding `any` vs `comparable` vs custom constraints
- **Zero values** - Using `var zero T` to get the zero value of a generic type
- **Comma-ok pattern** - Returning (value, bool) for operations that might fail
- **Map efficiency** - Using `map[T]struct{}` for set implementations
- **Generic constructors** - Functions like `NewStack[T]()` that return generic types
- **Real-world patterns** - These are actual data structures used in production code

---

**Next Exercise:** `16_graph_basics` - Graph representations and basic traversals
