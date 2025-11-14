package expressionevaluator

import "testing"

func TestSimpleAddition(t *testing.T) {
	tests := []struct {
		name  string
		expr  string
		want  int
		isErr bool
	}{
		{"single number", "5", 5, false},
		{"simple add", "2 + 3", 5, false},
		{"simple subtract", "10 - 3", 7, false},
		{"multiple adds", "1 + 2 + 3", 6, false},
		{"multiple subtracts", "10 - 2 - 3", 5, false},
		{"mixed add/subtract", "10 + 5 - 3", 12, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.expr)
			if (err != nil) != tt.isErr {
				t.Errorf("Evaluate(%q) error = %v, wantErr %v", tt.expr, err, tt.isErr)
				return
			}
			if got != tt.want {
				t.Errorf("Evaluate(%q) = %v, want %v", tt.expr, got, tt.want)
			}
		})
	}
}

func TestMultiplicationPrecedence(t *testing.T) {
	tests := []struct {
		name  string
		expr  string
		want  int
		isErr bool
	}{
		{"simple multiply", "2 * 3", 6, false},
		{"simple divide", "10 / 2", 5, false},
		{"multiply before add", "2 + 3 * 4", 14, false},
		{"add then multiply", "3 * 4 + 2", 14, false},
		{"multiply in middle", "1 + 2 * 3 + 4", 11, false},
		{"division before subtract", "10 - 6 / 2", 7, false},
		{"multiple operations", "2 * 3 + 4 * 5", 26, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.expr)
			if (err != nil) != tt.isErr {
				t.Errorf("Evaluate(%q) error = %v, wantErr %v", tt.expr, err, tt.isErr)
				return
			}
			if got != tt.want {
				t.Errorf("Evaluate(%q) = %v, want %v", tt.expr, got, tt.want)
			}
		})
	}
}

func TestParentheses(t *testing.T) {
	tests := []struct {
		name  string
		expr  string
		want  int
		isErr bool
	}{
		{"parens override precedence", "(2 + 3) * 4", 20, false},
		{"nested parens", "((1 + 2) * 3)", 9, false},
		{"multiple paren groups", "(1 + 2) * (3 + 4)", 21, false},
		{"parens with subtraction", "(10 - 2) * 3", 24, false},
		{"complex nested", "2 * (3 + (4 * 5))", 46, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.expr)
			if (err != nil) != tt.isErr {
				t.Errorf("Evaluate(%q) error = %v, wantErr %v", tt.expr, err, tt.isErr)
				return
			}
			if got != tt.want {
				t.Errorf("Evaluate(%q) = %v, want %v", tt.expr, got, tt.want)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	tests := []struct {
		name  string
		expr  string
		want  int
		isErr bool
	}{
		{"simple division", "10 / 2", 5, false},
		{"integer division", "7 / 2", 3, false},
		{"division precedence", "10 / 2 + 3", 8, false},
		{"multiple divisions", "20 / 2 / 2", 5, false},
		{"division by zero", "10 / 0", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.expr)
			if (err != nil) != tt.isErr {
				t.Errorf("Evaluate(%q) error = %v, wantErr %v", tt.expr, err, tt.isErr)
				return
			}
			if !tt.isErr && got != tt.want {
				t.Errorf("Evaluate(%q) = %v, want %v", tt.expr, got, tt.want)
			}
		})
	}
}

func TestComplexExpressions(t *testing.T) {
	tests := []struct {
		name  string
		expr  string
		want  int
		isErr bool
	}{
		{"complex 1", "2 + 3 * 4 - 5", 9, false},
		{"complex 2", "10 / 2 + 3 * 4 - 1", 16, false},
		{"complex 3", "(10 + 5) / 3", 5, false},
		{"complex 4", "2 * (3 + 4) * 5", 70, false},
		{"with spaces", "  2  +  3  *  4  ", 14, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.expr)
			if (err != nil) != tt.isErr {
				t.Errorf("Evaluate(%q) error = %v, wantErr %v", tt.expr, err, tt.isErr)
				return
			}
			if got != tt.want {
				t.Errorf("Evaluate(%q) = %v, want %v", tt.expr, got, tt.want)
			}
		})
	}
}

func TestErrorCases(t *testing.T) {
	tests := []struct {
		name string
		expr string
	}{
		{"empty expression", ""},
		{"missing operand", "2 +"},
		{"missing operator", "2 3"},
		{"unmatched left paren", "(2 + 3"},
		{"unmatched right paren", "2 + 3)"},
		{"invalid character", "2 & 3"},
		{"double operator", "2 ++ 3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Evaluate(tt.expr)
			if err == nil {
				t.Errorf("Evaluate(%q) expected error, got nil", tt.expr)
			}
		})
	}
}

func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		expr  string
		want  int
		isErr bool
	}{
		{"zero", "0", 0, false},
		{"negative result", "5 - 10", -5, false},
		{"multiple zeros", "0 + 0 + 0", 0, false},
		{"multiply by zero", "5 * 0", 0, false},
		{"large numbers", "100 + 200", 300, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.expr)
			if (err != nil) != tt.isErr {
				t.Errorf("Evaluate(%q) error = %v, wantErr %v", tt.expr, err, tt.isErr)
				return
			}
			if got != tt.want {
				t.Errorf("Evaluate(%q) = %v, want %v", tt.expr, got, tt.want)
			}
		})
	}
}
