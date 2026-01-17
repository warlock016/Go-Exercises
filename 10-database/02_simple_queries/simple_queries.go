package simple_queries

import "database/sql"

// User represents a user in the database
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

// GetUserByID retrieves a single user by ID
func GetUserByID(db *sql.DB, id int) (*User, error) {
	// TODO(human): Implement
	return nil, nil
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(db *sql.DB) ([]User, error) {
	// TODO(human): Implement
	return nil, nil
}

// GetUsersByAge retrieves users with age greater than minAge
func GetUsersByAge(db *sql.DB, minAge int) ([]User, error) {
	// TODO(human): Implement
	return nil, nil
}

// CountUsers returns the total number of users
func CountUsers(db *sql.DB) (int, error) {
	// TODO(human): Implement
	return 0, nil
}
