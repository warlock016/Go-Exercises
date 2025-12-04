package dependency_injection

import (
	"errors"
	"testing"
)

// Mock implementations for testing
type MockStore struct {
	users      map[int]string
	getError   error
	saveError  error
	getCalled  bool
	saveCalled bool
}

func (m *MockStore) GetUser(id int) (string, error) {
	m.getCalled = true
	if m.getError != nil {
		return "", m.getError
	}
	if name, ok := m.users[id]; ok {
		return name, nil
	}
	return "", errors.New("user not found")
}

func (m *MockStore) SaveUser(id int, name string) error {
	m.saveCalled = true
	if m.saveError != nil {
		return m.saveError
	}
	if m.users == nil {
		m.users = make(map[int]string)
	}
	m.users[id] = name
	return nil
}

type MockLogger struct {
	messages []string
}

func (m *MockLogger) Log(message string) {
	m.messages = append(m.messages, message)
}

func TestNewUserService(t *testing.T) {
	store := &MockStore{}
	logger := &MockLogger{}

	service := NewUserService(store, logger)

	if service == nil {
		t.Fatal("NewUserService returned nil")
	}
}

func TestGetUserSuccess(t *testing.T) {
	store := &MockStore{
		users: map[int]string{1: "Alice", 2: "Bob"},
	}
	logger := &MockLogger{}
	service := NewUserService(store, logger)

	name, err := service.GetUser(1)

	if err != nil {
		t.Fatalf("GetUser returned error: %v", err)
	}

	if name != "Alice" {
		t.Errorf("GetUser(1) = %q, want %q", name, "Alice")
	}

	if !store.getCalled {
		t.Error("Store.GetUser was not called")
	}

	if len(logger.messages) == 0 {
		t.Error("Logger was not called")
	}
}

func TestGetUserNotFound(t *testing.T) {
	store := &MockStore{
		users: map[int]string{},
	}
	logger := &MockLogger{}
	service := NewUserService(store, logger)

	_, err := service.GetUser(999)

	if err == nil {
		t.Error("GetUser should return error for non-existent user")
	}
}

func TestGetUserStoreError(t *testing.T) {
	store := &MockStore{
		getError: errors.New("database error"),
	}
	logger := &MockLogger{}
	service := NewUserService(store, logger)

	_, err := service.GetUser(1)

	if err == nil {
		t.Error("GetUser should return error when store fails")
	}
}

func TestCreateUserSuccess(t *testing.T) {
	store := &MockStore{}
	logger := &MockLogger{}
	service := NewUserService(store, logger)

	err := service.CreateUser(1, "Charlie")

	if err != nil {
		t.Fatalf("CreateUser returned error: %v", err)
	}

	if !store.saveCalled {
		t.Error("Store.SaveUser was not called")
	}

	if len(logger.messages) == 0 {
		t.Error("Logger was not called")
	}

	// Verify user was saved
	if store.users[1] != "Charlie" {
		t.Errorf("User was not saved correctly, got %q", store.users[1])
	}
}

func TestCreateUserEmptyName(t *testing.T) {
	store := &MockStore{}
	logger := &MockLogger{}
	service := NewUserService(store, logger)

	err := service.CreateUser(1, "")

	if err == nil {
		t.Error("CreateUser should return error for empty name")
	}

	if store.saveCalled {
		t.Error("Store.SaveUser should not be called for invalid input")
	}
}

func TestCreateUserStoreError(t *testing.T) {
	store := &MockStore{
		saveError: errors.New("database error"),
	}
	logger := &MockLogger{}
	service := NewUserService(store, logger)

	err := service.CreateUser(1, "Dave")

	if err == nil {
		t.Error("CreateUser should return error when store fails")
	}
}

func TestDependencyInjectionIndependence(t *testing.T) {
	// Test that multiple services with different dependencies work independently
	store1 := &MockStore{users: map[int]string{1: "User1"}}
	logger1 := &MockLogger{}
	service1 := NewUserService(store1, logger1)

	store2 := &MockStore{users: map[int]string{1: "User2"}}
	logger2 := &MockLogger{}
	service2 := NewUserService(store2, logger2)

	name1, _ := service1.GetUser(1)
	name2, _ := service2.GetUser(1)

	if name1 != "User1" {
		t.Errorf("service1.GetUser(1) = %q, want User1", name1)
	}

	if name2 != "User2" {
		t.Errorf("service2.GetUser(1) = %q, want User2", name2)
	}

	// Verify independent logging
	if len(logger1.messages) != 1 || len(logger2.messages) != 1 {
		t.Error("Loggers should be independent")
	}
}

func TestMultipleOperations(t *testing.T) {
	store := &MockStore{}
	logger := &MockLogger{}
	service := NewUserService(store, logger)

	// Create users
	service.CreateUser(1, "Alice")
	service.CreateUser(2, "Bob")

	// Get users
	name1, _ := service.GetUser(1)
	name2, _ := service.GetUser(2)

	if name1 != "Alice" || name2 != "Bob" {
		t.Error("Multiple operations failed")
	}

	// Verify logging happened for all operations
	if len(logger.messages) != 4 { // 2 creates + 2 gets
		t.Errorf("Expected 4 log messages, got %d", len(logger.messages))
	}
}
