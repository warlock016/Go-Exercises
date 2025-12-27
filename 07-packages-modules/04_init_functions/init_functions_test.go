package initfuncs

import (
	"strings"
	"testing"
)

func TestIsInitialized(t *testing.T) {
	// init() should have run automatically
	if !IsInitialized() {
		t.Error("Package should be initialized automatically via init()")
	}
}

func TestDefaultHandlers(t *testing.T) {
	// Test that default handlers were registered in init()
	tests := []struct {
		name        string
		handlerName string
		input       string
		wantFunc    func(string) string
	}{
		{
			name:        "Uppercase handler exists",
			handlerName: "uppercase",
			input:       "hello",
			wantFunc:    strings.ToUpper,
		},
		{
			name:        "Lowercase handler exists",
			handlerName: "lowercase",
			input:       "HELLO",
			wantFunc:    strings.ToLower,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler, ok := GetHandler(tt.handlerName)
			if !ok {
				t.Errorf("GetHandler(%q) handler not found - should be registered in init()", tt.handlerName)
				return
			}

			got := handler(tt.input)
			want := tt.wantFunc(tt.input)

			if got != want {
				t.Errorf("Handler %q(%q) = %q, want %q", tt.handlerName, tt.input, got, want)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	// Test registering a new handler
	called := false
	testHandler := func(s string) string {
		called = true
		return "test-" + s
	}

	Register("test", testHandler)

	handler, ok := GetHandler("test")
	if !ok {
		t.Error("GetHandler(\"test\") handler not found after Register()")
		return
	}

	result := handler("input")
	if !called {
		t.Error("Registered handler was not called")
	}
	if result != "test-input" {
		t.Errorf("Handler returned %q, want %q", result, "test-input")
	}
}

func TestGetHandlerNotFound(t *testing.T) {
	_, ok := GetHandler("nonexistent")
	if ok {
		t.Error("GetHandler(\"nonexistent\") should return false for non-existent handler")
	}
}

func TestRegisterOverwrite(t *testing.T) {
	// Register a handler
	Register("overwrite", func(s string) string {
		return "first"
	})

	// Overwrite it
	Register("overwrite", func(s string) string {
		return "second"
	})

	handler, ok := GetHandler("overwrite")
	if !ok {
		t.Fatal("Handler should exist")
	}

	result := handler("test")
	if result != "second" {
		t.Errorf("Handler should be overwritten, got %q, want %q", result, "second")
	}
}

func TestMultipleHandlers(t *testing.T) {
	// Register multiple custom handlers
	handlers := map[string]Handler{
		"reverse": func(s string) string {
			runes := []rune(s)
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			return string(runes)
		},
		"double": func(s string) string {
			return s + s
		},
		"prefix": func(s string) string {
			return ">> " + s
		},
	}

	// Register all
	for name, handler := range handlers {
		Register(name, handler)
	}

	// Verify all are accessible
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"reverse", "hello", "olleh"},
		{"double", "ab", "abab"},
		{"prefix", "test", ">> test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler, ok := GetHandler(tt.name)
			if !ok {
				t.Errorf("GetHandler(%q) not found", tt.name)
				return
			}

			got := handler(tt.input)
			if got != tt.want {
				t.Errorf("Handler %q(%q) = %q, want %q", tt.name, tt.input, got, tt.want)
			}
		})
	}
}

func TestInitRunsOnce(t *testing.T) {
	// This is a conceptual test
	// init() runs exactly once per package, before any other code
	// We can't test this directly, but we verify the effects

	t.Log("init() runs automatically once before any tests")
	t.Log("This is why IsInitialized() returns true")
	t.Log("And why default handlers are already registered")

	if !IsInitialized() {
		t.Error("If this fails, init() didn't run properly")
	}

	// Verify we have at least the default handlers
	_, hasUppercase := GetHandler("uppercase")
	_, hasLowercase := GetHandler("lowercase")

	if !hasUppercase || !hasLowercase {
		t.Error("Default handlers should be registered in init()")
	}
}
