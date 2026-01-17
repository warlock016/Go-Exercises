package channel_directions

import (
	"slices"
	"sort"
	"sync"
	"testing"
	"time"
)

// =============================================================================
// Channel Directions Test Suite
// =============================================================================
//
// Write tests for each function following these guidelines.
// Remember to run tests with: go test -race -count=10 ./...
//
// =============================================================================

func sliceToChan(slice []int) <-chan int {

	ch := make(chan int)
	go func() {
		for _, v := range slice {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func chanToSlice(t *testing.T, ch <-chan int) []int {
	t.Helper()
	const timeout time.Duration = 5 * time.Second
	slice := []int{}

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return slice
			}
			slice = append(slice, v)
		case <-time.After(timeout):
			t.Fatal("chanToSlice: timeout - channel never closed")
		}
	}
}

func sortedEqual(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	slices.Sort(a)
	slices.Sort(b)

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// =============================================================================
// TestNumbers
// =============================================================================
//
// Test cases to implement:
//
// 1. "emits correct sequence"
//    - Call Numbers(5)
//    - Collect all values from the channel
//    - Verify you receive exactly [1, 2, 3, 4, 5] in order
//
// 2. "closes after completion"
//    - Call Numbers(3)
//    - Drain the channel completely
//    - Verify the channel is closed (receive returns ok=false)
//
// 3. "handles zero"
//    - Call Numbers(0)
//    - Verify the channel closes immediately with no values
//    - Use a timeout to avoid hanging if implementation is wrong
//
// 4. "handles negative"
//    - Call Numbers(-5)
//    - Verify the channel closes immediately with no values
//
// Edge cases to consider:
//    - What if n is very large? (don't test 1 million, but think about it)
//    - Is the function non-blocking? (returns immediately before emitting)

// TODO(human): Implement TestNumbers
func TestGenerator(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  []int
	}{
		{
			name:  "emits correct sequence",
			input: 5,
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "closes after completion",
			input: 3,
			want:  []int{1, 2, 3},
		},
		{
			name:  "handles zero",
			input: 0,
			want:  []int{},
		},
		{
			name:  "handles negative",
			input: -5,
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			resCh := Generator(tt.input)
			got := chanToSlice(t, resCh)

			if len(got) != len(tt.want) {
				t.Fatalf("Unexpected slice length: got %d, want %d", len(got), len(tt.want))
			}

			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("unexpected element: got %d, want %d", v, tt.want[i])
				}
			}

		})
	}
}

// =============================================================================
// TestSquare
// =============================================================================
//
// Test cases to implement:
//
// 1. "squares all values"
//    - Create input channel with values [1, 2, 3, 4]
//    - Pass to Square()
//    - Verify output is [1, 4, 9, 16]
//
// 2. "handles empty channel"
//    - Create and immediately close an empty channel
//    - Pass to Square()
//    - Verify output channel closes with no values
//
// 3. "handles negative numbers"
//    - Input: [-2, -1, 0, 1, 2]
//    - Output: [4, 1, 0, 1, 4]
//
// 4. "closes output when input closes"
//    - Verify that ranging over output terminates
//
// Testing tip:
//    - Create a helper to convert a slice to a channel for easier testing
//    - Example: func sliceToChan(vals []int) <-chan int

// TODO(human): Implement TestSquare
func TestSquare(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "all values",
			input: []int{1, 2, 3, 4},
			want:  []int{1, 4, 9, 16},
		},
		{
			name:  "empty channel",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "negative number",
			input: []int{-2, -1, 0, 1, 2},
			want:  []int{4, 1, 0, 1, 4},
		},
		{
			name:  "closed input",
			input: []int{0, 1, 2, 3, 4, 5},
			want:  []int{0, 1, 4, 9, 16, 25},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inCh := sliceToChan(tt.input)
			outCh := Square(inCh)

			got := chanToSlice(t, outCh)

			if len(got) != len(tt.input) {
				t.Errorf("unexpected slice length: got %d, want %d", len(got), len(tt.input))
			}

			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("unexpected value: got %d, want %d", v, tt.input[i])
				}
			}
		})
	}
}

