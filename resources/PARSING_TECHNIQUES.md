# Parsing Techniques: A Practical Guide

## Table of Contents

1. [What is Parsing?](#what-is-parsing)
2. [The Parsing Pipeline](#the-parsing-pipeline)
3. [Common Parsing Patterns](#common-parsing-patterns)
4. [Lexical Analysis (Tokenization)](#lexical-analysis-tokenization)
5. [Syntax Analysis (Parsing Proper)](#syntax-analysis-parsing-proper)
6. [Error Handling Strategies](#error-handling-strategies)
7. [Performance Considerations](#performance-considerations)
8. [Real-World Examples](#real-world-examples)

---

## What is Parsing?

**Parsing** transforms unstructured text into structured data.

```
Input (text):      "name=John&age=30"
Output (structure): map[string]string{"name": "John", "age": "30"}
```

### Why Parsing is Hard

1. **Ambiguity:** "211" could mean "2 ones and 1 one" or "211 of something"
2. **Context-sensitivity:** Behavior depends on what came before
3. **Error recovery:** What to do with invalid input?
4. **Performance:** Parse millions of lines efficiently

---

## The Parsing Pipeline

Most parsers follow this flow:

```
Raw Input
   ↓
[Lexer/Tokenizer] ────> Tokens
   ↓
[Parser] ──────────────> Abstract Syntax Tree (AST) or Data Structure
   ↓
[Semantic Analysis] ───> Validated/Processed Result
```

### Example: Parsing "3a2b"

```
Input: "3a2b"

LEXER:
  → [NUMBER("3"), CHAR("a"), NUMBER("2"), CHAR("b")]

PARSER:
  → [Run(count=3, char='a'), Run(count=2, char='b')]

SEMANTIC:
  → "aaabb"
```

---

## Common Parsing Patterns

### Pattern 1: Greedy Consumption

**Read as much as possible** before making a decision.

```go
func parseNumber(input string, i int) (string, int) {
    num := ""
    for i < len(input) && isDigit(input[i]) {
        num += string(input[i])
        i++
    }
    return num, i
}
```

**Use when:**
- Multi-digit numbers
- Variable-length tokens
- Simple tokenization

**Example:** "123abc" → read "123", then stop at 'a'

---

### Pattern 2: Lookahead

**Peek at next character(s)** before deciding.

```go
func parseString(input string, i int) (string, int) {
    str := ""
    i++ // Skip opening quote

    for i < len(input) {
        if input[i] == '"' {
            // Lookahead: is it escaped?
            if i+1 < len(input) && input[i+1] == '"' {
                str += "\""
                i += 2 // Skip both quotes
            } else {
                i++ // Skip closing quote
                break
            }
        } else {
            str += string(input[i])
            i++
        }
    }

    return str, i
}
```

**Use when:**
- Escape sequences
- Ambiguous delimiters
- Context-dependent parsing

**Example:** CSV with `""` escaping

---

### Pattern 3: Recursive Descent

**Parse nested structures** by calling parsing functions recursively.

```go
func parseExpression(tokens []Token, i int) (Expr, int) {
    left, i := parseTerm(tokens, i)

    for i < len(tokens) && (tokens[i].Type == PLUS || tokens[i].Type == MINUS) {
        op := tokens[i]
        i++
        right, newI := parseTerm(tokens, i)
        i = newI
        left = BinaryExpr{Left: left, Op: op, Right: right}
    }

    return left, i
}

func parseTerm(tokens []Token, i int) (Expr, int) {
    left, i := parseFactor(tokens, i)

    for i < len(tokens) && (tokens[i].Type == MULT || tokens[i].Type == DIV) {
        op := tokens[i]
        i++
        right, newI := parseFactor(tokens, i)
        i = newI
        left = BinaryExpr{Left: left, Op: op, Right: right}
    }

    return left, i
}

func parseFactor(tokens []Token, i int) (Expr, int) {
    if tokens[i].Type == LPAREN {
        i++ // Skip '('
        expr, newI := parseExpression(tokens, i)
        i = newI
        i++ // Skip ')'
        return expr, i
    }

    // Base case: number
    return NumberExpr{Value: tokens[i].Value}, i+1
}
```

**Use when:**
- Grammar has rules (BNF notation)
- Nested structures (parentheses, JSON, XML)
- Operator precedence

**Example:** Parse "2 + 3 * 4" → `2 + (3 * 4)`

---

### Pattern 4: Delimiter-Based Splitting

**Split first, parse second.**

```go
func parseCSV(line string) []string {
    fields := []string{}
    field := ""
    inQuote := false

    for _, char := range line {
        if char == '"' {
            inQuote = !inQuote
        } else if char == ',' && !inQuote {
            fields = append(fields, field)
            field = ""
        } else {
            field += string(char)
        }
    }

    fields = append(fields, field) // Last field
    return fields
}
```

**Use when:**
- Clear delimiters (comma, semicolon, pipe)
- Flat structures
- No nesting

**Example:** CSV, TSV, pipe-delimited

---

### Pattern 5: Regular Expression

**Use regex for simple patterns.**

```go
import "regexp"

func parseEmail(input string) (string, string, error) {
    re := regexp.MustCompile(`^([a-zA-Z0-9._%+-]+)@([a-zA-Z0-9.-]+\.[a-zA-Z]{2,})$`)
    matches := re.FindStringSubmatch(input)

    if matches == nil {
        return "", "", errors.New("invalid email")
    }

    return matches[1], matches[2], nil // username, domain
}
```

**Use when:**
- Well-defined patterns
- No complex nesting
- Performance not critical

**Don't use when:**
- Nested structures (regex can't parse HTML!)
- Complex state machines
- Need detailed error messages

---

## Lexical Analysis (Tokenization)

### What is a Token?

A **token** is a meaningful unit:

```
Input: "x = 42 + y"

Tokens:
  [IDENTIFIER("x"), EQUALS, NUMBER(42), PLUS, IDENTIFIER("y")]
```

### Token Types

```go
type TokenType int

const (
    NUMBER TokenType = iota
    STRING
    IDENTIFIER
    OPERATOR
    DELIMITER
    KEYWORD
    EOF
)

type Token struct {
    Type  TokenType
    Value string
    Line  int  // For error messages
    Col   int
}
```

### Lexer Pattern: Character-at-a-Time

```go
type Lexer struct {
    input string
    pos   int
}

func (l *Lexer) NextToken() Token {
    // Skip whitespace
    for l.pos < len(l.input) && isWhitespace(l.input[l.pos]) {
        l.pos++
    }

    if l.pos >= len(l.input) {
        return Token{Type: EOF}
    }

    char := l.input[l.pos]

    // Number
    if isDigit(char) {
        return l.readNumber()
    }

    // String
    if char == '"' {
        return l.readString()
    }

    // Identifier/Keyword
    if isLetter(char) {
        return l.readIdentifier()
    }

    // Operator
    if isOperator(char) {
        l.pos++
        return Token{Type: OPERATOR, Value: string(char)}
    }

    // Unknown
    l.pos++
    return Token{Type: UNKNOWN, Value: string(char)}
}

func (l *Lexer) readNumber() Token {
    start := l.pos
    for l.pos < len(l.input) && isDigit(l.input[l.pos]) {
        l.pos++
    }
    return Token{Type: NUMBER, Value: l.input[start:l.pos]}
}
```

### Benefits of Separate Lexer

1. **Separation of concerns:** Lexer handles characters, parser handles structure
2. **Reusability:** Same lexer for multiple parsers
3. **Error messages:** Track line/column in lexer
4. **Performance:** Lexer can be optimized independently

---

## Syntax Analysis (Parsing Proper)

### Grammar Notation (BNF)

**Backus-Naur Form** describes syntax rules:

```
<expression> ::= <term> | <expression> "+" <term>
<term>       ::= <factor> | <term> "*" <factor>
<factor>     ::= <number> | "(" <expression> ")"
```

**Translates to code:**

```go
func parseExpression() Expr {
    left := parseTerm()
    for currentToken == PLUS {
        consume(PLUS)
        right := parseTerm()
        left = Add(left, right)
    }
    return left
}
```

### Top-Down vs Bottom-Up

**Top-Down (Recursive Descent):**
- Start with top rule
- Break down into smaller parts
- Easy to write by hand
- Example: Most hand-written parsers

**Bottom-Up (Shift-Reduce):**
- Start with tokens
- Build up to larger structures
- Hard to write by hand (use tools)
- Example: Yacc, Bison

---

## Error Handling Strategies

### Strategy 1: Fail Fast

**Stop on first error.**

```go
func parse(input string) (Result, error) {
    if !validate(input) {
        return nil, errors.New("invalid format")
    }
    // Continue parsing
}
```

**Pros:** Simple, clear
**Cons:** User only sees one error at a time

---

### Strategy 2: Error Recovery

**Try to continue parsing after error.**

```go
func parse(tokens []Token) (Result, []error) {
    var errors []error

    for i := 0; i < len(tokens); {
        token := tokens[i]

        if !isValid(token) {
            errors = append(errors, fmt.Errorf("invalid token at %d", i))
            i = skipToNextValid(tokens, i) // Skip to next valid token
            continue
        }

        // Parse normally
        i++
    }

    return result, errors
}
```

**Pros:** Show multiple errors
**Cons:** More complex, may give misleading errors

---

### Strategy 3: Panic Mode

**Skip tokens until synchronization point.**

```go
func recoverFromError(tokens []Token, i int) int {
    // Skip until we find a semicolon (statement end)
    for i < len(tokens) && tokens[i].Type != SEMICOLON {
        i++
    }
    return i
}
```

**Use in:** Compilers, where you want to report many errors in one pass

---

### Strategy 4: Best-Effort Parsing

**Return partial results.**

```go
func parseCSV(line string) ([]string, error) {
    fields := []string{}
    field := ""

    for i, char := range line {
        if char == ',' {
            fields = append(fields, field)
            field = ""
        } else if char == '\n' {
            return fields, fmt.Errorf("unexpected newline at pos %d", i)
        } else {
            field += string(char)
        }
    }

    fields = append(fields, field)
    return fields, nil // Return what we got
}
```

**Pros:** Useful output even with errors
**Cons:** May hide bugs

---

## Performance Considerations

### 1. Avoid String Concatenation

```go
// ❌ SLOW - creates new string each time
result := ""
for _, char := range input {
    result += string(char) // O(n²)
}

// ✓ FAST - strings.Builder
var result strings.Builder
for _, char := range input {
    result.WriteRune(char) // O(n)
}
```

### 2. Minimize Allocations

```go
// ❌ SLOW - converts string to rune slice
runes := []rune(input)

// ✓ FAST - iterate directly (for simple ASCII)
for i := 0; i < len(input); i++ {
    char := input[i]
    // ...
}

// ✓ FAST - range for Unicode
for _, r := range input {
    // r is already a rune
}
```

### 3. Preallocate Slices

```go
// ❌ SLOW - slice grows dynamically
var tokens []Token
for ... {
    tokens = append(tokens, token) // May reallocate
}

// ✓ FAST - preallocate if size known
tokens := make([]Token, 0, estimatedSize)
for ... {
    tokens = append(tokens, token) // Less reallocation
}
```

### 4. Use Rune Slices for Unicode

```go
// ❌ WRONG - byte-indexed access breaks Unicode
input := "café"
char := input[3] // 'Ã' (wrong!)

// ✓ RIGHT - convert to runes first
runes := []rune(input)
char := runes[3] // 'é' ✓
```

### 5. Avoid Regex for Simple Cases

```go
// ❌ SLOW - regex overhead
re := regexp.MustCompile(`^\d+$`)
if re.MatchString(input) { ... }

// ✓ FAST - simple loop
isNumber := true
for _, r := range input {
    if !unicode.IsDigit(r) {
        isNumber = false
        break
    }
}
```

---

## Real-World Examples

### Example 1: JSON Parser (Simplified)

```go
type JSONValue interface{}

func parseJSON(input string) (JSONValue, error) {
    input = strings.TrimSpace(input)

    if len(input) == 0 {
        return nil, errors.New("empty input")
    }

    switch input[0] {
    case '{':
        return parseObject(input)
    case '[':
        return parseArray(input)
    case '"':
        return parseString(input)
    case 't', 'f':
        return parseBool(input)
    case 'n':
        return parseNull(input)
    default:
        return parseNumber(input)
    }
}

func parseObject(input string) (map[string]JSONValue, error) {
    result := make(map[string]JSONValue)
    input = input[1 : len(input)-1] // Remove { }

    pairs := splitByComma(input) // Simplified

    for _, pair := range pairs {
        parts := strings.SplitN(pair, ":", 2)
        if len(parts) != 2 {
            return nil, errors.New("invalid object")
        }

        key, _ := parseString(parts[0])
        value, _ := parseJSON(parts[1])

        result[key.(string)] = value
    }

    return result, nil
}
```

### Example 2: Calculator (Infix Notation)

```go
// Parse: "2 + 3 * 4" → 14

func evaluate(expr string) (int, error) {
    tokens := tokenize(expr) // ["2", "+", "3", "*", "4"]
    return parseExpression(tokens, 0)
}

func parseExpression(tokens []string, i int) (int, int, error) {
    left, i, err := parseTerm(tokens, i)
    if err != nil {
        return 0, i, err
    }

    for i < len(tokens) && (tokens[i] == "+" || tokens[i] == "-") {
        op := tokens[i]
        i++

        right, newI, err := parseTerm(tokens, i)
        if err != nil {
            return 0, i, err
        }
        i = newI

        if op == "+" {
            left += right
        } else {
            left -= right
        }
    }

    return left, i, nil
}

func parseTerm(tokens []string, i int) (int, int, error) {
    left, i, err := parseFactor(tokens, i)
    if err != nil {
        return 0, i, err
    }

    for i < len(tokens) && (tokens[i] == "*" || tokens[i] == "/") {
        op := tokens[i]
        i++

        right, newI, err := parseFactor(tokens, i)
        if err != nil {
            return 0, i, err
        }
        i = newI

        if op == "*" {
            left *= right
        } else {
            if right == 0 {
                return 0, i, errors.New("division by zero")
            }
            left /= right
        }
    }

    return left, i, nil
}

func parseFactor(tokens []string, i int) (int, int, error) {
    if tokens[i] == "(" {
        i++ // Skip '('
        result, newI, err := parseExpression(tokens, i)
        if err != nil {
            return 0, i, err
        }
        i = newI
        i++ // Skip ')'
        return result, i, nil
    }

    // Number
    num, err := strconv.Atoi(tokens[i])
    if err != nil {
        return 0, i, err
    }
    return num, i + 1, nil
}
```

### Example 3: URL Parser

```go
type URL struct {
    Scheme string
    Host   string
    Port   string
    Path   string
    Query  map[string]string
}

func parseURL(input string) (*URL, error) {
    url := &URL{}

    // STATE 1: Read scheme (http, https, ftp)
    i := strings.Index(input, "://")
    if i == -1 {
        return nil, errors.New("missing scheme")
    }
    url.Scheme = input[:i]
    input = input[i+3:]

    // STATE 2: Read host (and optional port)
    i = strings.IndexAny(input, ":/")
    if i == -1 {
        url.Host = input
        return url, nil
    }

    url.Host = input[:i]
    input = input[i:]

    // STATE 3: Check for port
    if input[0] == ':' {
        input = input[1:]
        i = strings.IndexAny(input, "/")
        if i == -1 {
            url.Port = input
            return url, nil
        }
        url.Port = input[:i]
        input = input[i:]
    }

    // STATE 4: Read path
    i = strings.Index(input, "?")
    if i == -1 {
        url.Path = input
        return url, nil
    }

    url.Path = input[:i]
    input = input[i+1:]

    // STATE 5: Parse query params
    url.Query = parseQueryParams(input)

    return url, nil
}
```

---

## Key Takeaways

1. **Break parsing into stages:** Lexer → Parser → Semantic analysis
2. **Choose the right technique:** Greedy, lookahead, recursive descent, regex
3. **Handle edge cases:** Empty input, malformed data, EOF
4. **Optimize carefully:** strings.Builder, preallocate, avoid regex for simple cases
5. **Error handling matters:** Decide fail-fast vs error recovery early
6. **Test incrementally:** Build up from simple to complex cases

---

## Common Parsing Checklist

- [ ] Empty input handled
- [ ] Single element handled
- [ ] Multi-element handled
- [ ] Delimiters at start/end handled
- [ ] Consecutive delimiters handled
- [ ] Escape sequences handled
- [ ] Unicode handled (use runes!)
- [ ] Bounds checking (no index out of range)
- [ ] End-of-input handled (flush buffers)
- [ ] Error messages helpful
- [ ] Performance acceptable

---

## Further Reading

- **"Crafting Interpreters"** by Robert Nystrom - Best intro to parsing
- **"Engineering a Compiler"** by Cooper & Torczon - Advanced techniques
- **Go standard library:** `go/parser` package for Go code parsing
- **Parsing Expression Grammars (PEG)** - Modern alternative to BNF

---

## Next Steps

1. Read `STATE_MACHINE_GUIDE.md` for state-based parsing
2. Read `DEBUGGING_ALGORITHMS.md` for troubleshooting
3. Practice with the 3 quick exercises (key-value, CSV parsers)
4. Complete the Algorithms & State Machines module

**Remember:** Parsing is everywhere - JSON, XML, CSV, config files, programming languages, protocols. Master these patterns and you can parse anything!

---

**Version:** 1.0
**Last Updated:** 2025-11-14
**Related:** STATE_MACHINE_GUIDE.md, DEBUGGING_ALGORITHMS.md, ENCODING_DECODING_PATTERNS.md
