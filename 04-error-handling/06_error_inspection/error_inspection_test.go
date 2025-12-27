package error_inspection

import (
	"errors"
	"fmt"
	"testing"
)

func TestCheckErrorType(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"not found sentinel", ErrNotFound, "not_found"},
		{"timeout sentinel", ErrTimeout, "timeout"},
		{"nil error", nil, "no_error"},
		{"wrapped not found", fmt.Errorf("db error: %w", ErrNotFound), "not_found"},
		{"wrapped timeout", fmt.Errorf("connection: %w", ErrTimeout), "timeout"},
		{"double wrapped not found", fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", ErrNotFound)), "not_found"},
		{"unknown error", errors.New("something else"), "unknown"},
		{"permission error", ErrPermission, "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckErrorType(tt.err)
			if got != tt.want {
				t.Errorf("CheckErrorType() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestExtractHTTPStatus(t *testing.T) {
	tests := []struct {
		name      string
		err       error
		wantCode  int
		wantFound bool
	}{
		{
			name:      "direct HTTPError 404",
			err:       &HTTPError{StatusCode: 404, Message: "Not Found"},
			wantCode:  404,
			wantFound: true,
		},
		{
			name:      "direct HTTPError 500",
			err:       &HTTPError{StatusCode: 500, Message: "Internal Server Error"},
			wantCode:  500,
			wantFound: true,
		},
		{
			name:      "wrapped HTTPError",
			err:       fmt.Errorf("request failed: %w", &HTTPError{StatusCode: 403, Message: "Forbidden"}),
			wantCode:  403,
			wantFound: true,
		},
		{
			name:      "double wrapped HTTPError",
			err:       fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", &HTTPError{StatusCode: 401, Message: "Unauthorized"})),
			wantCode:  401,
			wantFound: true,
		},
		{
			name:      "nil error",
			err:       nil,
			wantCode:  0,
			wantFound: false,
		},
		{
			name:      "non-HTTP error",
			err:       errors.New("generic error"),
			wantCode:  0,
			wantFound: false,
		},
		{
			name:      "sentinel error",
			err:       ErrNotFound,
			wantCode:  0,
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, found := ExtractHTTPStatus(tt.err)
			if code != tt.wantCode {
				t.Errorf("ExtractHTTPStatus() code = %d, want %d", code, tt.wantCode)
			}
			if found != tt.wantFound {
				t.Errorf("ExtractHTTPStatus() found = %v, want %v", found, tt.wantFound)
			}
		})
	}
}

func TestIsTemporaryError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "direct temporary error",
			err:  &TemporaryError{Err: errors.New("network hiccup")},
			want: true,
		},
		{
			name: "wrapped temporary error",
			err:  fmt.Errorf("operation failed: %w", &TemporaryError{Err: errors.New("retry later")}),
			want: true,
		},
		{
			name: "timeout is temporary",
			err:  ErrTimeout,
			want: true,
		},
		{
			name: "not found is not temporary",
			err:  ErrNotFound,
			want: false,
		},
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "generic error",
			err:  errors.New("some error"),
			want: false,
		},
		{
			name: "HTTP 503 is temporary",
			err:  &HTTPError{StatusCode: 503, Message: "Service Unavailable"},
			want: true,
		},
		{
			name: "HTTP 404 is not temporary",
			err:  &HTTPError{StatusCode: 404, Message: "Not Found"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsTemporaryError(tt.err)
			if got != tt.want {
				t.Errorf("IsTemporaryError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleBasedOnType(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "not found error",
			err:  ErrNotFound,
			want: "not_found: return 404",
		},
		{
			name: "timeout error",
			err:  ErrTimeout,
			want: "timeout: retry operation",
		},
		{
			name: "HTTP 500",
			err:  &HTTPError{StatusCode: 500, Message: "Server Error"},
			want: "http_error: status 500",
		},
		{
			name: "HTTP 403",
			err:  &HTTPError{StatusCode: 403, Message: "Forbidden"},
			want: "http_error: status 403",
		},
		{
			name: "temporary error",
			err:  &TemporaryError{Err: errors.New("network")},
			want: "temporary: will retry",
		},
		{
			name: "nil error",
			err:  nil,
			want: "no_error: success",
		},
		{
			name: "unknown error",
			err:  errors.New("mystery"),
			want: "unknown: log and alert",
		},
		{
			name: "wrapped not found",
			err:  fmt.Errorf("db: %w", ErrNotFound),
			want: "not_found: return 404",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HandleBasedOnType(tt.err)
			if got != tt.want {
				t.Errorf("HandleBasedOnType() = %q, want %q", got, tt.want)
			}
		})
	}
}
