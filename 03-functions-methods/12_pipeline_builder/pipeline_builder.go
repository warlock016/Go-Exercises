package pipeline_builder

// IntTransform is a function that transforms an integer
type IntTransform func(int) int

// Pipeline chains transformations
type Pipeline struct {
	// TODO(human): Define fields
	pending []IntTransform
}

// NewPipeline creates a new empty pipeline
func NewPipeline() *Pipeline {
	// TODO(human): Implement
	return &Pipeline{
		pending: make([]IntTransform, 0),
	}
}

// Add adds a transformation to the pipeline
func (p *Pipeline) Add(transform IntTransform) *Pipeline {
	// TODO(human): Implement
	p.pending = append(p.pending, transform)
	return p
}

// Execute runs all transformations on the input value
func (p *Pipeline) Execute(value int) int {
	// TODO(human): Implement
	for _, t := range p.pending {
		value = t(value)
	}
	return value
}

// Add returns a transform that adds n to the input
func Add(n int) IntTransform {
	// TODO(human): Implement
	return func(i int) int {
		return i + n
	}
}

// Multiply returns a transform that multiplies the input by n
func Multiply(n int) IntTransform {
	// TODO(human): Implement
	return func(i int) int {
		return i * n
	}
}

// Square returns a transform that squares the input
func Square() IntTransform {
	// TODO(human): Implement
	return func(i int) int {
		return i * i
	}
}
