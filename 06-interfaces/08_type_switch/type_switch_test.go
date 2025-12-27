package type_switch

import "testing"

func TestStringify(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"String", "hello", "hello"},
		{"Int", 42, "42"},
		{"Float", 3.14, "3.14"},
		{"Bool true", true, "true"},
		{"Bool false", false, "false"},
		{"Unknown type", []int{1, 2, 3}, "unknown type"},
		{"Nil", nil, "unknown type"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Stringify(tt.input)
			if got != tt.want {
				t.Errorf("Stringify(%v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestTypeName(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"String type", "hello", "string"},
		{"Int type", 42, "int"},
		{"Float type", 3.14, "float64"},
		{"Bool type", true, "bool"},
		{"Slice type", []int{1, 2}, "other"},
		{"Struct type", struct{}{}, "other"},
		{"Nil", nil, "other"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TypeName(tt.input)
			if got != tt.want {
				t.Errorf("TypeName(%T) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name   string
		a      any
		b      any
		want   any
		wantOK bool
	}{
		{"Int add", 1, 2, 3, true},
		{"Float add", 1.5, 2.5, 4.0, true},
		{"String concat", "hello", " world", "hello world", true},
		{"Different types", 1, "hello", nil, false},
		{"Int and float", 1, 2.5, nil, false},
		{"Zero ints", 0, 0, 0, true},
		{"Empty strings", "", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Add(tt.a, tt.b)

			if ok != tt.wantOK {
				t.Errorf("Add(%v, %v) ok = %v, want %v", tt.a, tt.b, ok, tt.wantOK)
			}

			if ok && got != tt.want {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestAddIntegration(t *testing.T) {
	// Test multiple adds
	result, ok := Add(1, 2)
	if !ok || result != 3 {
		t.Errorf("Add(1, 2) = (%v, %v), want (3, true)", result, ok)
	}

	// Chain with another add (if result is int)
	if intResult, ok := result.(int); ok {
		finalResult, ok := Add(intResult, 5)
		if !ok || finalResult != 8 {
			t.Errorf("Add(3, 5) = (%v, %v), want (8, true)", finalResult, ok)
		}
	}
}
