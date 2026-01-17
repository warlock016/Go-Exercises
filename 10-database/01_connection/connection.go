package connection

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// OpenDatabase opens a SQLite database at the given path
func OpenDatabase(path string) (*sql.DB, error) {
	// TODO(human): Implement
	return nil, nil
}

// VerifyConnection checks if the database is reachable
func VerifyConnection(db *sql.DB) error {
	// TODO(human): Implement
	return nil
}

// CloseDatabase closes the database connection
func CloseDatabase(db *sql.DB) error {
	// TODO(human): Implement
	return nil
}

// WithDatabase opens a database, runs a function, then closes it
func WithDatabase(path string, fn func(*sql.DB) error) error {
	// TODO(human): Implement
	return nil
}
