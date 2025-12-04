package error_wrapping

// OpenAndReadFile opens a file and reads it, wrapping errors at each step
func OpenAndReadFile(filename string) (string, error) {
	// TODO(human): Implement
	return "", nil
}

// ProcessUser validates and saves user, wrapping errors with context
func ProcessUser(name, email string) error {
	// TODO(human): Implement
	return nil
}

// FetchAndParse fetches URL and parses JSON, wrapping errors
func FetchAndParse(url string) (map[string]interface{}, error) {
	// TODO(human): Implement
	return nil, nil
}

// UnwrapOnce unwraps one level of error wrapping
func UnwrapOnce(err error) error {
	// TODO(human): Implement
	return nil
}

// UnwrapAll repeatedly unwraps until reaching the root error
func UnwrapAll(err error) error {
	// TODO(human): Implement
	return nil
}
