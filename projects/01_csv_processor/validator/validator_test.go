package validator_test

import (
	"testing"

	"github.com/warlock016/csv_processor/config"
	"github.com/warlock016/csv_processor/validator"
)

func CreateTestConfig(t *testing.T) *config.ParserConfig {
	t.Helper()
	return &config.ParserConfig{
		ResourcePath: "../testdata/sample.csv",
		Delimiter:    ";",
		Encoding:     "utf-16le",
		HeaderRows:   1,
		DateConfig: config.DateTimeConfig{
			Detection: "index",
			Index:     0,
			Formats:   []string{"2006-01-02 15:04"},
		},
		ColConfig: []config.ColumnConfig{
			{
				Name:     "timestamp",
				Type:     "datetime",
				Aliases:  []string{"timesstamp"},
				Required: true,
			},
			{
				Name:     "value",
				Type:     "float",
				Aliases:  []string{"Test Column"},
				Required: true,
			},
		},
	}
}

func TestConvertStringToRune(t *testing.T) {
	// t.Skip("not implemented yet")
	tests := []struct {
		input    string
		expected rune
	}{
		{";", ';'},
		{",", ','},
		{"", 0},
		{"abc", 'a'},
	}

	for _, test := range tests {
		result := validator.ConvertStringToRune(test.input)
		if result != test.expected {
			t.Errorf("convertStringToRune(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestValidateRawFile(t *testing.T) {
	cfg := CreateTestConfig(t)

	rawData, errs := validator.ValidateRawFile(cfg)
	if len(errs.Errors) > 0 {
		t.Errorf("Expected no errors, but got: %v", errs.Errors)
	}
	if rawData == nil {
		t.Fatal("Expected rawData to be non-nil")
	}

	if len(rawData.Header) != cfg.HeaderRows {
		t.Errorf("Expected %d header rows, but got: %d", cfg.HeaderRows, len(rawData.Header))
	}

	if len(rawData.Body) == 0 {
		t.Error("Expected body to have 3 data rows, got: 0")
	}

	if len(rawData.Body) != 3 {
		t.Errorf("Expected 3 data rows, but got: %d", len(rawData.Body))
	}
}
