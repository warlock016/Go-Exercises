package validation_errors

// TODO(human): Define ValidationErrors struct with Errors []FieldError

// TODO(human): Define FieldError struct with Field and Message

// TODO(human): Implement Error() method on *ValidationErrors

// Add adds a field error to the collection
func (ve *ValidationErrors) Add(field, message string) {
	// TODO(human): Implement
}

// HasErrors returns true if there are any errors
func (ve *ValidationErrors) HasErrors() bool {
	// TODO(human): Implement
	return false
}

// ValidateUser validates all user fields and returns all errors
func ValidateUser(name, email string, age int) error {
	// TODO(human): Implement
	return nil
}

// ValidateProduct validates product fields
func ValidateProduct(name string, price float64, quantity int) error {
	// TODO(human): Implement
	return nil
}
