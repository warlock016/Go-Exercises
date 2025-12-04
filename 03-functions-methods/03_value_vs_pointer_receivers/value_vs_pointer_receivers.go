package value_vs_pointer_receivers

// Counter tracks a count value
type Counter struct {
	count int
}

// Value returns the current count
func (c Counter) Value() int {
	// TODO(human): Implement
	return c.count
}

// Increment increases the count by 1
func (c *Counter) Increment() {
	// TODO(human): Implement
	// newCount := c.count
	// newCount++
	// c.count = newCount
	c.count += 1
}

// Add increases the count by n
func (c *Counter) Add(n int) {
	// TODO(human): Implement
	c.count += n
}

// Reset sets the count back to 0
func (c *Counter) Reset() {
	// TODO(human): Implement
	c.count = 0
}

// Rectangle represents a rectangle with width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of the rectangle
func (r Rectangle) Area() float64 {
	// TODO(human): Implement

	return r.Height * r.Width
}

// Perimeter calculates the perimeter of the rectangle
func (r Rectangle) Perimeter() float64 {
	// TODO(human): Implement
	return 2 * (r.Height + r.Width)
}

// Scale multiplies both dimensions by the given factor
func (r *Rectangle) Scale(factor float64) {
	// TODO(human): Implement
	r.Height *= factor
	r.Width *= factor
}

// SetDimensions updates the width and height
func (r *Rectangle) SetDimensions(width, height float64) {
	// TODO(human): Implement
	r.Height = height
	r.Width = width
}
