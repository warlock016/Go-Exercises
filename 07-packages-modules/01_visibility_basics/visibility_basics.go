package visibility

import (
	"errors"
)

// User represents a user in the system
type User struct {
	// TODO(human): Add Name field (exported) and password field (unexported)
	Name     string
	password string
}

// config holds package-level configuration
type config struct {
	// TODO(human): Add minPasswordLength and maxNameLength fields
	minPasswordLength int
	maxNameLength     int
}

var cfg = config{
	minPasswordLength: 8,
	maxNameLength:     50,
}

// NewUser creates a new user with validation
func NewUser(name, password string) (*User, error) {
	// TODO(human): Implement
	err := validate(name, password)
	return &User{
		Name:     name,
		password: password,
	}, err
}

// validate checks if name and password meet requirements
func validate(name, password string) error {
	// TODO(human): Implement

	user := []rune(name)
	if len(user) > cfg.maxNameLength || len(user) == 0 {
		return errors.New("invalid username length")
	}

	passwd := []rune(password)
	if len(passwd) < cfg.minPasswordLength {
		return errors.New("invalid password length")
	}

	return nil
}

// CheckPassword verifies if the provided password matches
func (u *User) CheckPassword(password string) bool {
	// TODO(human): Implement
	if password == u.password {
		return true
	}
	return false
}
