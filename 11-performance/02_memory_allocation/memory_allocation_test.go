package memory_allocation

import (
	"strings"
	"testing"
)

// =============================================================================
// Unit Tests
// =============================================================================

func TestConcatStrings(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{[]string{}, ""},
		{[]string{"hello"}, "hello"},
		{[]string{"hello", " ", "world"}, "hello world"},
	}

	for _, tt := range tests {
		if got := ConcatStrings(tt.input); got != tt.want {
			t.Errorf("ConcatStrings(%v) = %q, want %q", tt.input, got, tt.want)
		}
		if got := ConcatStringsOptimized(tt.input); got != tt.want {
			t.Errorf("ConcatStringsOptimized(%v) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestProcessItems(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	want := []int{2, 4, 6, 8, 10}

	got := ProcessItems(input)
	if len(got) != len(want) {
		t.Fatalf("ProcessItems() len = %d, want %d", len(got), len(want))
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("ProcessItems()[%d] = %d, want %d", i, v, want[i])
		}
	}

	// Verify original not modified
	if input[0] != 1 {
		t.Error("ProcessItems() modified original slice")
	}
}

func TestProcessItemsInPlace(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	want := []int{2, 4, 6, 8, 10}

	got := ProcessItemsInPlace(input)
	for i, v := range got {
		if v != want[i] {
			t.Errorf("ProcessItemsInPlace()[%d] = %d, want %d", i, v, want[i])
		}
	}
}

func TestCreateUsers(t *testing.T) {
	names := []string{"Alice", "Bob", "Charlie"}

	ptrs := CreateUsers(names)
	if len(ptrs) != 3 {
		t.Errorf("CreateUsers() len = %d, want 3", len(ptrs))
	}

	vals := CreateUsersOptimized(names)
	if len(vals) != 3 {
		t.Errorf("CreateUsersOptimized() len = %d, want 3", len(vals))
	}
}

// =============================================================================
// Benchmarks
// =============================================================================

func BenchmarkConcatStrings(b *testing.B) {
	strs := strings.Split("the quick brown fox jumps over the lazy dog", " ")

	b.Run("Naive", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ConcatStrings(strs)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ConcatStringsOptimized(strs)
		}
	})
}

func BenchmarkProcessItems(b *testing.B) {
	items := make([]int, 1000)
	for i := range items {
		items[i] = i
	}

	b.Run("NewSlice", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ProcessItems(items)
		}
	})

	b.Run("InPlace", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// Copy to avoid accumulation
			data := make([]int, len(items))
			copy(data, items)
			ProcessItemsInPlace(data)
		}
	})
}

func BenchmarkCreateUsers(b *testing.B) {
	names := make([]string, 100)
	for i := range names {
		names[i] = "User"
	}

	b.Run("Pointers", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			CreateUsers(names)
		}
	})

	b.Run("Values", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			CreateUsersOptimized(names)
		}
	})
}
