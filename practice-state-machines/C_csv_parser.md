# Exercise C: Quote-Aware CSV Parser

**Time Estimate:** 30-40 minutes
**Difficulty:** Medium-Hard
**Goal:** Practice nested states (inside/outside quotes)

## The Problem

Parse CSV (Comma-Separated Values) with quoted fields:

```
Alice,30,Engineer
Bob,"New York, NY",Designer
Carol,"Said ""Hi""",Manager
```

Fields containing commas MUST be quoted.
Quotes inside quoted fields are doubled: `""` → `"`

## Function Signature

```go
func ParseCSV(line string) []string
```

## Examples

```go
ParseCSV("Alice,30,Engineer")
// → []string{"Alice", "30", "Engineer"}

ParseCSV("Bob,\"New York, NY\",Designer")
// → []string{"Bob", "New York, NY", "Designer"}

ParseCSV("Carol,\"Said \"\"Hi\"\"\",Manager")
// → []string{"Carol", "Said \"Hi\"", "Manager"}

ParseCSV("")
// → []string{} (empty slice)

ParseCSV("single")
// → []string{"single"}
```

## State Machine Design

```
STATE 1: OUTSIDE_QUOTE (normal parsing)
  - If char == ',' → save field, reset buffer
  - If char == '"' → switch to INSIDE_QUOTE
  - Otherwise → accumulate into buffer

STATE 2: INSIDE_QUOTE (quoted field)
  - If char == '"' → check next char:
      - If next is '"' → add single quote to buffer, skip next char
      - If next is ',' or end → save field, switch to OUTSIDE_QUOTE
  - Otherwise → accumulate into buffer
```

## Implementation Hints

```go
func ParseCSV(line string) []string {
    var fields []string

    if len(line) == 0 {
        return fields
    }

    field := ""
    inQuote := false

    for i := 0; i < len(line); i++ {
        char := line[i]

        if inQuote {
            // STATE: INSIDE_QUOTE
            if char == '"' {
                // Check if doubled quote or end of quoted field
                if i+1 < len(line) && line[i+1] == '"' {
                    field += "\""
                    i++ // Skip next quote
                } else {
                    inQuote = false
                }
            } else {
                field += string(char)
            }
        } else {
            // STATE: OUTSIDE_QUOTE
            if char == ',' {
                fields = append(fields, field)
                field = ""
            } else if char == '"' {
                inQuote = true
            } else {
                field += string(char)
            }
        }
    }

    // Don't forget the last field!
    fields = append(fields, field)

    return fields
}
```

## Your Task

1. Create `csv_parser.go`
2. Implement `ParseCSV` using the pattern above
3. Test with all the examples
4. Pay special attention to the doubled-quote handling

## Common Pitfalls

1. **Forgetting last field:** After loop, append final field!
2. **Not handling doubled quotes:** `""` inside quoted field should become `"`
3. **Index out of bounds:** When checking `line[i+1]`, always verify `i+1 < len(line)`
4. **State confusion:** Make sure you know which state you're in

## Tricky Test Cases

```go
// Edge case: Quoted field at end
ParseCSV("a,b,\"quoted\"")
// → []string{"a", "b", "quoted"}

// Edge case: Empty quoted field
ParseCSV("a,\"\",c")
// → []string{"a", "", "c"}

// Edge case: Quote at start
ParseCSV("\"quoted\",b,c")
// → []string{"quoted", "b", "c"}
```

## Key Learning

- **Nested states:** `inQuote` boolean creates two parsing contexts
- **Lookahead:** Checking `line[i+1]` to distinguish `""` from `"` end-quote
- **Index manipulation:** Manually incrementing `i` to skip characters
- **Bounds checking:** Always verify array access is in range

## After Completing

Try extending:
- Handle escaped characters: `\n`, `\t` in quoted fields
- Handle multiple delimiters (tab, semicolon)
- Parse multiple lines into `[][]string`
- Add validation (unclosed quotes should error)

## Connection to Run-Length Encoding

Notice the similarities:
- **State-based parsing:** Different behavior based on state
- **Escape sequences:** Doubled quotes, just like doubled digits
- **Bounds checking:** Always check before accessing `line[i+1]`
- **Manual index control:** `for i := 0; i < len(line); i++` pattern

This CSV parser uses the SAME patterns you saw in the Decode function!

**Time to complete:** 30-40 minutes
