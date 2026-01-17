package connection

import (
	"errors"
	"testing"
)

func TestOpenDatabase(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "in-memory database",
			path:    ":memory:",
			wantErr: false,
		},
		{
			name:    "file database",
			path:    "test.db",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := OpenDatabase(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if db == nil && !tt.wantErr {
				t.Error("OpenDatabase() returned nil db")
				return
			}
			if db != nil {
				db.Close()
			}
		})
	}
}

func TestVerifyConnection(t *testing.T) {
	db, err := OpenDatabase(":memory:")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	t.Run("healthy connection", func(t *testing.T) {
		if err := VerifyConnection(db); err != nil {
			t.Errorf("VerifyConnection() error = %v, want nil", err)
		}
	})

	t.Run("closed connection", func(t *testing.T) {
		closedDB, _ := OpenDatabase(":memory:")
		closedDB.Close()

		if err := VerifyConnection(closedDB); err == nil {
			t.Error("VerifyConnection() on closed db should return error")
		}
	})
}

func TestCloseDatabase(t *testing.T) {
	db, err := OpenDatabase(":memory:")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	if err := CloseDatabase(db); err != nil {
		t.Errorf("CloseDatabase() error = %v", err)
	}

	// Verify it's actually closed
	if err := db.Ping(); err == nil {
		t.Error("Database should be closed")
	}
}

func TestWithDatabase(t *testing.T) {
	t.Run("successful operation", func(t *testing.T) {
		called := false
		err := WithDatabase(":memory:", func(db *sql.DB) error {
			called = true
			return nil
		})

		if err != nil {
			t.Errorf("WithDatabase() error = %v", err)
		}
		if !called {
			t.Error("Function was not called")
		}
	})

	t.Run("function returns error", func(t *testing.T) {
		expectedErr := errors.New("test error")
		err := WithDatabase(":memory:", func(db *sql.DB) error {
			return expectedErr
		})

		if err != expectedErr {
			t.Errorf("WithDatabase() error = %v, want %v", err, expectedErr)
		}
	})

	t.Run("database is usable in function", func(t *testing.T) {
		err := WithDatabase(":memory:", func(db *sql.DB) error {
			return db.Ping()
		})

		if err != nil {
			t.Errorf("Database not usable: %v", err)
		}
	})
}
