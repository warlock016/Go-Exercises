# Exercise 08: Collections

## 🎯 Learning Goal
Implement classic data structures (Stack and Queue) using Go slices. Understand LIFO vs FIFO behavior, receiver methods, and pointer receivers for mutation.

## 📝 Problem Description

Stacks and Queues are fundamental data structures used everywhere in programming. A Stack is Last-In-First-Out (LIFO) - like a stack of plates. A Queue is First-In-First-Out (FIFO) - like a line at the store.

**Key Concepts:**
- Methods with receivers (not standalone functions!)
- Pointer receivers for mutation
- Value receivers for read-only operations
- Multiple return values for optional results
- Implementing abstract data types with slices

## 🔧 Type Definitions & Methods

```go
// Stack implements a LIFO (Last In First Out) data structure
type Stack struct {
	items []int
}

// Push adds a value to the top of the stack
func (s *Stack) Push(value int)

// Pop removes and returns the top value (returns value, true if success)
func (s *Stack) Pop() (int, bool)

// Peek returns the top value without removing it (returns value, true if success)
func (s *Stack) Peek() (int, bool)

// IsEmpty returns true if stack has no elements
func (s *Stack) IsEmpty() bool

// Size returns the number of elements in the stack
func (s *Stack) Size() int

// Queue implements a FIFO (First In First Out) data structure
type Queue struct {
	items []int
}

// Enqueue adds a value to the back of the queue
func (q *Queue) Enqueue(value int)

// Dequeue removes and returns the front value (returns value, true if success)
func (q *Queue) Dequeue() (int, bool)

// Front returns the front value without removing it (returns value, true if success)
func (q *Queue) Front() (int, bool)

// IsEmpty returns true if queue has no elements
func (q *Queue) IsEmpty() bool

// Size returns the number of elements in the queue
func (q *Queue) Size() int
```

## 💡 Examples

```go
// Stack usage (LIFO)
var stack Stack
stack.Push(1)
stack.Push(2)
stack.Push(3)

val, ok := stack.Pop()  // val=3, ok=true (last in, first out)
val, ok = stack.Pop()   // val=2, ok=true
val, ok = stack.Peek()  // val=1, ok=true (doesn't remove)
val, ok = stack.Pop()   // val=1, ok=true
val, ok = stack.Pop()   // val=0, ok=false (empty)

// Queue usage (FIFO)
var queue Queue
queue.Enqueue(1)
queue.Enqueue(2)
queue.Enqueue(3)

val, ok := queue.Dequeue()  // val=1, ok=true (first in, first out)
val, ok = queue.Dequeue()   // val=2, ok=true
val, ok = queue.Front()     // val=3, ok=true (doesn't remove)
val, ok = queue.Dequeue()   // val=3, ok=true
val, ok = queue.Dequeue()   // val=0, ok=false (empty)
```

## 📋 Instructions

### Stack Implementation
1. **Push:** Append to `s.items` slice
2. **Pop:** Check if empty, return last element and remove it using slice expression
3. **Peek:** Check if empty, return last element without removing
4. **IsEmpty:** Return `len(s.items) == 0`
5. **Size:** Return `len(s.items)`

### Queue Implementation
1. **Enqueue:** Append to `q.items` slice (add to back)
2. **Dequeue:** Check if empty, return first element and remove it using slice expression
3. **Front:** Check if empty, return first element without removing
4. **IsEmpty:** Return `len(q.items) == 0`
5. **Size:** Return `len(q.items)`

## 🧪 Testing

```bash
go test -v
```

Expected test count: ~40+ tests

## 🤔 Think About

1. **Why use pointer receivers for Push/Pop/Enqueue/Dequeue?**
   - These methods modify the struct, so need pointer to change original

2. **Why use value receivers for IsEmpty/Size?**
   - These are read-only operations, don't need to modify

3. **Why return (int, bool) instead of just int?**
   - Need to distinguish "popped 0" from "stack empty"

4. **How do you remove the first element from a slice?**
   - Use slice expression: `items[1:]` (everything except first)

5. **How do you remove the last element from a slice?**
   - Use slice expression: `items[:len(items)-1]` (everything except last)

6. **What's the performance difference between Stack and Queue?**
   - Stack operations are O(1), Queue's Dequeue is O(n) due to slice shifting!

## 💡 Hints

<details>
<summary>Hint 1: Stack Push and Pop</summary>

```go
func (s *Stack) Push(value int) {
    s.items = append(s.items, value)
}

func (s *Stack) Pop() (int, bool) {
    if len(s.items) == 0 {
        return 0, false
    }
    lastIndex := len(s.items) - 1
    value := s.items[lastIndex]
    s.items = s.items[:lastIndex]  // Remove last element
    return value, true
}
```
</details>

<details>
<summary>Hint 2: Stack Peek</summary>

```go
func (s *Stack) Peek() (int, bool) {
    if len(s.items) == 0 {
        return 0, false
    }
    return s.items[len(s.items)-1], true
}
```
</details>

<details>
<summary>Hint 3: Queue Enqueue and Dequeue</summary>

```go
func (q *Queue) Enqueue(value int) {
    q.items = append(q.items, value)
}

func (q *Queue) Dequeue() (int, bool) {
    if len(q.items) == 0 {
        return 0, false
    }
    value := q.items[0]
    q.items = q.items[1:]  // Remove first element
    return value, true
}
```
</details>

<details>
<summary>Hint 4: IsEmpty and Size (works for both)</summary>

```go
// For Stack
func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}

func (s *Stack) Size() int {
    return len(s.items)
}

// For Queue - same logic
func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}

func (q *Queue) Size() int {
    return len(q.items)
}
```
</details>

<details>
<summary>Complete Solution</summary>

```go
type Stack struct {
    items []int
}

func (s *Stack) Push(value int) {
    s.items = append(s.items, value)
}

func (s *Stack) Pop() (int, bool) {
    if len(s.items) == 0 {
        return 0, false
    }
    lastIndex := len(s.items) - 1
    value := s.items[lastIndex]
    s.items = s.items[:lastIndex]
    return value, true
}

func (s *Stack) Peek() (int, bool) {
    if len(s.items) == 0 {
        return 0, false
    }
    return s.items[len(s.items)-1], true
}

func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}

func (s *Stack) Size() int {
    return len(s.items)
}

type Queue struct {
    items []int
}

func (q *Queue) Enqueue(value int) {
    q.items = append(q.items, value)
}

func (q *Queue) Dequeue() (int, bool) {
    if len(q.items) == 0 {
        return 0, false
    }
    value := q.items[0]
    q.items = q.items[1:]
    return value, true
}

func (q *Queue) Front() (int, bool) {
    if len(q.items) == 0 {
        return 0, false
    }
    return q.items[0], true
}

func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}

func (q *Queue) Size() int {
    return len(q.items)
}
```
</details>

## 🎓 What This Teaches

- **Methods with receivers** - Attaching functions to types
- **Pointer vs value receivers** - When to use each
- **LIFO vs FIFO** - Understanding different access patterns
- **Optional return values** - Using (value, bool) pattern
- **Slice manipulation** - Adding and removing elements efficiently
- **Abstract data types** - Implementing interfaces with slices
- **Encapsulation** - Hiding implementation details behind methods
- **Performance awareness** - Understanding O(1) vs O(n) operations

---

**Module Complete!** You've mastered Tier 2 application patterns. Next: Tier 3 Integration exercises.
