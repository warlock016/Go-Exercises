package internal

// CanImport determines if importer can import target package
func CanImport(importer, target string) bool {
	// TODO(human): Implement internal/ visibility rules
	return false
}

// ExplainInternalRule returns an explanation of the internal/ directory rule
func ExplainInternalRule() string {
	// TODO(human): Return a clear explanation
	return ""
}

// ValidateStructure finds invalid imports in a list of import paths
func ValidateStructure(paths []string) []string {
	// TODO(human): Implement
	return nil
}
