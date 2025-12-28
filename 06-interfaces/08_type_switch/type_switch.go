package type_switch

import (
	"strconv"
)

// Stringify converts any value to a string representation
func Stringify(v any) string {
	// TODO(human): Implement using type switch
	var res string
	switch v := v.(type) {
	case string:
		res = v
	case int:
		res = strconv.Itoa(v)
	case float32:
		res = strconv.FormatFloat(float64(v), 'f', 2, 32)
	case float64:
		res = strconv.FormatFloat(v, 'f', 2, 64)
	case bool:
		if v {
			res = "true"
		} else {
			res = "false"
		}
	default:
		res = "unknown type"
	}
	return res
}

// TypeName returns the type name of the value
func TypeName(v any) string {
	// TODO(human): Implement using type switch
	var res string

	switch v.(type) {
	case string:
		res = "string"
	case bool:
		res = "bool"
	case int:
		res = "int"
	case float32:
		res = "float32"
	case float64:
		res = "float64"
	case rune:
		res = "rune"
	default:
		res = "other"
	}
	return res
}

// Add adds two values if they are the same type
func Add(a, b any) (any, bool) {
	// TODO(human): Implement using type switch
	switch a := a.(type) {
	case string:
		if _, ok := b.(string); ok {
			return a + b.(string), true
		}
	case int:
		if _, ok := b.(int); ok {
			return a + b.(int), true
		}
	case float32:
		if _, ok := b.(float32); ok {
			return a + b.(float32), true
		}
	case float64:
		if _, ok := b.(float64); ok {
			return a + b.(float64), true
		}
	default:
		return nil, false
	}
	return nil, false
}
