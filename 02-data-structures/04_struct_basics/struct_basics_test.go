package struct_basics

import (
	"math"
	"testing"
)

func TestCreatePerson(t *testing.T) {
	tests := []struct {
		name     string
		nameParm string
		age      int
	}{
		{"basic person", "Alice", 30},
		{"empty name", "", 25},
		{"zero age", "Bob", 0},
		{"negative age", "Charlie", -5},
		{"large age", "Dave", 150},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreatePerson(tt.nameParm, tt.age)
			if got.Name != tt.nameParm {
				t.Errorf("CreatePerson(%q, %d).Name = %q, want %q", tt.nameParm, tt.age, got.Name, tt.nameParm)
			}
			if got.Age != tt.age {
				t.Errorf("CreatePerson(%q, %d).Age = %d, want %d", tt.nameParm, tt.age, got.Age, tt.age)
			}
		})
	}
}

func TestCreateZeroPerson(t *testing.T) {
	got := CreateZeroPerson()
	if got.Name != "" {
		t.Errorf("CreateZeroPerson().Name = %q, want empty string", got.Name)
	}
	if got.Age != 0 {
		t.Errorf("CreateZeroPerson().Age = %d, want 0", got.Age)
	}
}

func TestUpdatePersonAge(t *testing.T) {
	tests := []struct {
		name   string
		person Person
		newAge int
	}{
		{"update from zero", Person{Name: "Alice", Age: 0}, 30},
		{"increase age", Person{Name: "Bob", Age: 25}, 26},
		{"decrease age", Person{Name: "Charlie", Age: 50}, 30},
		{"set to zero", Person{Name: "Dave", Age: 40}, 0},
		{"negative age", Person{Name: "Eve", Age: 30}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := tt.person
			got := UpdatePersonAge(tt.person, tt.newAge)

			// Check the returned person has updated age
			if got.Age != tt.newAge {
				t.Errorf("UpdatePersonAge(%v, %d).Age = %d, want %d", tt.person, tt.newAge, got.Age, tt.newAge)
			}

			// Check the name is preserved
			if got.Name != tt.person.Name {
				t.Errorf("UpdatePersonAge(%v, %d).Name = %q, want %q", tt.person, tt.newAge, got.Name, tt.person.Name)
			}

			// Check original is not modified (value semantics)
			if tt.person != original {
				t.Errorf("UpdatePersonAge modified the original person: got %v, want %v", tt.person, original)
			}
		})
	}
}

func TestComparePersons(t *testing.T) {
	tests := []struct {
		name string
		p1   Person
		p2   Person
		want bool
	}{
		{"identical persons", Person{Name: "Alice", Age: 30}, Person{Name: "Alice", Age: 30}, true},
		{"different names", Person{Name: "Alice", Age: 30}, Person{Name: "Bob", Age: 30}, false},
		{"different ages", Person{Name: "Alice", Age: 30}, Person{Name: "Alice", Age: 31}, false},
		{"both zero", Person{}, Person{}, true},
		{"completely different", Person{Name: "Alice", Age: 30}, Person{Name: "Bob", Age: 25}, false},
		{"empty name same age", Person{Name: "", Age: 25}, Person{Name: "", Age: 25}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ComparePersons(tt.p1, tt.p2)
			if got != tt.want {
				t.Errorf("ComparePersons(%v, %v) = %v, want %v", tt.p1, tt.p2, got, tt.want)
			}
		})
	}
}

func TestCreatePoint(t *testing.T) {
	tests := []struct {
		name  string
		x     int
		y     int
		wantX int
		wantY int
	}{
		{"origin", 0, 0, 0, 0},
		{"positive coordinates", 3, 4, 3, 4},
		{"negative coordinates", -3, -4, -3, -4},
		{"mixed coordinates", 5, -5, 5, -5},
		{"large coordinates", 1000, 2000, 1000, 2000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreatePoint(tt.x, tt.y)
			if got.X != tt.wantX {
				t.Errorf("CreatePoint(%d, %d).X = %d, want %d", tt.x, tt.y, got.X, tt.wantX)
			}
			if got.Y != tt.wantY {
				t.Errorf("CreatePoint(%d, %d).Y = %d, want %d", tt.x, tt.y, got.Y, tt.wantY)
			}
		})
	}
}

func TestDistanceFromOrigin(t *testing.T) {
	tests := []struct {
		name  string
		point Point
		want  float64
	}{
		{"origin", Point{X: 0, Y: 0}, 0.0},
		{"3-4-5 triangle", Point{X: 3, Y: 4}, 5.0},
		{"on x-axis", Point{X: 5, Y: 0}, 5.0},
		{"on y-axis", Point{X: 0, Y: 12}, 12.0},
		{"5-12-13 triangle", Point{X: 5, Y: 12}, 13.0},
		{"negative coordinates", Point{X: -3, Y: -4}, 5.0},
		{"mixed signs", Point{X: 8, Y: -6}, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DistanceFromOrigin(tt.point)
			if math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("DistanceFromOrigin(%v) = %v, want %v", tt.point, got, tt.want)
			}
		})
	}
}

func TestCreateBook(t *testing.T) {
	tests := []struct {
		name       string
		title      string
		author     string
		pages      int
		wantAvail  bool
	}{
		{"classic novel", "1984", "George Orwell", 328, true},
		{"technical book", "The Go Programming Language", "Donovan & Kernighan", 380, true},
		{"empty title", "", "Unknown", 100, true},
		{"zero pages", "Empty Book", "Nobody", 0, true},
		{"large book", "War and Peace", "Leo Tolstoy", 1225, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateBook(tt.title, tt.author, tt.pages)
			if got.Title != tt.title {
				t.Errorf("CreateBook(%q, %q, %d).Title = %q, want %q", tt.title, tt.author, tt.pages, got.Title, tt.title)
			}
			if got.Author != tt.author {
				t.Errorf("CreateBook(%q, %q, %d).Author = %q, want %q", tt.title, tt.author, tt.pages, got.Author, tt.author)
			}
			if got.Pages != tt.pages {
				t.Errorf("CreateBook(%q, %q, %d).Pages = %d, want %d", tt.title, tt.author, tt.pages, got.Pages, tt.pages)
			}
			if got.Available != tt.wantAvail {
				t.Errorf("CreateBook(%q, %q, %d).Available = %v, want %v", tt.title, tt.author, tt.pages, got.Available, tt.wantAvail)
			}
		})
	}
}

func TestIsBookAvailable(t *testing.T) {
	tests := []struct {
		name string
		book Book
		want bool
	}{
		{
			"available book",
			Book{Title: "1984", Author: "Orwell", Pages: 328, Available: true},
			true,
		},
		{
			"unavailable book",
			Book{Title: "1984", Author: "Orwell", Pages: 328, Available: false},
			false,
		},
		{
			"zero value book",
			Book{},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsBookAvailable(tt.book)
			if got != tt.want {
				t.Errorf("IsBookAvailable(%v) = %v, want %v", tt.book, got, tt.want)
			}
		})
	}
}
