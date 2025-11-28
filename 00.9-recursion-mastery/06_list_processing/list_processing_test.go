package list_processing

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		name   string
		nested any
		want   []int
	}{
		{"Single integer", 5, []int{5}},
		{"Flat slice", []any{1, 2, 3}, []int{1, 2, 3}},
		{"One level nesting", []any{1, []any{2, 3}}, []int{1, 2, 3}},
		{"Two level nesting", []any{1, []any{2, []any{3}}}, []int{1, 2, 3}},
		{"Complex nesting", []any{1, []any{2, 3}, []any{4, []any{5, 6}}}, []int{1, 2, 3, 4, 5, 6}},
		{"Empty slice", []any{}, []int{}},
		{"Nested empties", []any{1, []any{}, 2}, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Flatten(tt.nested)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten(%v) = %v, want %v", tt.nested, got, tt.want)
			}
		})
	}
}

func TestDeepSum(t *testing.T) {
	tests := []struct {
		name   string
		nested any
		want   int
	}{
		{"Single integer", 5, 5},
		{"Flat slice", []any{1, 2, 3}, 6},
		{"One level nesting", []any{1, []any{2, 3}}, 6},
		{"Two level nesting", []any{1, []any{2, []any{3, 4}}}, 10},
		{"Complex nesting", []any{1, []any{2, 3}, []any{4, []any{5, 6}}}, 21},
		{"With zeros", []any{0, []any{1, 0}, 2}, 3},
		{"Empty slice", []any{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeepSum(tt.nested)
			if got != tt.want {
				t.Errorf("DeepSum(%v) = %d, want %d", tt.nested, got, tt.want)
			}
		})
	}
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name   string
		nested any
		want   int
	}{
		{"Single integer", 5, 1},
		{"Flat slice", []any{1, 2, 3}, 1},
		{"One level nesting", []any{1, []any{2}}, 2},
		{"Two level nesting", []any{1, []any{2, []any{3}}}, 3},
		{"Unbalanced tree", []any{1, []any{2}, []any{3, []any{4}}}, 3},
		{"Deep nesting", []any{[]any{[]any{[]any{1}}}}, 4},
		{"Empty slice", []any{}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxDepth(tt.nested)
			if got != tt.want {
				t.Errorf("MaxDepth(%v) = %d, want %d", tt.nested, got, tt.want)
			}
		})
	}
}

func TestFilterNested(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	isPositive := func(n int) bool { return n > 0 }
	greaterThan3 := func(n int) bool { return n > 3 }

	tests := []struct {
		name      string
		nested    any
		predicate func(int) bool
		want      []int
	}{
		{"Filter even - flat", []any{1, 2, 3, 4}, isEven, []int{2, 4}},
		{"Filter even - nested", []any{1, []any{2, 3}, 4}, isEven, []int{2, 4}},
		{"Filter even - complex", []any{1, []any{2, 3}, []any{4, []any{5, 6}}}, isEven, []int{2, 4, 6}},
		{"Filter positive", []any{-1, []any{2, -3}, 4}, isPositive, []int{2, 4}},
		{"Filter > 3", []any{1, 2, []any{3, 4, 5}}, greaterThan3, []int{4, 5}},
		{"No matches", []any{1, 3, 5}, isEven, []int{}},
		{"All match", []any{2, []any{4, 6}}, isEven, []int{2, 4, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterNested(tt.nested, tt.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterNested(%v, predicate) = %v, want %v", tt.nested, got, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkFlatten(b *testing.B) {
	nested := []any{1, []any{2, 3}, []any{4, []any{5, 6}}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Flatten(nested)
	}
}

func BenchmarkDeepSum(b *testing.B) {
	nested := []any{1, []any{2, 3}, []any{4, []any{5, 6}}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepSum(nested)
	}
}

func BenchmarkMaxDepth(b *testing.B) {
	nested := []any{1, []any{2, []any{3, []any{4}}}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxDepth(nested)
	}
}
