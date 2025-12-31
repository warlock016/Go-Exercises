package goroutine_basics

import (
	"bytes"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestPrintMessages(t *testing.T) {
	// This test just ensures the function doesn't panic
	// We can't reliably test output since goroutines may not run
	PrintMessages([]string{"hello", "world"})
	// Give goroutines a tiny chance to run (not reliable, just for demo)
}

func TestPrintMessagesSync(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	messages := []string{"alpha", "beta", "gamma"}
	PrintMessagesSync(messages)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Check all messages appear (order may vary)
	for _, msg := range messages {
		if !strings.Contains(output, msg) {
			t.Errorf("PrintMessagesSync() output missing %q, got: %q", msg, output)
		}
	}
}

func TestCounterUnsafe(t *testing.T) {
	// Run multiple times to demonstrate inconsistency
	results := make(map[int]int)
	n := 1000

	for i := 0; i < 10; i++ {
		result := CounterUnsafe(n)
		results[result]++
	}

	// With a race condition, we expect varied results
	// If counter is always exactly n, the race condition isn't demonstrated
	// (but we can't fail the test since mutex might accidentally work)
	t.Logf("CounterUnsafe results over 10 runs with n=%d: %v", n, results)

	// At minimum, check it runs without panic
	if len(results) == 0 {
		t.Error("CounterUnsafe() produced no results")
	}
}

func TestCounterSafe(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero goroutines", 0, 0},
		{"one goroutine", 1, 1},
		{"ten goroutines", 10, 10},
		{"hundred goroutines", 100, 100},
		{"thousand goroutines", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CounterSafe(tt.n)
			if got != tt.want {
				t.Errorf("CounterSafe(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestCounterSafeConsistency(t *testing.T) {
	// Run multiple times to verify consistency
	n := 1000
	for i := 0; i < 10; i++ {
		got := CounterSafe(n)
		if got != n {
			t.Errorf("CounterSafe(%d) run %d = %d, want %d", n, i, got, n)
		}
	}
}

// Benchmark to compare unsafe vs safe counter
func BenchmarkCounterUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CounterUnsafe(100)
	}
}

func BenchmarkCounterSafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CounterSafe(100)
	}
}

// Example showing WaitGroup pattern
func ExamplePrintMessagesSync() {
	// Demonstrates the sync pattern
	var wg sync.WaitGroup
	messages := []string{"a", "b", "c"}

	for _, msg := range messages {
		wg.Add(1)
		go func(m string) {
			defer wg.Done()
			_ = m // Would print in real implementation
		}(msg)
	}
	wg.Wait()
	// Output is unpredictable due to goroutine scheduling
}
