package struct_basics

import "math"

// TODO(human): Import required packages

// TODO(human): Define a Person struct with Name (string) and Age (int) fields
type Person struct {
	Name string
	Age  int
}

// TODO(human): Define a Point struct with X and Y (int) fields
type Point struct {
	X int
	Y int
}

// TODO(human): Define a Book struct with Title, Author (string), Pages (int), and Available (bool) fields
type Book struct {
	Title, Author string
	Pages         int
	Available     bool
}

// CreatePerson creates and returns a Person with the given name and age
func CreatePerson(name string, age int) Person {

	p := Person{name, age}

	return p
}

// CreateZeroPerson returns a zero-valued Person (empty name, age 0)
func CreateZeroPerson() Person {

	p := Person{}
	return p
}

// UpdatePersonAge returns a new Person with the age updated
func UpdatePersonAge(p Person, newAge int) Person {
	p.Age = newAge
	return p
}

// ComparePersons returns true if both persons have the same name and age
func ComparePersons(p1, p2 Person) bool {
	return p1 == p2
}

// CreatePoint creates and returns a Point with the given x and y coordinates
func CreatePoint(x, y int) Point {
	return Point{x, y}
}

// DistanceFromOrigin calculates the Euclidean distance from the origin (0, 0)
// Formula: sqrt(x^2 + y^2)
func DistanceFromOrigin(p Point) float64 {
	return math.Sqrt(math.Pow(float64(p.X), 2) + math.Pow(float64(p.Y), 2))
}

// CreateBook creates and returns a Book with Available set to true by default
func CreateBook(title, author string, pages int) Book {
	return Book{title, author, pages, true}
}

// IsBookAvailable returns the availability status of the book
func IsBookAvailable(b Book) bool {
	return b.Available
}
