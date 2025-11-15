package variables

// DeclareInteger declares and returns an integer with value 42
func DeclareInteger() int {
	// TODO(human): Declare and return an int variable with value 42
	var number int = 42
	return number
}

// DeclareFloat declares and returns a float64 with value 3.14
func DeclareFloat() float64 {
	// TODO(human): Declare and return a float64 variable with value 3.14
	var number float64 = 3.14
	return number
}

// DeclareString declares and returns a string with value "Hello, Go!"
func DeclareString() string {
	// TODO(human): Declare and return a string variable with value "Hello, Go!"
	var text string = "Hello, Go!"
	return text
}

// DeclareBoolean declares and returns a bool with value true
func DeclareBoolean() bool {
	// TODO(human): Declare and return a bool variable with value true
	var boolVar bool = true
	return boolVar
}

// DeclareMultiple declares and returns multiple variables:
// name: "Alice", age: 25, height: 5.6
func DeclareMultiple() (string, int, float64) {
	// TODO(human): Declare and return name, age, and height
	var name string = "Alice"
	var age int = 25
	var height float64 = 5.6

	return name, age, height
}
