package error_creation

import (
	"fmt"
	"strings"
	"unicode"
)

// ValidateAge returns an error if age is negative or unrealistic
func ValidateAge(age int) error {
	// TODO(human): Implement

	if age < 0 {
		return fmt.Errorf("negative")
	} else if age > 150 {
		return fmt.Errorf("must be between")
	}
	return nil
}

// ValidateEmail returns an error if email doesn't contain @
func ValidateEmail(email string) error {
	// TODO(human): Implement
	if !strings.Contains(email, "@") {
		return fmt.Errorf("missing @")
	}
	return nil
}

// WithdrawMoney returns error if amount is negative or exceeds balance
func WithdrawMoney(balance, amount float64) (float64, error) {
	// TODO(human): Implement
	if amount < 0 {
		return balance - amount, fmt.Errorf("positive")
	}
	if amount == 0 {
		return balance, fmt.Errorf("positive")
	}
	if balance-amount < 0 {
		return 0, fmt.Errorf("insufficient")
	} else {
		return balance - amount, nil
	}
}

// FormatUserError creates a descriptive error message for user operations
func FormatUserError(operation, username, reason string) error {
	// TODO(human): Implement
	result := fmt.Sprintf("op: %s, u: %s, reason: %s", operation, username, reason)
	return fmt.Errorf("%s", result)
}

// ValidatePassword returns detailed error describing password requirements
func ValidatePassword(password string) error {
	// TODO(human): Implement

	if len(password) < 8 {
		return fmt.Errorf("at least 8")
	}

	hasUpper := false
	hasDigit := false
	hasChar := false
	for _, v := range password {
		if unicode.IsSpace(v) {
			return fmt.Errorf("password contains white space")
		}
		if unicode.IsDigit(v) && !hasDigit {
			hasDigit = true
		} else if !unicode.IsLower(v) && !hasUpper {
			hasUpper = true
		} else if unicode.IsGraphic(v) && !hasChar {
			hasChar = true
		}
	}
	if !hasDigit {
		return fmt.Errorf("missing digit")
	} else if !hasUpper {
		return fmt.Errorf("missing uppercase character")
	} else if !hasChar {
		return fmt.Errorf("missing special character")
	}
	// fmt.Printf("d: %v, U: %v, c: %v\n", hasDigit, hasUpper, hasChar)
	return nil
}

// ParseConfig returns error with filename and line number if parsing fails
func ParseConfig(filename string, lineNumber int, content string) error {
	// TODO(human): Implement
	errResult := fmt.Sprintf("f: %s, ln: %d, content: %s", filename, lineNumber, content)
	return fmt.Errorf("%s", errResult)
}
