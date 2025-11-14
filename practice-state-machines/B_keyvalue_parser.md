# Exercise B: Key-Value Parser

**Time Estimate:** 30-40 minutes
**Difficulty:** Medium
**Goal:** Practice state-based parsing with delimiters

## The Problem

Parse URL query string format: `"name=John&age=30&city=NYC"`

Output: `map[string]string{"name": "John", "age": "30", "city": "NYC"}`

## Function Signature

```go
func ParseKeyValue(s string) map[string]string
```

## Examples

```go
ParseKeyValue("name=John&age=30")
// → map[string]string{"name": "John", "age": "30"}

ParseKeyValue("x=1&y=2&z=3")
// → map[string]string{"x": "1", "y": "2", "z": "3"}

ParseKeyValue("")
// → map[string]string{} (empty map)

ParseKeyValue("single=value")
// → map[string]string{"single": "value"}
```

## State Machine Design

```
STATE 1: Reading Key
  - If char == '=' → switch to STATE 2
  - Otherwise → accumulate into key

STATE 2: Reading Value
  - If char == '&' → save key/value pair, reset, go to STATE 1
  - Otherwise → accumulate into value

End of string:
  - Save final key/value pair
```

## Implementation Pattern

```go
func ParseKeyValue(s string) map[string]string {
    result := make(map[string]string)

    if len(s) == 0 {
        return result
    }

    key := ""
    value := ""
    state := "READING_KEY" // or "READING_VALUE"

    for i := 0; i < len(s); i++ {
        char := s[i]

        switch state {
        case "READING_KEY":
            if char == '=' {
                state = "READING_VALUE"
            } else {
                key += string(char)
            }

        case "READING_VALUE":
            if char == '&' {
                result[key] = value
                key = ""
                value = ""
                state = "READING_KEY"
            } else {
                value += string(char)
            }
        }
    }

    // Don't forget the last pair!
    if key != "" {
        result[key] = value
    }

    return result
}
```

## Your Task

1. Create `keyvalue_parser.go`
2. Implement `ParseKeyValue` using the pattern above
3. Test with the examples
4. Make sure you handle the last key/value pair (after loop ends)

## Common Mistakes

1. **Forgetting the last pair:** After loop ends, you still have key/value in buffers!
2. **Not resetting buffers:** After `&`, reset both key and value to ""
3. **Off-by-one errors:** Make sure you don't skip characters when switching states

## Key Learning

- **State-based parsing:** State determines how to interpret each character
- **Delimiter handling:** '=' and '&' trigger state changes
- **Buffer management:** Accumulate characters until delimiter found
- **End-of-input handling:** Don't forget data still in buffers!

## After Completing

Try extending:
- Handle empty values: "key=&other=value"
- Handle URL encoding: "name=John%20Doe" → "name=John Doe"
- Validate format (must have '=' for each pair)

**Time to complete:** 30-40 minutes
