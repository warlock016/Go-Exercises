package errors

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// Network/Transport errors
	ErrTimeout error = errors.New("request timed out")
	ErrNetwork error = errors.New("network error")

	// API-specific errors
	ErrNotFound     error = errors.New("not found")
	ErrRateLimited  error = errors.New("rate limited")
	ErrUnauthorized error = errors.New("unauthorized")

	// Input errors
	ErrInvalidInput error = errors.New("invalid input")
)

// OpenMeteo API errors
type WeatherAPIError struct {
	StatusCode int
	Message    string
	Endpoint   string
	Latitude   float64
	Longitude  float64
	ErrorType  error
}

func (e *WeatherAPIError) Error() string {
	return fmt.Sprintf("%d, %s, %s, %.2f, %.2f", e.StatusCode, e.Message, e.Endpoint, e.Latitude, e.Longitude)
}

func (e *WeatherAPIError) Unwrap() error {
	return e.ErrorType
}

// Geocoding API errors
type GeocodeError struct {
	StatusCode int
	Message    string
	Query      string
	ErrorType  error
}

func (e *GeocodeError) Error() string {
	return fmt.Sprintf("%d, %s, %s", e.StatusCode, e.Message, e.Query)
}

func (e *GeocodeError) Unwrap() error {
	return e.ErrorType
}

// Validation Errors
type FieldError struct {
	Field   string
	Message string
}

type ValidationErrors struct {
	Errors []FieldError
}

func (e *ValidationErrors) Error() string {

	var result strings.Builder

	for _, v := range e.Errors {
		result.WriteString("{" + v.Field + "}: {" + v.Message + "}\n")
	}

	return result.String()
}

func (e *ValidationErrors) Add(field, message string) {
	e.Errors = append(e.Errors, FieldError{Field: field, Message: message})
}

func (e *ValidationErrors) HasErrors() bool {
	if len(e.Errors) == 0 {
		return false
	}
	return true
}
