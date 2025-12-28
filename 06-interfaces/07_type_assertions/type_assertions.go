package type_assertions

import "math"

// TODO(human): Define Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// TODO(human): Define Circle struct
type Circle struct {
	Radius float64
}

// TODO(human): Define Rectangle struct
type Rectangle struct {
	Height float64
	Width  float64
}

// TODO(human): Implement Area() for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// TODO(human): Implement Area() for Rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

// Diameter returns the diameter if s is a Circle, otherwise returns 0
func Diameter(s Shape) float64 {
	// TODO(human): Implement using type assertion
	if c, ok := s.(Circle); ok {
		return c.Radius * 2
	}
	return 0
}

// GetRadius safely extracts the radius from a Circle
func GetRadius(s Shape) (float64, bool) {
	// TODO(human): Implement using comma-ok idiom
	if c, ok := s.(Circle); ok {
		return c.Radius, true
	}
	return 0, false
}
