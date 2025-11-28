package list_processing

// Flatten flattens a nested slice structure into a single slice
func Flatten(nested any) []int {
	// TODO(human): Implement

	result := []int{}

	if num, ok := nested.(int); ok {
		result = append(result, num)
	}

	if slice, ok := nested.([]any); ok {
		for _, element := range slice {
			result = append(result, Flatten(element)...)
		}
	}

	return result
}

// DeepSum calculates the sum of all integers in a nested structure
func DeepSum(nested any) int {
	// TODO(human): Implement

	result := 0

	if num, ok := nested.(int); ok {
		result += num
	}

	if slice, ok := nested.([]any); ok {
		for _, element := range slice {
			result += DeepSum(element)
		}
	}

	return result
}

// MaxDepth finds the maximum nesting depth
func MaxDepth(nested any) int {
	// TODO(human): Implement

	depth := 0
	if _, ok := nested.(int); ok {
		return 1
	}

	if slice, ok := nested.([]any); ok {
		for _, element := range slice {
			if _, ok := element.(int); !ok {
				childDepth := MaxDepth(element)
				depth = max(childDepth, depth)
			}
		}
	}

	return depth + 1
}

// FilterNested keeps only elements that satisfy the predicate
func FilterNested(nested any, predicate func(int) bool) []int {
	// TODO(human): Implement
	result := []int{}

	if val, ok := nested.(int); ok {
		if predicate(val) {
			result = append(result, val)
		}
	}

	if slice, ok := nested.([]any); ok {
		for _, val := range slice {
			result = append(result, FilterNested(val, predicate)...)
		}
	}

	return result
}
