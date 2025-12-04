package method_sets

import (
	"slices"
	"strconv"
	"strings"
)

// Temperature represents a temperature in Celsius
type Temperature float64

// Celsius returns the temperature in Celsius
func (t Temperature) Celsius() float64 {
	// TODO(human): Implement
	result := float64(t)
	return result
}

// Fahrenheit converts and returns the temperature in Fahrenheit
func (t Temperature) Fahrenheit() float64 {
	// TODO(human): Implement
	result := (float64(t)*9)/5 + 32
	return result
}

// Kelvin converts and returns the temperature in Kelvin
func (t Temperature) Kelvin() float64 {
	// TODO(human): Implement
	result := float64(t) + 273.15
	return result
}

// StringList is a slice of strings with utility methods
type StringList []string

// Join combines all strings with a separator
func (s StringList) Join(separator string) string {
	// TODO(human): Implement
	var result strings.Builder

	for i, v := range s {
		result.WriteString(v)
		if i < len(s)-1 {
			result.WriteString(separator)
		}
	}

	return result.String()
}

// Contains checks if the list contains the given item
func (s StringList) Contains(item string) bool {
	// TODO(human): Implement
	if len(s) == 0 {
		return false
	}

	return slices.Contains(s, item)
}

// Filter returns a new StringList containing only items matching the predicate
func (s StringList) Filter(predicate func(string) bool) StringList {
	// TODO(human): Implement
	result := StringList{}

	for _, v := range s {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Append adds items to the list
func (s *StringList) Append(items ...string) {
	// TODO(human): Implement

	for _, v := range items {
		*s = append(*s, v)
	}
}

// Bytes represents a number of bytes
type Bytes int64

// String formats the bytes with an appropriate unit
func (b Bytes) String() string {
	// TODO(human): Implement

	var ref Bytes = 1024
	var mult Bytes = ref
	var dec int = 0

	var value float64
	var result strings.Builder

	for b/mult > 0 {
		value = float64(b / mult)
		dec += 1
		mult *= mult
	}

	switch dec {
	case 0:
		result.WriteString(strconv.FormatInt(int64(b), 10))
		result.WriteString(" B")
	case 1:
		// if value > 1000 {
		// 	value /= 1024
		// }
		result.WriteString(strconv.FormatFloat(value, 'f', 2, 64))
		result.WriteString(" KB")
	case 2:
		if value > 1000 {
			value /= 1024
			result.WriteString(strconv.FormatFloat(value, 'f', 2, 64))
			result.WriteString(" GB")
		} else {
			result.WriteString(strconv.FormatFloat(value, 'f', 2, 64))
			result.WriteString(" MB")
		}
	case 3:
		if value > 1000 {
			value /= 1024
		}
		result.WriteString(strconv.FormatFloat(value, 'f', 2, 64))
		result.WriteString(" GB")
	// case 4:
	default:
	}

	return result.String()
}

// Kilobytes returns the value in kilobytes
func (b Bytes) Kilobytes() float64 {
	// TODO(human): Implement
	return float64(b) / 1024
}

// Megabytes returns the value in megabytes
func (b Bytes) Megabytes() float64 {
	// TODO(human): Implement
	return float64(b) / (1024 * 1024)
}

// Gigabytes returns the value in gigabytes
func (b Bytes) Gigabytes() float64 {
	// TODO(human): Implement
	return float64(b) / (1024 * 1024 * 1024)
}
