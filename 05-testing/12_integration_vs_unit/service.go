package service

import "fmt"

// Database interface (for mocking)
type Database interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

// Service depends on database
type Service struct {
	DB Database
}

// GetUser retrieves user from database
func (s *Service) GetUser(id string) (string, error) {
	return s.DB.Get("user:" + id)
}

// ValidateUser checks if user data is valid (pure function, no dependencies)
func ValidateUser(name, email string) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	return nil
}

// MockDB for testing
type MockDB struct {
	Data map[string]string
}

func (m *MockDB) Get(key string) (string, error) {
	if val, ok := m.Data[key]; ok {
		return val, nil
	}
	return "", fmt.Errorf("key not found: %s", key)
}

func (m *MockDB) Set(key, value string) error {
	m.Data[key] = value
	return nil
}
