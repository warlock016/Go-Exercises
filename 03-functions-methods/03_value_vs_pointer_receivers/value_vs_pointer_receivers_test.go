package value_vs_pointer_receivers

import (
	"math"
	"testing"
)

func TestCounterValue(t *testing.T) {
	c := Counter{count: 5}
	if got := c.Value(); got != 5 {
		t.Errorf("Counter.Value() = %d, want 5", got)
	}

	// Test zero value
	var c2 Counter
	if got := c2.Value(); got != 0 {
		t.Errorf("Counter.Value() = %d, want 0", got)
	}
}

func TestCounterIncrement(t *testing.T) {
	c := Counter{count: 0}

	c.Increment()
	if got := c.Value(); got != 1 {
		t.Errorf("After Increment(), Value() = %d, want 1", got)
	}

	c.Increment()
	c.Increment()
	if got := c.Value(); got != 3 {
		t.Errorf("After 3 Increments, Value() = %d, want 3", got)
	}
}

func TestCounterAdd(t *testing.T) {
	tests := []struct {
		name     string
		initial  int
		add      int
		expected int
	}{
		{"add positive", 0, 5, 5},
		{"add to existing", 10, 3, 13},
		{"add negative", 10, -3, 7},
		{"add zero", 5, 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Counter{count: tt.initial}
			c.Add(tt.add)
			if got := c.Value(); got != tt.expected {
				t.Errorf("After Add(%d), Value() = %d, want %d", tt.add, got, tt.expected)
			}
		})
	}
}

func TestCounterReset(t *testing.T) {
	c := Counter{count: 42}
	c.Reset()
	if got := c.Value(); got != 0 {
		t.Errorf("After Reset(), Value() = %d, want 0", got)
	}

	// Reset already zero counter
	c.Reset()
	if got := c.Value(); got != 0 {
		t.Errorf("After Reset() on zero counter, Value() = %d, want 0", got)
	}
}

func TestCounterMutability(t *testing.T) {
	// Test that pointer receiver methods actually modify the counter
	c := Counter{count: 0}
	c.Increment()
	c.Add(5)
	c.Increment()

	if got := c.Value(); got != 7 {
		t.Errorf("After Increment, Add(5), Increment, Value() = %d, want 7", got)
	}

	c.Reset()
	if got := c.Value(); got != 0 {
		t.Errorf("After Reset(), Value() = %d, want 0", got)
	}
}

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		name   string
		width  float64
		height float64
		want   float64
	}{
		{"square", 5, 5, 25},
		{"rectangle", 10, 5, 50},
		{"decimal", 3.5, 2.5, 8.75},
		{"zero width", 0, 5, 0},
		{"zero height", 5, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{Width: tt.width, Height: tt.height}
			got := r.Area()
			if math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("Rectangle{%f, %f}.Area() = %f, want %f",
					tt.width, tt.height, got, tt.want)
			}
		})
	}
}

func TestRectanglePerimeter(t *testing.T) {
	tests := []struct {
		name   string
		width  float64
		height float64
		want   float64
	}{
		{"square", 5, 5, 20},
		{"rectangle", 10, 5, 30},
		{"decimal", 3.5, 2.5, 12},
		{"zero dimensions", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{Width: tt.width, Height: tt.height}
			got := r.Perimeter()
			if math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("Rectangle{%f, %f}.Perimeter() = %f, want %f",
					tt.width, tt.height, got, tt.want)
			}
		})
	}
}

func TestRectangleScale(t *testing.T) {
	tests := []struct {
		name       string
		width      float64
		height     float64
		factor     float64
		wantWidth  float64
		wantHeight float64
	}{
		{"double", 10, 5, 2, 20, 10},
		{"half", 10, 5, 0.5, 5, 2.5},
		{"triple", 3, 4, 3, 9, 12},
		{"identity", 10, 5, 1, 10, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{Width: tt.width, Height: tt.height}
			r.Scale(tt.factor)

			if math.Abs(r.Width-tt.wantWidth) > 0.0001 {
				t.Errorf("After Scale(%f), Width = %f, want %f",
					tt.factor, r.Width, tt.wantWidth)
			}
			if math.Abs(r.Height-tt.wantHeight) > 0.0001 {
				t.Errorf("After Scale(%f), Height = %f, want %f",
					tt.factor, r.Height, tt.wantHeight)
			}
		})
	}
}

func TestRectangleSetDimensions(t *testing.T) {
	r := Rectangle{Width: 10, Height: 5}
	r.SetDimensions(20, 15)

	if r.Width != 20 {
		t.Errorf("After SetDimensions(20, 15), Width = %f, want 20", r.Width)
	}
	if r.Height != 15 {
		t.Errorf("After SetDimensions(20, 15), Height = %f, want 15", r.Height)
	}

	// Test zero values
	r.SetDimensions(0, 0)
	if r.Width != 0 || r.Height != 0 {
		t.Errorf("After SetDimensions(0, 0), got (%f, %f), want (0, 0)",
			r.Width, r.Height)
	}
}

func TestRectangleMutability(t *testing.T) {
	// Test that methods work together correctly
	r := Rectangle{Width: 10, Height: 5}

	originalArea := r.Area()
	if originalArea != 50 {
		t.Errorf("Initial Area() = %f, want 50", originalArea)
	}

	r.Scale(2)
	scaledArea := r.Area()
	if scaledArea != 200 {
		t.Errorf("After Scale(2), Area() = %f, want 200", scaledArea)
	}

	r.SetDimensions(8, 4)
	newArea := r.Area()
	if newArea != 32 {
		t.Errorf("After SetDimensions(8, 4), Area() = %f, want 32", newArea)
	}
}
