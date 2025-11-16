package collections

import (
	"testing"
)

// Stack Tests

func TestStackPush(t *testing.T) {
	var stack Stack

	stack.Push(1)
	if stack.Size() != 1 {
		t.Errorf("After Push(1), Size() = %d, want 1", stack.Size())
	}

	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Errorf("After 3 pushes, Size() = %d, want 3", stack.Size())
	}
}

func TestStackPop(t *testing.T) {
	tests := []struct {
		name      string
		pushes    []int
		pops      int
		wantVals  []int
		wantOks   []bool
		wantSize  int
	}{
		{
			"single element",
			[]int{42},
			1,
			[]int{42},
			[]bool{true},
			0,
		},
		{
			"multiple elements LIFO",
			[]int{1, 2, 3},
			3,
			[]int{3, 2, 1},
			[]bool{true, true, true},
			0,
		},
		{
			"pop from empty",
			[]int{},
			1,
			[]int{0},
			[]bool{false},
			0,
		},
		{
			"pop more than pushed",
			[]int{1, 2},
			3,
			[]int{2, 1, 0},
			[]bool{true, true, false},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stack Stack
			for _, val := range tt.pushes {
				stack.Push(val)
			}

			for i := 0; i < tt.pops; i++ {
				val, ok := stack.Pop()
				if val != tt.wantVals[i] {
					t.Errorf("Pop() #%d value = %d, want %d", i+1, val, tt.wantVals[i])
				}
				if ok != tt.wantOks[i] {
					t.Errorf("Pop() #%d ok = %v, want %v", i+1, ok, tt.wantOks[i])
				}
			}

			if stack.Size() != tt.wantSize {
				t.Errorf("Final Size() = %d, want %d", stack.Size(), tt.wantSize)
			}
		})
	}
}

func TestStackPeek(t *testing.T) {
	tests := []struct {
		name     string
		pushes   []int
		wantVal  int
		wantOk   bool
		wantSize int
	}{
		{"empty stack", []int{}, 0, false, 0},
		{"single element", []int{42}, 42, true, 1},
		{"multiple elements", []int{1, 2, 3}, 3, true, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stack Stack
			for _, val := range tt.pushes {
				stack.Push(val)
			}

			val, ok := stack.Peek()
			if val != tt.wantVal {
				t.Errorf("Peek() value = %d, want %d", val, tt.wantVal)
			}
			if ok != tt.wantOk {
				t.Errorf("Peek() ok = %v, want %v", ok, tt.wantOk)
			}
			if stack.Size() != tt.wantSize {
				t.Errorf("Size after Peek() = %d, want %d (should not change)", stack.Size(), tt.wantSize)
			}
		})
	}
}

func TestStackIsEmpty(t *testing.T) {
	var stack Stack

	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}

	stack.Push(1)
	if stack.IsEmpty() {
		t.Error("Stack with element should not be empty")
	}

	stack.Pop()
	if !stack.IsEmpty() {
		t.Error("Stack after popping all elements should be empty")
	}
}

func TestStackSize(t *testing.T) {
	var stack Stack

	if stack.Size() != 0 {
		t.Errorf("New stack Size() = %d, want 0", stack.Size())
	}

	for i := 1; i <= 5; i++ {
		stack.Push(i)
		if stack.Size() != i {
			t.Errorf("After %d pushes, Size() = %d, want %d", i, stack.Size(), i)
		}
	}

	for i := 4; i >= 0; i-- {
		stack.Pop()
		if stack.Size() != i {
			t.Errorf("After pop, Size() = %d, want %d", stack.Size(), i)
		}
	}
}

// Queue Tests

func TestQueueEnqueue(t *testing.T) {
	var queue Queue

	queue.Enqueue(1)
	if queue.Size() != 1 {
		t.Errorf("After Enqueue(1), Size() = %d, want 1", queue.Size())
	}

	queue.Enqueue(2)
	queue.Enqueue(3)
	if queue.Size() != 3 {
		t.Errorf("After 3 enqueues, Size() = %d, want 3", queue.Size())
	}
}

func TestQueueDequeue(t *testing.T) {
	tests := []struct {
		name      string
		enqueues  []int
		dequeues  int
		wantVals  []int
		wantOks   []bool
		wantSize  int
	}{
		{
			"single element",
			[]int{42},
			1,
			[]int{42},
			[]bool{true},
			0,
		},
		{
			"multiple elements FIFO",
			[]int{1, 2, 3},
			3,
			[]int{1, 2, 3},
			[]bool{true, true, true},
			0,
		},
		{
			"dequeue from empty",
			[]int{},
			1,
			[]int{0},
			[]bool{false},
			0,
		},
		{
			"dequeue more than enqueued",
			[]int{1, 2},
			3,
			[]int{1, 2, 0},
			[]bool{true, true, false},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var queue Queue
			for _, val := range tt.enqueues {
				queue.Enqueue(val)
			}

			for i := 0; i < tt.dequeues; i++ {
				val, ok := queue.Dequeue()
				if val != tt.wantVals[i] {
					t.Errorf("Dequeue() #%d value = %d, want %d", i+1, val, tt.wantVals[i])
				}
				if ok != tt.wantOks[i] {
					t.Errorf("Dequeue() #%d ok = %v, want %v", i+1, ok, tt.wantOks[i])
				}
			}

			if queue.Size() != tt.wantSize {
				t.Errorf("Final Size() = %d, want %d", queue.Size(), tt.wantSize)
			}
		})
	}
}

