package wordcount

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name string
		text string
		want map[string]int
	}{
		{
			"Repeated words",
			"hello world hello",
			map[string]int{"hello": 2, "world": 1},
		},
		{
			"All same word",
			"go go go",
			map[string]int{"go": 3},
		},
		{
			"Single word",
			"hello",
			map[string]int{"hello": 1},
		},
		{
			"Empty string",
			"",
			map[string]int{},
		},
		{
			"Multiple different words",
			"the quick brown fox jumps over the lazy dog",
			map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWords(tt.text)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountWords(%q) = %v, want %v", tt.text, got, tt.want)
			}
		})
	}
}
