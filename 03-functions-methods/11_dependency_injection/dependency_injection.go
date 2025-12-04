package dependency_injection

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
}

// NewUserService creates a new UserService with dependencies
func NewUserService(store UserStore, logger Logger) *UserService {
	// TODO(human): Implement
	return nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id int) (string, error) {
	// TODO(human): Implement
	return "", nil
}

// CreateUser creates a new user
func (s *UserService) CreateUser(id int, name string) error {
	// TODO(human): Implement
	return nil
}
