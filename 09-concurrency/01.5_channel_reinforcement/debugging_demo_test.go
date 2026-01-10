package channel_reinforcement

import (
	"testing"
	"time"
)

// =============================================================================
// DEBUGGING DEMO TESTS
// =============================================================================
//
// Run these tests to observe the bugs:
//   go test -v -run TestBuggy -timeout 5s
//   go test -race -run TestBuggy -timeout 5s
//
// Use Delve for deeper investigation:
//   dlv test -- -test.run TestBuggyWorkerPool -test.timeout 5s
//
// =============================================================================

func TestBuggyWorkerPool(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{"simple", []int{1, 2, 3}, 12},           // (1+2+3)*2 = 12
		{"six_values", []int{1, 2, 3, 4, 5, 6}, 42}, // (1+2+3+4+5+6)*2 = 42
		{"empty", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Add timeout to detect deadlock
			done := make(chan int)
			go func() {
				done <- BuggyWorkerPool(tt.values)
			}()

			select {
			case got := <-done:
				if got != tt.want {
					t.Errorf("BuggyWorkerPool(%v) = %d, want %d", tt.values, got, tt.want)
				}
			case <-time.After(2 * time.Second):
				t.Fatalf("BuggyWorkerPool(%v) timed out - likely deadlock", tt.values)
			}
		})
	}
}

func TestBuggyPingPong(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"three", 3, 3},
		{"five", 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			done := make(chan int)
			go func() {
				done <- BuggyPingPong(tt.n)
			}()

			select {
			case got := <-done:
				if got != tt.want {
					t.Errorf("BuggyPingPong(%d) = %d, want %d", tt.n, got, tt.want)
				}
			case <-time.After(2 * time.Second):
				t.Fatalf("BuggyPingPong(%d) timed out - likely deadlock", tt.n)
			}
		})
	}
}

func TestBuggyFanOut(t *testing.T) {
	tests := []struct {
		name       string
		values     []int
		numWorkers int
		want       int
	}{
		{"simple", []int{1, 2, 3, 4}, 2, 10},
		{"four_workers", []int{1, 2, 3, 4, 5, 6, 7, 8}, 4, 36},
		{"empty", []int{}, 2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			done := make(chan int)
			go func() {
				done <- BuggyFanOut(tt.values, tt.numWorkers)
			}()

			select {
			case got := <-done:
				if got != tt.want {
					t.Errorf("BuggyFanOut(%v, %d) = %d, want %d", tt.values, tt.numWorkers, got, tt.want)
				}
			case <-time.After(2 * time.Second):
				t.Fatalf("BuggyFanOut(%v, %d) timed out - likely deadlock", tt.values, tt.numWorkers)
			}
		})
	}
}

// =============================================================================
// HINTS (Don't read until you've tried debugging!)
// =============================================================================
//
// BuggyWorkerPool hints:
// <details>
// 1. Look at the results slice - what are its elements?
// 2. How do workers know which result channel to use?
// </details>
//
// BuggyPingPong hints:
// <details>
// 1. Trace the flow: ping sends to pong, pong sends back to ping...
// 2. After n bounces, where is the final value?
// 3. What happens after the loops end?
// </details>
//
// BuggyFanOut hints:
// <details>
// 1. Run with -race flag
// 2. Look carefully at the inner loop index
// 3. What if len(values) is not evenly divisible by numWorkers?
// </details>
//
