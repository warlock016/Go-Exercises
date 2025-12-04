package api_error_format

// ProblemDetail represents RFC 7807 Problem Details
type ProblemDetail struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

// NewProblemDetail creates a problem detail from an error
func NewProblemDetail(err error, requestPath string) ProblemDetail {
	// TODO(human): Implement
	return ProblemDetail{}
}

// ValidationProblemDetail creates a problem detail for validation errors
func ValidationProblemDetail(errors []string, requestPath string) ProblemDetail {
	// TODO(human): Implement
	return ProblemDetail{}
}
