package internal

import (
	"reflect"
	"strings"
	"testing"
)

func TestCanImportNoInternal(t *testing.T) {
	// Regular packages can always be imported
	tests := []struct {
		importer string
		target   string
	}{
		{"myapp/handler", "myapp/models"},
		{"external/app", "myapp/public"},
		{"a/b/c", "x/y/z"},
	}

	for _, tt := range tests {
		t.Run(tt.importer+"->"+tt.target, func(t *testing.T) {
			if !CanImport(tt.importer, tt.target) {
				t.Errorf("CanImport(%q, %q) = false, want true (no internal/ restriction)",
					tt.importer, tt.target)
			}
		})
	}
}

func TestCanImportInternalAllowed(t *testing.T) {
	// These imports should be allowed
	tests := []struct {
		name     string
		importer string
		target   string
	}{
		{
			name:     "Parent can import internal",
			importer: "myapp",
			target:   "myapp/internal/auth",
		},
		{
			name:     "Sibling can import internal",
			importer: "myapp/handler",
			target:   "myapp/internal/auth",
		},
		{
			name:     "Nested sibling can import internal",
			importer: "myapp/api/v1",
			target:   "myapp/internal/database",
		},
		{
			name:     "Same level as internal",
			importer: "myapp/service",
			target:   "myapp/internal/config",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !CanImport(tt.importer, tt.target) {
				t.Errorf("CanImport(%q, %q) = false, want true (%s)",
					tt.importer, tt.target, tt.name)
			}
		})
	}
}

func TestCanImportInternalForbidden(t *testing.T) {
	// These imports should NOT be allowed
	tests := []struct {
		name     string
		importer string
		target   string
	}{
		{
			name:     "External package cannot import internal",
			importer: "external/app",
			target:   "myapp/internal/auth",
		},
		{
			name:     "Sibling tree cannot import internal",
			importer: "otherapp",
			target:   "myapp/internal/database",
		},
		{
			name:     "Cousin package cannot import internal",
			importer: "mymodule/other",
			target:   "mymodule/myapp/internal/auth",
		},
		{
			name:     "Unrelated cannot import internal",
			importer: "completely/different",
			target:   "myapp/internal/config",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if CanImport(tt.importer, tt.target) {
				t.Errorf("CanImport(%q, %q) = true, want false (%s)",
					tt.importer, tt.target, tt.name)
			}
		})
	}
}

func TestExplainInternalRule(t *testing.T) {
	explanation := ExplainInternalRule()

	if explanation == "" {
		t.Fatal("ExplainInternalRule() returned empty string")
	}

	// Should mention key concepts
	lowerExpl := strings.ToLower(explanation)
	keywords := []string{"internal", "import", "package"}

	for _, keyword := range keywords {
		if !strings.Contains(lowerExpl, keyword) {
			t.Errorf("Explanation should mention %q, got: %q", keyword, explanation)
		}
	}

	t.Logf("Explanation: %s", explanation)
}

func TestValidateStructureNoViolations(t *testing.T) {
	paths := []string{
		"myapp imports myapp/internal/auth",
		"myapp/handler imports myapp/internal/auth",
		"external imports public/api",
	}

	violations := ValidateStructure(paths)

	if len(violations) != 0 {
		t.Errorf("ValidateStructure() found violations in valid structure: %v", violations)
	}
}

func TestValidateStructureWithViolations(t *testing.T) {
	paths := []string{
		"myapp imports myapp/internal/auth",              // Valid
		"external imports myapp/internal/auth",           // INVALID
		"otherapp/handler imports myapp/internal/config", // INVALID
	}

	violations := ValidateStructure(paths)

	expectedCount := 2
	if len(violations) != expectedCount {
		t.Errorf("ValidateStructure() found %d violations, want %d: %v",
			len(violations), expectedCount, violations)
	}

	// Should contain the invalid imports
	expected := []string{
		"external imports myapp/internal/auth",
		"otherapp/handler imports myapp/internal/config",
	}

	if !reflect.DeepEqual(violations, expected) {
		t.Errorf("ValidateStructure() = %v, want %v", violations, expected)
	}
}

func TestValidateStructureMixedPaths(t *testing.T) {
	paths := []string{
		"myapp/api imports myapp/internal/database",      // Valid
		"myapp/service imports myapp/internal/config",    // Valid
		"external/client imports myapp/internal/auth",    // Invalid
		"myapp/handler imports public/utils",             // Valid (no internal)
		"other/app imports myapp/internal/secret",        // Invalid
	}

	violations := ValidateStructure(paths)

	expectedCount := 2
	if len(violations) != expectedCount {
		t.Errorf("ValidateStructure() found %d violations, want %d: %v",
			len(violations), expectedCount, violations)
	}
}

func TestRealWorldScenario(t *testing.T) {
	// Simulating a real project structure
	scenarios := []struct {
		importer string
		target   string
		allowed  bool
		reason   string
	}{
		{
			importer: "github.com/user/proj/cmd/server",
			target:   "github.com/user/proj/internal/config",
			allowed:  true,
			reason:   "cmd/server is in proj/, can access proj/internal/",
		},
		{
			importer: "github.com/user/proj/pkg/client",
			target:   "github.com/user/proj/internal/auth",
			allowed:  true,
			reason:   "pkg/client is in proj/, can access proj/internal/",
		},
		{
			importer: "github.com/other/project",
			target:   "github.com/user/proj/internal/secret",
			allowed:  false,
			reason:   "external module cannot access proj/internal/",
		},
		{
			importer: "github.com/user/proj/cmd/server",
			target:   "github.com/user/proj/pkg/models",
			allowed:  true,
			reason:   "no internal/, always allowed",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.reason, func(t *testing.T) {
			got := CanImport(scenario.importer, scenario.target)
			if got != scenario.allowed {
				t.Errorf("CanImport(%q, %q) = %v, want %v\nReason: %s",
					scenario.importer, scenario.target, got, scenario.allowed, scenario.reason)
			}
		})
	}
}

func TestInternalAtDifferentLevels(t *testing.T) {
	// Projects can have internal/ at different levels
	tests := []struct {
		name     string
		importer string
		target   string
		allowed  bool
	}{
		{
			name:     "Root level internal",
			importer: "myapp/service",
			target:   "myapp/internal/shared",
			allowed:  true,
		},
		{
			name:     "Nested internal",
			importer: "myapp/api/handler",
			target:   "myapp/api/internal/middleware",
			allowed:  true,
		},
		{
			name:     "Cannot access sibling's internal",
			importer: "myapp/service",
			target:   "myapp/api/internal/middleware",
			allowed:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanImport(tt.importer, tt.target)
			if got != tt.allowed {
				t.Errorf("CanImport(%q, %q) = %v, want %v",
					tt.importer, tt.target, got, tt.allowed)
			}
		})
	}
}
