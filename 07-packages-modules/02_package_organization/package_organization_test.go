package organization

import (
	"strings"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name      string
		a         int
		b         int
		operation string
		want      int
		wantErr   bool
	}{
		// Add tests
		{
			name:      "Add positive numbers",
			a:         5,
			b:         3,
			operation: "add",
			want:      8,
			wantErr:   false,
		},
		{
			name:      "Add negative numbers",
			a:         -5,
			b:         -3,
			operation: "add",
			want:      -8,
			wantErr:   false,
		},
		// Subtract tests
		{
			name:      "Subtract",
			a:         10,
			b:         3,
			operation: "subtract",
			want:      7,
			wantErr:   false,
		},
		{
			name:      "Subtract resulting in negative",
			a:         3,
			b:         10,
			operation: "subtract",
			want:      -7,
			wantErr:   false,
		},
		// Multiply tests
		{
			name:      "Multiply",
			a:         4,
			b:         7,
			operation: "multiply",
			want:      28,
			wantErr:   false,
		},
		{
			name:      "Multiply by zero",
			a:         5,
			b:         0,
			operation: "multiply",
			want:      0,
			wantErr:   false,
		},
		// Divide tests
		{
			name:      "Divide evenly",
			a:         10,
			b:         2,
			operation: "divide",
			want:      5,
			wantErr:   false,
		},
		{
			name:      "Divide with remainder (integer division)",
			a:         10,
			b:         3,
			operation: "divide",
			want:      3,
			wantErr:   false,
		},
		{
			name:      "Divide by zero",
			a:         10,
			b:         0,
			operation: "divide",
			want:      0,
			wantErr:   true,
		},
		// Error cases
		{
			name:      "Unknown operation",
			a:         5,
			b:         3,
			operation: "modulo",
			want:      0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.a, tt.b, tt.operation)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Calculate(%d, %d, %q) expected error, got nil", tt.a, tt.b, tt.operation)
				}
			} else {
				if err != nil {
					t.Errorf("Calculate(%d, %d, %q) unexpected error: %v", tt.a, tt.b, tt.operation, err)
				}
				if got != tt.want {
					t.Errorf("Calculate(%d, %d, %q) = %d, want %d", tt.a, tt.b, tt.operation, got, tt.want)
				}
			}
		})
	}
}

func TestDivideByZeroError(t *testing.T) {
	_, err := Calculate(10, 0, "divide")
	if err == nil {
		t.Error("Calculate(10, 0, \"divide\") should return error for division by zero")
	}

	errMsg := strings.ToLower(err.Error())
	if !strings.Contains(errMsg, "zero") && !strings.Contains(errMsg, "divide") {
		t.Errorf("Error message should mention division by zero, got: %q", err.Error())
	}
}

func TestAllOperations(t *testing.T) {
	operations := []struct {
		op   string
		a, b int
		want int
	}{
		{"add", 2, 3, 5},
		{"subtract", 10, 4, 6},
		{"multiply", 6, 7, 42},
		{"divide", 20, 4, 5},
	}

	for _, op := range operations {
		t.Run(op.op, func(t *testing.T) {
			got, err := Calculate(op.a, op.b, op.op)
			if err != nil {
				t.Fatalf("Calculate(%d, %d, %q) unexpected error: %v", op.a, op.b, op.op, err)
			}
			if got != op.want {
				t.Errorf("Calculate(%d, %d, %q) = %d, want %d", op.a, op.b, op.op, got, op.want)
			}
		})
	}
}
