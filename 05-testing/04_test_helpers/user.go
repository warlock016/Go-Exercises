package user

import (
	"errors"
	"strings"
)

// User represents a user with name and email
type User struct {
	Name  string
	Email string
}

// NewUser creates a new user with validation
func NewUser(name, email string) (*User, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if !strings.Contains(email, "@") {
		return nil, errors.New("invalid email address")
	}

	return &User{
		Name:  name,
		Email: email,
	}, nil
}

// Validate checks if the user data is valid
func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name cannot be empty")
	}
	if !strings.Contains(u.Email, "@") {
		return errors.New("invalid email address")
	}
	return nil
}

// UpdateEmail updates the user's email with validation
func (u *User) UpdateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return errors.New("invalid email address")
	}
	u.Email = email
	return nil
}
