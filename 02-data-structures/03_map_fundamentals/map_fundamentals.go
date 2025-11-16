package map_fundamentals

import (
	"slices"
)

// CreateMapLiteral creates and returns a map with keys "one", "two", "three"
// and values 1, 2, 3 respectively
func CreateMapLiteral() map[string]int {

	return map[string]int{"one": 1, "two": 2, "three": 3}
}

// CreateMapWithMake creates an empty map with the specified initial capacity
func CreateMapWithMake(size int) map[string]int {

	return make(map[string]int, size)
}

// GetValue safely retrieves a value from the map using the comma-ok idiom
// Returns the value and true if the key exists, or zero value and false if not
func GetValue(m map[string]int, key string) (int, bool) {

	val, ok := m[key]
	return val, ok
}

// SetValue sets the key to the specified value in the map
func SetValue(m map[string]int, key string, value int) {
	m[key] = value
}

// DeleteKey removes the key from the map
// If the key doesn't exist, this is a no-op (safe to call)
func DeleteKey(m map[string]int, key string) {

	for range m {
		delete(m, key)
	}
}

// GetKeys returns a slice containing all keys from the map
// The order is not guaranteed (maps are unordered)
func GetKeys(m map[string]int) []string {

	// slice := make([]string, len(m))

	// sliceA := make([]string, len(m))
	// sliceB := []string{}
	// fmt.Printf("make %s // empty: %s", sliceA, sliceB)

	slice := []string{}

	for k := range m {
		slice = append(slice, k)
	}

	slices.Sort(slice)

	return slice
}

// CountOccurrences counts how many times each word appears in the slice
// Returns a map where keys are words and values are their frequencies
func CountOccurrences(words []string) map[string]int {

	result := make(map[string]int, len(words))

	for _, v := range words {
		result[v]++
	}
	return result
}
