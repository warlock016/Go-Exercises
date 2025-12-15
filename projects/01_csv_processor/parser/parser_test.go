package parser_test

import (
	"testing"
	"time"

	"github.com/warlock016/csv_processor/config"
	"github.com/warlock016/csv_processor/parser"
	"github.com/warlock016/csv_processor/types"
)

func setupTestParserConfig(t *testing.T) *config.ParserConfig {
	t.Helper()
	tz, err := time.LoadLocation("America/Costa_Rica")
	if err != nil {
		t.Fatalf("failed to parse location: %v", err)
	}
	return &config.ParserConfig{
		Name:       "Test Provider",
		Encoding:   "utf-8",
		Delimiter:  ",",
		SkipRows:   0,
		HeaderRows: 2,
		DateConfig: config.DateTimeConfig{
			Detection: "index",
			Index:     0,
			Formats:   []string{"2006-01-02 15:04:05"},
			Timezone:  tz,
		},
		ColConfig: []config.ColumnConfig{
			{
				Name:     "Timestamp",
				Type:     "datetime",
				Aliases:  []string{"Date", "Time"},
				Required: true,
			},
			{
				Name:     "Value",
				Type:     "float",
				Aliases:  []string{"Measurement"},
				Required: true,
			},
		},
		ResourcePath: "test_data/test_file.csv",
	}
}

func setupTestRawData(t *testing.T) *types.RawData {
	t.Helper()
	return &types.RawData{
		Header: [][]string{
			{"Timestamp", "Value"},
			{"", ""},
		},
		HeaderWidth: 2,
		Body: [][]string{
			{"2024-01-01 00:00:00", "12.34"},
			{"2024-01-01 01:00:00", "56.78"},
			{"2024-01-01 02:00:00", "90.12"},
		},
		BodyWidth: 2,
		Source:    "test_data/test_file.csv",
	}
}

func TestParseRawData(t *testing.T) {
	cfg := setupTestParserConfig(t)
	raw := setupTestRawData(t)

	parsedData, procErrors := parser.ParseRawData(cfg, raw)

	if procErrors.HasFatalErrors() {
		for _, e := range procErrors.Errors {
			t.Logf("Processing Error: %s: %s: %s ln: %d col: %d", e.Stage, e.Message, e.Value, e.Line, e.Column)
		}
		t.Fatal("unexpected errors during parsing")
	}
	if procErrors.HasWarnings() {
		for _, e := range procErrors.Warnings {
			t.Logf("%s: %s: %s ln: %d col: %d", e.Stage, e.Message, e.Value, e.Line, e.Column)
		}
		t.Fatal("unexpected errors during parsing")
	}
	if parsedData == nil {
		t.Fatal("parsedData is nil")
	}
	if len(parsedData.Time) != len(raw.Body) {
		t.Fatalf("expected %d time entries, got %d", len(raw.Body), len(parsedData.Time))
	}

	expectedLabels := []string{"Timestamp", "Value"}
	for i, label := range expectedLabels {
		if parsedData.Labels[i] != label {
			t.Fatalf("expected label %s at index %d, got %s", label, i, parsedData.Labels[i])
		}
	}

	valueData, exists := parsedData.Datapoints["Value"]
	if !exists {
		t.Fatal("Value datapoints not found")
	}
	if len(valueData) != len(raw.Body) {
		t.Fatalf("expected %d Value datapoints, got %d", len(raw.Body), len(valueData))
	}
}
