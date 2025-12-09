package closure_syntax

import (
	"strings"
	"testing"
	"time"
)

// =============================================================================
// Exercise 1: StatefulOperation
// =============================================================================

func TestStatefulOperation(t *testing.T) {
	tests := []struct {
		name      string
		initial   int
		adds      []int
		wantFinal int
	}{
		{"start at zero", 0, []int{5, 3, 2}, 10},
		{"start at 10", 10, []int{5}, 15},
		{"negative adds", 100, []int{-30, -20}, 50},
		{"no adds", 42, []int{}, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			add, get := StatefulOperation(tt.initial)

			for _, v := range tt.adds {
				add(v)
			}

			if got := get(); got != tt.wantFinal {
				t.Errorf("getValue() = %d, want %d", got, tt.wantFinal)
			}
		})
	}
}

func TestStatefulOperation_Independence(t *testing.T) {
	add1, get1 := StatefulOperation(0)
	add2, get2 := StatefulOperation(100)

	add1(50)
	add1(50)
	add2(0) // Use add2 to verify it doesn't affect get1

	if got := get2(); got != 100 {
		t.Errorf("get2() = %d, want 100 (should be independent)", got)
	}
	if got := get1(); got != 100 {
		t.Errorf("get1() = %d, want 100", got)
	}
}

// =============================================================================
// Exercise 2: UndoRedo
// =============================================================================

func TestUndoRedo(t *testing.T) {
	do, undo, redo := UndoRedo()

	// Initial state - nothing to undo
	_, ok := undo()
	if ok {
		t.Error("undo should return false when nothing to undo")
	}

	// Do some operations
	do(10)
	do(20)
	do(30)

	// Undo
	val, ok := undo()
	if !ok || val != 20 {
		t.Errorf("undo() = %d, %v; want 20, true", val, ok)
	}

	val, ok = undo()
	if !ok || val != 10 {
		t.Errorf("undo() = %d, %v; want 10, true", val, ok)
	}

	// Redo
	val, ok = redo()
	if !ok || val != 20 {
		t.Errorf("redo() = %d, %v; want 20, true", val, ok)
	}

	// Undo again
	val, ok = undo()
	if !ok || val != 10 {
		t.Errorf("undo() = %d, %v; want 10, true", val, ok)
	}

	// Undo to empty
	val, ok = undo()
	if !ok || val != 0 {
		t.Errorf("undo() = %d, %v; want 0, true (initial state)", val, ok)
	}

	// Nothing more to undo
	_, ok = undo()
	if ok {
		t.Error("undo should return false when at initial state")
	}
}

func TestUndoRedo_RedoClearedOnNewDo(t *testing.T) {
	do, undo, redo := UndoRedo()

	do(10)
	do(20)
	undo() // Back to 10

	// New do should clear redo stack
	do(30)

	_, ok := redo()
	if ok {
		t.Error("redo should return false after new do() clears redo stack")
	}
}

// =============================================================================
// Exercise 3: Logger
// =============================================================================

func TestLogger(t *testing.T) {
	double := func(x int) int { return x * 2 }

	logged := Logger("TEST", double)

	// Should return correct value
	result := logged(5)
	if result != 10 {
		t.Errorf("logged(5) = %d, want 10", result)
	}

	result = logged(7)
	if result != 14 {
		t.Errorf("logged(7) = %d, want 14", result)
	}
}

func TestLogger_PreservesFunction(t *testing.T) {
	identity := func(x int) int { return x }
	addOne := func(x int) int { return x + 1 }

	loggedId := Logger("ID", identity)
	loggedAdd := Logger("ADD", addOne)

	if got := loggedId(42); got != 42 {
		t.Errorf("loggedId(42) = %d, want 42", got)
	}
	if got := loggedAdd(42); got != 43 {
		t.Errorf("loggedAdd(42) = %d, want 43", got)
	}
}

// =============================================================================
// Exercise 4: Compose
// =============================================================================

