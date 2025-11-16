package generic_collections

import (
	"testing"
)

// Stack Tests

func TestStackPushPop(t *testing.T) {
	stack := NewStack[int]()

	// Push elements
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Pop elements (LIFO order)
	val, ok := stack.Pop()
	if !ok || val != 3 {
		t.Errorf("Pop() = (%d, %v), want (3, true)", val, ok)
	}

	val, ok = stack.Pop()
	if !ok || val != 2 {
		t.Errorf("Pop() = (%d, %v), want (2, true)", val, ok)
	}

	val, ok = stack.Pop()
	if !ok || val != 1 {
		t.Errorf("Pop() = (%d, %v), want (1, true)", val, ok)
	}

	// Pop from empty stack
	val, ok = stack.Pop()
	if ok || val != 0 {
		t.Errorf("Pop() on empty stack = (%d, %v), want (0, false)", val, ok)
	}
}

func TestStackPeek(t *testing.T) {
	stack := NewStack[string]()

	// Peek empty stack
	val, ok := stack.Peek()
	if ok || val != "" {
		t.Errorf("Peek() on empty stack = (%q, %v), want (\"\", false)", val, ok)
	}

	stack.Push("hello")
	stack.Push("world")

	// Peek should return top without removing
	val, ok = stack.Peek()
	if !ok || val != "world" {
		t.Errorf("Peek() = (%q, %v), want (\"world\", true)", val, ok)
	}

	// Peek again - should still be there
	val, ok = stack.Peek()
	if !ok || val != "world" {
		t.Errorf("Second Peek() = (%q, %v), want (\"world\", true)", val, ok)
	}

	// Size should still be 2
	if stack.Size() != 2 {
		t.Errorf("Size() = %d, want 2", stack.Size())
	}
}

func TestStackIsEmptyAndSize(t *testing.T) {
	stack := NewStack[int]()

	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}
	if stack.Size() != 0 {
		t.Errorf("New stack Size() = %d, want 0", stack.Size())
	}

	stack.Push(1)
	if stack.IsEmpty() {
		t.Error("Stack with element should not be empty")
	}
	if stack.Size() != 1 {
		t.Errorf("Stack Size() = %d, want 1", stack.Size())
	}

	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Errorf("Stack Size() = %d, want 3", stack.Size())
	}

	stack.Pop()
	stack.Pop()
	stack.Pop()
	if !stack.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}
}

// Queue Tests

func TestQueueEnqueueDequeue(t *testing.T) {
	queue := NewQueue[int]()

	// Enqueue elements
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// Dequeue elements (FIFO order)
	val, ok := queue.Dequeue()
	if !ok || val != 1 {
		t.Errorf("Dequeue() = (%d, %v), want (1, true)", val, ok)
	}

	val, ok = queue.Dequeue()
	if !ok || val != 2 {
		t.Errorf("Dequeue() = (%d, %v), want (2, true)", val, ok)
	}

	val, ok = queue.Dequeue()
	if !ok || val != 3 {
		t.Errorf("Dequeue() = (%d, %v), want (3, true)", val, ok)
	}

	// Dequeue from empty queue
	val, ok = queue.Dequeue()
	if ok || val != 0 {
		t.Errorf("Dequeue() on empty queue = (%d, %v), want (0, false)", val, ok)
	}
}

func TestQueueFront(t *testing.T) {
	queue := NewQueue[string]()

	// Front on empty queue
	val, ok := queue.Front()
	if ok || val != "" {
		t.Errorf("Front() on empty queue = (%q, %v), want (\"\", false)", val, ok)
	}

	queue.Enqueue("first")
	queue.Enqueue("second")

	// Front should return first without removing
	val, ok = queue.Front()
	if !ok || val != "first" {
		t.Errorf("Front() = (%q, %v), want (\"first\", true)", val, ok)
	}

	// Front again - should still be there
	val, ok = queue.Front()
	if !ok || val != "first" {
		t.Errorf("Second Front() = (%q, %v), want (\"first\", true)", val, ok)
	}

	// Size should still be 2
	if queue.Size() != 2 {
		t.Errorf("Size() = %d, want 2", queue.Size())
	}
}

func TestQueueIsEmptyAndSize(t *testing.T) {
	queue := NewQueue[int]()

	if !queue.IsEmpty() {
		t.Error("New queue should be empty")
	}
	if queue.Size() != 0 {
		t.Errorf("New queue Size() = %d, want 0", queue.Size())
	}

	queue.Enqueue(1)
	if queue.IsEmpty() {
		t.Error("Queue with element should not be empty")
	}
	if queue.Size() != 1 {
		t.Errorf("Queue Size() = %d, want 1", queue.Size())
	}

	queue.Enqueue(2)
	queue.Enqueue(3)
	if queue.Size() != 3 {
		t.Errorf("Queue Size() = %d, want 3", queue.Size())
	}

	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	if !queue.IsEmpty() {
		t.Error("Queue should be empty after dequeueing all elements")
	}
}

// Set Tests

func TestSetAddAndContains(t *testing.T) {
	set := NewSet[string]()

	// Initially empty
	if set.Contains("apple") {
		t.Error("Empty set should not contain 'apple'")
	}

	// Add elements
	set.Add("apple")
	if !set.Contains("apple") {
		t.Error("Set should contain 'apple' after adding")
	}

	set.Add("banana")
	set.Add("cherry")
	if !set.Contains("banana") || !set.Contains("cherry") {
		t.Error("Set should contain all added elements")
	}

	// Add duplicate - should have no effect
	set.Add("apple")
	if set.Size() != 3 {
		t.Errorf("Set Size() = %d, want 3 (duplicates should be ignored)", set.Size())
	}
}

