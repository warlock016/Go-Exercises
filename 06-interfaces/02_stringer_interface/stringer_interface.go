package stringer_interface

import "fmt"

// TODO(human): Define Temperature struct
type Temperature struct {
	Value float64
	Unit  string
}

// TODO(human): Define Point struct
type Point struct {
	X, Y int
}

// TODO(human): Define Duration struct
type Duration struct {
	Hours, Minutes int
}

// TODO(human): Implement String() method for Temperature
func (t Temperature) String() string {
	return fmt.Sprintf("%.1f°%s", t.Value, t.Unit)
}

// TODO(human): Implement String() method for Point
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// TODO(human): Implement String() method for Duration
func (d Duration) String() string {
	return fmt.Sprintf("%dh %dm", d.Hours, d.Minutes)
}
