package structure

import (
	"reflect"
	"strings"
	"testing"
)

func TestPackageInfoCreation(t *testing.T) {
	pkg := PackageInfo{
		Name:         "models",
		Imports:      []string{},
		ExportsTypes: true,
		IsInternal:   false,
	}

	if pkg.Name != "models" {
		t.Errorf("PackageInfo.Name = %q, want %q", pkg.Name, "models")
	}
	if len(pkg.Imports) != 0 {
		t.Errorf("PackageInfo.Imports = %v, want []", pkg.Imports)
	}
	if !pkg.ExportsTypes {
		t.Error("PackageInfo.ExportsTypes should be true")
	}
	if pkg.IsInternal {
		t.Error("PackageInfo.IsInternal should be false")
	}
}

func TestValidateDesignGoodStructure(t *testing.T) {
	// Well-designed layered architecture
	design := ModuleDesign{
		Packages: []PackageInfo{
			{
				Name:         "models",
				Imports:      []string{},
				ExportsTypes: true,
				IsInternal:   false,
			},
			{
				Name:         "repository",
				Imports:      []string{"models"},
				ExportsTypes: false,
				IsInternal:   true,
			},
			{
				Name:         "service",
				Imports:      []string{"models", "repository"},
				ExportsTypes: false,
				IsInternal:   true,
			},
			{
				Name:         "handlers",
				Imports:      []string{"models", "service"},
				ExportsTypes: false,
				IsInternal:   false,
			},
		},
	}

	problems := ValidateDesign(design)

	if len(problems) != 0 {
		t.Errorf("ValidateDesign() found problems in good structure: %v", problems)
	}
}

