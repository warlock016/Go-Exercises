package cycles

// Dependency represents an import relationship
type Dependency struct {
	// TODO(human): Add From and To fields
	From, To string
}

// AnalyzeDependency creates a dependency relationship
func AnalyzeDependency(from, to string) Dependency {
	// TODO(human): Implement
	return Dependency{
		From: from,
		To:   to,
	}
}

// DetectCycle checks if dependencies form a cycle
func DetectCycle(deps []Dependency) bool {
	// TODO(human): Implement cycle detection

	graph := make(map[string][]string)
	for _, dep := range deps {
		graph[dep.From] = append(graph[dep.From], dep.To)
	}

	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	for pkg := range graph {
		if hasCycle(pkg, graph, visited, recStack) {
			return true
		}
	}

	return false
}

// SuggestFix provides advice on how to break an import cycle
func SuggestFix(cycle []string) string {
	// TODO(human): Implement
	return ""
}

func hasCycle(pkg string, graph map[string][]string, visited, recStack map[string]bool) bool {

	if recStack[pkg] {
		return true
	}

	if visited[pkg] {
		return false
	}

	visited[pkg] = true
	recStack[pkg] = true

	for _, dep := range graph[pkg] {
		// recursive function call
		if hasCycle(dep, graph, visited, recStack) {
			return true
		}
	}

	recStack[pkg] = false
	return false
}
