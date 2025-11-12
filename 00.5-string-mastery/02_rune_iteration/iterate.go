package iterate

import "fmt"

// CollectBytes returns a slice containing all bytes from the string
func CollectBytes(s string) []byte {
	// TODO(human): Convert string to []byte
	// Hint: Use direct conversion: []byte(s)
	return []byte(s)
}

// CollectRunes returns a slice containing all runes from the string
func CollectRunes(s string) []rune {
	// TODO(human): Convert string to []rune
	// Hint: Use direct conversion: []rune(s)
	// OR use range to iterate and collect
	return []rune(s)
}

// GetRuneAt returns the rune at the given index (0-based)
// Returns an error if index is out of bounds
func GetRuneAt(s string, index int) (rune, error) {
	// TODO(human): Get the Nth rune from the string
	// Hint: Convert to []rune first, then check bounds

	col := []rune(s)

	if index >= len(col) || index < 0 {
		return 0, fmt.Errorf("invalid index: out of bounds")
	}

	return col[index], nil
}
