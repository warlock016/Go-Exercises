package error_wrapping

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOpenAndReadFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	content := "Hello, World!"
	os.WriteFile(testFile, []byte(content), 0644)

	tests := []struct {
		name     string
		filename string
		want     string
		wantErr  bool
		checkIs  error
	}{
		{"existing file", testFile, content, false, nil},
		{"nonexistent file", filepath.Join(tmpDir, "missing.txt"), "", true, os.ErrNotExist},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OpenAndReadFile(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenAndReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("OpenAndReadFile() = %q, want %q", got, tt.want)
			}
			if tt.checkIs != nil && !errors.Is(err, tt.checkIs) {
				t.Errorf("OpenAndReadFile() error chain doesn't contain %v", tt.checkIs)
			}
		})
	}
}

func TestProcessUser(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		email    string
		wantErr  bool
	}{
		{"valid user", "John", "john@example.com", false},
		{"empty name", "", "john@example.com", true},
		{"invalid email", "John", "notanemail", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ProcessUser(tt.userName, tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Check that error contains context
			if tt.wantErr && err != nil && !strings.Contains(err.Error(), "process") {
				t.Errorf("ProcessUser() error should contain context, got %v", err)
			}
		})
	}
}

func TestUnwrapOnce(t *testing.T) {
	inner := errors.New("inner error")
	wrapped := fmt.Errorf("outer error: %w", inner)

	tests := []struct {
		name string
		err  error
		want error
	}{
		{"nil error", nil, nil},
		{"unwrappable error", inner, nil},
		{"wrapped error", wrapped, inner},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnwrapOnce(tt.err)
			if got != tt.want {
				t.Errorf("UnwrapOnce() = %v, want %v", got, tt.want)
			}
		})
	}
}
