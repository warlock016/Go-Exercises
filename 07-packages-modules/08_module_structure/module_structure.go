package structure

// PackageInfo represents information about a package
type PackageInfo struct {
	// TODO(human): Add fields: Name, Imports, ExportsTypes, IsInternal
}

// ModuleDesign represents a module's architecture
type ModuleDesign struct {
	// TODO(human): Add Packages field
}

// ValidateDesign checks for structural problems
func ValidateDesign(design ModuleDesign) []string {
	// TODO(human): Implement validation rules
	return nil
}

// DetectCycles finds import cycles in the design
func DetectCycles(design ModuleDesign) [][]string {
	// TODO(human): Implement cycle detection
	return nil
}

// CalculateCoupling measures how many packages this package depends on
func CalculateCoupling(pkg PackageInfo) int {
	// TODO(human): Implement
	return 0
}

// SuggestStructure recommends a package layout
func SuggestStructure(pkgCount int, hasAPI, hasDB bool) []string {
	// TODO(human): Implement
	return nil
}
