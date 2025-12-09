package repository_errors

import "fmt"

// User represents a domain user
type User struct {
	ID    int
	Name  string
	Email string
}

// UserRepository handles user data access
type UserRepository struct {
	users map[int]User
}

// NewUserRepository creates a new UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[int]User),
	}
}

// FindByID returns user by ID or ErrNotFound
func (r *UserRepository) FindByID(id int) (User, error) {
	// TODO(human): Implement

	if user, exists := r.users[id]; exists {
		return user, nil
	}

	return User{}, fmt.Errorf("ErrNotFound")
}

// Create creates a user or returns ErrAlreadyExists
func (r *UserRepository) Create(user User) error {
	// TODO(human): Implement
	if _, exists := r.users[user.ID]; !exists {
		r.users[user.ID] = user
		return nil
	}
	return fmt.Errorf("user already exists")
}

// Update updates a user or returns ErrNotFound
func (r *UserRepository) Update(user User) error {
	// TODO(human): Implement
	if _, exists := r.users[user.ID]; !exists {
		return fmt.Errorf("user not found")
	}
	r.users[user.ID] = user
	return nil
}

// Delete deletes a user or returns ErrNotFound
func (r *UserRepository) Delete(id int) error {
	// TODO(human): Implement
	if _, exists := r.users[id]; !exists {
		return fmt.Errorf("user not found")
	}
	delete(r.users, id)
	return nil
}