func TestValidateDesignHighCoupling(t *testing.T) {
	// Package with too many dependencies
	design := ModuleDesign{
		Packages: []PackageInfo{
			{
				Name:         "badpackage",
				Imports:      []string{"a", "b", "c", "d", "e", "f", "g"},
				ExportsTypes: false,
				IsInternal:   false,
			},
		},
	}

	problems := ValidateDesign(design)

	if len(problems) == 0 {
		t.Error("ValidateDesign() should detect high coupling")
	}

	// Should mention coupling or dependencies
	found := false
	for _, problem := range problems {
		if strings.Contains(strings.ToLower(problem), "dependenc") ||
			strings.Contains(strings.ToLower(problem), "coupling") {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("ValidateDesign() should mention dependencies/coupling, got: %v", problems)
	}
}

func TestDetectCyclesNoCycle(t *testing.T) {
	// Linear dependency chain (no cycle)
	design := ModuleDesign{
		Packages: []PackageInfo{
			{Name: "A", Imports: []string{"B"}},
			{Name: "B", Imports: []string{"C"}},
			{Name: "C", Imports: []string{}},
		},
	}

	cycles := DetectCycles(design)

	if len(cycles) != 0 {
		t.Errorf("DetectCycles() found cycles when there are none: %v", cycles)
	}
}

func TestDetectCyclesSimpleCycle(t *testing.T) {
	// A -> B -> A (cycle)
	design := ModuleDesign{
		Packages: []PackageInfo{
			{Name: "A", Imports: []string{"B"}},
			{Name: "B", Imports: []string{"A"}},
		},
	}

	cycles := DetectCycles(design)

	if len(cycles) == 0 {
		t.Error("DetectCycles() should detect A <-> B cycle")
	}

	// Should find a cycle involving A and B
	foundCycle := false
	for _, cycle := range cycles {
		hasA := false
		hasB := false
		for _, pkg := range cycle {
			if pkg == "A" {
				hasA = true
			}
			if pkg == "B" {
				hasB = true
			}
		}
		if hasA && hasB {
			foundCycle = true
			break
		}
	}

	if !foundCycle {
		t.Errorf("DetectCycles() should find cycle with A and B, got: %v", cycles)
	}
}

func TestDetectCyclesComplexCycle(t *testing.T) {
	// A -> B -> C -> A (triangle cycle)
	design := ModuleDesign{
		Packages: []PackageInfo{
			{Name: "A", Imports: []string{"B"}},
			{Name: "B", Imports: []string{"C"}},
			{Name: "C", Imports: []string{"A"}},
		},
	}

	cycles := DetectCycles(design)

	if len(cycles) == 0 {
		t.Error("DetectCycles() should detect A -> B -> C -> A cycle")
	}
}

func TestCalculateCoupling(t *testing.T) {
	tests := []struct {
		name    string
		pkg     PackageInfo
		want    int
	}{
		{
			name: "No dependencies",
			pkg: PackageInfo{
				Name:    "models",
				Imports: []string{},
			},
			want: 0,
		},
		{
			name: "One dependency",
			pkg: PackageInfo{
				Name:    "repository",
				Imports: []string{"models"},
			},
			want: 1,
		},
		{
			name: "Multiple dependencies",
			pkg: PackageInfo{
				Name:    "service",
				Imports: []string{"models", "repository", "config"},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateCoupling(tt.pkg)
			if got != tt.want {
				t.Errorf("CalculateCoupling(%v) = %d, want %d", tt.pkg.Name, got, tt.want)
			}
		})
	}
}

func TestSuggestStructureMinimal(t *testing.T) {
	// Simple app: no API, no DB
	structure := SuggestStructure(3, false, false)

	if len(structure) == 0 {
		t.Error("SuggestStructure() should return at least some packages")
	}

	// Should at least suggest cmd/
	hasCorePackages := false
	for _, pkg := range structure {
		if strings.Contains(pkg, "cmd") || strings.Contains(pkg, "models") {
			hasCorePackages = true
			break
		}
	}

	if !hasCorePackages {
		t.Errorf("SuggestStructure() should include core packages, got: %v", structure)
	}
}

func TestSuggestStructureWithAPI(t *testing.T) {
	structure := SuggestStructure(5, true, false)

	// Should suggest API-related packages
	hasAPI := false
	for _, pkg := range structure {
		if strings.Contains(strings.ToLower(pkg), "api") ||
			strings.Contains(strings.ToLower(pkg), "handler") {
			hasAPI = true
			break
		}
	}

	if !hasAPI {
		t.Errorf("SuggestStructure(hasAPI=true) should include API packages, got: %v", structure)
	}
}

func TestSuggestStructureWithDB(t *testing.T) {
	structure := SuggestStructure(5, false, true)

	// Should suggest database-related packages
	hasDB := false
	for _, pkg := range structure {
		lowerPkg := strings.ToLower(pkg)
		if strings.Contains(lowerPkg, "database") ||
			strings.Contains(lowerPkg, "repository") ||
			strings.Contains(lowerPkg, "repo") {
			hasDB = true
			break
		}
	}

	if !hasDB {
		t.Errorf("SuggestStructure(hasDB=true) should include DB packages, got: %v", structure)
	}
}

func TestSuggestStructureComplete(t *testing.T) {
	structure := SuggestStructure(8, true, true)

	// Should have a comprehensive structure
	if len(structure) < 4 {
		t.Errorf("SuggestStructure(8, true, true) should suggest at least 4 packages, got %d: %v",
			len(structure), structure)
	}

	// Check for key layers
	layers := map[string]bool{
		"models":     false,
		"internal":   false,
		"api/cmd":    false, // Either API or cmd
	}

	for _, pkg := range structure {
		lowerPkg := strings.ToLower(pkg)
		if strings.Contains(lowerPkg, "model") {
			layers["models"] = true
		}
		if strings.Contains(lowerPkg, "internal") {
			layers["internal"] = true
		}
		if strings.Contains(lowerPkg, "api") || strings.Contains(lowerPkg, "cmd") {
			layers["api/cmd"] = true
		}
	}

	for layer, found := range layers {
		if !found {
			t.Logf("Warning: SuggestStructure() might want to include %s layer", layer)
		}
	}
}

func TestRealWorldDesign(t *testing.T) {
	// Simulate a real web application structure
	design := ModuleDesign{
		Packages: []PackageInfo{
			{
				Name:         "models",
				Imports:      []string{},
				ExportsTypes: true,
				IsInternal:   false,
			},
			{
				Name:         "database",
				Imports:      []string{"models"},
				ExportsTypes: false,
				IsInternal:   true,
			},
			{
				Name:         "repository",
				Imports:      []string{"models", "database"},
				ExportsTypes: false,
				IsInternal:   true,
			},
			{
				Name:         "service",
				Imports:      []string{"models", "repository"},
				ExportsTypes: false,
				IsInternal:   true,
			},
			{
				Name:         "handlers",
				Imports:      []string{"models", "service"},
				ExportsTypes: false,
				IsInternal:   false,
			},
		},
	}

	problems := ValidateDesign(design)
	cycles := DetectCycles(design)

	if len(problems) > 0 {
		t.Logf("Design problems: %v", problems)
	}

	if len(cycles) > 0 {
		t.Errorf("Real-world design should not have cycles: %v", cycles)
	}

	// Check coupling
	for _, pkg := range design.Packages {
		coupling := CalculateCoupling(pkg)
		t.Logf("Package %s has coupling of %d", pkg.Name, coupling)

		if coupling > 5 {
			t.Errorf("Package %s has high coupling (%d dependencies)", pkg.Name, coupling)
		}
	}
}

func TestModuleDesignEquality(t *testing.T) {
	design1 := ModuleDesign{
		Packages: []PackageInfo{
			{Name: "A", Imports: []string{"B"}},
		},
	}

	design2 := ModuleDesign{
		Packages: []PackageInfo{
			{Name: "A", Imports: []string{"B"}},
		},
	}

	if !reflect.DeepEqual(design1, design2) {
		t.Error("Identical ModuleDesigns should be equal")
	}
}
