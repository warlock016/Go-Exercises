package implicit_satisfaction

// TODO(human): Define the Speaker interface

type Speaker interface {
	Speak() string
}

// TODO(human): Define Dog, Cat, and Robot types
type Dog struct{}
type Cat struct{}
type Robot struct{}

// TODO(human): Implement Speak() method for Dog
func (d Dog) Speak() string {
	return "Woof!"
}

// TODO(human): Implement Speak() method for Cat
func (c Cat) Speak() string {
	return "Meow!"
}

// TODO(human): Implement Speak() method for Robot
func (r Robot) Speak() string {
	return "Beep boop!"
}

// MakeSpeak calls the Speak method on any Speaker
func MakeSpeak(s Speaker) string {
	// TODO(human): Implement
	return s.Speak()
}
