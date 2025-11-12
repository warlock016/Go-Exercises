# Exercise 09: Word Wrapper

**Concept:** Text wrapping algorithm with word boundary detection
**Difficulty:** Hard
**Estimated Time:** 35-40 minutes

## Learning Goal

Implement a sophisticated text wrapping algorithm that breaks lines at word boundaries while respecting maximum line width. This exercise combines string building, word splitting, length tracking, and algorithmic thinking. You'll learn how text editors, terminals, and formatters handle line wrapping.

## The Problem

Text wrapping is everywhere:
- Terminal output (wrapping at 80 columns)
- Text editors (word wrap mode)
- Email clients (wrapping at 72 characters)
- Code formatters (keeping lines under max length)
- PDF generators (fitting text in a column)

The challenge: Break lines to fit within a maximum width WITHOUT breaking in the middle of words.

```go
text := "The quick brown fox jumps over the lazy dog"
maxWidth := 20

result:
"The quick brown fox\njumps over the lazy\ndog"
//  ↑ 19 chars        ↑ 19 chars         ↑ 3 chars
```

## Your Task

Implement a word wrapping function that:
1. Wraps text to fit within `maxWidth` characters per line
2. Breaks only at word boundaries (spaces)
3. Preserves single spaces between words
4. Returns lines separated by `\n`
5. Handles edge cases gracefully

```go
func WrapText(text string, maxWidth int) string
```

## Examples

### Basic Examples
```go
WrapText("hello world", 20)
// → "hello world"
// (fits in one line)

WrapText("hello world", 5)
// → "hello\nworld"
// (each word on its own line)

WrapText("The quick brown fox jumps over the lazy dog", 20)
// → "The quick brown fox\njumps over the lazy\ndog"
```

### Edge Cases
```go
WrapText("", 10)
// → ""
// (empty input)

WrapText("hello", 10)
// → "hello"
// (single word)

WrapText("one two three four five", 10)
// → "one two\nthree four\nfive"

WrapText("verylongword", 5)
// → "verylongword"
// (word longer than maxWidth - keep it on one line)

WrapText("word  word", 10)
// → "word word"
// (normalize multiple spaces to single space)

WrapText("a b c d e f g", 5)
// → "a b c\nd e f\ng"
```

## Instructions

1. Open `word_wrapper.go`
2. Implement the `WrapText` function
3. Run `go test -v` to verify your solution
4. Pay attention to edge cases and boundary conditions

## Algorithm Walkthrough

Let's wrap `"The quick brown fox"` with `maxWidth = 10`:

```
words: ["The", "quick", "brown", "fox"]
maxWidth: 10

Line 1:
  - Add "The" (3 chars) → currentLine = "The"
  - Add " " + "quick" (1 + 5 = 6 chars) → "The quick" (9 chars total)
  - Try "brown": "The quick brown" = 15 chars > 10
  - WRAP! Output "The quick", start new line

Line 2:
  - Add "brown" (5 chars) → currentLine = "brown"
  - Add " " + "fox" (1 + 3 = 4 chars) → "brown fox" (9 chars total)
  - No more words

Result: "The quick\nbrown fox"
```

## Hints

### Hint 1: Algorithm Structure

```go
import "strings"

func WrapText(text string, maxWidth int) string {
    // 1. Handle edge cases (empty string, maxWidth <= 0)
    // 2. Split text into words
    // 3. Build lines by adding words until maxWidth is reached
    // 4. Join lines with \n
}
```

### Hint 2: Splitting Words

```go
words := strings.Fields(text)
// Fields splits on whitespace and removes empty strings
// "hello  world" → ["hello", "world"]
```

### Hint 3: Building Lines

```go
var lines []string
var currentLine strings.Builder

for _, word := range words {
    // Check if adding this word would exceed maxWidth
    // If yes: save current line, start new line
    // If no: add word to current line
}

// Don't forget to add the last line!
if currentLine.Len() > 0 {
    lines = append(lines, currentLine.String())
}

return strings.Join(lines, "\n")
```

### Hint 4: Length Calculation

The tricky part is calculating the length BEFORE adding a word:

```go
// Length if we add this word
newLength := currentLine.Len()
if currentLine.Len() > 0 {
    newLength++ // +1 for the space before the word
}
newLength += len(word)

if newLength > maxWidth {
    // Would exceed limit - wrap
    lines = append(lines, currentLine.String())
    currentLine.Reset()
    currentLine.WriteString(word)
} else {
    // Fits - add to current line
    if currentLine.Len() > 0 {
        currentLine.WriteRune(' ')
    }
    currentLine.WriteString(word)
}
```

### Hint 5: Full Solution Structure

