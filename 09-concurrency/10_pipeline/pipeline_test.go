package pipeline

import (
	"testing"
	"time"
)

func TestGenerator(t *testing.T) {
	t.Run("emits all values", func(t *testing.T) {
		gen := Generator(1, 2, 3, 4, 5)
		if gen == nil {
			t.Fatal("Generator() returned nil")
		}

		var got []int
		for v := range gen {
			got = append(got, v)
		}

		if len(got) != 5 {
			t.Errorf("Generator() produced %d values, want 5", len(got))
		}

		for i, v := range got {
			if v != i+1 {
				t.Errorf("got[%d] = %d, want %d", i, v, i+1)
			}
		}
	})

	t.Run("empty generator", func(t *testing.T) {
		gen := Generator[int]()
		if gen == nil {
			t.Fatal("Generator() returned nil")
		}

		count := 0
		for range gen {
			count++
		}

		if count != 0 {
			t.Errorf("Empty Generator() produced %d values", count)
		}
	})

	t.Run("closes channel", func(t *testing.T) {
		gen := Generator(1)
		<-gen // consume value

		select {
		case _, ok := <-gen:
			if ok {
				t.Error("Generator channel should be closed")
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Generator channel did not close")
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("filters values", func(t *testing.T) {
		input := Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		stage := Filter(func(n int) bool { return n%2 == 0 })
		if stage == nil {
			t.Fatal("Filter() returned nil")
		}

		output := stage(input)
		if output == nil {
			t.Fatal("Filter stage returned nil")
		}

		var got []int
		for v := range output {
			got = append(got, v)
		}

		want := []int{2, 4, 6, 8, 10}
		if len(got) != len(want) {
			t.Errorf("Filter() produced %d values, want %d", len(got), len(want))
		}
		for i, v := range got {
			if v != want[i] {
				t.Errorf("got[%d] = %d, want %d", i, v, want[i])
			}
		}
	})

	t.Run("no matches", func(t *testing.T) {
		input := Generator(1, 3, 5)
		output := Filter(func(n int) bool { return n%2 == 0 })(input)

		count := 0
		for range output {
			count++
		}

		if count != 0 {
			t.Errorf("Filter with no matches produced %d values", count)
		}
	})
}

func TestMap(t *testing.T) {
	t.Run("transforms values", func(t *testing.T) {
		input := Generator(1, 2, 3)
		stage := Map(func(n int) int { return n * 2 })
		if stage == nil {
			t.Fatal("Map() returned nil")
		}

		output := stage(input)
		if output == nil {
			t.Fatal("Map stage returned nil")
		}

		var got []int
		for v := range output {
			got = append(got, v)
		}

		want := []int{2, 4, 6}
		for i, v := range got {
			if v != want[i] {
				t.Errorf("got[%d] = %d, want %d", i, v, want[i])
			}
		}
	})

	t.Run("type conversion", func(t *testing.T) {
		input := Generator(1, 2, 3)
		output := Map(func(n int) string {
			return string(rune('A' + n - 1))
		})(input)

		var got []string
		for v := range output {
			got = append(got, v)
		}

		want := []string{"A", "B", "C"}
		for i, v := range got {
			if v != want[i] {
				t.Errorf("got[%d] = %q, want %q", i, v, want[i])
			}
		}
	})
}

func TestTake(t *testing.T) {
	t.Run("takes first n", func(t *testing.T) {
		input := Generator(1, 2, 3, 4, 5)
		stage := Take[int](3)
		if stage == nil {
			t.Fatal("Take() returned nil")
		}

		output := stage(input)
		if output == nil {
			t.Fatal("Take stage returned nil")
		}

		var got []int
		for v := range output {
			got = append(got, v)
		}

		if len(got) != 3 {
			t.Errorf("Take(3) produced %d values, want 3", len(got))
		}

		want := []int{1, 2, 3}
		for i, v := range got {
			if v != want[i] {
				t.Errorf("got[%d] = %d, want %d", i, v, want[i])
			}
		}
	})

	t.Run("take more than available", func(t *testing.T) {
		input := Generator(1, 2)
		output := Take[int](5)(input)

		var got []int
		for v := range output {
			got = append(got, v)
		}

		if len(got) != 2 {
			t.Errorf("Take(5) from 2 values produced %d values, want 2", len(got))
		}
	})

	t.Run("take zero", func(t *testing.T) {
		input := Generator(1, 2, 3)
		output := Take[int](0)(input)

		count := 0
		for range output {
			count++
		}

		if count != 0 {
			t.Errorf("Take(0) produced %d values, want 0", count)
		}
	})
}

func TestSkip(t *testing.T) {
	t.Run("skips first n", func(t *testing.T) {
		input := Generator(1, 2, 3, 4, 5)
		stage := Skip[int](2)
		if stage == nil {
			t.Fatal("Skip() returned nil")
		}

		output := stage(input)
		if output == nil {
			t.Fatal("Skip stage returned nil")
		}

		var got []int
		for v := range output {
			got = append(got, v)
		}

		want := []int{3, 4, 5}
		if len(got) != len(want) {
			t.Errorf("Skip(2) produced %d values, want %d", len(got), len(want))
		}
		for i, v := range got {
			if v != want[i] {
				t.Errorf("got[%d] = %d, want %d", i, v, want[i])
			}
		}
	})

	t.Run("skip more than available", func(t *testing.T) {
		input := Generator(1, 2)
		output := Skip[int](5)(input)

		count := 0
		for range output {
			count++
		}

		if count != 0 {
			t.Errorf("Skip(5) from 2 values produced %d values", count)
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("sums values", func(t *testing.T) {
		input := Generator(1, 2, 3, 4, 5)
		reducer := Reduce(0, func(acc, n int) int { return acc + n })
		if reducer == nil {
			t.Fatal("Reduce() returned nil")
		}

		result := reducer(input)
		if result != 15 {
			t.Errorf("Reduce sum = %d, want 15", result)
		}
	})

	t.Run("concatenates strings", func(t *testing.T) {
		input := Generator("a", "b", "c")
		result := Reduce("", func(acc, s string) string { return acc + s })(input)

		if result != "abc" {
			t.Errorf("Reduce concat = %q, want 'abc'", result)
		}
	})

	t.Run("empty input", func(t *testing.T) {
		input := Generator[int]()
		result := Reduce(42, func(acc, n int) int { return acc + n })(input)

		if result != 42 {
			t.Errorf("Reduce empty = %d, want initial value 42", result)
		}
	})
}

func TestPipeline(t *testing.T) {
	t.Run("chains stages", func(t *testing.T) {
		input := Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

		output := Pipeline(input,
			Filter(func(n int) bool { return n%2 == 0 }),
			Map(func(n int) int { return n * 2 }),
			Take[int](3),
		)

		if output == nil {
			t.Fatal("Pipeline() returned nil")
		}

		var got []int
		for v := range output {
			got = append(got, v)
		}

		want := []int{4, 8, 12}
		if len(got) != len(want) {
			t.Fatalf("Pipeline produced %d values, want %d", len(got), len(want))
		}
		for i, v := range got {
			if v != want[i] {
				t.Errorf("got[%d] = %d, want %d", i, v, want[i])
			}
		}
	})

	t.Run("no stages", func(t *testing.T) {
		input := Generator(1, 2, 3)
		output := Pipeline(input)

		var got []int
		for v := range output {
			got = append(got, v)
		}

		if len(got) != 3 {
			t.Errorf("Pipeline with no stages produced %d values, want 3", len(got))
		}
	})

	t.Run("complex pipeline", func(t *testing.T) {
		input := Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

		output := Pipeline(input,
			Skip[int](2),                                   // 3,4,5,6,7,8,9,10
			Filter(func(n int) bool { return n%2 == 1 }),   // 3,5,7,9
			Map(func(n int) int { return n * 10 }),         // 30,50,70,90
			Take[int](2),                                   // 30,50
		)

		var got []int
		for v := range output {
			got = append(got, v)
		}

		want := []int{30, 50}
		if len(got) != len(want) {
			t.Fatalf("Complex pipeline produced %d values, want %d", len(got), len(want))
		}
		for i, v := range got {
			if v != want[i] {
				t.Errorf("got[%d] = %d, want %d", i, v, want[i])
			}
		}
	})
}

func TestPipelineAsync(t *testing.T) {
	t.Run("parallel processing", func(t *testing.T) {
		input := Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

		output := PipelineAsync(input, 3,
			Map(func(n int) int {
				time.Sleep(50 * time.Millisecond)
				return n * 2
			}),
		)

		if output == nil {
			t.Skip("PipelineAsync not implemented (bonus)")
		}

		start := time.Now()
		count := 0
		for range output {
			count++
		}
		elapsed := time.Since(start)

		if count != 10 {
			t.Errorf("PipelineAsync produced %d values, want 10", count)
		}

		// With 3 workers, should take ~200ms instead of ~500ms
		if elapsed > 300*time.Millisecond {
			t.Errorf("PipelineAsync took %v, expected parallel execution", elapsed)
		}
	})
}

// Benchmarks
func BenchmarkPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		output := Pipeline(input,
			Filter(func(n int) bool { return n%2 == 0 }),
			Map(func(n int) int { return n * 2 }),
		)
		for range output {
		}
	}
}

func BenchmarkReduce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		Reduce(0, func(acc, n int) int { return acc + n })(input)
	}
}