// =============================================================================
// TestFilter
// =============================================================================
//
// Test cases to implement:
//
// 1. "filters even numbers"
//    - Input: [1, 2, 3, 4, 5, 6]
//    - Predicate: n % 2 == 0
//    - Output: [2, 4, 6]
//
// 2. "filters with no matches"
//    - Input: [1, 3, 5, 7]
//    - Predicate: n % 2 == 0
//    - Output: [] (empty, but channel still closes)
//
// 3. "filters with all matches"
//    - Input: [2, 4, 6, 8]
//    - Predicate: n % 2 == 0
//    - Output: [2, 4, 6, 8]
//
// 4. "handles empty input"
//    - Empty input channel
//    - Any predicate
//    - Output: [] (channel closes immediately)
//
// 5. "custom predicate - greater than threshold"
//    - Input: [1, 5, 10, 15, 20]
//    - Predicate: n > 10
//    - Output: [15, 20]

// TODO(human): Implement TestFilter
func TestFilter(t *testing.T) {
	var (
		isEven    = func(n int) bool { return n%2 == 0 }
		isGreater = func(n int) bool { return n > 10 }
	)

	tests := []struct {
		name  string
		input []int
		fn    func(int) bool
		want  []int
	}{
		{
			name:  "even numbers",
			input: []int{0, 1, 2, 3, 4, 5},
			fn:    isEven,
			want:  []int{0, 2, 4},
		},
		{
			name:  "no matches",
			input: []int{1, 3, 5, 7, 9},
			fn:    isEven,
			want:  []int{},
		},
		{
			name:  "all matches",
			input: []int{0, 2, 4, 6, 8},
			fn:    isEven,
			want:  []int{0, 2, 4, 6, 8},
		},
		{
			name:  "empty input",
			input: []int{},
			fn:    isGreater,
			want:  []int{},
		},
		{
			name:  "custom filter - greater than threshold",
			input: []int{0, 5, 10, 15, 20},
			fn:    isGreater,
			want:  []int{15, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inCh := sliceToChan(tt.input)
			outCh := Filter(inCh, tt.fn)
			got := chanToSlice(t, outCh)

			if len(got) != len(tt.want) {
				t.Errorf("unexpected slice length: got %d, want %d", len(got), len(tt.want))
			}

			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("unexpected value: got %d, want %d", v, tt.want[i])
				}
			}
		})
	}
}

// =============================================================================
// TestSum
// =============================================================================
//
// Test cases to implement:
//
// 1. "sums all values"
//    - Input: [1, 2, 3, 4, 5]
//    - Expected: 15
//
// 2. "handles empty channel"
//    - Input: [] (closed immediately)
//    - Expected: 0
//
// 3. "handles single value"
//    - Input: [42]
//    - Expected: 42
//
// 4. "handles negative values"
//    - Input: [-5, 10, -3, 8]
//    - Expected: 10
//
// 5. "blocks until channel closes"
//    - Create a channel, send some values with delay, then close
//    - Verify Sum doesn't return until channel is closed
//    - Use goroutine + timer to verify blocking behavior

