package pipeline_builder

// IntTransform is a function that transforms an integer
type IntTransform func(int) int

// Pipeline chains transformations
type Pipeline struct {
	// TODO(human): Define fields
}

// NewPipeline creates a new empty pipeline
func NewPipeline() *Pipeline {
	// TODO(human): Implement
	return nil
}

// Add adds a transformation to the pipeline
func (p *Pipeline) Add(transform IntTransform) *Pipeline {
	// TODO(human): Implement
	return nil
}

// Execute runs all transformations on the input value
func (p *Pipeline) Execute(value int) int {
	// TODO(human): Implement
	return 0
}

// Add returns a transform that adds n to the input
func Add(n int) IntTransform {
	// TODO(human): Implement
	return nil
}

// Multiply returns a transform that multiplies the input by n
func Multiply(n int) IntTransform {
	// TODO(human): Implement
	return nil
}

// Square returns a transform that squares the input
func Square() IntTransform {
	// TODO(human): Implement
	return nil
}
