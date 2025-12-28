package imports

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	str "strings"
)

// StandardImport converts a map to JSON string using standard imports
func StandardImport(data map[string]int) string {
	// TODO(human): Implement using encoding/json
	res, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(res)
}

// AliasedImport converts text to uppercase and base64 encodes it using aliased imports
func AliasedImport(text string) (string, error) {
	// TODO(human): Implement using aliased strings and base64 packages
	upper := str.ToUpper(text)
	encoded := base64.StdEncoding.EncodeToString([]byte(upper))
	return encoded, nil
}

// GroupedImports generates n formatted strings
func GroupedImports(n int) []string {
	// TODO(human): Implement
	result := make([]string, n)

	for i := range result {
		result[i] = fmt.Sprintf("item-%d", i)
	}
	return result
}

// BlankImportExample returns an explanation of blank imports
func BlankImportExample() string {
	// TODO(human): Return a description of when/why to use blank imports
	return "Blank imports are used for side effects, like registering database drivers: import _ \"github.com/lib/pq\""
}
