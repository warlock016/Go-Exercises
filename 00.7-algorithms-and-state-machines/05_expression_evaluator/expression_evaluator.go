package expressionevaluator

import (
	"fmt"
	_ "strconv"
	_ "strings"
	_ "unicode"
)

// Token represents a single token in the expression
type Token struct {
	Type  string // "NUMBER", "PLUS", "MINUS", "MUL", "DIV", "LPAREN", "RPAREN", "EOF"
	Value int    // Only used for NUMBER tokens
}

// Parser holds the state for parsing
type Parser struct {
	tokens []Token
	pos    int
}

// Evaluate parses and evaluates a mathematical expression
func Evaluate(expr string) (int, error) {
	// TODO(human): Implement the main evaluation function
	//
	// Steps:
	// 1. Tokenize the expression using tokenize()
	// 2. Create a Parser with those tokens
	// 3. Call parseExpression() to get the result
	// 4. Check if there are unexpected tokens remaining
	//
	// Hint: If parser.pos < len(parser.tokens)-1 after parsing,
	// there are leftover tokens (error condition)

	return 0, nil
}

// tokenize converts a string expression into tokens
func tokenize(expr string) ([]Token, error) {
	// TODO(human): Implement tokenization
	//
	// Algorithm:
	// 1. Iterate through each character
	// 2. Skip whitespace
	// 3. If digit, collect all consecutive digits into a number
	// 4. If operator or paren, create appropriate token
	// 5. Add EOF token at the end
	//
	// Hint: Use unicode.IsDigit() to check for digits
	// Hint: Use strconv.Atoi() to convert string to int
	//
	// Example: "2 + 3" → [NUMBER(2), PLUS, NUMBER(3), EOF]

	// for _, r := range expr {
	// 	switch unicode.IsDigit(r) {
	// 	case true:
	// 	case false:
	// 	}
	// }

	return nil, nil
}

// parseExpression handles addition and subtraction (lowest precedence)
func (p *Parser) parseExpression() (int, error) {
	// TODO(human): Implement expression parsing
	//
	// Pattern:
	// 1. Get first term by calling parseTerm()
	// 2. While current token is + or -:
	//    a. Remember the operator
	//    b. Advance to next token
	//    c. Get next term by calling parseTerm()
	//    d. Apply operator to result
	// 3. Return final result
	//
	// Hint: This is where you handle left-to-right evaluation
	// for operators of the same precedence

	return 0, nil
}

// parseTerm handles multiplication and division (higher precedence)
func (p *Parser) parseTerm() (int, error) {
	// TODO(human): Implement term parsing
	//
	// Pattern: Similar to parseExpression() but:
	// 1. Call parseFactor() instead of parseTerm()
	// 2. Check for * and / instead of + and -
	// 3. Handle division by zero error
	//
	// Think: Why does calling parseFactor() make multiplication
	// happen before addition?

	return 0, nil
}

// parseFactor handles numbers and parentheses (highest precedence)
func (p *Parser) parseFactor() (int, error) {
	// TODO(human): Implement factor parsing
	//
	// Cases:
	// 1. If current token is NUMBER:
	//    - Get the value
	//    - Advance position
	//    - Return the value
	//
	// 2. If current token is LPAREN:
	//    - Advance past the (
	//    - Recursively call parseExpression()
	//    - Expect RPAREN (error if not)
	//    - Advance past the )
	//    - Return the result
	//
	// 3. Otherwise: error (unexpected token)
	//
	// Think: Why does recursively calling parseExpression()
	// make parentheses work correctly?

	return 0, nil
}

// currentToken returns the current token
func (p *Parser) currentToken() Token {
	if p.pos >= len(p.tokens) {
		return Token{Type: "EOF"}
	}
	return p.tokens[p.pos]
}

// advance moves to the next token
func (p *Parser) advance() {
	p.pos++
}

// expect checks if current token matches expected type and advances
func (p *Parser) expect(tokenType string) error {
	if p.currentToken().Type != tokenType {
		return fmt.Errorf("expected %s, got %s", tokenType, p.currentToken().Type)
	}
	p.advance()
	return nil
}
