package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/warlock016/csv_processor/config"
	apiErrors "github.com/warlock016/csv_processor/errors"
	"github.com/warlock016/csv_processor/types"
)

func ParseFloat(s string) (float64, error) {
	if s == "" {
		return 0, fmt.Errorf("empty string")
	}
	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, fmt.Errorf("parse float error: %w", err)
	}
	return result, nil
}

func ParseRawData(cfg *config.ParserConfig, raw *types.RawData) (*types.ParsedData, *apiErrors.ProcessingErrors) {

	errors := apiErrors.ProcessingErrors{
		Errors:   make([]apiErrors.FieldError, 0, 10),
		Warnings: make([]apiErrors.FieldError, 0, 10),
	}
	if cfg == nil {
		errors.AddError("raw data", "config is nil", "", 0, 0)
	}
	if raw == nil {
		errors.AddError("raw data", "raw data is nil", "", 0, 0)
		return nil, &errors
	}
	if len(raw.Header) == 0 {
		errors.AddError("raw data", "unexpected empty header", "", 0, 0)
		return nil, &errors
	}
	if len(raw.Body) == 0 {
		errors.AddError("raw data", "unexpected empty body", "", 0, 0)
		return nil, &errors
	}

	result := &types.ParsedData{
		Time:       make([]time.Time, 0, len(raw.Body)),
		Timezone:   cfg.DateConfig.Timezone,
		Labels:     make([]string, 0, len(raw.Header[0])),
		Datapoints: make(map[string][]float64),
		Metadata:   make(map[string][]string),
	}

	// Parse headers row-wise
	for i, headerRow := range raw.Header {
		if len(headerRow) != raw.HeaderWidth {
			errors.AddWarning("header parsing", "inconsistent header row length", strings.Join(headerRow, ";"), i, 0)
		}
		switch i {
		case 0:
			for j, key := range headerRow {
				if key == "" {
					errors.AddWarning("header parsing", "empty header label found", raw.Source, i+1, j+1)
				}

				// append all so that we can reference by index later
				result.Labels = append(result.Labels, key)
				result.Datapoints[key] = make([]float64, 0, len(raw.Body))
				result.Metadata[key] = make([]string, 0, len(raw.Body))
			}
		default:

			// Additional header rows can be processed here if needed
		}
	}

	// Parse body row-wise
	for rowIdx, row := range raw.Body {
		// Parse date/time from the specified column

		if len(row) != raw.BodyWidth {
			errors.AddError("body parsing", "inconsistent body row length", strings.Join(row, ";"), rowIdx, 0)
		}

		dateStr := row[cfg.DateConfig.Index]

		for options, format := range cfg.DateConfig.Formats {
			parsedTime, err := time.ParseInLocation(format, dateStr, cfg.DateConfig.Timezone)
			if err == nil {
				result.Time = append(result.Time, parsedTime)
				break
			}

			if options == len(cfg.DateConfig.Formats)-1 {
				errors.AddError("datetime parsing", fmt.Sprintf("unable to parse datetime: %v", err), raw.Source, rowIdx+1, cfg.DateConfig.Index+1)
				result.Time = append(result.Time, time.Time{})
			}
		}

		for colIdx, cell := range row {
			if colIdx == cfg.DateConfig.Index {
				continue // skip date column
			}
			label := result.Labels[colIdx]
			if label == "" {
				continue // skip columns without labels
			}
			value, err := ParseFloat(cell)
			if err != nil {
				errors.AddWarning("data parsing", fmt.Sprintf("parsed expected float as metadata: %v", err), raw.Source, rowIdx+1, colIdx+1)
				result.Metadata[label] = append(result.Metadata[label], cell) // or use NaN if preferred
			} else {
				result.Datapoints[label] = append(result.Datapoints[label], value)
			}
		}
	}

	if len(result.Time) == 0 {
		errors.AddError("data parsing", "no valid time entries parsed", "", 0, 0)
	}
	if len(result.Labels) == 0 {
		errors.AddError("data parsing", "no valid data labels parsed", "", 0, 0)
	}

	return result, &errors
}
