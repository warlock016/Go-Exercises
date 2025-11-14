# Exercise 02: String Tokenizer

**Difficulty:** Easy-Medium
**Time:** 30-40 minutes
**Concepts:** Character-by-character parsing, state machines, accumulation patterns

## Learning Goal

Learn to build a state machine that processes input character-by-character, accumulating data in one state and outputting results when transitioning to another state.

## The Problem

Implement a tokenizer that splits a string into words, treating any whitespace (spaces, tabs, newlines) as delimiters. Multiple consecutive whitespace characters should be treated as a single delimiter.

For example:
- `"hello world  test"` → `["hello", "world", "test"]` (note two spaces between "world" and "test")
- `"  hello  "` → `["hello"]` (leading and trailing spaces ignored)

Your tokenizer should:
- Process the string character-by-character
- Track two states: IN_WORD (accumulating characters) and SKIPPING_WHITESPACE (between words)
- Handle empty strings, single words, multiple spaces, and leading/trailing whitespace

## Function Signatures

```go
func Tokenize(s string) []string
```

## Examples

```go
Tokenize("hello world")
// → ["hello", "world"]

Tokenize("hello world  test")
// → ["hello", "world", "test"]

Tokenize("  hello  ")
// → ["hello"]

Tokenize("")
// → []

Tokenize("single")
// → ["single"]

Tokenize("   ")
// → []
```

## Instructions

1. Create a slice to store the result words
2. Create a buffer (strings.Builder) to accumulate characters for the current word
3. Iterate through each character (rune) in the string
4. Track whether you're currently IN_WORD or SKIPPING_WHITESPACE
5. When you encounter a non-whitespace character:
   - If SKIPPING_WHITESPACE, transition to IN_WORD
   - Add the character to the buffer
6. When you encounter a whitespace character:
   - If IN_WORD, flush the buffer to the result slice and transition to SKIPPING_WHITESPACE
   - If SKIPPING_WHITESPACE, continue skipping
7. After the loop, check if there's a final word in the buffer to flush

## Hints

### Basic Hint

Think about the two states:
- **IN_WORD**: Accumulating characters into the current word
- **SKIPPING_WHITESPACE**: Between words, waiting for the next word to start

You transition between these states when you see whitespace (word → skipping) or non-whitespace (skipping → word).

### Intermediate Hint

Structure your code like this:

```go
var result []string
var buffer strings.Builder
inWord := false

for _, char := range s {
    if isWhitespace(char) {
        // If we were in a word, we need to finish it
        // ...
    } else {
        // We're in a non-whitespace character, add to buffer
        // ...
    }
}

// Don't forget to handle the last word if buffer isn't empty!
```

You'll need a helper function to check if a character is whitespace:
```go
func isWhitespace(r rune) bool {
    return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}
```

### Complete Solution Hint

Here's the full algorithm:

```go
func Tokenize(s string) []string {
    var result []string
    var buffer strings.Builder
    inWord := false

    for _, char := range s {
        if isWhitespace(char) {
            if inWord {
                // Transition: IN_WORD → SKIPPING_WHITESPACE
                result = append(result, buffer.String())
                buffer.Reset()
                inWord = false
            }
            // If !inWord, we're already skipping whitespace, continue
        } else {
            // Non-whitespace character
            if !inWord {
                // Transition: SKIPPING_WHITESPACE → IN_WORD
                inWord = true
            }
            buffer.WriteRune(char)
        }
    }

    // Handle final word
    if inWord {
        result = append(result, buffer.String())
    }

    return result
}
```

## Think About

1. What happens if you forget to check the buffer after the loop?
2. Why do we need the `inWord` boolean? Could we just check if `buffer.Len() > 0`?
3. How would you modify this to tokenize by a specific delimiter (like commas) instead of whitespace?
4. What's the time complexity of this algorithm?

## What This Teaches

- **State tracking:** Using a boolean flag to represent state
- **Accumulation pattern:** Building up results character-by-character
- **Flushing buffers:** Knowing when to output accumulated data
- **Edge case handling:** Empty strings, trailing data, multiple delimiters
- **Character-by-character parsing:** Foundation for more complex parsers

Ready to implement? Open `string_tokenizer.go` and look for `TODO(human)` markers!
