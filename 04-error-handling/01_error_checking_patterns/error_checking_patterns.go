package error_checking_patterns

import (
	"fmt"
	"os"
	"strconv"
)

// Divide performs division and returns an error if divisor is zero
func Divide(a, b float64) (float64, error) {
	// TODO(human): Implement
	if b == 0 {
		return 0, fmt.Errorf("invalid division by zero")
	}
	return a / b, nil
}

// ReadFileLength reads a file and returns its length in bytes
func ReadFileLength(filename string) (int64, error) {
	// TODO(human): Implement
	if len(filename) == 0 {
		return 0, fmt.Errorf("error: empty string argument")
	}
	result, err := os.ReadFile(filename)
	if err != nil {
		return 0, fmt.Errorf("error: %v", err)
	}
	count := len(result)
	return int64(count), err
}

// ParseAndDouble parses a string to int, doubles it, returns error if parsing fails
func ParseAndDouble(s string) (int, error) {
	// TODO(human): Implement
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return num * 2, err
}

// ChainOperations performs multiple operations that can fail
func ChainOperations(a, b, c, d float64) (float64, error) {
	// TODO(human): Implement
	result, err := Divide(a, b)
	if err != nil {
		return 0, err
	}
	result = result * c
	result = result + d
	return result, nil
}

// ProcessFiles reads multiple files and returns total byte count
func ProcessFiles(filenames []string) (int64, error) {
	// TODO(human): Implement

	var totalCount int64
	var err error
	var length int64

	for _, n := range filenames {
		length, err = ReadFileLength(n)
		if err != nil {
			return totalCount, err
		} else {
			totalCount += length
		}
	}
	return totalCount, err
}

// SafeIndexAccess returns the element at index i, or error if out of bounds
func SafeIndexAccess(slice []int, i int) (int, error) {
	// TODO(human): Implement
	if i < 0 || i >= len(slice) {
		return 0, fmt.Errorf("index out of range")
	}
	return slice[i], nil
}
