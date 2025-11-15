package convert

import (
	"fmt"
	"strconv"
)

// IntToFloat converts an integer to a float64
func IntToFloat(n int) float64 {
	// TODO(human): Convert int to float64
	return float64(n)
}

// FloatToInt converts a float64 to an integer (truncating decimals)
func FloatToInt(f float64) int {
	// TODO(human): Convert float64 to int
	return int(f)
}

// StringToInt converts a string to an integer
// Returns an error if the string is not a valid integer
func StringToInt(s string) (int, error) {
	// TODO(human): Convert string to int using strconv.Atoi
	val, err := strconv.Atoi(s)

	if err != nil {
		return 0, fmt.Errorf("invalid Atoi conversion: %v", err)
	}

	return val, nil
}

// IntToString converts an integer to a string
func IntToString(n int) string {
	// TODO(human): Convert int to string using strconv.Itoa
	val := strconv.Itoa(n)

	return val
}
