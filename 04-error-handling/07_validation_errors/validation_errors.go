package validation_errors

import (
	"fmt"
	"strings"
)

// TODO(human): Define ValidationErrors struct with Errors []FieldError
type ValidationErrors struct {
	Errors []FieldError
}

// TODO(human): Define FieldError struct with Field and Message
type FieldError struct {
	Field   string
	Message string
}

// TODO(human): Implement Error() method on *ValidationErrors
func (ve *ValidationErrors) Error() string {
	return fmt.Sprintf("%s", ve.Errors)
}

// Add adds a field error to the collection
func (ve *ValidationErrors) Add(field, message string) {
	// TODO(human): Implement
	new := FieldError{
		Field:   field,
		Message: message,
	}
	ve.Errors = append(ve.Errors, new)
}

// HasErrors returns true if there are any errors
func (ve *ValidationErrors) HasErrors() bool {
	// TODO(human): Implement
	if len(ve.Errors) != 0 {
		return true
	}
	return false
}

// ValidateUser validates all user fields and returns all errors
func ValidateUser(name, email string, age int) error {
	// TODO(human): Implement
	// newFieldErr := FieldError{}
	validErrors := &ValidationErrors{
		Errors: []FieldError{},
	}
	if name == "" {
		validErrors.Errors = append(validErrors.Errors, FieldError{Field: "name", Message: "cannot be empty"})
	}
	if !strings.Contains(email, "@") {
		validErrors.Errors = append(validErrors.Errors, FieldError{Field: "email", Message: "invalid email"})
	}
	if age < 0 {
		validErrors.Errors = append(validErrors.Errors, FieldError{Field: "age", Message: "negative age"})
	}
	if age > 150 {
		validErrors.Errors = append(validErrors.Errors, FieldError{Field: "age", Message: "exceeds max"})
	}
	if validErrors.HasErrors() {
		return validErrors
	}
	return nil
}

// ValidateProduct validates product fields
func ValidateProduct(name string, price float64, quantity int) error {
	// TODO(human): Implement
	validErrors := &ValidationErrors{
		Errors: []FieldError{},
	}

	if name == "" {
		validErrors.Errors = append(validErrors.Errors, FieldError{Field: "name", Message: "cannot be empty"})
	}
	if price < 0 {
		validErrors.Errors = append(validErrors.Errors, FieldError{Field: "price", Message: "negative"})
	}
	if quantity < 1 {
		validErrors.Errors = append(validErrors.Errors, FieldError{Field: "quantity", Message: "invalid quantity < 1"})
	}
	if validErrors.HasErrors() {
		return validErrors
	}
	return nil
}
