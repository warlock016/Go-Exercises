package visibility

// User represents a user in the system
type User struct {
	// TODO(human): Add Name field (exported) and password field (unexported)
}

// config holds package-level configuration
type config struct {
	// TODO(human): Add minPasswordLength and maxNameLength fields
}

// NewUser creates a new user with validation
func NewUser(name, password string) (*User, error) {
	// TODO(human): Implement
	return nil, nil
}

// validate checks if name and password meet requirements
func validate(name, password string) error {
	// TODO(human): Implement
	return nil
}

// CheckPassword verifies if the provided password matches
func (u *User) CheckPassword(password string) bool {
	// TODO(human): Implement
	return false
}
