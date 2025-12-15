package client

import "fmt"

type HAError struct {
	Stage   string
	Message string
	Errors  []error
}

type HTTPError struct {
}

type PaseError struct {
}

func Error(err error) string {
	return fmt.Sprintf("%v", err)
}

func Unwrap(err error) error {
	return Unwrap(err)
}
