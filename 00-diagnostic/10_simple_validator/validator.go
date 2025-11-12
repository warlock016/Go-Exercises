package validator

import "strings"

// TODO(human): Define the Validator interface
type Validator interface {
	Validate(s string) bool
}

// TODO(human): Define the EmailValidator struct

type EmailValidator struct {
}

// TODO(human): Implement the Validate method for EmailValidator

func (ev EmailValidator) Validate(s string) bool {
	return strings.Contains(s, "@")
}
