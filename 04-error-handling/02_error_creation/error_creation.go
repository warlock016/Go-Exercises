package error_creation

// ValidateAge returns an error if age is negative or unrealistic
func ValidateAge(age int) error {
	// TODO(human): Implement
	return nil
}

// ValidateEmail returns an error if email doesn't contain @
func ValidateEmail(email string) error {
	// TODO(human): Implement
	return nil
}

// WithdrawMoney returns error if amount is negative or exceeds balance
func WithdrawMoney(balance, amount float64) (float64, error) {
	// TODO(human): Implement
	return 0, nil
}

// FormatUserError creates a descriptive error message for user operations
func FormatUserError(operation, username, reason string) error {
	// TODO(human): Implement
	return nil
}

// ValidatePassword returns detailed error describing password requirements
func ValidatePassword(password string) error {
	// TODO(human): Implement
	return nil
}

// ParseConfig returns error with filename and line number if parsing fails
func ParseConfig(filename string, lineNumber int, content string) error {
	// TODO(human): Implement
	return nil
}
