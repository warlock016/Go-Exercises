package error_checking_patterns

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"basic division", 10, 2, 5, false},
		{"division by zero", 10, 0, 0, true},
		{"negative numbers", -10, 2, -5, false},
		{"fractional result", 5, 2, 2.5, false},
		{"divide zero", 0, 5, 0, false},
		{"both negative", -10, -2, 5, false},
		{"large numbers", 1000000, 1000, 1000, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide(%v, %v) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Divide(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestReadFileLength(t *testing.T) {
	// Create a temporary test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	content := []byte("Hello, World!")
	if err := os.WriteFile(testFile, content, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	tests := []struct {
		name     string
		filename string
		want     int64
		wantErr  bool
	}{
		{"existing file", testFile, int64(len(content)), false},
		{"nonexistent file", filepath.Join(tmpDir, "nonexistent.txt"), 0, true},
		{"empty path", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFileLength(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileLength(%v) error = %v, wantErr %v", tt.filename, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ReadFileLength(%v) = %v, want %v", tt.filename, got, tt.want)
			}
		})
	}
}

func TestParseAndDouble(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{"valid positive", "42", 84, false},
		{"valid negative", "-10", -20, false},
		{"zero", "0", 0, false},
		{"invalid input", "not a number", 0, true},
		{"empty string", "", 0, true},
		{"float string", "3.14", 0, true},
		{"large number", "1000000", 2000000, false},
		{"whitespace", "  42  ", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAndDouble(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAndDouble(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ParseAndDouble(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestChainOperations(t *testing.T) {
	tests := []struct {
		name       string
		a, b, c, d float64
		want       float64
		wantErr    bool
	}{
		{"all valid", 10, 2, 3, 5, 20, false}, // (10/2)*3 + 5 = 20
		{"division by zero", 10, 0, 3, 5, 0, true},
		{"zero result", 0, 2, 5, 0, 0, false}, // (0/2)*5 + 0 = 0
		{"negative numbers", -10, 2, -3, 5, 20, false}, // (-10/2)*-3 + 5 = 20
		{"fractional", 7, 2, 4, 1, 15, false}, // (7/2)*4 + 1 = 15
		{"multiply by zero", 10, 2, 0, 5, 5, false}, // (10/2)*0 + 5 = 5
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ChainOperations(tt.a, tt.b, tt.c, tt.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChainOperations(%v, %v, %v, %v) error = %v, wantErr %v",
					tt.a, tt.b, tt.c, tt.d, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ChainOperations(%v, %v, %v, %v) = %v, want %v",
					tt.a, tt.b, tt.c, tt.d, got, tt.want)
			}
		})
	}
}

func TestProcessFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Create test files
	file1 := filepath.Join(tmpDir, "file1.txt")
	file2 := filepath.Join(tmpDir, "file2.txt")
	file3 := filepath.Join(tmpDir, "file3.txt")

	os.WriteFile(file1, []byte("Hello"), 0644)      // 5 bytes
	os.WriteFile(file2, []byte("World!"), 0644)     // 6 bytes
	os.WriteFile(file3, []byte("Go"), 0644)         // 2 bytes

	tests := []struct {
		name      string
		filenames []string
		want      int64
		wantErr   bool
	}{
		{"single file", []string{file1}, 5, false},
		{"multiple files", []string{file1, file2, file3}, 13, false},
		{"empty list", []string{}, 0, false},
		{"nonexistent file", []string{file1, filepath.Join(tmpDir, "missing.txt")}, 0, true},
		{"all nonexistent", []string{filepath.Join(tmpDir, "missing.txt")}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProcessFiles(tt.filenames)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessFiles(%v) error = %v, wantErr %v", tt.filenames, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ProcessFiles(%v) = %v, want %v", tt.filenames, got, tt.want)
			}
		})
	}
}

func TestSafeIndexAccess(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}

	tests := []struct {
		name    string
		slice   []int
		index   int
		want    int
		wantErr bool
	}{
		{"first element", slice, 0, 10, false},
		{"middle element", slice, 2, 30, false},
		{"last element", slice, 4, 50, false},
		{"negative index", slice, -1, 0, true},
		{"index too large", slice, 5, 0, true},
		{"way out of bounds", slice, 100, 0, true},
		{"empty slice", []int{}, 0, 0, true},
		{"single element valid", []int{42}, 0, 42, false},
		{"single element invalid", []int{42}, 1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SafeIndexAccess(tt.slice, tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("SafeIndexAccess(%v, %v) error = %v, wantErr %v",
					tt.slice, tt.index, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("SafeIndexAccess(%v, %v) = %v, want %v",
					tt.slice, tt.index, got, tt.want)
			}
		})
	}
}
