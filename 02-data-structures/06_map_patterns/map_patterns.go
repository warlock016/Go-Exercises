package map_patterns

import (
	"maps"
	"slices"
	"strings"
)

// CountWords returns word frequency map from text
func CountWords(text string) map[string]int {
	// TODO(human): Return map of word counts from text

	words := strings.Fields(text)
	counts := map[string]int{}

	for _, w := range words {
		counts[w]++
	}

	return counts
}

// GroupByLength groups strings by their length
func GroupByLength(words []string) map[int][]string {
	// TODO(human): Return map grouping words by their length

	result := map[int][]string{}
	// charCount := map[]

	for _, v := range words {
		charCount := 0

		for range v {
			charCount++
		}

		result[charCount] = append(result[charCount], v)

	}
	return result
}

// InvertMap inverts keys/values (values become keys, original keys in slice)
func InvertMap(m map[string]int) map[int][]string {
	// TODO(human): Return map with inverted keys and values

	result := map[int][]string{}

	for k, v := range m {
		result[v] = append(result[v], k)
	}

	return result
}

// MergeMaps combines two maps (m2 values overwrite m1 on conflicts)
func MergeMaps(m1, m2 map[string]int) map[string]int {
	// TODO(human): Return new map combining m1 and m2

	result := map[string]int{}

	maps.Copy(result, m1)
	maps.Copy(result, m2)

	return result
}

// MapKeys returns sorted slice of all keys
func MapKeys(m map[string]int) []string {
	// TODO(human): Return sorted slice of all keys from map

	result := []string{}

	for k := range m {
		result = append(result, k)
	}

	slices.Sort(result)

	return result
}

// MapValues returns slice of all values (order unspecified)
func MapValues(m map[string]int) []int {
	// TODO(human): Return slice of all values from map
	result := []int{}

	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// FilterMap returns new map with entries matching predicate
func FilterMap(m map[string]int, predicate func(string, int) bool) map[string]int {
	// TODO(human): Return new map containing only entries where predicate returns true

	result := map[string]int{}

	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}
