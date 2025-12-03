package expressionevaluator

import (
	"fmt"
	"strconv"
	_ "strconv"
	"strings"
	_ "strings"
	"unicode"
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

	tokens, err := tokenize(expr)

	fmt.Printf("Tokens: %v\n", tokens)

	if err != nil {
		return 0, fmt.Errorf("invalid tokens")
	}

	newParser := Parser{
		tokens: make([]Token, 0),
		// pos:    0,
	}

	newParser.tokens = append(newParser.tokens, tokens...)

	result, err := newParser.parseExpression()

	// if newParser.pos < len(newParser.tokens)-1 {
	// 	return 0, fmt.Errorf("leftover tokens")
	// }

	if newParser.currentToken().Type != "EOF" {
		return 0, fmt.Errorf("leftover tokens")
	}

	return result, err
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

	result := []Token{}
	numState := false
	var currentNum strings.Builder

	for i, v := range expr {
		if unicode.IsDigit(v) {
			numState = true
			currentNum.WriteRune(v)

			// if we reached the end, then convert the digit string to int and append the Token
			if i == len(expr)-1 {
				value, err := strconv.Atoi(currentNum.String())
				if err != nil {
					return nil, fmt.Errorf("string to int conversion error")
				}
				result = append(result, Token{Type: "NUMBER", Value: value})
			}
		} else {
			// if we encounter a non-digit character, then we need to create the digit token and append it, before proceeding with the character parsing
			if numState {
				numState = false
				value, err := strconv.Atoi(currentNum.String())
				if err != nil {
					return nil, fmt.Errorf("string to int conversion error")
				}
				result = append(result, Token{Type: "NUMBER", Value: value})
				currentNum.Reset()
			}

			switch v {
			case '+':
				result = append(result, Token{Type: "PLUS"})
			case '-':
				result = append(result, Token{Type: "MINUS"})
			case '/':
				result = append(result, Token{Type: "DIV"})
			case '*':
				result = append(result, Token{Type: "MUL"})
			case '(':
				result = append(result, Token{Type: "LPAREN"})
			case ')':
				result = append(result, Token{Type: "RPAREN"})
			case ' ':
				continue
			default:
				return nil, fmt.Errorf("unknown token %v", v)
			}
		}
	}

	result = append(result, Token{Type: "EOF"})
	return result, nil
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

	var result int
	var err error = nil

	result, err = p.parseTerm()

	if err != nil {
		return 0, err
	}

	for p.currentToken().Type == "MINUS" || p.currentToken().Type == "PLUS" {
		op := p.currentToken().Type
		p.advance()

		right, err := p.parseTerm()

		if err != nil {
			return 0, err
		}

		switch op {
		case "PLUS":
			result += right
		case "MINUS":
			result -= right
		default:
			return 0, fmt.Errorf("invalid operation")
		}
	}

	return result, err
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

	var result int
	var err error = nil

	result, err = p.parseFactor()

	if err != nil {
		return 0, err
	}

	for p.currentToken().Type == "MUL" || p.currentToken().Type == "DIV" {
		op := p.currentToken().Type
		p.advance()

		right, err := p.parseFactor()

		if err != nil {
			return 0, err
		}

		if op == "MUL" {
			result *= right
		} else if op == "DIV" && right != 0 {
			result /= right
		} else {
			return 0, fmt.Errorf("invalid division by zero")
		}
	}

	return result, err
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

	var result int
	var err error = nil

	if p.currentToken().Type == "NUMBER" {
		result = p.currentToken().Value
		p.advance()
		// return result, nil
	} else if p.currentToken().Type == "LPAREN" {
		p.advance()
		result, err = p.parseExpression()
		if err != nil {
			return 0, fmt.Errorf("error parsing expression")
		}
		err = p.expect("RPAREN")

		if err != nil {
			return 0, fmt.Errorf("invalid expression")
		}
	} else {
		return 0, fmt.Errorf("unexpected token")
	}

	return result, err
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
