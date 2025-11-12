package person

import "testing"

func TestPersonGreet(t *testing.T) {
	tests := []struct {
		name   string
		person Person
		want   string
	}{
		{
			"Alice",
			Person{Name: "Alice", Age: 25},
			"Hello, my name is Alice and I am 25 years old.",
		},
		{
			"Bob",
			Person{Name: "Bob", Age: 30},
			"Hello, my name is Bob and I am 30 years old.",
		},
		{
			"Charlie",
			Person{Name: "Charlie", Age: 18},
			"Hello, my name is Charlie and I am 18 years old.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.person.Greet()
			if got != tt.want {
				t.Errorf("Person.Greet() = %q, want %q", got, tt.want)
			}
		})
	}
}
