package jsonliteparser

import (
	"fmt"
	_ "strconv"
	_ "strings"
	"unicode"
)

// JSONParser holds the parsing state
type JSONParser struct {
	input string
	pos   int
}

// ParseJSON parses a JSON string and returns the resulting value
func ParseJSON(input string) (interface{}, error) {
	// TODO(human): Implement JSON parsing
	//
	// Steps:
	// 1. Create a JSONParser with the input
	// 2. Call parseValue() to parse the JSON
	// 3. Check that all input was consumed (no trailing characters)
	//
	// Hint: After parsing, skip whitespace and check if pos == len(input)

	return nil, nil
}

// parseValue determines the type and parses the appropriate value
func (p *JSONParser) parseValue() (interface{}, error) {
	// TODO(human): Implement value parsing
	//
	// Algorithm:
	// 1. Skip any whitespace
	// 2. Look at the current character (peek)
	// 3. Based on the character, call the appropriate parser:
	//    '{' → parseObject()
	//    '[' → parseArray()
	//    '"' → parseString()
	//    't' → parseTrue()
	//    'f' → parseFalse()
	//    'n' → parseNull()
	//    digit or '-' → parseNumber()
	// 4. Return the parsed value
	//
	// Hint: Use a switch statement on p.peek()

	return nil, nil
}

// parseObject parses a JSON object: { "key": value, ... }
func (p *JSONParser) parseObject() (map[string]interface{}, error) {
	// TODO(human): Implement object parsing
	//
	// Algorithm:
	// 1. Consume the opening '{'
	// 2. Create a map to store key-value pairs
	// 3. Check for empty object ('}' immediately)
	// 4. Loop:
	//    a. Parse key (must be a string)
	//    b. Skip whitespace and expect ':'
	//    c. Parse value (call parseValue - this is the recursion!)
	//    d. Store key-value pair in map
	//    e. Skip whitespace and check next character:
	//       - '}' → done, break
	//       - ',' → continue to next pair
	//       - else → error
	// 5. Return the map
	//
	// Think: Why does calling parseValue() for the value allow
	// nested objects to work?

	return nil, nil
}

// parseArray parses a JSON array: [ value, value, ... ]
func (p *JSONParser) parseArray() ([]interface{}, error) {
	// TODO(human): Implement array parsing
	//
	// Algorithm: Similar to parseObject but:
	// 1. Consume opening '['
	// 2. Create a slice to store values
	// 3. Check for empty array (']' immediately)
	// 4. Loop:
	//    a. Parse value (call parseValue)
	//    b. Append to slice
	//    c. Skip whitespace and check next character:
	//       - ']' → done, break
	//       - ',' → continue to next value
	//       - else → error
	// 5. Return the slice
	//
	// Think: Arrays can contain objects, and objects can contain
	// arrays. How does recursion handle this?

	return nil, nil
}

// parseString parses a JSON string: "text"
func (p *JSONParser) parseString() (string, error) {
	// TODO(human): Implement string parsing
	//
	// Algorithm:
	// 1. Consume opening '"'
	// 2. Build the string character by character:
	//    - If character is '"', we're done
	//    - If character is '\\', handle escape sequence:
	//      - '\\n' → newline
	//      - '\\t' → tab
	//      - '\\"' → quote
	//      - '\\\\' → backslash
	//    - Otherwise, add character to result
	// 3. Consume closing '"'
	// 4. Return the string
	//
	// Hint: Use strings.Builder for efficient string building
	// Hint: Check for end of input (unterminated string error)

	return "", nil
}

// parseNumber parses a JSON number: 123, -45, 3.14
func (p *JSONParser) parseNumber() (float64, error) {
	// TODO(human): Implement number parsing
	//
	// Algorithm:
	// 1. Collect all consecutive characters that form a number:
	//    - Digits (0-9)
	//    - Minus sign (-) at start
	//    - Decimal point (.)
	// 2. Use strconv.ParseFloat() to convert to float64
	// 3. Return the number
	//
	// Hint: Build a string of the number characters, then parse it
	// Hint: Check for at least one digit

	return 0, nil
}

// parseTrue parses the literal "true"
func (p *JSONParser) parseTrue() (bool, error) {
	// TODO(human): Implement true parsing
	//
	// Algorithm:
	// 1. Check that next 4 characters are "true"
	// 2. If yes, consume them and return true
	// 3. If no, return error
	//
	// Hint: Use p.expect("true")

	return false, nil
}

// parseFalse parses the literal "false"
func (p *JSONParser) parseFalse() (bool, error) {
	// TODO(human): Implement false parsing
	//
	// Similar to parseTrue, but for "false"

	return false, nil
}

// parseNull parses the literal "null"
func (p *JSONParser) parseNull() (interface{}, error) {
	// TODO(human): Implement null parsing
	//
	// Algorithm:
	// 1. Check that next 4 characters are "null"
	// 2. If yes, consume them and return nil
	// 3. If no, return error

	return nil, nil
}

// Helper methods

// peek returns the current character without advancing
func (p *JSONParser) peek() byte {
	if p.pos >= len(p.input) {
		return 0 // EOF
	}
	return p.input[p.pos]
}

// consume returns the current character and advances position
func (p *JSONParser) consume() byte {
	if p.pos >= len(p.input) {
		return 0 // EOF
	}
	ch := p.input[p.pos]
	p.pos++
	return ch
}

// skipWhitespace skips spaces, tabs, newlines, carriage returns
func (p *JSONParser) skipWhitespace() {
	for p.pos < len(p.input) && unicode.IsSpace(rune(p.input[p.pos])) {
		p.pos++
	}
}

// expect checks if the next characters match the expected string
func (p *JSONParser) expect(expected string) error {
	if p.pos+len(expected) > len(p.input) {
		return fmt.Errorf("unexpected end of input, expected %s", expected)
	}
	actual := p.input[p.pos : p.pos+len(expected)]
	if actual != expected {
		return fmt.Errorf("expected %s, got %s", expected, actual)
	}
	p.pos += len(expected)
	return nil
}
