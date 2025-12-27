package cycles

import (
	"strings"
	"testing"
)

func TestAnalyzeDependency(t *testing.T) {
	tests := []struct {
		from string
		to   string
	}{
		{"A", "B"},
		{"user", "order"},
		{"service", "repository"},
	}

	for _, tt := range tests {
		t.Run(tt.from+"->"+tt.to, func(t *testing.T) {
			dep := AnalyzeDependency(tt.from, tt.to)

			if dep.From != tt.from {
				t.Errorf("Dependency.From = %q, want %q", dep.From, tt.from)
			}
			if dep.To != tt.to {
				t.Errorf("Dependency.To = %q, want %q", dep.To, tt.to)
			}
		})
	}
}

func TestDetectCycleNoCycle(t *testing.T) {
	tests := []struct {
		name string
		deps []Dependency
	}{
		{
			name: "No dependencies",
			deps: []Dependency{},
		},
		{
			name: "Linear chain",
			deps: []Dependency{
				{"A", "B"},
				{"B", "C"},
				{"C", "D"},
			},
		},
		{
			name: "Tree structure",
			deps: []Dependency{
				{"A", "B"},
				{"A", "C"},
				{"B", "D"},
				{"C", "D"},
			},
		},
		{
			name: "Single dependency",
			deps: []Dependency{
				{"A", "B"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DetectCycle(tt.deps) {
				t.Errorf("DetectCycle() = true, expected false (no cycle)")
			}
		})
	}
}

func TestDetectCycleWithCycle(t *testing.T) {
	tests := []struct {
		name string
		deps []Dependency
	}{
		{
			name: "Direct cycle (2 packages)",
			deps: []Dependency{
				{"A", "B"},
				{"B", "A"},
			},
		},
		{
			name: "Triangle cycle (3 packages)",
			deps: []Dependency{
				{"A", "B"},
				{"B", "C"},
				{"C", "A"},
			},
		},
		{
			name: "Longer cycle (4 packages)",
			deps: []Dependency{
				{"A", "B"},
				{"B", "C"},
				{"C", "D"},
				{"D", "A"},
			},
		},
		{
			name: "Cycle with extra dependencies",
			deps: []Dependency{
				{"A", "B"},
				{"B", "C"},
				{"C", "A"}, // Cycle
				{"D", "A"}, // Extra
			},
		},
		{
			name: "Self-cycle",
			deps: []Dependency{
				{"A", "A"}, // Package imports itself
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !DetectCycle(tt.deps) {
				t.Errorf("DetectCycle() = false, expected true (cycle exists)")
			}
		})
	}
}

func TestSuggestFix(t *testing.T) {
	tests := []struct {
		name         string
		cycle        []string
		shouldMention string
	}{
		{
			name:         "No cycle",
			cycle:        []string{},
			shouldMention: "no",
		},
		{
			name:         "Two package cycle",
			cycle:        []string{"A", "B"},
			shouldMention: "merge", // or "extract"
		},
		{
			name:         "Three package cycle",
			cycle:        []string{"user", "order", "user"},
			shouldMention: "extract", // or "types"
		},
		{
			name:         "Longer cycle",
			cycle:        []string{"A", "B", "C", "D", "E"},
			shouldMention: "interface", // or "inversion"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suggestion := SuggestFix(tt.cycle)

			if suggestion == "" {
				t.Error("SuggestFix() returned empty string")
				return
			}

			lowerSuggestion := strings.ToLower(suggestion)
			lowerMention := strings.ToLower(tt.shouldMention)

			if !strings.Contains(lowerSuggestion, lowerMention) {
				t.Logf("Suggestion: %q", suggestion)
				t.Logf("Expected to mention: %q", tt.shouldMention)
				// Don't fail - there are multiple valid suggestions
				t.Logf("Note: This is informational. Your suggestion is valid if it addresses the cycle.")
			}
		})
	}
}

func TestRealWorldScenario(t *testing.T) {
	// Simulate a real scenario: user package and order package depend on each other
	deps := []Dependency{
		{"user", "order"},   // user needs order.Order type
		{"order", "user"},   // order needs user.User type
		{"order", "database"}, // order also uses database
	}

	if !DetectCycle(deps) {
		t.Error("Should detect cycle between user and order packages")
	}

	cycle := []string{"user", "order", "user"}
	suggestion := SuggestFix(cycle)

	t.Logf("Detected cycle: user <-> order")
	t.Logf("Suggestion: %s", suggestion)

	// Verify suggestion is helpful
	if len(suggestion) < 10 {
		t.Error("Suggestion should be descriptive")
	}
}

func TestComplexDependencyGraph(t *testing.T) {
	// More complex scenario
	deps := []Dependency{
		// Layer 1: models (no dependencies)

		// Layer 2: repository
		{"repository", "models"},

		// Layer 3: service
		{"service", "models"},
		{"service", "repository"},

		// Layer 4: handlers
		{"handlers", "service"},
		{"handlers", "models"},

		// No cycles - this is good architecture!
	}

	if DetectCycle(deps) {
		t.Error("This layered architecture should not have cycles")
	}
}

func TestBadArchitecture(t *testing.T) {
	// Bad architecture with cycles
	deps := []Dependency{
		{"handlers", "service"},
		{"service", "repository"},
		{"repository", "handlers"}, // BAD: repository shouldn't know about handlers
	}

	if !DetectCycle(deps) {
		t.Error("Should detect cycle in bad architecture")
	}
}
