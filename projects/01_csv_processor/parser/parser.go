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
	errors := &apiErrors.ProcessingErrors{
		Errors:   make([]apiErrors.FieldError, 0),
		Warnings: make([]apiErrors.FieldError, 0),
	}

	result, errors = parseRawHeaders(raw, errors)
	if errors.HasFatalErrors() {
		return nil, errors
	}

	// fmt.Printf("%v\n", result.Labels)

	result, errors = parseRawBody(raw, result, errors)
	if errors.HasFatalErrors() {
		return nil, errors
	}

	if len(result.Time) == 0 {
		errors.AddError("ERR: data parsing", "no valid time entries parsed", "", 0, 0)
	}
	if len(result.Labels) == 0 {
		errors.AddError("ERR: data parsing", "no valid data labels parsed", "", 0, 0)
	}
	if len(result.Datapoints) == 0 {
		errors.AddError("ERR: data parsing", "failed to extract time-series", "", 0, 0)
	}

	for label, series := range result.Datapoints {
		if len(series) != len(result.Time) {
			errors.AddError("ERR: data parsing", fmt.Sprintf("invalid \"%s\" series, len: %d, want: %d, delta: %d", label, len(series), len(result.Time), len(series)-len(result.Time)), "", 0, 0)
		}
	}

	if errors.HasFatalErrors() {
		return nil, errors
	}
	return result, errors
}

func parseRawHeaders(raw *types.RawData, errors *apiErrors.ProcessingErrors) (*types.ParsedData, *apiErrors.ProcessingErrors) {

	result := &types.ParsedData{
		Time:          make([]time.Time, 0, len(raw.Body)),
		Timezone:      raw.Timezone,
		Labels:        make([]string, 0, len(raw.Header[0])),
		Datapoints:    make(map[string][]float64),
		Metadata:      make(map[string][]string),
		LabelColIndex: make(map[int]int),
		Source:        raw.Source,
	}

	// *HEADERS* Parse headers row-wise (record labels, ...)
	for row, headerRow := range raw.Header {
		// Malformed Header: we skip header rows that contain only only one column or are completely empty //
		if len(headerRow) < 2 {
			errors.AddWarning("WRN: header parsing", "row missing data", raw.Source, row, 0)
			continue
		}
		// // check if the row length matches the expected
		// if len(headerRow) != raw.HeaderWidth {
		// 	errors.AddWarning("WRN: header parsing", "inconsistent header row length", raw.Source, row, 0)
		// }

		switch row {
		case 0: // this case parses only the first row
			labelIdx := 0
			for col, label := range headerRow {
				if col == raw.DatetimeIndex {
					continue
				}
				if label == "" {
					errors.AddWarning("WRN: header parsing", "empty header label found", raw.Source, row, col)
					continue
				}
				result.LabelColIndex[col] = labelIdx
				labelIdx++
				result.Labels = append(result.Labels, label)
				result.Datapoints[label] = make([]float64, 0, len(raw.Body))
			}
		default: // Additional multi-row header rows can be processed here if needed
		}
	}

	return result, errors
}

