package simple_queries

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			age INTEGER NOT NULL
		)
	`)
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	_, err = db.Exec(`
		INSERT INTO users (name, email, age) VALUES
		('Alice', 'alice@example.com', 30),
		('Bob', 'bob@example.com', 25),
		('Charlie', 'charlie@example.com', 35)
	`)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	return db
}

func TestGetUserByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	t.Run("existing user", func(t *testing.T) {
		user, err := GetUserByID(db, 1)
		if err != nil {
			t.Fatalf("GetUserByID() error = %v", err)
		}
		if user == nil {
			t.Fatal("GetUserByID() returned nil user")
		}
		if user.Name != "Alice" {
			t.Errorf("Name = %q, want %q", user.Name, "Alice")
		}
	})

	t.Run("non-existent user", func(t *testing.T) {
		user, err := GetUserByID(db, 999)
		if err != sql.ErrNoRows {
			t.Errorf("Expected sql.ErrNoRows, got %v", err)
		}
		if user != nil {
			t.Error("Expected nil user for non-existent ID")
		}
	})
}

func TestGetAllUsers(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	users, err := GetAllUsers(db)
	if err != nil {
		t.Fatalf("GetAllUsers() error = %v", err)
	}
	if len(users) != 3 {
		t.Errorf("Got %d users, want 3", len(users))
	}
}

func TestGetUsersByAge(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	users, err := GetUsersByAge(db, 28)
	if err != nil {
		t.Fatalf("GetUsersByAge() error = %v", err)
	}
	if len(users) != 2 {
		t.Errorf("Got %d users, want 2", len(users))
	}
}

func TestCountUsers(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	count, err := CountUsers(db)
	if err != nil {
		t.Fatalf("CountUsers() error = %v", err)
	}
	if count != 3 {
		t.Errorf("Count = %d, want 3", count)
	}
}
