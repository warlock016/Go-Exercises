# Exercise 05: Expression Evaluator

**Difficulty:** Medium-Hard
**Time:** 50-60 minutes
**Concepts:** Recursive descent parsing, operator precedence, tokenization, state tracking

## Learning Goal

Learn to build a mathematical expression parser that respects operator precedence using recursive descent parsing. This teaches how to break down complex parsing problems into smaller, manageable parsing functions.

## The Problem

Implement an expression evaluator that can parse and evaluate mathematical expressions with correct operator precedence:

```
"2 + 3 * 4"  → 14  (not 20, because * has higher precedence)
"(2 + 3) * 4" → 20  (parentheses override precedence)
"10 - 2 - 3"  → 5   (left-to-right for same precedence)
```

Your evaluator should:
- Handle basic operators: `+`, `-`, `*`, `/`
- Respect operator precedence (multiplication and division before addition and subtraction)
- Support parentheses to override precedence
- Parse expressions left-to-right for operators of the same precedence

## Function Signatures

```go
func Evaluate(expr string) (int, error)
```

## Examples

```go
result, err := Evaluate("2 + 3")
// result: 5, err: nil

result, err = Evaluate("2 + 3 * 4")
// result: 14, err: nil (3*4 = 12, then 2+12 = 14)

result, err = Evaluate("(2 + 3) * 4")
// result: 20, err: nil (parentheses first)

result, err = Evaluate("10 / 2 + 3")
// result: 8, err: nil (10/2 = 5, then 5+3 = 8)

result, err = Evaluate("10 - 2 - 3")
// result: 5, err: nil (left-to-right: 10-2 = 8, then 8-3 = 5)

result, err = Evaluate("(1 + 2) * (3 + 4)")
// result: 21, err: nil

result, err = Evaluate("2 +")
// result: 0, err: error ("unexpected end of expression")
```

## Instructions

1. First, implement a simple tokenizer that breaks the expression into tokens (numbers, operators, parentheses)
2. Implement a recursive descent parser with three levels:
   - `parseExpression()` - handles + and - (lowest precedence)
   - `parseTerm()` - handles * and / (higher precedence)
   - `parseFactor()` - handles numbers and parentheses (highest precedence)
3. Each parsing function should call the next higher precedence function
4. Handle whitespace by skipping it during tokenization

## Hints

### Basic Hint: Operator Precedence Hierarchy

Think of precedence levels:
- Level 1 (lowest): `+` and `-`
- Level 2 (higher): `*` and `/`
- Level 3 (highest): numbers and `()`

### Intermediate Hint: Recursive Descent Structure

```go
// Expression: handles + and -
func parseExpression() int {
    result := parseTerm()  // Get first term
    for current token is + or - {
        operator := current token
        advance to next token
        right := parseTerm()
        if operator is + {
            result += right
        } else {
            result -= right
        }
    }
    return result
}

// Term: handles * and /
func parseTerm() int {
    result := parseFactor()  // Get first factor
    for current token is * or / {
        // Similar pattern...
    }
    return result
}

// Factor: handles numbers and parentheses
func parseFactor() int {
    if current token is number {
        return the number
    }
    if current token is ( {
        skip the (
        result := parseExpression()  // Recursively parse inside parens
        skip the )
        return result
    }
}
```

### Advanced Hint: Complete Approach

1. **Tokenize:** Convert string to list of tokens
   ```go
   type Token struct {
       Type  string // "NUMBER", "PLUS", "MINUS", "MUL", "DIV", "LPAREN", "RPAREN"
       Value int    // For NUMBER tokens
   }
   ```

2. **Parser State:** Track current position in token list
   ```go
   type Parser struct {
       tokens []Token
       pos    int
   }
   ```

3. **Error Handling:** Return errors for:
   - Unexpected end of expression
   - Invalid characters
   - Mismatched parentheses
   - Division by zero

## Think About

1. Why does calling `parseTerm()` from `parseExpression()` automatically handle precedence correctly?
2. How would you extend this to support more operators (e.g., `^` for exponentiation)?
3. What happens if you have nested parentheses like `((1 + 2) * 3)`?
4. How would you modify this to support floating-point numbers?

## What This Teaches

- **Recursive descent parsing:** Breaking down parsing into hierarchical functions
- **Operator precedence:** How to encode precedence rules in code structure
- **Tokenization:** Converting raw strings into structured tokens
- **State management:** Tracking position in a token stream
- **Error handling:** Detecting and reporting syntax errors

Ready to implement? Open `expression_evaluator.go` and look for `TODO(human)` markers!
