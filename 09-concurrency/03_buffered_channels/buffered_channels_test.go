package buffered_channels

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestBufferDemo(t *testing.T) {
	tests := []struct {
		name       string
		values     []int
		bufferSize int
		want       []int
	}{
		{"empty", []int{}, 5, []int{}},
		{"single value", []int{1}, 1, []int{1}},
		{"buffer equals values", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"buffer smaller", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 4, 5}},
		{"buffer larger", []int{1, 2}, 10, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BufferDemo(tt.values, tt.bufferSize)
			if len(got) != len(tt.want) {
				t.Errorf("BufferDemo() returned %d values, want %d", len(got), len(tt.want))
				return
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("BufferDemo()[%d] = %d, want %d", i, v, tt.want[i])
				}
			}
		})
	}
}

func TestProducer(t *testing.T) {
	tests := []struct {
		name       string
		values     []int
		bufferSize int
	}{
		{"empty", []int{}, 5},
		{"single", []int{42}, 1},
		{"multiple", []int{1, 2, 3, 4, 5}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := Producer(tt.values, tt.bufferSize)
			if ch == nil {
				t.Fatal("Producer() returned nil channel")
			}

			var got []int
			for v := range ch {
				got = append(got, v)
			}

			if len(got) != len(tt.values) {
				t.Errorf("Producer() produced %d values, want %d", len(got), len(tt.values))
			}
		})
	}
}

func TestConsumer(t *testing.T) {
	tests := []struct {
		name   string
		values []int
	}{
		{"empty", []int{}},
		{"single", []int{42}},
		{"multiple", []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan int, len(tt.values))
			for _, v := range tt.values {
				ch <- v
			}
			close(ch)

			got := Consumer(ch)
			if len(got) != len(tt.values) {
				t.Errorf("Consumer() returned %d values, want %d", len(got), len(tt.values))
			}
		})
	}
}

func TestProducerConsumer(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := Producer(values, 3)
	got := Consumer(ch)

	if len(got) != len(values) {
		t.Errorf("Producer/Consumer pipeline returned %d values, want %d", len(got), len(values))
	}

	for i, v := range got {
		if v != values[i] {
			t.Errorf("Value at index %d = %d, want %d", i, v, values[i])
		}
	}
}

func TestBatchCollector(t *testing.T) {
	tests := []struct {
		name      string
		values    []int
		batchSize int
		want      [][]int
	}{
		{"empty", []int{}, 3, [][]int{}},
		{"exact batches", []int{1, 2, 3, 4, 5, 6}, 3, [][]int{{1, 2, 3}, {4, 5, 6}}},
		{"partial last", []int{1, 2, 3, 4, 5, 6, 7}, 3, [][]int{{1, 2, 3}, {4, 5, 6}, {7}}},
		{"single batch", []int{1, 2}, 5, [][]int{{1, 2}}},
		{"batch size 1", []int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := Producer(tt.values, 10)
			batches := BatchCollector(in, tt.batchSize)
			if batches == nil {
				t.Fatal("BatchCollector() returned nil channel")
			}

			var got [][]int
			for batch := range batches {
				got = append(got, batch)
			}

			if len(got) != len(tt.want) {
				t.Errorf("BatchCollector() produced %d batches, want %d", len(got), len(tt.want))
				return
			}

			for i, batch := range got {
				if len(batch) != len(tt.want[i]) {
					t.Errorf("Batch %d has %d items, want %d", i, len(batch), len(tt.want[i]))
					continue
				}
				for j, v := range batch {
					if v != tt.want[i][j] {
						t.Errorf("Batch %d[%d] = %d, want %d", i, j, v, tt.want[i][j])
					}
				}
			}
		})
	}
}

func TestSemaphore(t *testing.T) {
	limit := 3
	sem := Semaphore(limit)
	if sem == nil {
		t.Fatal("Semaphore() returned nil")
	}

	var concurrent int64
	var maxConcurrent int64
	var wg sync.WaitGroup

	// Run many tasks, track max concurrent
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem(func() {
				current := atomic.AddInt64(&concurrent, 1)
				// Track max
				for {
					max := atomic.LoadInt64(&maxConcurrent)
					if current <= max || atomic.CompareAndSwapInt64(&maxConcurrent, max, current) {
						break
					}
				}
				time.Sleep(10 * time.Millisecond) // Simulate work
				atomic.AddInt64(&concurrent, -1)
			})
		}()
	}

	wg.Wait()

	if maxConcurrent > int64(limit) {
		t.Errorf("Semaphore allowed %d concurrent, limit was %d", maxConcurrent, limit)
	}

	if maxConcurrent == 0 {
		t.Error("Semaphore appears to have blocked all work")
	}
}

func TestSemaphoreDoesWork(t *testing.T) {
	sem := Semaphore(2)
	var counter int64

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem(func() {
				atomic.AddInt64(&counter, 1)
			})
		}()
	}
	wg.Wait()

	if counter != 10 {
		t.Errorf("Semaphore completed %d tasks, want 10", counter)
	}
}

// Benchmarks
func BenchmarkBufferDemo(b *testing.B) {
	values := make([]int, 100)
	for i := range values {
		values[i] = i
	}

	b.Run("buffer=1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BufferDemo(values, 1)
		}
	})

	b.Run("buffer=10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BufferDemo(values, 10)
		}
	})

	b.Run("buffer=100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BufferDemo(values, 100)
		}
	})
}