func TestSetRemove(t *testing.T) {
	set := NewSet[int]()

	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.Remove(2)
	if set.Contains(2) {
		t.Error("Set should not contain 2 after removal")
	}
	if set.Size() != 2 {
		t.Errorf("Set Size() = %d, want 2", set.Size())
	}

	// Remove non-existent element - should not error
	set.Remove(99)
	if set.Size() != 2 {
		t.Errorf("Set Size() = %d, want 2 (removing non-existent element)", set.Size())
	}
}

func TestSetSize(t *testing.T) {
	set := NewSet[string]()

	if set.Size() != 0 {
		t.Errorf("New set Size() = %d, want 0", set.Size())
	}

	set.Add("a")
	set.Add("b")
	set.Add("c")
	if set.Size() != 3 {
		t.Errorf("Set Size() = %d, want 3", set.Size())
	}

	set.Add("a") // Duplicate
	if set.Size() != 3 {
		t.Errorf("Set Size() = %d, want 3 (after adding duplicate)", set.Size())
	}
}

func TestSetToSlice(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	slice := set.ToSlice()
	if len(slice) != 3 {
		t.Fatalf("ToSlice() length = %d, want 3", len(slice))
	}

	// Check all elements are present (order not guaranteed)
	found := make(map[int]bool)
	for _, val := range slice {
		found[val] = true
	}

	for _, expected := range []int{1, 2, 3} {
		if !found[expected] {
			t.Errorf("ToSlice() missing element %d", expected)
		}
	}
}

func TestSetWithDifferentTypes(t *testing.T) {
	// Test with int
	intSet := NewSet[int]()
	intSet.Add(42)
	if !intSet.Contains(42) {
		t.Error("Int set should contain 42")
	}

	// Test with string
	strSet := NewSet[string]()
	strSet.Add("hello")
	if !strSet.Contains("hello") {
		t.Error("String set should contain 'hello'")
	}

	// Test with float64
	floatSet := NewSet[float64]()
	floatSet.Add(3.14)
	if !floatSet.Contains(3.14) {
		t.Error("Float set should contain 3.14")
	}
}

// Pair Tests

func TestPairCreation(t *testing.T) {
	// Int and String
	pair1 := NewPair(42, "answer")
	if pair1.First != 42 || pair1.Second != "answer" {
		t.Errorf("NewPair(42, \"answer\") = {%v, %v}, want {42, \"answer\"}", pair1.First, pair1.Second)
	}

	// String and String
	pair2 := NewPair("key", "value")
	if pair2.First != "key" || pair2.Second != "value" {
		t.Errorf("NewPair(\"key\", \"value\") = {%v, %v}, want {\"key\", \"value\"}", pair2.First, pair2.Second)
	}

	// Float and Bool
	pair3 := NewPair(3.14, true)
	if pair3.First != 3.14 || pair3.Second != true {
		t.Errorf("NewPair(3.14, true) = {%v, %v}, want {3.14, true}", pair3.First, pair3.Second)
	}
}

func TestPairStructFields(t *testing.T) {
	pair := NewPair(100, "hundred")

	// Access fields directly
	if pair.First != 100 {
		t.Errorf("pair.First = %d, want 100", pair.First)
	}
	if pair.Second != "hundred" {
		t.Errorf("pair.Second = %q, want \"hundred\"", pair.Second)
	}

	// Modify fields
	pair.First = 200
	pair.Second = "two hundred"

	if pair.First != 200 || pair.Second != "two hundred" {
		t.Errorf("After modification, pair = {%d, %q}, want {200, \"two hundred\"}", pair.First, pair.Second)
	}
}

// Integration Tests

func TestStackWithDifferentTypes(t *testing.T) {
	// String stack
	strStack := NewStack[string]()
	strStack.Push("hello")
	strStack.Push("world")
	val, _ := strStack.Pop()
	if val != "world" {
		t.Errorf("String stack Pop() = %q, want \"world\"", val)
	}

	// Float stack
	floatStack := NewStack[float64]()
	floatStack.Push(1.5)
	floatStack.Push(2.5)
	fval, _ := floatStack.Pop()
	if fval != 2.5 {
		t.Errorf("Float stack Pop() = %f, want 2.5", fval)
	}

	// Bool stack
	boolStack := NewStack[bool]()
	boolStack.Push(true)
	boolStack.Push(false)
	bval, _ := boolStack.Pop()
	if bval != false {
		t.Errorf("Bool stack Pop() = %v, want false", bval)
	}
}

func TestQueueWithDifferentTypes(t *testing.T) {
	// String queue
	strQueue := NewQueue[string]()
	strQueue.Enqueue("first")
	strQueue.Enqueue("second")
	val, _ := strQueue.Dequeue()
	if val != "first" {
		t.Errorf("String queue Dequeue() = %q, want \"first\"", val)
	}

	// Float queue
	floatQueue := NewQueue[float64]()
	floatQueue.Enqueue(1.1)
	floatQueue.Enqueue(2.2)
	fval, _ := floatQueue.Dequeue()
	if fval != 1.1 {
		t.Errorf("Float queue Dequeue() = %f, want 1.1", fval)
	}
}
