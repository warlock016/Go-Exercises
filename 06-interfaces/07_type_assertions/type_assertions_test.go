package type_assertions

import (
	"math"
	"testing"
)

const epsilon = 0.0001

func TestCircleArea(t *testing.T) {
	tests := []struct {
		name   string
		circle Circle
		want   float64
	}{
		{"Radius 5", Circle{Radius: 5}, math.Pi * 25},
		{"Radius 1", Circle{Radius: 1}, math.Pi},
		{"Radius 0", Circle{Radius: 0}, 0},
		{"Radius 10", Circle{Radius: 10}, math.Pi * 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.circle.Area()
			if math.Abs(got-tt.want) > epsilon {
				t.Errorf("Circle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		name string
		rect Rectangle
		want float64
	}{
		{"3x4", Rectangle{Width: 3, Height: 4}, 12},
		{"5x5", Rectangle{Width: 5, Height: 5}, 25},
		{"1x10", Rectangle{Width: 1, Height: 10}, 10},
		{"0x5", Rectangle{Width: 0, Height: 5}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.rect.Area()
			if got != tt.want {
				t.Errorf("Rectangle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCirclePerimeter(t *testing.T) {
	tests := []struct {
		name   string
		circle Circle
		want   float64
	}{
		{"Radius 0", Circle{Radius: 0}, math.Pi * 2 * 0},
		{"Radius 1", Circle{Radius: 1}, math.Pi * 2 * 1},
		{"Radius 5", Circle{Radius: 5}, math.Pi * 2 * 5},
		{"Radius 10", Circle{Radius: 10}, math.Pi * 2 * 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.circle.Perimeter()
			if math.Abs(got-tt.want) > epsilon {
				t.Errorf("Circle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectanglePerimeter(t *testing.T) {
	tests := []struct {
		name string
		rect Rectangle
		want float64
	}{
		{"3x4", Rectangle{Width: 3, Height: 4}, 2 * (3 + 4)},
		{"5x5", Rectangle{Width: 5, Height: 5}, 2 * (5 + 5)},
		{"1x10", Rectangle{Width: 1, Height: 10}, 2 * (1 + 10)},
		{"0x5", Rectangle{Width: 0, Height: 5}, 2 * (0 + 5)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.rect.Perimeter()
			if math.Abs(got-tt.want) > epsilon {
				t.Errorf("Circle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiameter(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Circle radius 5", Circle{Radius: 5}, 10},
		{"Circle radius 1", Circle{Radius: 1}, 2},
		{"Circle radius 0", Circle{Radius: 0}, 0},
		{"Rectangle 3x4", Rectangle{Width: 3, Height: 4}, 0},
		{"Rectangle 5x5", Rectangle{Width: 5, Height: 5}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Diameter(tt.shape)
			if got != tt.want {
				t.Errorf("Diameter(%T) = %v, want %v", tt.shape, got, tt.want)
			}
		})
	}
}

func TestGetRadius(t *testing.T) {
	tests := []struct {
		name       string
		shape      Shape
		wantRadius float64
		wantOK     bool
	}{
		{"Circle radius 5", Circle{Radius: 5}, 5, true},
		{"Circle radius 1", Circle{Radius: 1}, 1, true},
		{"Circle radius 0", Circle{Radius: 0}, 0, true},
		{"Rectangle 3x4", Rectangle{Width: 3, Height: 4}, 0, false},
		{"Rectangle 5x5", Rectangle{Width: 5, Height: 5}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRadius, gotOK := GetRadius(tt.shape)

			if gotOK != tt.wantOK {
				t.Errorf("GetRadius(%T) ok = %v, want %v", tt.shape, gotOK, tt.wantOK)
			}

			if gotRadius != tt.wantRadius {
				t.Errorf("GetRadius(%T) radius = %v, want %v", tt.shape, gotRadius, tt.wantRadius)
			}
		})
	}
}

func TestImplementsShape(t *testing.T) {
	var _ Shape = Circle{}
	var _ Shape = Rectangle{}
	t.Log("✓ Both Circle and Rectangle implement Shape interface")
}

func TestTypeAssertionSafety(t *testing.T) {
	var s Shape = Rectangle{Width: 3, Height: 4}

	// This should NOT panic - using comma-ok idiom
	if _, ok := s.(Circle); ok {
		t.Error("Rectangle should not be assertable as Circle")
	}

	// This SHOULD work
	if _, ok := s.(Rectangle); !ok {
		t.Error("Rectangle should be assertable as Rectangle")
	}
}