// TODO(human): Implement TestSum
func TestSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "sum all values",
			input: []int{1, 2, 3, 4, 5},
			want:  15,
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  0,
		},
		{
			name:  "single value",
			input: []int{5},
			want:  5,
		},
		{
			name:  "negative values",
			input: []int{-5, -1, 0, 1, 5},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inCh := sliceToChan(tt.input)
			got := Sum(inCh)

			if got != tt.want {
				t.Errorf("unexpected sum result: got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestSumEmptyChannel(t *testing.T) {

	inCh := make(chan int)
	done := make(chan struct{})

	duration := time.Second * 1

	go func() {
		// this delay should be longer than "duration" to ensure that the main timer finishes before the sum function
		time.Sleep(2 * time.Second)
		close(inCh)
	}()

	go func() {
		Sum(inCh)
		done <- struct{}{}
	}()

	timer := time.NewTimer(duration)

	for {
		select {
		case <-done:
			t.Error("unexpected finish before timer")
			return
		case <-timer.C:
			return
		}
	}
}

// =============================================================================
// TestMerge
// =============================================================================
//
// Test cases to implement:
//
// 1. "merges two channels"
//    - ch1: [1, 2, 3], ch2: [4, 5, 6]
//    - Merged should contain all 6 values (order may vary)
//    - Sort results before comparing if needed
//
// 2. "merges three channels"
//    - Three channels with different values
//    - Verify all values present in output
//
// 3. "handles one empty channel"
//    - ch1: [1, 2, 3], ch2: [] (closed immediately)
//    - Output should contain [1, 2, 3]
//
// 4. "handles all empty channels"
//    - All input channels closed immediately
//    - Output should close with no values
//
// 5. "handles no channels"
//    - Call Merge() with no arguments
//    - Output should close immediately
//
// 6. "closes only after all inputs close"
//    - Use channels with delayed sends
//    - Verify output doesn't close prematurely
//
// Race condition note:
//    - Order of merged values is non-deterministic
//    - Sort results or use a set/map for comparison

// TODO(human): Implement TestMerge
func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  []int
	}{
		{
			name:  "merges two channels",
			input: [][]int{{1, 2, 3}, {4, 5, 6}},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "merges three channels",
			input: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "handles one empty channel",
			input: [][]int{{1, 2, 3}, {}},
			want:  []int{1, 2, 3},
		},
		// these test cases require separate test functions
		{
			name:  "handles no channels",
			input: nil,
			want:  []int{},
		},
		// {
		// 	name: "closes after all inputs close",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			channels := make([]<-chan int, len(tt.input))
			for i, slice := range tt.input {
				channels[i] = sliceToChan(slice)
			}

			got := chanToSlice(t, Merge(channels...))
			sort.Ints(got)

			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("unexpected result: got %d, want %d", v, tt.want[i])
				}
			}
		})
	}
}

func TestMergeClosingOrder(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	out := Merge(ch1, ch2, ch3)

	done := make(chan struct{})
	go func() {
		for range out {
		}
		close(done)
	}()

	close(ch1)
	// TODO: verify out is still open
	select {
	case <-done:
		t.Fatal("out closed too early after ch1")
	case <-time.After(100 * time.Millisecond):
	}

	close(ch2)
	// TODO: verify out is still open
	select {
	case <-done:
		t.Fatal("out closed too early after ch1")
	case <-time.After(100 * time.Millisecond):
	}

	close(ch3)
	// TODO: verify out is closed -> we closed all input channels already
	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("out closed too early after ch1")
	}

}

// =============================================================================
// TestTee
// =============================================================================
//
// Test cases to implement:
//
// 1. "broadcasts to both outputs"
//    - Input: [1, 2, 3]
//    - Both output channels should receive [1, 2, 3]
//    - Must consume BOTH outputs concurrently to avoid deadlock!
//
// 2. "handles empty input"
//    - Empty input channel (closed immediately)
//    - Both outputs should close with no values
//
// 3. "handles single value"
//    - Input: [42]
//    - Both outputs receive [42]
//
// CRITICAL testing note for Tee:
//    - If you drain ch1 completely before reading ch2, you may DEADLOCK
//    - Tee sends to BOTH channels for each input value
//    - If one output buffer fills, Tee blocks until both can receive
//    - You MUST read from both outputs concurrently
//
// Concurrent consumption pattern:
//    ch1, ch2 := Tee(input)
//    var wg sync.WaitGroup
//    var vals1, vals2 []int
//    wg.Add(2)
//    go func() { defer wg.Done(); for v := range ch1 { vals1 = append(vals1, v) } }()
//    go func() { defer wg.Done(); for v := range ch2 { vals2 = append(vals2, v) } }()
//    wg.Wait()
//    // Now compare vals1 and vals2

