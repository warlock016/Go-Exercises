package map_fundamentals

import (
	"sort"
	"testing"
)

func TestCreateMapLiteral(t *testing.T) {
	got := CreateMapLiteral()

	if got == nil {
		t.Fatal("CreateMapLiteral() returned nil")
	}

	want := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	if len(got) != len(want) {
		t.Fatalf("CreateMapLiteral() has %d elements, want %d", len(got), len(want))
	}

	for key, wantVal := range want {
		gotVal, ok := got[key]
		if !ok {
			t.Errorf("CreateMapLiteral() missing key %q", key)
		}
		if gotVal != wantVal {
			t.Errorf("CreateMapLiteral()[%q] = %d, want %d", key, gotVal, wantVal)
		}
	}
}

func TestCreateMapWithMake(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"zero capacity", 0},
		{"small capacity", 10},
		{"medium capacity", 100},
		{"large capacity", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateMapWithMake(tt.size)
			if got == nil {
				t.Fatalf("CreateMapWithMake(%d) returned nil", tt.size)
			}
			if len(got) != 0 {
				t.Errorf("CreateMapWithMake(%d) has %d elements, want 0", tt.size, len(got))
			}
		})
	}
}

func TestGetValue(t *testing.T) {
	tests := []struct {
		name      string
		m         map[string]int
		key       string
		wantVal   int
		wantFound bool
	}{
		{"key exists", map[string]int{"age": 25}, "age", 25, true},
		{"key doesn't exist", map[string]int{"age": 25}, "name", 0, false},
		{"empty map", map[string]int{}, "key", 0, false},
		{"zero value exists", map[string]int{"count": 0}, "count", 0, true},
		{"multiple keys, find first", map[string]int{"a": 1, "b": 2, "c": 3}, "a", 1, true},
		{"multiple keys, find last", map[string]int{"a": 1, "b": 2, "c": 3}, "c", 3, true},
		{"multiple keys, not found", map[string]int{"a": 1, "b": 2, "c": 3}, "d", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, gotFound := GetValue(tt.m, tt.key)
			if gotVal != tt.wantVal {
				t.Errorf("GetValue(%v, %q) value = %d, want %d", tt.m, tt.key, gotVal, tt.wantVal)
			}
			if gotFound != tt.wantFound {
				t.Errorf("GetValue(%v, %q) found = %v, want %v", tt.m, tt.key, gotFound, tt.wantFound)
			}
		})
	}
}

func TestSetValue(t *testing.T) {
	tests := []struct {
		name  string
		m     map[string]int
		key   string
		value int
	}{
		{"set in empty map", make(map[string]int), "key", 42},
		{"set new key", map[string]int{"a": 1}, "b", 2},
		{"update existing key", map[string]int{"key": 1}, "key", 2},
		{"set zero value", make(map[string]int), "zero", 0},
		{"set negative value", make(map[string]int), "neg", -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetValue(tt.m, tt.key, tt.value)
			got, ok := tt.m[tt.key]
			if !ok {
				t.Errorf("SetValue(%v, %q, %d) did not set the key", tt.m, tt.key, tt.value)
			}
			if got != tt.value {
				t.Errorf("After SetValue(%q, %d), got %d", tt.key, tt.value, got)
			}
		})
	}
}

func TestDeleteKey(t *testing.T) {
	tests := []struct {
		name       string
		m          map[string]int
		key        string
		wantRemain map[string]int
	}{
		{
			"delete existing key",
			map[string]int{"a": 1, "b": 2, "c": 3},
			"b",
			map[string]int{"a": 1, "c": 3},
		},
		{
			"delete non-existent key",
			map[string]int{"a": 1},
			"b",
			map[string]int{"a": 1},
		},
		{
			"delete from empty map",
			map[string]int{},
			"key",
			map[string]int{},
		},
		{
			"delete only key",
			map[string]int{"only": 1},
			"only",
			map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteKey(tt.m, tt.key)
			if len(tt.m) != len(tt.wantRemain) {
				t.Fatalf("After DeleteKey, map has %d elements, want %d", len(tt.m), len(tt.wantRemain))
			}
			for key, wantVal := range tt.wantRemain {
				gotVal, ok := tt.m[key]
				if !ok {
					t.Errorf("After DeleteKey, map missing key %q", key)
				}
				if gotVal != wantVal {
					t.Errorf("After DeleteKey, map[%q] = %d, want %d", key, gotVal, wantVal)
				}
			}
		})
	}
}

func TestGetKeys(t *testing.T) {
	tests := []struct {
		name     string
		m        map[string]int
		wantKeys []string
	}{
		{"empty map", map[string]int{}, []string{}},
		{"single key", map[string]int{"a": 1}, []string{"a"}},
		{"multiple keys", map[string]int{"x": 1, "y": 2, "z": 3}, []string{"x", "y", "z"}},
		{"five keys", map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, []string{"a", "b", "c", "d", "e"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetKeys(tt.m)
			if len(got) != len(tt.wantKeys) {
				t.Fatalf("GetKeys(%v) returned %d keys, want %d", tt.m, len(got), len(tt.wantKeys))
			}

			// Sort both slices since map iteration order is random
			sort.Strings(got)
			sort.Strings(tt.wantKeys)

			for i := range tt.wantKeys {
				if got[i] != tt.wantKeys[i] {
					t.Errorf("GetKeys(%v)[%d] = %q, want %q", tt.m, i, got[i], tt.wantKeys[i])
				}
			}
		})
	}
}

func TestCountOccurrences(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  map[string]int
	}{
		{
			"empty slice",
			[]string{},
			map[string]int{},
		},
		{
			"single word",
			[]string{"hello"},
			map[string]int{"hello": 1},
		},
		{
			"all unique",
			[]string{"cat", "dog", "bird"},
			map[string]int{"cat": 1, "dog": 1, "bird": 1},
		},
		{
			"some duplicates",
			[]string{"cat", "dog", "cat", "bird", "dog", "cat"},
			map[string]int{"cat": 3, "dog": 2, "bird": 1},
		},
		{
			"all same",
			[]string{"test", "test", "test", "test"},
			map[string]int{"test": 4},
		},
		{
			"mixed case",
			[]string{"Apple", "apple", "APPLE"},
			map[string]int{"Apple": 1, "apple": 1, "APPLE": 1},
		},
		{
			"complex example",
			[]string{"go", "is", "fun", "go", "is", "powerful", "go", "go"},
			map[string]int{"go": 4, "is": 2, "fun": 1, "powerful": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountOccurrences(tt.words)
			if len(got) != len(tt.want) {
				t.Fatalf("CountOccurrences(%v) returned %d keys, want %d", tt.words, len(got), len(tt.want))
			}
			for key, wantCount := range tt.want {
				gotCount, ok := got[key]
				if !ok {
					t.Errorf("CountOccurrences(%v) missing key %q", tt.words, key)
				}
				if gotCount != wantCount {
					t.Errorf("CountOccurrences(%v)[%q] = %d, want %d", tt.words, key, gotCount, wantCount)
				}
			}
		})
	}
}
