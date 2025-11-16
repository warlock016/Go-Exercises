package map_patterns

import (
	"reflect"
	"sort"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name string
		text string
		want map[string]int
	}{
		{"simple words", "hello world", map[string]int{"hello": 1, "world": 1}},
		{"repeated words", "hello world hello", map[string]int{"hello": 2, "world": 1}},
		{"single word", "hello", map[string]int{"hello": 1}},
		{"empty string", "", map[string]int{}},
		{"multiple spaces", "hello  world", map[string]int{"hello": 1, "world": 1}},
		{"many repeats", "a b a c a b", map[string]int{"a": 3, "b": 2, "c": 1}},
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

func TestGroupByLength(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  map[int][]string
	}{
		{
			"mixed lengths",
			[]string{"cat", "elephant", "dog", "ant"},
			map[int][]string{3: {"cat", "dog", "ant"}, 8: {"elephant"}},
		},
		{
			"all same length",
			[]string{"cat", "dog", "ant"},
			map[int][]string{3: {"cat", "dog", "ant"}},
		},
		{
			"empty slice",
			[]string{},
			map[int][]string{},
		},
		{
			"single word",
			[]string{"hello"},
			map[int][]string{5: {"hello"}},
		},
		{
			"various lengths",
			[]string{"a", "bb", "ccc", "dd", "e"},
			map[int][]string{1: {"a", "e"}, 2: {"bb", "dd"}, 3: {"ccc"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupByLength(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupByLength(%v) = %v, want %v", tt.words, got, tt.want)
			}
		})
	}
}

func TestInvertMap(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		want map[int][]string
	}{
		{
			"unique values",
			map[string]int{"a": 1, "b": 2, "c": 3},
			map[int][]string{1: {"a"}, 2: {"b"}, 3: {"c"}},
		},
		{
			"duplicate values",
			map[string]int{"a": 1, "b": 2, "c": 1},
			map[int][]string{1: {"a", "c"}, 2: {"b"}},
		},
		{
			"all same value",
			map[string]int{"a": 1, "b": 1, "c": 1},
			map[int][]string{1: {"a", "b", "c"}},
		},
		{
			"empty map",
			map[string]int{},
			map[int][]string{},
		},
		{
			"single entry",
			map[string]int{"key": 42},
			map[int][]string{42: {"key"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InvertMap(tt.m)

			// Need to sort slices for comparison since map iteration is random
			for key := range got {
				sort.Strings(got[key])
			}
			for key := range tt.want {
				sort.Strings(tt.want[key])
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvertMap(%v) = %v, want %v", tt.m, got, tt.want)
			}
		})
	}
}

func TestMergeMaps(t *testing.T) {
	tests := []struct {
		name string
		m1   map[string]int
		m2   map[string]int
		want map[string]int
	}{
		{
			"no overlap",
			map[string]int{"a": 1, "b": 2},
			map[string]int{"c": 3, "d": 4},
			map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
		},
		{
			"with overlap",
			map[string]int{"a": 1, "b": 2},
			map[string]int{"b": 3, "c": 4},
			map[string]int{"a": 1, "b": 3, "c": 4},
		},
		{
			"m1 empty",
			map[string]int{},
			map[string]int{"a": 1, "b": 2},
			map[string]int{"a": 1, "b": 2},
		},
		{
			"m2 empty",
			map[string]int{"a": 1, "b": 2},
			map[string]int{},
			map[string]int{"a": 1, "b": 2},
		},
		{
			"both empty",
			map[string]int{},
			map[string]int{},
			map[string]int{},
		},
		{
			"complete overlap",
			map[string]int{"a": 1, "b": 2},
			map[string]int{"a": 10, "b": 20},
			map[string]int{"a": 10, "b": 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeMaps(tt.m1, tt.m2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeMaps(%v, %v) = %v, want %v", tt.m1, tt.m2, got, tt.want)
			}
		})
	}
}

func TestMapKeys(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		want []string
	}{
		{"three keys", map[string]int{"z": 1, "a": 2, "m": 3}, []string{"a", "m", "z"}},
		{"single key", map[string]int{"only": 1}, []string{"only"}},
		{"empty map", map[string]int{}, []string{}},
		{"alphabetical", map[string]int{"apple": 1, "banana": 2, "cherry": 3}, []string{"apple", "banana", "cherry"}},
		{"reverse alpha", map[string]int{"z": 1, "y": 2, "x": 3}, []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeys(tt.m)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapKeys(%v) = %v, want %v", tt.m, got, tt.want)
			}
		})
	}
}

func TestMapValues(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
	}{
		{"three values", map[string]int{"a": 1, "b": 2, "c": 3}},
		{"single value", map[string]int{"only": 42}},
		{"empty map", map[string]int{}},
		{"duplicate values", map[string]int{"a": 1, "b": 1, "c": 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapValues(tt.m)

			// Check length
			if len(got) != len(tt.m) {
				t.Errorf("MapValues(%v) length = %d, want %d", tt.m, len(got), len(tt.m))
			}

			// Check all values are present (order doesn't matter)
			gotMap := make(map[int]int)
			for _, v := range got {
				gotMap[v]++
			}
			wantMap := make(map[int]int)
			for _, v := range tt.m {
				wantMap[v]++
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapValues(%v) values = %v, want %v", tt.m, got, tt.m)
			}
		})
	}
}

func TestFilterMap(t *testing.T) {
	tests := []struct {
		name      string
		m         map[string]int
		predicate func(string, int) bool
		want      map[string]int
	}{
		{
			"values > 5",
			map[string]int{"a": 3, "b": 10, "c": 7},
			func(k string, v int) bool { return v > 5 },
			map[string]int{"b": 10, "c": 7},
		},
		{
			"keys starting with 'a'",
			map[string]int{"apple": 1, "banana": 2, "apricot": 3},
			func(k string, v int) bool { return k[0] == 'a' },
			map[string]int{"apple": 1, "apricot": 3},
		},
		{
			"even values",
			map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			func(k string, v int) bool { return v%2 == 0 },
			map[string]int{"b": 2, "d": 4},
		},
		{
			"all match",
			map[string]int{"a": 1, "b": 2},
			func(k string, v int) bool { return true },
			map[string]int{"a": 1, "b": 2},
		},
		{
			"none match",
			map[string]int{"a": 1, "b": 2},
			func(k string, v int) bool { return false },
			map[string]int{},
		},
		{
			"empty map",
			map[string]int{},
			func(k string, v int) bool { return true },
			map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterMap(tt.m, tt.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterMap(%v) = %v, want %v", tt.m, got, tt.want)
			}
		})
	}
}