// TODO(human): Implement TestTee
func TestTee(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "broadcasts to both outputs",
			input: []int{1, 2, 3},
		},
		{
			name:  "handles empty input",
			input: []int{},
		},
		{
			name:  "handles single value",
			input: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			inCh := sliceToChan(tt.input)
			chA, chB := Tee(inCh)

			var resA []int
			var resB []int

			var wg sync.WaitGroup
			wg.Add(2)

			wg.Go(func() {
				defer wg.Done()
				resA = chanToSlice(t, chA)
			})

			wg.Go(func() {
				defer wg.Done()
				resB = chanToSlice(t, chB)
			})

			wg.Wait()

			if len(resA) != len(tt.input) || len(resB) != len(tt.input) {
				t.Errorf("unexpected result: got %v, want %v", resA, tt.input)
			}

			if !sortedEqual(resA, resB) {
				t.Errorf("unexpected unmatching results: A: %v, B: %v", resA, resB)
			}
		})
	}
}

// =============================================================================
// TestPipeline (Integration Test)
// =============================================================================
//
// Test the full pipeline composition:
//
// 1. "numbers -> square -> sum"
//    - Numbers(5) -> Square -> Sum
//    - Expected: 1² + 2² + 3² + 4² + 5² = 1 + 4 + 9 + 16 + 25 = 55
//
// 2. "numbers -> filter(even) -> sum"
//    - Numbers(10) -> Filter(even) -> Sum
//    - Expected: 2 + 4 + 6 + 8 + 10 = 30
//
// 3. "numbers -> square -> filter(>10) -> sum"
//    - Numbers(5) -> Square -> Filter(>10) -> Sum
//    - Squares: 1, 4, 9, 16, 25
//    - Filter >10: 16, 25
//    - Sum: 41
//
// 4. "complex pipeline with tee and merge"
//    - nums := Numbers(5)
//    - a, b := Tee(nums)
//    - squared := Square(a)
//    - filtered := Filter(b, isOdd)
//    - merged := Merge(squared, filtered)
//    - Collect and verify all values present

// TODO(human): Implement TestPipeline
func TestPipelineSimple(t *testing.T) {

	got := Sum(Square(Generator(5)))
	want := 55

	if got != want {
		t.Errorf("unexpected result: got %d, want %d", got, want)
	}
}

func TestPipelineEvenSum(t *testing.T) {
	evens := func(i int) bool { return i%2 == 0 }
	got := Sum(Filter(Generator(5), evens))
	want := 6

	if got != want {
		t.Errorf("unexpected result: got %d, want %d", got, want)
	}
}

func TestPipelineSquareFilterSum(t *testing.T) {
	filter := func(i int) bool { return i > 10 }
	got := Sum(Filter(Square(Generator(5)), filter))
	want := 41

	if got != want {
		t.Errorf("unexpected result: got %d, want %d", got, want)
	}
}

func TestPipelineComplex(t *testing.T) {
	isOdd := func(i int) bool { return i%2 != 0 }

	nums := Generator(5) // 1, 2, 3, 4, 5
	a, b := Tee(nums)
	squared := Square(a)               // 1, 4, 9, 16, 25
	filtered := Filter(b, isOdd)       // 1, 3, 5
	merged := Merge(squared, filtered) // 1, 1, 3, 4, 5, 9, 16, 25

	got := chanToSlice(t, merged)
	want := []int{1, 1, 3, 4, 5, 9, 16, 25}

	if !sortedEqual(got, want) {
		t.Errorf("unexpected result: got %v, want %v", got, want)
	}
}

// =============================================================================
// Benchmark Suggestions (Optional)
// =============================================================================
//
// If you want to practice benchmarking concurrent code:
//
// BenchmarkNumbers - measure throughput of generator
// BenchmarkSquare - measure pipeline stage throughput
// BenchmarkMerge - compare merging 2 vs 4 vs 8 channels
//
// Remember: benchmarks with goroutines can be tricky because setup
// costs are amortized differently. Use b.ResetTimer() after setup.

// =============================================================================
// Helper Function Suggestions
// =============================================================================
//
// Consider implementing these helpers to make testing easier:
//
// 1. sliceToChan(vals []int) <-chan int
//    - Converts a slice to a channel (useful for creating test inputs)
//
// 2. chanToSlice(ch <-chan int) []int
//    - Drains a channel into a slice (useful for collecting outputs)
//
// 3. sortedEqual(a, b []int) bool
//    - Compares two slices ignoring order (useful for Merge tests)
//
// These helpers make tests more readable and reduce boilerplate.
