package errors

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNotFound     error = errors.New("not found")
	ErrInvalidInput error = errors.New("invalid input")
)

// Error types:
// 1. Missing file (incorrect path or missing file)
// 2. Empty or malformed file (read error or "")
// 3.
// type APIError struct {
// }

// type HTTP struct {
// }

// type ParseError struct {
// }

type FieldError struct {
	Stage   string
	Message string
	Line    int
	Column  int
	Value   string
}

type ProcessingErrors struct {
	Errors   []FieldError
	Warnings []FieldError
}

// returns content of all errors as string: message, line & position
func (e *ProcessingErrors) Error() string {
	var result strings.Builder

	if len(e.Errors) != 0 {
		result.WriteString("Fatal Errors:\n")
		for _, v := range e.Errors {
			fmt.Fprintf(&result, "ERR: %s: %s: %s ln: %d col: %d\n", v.Stage, v.Message, v.Value, v.Line, v.Column)
		}
	}
	return result.String()
}

// returns content of all errors as string: message, line & position
func (e *ProcessingErrors) Warning() string {
	var result strings.Builder

	if len(e.Warnings) != 0 {
		result.WriteString("Warnings\n")
		for _, v := range e.Warnings {
			fmt.Fprintf(&result, "WRN: %s: %s: %s ln: %d col: %d\n", v.Stage, v.Message, v.Value, v.Line, v.Column)
		}
	}
	return result.String()
}

func (e *ProcessingErrors) Summary() string {
	var result strings.Builder

	if e.HasFatalErrors() {
		result.WriteString(e.Error())
	}
	if e.HasWarnings() {
		result.WriteString(e.Warning())
	}
	fmt.Fprintf(&result, "\n%d errors, %d warnings\n", len(e.Errors), len(e.Warnings))

	return result.String()
}

// adds error content to current error list
func (e *ProcessingErrors) AddError(stage, msg, value string, ln, col int) {
	e.Errors = append(e.Errors, FieldError{Stage: stage, Message: msg, Line: ln, Column: col, Value: value})
}

func (e *ProcessingErrors) AddWarning(stage, msg, value string, ln, col int) {
	e.Warnings = append(e.Warnings, FieldError{Stage: stage, Message: msg, Line: ln, Column: col, Value: value})
}

// checks if there are any errors in the list
func (e *ProcessingErrors) HasWarnings() bool {
	return len(e.Warnings) > 0
}

func (e *ProcessingErrors) HasFatalErrors() bool {
	return len(e.Errors) > 0
}
