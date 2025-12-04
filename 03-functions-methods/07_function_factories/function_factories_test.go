package function_factories

import (
	"math"
	"testing"
)

func TestMakeAdder(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{5, 3, 8},
		{10, -5, 5},
		{0, 42, 42},
		{-5, -3, -8},
	}

	for _, tt := range tests {
		adder := MakeAdder(tt.x)
		got := adder(tt.y)
		if got != tt.want {
			t.Errorf("MakeAdder(%d)(%d) = %d, want %d", tt.x, tt.y, got, tt.want)
		}
	}

	// Test independence
	add5 := MakeAdder(5)
	add10 := MakeAdder(10)
	if add5(3) != 8 || add10(3) != 13 {
		t.Error("Multiple adders should be independent")
	}
}

func TestMakeGreeter(t *testing.T) {
	tests := []struct {
		greeting string
		name     string
		want     string
	}{
		{"Hello", "Alice", "Hello, Alice!"},
		{"Hola", "Bob", "Hola, Bob!"},
		{"Hi", "World", "Hi, World!"},
		{"Greetings", "Friend", "Greetings, Friend!"},
	}

	for _, tt := range tests {
		greeter := MakeGreeter(tt.greeting)
		got := greeter(tt.name)
		if got != tt.want {
			t.Errorf("MakeGreeter(%q)(%q) = %q, want %q", tt.greeting, tt.name, got, tt.want)
		}
	}
}

func TestMakeValidator(t *testing.T) {
	inRange := MakeValidator(1, 10)

	tests := []struct {
		n    int
		want bool
	}{
		{1, true},
		{5, true},
		{10, true},
		{0, false},
		{11, false},
		{-5, false},
		{15, false},
	}

	for _, tt := range tests {
		got := inRange(tt.n)
		if got != tt.want {
			t.Errorf("inRange(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}

	// Test different range
	age := MakeValidator(18, 65)
	if !age(25) || age(10) || age(70) {
		t.Error("Validator with different range failed")
	}
}

func TestMakeFormatter(t *testing.T) {
	tests := []struct {
		prefix string
		suffix string
		input  string
		want   string
	}{
		{"[", "]", "test", "[test]"},
		{"(", ")", "hello", "(hello)"},
		{"<", ">", "tag", "<tag>"},
		{"", "", "plain", "plain"},
		{"**", "**", "bold", "**bold**"},
	}

	for _, tt := range tests {
		formatter := MakeFormatter(tt.prefix, tt.suffix)
		got := formatter(tt.input)
		if got != tt.want {
			t.Errorf("MakeFormatter(%q, %q)(%q) = %q, want %q",
				tt.prefix, tt.suffix, tt.input, got, tt.want)
		}
	}
}

func TestMakePowerFunction(t *testing.T) {
	tests := []struct {
		exponent int
		base     float64
		want     float64
	}{
		{2, 4.0, 16.0},
		{3, 3.0, 27.0},
		{2, 5.0, 25.0},
		{3, 2.0, 8.0},
		{1, 42.0, 42.0},
		{0, 5.0, 1.0},
		{4, 2.0, 16.0},
	}

	for _, tt := range tests {
		power := MakePowerFunction(tt.exponent)
		got := power(tt.base)
		if math.Abs(got-tt.want) > 0.0001 {
			t.Errorf("MakePowerFunction(%d)(%f) = %f, want %f",
				tt.exponent, tt.base, got, tt.want)
		}
	}

	// Test independence
	square := MakePowerFunction(2)
	cube := MakePowerFunction(3)
	if square(4.0) != 16.0 || cube(4.0) != 64.0 {
		t.Error("Multiple power functions should be independent")
	}
}
