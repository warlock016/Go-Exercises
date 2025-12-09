package dependency_injection

import "fmt"

// UserStore defines the interface for user storage
type UserStore interface {
	GetUser(id int) (string, error)
	SaveUser(id int, name string) error
}

// Logger defines the interface for logging
type Logger interface {
	Log(message string)
}

// UserService handles user operations with injected dependencies
type UserService struct {
	// TODO(human): Define fields
	store  UserStore
	logger Logger
}

// NewUserService creates a new UserService with dependencies
func NewUserService(store UserStore, logger Logger) *UserService {
	// TODO(human): Implement
	return &UserService{
		store:  store,
		logger: logger,
	}
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id int) (string, error) {
	// TODO(human): Implement
	s.logger.Log("Getting user")
	return s.store.GetUser(id)
}

// CreateUser creates a new user
func (s *UserService) CreateUser(id int, name string) error {
	// TODO(human): Implement
	if len(name) != 0 {
		s.logger.Log("Creating user")
		return s.store.SaveUser(id, name)
	} else {
		return fmt.Errorf("empty name")
	}
}
