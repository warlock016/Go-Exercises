package imports

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestStandardImport(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]int
		want  map[string]int // We'll parse it back to compare
	}{
		{
			name:  "Simple map",
			input: map[string]int{"a": 1, "b": 2},
			want:  map[string]int{"a": 1, "b": 2},
		},
		{
			name:  "Empty map",
			input: map[string]int{},
			want:  map[string]int{},
		},
		{
			name:  "Single entry",
			input: map[string]int{"count": 42},
			want:  map[string]int{"count": 42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StandardImport(tt.input)

			// Verify it's valid JSON by unmarshaling
			var result map[string]int
			err := json.Unmarshal([]byte(got), &result)
			if err != nil {
				t.Errorf("StandardImport() returned invalid JSON: %v, error: %v", got, err)
			}

			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("StandardImport(%v) = %v, want %v", tt.input, result, tt.want)
			}
		})
	}
}

func TestAliasedImport(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Simple text",
			input:   "hello",
			wantErr: false,
		},
		{
			name:    "Empty string",
			input:   "",
			wantErr: false,
		},
		{
			name:    "Mixed case",
			input:   "Hello World",
			wantErr: false,
		},
		{
			name:    "Special characters",
			input:   "test!@#",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AliasedImport(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Error("AliasedImport() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("AliasedImport(%q) unexpected error: %v", tt.input, err)
				return
			}

			// Decode base64 to verify
			decoded, err := base64.StdEncoding.DecodeString(got)
			if err != nil {
				t.Errorf("AliasedImport(%q) returned invalid base64: %v", tt.input, err)
				return
			}

			// Should be uppercase version of input
			expected := strings.ToUpper(tt.input)
			if string(decoded) != expected {
				t.Errorf("AliasedImport(%q) decoded = %q, want %q", tt.input, string(decoded), expected)
			}
		})
	}
}

func TestGroupedImports(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int // Expected length
	}{
		{
			name:  "Zero items",
			input: 0,
			want:  0,
		},
		{
			name:  "Single item",
			input: 1,
			want:  1,
		},
		{
			name:  "Multiple items",
			input: 5,
			want:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupedImports(tt.input)

			if len(got) != tt.want {
				t.Errorf("GroupedImports(%d) returned %d items, want %d", tt.input, len(got), tt.want)
			}

			// Verify format: "item-N"
			for i, item := range got {
				expected := "item-" + string(rune('0'+i))
				if i < 10 && item != expected {
					// Simple check for single digits
					if !strings.HasPrefix(item, "item-") {
						t.Errorf("GroupedImports(%d)[%d] = %q, should start with 'item-'", tt.input, i, item)
					}
				}
			}
		})
	}
}

func TestBlankImportExample(t *testing.T) {
	got := BlankImportExample()

	if got == "" {
		t.Error("BlankImportExample() returned empty string")
	}

	// Should mention key concepts
	lowerGot := strings.ToLower(got)
	keywords := []string{"blank", "import", "side", "effect"}

	foundKeywords := 0
	for _, keyword := range keywords {
		if strings.Contains(lowerGot, keyword) {
			foundKeywords++
		}
	}

	if foundKeywords < 2 {
		t.Errorf("BlankImportExample() should mention blank imports and side effects, got: %q", got)
	}
}

func TestImportOrganization(t *testing.T) {
	// This test is informational - it verifies that the student
	// understands import organization by checking their code compiles
	// and follows conventions.

	t.Log("Import organization best practices:")
	t.Log("1. Group standard library imports together")
	t.Log("2. Separate external packages with a blank line")
	t.Log("3. Separate internal packages with a blank line")
	t.Log("4. Use aliases when package names conflict or for clarity")
	t.Log("5. Use blank imports (_) only for side effects")
}
