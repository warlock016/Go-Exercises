package parser

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	apiErrors "github.com/warlock016/csv_processor/errors"
	"github.com/warlock016/csv_processor/types"
)

func ParseRawData(raw *types.RawData) (*types.ParsedData, *apiErrors.ProcessingErrors) {

	var result *types.ParsedData
	var errors *apiErrors.ProcessingErrors

	result, errors = parseRawHeaders(raw)
	if errors.HasFatalErrors() {
		return nil, errors
	}

	result, errors = parseRawBody(raw, result)
	if errors.HasFatalErrors() {
		return nil, errors
	}

	if len(result.Time) == 0 {
		errors.AddError("data parsing", "no valid time entries parsed", "", 0, 0)
	}
	if len(result.Labels) == 0 {
		errors.AddError("data parsing", "no valid data labels parsed", "", 0, 0)
	}
	if len(result.Datapoints) == 0 {
		errors.AddError("data parsing", "failed to extract time-series", "", 0, 0)
	}

	for label, series := range result.Datapoints {
		if len(series) != len(result.Time) {
			errors.AddError("data parsing", fmt.Sprintf("invalid \"%s\" series, len: %d, want: %d, delta: %d", label, len(series), len(result.Time), len(series)-len(result.Time)), "", 0, 0)
		}
	}

	if errors.HasFatalErrors() {
		return nil, errors
	}
	return result, errors
}

func parseRawHeaders(raw *types.RawData) (*types.ParsedData, *apiErrors.ProcessingErrors) {

	errors := apiErrors.ProcessingErrors{
		Errors:   make([]apiErrors.FieldError, 0, 10),
		Warnings: make([]apiErrors.FieldError, 0, 10),
	}

	result := &types.ParsedData{
		Time:        make([]time.Time, 0, len(raw.Body)),
		Timezone:    raw.Timezone,
		Labels:      make([]string, 0, len(raw.Header[0])),
		Datapoints:  make(map[string][]float64),
		Metadata:    make(map[string][]string),
		SkippedCols: make(map[int]bool),
	}

	// *HEADERS* Parse headers row-wise (record labels, ...)
	for row, headerRow := range raw.Header {
		// Malformed Header: we skip header rows that contain only only one column or are completely empty //
		if len(headerRow) < 2 {
			errors.AddWarning("header parsing", "row missing data", strings.Join(headerRow, ";"), row, 0)
			continue
		}
		// check if the row length matches the expected
		if len(headerRow) != raw.HeaderWidth {
			errors.AddWarning("header parsing", "inconsistent header row length", strings.Join(headerRow, ";"), row, 0)
		}

		switch row {
		case 0: // this case parses only the first row
			for col, label := range headerRow {
				switch label {
				case "":
					// mark as skipped and warn if key is empty
					// skippedCols[col] = true
					result.SkippedCols[col] = true
					errors.AddWarning("header parsing", "empty header label found", raw.Source, row, row)
					continue
				default:
					result.Labels = append(result.Labels, label)
					// initialize the float slice for the given key, so that we can later push data
					if col != raw.DatetimeIndex {
						result.Datapoints[label] = make([]float64, 0, len(raw.Body))
					}
				}
			}
		default: // Additional multi-row header rows can be processed here if needed
		}
	}

	return result, &errors
}