func TestQueueFront(t *testing.T) {
	tests := []struct {
		name     string
		enqueues []int
		wantVal  int
		wantOk   bool
		wantSize int
	}{
		{"empty queue", []int{}, 0, false, 0},
		{"single element", []int{42}, 42, true, 1},
		{"multiple elements", []int{1, 2, 3}, 1, true, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var queue Queue
			for _, val := range tt.enqueues {
				queue.Enqueue(val)
			}

			val, ok := queue.Front()
			if val != tt.wantVal {
				t.Errorf("Front() value = %d, want %d", val, tt.wantVal)
			}
			if ok != tt.wantOk {
				t.Errorf("Front() ok = %v, want %v", ok, tt.wantOk)
			}
			if queue.Size() != tt.wantSize {
				t.Errorf("Size after Front() = %d, want %d (should not change)", queue.Size(), tt.wantSize)
			}
		})
	}
}

func TestQueueIsEmpty(t *testing.T) {
	var queue Queue

	if !queue.IsEmpty() {
		t.Error("New queue should be empty")
	}

	queue.Enqueue(1)
	if queue.IsEmpty() {
		t.Error("Queue with element should not be empty")
	}

	queue.Dequeue()
	if !queue.IsEmpty() {
		t.Error("Queue after dequeueing all elements should be empty")
	}
}

func TestQueueSize(t *testing.T) {
	var queue Queue

	if queue.Size() != 0 {
		t.Errorf("New queue Size() = %d, want 0", queue.Size())
	}

	for i := 1; i <= 5; i++ {
		queue.Enqueue(i)
		if queue.Size() != i {
			t.Errorf("After %d enqueues, Size() = %d, want %d", i, queue.Size(), i)
		}
	}

	for i := 4; i >= 0; i-- {
		queue.Dequeue()
		if queue.Size() != i {
			t.Errorf("After dequeue, Size() = %d, want %d", queue.Size(), i)
		}
	}
}

func TestQueueFIFOOrder(t *testing.T) {
	var queue Queue

	// Enqueue 1-5
	for i := 1; i <= 5; i++ {
		queue.Enqueue(i)
	}

	// Dequeue should return 1-5 in order
	for i := 1; i <= 5; i++ {
		val, ok := queue.Dequeue()
		if !ok {
			t.Fatalf("Dequeue() #%d failed, want success", i)
		}
		if val != i {
			t.Errorf("Dequeue() #%d = %d, want %d (FIFO order)", i, val, i)
		}
	}
}

func TestStackLIFOOrder(t *testing.T) {
	var stack Stack

	// Push 1-5
	for i := 1; i <= 5; i++ {
		stack.Push(i)
	}

	// Pop should return 5-1 in reverse order
	for i := 5; i >= 1; i-- {
		val, ok := stack.Pop()
		if !ok {
			t.Fatalf("Pop() for value %d failed, want success", i)
		}
		if val != i {
			t.Errorf("Pop() = %d, want %d (LIFO order)", val, i)
		}
	}
}

func TestMixedStackOperations(t *testing.T) {
	var stack Stack

	stack.Push(1)
	stack.Push(2)
	val, _ := stack.Pop() // Remove 2
	if val != 2 {
		t.Errorf("Pop() = %d, want 2", val)
	}

	stack.Push(3)
	stack.Push(4)

	val, _ = stack.Peek() // Should be 4
	if val != 4 {
		t.Errorf("Peek() = %d, want 4", val)
	}

	if stack.Size() != 3 {
		t.Errorf("Size() = %d, want 3 (elements: 1, 3, 4)", stack.Size())
	}
}

func TestMixedQueueOperations(t *testing.T) {
	var queue Queue

	queue.Enqueue(1)
	queue.Enqueue(2)
	val, _ := queue.Dequeue() // Remove 1
	if val != 1 {
		t.Errorf("Dequeue() = %d, want 1", val)
	}

	queue.Enqueue(3)
	queue.Enqueue(4)

	val, _ = queue.Front() // Should be 2
	if val != 2 {
		t.Errorf("Front() = %d, want 2", val)
	}

	if queue.Size() != 3 {
		t.Errorf("Size() = %d, want 3 (elements: 2, 3, 4)", queue.Size())
	}
}
