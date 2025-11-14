# Exercise 06: JSON Lite Parser

**Difficulty:** Hard
**Time:** 60-70 minutes
**Concepts:** Recursive parsing, state machines, complex data structures, character-by-character processing

## Learning Goal

Learn to build a JSON parser from scratch that handles nested objects and arrays. This teaches recursive descent parsing for complex, nested data structures and reinforces state management during parsing.

## The Problem

Implement a simple JSON parser that can parse a subset of JSON:

```json
{"name":"John","age":30}
```

Your parser should handle:
- Objects: `{ "key": value, ... }`
- Arrays: `[ value, value, ... ]`
- Strings: `"text"`
- Numbers: `123`, `-45`
- Booleans: `true`, `false`
- Null: `null`
- Nested structures: objects in objects, arrays in arrays, etc.

The parser returns `interface{}` which can be:
- `map[string]interface{}` for objects
- `[]interface{}` for arrays
- `string` for strings
- `float64` for numbers
- `bool` for booleans
- `nil` for null

## Function Signatures

```go
func ParseJSON(input string) (interface{}, error)
```

## Examples

```go
// Simple object
result, err := ParseJSON(`{"name":"John"}`)
// result: map[string]interface{}{"name": "John"}

// Object with multiple types
result, err = ParseJSON(`{"name":"John","age":30,"active":true}`)
// result: map[string]interface{}{"name":"John", "age":30.0, "active":true}

// Nested object
result, err = ParseJSON(`{"user":{"name":"John","age":30}}`)
// result: map[string]interface{}{"user": map[string]interface{}{"name":"John", "age":30.0}}

// Array
result, err = ParseJSON(`[1,2,3]`)
// result: []interface{}{1.0, 2.0, 3.0}

// Array of objects
result, err = ParseJSON(`[{"name":"John"},{"name":"Jane"}]`)
// result: []interface{}{map[string]interface{}{"name":"John"}, map[string]interface{}{"name":"Jane"}}

// Mixed array
result, err = ParseJSON(`[1,"text",true,null]`)
// result: []interface{}{1.0, "text", true, nil}

// Empty structures
result, err = ParseJSON(`{}`)
// result: map[string]interface{}{}

result, err = ParseJSON(`[]`)
// result: []interface{}{}
```

## Instructions

1. Create a parser struct that tracks position in the input string
2. Implement helper methods for peeking/consuming characters and skipping whitespace
3. Implement `parseValue()` that determines what type of value to parse
4. Implement specific parsers for each type:
   - `parseObject()` for objects
   - `parseArray()` for arrays
   - `parseString()` for strings
   - `parseNumber()` for numbers
   - `parseTrue()`, `parseFalse()`, `parseNull()` for literals
5. Use recursion for nested structures

## Hints

### Basic Hint: Parser Structure

```go
type JSONParser struct {
    input string
    pos   int
}

func (p *JSONParser) peek() byte {
    // Return current character without advancing
}

func (p *JSONParser) consume() byte {
    // Return current character and advance position
}

func (p *JSONParser) skipWhitespace() {
    // Skip spaces, tabs, newlines
}
```

### Intermediate Hint: Value Parsing

```go
func (p *JSONParser) parseValue() (interface{}, error) {
    p.skipWhitespace()

    switch p.peek() {
    case '{':
        return p.parseObject()
    case '[':
        return p.parseArray()
    case '"':
        return p.parseString()
    case 't':
        return p.parseTrue()
    case 'f':
        return p.parseFalse()
    case 'n':
        return p.parseNull()
    default:
        // Must be a number (or error)
        return p.parseNumber()
    }
}
```

### Advanced Hint: Object Parsing

```go
func (p *JSONParser) parseObject() (map[string]interface{}, error) {
    result := make(map[string]interface{})

    p.consume() // Skip opening {
    p.skipWhitespace()

    // Handle empty object
    if p.peek() == '}' {
        p.consume()
        return result, nil
    }

    for {
        // Parse key (must be string)
        key, err := p.parseString()
        if err != nil {
            return nil, err
        }

        // Expect colon
        p.skipWhitespace()
        if p.consume() != ':' {
            return nil, error("expected :")
        }

        // Parse value (recursive!)
        value, err := p.parseValue()
        if err != nil {
            return nil, err
        }

        result[key] = value

        p.skipWhitespace()
        next := p.consume()
        if next == '}' {
            break  // End of object
        }
        if next != ',' {
            return nil, error("expected , or }")
        }
        p.skipWhitespace()
    }

    return result, nil
}
```

### String Parsing Hint

For strings, you need to:
1. Skip opening `"`
2. Collect characters until closing `"`
3. Handle escape sequences: `\"`, `\\`, `\n`, `\t`
4. Return error if string is unterminated

### Number Parsing Hint

For numbers:
1. Collect all consecutive digits (and `-` for negative, `.` for decimal)
2. Use `strconv.ParseFloat()` to convert to float64
3. JSON numbers are always treated as float64 in Go

## Think About

1. Why is recursion necessary for parsing nested structures?
2. How would you handle escape sequences in strings (like `\n` or `\"`)?
3. What happens if you have deeply nested structures (100 levels deep)?
4. How would you add line/column number tracking for better error messages?

## What This Teaches

- **Recursive parsing:** Handling arbitrarily nested structures
- **State management:** Tracking position while parsing character-by-character
- **Type handling:** Working with `interface{}` for dynamic types
- **Character processing:** Building up complex structures from individual characters
- **Error handling:** Providing meaningful error messages for malformed input

Ready to implement? Open `json_lite_parser.go` and look for `TODO(human)` markers!