func parseRawBody(raw *types.RawData, result *types.ParsedData, errors *apiErrors.ProcessingErrors) (*types.ParsedData, *apiErrors.ProcessingErrors) {

	ParseTimestamp := newTimestampParser(raw.DatetimeFormats, raw.Timezone)

	if raw.BodyWidth != raw.HeaderWidth {
		errors.AddWarning("data parsing", fmt.Sprintf("width mismatch: header: %d body: %d", raw.HeaderWidth, raw.BodyWidth), raw.Source, 0, 0)
	}

rowLoop:
	for row, bodyRow := range raw.Body {
		// row-wise validation
		var gaps int
		switch {
		case len(bodyRow) < 2:
			errors.AddWarning("WRN: body parsing", "row missing data", raw.Source, row, 0)
			continue
		case len(bodyRow) < raw.BodyWidth: // initial ops for detecting datetime parsing format
			// log warning, currently missing value handling strategy!
			gaps = raw.BodyWidth - len(bodyRow)
			errors.AddWarning("data parsing", "skipping gap row", raw.Source, row, 0)
		case len(bodyRow) > raw.BodyWidth: // current row overflow (extra delimiters or )
			right := bodyRow[raw.BodyWidth:]
			validDigits := 0
			for j, k := range right {
				switch k {
				case "":
					errors.AddWarning("WRN: data parsing", "malformed delimiter", raw.Source, row, len(bodyRow)+j)
				default:
					validDigits++
					errors.AddError("ERR: data parsing", "skipping out of bounds value", raw.Source, row, len(bodyRow)+j)
				}
			}
			if validDigits > 0 {
				errors.AddError("ERR: data parsing", "row overflow", raw.Source, row, 0)
			}
			bodyRow = bodyRow[:raw.BodyWidth] // discard excess in any way
		}

		// column-wise validation
		for col, cell := range bodyRow {

			switch col {
			case raw.DatetimeIndex: // we encountered the timestamp column
				ts, err := ParseTimestamp(cell)
				if err != nil {
					errors.AddError("ERR: timestamp processing", "invalid timestamp", raw.Source, row, col)
					continue rowLoop // skips the entire "outer" row iteration after enountering an invalid timestamp
				}
				result.Time = append(result.Time, ts)
				continue
			// case col >= len(result.Labels): // we encountered some cells out of boundaries
			// 	if cell == "" {
			// 		errors.AddWarning("WRN: cell processing", "extra-delimiter", raw.Source, row, col)
			// 	} else {
			// 		errors.AddError("ERR: cell processing", "exceeded expected column width, valid value skipped", raw.Source, row, col)
			// 	}
			// 	continue
			default:
				labelIdx, ok := result.LabelColIndex[col]
				if !ok {
					continue
				}
				label := result.Labels[labelIdx]
				value, err := parseFloat(cell, raw.DigitSeparator)
				if err != nil {
					errors.AddWarning("WRN: data parsing", "parsed NaN value", raw.Source, row, col)
					result.Metadata[label] = append(result.Metadata[label], "NaN") // or use NaN if preferred
					result.Datapoints[label] = append(result.Datapoints[label], math.NaN())
				} else {
					result.Metadata[label] = append(result.Metadata[label], "float64") // or use NaN if preferred
					result.Datapoints[label] = append(result.Datapoints[label], value)
				}
			}
		}

		if len(bodyRow) < raw.HeaderWidth {
			for col := len(bodyRow); col < raw.HeaderWidth; col++ {
				labelIdx, ok := result.LabelColIndex[col]
				if !ok {
					// fmt.Printf("failed to find label %d\n", col)
					continue
				}
				label := result.Labels[labelIdx]
				result.Metadata[label] = append(result.Metadata[label], "NaN") // or use NaN if preferred
				result.Datapoints[label] = append(result.Datapoints[label], math.NaN())
			}
			errors.AddWarning("WRN: data parsing", fmt.Sprintf("filled %d gaps with NaN", gaps), raw.Source, row, 0)
		}
	}

	return result, errors
}

func newTimestampParser(formats []string, timezone *time.Location) func(string) (time.Time, error) {

	layout := ""

	return func(s string) (time.Time, error) {

		switch layout {
		case "":
			for _, format := range formats {
				result, err := time.ParseInLocation(format, s, timezone)
				if err == nil {
					layout = format
					return result, nil
				}
			}
			return time.Time{}, fmt.Errorf("unknown datetime format: %s", s)
		default:
			result, err := time.ParseInLocation(layout, s, timezone)
			if err != nil {
				for _, format := range formats {
					result, err = time.ParseInLocation(format, s, timezone)
					if err == nil {
						layout = format
						return result, nil
					}
				}
				return time.Time{}, fmt.Errorf("failed to parse datetime %s: %v", s, err)
			}
			return result, nil
		}
	}
}

func parseFloat(s, sep string) (float64, error) {

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
