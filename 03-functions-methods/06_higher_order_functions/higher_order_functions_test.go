package higher_order_functions

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		fn    func(int) int
		want  []int
	}{
		{
			"double",
			[]int{1, 2, 3, 4},
			func(n int) int { return n * 2 },
			[]int{2, 4, 6, 8},
		},
		{
			"add ten",
			[]int{1, 2, 3},
			func(n int) int { return n + 10 },
			[]int{11, 12, 13},
		},
		{
			"square",
			[]int{1, 2, 3, 4},
			func(n int) int { return n * n },
			[]int{1, 4, 9, 16},
		},
		{
			"empty slice",
			[]int{},
			func(n int) int { return n * 2 },
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.slice, tt.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
	}{
		{
			"even numbers",
			[]int{1, 2, 3, 4, 5, 6},
			func(n int) bool { return n%2 == 0 },
			[]int{2, 4, 6},
		},
		{
			"greater than 3",
			[]int{1, 2, 3, 4, 5},
			func(n int) bool { return n > 3 },
			[]int{4, 5},
		},
		{
			"negative numbers",
			[]int{-2, -1, 0, 1, 2},
			func(n int) bool { return n < 0 },
			[]int{-2, -1},
		},
		{
			"none match",
			[]int{1, 2, 3},
			func(n int) bool { return n > 10 },
			[]int{},
		},
		{
			"all match",
			[]int{2, 4, 6},
			func(n int) bool { return n%2 == 0 },
			[]int{2, 4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.slice, tt.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		initial int
		fn      func(int, int) int
		want    int
	}{
		{
			"sum",
			[]int{1, 2, 3, 4},
			0,
			func(acc, n int) int { return acc + n },
			10,
		},
		{
			"product",
			[]int{1, 2, 3, 4},
			1,
			func(acc, n int) int { return acc * n },
			24,
		},
		{
			"max",
			[]int{5, 2, 9, 1, 7},
			0,
			func(acc, n int) int {
				if n > acc {
					return n
				}
				return acc
			},
			9,
		},
		{
			"empty slice",
			[]int{},
			42,
			func(acc, n int) int { return acc + n },
			42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.slice, tt.initial, tt.fn)
			if got != tt.want {
				t.Errorf("Reduce() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	double := func(n int) int { return n * 2 }
	addTen := func(n int) int { return n + 10 }
	square := func(n int) int { return n * n }

	tests := []struct {
		name string
		f    func(int) int
		g    func(int) int
		x    int
		want int
	}{
		{"double then addTen", double, addTen, 5, 30},   // double(addTen(5)) = double(15) = 30
		{"addTen then double", addTen, double, 5, 20},   // addTen(double(5)) = addTen(10) = 20
		{"square then double", double, square, 3, 18},   // double(square(3)) = double(9) = 18
		{"double then square", square, double, 3, 36},   // square(double(3)) = square(6) = 36
		{"addTen then addTen", addTen, addTen, 5, 25},   // addTen(addTen(5)) = addTen(15) = 25
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			composed := Compose(tt.f, tt.g)
			got := composed(tt.x)
			if got != tt.want {
				t.Errorf("Compose(f, g)(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	// Test by collecting results
	var results []int
	ForEach([]int{1, 2, 3, 4}, func(n int) {
		results = append(results, n*2)
	})

	want := []int{2, 4, 6, 8}
	if !reflect.DeepEqual(results, want) {
		t.Errorf("ForEach collected %v, want %v", results, want)
	}

	// Test with empty slice
	var count int
	ForEach([]int{}, func(n int) {
		count++
	})
	if count != 0 {
		t.Errorf("ForEach on empty slice called function %d times, want 0", count)
	}
}

func TestAny(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      bool
	}{
		{
			"has even",
			[]int{1, 3, 4, 5},
			func(n int) bool { return n%2 == 0 },
			true,
		},
		{
			"no even",
			[]int{1, 3, 5},
			func(n int) bool { return n%2 == 0 },
			false,
		},
		{
			"has negative",
			[]int{1, 2, -1, 3},
			func(n int) bool { return n < 0 },
			true,
		},
		{
			"empty slice",
			[]int{},
			func(n int) bool { return true },
			false,
		},
		{
			"all match",
			[]int{2, 4, 6},
			func(n int) bool { return n%2 == 0 },
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Any(tt.slice, tt.predicate)
			if got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      bool
	}{
		{
			"all even",
			[]int{2, 4, 6},
			func(n int) bool { return n%2 == 0 },
			true,
		},
		{
			"not all even",
			[]int{2, 3, 4},
			func(n int) bool { return n%2 == 0 },
			false,
		},
		{
			"all positive",
			[]int{1, 2, 3},
			func(n int) bool { return n > 0 },
			true,
		},
		{
			"empty slice",
			[]int{},
			func(n int) bool { return false },
			true, // vacuous truth
		},
		{
			"none match",
			[]int{1, 3, 5},
			func(n int) bool { return n%2 == 0 },
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := All(tt.slice, tt.predicate)
			if got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMap(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	double := func(n int) int { return n * 2 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Map(slice, double)
	}
}

func BenchmarkFilter(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	isEven := func(n int) bool { return n%2 == 0 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(slice, isEven)
	}
}

func BenchmarkReduce(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	add := func(acc, n int) int { return acc + n }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(slice, 0, add)
	}
}
