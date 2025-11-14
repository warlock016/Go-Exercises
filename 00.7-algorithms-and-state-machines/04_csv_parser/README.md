# Exercise 04: CSV Parser

**Difficulty:** Medium
**Time:** 40-50 minutes
**Concepts:** Quote-aware parsing, escape sequences, lookahead, state machines

## Learning Goal

Learn to parse structured data with special characters and escape sequences. This teaches you how to handle "contexts" where the same character means different things (comma inside quotes vs outside quotes).

## The Problem

Implement a CSV (Comma-Separated Values) parser that handles quoted fields and escape sequences.

**Rules:**
1. Fields are separated by commas
2. Fields can be wrapped in double quotes: `"field"`
3. Commas inside quotes are part of the field value, not delimiters
4. Quotes inside quoted fields are escaped by doubling: `""` becomes `"`

**Examples:**
- `Alice,Bob,Charlie` → `["Alice", "Bob", "Charlie"]`
- `"Alice","Bob","Charlie"` → `["Alice", "Bob", "Charlie"]`
- `Alice,"New York, NY",30` → `["Alice", "New York, NY", "30"]`
- `"Say ""hello""",world` → `["Say "hello"", "world"]`
- `,,empty` → `["", "", "empty"]`

## Function Signatures

```go
func ParseCSV(line string) []string
```

## Examples

```go
ParseCSV("Alice,Bob,Charlie")
// → ["Alice", "Bob", "Charlie"]

ParseCSV("Alice,30,Engineer")
// → ["Alice", "30", "Engineer"]

ParseCSV(`"Alice","Bob","Charlie"`)
// → ["Alice", "Bob", "Charlie"]

ParseCSV(`Alice,"New York, NY",30`)
// → ["Alice", "New York, NY", "30"]
// Note: comma inside quotes is part of the field

ParseCSV(`"Say ""hello""",world`)
// → ["Say "hello"", "world"]
// Note: doubled quotes become single quotes

ParseCSV(",,")
// → ["", "", ""]

ParseCSV("")
// → [""]
```

## Instructions

1. Create a result slice and a buffer (strings.Builder)
2. Track state: are we inside quotes or outside quotes?
3. Iterate through characters with an index-based loop (need lookahead for `""`)
4. **Outside quotes:**
   - Comma → flush buffer to result, reset buffer
   - Quote → enter quote mode
   - Other → add to buffer
5. **Inside quotes:**
   - Quote → check next character:
     - If next is also quote → add single quote to buffer, skip both
     - If next is comma or end → exit quote mode, don't add the quote
   - Other → add to buffer
6. After loop: flush final field

## Hints

### Basic Hint

The key challenge is the **quote state**. When you're inside quotes, commas are regular characters. When outside quotes, commas are delimiters.

Use a boolean `inQuotes` to track which state you're in.

### Intermediate Hint

You'll need an index-based loop (not `range`) because you need to look ahead:

```go
for i := 0; i < len(line); i++ {
    char := line[i]

    if inQuotes {
        if char == '"' {
            // Check if next char is also a quote (escaped quote)
            if i+1 < len(line) && line[i+1] == '"' {
                // Doubled quote: add single quote and skip next char
                buffer.WriteRune('"')
                i++ // Skip the second quote
            } else {
                // Single quote: exit quote mode
                inQuotes = false
            }
        } else {
            // Regular character inside quotes
            buffer.WriteByte(char)
        }
    } else {
        // Outside quotes logic...
    }
}
```

### Complete Solution Hint

Here's the full algorithm:

```go
func ParseCSV(line string) []string {
    var result []string
    var buffer strings.Builder
    inQuotes := false

    for i := 0; i < len(line); i++ {
        char := line[i]

        if inQuotes {
            if char == '"' {
                // Look ahead for doubled quote
                if i+1 < len(line) && line[i+1] == '"' {
                    buffer.WriteRune('"')
                    i++ // Skip next quote
                } else {
                    // End of quoted field
                    inQuotes = false
                }
            } else {
                buffer.WriteByte(char)
            }
        } else {
            // Outside quotes
            if char == ',' {
                // Field delimiter
                result = append(result, buffer.String())
                buffer.Reset()
            } else if char == '"' {
                // Start of quoted field
                inQuotes = true
            } else {
                // Regular character
                buffer.WriteByte(char)
            }
        }
    }

    // Flush final field
    result = append(result, buffer.String())

    return result
}
```

**Key insights:**
- Quote mode changes the meaning of commas
- Doubled quotes (`""`) are escape sequences for literal quotes
- Need lookahead to distinguish `""` from `"`
- Always flush the final field after the loop

## Think About

1. Why do we need an index-based loop instead of `range`?
2. What happens if the CSV line has mismatched quotes (e.g., `"Alice,Bob`)?
3. How would you extend this to handle multi-line CSV (newlines inside quotes)?
4. What's the time complexity of this parser?

## What This Teaches

- **Context-sensitive parsing:** Same character means different things in different states
- **Escape sequences:** Handling special character representations
- **Lookahead:** Checking future characters to make decisions
- **Index management:** When `range` isn't enough
- **Real-world parsing:** CSV is everywhere (Excel, databases, APIs)

Ready to implement? Open `csv_parser.go` and look for `TODO(human)` markers!