func TestCompose(t *testing.T) {
	addOne := func(x int) int { return x + 1 }
	double := func(x int) int { return x * 2 }
	square := func(x int) int { return x * x }

	tests := []struct {
		name  string
		f, g  func(int) int
		input int
		want  int
	}{
		{"double(addOne(x))", double, addOne, 5, 12}, // (5+1)*2
		{"addOne(double(x))", addOne, double, 5, 11}, // (5*2)+1
		{"square(addOne(x))", square, addOne, 3, 16}, // (3+1)^2
		{"addOne(square(x))", addOne, square, 3, 10}, // (3^2)+1
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			composed := Compose(tt.f, tt.g)
			if got := composed(tt.input); got != tt.want {
				t.Errorf("Compose(f,g)(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestCompose_Identity(t *testing.T) {
	identity := func(x int) int { return x }
	double := func(x int) int { return x * 2 }

	// Compose with identity should be equivalent to original
	composed := Compose(double, identity)
	if got := composed(5); got != 10 {
		t.Errorf("Compose(double, identity)(5) = %d, want 10", got)
	}

	composed = Compose(identity, double)
	if got := composed(5); got != 10 {
		t.Errorf("Compose(identity, double)(5) = %d, want 10", got)
	}
}

// =============================================================================
// Exercise 5: TimingMiddleware
// =============================================================================

func TestTimingMiddleware(t *testing.T) {
	// Function that takes known time
	slowFunc := func(x int) int {
		time.Sleep(50 * time.Millisecond)
		return x * 2
	}

	fastFunc := func(x int) int {
		return x * 2
	}

	// Threshold of 30ms - slowFunc should trigger, fastFunc should not
	middleware := TimingMiddleware(30 * time.Millisecond)

	wrappedSlow := middleware(slowFunc)
	wrappedFast := middleware(fastFunc)

	// Both should return correct values
	if got := wrappedSlow(5); got != 10 {
		t.Errorf("wrappedSlow(5) = %d, want 10", got)
	}

	if got := wrappedFast(5); got != 10 {
		t.Errorf("wrappedFast(5) = %d, want 10", got)
	}
}

func TestTimingMiddleware_ThreeLevels(t *testing.T) {
	// Test that middle level (the wrapper factory) works correctly
	mw1 := TimingMiddleware(10 * time.Millisecond)
	mw2 := TimingMiddleware(100 * time.Millisecond)

	fn := func(x int) int { return x }

	wrapped1 := mw1(fn)
	wrapped2 := mw2(fn)

	// Each should work independently
	if got := wrapped1(42); got != 42 {
		t.Errorf("wrapped1(42) = %d, want 42", got)
	}
	if got := wrapped2(42); got != 42 {
		t.Errorf("wrapped2(42) = %d, want 42", got)
	}
}

// =============================================================================
// Exercise 6: PartialApply
// =============================================================================

func TestPartialApply(t *testing.T) {
	add := func(a, b int) int { return a + b }
	multiply := func(a, b int) int { return a * b }
	subtract := func(a, b int) int { return a - b }

	tests := []struct {
		name   string
		fn     func(int, int) int
		first  int
		second int
		want   int
	}{
		{"add10", add, 10, 5, 15},
		{"add10 different", add, 10, 20, 30},
		{"triple", multiply, 3, 7, 21},
		{"subtract from 100", subtract, 100, 30, 70},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			partial := PartialApply(tt.fn, tt.first)
			if got := partial(tt.second); got != tt.want {
				t.Errorf("partial(%d) = %d, want %d", tt.second, got, tt.want)
			}
		})
	}
}

func TestPartialApply_Independence(t *testing.T) {
	add := func(a, b int) int { return a + b }

	add10 := PartialApply(add, 10)
	add100 := PartialApply(add, 100)

	if got := add10(5); got != 15 {
		t.Errorf("add10(5) = %d, want 15", got)
	}
	if got := add100(5); got != 105 {
		t.Errorf("add100(5) = %d, want 105", got)
	}
}

// =============================================================================
// Exercise 7: Pipeline
// =============================================================================

func TestPipeline(t *testing.T) {
	add, exec := Pipeline()

	add(func(x int) int { return x + 1 })
	add(func(x int) int { return x * 2 })
	add(func(x int) int { return x - 3 })

	// (5+1)*2-3 = 9
	if got := exec(5); got != 9 {
		t.Errorf("exec(5) = %d, want 9", got)
	}

	// (10+1)*2-3 = 19
	if got := exec(10); got != 19 {
		t.Errorf("exec(10) = %d, want 19", got)
	}
}

func TestPipeline_Empty(t *testing.T) {
	_, exec := Pipeline()

	// Empty pipeline should return input unchanged
	if got := exec(42); got != 42 {
		t.Errorf("empty pipeline exec(42) = %d, want 42", got)
	}
}

func TestPipeline_Independence(t *testing.T) {
	add1, exec1 := Pipeline()
	add2, exec2 := Pipeline()

	add1(func(x int) int { return x * 2 })
	add2(func(x int) int { return x + 100 })

	if got := exec1(5); got != 10 {
		t.Errorf("exec1(5) = %d, want 10", got)
	}
	if got := exec2(5); got != 105 {
		t.Errorf("exec2(5) = %d, want 105", got)
	}
}

// =============================================================================
// Exercise 8: Builder
// =============================================================================

func TestOptionBuilder(t *testing.T) {
	builder, getOpts := NewOptionBuilder()

	builder("verbose")("debug")("color")

	opts := getOpts()
	if len(opts) != 3 {
		t.Fatalf("getOpts() returned %d options, want 3", len(opts))
	}

	want := []string{"verbose", "debug", "color"}
	for i, w := range want {
		if opts[i] != w {
			t.Errorf("opts[%d] = %q, want %q", i, opts[i], w)
		}
	}
}

func TestOptionBuilder_Empty(t *testing.T) {
	_, getOpts := NewOptionBuilder()

	opts := getOpts()
	if len(opts) != 0 {
		t.Errorf("empty builder getOpts() returned %d options, want 0", len(opts))
	}
}

func TestOptionBuilder_Independence(t *testing.T) {
	b1, get1 := NewOptionBuilder()
	b2, get2 := NewOptionBuilder()

	b1("a")("b")("c")
	b2("x")

	if len(get1()) != 3 {
		t.Errorf("get1() returned %d options, want 3", len(get1()))
	}
	if len(get2()) != 1 {
		t.Errorf("get2() returned %d options, want 1", len(get2()))
	}
}

// =============================================================================
// Exercise 9: EventBus
// =============================================================================

func TestEventBus(t *testing.T) {
	pub, sub, closeBus := EventBus(10)

	ch := sub()

	pub("hello")
	pub("world")

	select {
	case msg := <-ch:
		if msg != "hello" {
			t.Errorf("first message = %q, want 'hello'", msg)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout waiting for first message")
	}

	select {
	case msg := <-ch:
		if msg != "world" {
			t.Errorf("second message = %q, want 'world'", msg)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout waiting for second message")
	}

	closeBus()
}

func TestEventBus_CloseStopsReceiving(t *testing.T) {
	pub, sub, closeBus := EventBus(10)

	ch := sub()
	pub("before close")
	closeBus()

	// Should be able to read buffered message
	select {
	case msg := <-ch:
		if msg != "before close" {
			t.Errorf("got %q, want 'before close'", msg)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout reading buffered message")
	}

	// Channel should be closed, read returns zero value
	select {
	case msg, ok := <-ch:
		if ok {
			t.Errorf("channel should be closed, got msg=%q", msg)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout - channel should be closed")
	}
}

func TestEventBus_SubscribeReturnsReceiveOnly(t *testing.T) {
	_, sub, closeBus := EventBus(10)
	defer closeBus()

	ch := sub()

	// This is a compile-time check - if ch allows sending, this would compile:
	// ch <- "test"  // Should NOT compile - ch is <-chan string
	_ = ch
}

// =============================================================================
// Exercise 10: MiddlewareChain
// =============================================================================

func TestChain(t *testing.T) {
	upper := func(next Handler) Handler {
		return func(s string) string {
			return next(strings.ToUpper(s))
		}
	}

	exclaim := func(next Handler) Handler {
		return func(s string) string {
			return next(s) + "!"
		}
	}

	bracket := func(next Handler) Handler {
		return func(s string) string {
			return "[" + next(s) + "]"
		}
	}

	base := func(s string) string { return s }

	tests := []struct {
		name        string
		middlewares []Middleware
		input       string
		want        string
	}{
		{"single - upper", []Middleware{upper}, "hello", "HELLO"},
		{"single - exclaim", []Middleware{exclaim}, "hello", "hello!"},
		{"upper then exclaim", []Middleware{upper, exclaim}, "hello", "HELLO!"},
		{"exclaim then upper", []Middleware{exclaim, upper}, "hello", "HELLO!"}, // upper runs on input first
		{"three middlewares", []Middleware{bracket, upper, exclaim}, "hi", "[HI!]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chained := Chain(tt.middlewares...)
			handler := chained(base)
			if got := handler(tt.input); got != tt.want {
				t.Errorf("handler(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestChain_Empty(t *testing.T) {
	base := func(s string) string { return s }

	chained := Chain() // No middleware
	handler := chained(base)

	if got := handler("hello"); got != "hello" {
		t.Errorf("empty chain handler('hello') = %q, want 'hello'", got)
	}
}

func TestChain_Order(t *testing.T) {
	// Track execution order
	var order []string

	first := func(next Handler) Handler {
		return func(s string) string {
			order = append(order, "first-before")
			result := next(s)
			order = append(order, "first-after")
			return result
		}
	}

	second := func(next Handler) Handler {
		return func(s string) string {
			order = append(order, "second-before")
			result := next(s)
			order = append(order, "second-after")
			return result
		}
	}

	base := func(s string) string {
		order = append(order, "base")
		return s
	}

	chained := Chain(first, second)
	handler := chained(base)
	handler("test")

	// Expected order: first wraps second wraps base
	// So: first-before, second-before, base, second-after, first-after
	want := []string{"first-before", "second-before", "base", "second-after", "first-after"}

	if len(order) != len(want) {
		t.Fatalf("order = %v, want %v", order, want)
	}
	for i, w := range want {
		if order[i] != w {
			t.Errorf("order[%d] = %q, want %q", i, order[i], w)
		}
	}
}

// =============================================================================
// Benchmarks
// =============================================================================

func BenchmarkCompose(b *testing.B) {
	addOne := func(x int) int { return x + 1 }
	double := func(x int) int { return x * 2 }
	composed := Compose(double, addOne)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		composed(i)
	}
}

func BenchmarkPipeline(b *testing.B) {
	add, exec := Pipeline()
	add(func(x int) int { return x + 1 })
	add(func(x int) int { return x * 2 })
	add(func(x int) int { return x - 3 })

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		exec(i)
	}
}

func BenchmarkMiddlewareChain(b *testing.B) {
	upper := func(next Handler) Handler {
		return func(s string) string {
			return next(strings.ToUpper(s))
		}
	}
	base := func(s string) string { return s }
	handler := Chain(upper, upper, upper)(base)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler("hello")
	}
}
