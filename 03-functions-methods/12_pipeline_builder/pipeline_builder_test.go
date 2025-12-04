package pipeline_builder

import "testing"

func TestPipelineEmpty(t *testing.T) {
	pipeline := NewPipeline()
	result := pipeline.Execute(42)

	if result != 42 {
		t.Errorf("Empty pipeline should return input unchanged, got %d", result)
	}
}

func TestPipelineSingleTransform(t *testing.T) {
	pipeline := NewPipeline().Add(Add(10))
	result := pipeline.Execute(5)

	if result != 15 {
		t.Errorf("Pipeline.Execute(5) = %d, want 15", result)
	}
}

func TestPipelineMultipleTransforms(t *testing.T) {
	// 5 -> +10 -> *2 -> square
	// 5 -> 15 -> 30 -> 900
	pipeline := NewPipeline().
		Add(Add(10)).
		Add(Multiply(2)).
		Add(Square())

	result := pipeline.Execute(5)

	if result != 900 {
		t.Errorf("Pipeline.Execute(5) = %d, want 900", result)
	}
}

func TestPipelineReuse(t *testing.T) {
	pipeline := NewPipeline().
		Add(Multiply(2)).
		Add(Add(10))

	tests := []struct {
		input int
		want  int
	}{
		{5, 20},   // 5*2 + 10 = 20
		{10, 30},  // 10*2 + 10 = 30
		{0, 10},   // 0*2 + 10 = 10
		{-5, 0},   // -5*2 + 10 = 0
	}

	for _, tt := range tests {
		got := pipeline.Execute(tt.input)
		if got != tt.want {
			t.Errorf("Pipeline.Execute(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestAddTransform(t *testing.T) {
	tests := []struct {
		n     int
		input int
		want  int
	}{
		{10, 5, 15},
		{-5, 10, 5},
		{0, 42, 42},
		{100, -50, 50},
	}

	for _, tt := range tests {
		transform := Add(tt.n)
		got := transform(tt.input)
		if got != tt.want {
			t.Errorf("Add(%d)(%d) = %d, want %d", tt.n, tt.input, got, tt.want)
		}
	}
}

func TestMultiplyTransform(t *testing.T) {
	tests := []struct {
		n     int
		input int
		want  int
	}{
		{2, 5, 10},
		{3, 4, 12},
		{-1, 5, -5},
		{0, 42, 0},
	}

	for _, tt := range tests {
		transform := Multiply(tt.n)
		got := transform(tt.input)
		if got != tt.want {
			t.Errorf("Multiply(%d)(%d) = %d, want %d", tt.n, tt.input, got, tt.want)
		}
	}
}

func TestSquareTransform(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 1},
		{2, 4},
		{5, 25},
		{-3, 9},
		{10, 100},
	}

	square := Square()
	for _, tt := range tests {
		got := square(tt.input)
		if got != tt.want {
			t.Errorf("Square()(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestPipelineComplexChain(t *testing.T) {
	// Test: ((5 + 3) * 2) ^ 2 + 10 = (8 * 2) ^ 2 + 10 = 16 ^ 2 + 10 = 256 + 10 = 266
	pipeline := NewPipeline().
		Add(Add(3)).
		Add(Multiply(2)).
		Add(Square()).
		Add(Add(10))

	result := pipeline.Execute(5)

	if result != 266 {
		t.Errorf("Pipeline.Execute(5) = %d, want 266", result)
	}
}

func TestPipelineOrder(t *testing.T) {
	// Verify that order matters
	// Pipeline 1: (5 + 10) * 2 = 30
	pipeline1 := NewPipeline().
		Add(Add(10)).
		Add(Multiply(2))

	// Pipeline 2: (5 * 2) + 10 = 20
	pipeline2 := NewPipeline().
		Add(Multiply(2)).
		Add(Add(10))

	result1 := pipeline1.Execute(5)
	result2 := pipeline2.Execute(5)

	if result1 != 30 {
		t.Errorf("Pipeline1.Execute(5) = %d, want 30", result1)
	}

	if result2 != 20 {
		t.Errorf("Pipeline2.Execute(5) = %d, want 20", result2)
	}

	if result1 == result2 {
		t.Error("Different pipeline orders should produce different results")
	}
}

func TestPipelineIndependence(t *testing.T) {
	// Test that pipelines are independent
	base := NewPipeline().Add(Multiply(2))

	pipeline1 := NewPipeline().Add(Multiply(2)).Add(Add(10))
	pipeline2 := NewPipeline().Add(Multiply(2)).Add(Add(20))

	result1 := pipeline1.Execute(5)
	result2 := pipeline2.Execute(5)

	if result1 != 20 {
		t.Errorf("Pipeline1.Execute(5) = %d, want 20", result1)
	}

	if result2 != 30 {
		t.Errorf("Pipeline2.Execute(5) = %d, want 30", result2)
	}

	// Base pipeline should still work
	baseResult := base.Execute(5)
	if baseResult != 10 {
		t.Errorf("Base pipeline affected by other pipelines")
	}
}

func BenchmarkPipeline(b *testing.B) {
	pipeline := NewPipeline().
		Add(Add(10)).
		Add(Multiply(2)).
		Add(Square()).
		Add(Add(5))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pipeline.Execute(5)
	}
}
