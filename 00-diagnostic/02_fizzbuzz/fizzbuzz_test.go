package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			"First 15 numbers",
			15,
			[]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
		{
			"First 5 numbers",
			5,
			[]string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			"Just 1",
			1,
			[]string{"1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FizzBuzz(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FizzBuzz(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
