package stringer_interface

import (
	"fmt"
	"testing"
)

func TestTemperatureString(t *testing.T) {
	tests := []struct {
		name string
		temp Temperature
		want string
	}{
		{"Fahrenheit", Temperature{Value: 72.5, Unit: "F"}, "72.5°F"},
		{"Celsius", Temperature{Value: 22.0, Unit: "C"}, "22.0°C"},
		{"Freezing", Temperature{Value: 0.0, Unit: "C"}, "0.0°C"},
		{"Negative", Temperature{Value: -10.5, Unit: "C"}, "-10.5°C"},
		{"Hot", Temperature{Value: 98.6, Unit: "F"}, "98.6°F"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.temp.String()
			if got != tt.want {
				t.Errorf("Temperature.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestPointString(t *testing.T) {
	tests := []struct {
		name  string
		point Point
		want  string
	}{
		{"Origin", Point{X: 0, Y: 0}, "(0, 0)"},
		{"Positive", Point{X: 3, Y: 5}, "(3, 5)"},
		{"Negative X", Point{X: -2, Y: 4}, "(-2, 4)"},
		{"Negative Y", Point{X: 7, Y: -3}, "(7, -3)"},
		{"Both negative", Point{X: -1, Y: -8}, "(-1, -8)"},
		{"Large values", Point{X: 100, Y: 200}, "(100, 200)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.point.String()
			if got != tt.want {
				t.Errorf("Point.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDurationString(t *testing.T) {
	tests := []struct {
		name string
		dur  Duration
		want string
	}{
		{"Two and half hours", Duration{Hours: 2, Minutes: 30}, "2h 30m"},
		{"One hour", Duration{Hours: 1, Minutes: 0}, "1h 0m"},
		{"Only minutes", Duration{Hours: 0, Minutes: 45}, "0h 45m"},
		{"Zero", Duration{Hours: 0, Minutes: 0}, "0h 0m"},
		{"Long duration", Duration{Hours: 12, Minutes: 15}, "12h 15m"},
		{"Almost an hour", Duration{Hours: 0, Minutes: 59}, "0h 59m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.dur.String()
			if got != tt.want {
				t.Errorf("Duration.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestStringerInterfaceWithFmt(t *testing.T) {
	// Test that fmt.Println uses our String() method
	temp := Temperature{Value: 72.5, Unit: "F"}
	point := Point{X: 3, Y: 5}
	dur := Duration{Hours: 2, Minutes: 30}

	// These use String() implicitly
	tempStr := fmt.Sprintf("%v", temp)
	if tempStr != "72.5°F" {
		t.Errorf("fmt.Sprintf with Temperature = %q, want %q", tempStr, "72.5°F")
	}

	pointStr := fmt.Sprintf("%s", point)
	if pointStr != "(3, 5)" {
		t.Errorf("fmt.Sprintf with Point = %q, want %q", pointStr, "(3, 5)")
	}

	durStr := fmt.Sprint(dur)
	if durStr != "2h 30m" {
		t.Errorf("fmt.Sprint with Duration = %q, want %q", durStr, "2h 30m")
	}
}

func TestImplementsFmtStringer(t *testing.T) {
	// Verify that our types satisfy fmt.Stringer
	var _ fmt.Stringer = Temperature{}
	var _ fmt.Stringer = Point{}
	var _ fmt.Stringer = Duration{}

	t.Log("✓ All types implement fmt.Stringer interface")
}

func ExampleTemperature_String() {
	temp := Temperature{Value: 72.5, Unit: "F"}
	fmt.Println(temp)
	// Output: 72.5°F
}

func ExamplePoint_String() {
	point := Point{X: 3, Y: 5}
	fmt.Println(point)
	// Output: (3, 5)
}

func ExampleDuration_String() {
	dur := Duration{Hours: 2, Minutes: 30}
	fmt.Println(dur)
	// Output: 2h 30m
}
