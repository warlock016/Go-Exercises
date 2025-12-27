package parser_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/warlock016/csv_processor/parser"
	"github.com/warlock016/csv_processor/types"
)

func setupTestRawData(t *testing.T) *types.RawData {
	t.Helper()

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Printf("failed to generate location: %v", err)
		return nil
	}

	return &types.RawData{
		Header: [][]string{
			{"Timestamp", "Value"},
			{"", ""},
		},
		HeaderStats: map[int]int{2: 2},
		HeaderWidth: 2,
		Body: [][]string{
			{"2024-01-01 00:00:00", "12.34"},
			{"2024-01-01 01:00:00", "56.78"},
			{"2024-01-01 02:00:00", "90.12"},
		},
		BodyStats: map[int]int{2: 3},
		BodyWidth: 2,
		Source:    "test_data/test_file.csv",

		DatetimeIndex:   0,
		DatetimeFormats: []string{"2006-01-2 15:04:05"},
		Timezone:        loc,
		DigitSeparator:  "",
	}
}

func TestParseRawData(t *testing.T) {
	raw := setupTestRawData(t)
	parsedData, procErrors := parser.ParseRawData(raw)

	if procErrors.HasFatalErrors() {
		for _, e := range procErrors.Errors {
			t.Logf("ERROR: Processing Error: %s: %s: %s ln: %d col: %d", e.Stage, e.Message, e.Value, e.Line, e.Column)
		}
		t.Fatal("unexpected errors during parsing")
	}
	if procErrors.HasWarnings() {
		for _, e := range procErrors.Warnings {
			t.Logf("WARN: %s: %s: %s ln: %d col: %d", e.Stage, e.Message, e.Value, e.Line, e.Column)
		}
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
	// fmt.Printf("TS[0]: %v\nTZ: %s\nHeaders: %v\nDatapoints: %v\nMetadata: %v\n", parsedData.Time[0], parsedData.Timezone, parsedData.Labels, parsedData.Datapoints, parsedData.Metadata)
}