func parseRawBody(raw *types.RawData, result *types.ParsedData) (*types.ParsedData, *apiErrors.ProcessingErrors) {

	errors := apiErrors.ProcessingErrors{
		Errors:   make([]apiErrors.FieldError, 0, 10),
		Warnings: make([]apiErrors.FieldError, 0, 10),
	}

	ParseTimestamp := NewTimestampParser(raw.DatetimeFormats, raw.Timezone)

	for row, bodyRow := range raw.Body {
		// row-wise validation
		var gaps int
		switch {
		case len(bodyRow) < 2:
			errors.AddWarning("body parsing", "row missing data", strings.Join(bodyRow, ";"), row, 0)
			continue
		case len(bodyRow) < raw.BodyWidth: // initial ops for detecting datetime parsing format
			// log warning, currently missing value handling strategy!
			gaps = raw.BodyWidth - len(bodyRow)
			errors.AddWarning("data parsing", "skipping gap row", bodyRow[0], row, 0)
			// continue
		case len(bodyRow) > raw.BodyWidth: // current row overflow (extra delimiters or )
			right := bodyRow[raw.BodyWidth:]
			validDigits := 0
			for j, k := range right {
				switch k {
				case "":
					errors.AddWarning("data parsing", "malformed delimiter", "", row, len(bodyRow)+j)
				default:
					validDigits++
					errors.AddError("data parsing", "skipping out of bounds value", k, row, len(bodyRow)+j)
				}
			}
			if validDigits > 0 {
				errors.AddError("data parsing", "row overflow", fmt.Sprintf("discarded %d values", validDigits), row, 0)
			}
			bodyRow = bodyRow[:raw.BodyWidth] // discard
		}

		// column-wise validation
		for col, cell := range bodyRow {
			switch {
			case result.SkippedCols[col]: // we marked this column as skipped due to malformed header
				continue
			case col == raw.DatetimeIndex: // we encountered the timestamp column
				ts, err := ParseTimestamp(cell)
				if err != nil {
					errors.AddError("timestamp processing", "invalid timestamp", cell, row, col)
				}
				result.Time = append(result.Time, ts)
				continue
			case col >= len(result.Labels): // we encountered some cells out of boundaries
				if cell == "" {
					errors.AddWarning("cell processing", "extra-delimiter", raw.Source, row, col)
				} else {
					errors.AddError("cell processing", "exceeded expected column width, valid value skipped", raw.Source, row, col)
				}
				continue
			default:
				label := result.Labels[col]
				value, err := ParseFloat(cell, raw.DigitSeparator)
				if err != nil {
					errors.AddWarning("data parsing", "parsed NaN value", cell, row, col)
					result.Metadata[label] = append(result.Metadata[label], "NaN") // or use NaN if preferred
					result.Datapoints[label] = append(result.Datapoints[label], math.NaN())
				} else {
					result.Metadata[label] = append(result.Metadata[label], "float64") // or use NaN if preferred
					result.Datapoints[label] = append(result.Datapoints[label], value)

					if col == len(bodyRow)-1 && len(bodyRow) < raw.BodyWidth {
						label := result.Labels[col]
						for range gaps {
							result.Metadata[label] = append(result.Metadata[label], "NaN") // or use NaN if preferred
							result.Datapoints[label] = append(result.Datapoints[label], math.NaN())
						}
						errors.AddWarning("data parsing", fmt.Sprintf("filled %d gaps with NaN for \"%s\"", gaps, label), raw.Source, row, 0)
					}
				}
			}
		}
	}
	return result, &errors
}

func ParseFloat(s, sep string) (float64, error) {

	if s == "" {
		return 0, fmt.Errorf("empty string")
	}
	if sep != "" {
		s = strings.ReplaceAll(s, sep, "")
	}

	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, fmt.Errorf("parse float error: %w", err)
	}
	return result, nil
}

func NewTimestampParser(formats []string, timezone *time.Location) func(string) (time.Time, error) {

	layout := ""
	// cfg.DateConfig.Formats

	return func(s string) (time.Time, error) {

		result := time.Time{}

		switch layout {
		case "":
			for cnt, format := range formats {
				result, err := time.ParseInLocation(format, s, timezone)
				if err == nil {
					layout = format
					return result, nil
				}
				if cnt == len(formats)-1 {
					return time.Time{}, fmt.Errorf("unknown datetime format: %s %v", s, err)
				}
			}
		default:
			result, err := time.ParseInLocation(layout, s, timezone)
			if err != nil {
				return time.Time{}, fmt.Errorf("failed to parse datetime %s: %v", s, err)
			}
			return result, nil
		}
		return result, fmt.Errorf("unknown error")
	}
}
