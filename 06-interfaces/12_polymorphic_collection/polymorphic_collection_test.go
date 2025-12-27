package polymorphic_collection

import "testing"

func TestTaskPriority(t *testing.T) {
	task := NewTask(5, "Test task")
	if task.Priority() != 5 {
		t.Errorf("Task.Priority() = %d, want 5", task.Priority())
	}
}

func TestMessagePriority(t *testing.T) {
	msg := NewMessage(3, "Test message")
	if msg.Priority() != 3 {
		t.Errorf("Message.Priority() = %d, want 3", msg.Priority())
	}
}

func TestPriorityQueuePush(t *testing.T) {
	pq := &PriorityQueue{}

	pq.Push(NewTask(5, "Task 1"))
	pq.Push(NewMessage(3, "Message 1"))

	if pq.Len() != 2 {
		t.Errorf("After 2 pushes, Len() = %d, want 2", pq.Len())
	}
}

func TestPriorityQueuePop(t *testing.T) {
	pq := &PriorityQueue{}

	pq.Push(NewTask(5, "Medium priority"))
	pq.Push(NewMessage(10, "High priority"))
	pq.Push(NewTask(3, "Low priority"))

	// Should pop in priority order: 10, 5, 3
	item1, ok := pq.Pop()
	if !ok {
		t.Fatal("Pop() should return true for non-empty queue")
	}
	if item1.Priority() != 10 {
		t.Errorf("First pop priority = %d, want 10", item1.Priority())
	}

	item2, ok := pq.Pop()
	if !ok {
		t.Fatal("Pop() should return true for non-empty queue")
	}
	if item2.Priority() != 5 {
		t.Errorf("Second pop priority = %d, want 5", item2.Priority())
	}

	item3, ok := pq.Pop()
	if !ok {
		t.Fatal("Pop() should return true for non-empty queue")
	}
	if item3.Priority() != 3 {
		t.Errorf("Third pop priority = %d, want 3", item3.Priority())
	}

	// Queue should be empty now
	if pq.Len() != 0 {
		t.Errorf("After 3 pops, Len() = %d, want 0", pq.Len())
	}
}

func TestPriorityQueuePopEmpty(t *testing.T) {
	pq := &PriorityQueue{}

	_, ok := pq.Pop()
	if ok {
		t.Error("Pop() on empty queue should return false")
	}
}

func TestPriorityQueuePeek(t *testing.T) {
	pq := &PriorityQueue{}

	pq.Push(NewTask(5, "Medium"))
	pq.Push(NewMessage(10, "High"))

	// Peek should return highest priority without removing
	item, ok := pq.Peek()
	if !ok {
		t.Fatal("Peek() should return true for non-empty queue")
	}
	if item.Priority() != 10 {
		t.Errorf("Peek() priority = %d, want 10", item.Priority())
	}

	// Length should be unchanged
	if pq.Len() != 2 {
		t.Errorf("After Peek, Len() = %d, want 2", pq.Len())
	}

	// Peek again - should return same item
	item2, _ := pq.Peek()
	if item2.Priority() != 10 {
		t.Errorf("Second Peek() priority = %d, want 10", item2.Priority())
	}
}

func TestPriorityQueuePeekEmpty(t *testing.T) {
	pq := &PriorityQueue{}

	_, ok := pq.Peek()
	if ok {
		t.Error("Peek() on empty queue should return false")
	}
}

func TestPriorityQueueLen(t *testing.T) {
	pq := &PriorityQueue{}

	if pq.Len() != 0 {
		t.Errorf("New queue Len() = %d, want 0", pq.Len())
	}

	pq.Push(NewTask(5, "Task"))
	if pq.Len() != 1 {
		t.Errorf("After 1 push, Len() = %d, want 1", pq.Len())
	}

	pq.Push(NewMessage(3, "Message"))
	if pq.Len() != 2 {
		t.Errorf("After 2 pushes, Len() = %d, want 2", pq.Len())
	}

	pq.Pop()
	if pq.Len() != 1 {
		t.Errorf("After 1 pop, Len() = %d, want 1", pq.Len())
	}
}

func TestPolymorphicCollection(t *testing.T) {
	// Demonstrate polymorphism - different types in same collection
	pq := &PriorityQueue{}

	task1 := NewTask(5, "Fix bug")
	msg1 := NewMessage(10, "Server down!")
	task2 := NewTask(3, "Update docs")
	msg2 := NewMessage(7, "Deploy ready")

	pq.Push(task1)
	pq.Push(msg1)
	pq.Push(task2)
	pq.Push(msg2)

	// All stored as Prioritizable interface
	// Should come out in priority order: 10, 7, 5, 3

	priorities := []int{}
	for pq.Len() > 0 {
		item, _ := pq.Pop()
		priorities = append(priorities, item.Priority())
	}

	want := []int{10, 7, 5, 3}
	for i, p := range want {
		if priorities[i] != p {
			t.Errorf("priorities[%d] = %d, want %d", i, priorities[i], p)
		}
	}
}

func TestImplementsPrioritizable(t *testing.T) {
	var _ Prioritizable = Task{}
	var _ Prioritizable = Message{}
	t.Log("✓ Both Task and Message implement Prioritizable interface")
}

func TestSamePriorityFIFO(t *testing.T) {
	pq := &PriorityQueue{}

	pq.Push(NewTask(5, "First"))
	pq.Push(NewMessage(5, "Second"))
	pq.Push(NewTask(5, "Third"))

	// When priorities are equal, should maintain FIFO order
	// (This is implementation-dependent; adjust test based on your implementation)
	for i := 0; i < 3; i++ {
		item, ok := pq.Pop()
		if !ok {
			t.Errorf("Pop %d failed", i)
		}
		if item.Priority() != 5 {
			t.Errorf("Pop %d priority = %d, want 5", i, item.Priority())
		}
	}
}
