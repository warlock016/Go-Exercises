package error_creation

import (
	"strings"
	"testing"
)

func TestValidateAge(t *testing.T) {
	tests := []struct {
		name    string
		age     int
		wantErr bool
		contains string
	}{
		{"valid age 0", 0, false, ""},
		{"valid age 25", 25, false, ""},
		{"valid age 150", 150, false, ""},
		{"negative age", -1, true, "negative"},
		{"negative age -100", -100, true, "negative"},
		{"too old 151", 151, true, "must be between"},
		{"way too old", 500, true, "must be between"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAge(tt.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAge(%d) error = %v, wantErr %v", tt.age, err, tt.wantErr)
				return
			}
			if tt.wantErr && !strings.Contains(err.Error(), tt.contains) {
				t.Errorf("ValidateAge(%d) error = %q, want to contain %q", tt.age, err.Error(), tt.contains)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{"valid email", "user@example.com", false},
		{"valid with plus", "user+tag@example.com", false},
		{"missing @", "userexample.com", true},
		{"empty string", "", true},
		{"only @", "@", false}, // Contains @, so technically valid for this simple check
		{"multiple @", "user@@example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmail(%q) error = %v, wantErr %v", tt.email, err, tt.wantErr)
			}
		})
	}
}

func TestWithdrawMoney(t *testing.T) {
	tests := []struct {
		name        string
		balance     float64
		amount      float64
		wantBalance float64
		wantErr     bool
		contains    string
	}{
		{"valid withdrawal", 100, 50, 50, false, ""},
		{"withdraw all", 100, 100, 0, false, ""},
		{"insufficient funds", 100, 150, 100, true, "insufficient"},
		{"negative amount", 100, -50, 100, true, "positive"},
		{"zero amount", 100, 0, 100, true, "positive"},
		{"small withdrawal", 1000.50, 0.50, 1000, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WithdrawMoney(tt.balance, tt.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithdrawMoney(%v, %v) error = %v, wantErr %v", tt.balance, tt.amount, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.wantBalance {
				t.Errorf("WithdrawMoney(%v, %v) = %v, want %v", tt.balance, tt.amount, got, tt.wantBalance)
			}
			if tt.wantErr && !strings.Contains(err.Error(), tt.contains) {
				t.Errorf("WithdrawMoney(%v, %v) error = %q, want to contain %q",
					tt.balance, tt.amount, err.Error(), tt.contains)
			}
		})
	}
}

func TestFormatUserError(t *testing.T) {
	tests := []struct {
		name      string
		operation string
		username  string
		reason    string
		contains  []string
	}{
		{
			"delete user not found",
			"delete",
			"john_doe",
			"user not found",
			[]string{"delete", "john_doe", "user not found"},
		},
		{
			"create duplicate",
			"create",
			"alice",
			"username already exists",
			[]string{"create", "alice", "already exists"},
		},
		{
			"update permission denied",
			"update",
			"bob",
			"permission denied",
			[]string{"update", "bob", "permission denied"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := FormatUserError(tt.operation, tt.username, tt.reason)
			if err == nil {
				t.Errorf("FormatUserError() returned nil, want error")
				return
			}
			errStr := err.Error()
			for _, substr := range tt.contains {
				if !strings.Contains(errStr, substr) {
					t.Errorf("FormatUserError() = %q, want to contain %q", errStr, substr)
				}
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
		contains string
	}{
		{"valid password", "Password1", false, ""},
		{"valid complex", "MyP@ssw0rd", false, ""},
		{"too short", "Pass1", true, "at least 8"},
		{"no digit", "Password", true, "digit"},
		{"no uppercase", "password1", true, "uppercase"},
		{"empty string", "", true, "at least 8"},
		{"only lowercase", "abcdefgh", true, "digit"},
		{"8 chars no digit", "Password", true, "digit"},
		{"8 chars no upper", "password1", true, "uppercase"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePassword(%q) error = %v, wantErr %v", tt.password, err, tt.wantErr)
				return
			}
			if tt.wantErr && !strings.Contains(err.Error(), tt.contains) {
				t.Errorf("ValidatePassword(%q) error = %q, want to contain %q",
					tt.password, err.Error(), tt.contains)
			}
		})
	}
}

func TestParseConfig(t *testing.T) {
	tests := []struct {
		name       string
		filename   string
		lineNumber int
		content    string
		contains   []string
	}{
		{
			"yaml parse error",
			"app.yaml",
			42,
			"invalid: [value",
			[]string{"app.yaml", "42", "invalid: [value"},
		},
		{
			"json error",
			"config.json",
			10,
			"unexpected token",
			[]string{"config.json", "10", "unexpected token"},
		},
		{
			"line 1 error",
			"settings.ini",
			1,
			"missing =",
			[]string{"settings.ini", "1", "missing ="},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ParseConfig(tt.filename, tt.lineNumber, tt.content)
			if err == nil {
				t.Errorf("ParseConfig() returned nil, want error")
				return
			}
			errStr := err.Error()
			for _, substr := range tt.contains {
				if !strings.Contains(errStr, substr) {
					t.Errorf("ParseConfig() = %q, want to contain %q", errStr, substr)
				}
			}
		})
	}
}
