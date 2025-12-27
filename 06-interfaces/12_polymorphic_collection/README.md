# Exercise 12: Polymorphic Collection

**Concept:** Slices of interfaces for runtime polymorphism
**Difficulty:** Medium-Hard
**Estimated Time:** 60 minutes

## Learning Goal

Learn to use interface slices to create collections of different types that share common behavior, enabling runtime polymorphism.

## The Problem

You need a collection of different types with shared behavior:

```go
// Can't do this - different types
items := []???{
    Task{Priority: 5},
    Message{Priority: 3},
    Alert{Priority: 10},
}

// Solution - slice of interface!
type Prioritizable interface {
    Priority() int
}

items := []Prioritizable{
    Task{priority: 5},
    Message{priority: 3},
    Alert{priority: 10},
}
```

## Your Task

Create a priority queue that works with any prioritizable items.

### Interface

**Prioritizable** - things with priority
- `Priority() int` - returns priority (higher = more important)

### Types

1. **Task** - `priority int`, `description string`
2. **Message** - `priority int`, `text string`

### Type

**PriorityQueue** - manages prioritizable items

**Fields:**
- `items []Prioritizable`

**Methods:**
- `Push(item Prioritizable)` - adds item
- `Pop() (Prioritizable, bool)` - removes and returns highest priority item
- `Peek() (Prioritizable, bool)` - returns highest priority without removing
- `Len() int` - returns number of items

## Function Signatures

```go
type Prioritizable interface {
    Priority() int
}

type Task struct {
    priority    int
    description string
}

type Message struct {
    priority int
    text     string
}

type PriorityQueue struct {
    items []Prioritizable
}

func (t Task) Priority() int
func (m Message) Priority() int

func (pq *PriorityQueue) Push(item Prioritizable)
func (pq *PriorityQueue) Pop() (Prioritizable, bool)
func (pq *PriorityQueue) Peek() (Prioritizable, bool)
func (pq *PriorityQueue) Len() int

func NewTask(priority int, description string) Task
func NewMessage(priority int, text string) Message
```

## Examples

```go
pq := &PriorityQueue{}

pq.Push(NewTask(5, "Medium task"))
pq.Push(NewMessage(10, "Urgent message"))
pq.Push(NewTask(3, "Low task"))

item, _ := pq.Pop()  // Returns message (priority 10)
item, _ := pq.Pop()  // Returns task (priority 5)
item, _ := pq.Pop()  // Returns task (priority 3)
```

## What This Teaches

- Interface slices
- Runtime polymorphism
- Priority queue pattern
- Mixed-type collections

---

**Congratulations! You've completed the Interfaces module!**
