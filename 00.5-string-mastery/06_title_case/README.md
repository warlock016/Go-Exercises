# Exercise 06: Title Case Converter

**Concept:** Unicode transformations and state tracking
**Difficulty:** Medium
**Estimated Time:** 25-30 minutes

## Learning Goal

Learn to transform strings using unicode case functions while tracking state (word boundaries). This exercise combines rune iteration, unicode package functions, and boolean logic to solve a common text formatting problem.

## The Problem

Converting text to "title case" means capitalizing the first letter of each word. Simple, right? But what defines a "word"? And what about edge cases?

```go
"hello world"           → "Hello World"
"hello  world"          → "Hello  World"  (preserve spacing)
"hello-world"           → "Hello-World"   (hyphen separates words)
"it's awesome!"         → "It's Awesome!" (apostrophes do not separate words)
"  leading spaces"      → "  Leading Spaces"
"ALREADY SHOUTING"      → "Already Shouting" (lowercase non-first letters)
```

## Your Task

Implement a title case converter that:
1. Capitalizes the first letter of each word
2. Lowercases all other letters
3. Treats spaces, punctuation, and hyphens as word boundaries
4. Handles Unicode characters correctly
5. Preserves the original whitespace and punctuation

## Function Signature

```go
func ToTitleCase(s string) string
```

## Examples

```go
ToTitleCase("hello world")              // → "Hello World"
ToTitleCase("the quick brown fox")      // → "The Quick Brown Fox"
ToTitleCase("hello-world")              // → "Hello-World"
ToTitleCase("it's a nice day")          // → "It's A Nice Day"
ToTitleCase("STOP SHOUTING")            // → "Stop Shouting"
ToTitleCase("café résumé")              // → "Café Résumé"
ToTitleCase("hello  world")             // → "Hello  World" (preserve double space)
ToTitleCase("  leading")                // → "  Leading"
ToTitleCase("")                         // → ""
ToTitleCase("123abc")                   // → "123Abc"
```

## Instructions

1. Open `title_case.go`
2. Implement the `ToTitleCase` function
3. Run `go test -v` to verify your solution
4. Try the examples manually to understand the behavior

## Approach

You'll need to:
1. Track whether you're at the start of a new word (state tracking)
2. Iterate through runes (not bytes!)
3. Use `unicode.IsLetter()` to detect letters
4. Use `unicode.ToUpper()` and `unicode.ToLower()` for conversion
5. Build the result with `strings.Builder`

### State Tracking Pattern

The key insight is maintaining a boolean flag:

```go
atWordStart := true  // Start of string is start of a word

for each rune:
    if rune is a letter:
        if atWordStart:
            capitalize it
            atWordStart = false
        else:
            lowercase it
    else:
        // Non-letter (space, punctuation, etc.)
        keep it as-is
        atWordStart = true  // Next letter will start a new word
```

## Hints

<details>
<summary>Hint 1: Function Structure</summary>

```go
import (
    "strings"
    "unicode"
)

func ToTitleCase(s string) string {
    var builder strings.Builder
    atWordStart := true

    // Iterate runes...
    // Apply logic based on atWordStart flag

    return builder.String()
}
```
</details>

<details>
<summary>Hint 2: Rune Iteration</summary>

```go
for _, r := range s {
    // r is a rune (character)
    // Check if it's a letter with unicode.IsLetter(r)
}
```
</details>

<details>
<summary>Hint 3: Letter vs Non-Letter Logic</summary>

```go
if unicode.IsLetter(r) {
    if atWordStart {
        builder.WriteRune(unicode.ToUpper(r))
        atWordStart = false
    } else {
        builder.WriteRune(unicode.ToLower(r))
    }
} else {
    // Non-letter: keep as-is, mark next letter as word start
    builder.WriteRune(r)
    atWordStart = true
}
```
</details>

<details>
<summary>Hint 4: Full Solution Approach</summary>

The complete algorithm:
1. Create a strings.Builder for efficiency
2. Initialize `atWordStart = true` (first character starts a word)
3. For each rune in the string:
   - If it's a letter:
     - If at word start: capitalize and set flag to false
     - Otherwise: lowercase
   - If it's not a letter:
     - Keep it unchanged
     - Set flag to true (next letter will be word start)
4. Return the built string
</details>

## Think About

1. **Why use a boolean flag instead of checking the previous character?**
   - State tracking is cleaner and handles edge cases better

2. **Why use `unicode.IsLetter()` instead of checking for spaces?**
   - Spaces aren't the only word separators (hyphens, punctuation, etc.)

3. **What happens with "it's"?**
   - The apostrophe is not a letter, so it triggers a new word
   - "it's" becomes "It'S" - which might not be desired!
   - Real title case is complex (this is a simplified version)

4. **How does this handle Unicode?**
   - `unicode.ToUpper()` and `unicode.ToLower()` work with all Unicode letters
   - Works with accented characters, Cyrillic, Greek, etc.

## What This Teaches

- **State tracking** - maintaining context across iterations
- **Unicode operations** - using the unicode package for transformations
- **Rune iteration** - processing characters correctly regardless of byte length
- **String building** - efficient string construction with strings.Builder
- **Edge case handling** - dealing with whitespace, punctuation, empty strings
- **Boolean logic** - using flags to control program flow

## Common Mistakes to Avoid

1. **Iterating bytes instead of runes**
   ```go
   // WRONG - breaks on non-ASCII
   for i := 0; i < len(s); i++ {
       b := s[i]  // This is a byte, not a character!
   }

   // RIGHT
   for _, r := range s {
       // r is a rune (character)
   }
   ```

2. **Using string concatenation instead of Builder**
   ```go
   // SLOW - creates new string each time
   result := ""
   result += string(r)

   // FAST - efficient building
   var builder strings.Builder
   builder.WriteRune(r)
   ```

3. **Only checking for spaces as separators**
   ```go
   // INCOMPLETE
   if r == ' ' {
       atWordStart = true
   }

   // COMPLETE
   if !unicode.IsLetter(r) {
       atWordStart = true
   }
   ```

4. **Forgetting to reset the flag**
   ```go
   // BUG - atWordStart never changes
   if unicode.IsLetter(r) {
       if atWordStart {
           builder.WriteRune(unicode.ToUpper(r))
           // Missing: atWordStart = false
       }
   }
   ```

## Challenge Extensions (Optional)

After completing the basic version, try these:

1. **Smart apostrophes**: Don't capitalize after apostrophes (it's → It's, not It'S)
2. **Articles exception**: Don't capitalize "a", "an", "the" unless they're first
3. **Preserve all-caps words**: Keep "NASA" as "NASA" not "Nasa"

## After Completing

You now understand:
- How to track state across iterations
- How to use the unicode package for transformations
- When to capitalize vs lowercase based on context
- How to handle edge cases in string processing

This pattern (state tracking + rune iteration) is used in many text processing tasks like parsers, validators, and formatters.

---

**Next up:** Exercise 07 - Character Counter (maps and frequency counting)