```go
func WrapText(text string, maxWidth int) string {
    // Edge cases
    if text == "" {
        return ""
    }
    if maxWidth <= 0 {
        return text
    }

    words := strings.Fields(text)
    if len(words) == 0 {
        return ""
    }

    var lines []string
    var currentLine strings.Builder

    for _, word := range words {
        // Calculate length if we add this word
        futureLen := currentLine.Len()
        if currentLine.Len() > 0 {
            futureLen++ // space
        }
        futureLen += len(word)

        if futureLen > maxWidth && currentLine.Len() > 0 {
            // Wrap: save current line, start new
            lines = append(lines, currentLine.String())
            currentLine.Reset()
            currentLine.WriteString(word)
        } else {
            // Add to current line
            if currentLine.Len() > 0 {
                currentLine.WriteRune(' ')
            }
            currentLine.WriteString(word)
        }
    }

    // Add last line
    if currentLine.Len() > 0 {
        lines = append(lines, currentLine.String())
    }

    return strings.Join(lines, "\n")
}
```

## Think About

1. **Why use strings.Fields instead of strings.Split(" ")?**
   - Fields handles multiple consecutive spaces correctly
   - Fields removes leading/trailing whitespace
   - `"  hello  world  "` → `["hello", "world"]` (no empty strings)

2. **What happens with words longer than maxWidth?**
   - They can't be broken (we don't hyphenate)
   - Options: keep on one line (exceeds max) or reject
   - Our implementation keeps them (practical choice)

3. **Why check `currentLine.Len() > 0` before wrapping?**
   - Prevents wrapping before adding the first word
   - If word is too long, it still goes on a line by itself

4. **How would you modify this for hard wrapping?**
   - Break words at character boundaries if needed
   - Insert hyphens at break points
   - Much more complex!

## What This Teaches

- **Algorithm design** - breaking down a complex problem into steps
- **State management** - tracking current line and future length
- **String building** - efficient line construction with strings.Builder
- **Word splitting** - using strings.Fields correctly
- **Edge case handling** - empty strings, long words, boundary conditions
- **String joining** - combining lines with separators
- **Look-ahead logic** - calculating future state before committing

## Common Mistakes to Avoid

1. **Not handling empty input**
   ```go
   // BUG - crashes on empty string
   words := strings.Fields(text)
   firstWord := words[0]  // Panic if words is empty!

   // FIX
   if len(words) == 0 {
       return ""
   }
   ```

2. **Forgetting the space between words**
   ```go
   // BUG - words run together
   currentLine.WriteString(word)

   // FIX
   if currentLine.Len() > 0 {
       currentLine.WriteRune(' ')
   }
   currentLine.WriteString(word)
   ```

3. **Not adding the last line**
   ```go
   // BUG - last line is lost
   for _, word := range words {
       // ... build lines ...
   }
   return strings.Join(lines, "\n")  // Missing last line!

   // FIX
   if currentLine.Len() > 0 {
       lines = append(lines, currentLine.String())
   }
   ```

4. **Wrong length calculation**
   ```go
   // BUG - doesn't account for space
   if currentLine.Len() + len(word) > maxWidth {

   // FIX - need +1 for space (if line isn't empty)
   futureLen := currentLine.Len()
   if currentLine.Len() > 0 {
       futureLen++
   }
   futureLen += len(word)
   if futureLen > maxWidth {
   ```

5. **Using string concatenation instead of Builder**
   ```go
   // SLOW - creates new string each time
   currentLine := ""
   currentLine += word

   // FAST - efficient building
   var currentLine strings.Builder
   currentLine.WriteString(word)
   ```

6. **Off-by-one in length check**
   ```go
   // BUG - should be > not >=
   if futureLen >= maxWidth {  // "hello" with max 5 shouldn't wrap!

   // FIX
   if futureLen > maxWidth {
   ```

## Challenge Extensions (Optional)

After completing the basic version, try these:

1. **WrapTextIndent**: Add indentation to wrapped lines
   ```go
   WrapText("long text", 20, "  ") // 2-space indent for wrapped lines
   ```

2. **WrapTextHyphenate**: Break long words with hyphens
   ```go
   WrapText("verylongword", 5) → "very-\nlongw-\nord"
   ```

3. **WrapTextJustify**: Justify text (add spaces to reach maxWidth)
   ```go
   // Each line except last should be exactly maxWidth
   ```

4. **WrapPreserveNewlines**: Respect existing \n in input
   ```go
   WrapText("line1\nline2", 20) // Keep both lines separate
   ```

## Real-World Applications

- **Terminal programs**: Formatting output to fit terminal width
- **Email clients**: Wrapping at 72 characters per RFC
- **Code formatters**: Keeping lines under 80/100 characters
- **Document generators**: Fitting text in PDF/HTML columns
- **Chat applications**: Wrapping messages in bubbles
- **Help text**: Formatting CLI help messages

## After Completing

You now understand:
- How to implement a line wrapping algorithm
- How to track state across iterations
- How to calculate future outcomes before committing
- How to handle edge cases in text processing
- Why text wrapping is more complex than it seems

This algorithm is used in almost every text-rendering system. You've just implemented a fundamental piece of text processing!

---

**Next up:** Exercise 10 - Run-Length Encoding (compression algorithm)
