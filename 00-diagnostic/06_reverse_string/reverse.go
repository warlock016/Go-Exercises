package reverse

import (
	"unicode/utf8"
)

// Reverse returns the string reversed, handling Unicode properly
func Reverse(s string) string {
	// TODO(human): Implement string reversal
	reverse := ""

	b := []byte(s)

	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		reverse += string(r)
		b = b[:len(b)-size]
	}

	return reverse
}
