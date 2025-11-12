package person

import "fmt"

// TODO(human): Define the Person struct with Name and Age fields

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	greet := fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
	return greet
}

// TODO(human): Implement the Greet method
