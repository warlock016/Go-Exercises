package validator

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/warlock016/csv_processor/config"
	apiErrors "github.com/warlock016/csv_processor/errors"
	"github.com/warlock016/csv_processor/types"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func ConvertStringToRune(s string) rune {
	if s == "" {
		return 0
	}
	result := []rune(s)
	return result[0]
}

func ParseStats(stat map[int]int) (int, error) {
	if len(stat) == 0 {
		return 0, fmt.Errorf("map does not contain keys")
	}
	var (
		ref    int
		refKey int
	)
	for k, v := range stat {
		if v > ref {
			ref = v
			refKey = k
		}
	}
	return refKey, nil
}

func ValidateRawFile(cfg *config.ParserConfig) (*types.RawData, apiErrors.ProcessingErrors) {

	result := types.RawData{
		Header:          make([][]string, 0, cfg.HeaderRows),
		HeaderWidth:     0,
		HeaderStats:     make(map[int]int),
		Body:            make([][]string, 0, 4380), // ~730*6 rows in a monthly set with 10 min intervals
		BodyWidth:       0,
		BodyStats:       make(map[int]int),
		Source:          cfg.ResourcePath,
		DatetimeIndex:   cfg.DateConfig.Index,
		DatetimeFormats: cfg.DateConfig.Formats,
		Timezone:        cfg.DateConfig.Timezone,
		DigitSeparator:  cfg.DigitSeparator,
	}

	newErrs := apiErrors.ProcessingErrors{
		Errors:   make([]apiErrors.FieldError, 0, 10),
		Warnings: make([]apiErrors.FieldError, 0, 10),
	}

	file, err := os.Open(cfg.ResourcePath)
	if err != nil {
		newErrs.AddError("file open", "unable to open file", cfg.ResourcePath, 0, 0)
		return nil, newErrs
	}
	defer file.Close()

	var reader io.Reader = file
	switch strings.ToLower(cfg.Encoding) {
	case "utf-8":
		// no special handling needed
	case "utf-16le":
		decoder := unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()
		reader = transform.NewReader(file, decoder)
	case "utf-16be":
		decoder := unicode.UTF16(unicode.BigEndian, unicode.UseBOM).NewDecoder()
		reader = transform.NewReader(file, decoder)
	default:
		newErrs.AddError("file encoding", "unsupported file encoding", cfg.Encoding, 0, 0)
		return nil, newErrs
	}
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1
	// csvReader.
	csvReader.Comma = ConvertStringToRune(cfg.Delimiter)
	if cfg.Comment != "" {
		csvReader.Comment = ConvertStringToRune(cfg.Comment)
	}

	rowIdx := 0

	for {
		// a record defines a row in CSV
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			newErrs.AddWarning("csv parsing", "problematic record", cfg.ResourcePath, rowIdx, 0)
		}

		// if true, then skip the row
		if rowIdx < cfg.SkipRows {
			rowIdx++
			continue
		}

		// ... else continue here
		switch {
		// we are in the header section
		case rowIdx < cfg.HeaderRows+cfg.SkipRows:
			result.Header = append(result.Header, record)
			result.HeaderStats[len(record)]++
		default:
			result.Body = append(result.Body, record)
			result.BodyStats[len(record)]++
		}
		rowIdx++
	}

	if len(result.Header) == 0 {
		newErrs.AddError("header validation", "no header rows found", cfg.ResourcePath, 0, 0)
		return nil, newErrs
	}
	headerStats, err := ParseStats(result.HeaderStats)
	if err != nil {
		newErrs.AddError("header stats validation", "unexpected empty stats map", cfg.ResourcePath, 0, 0)
	}
	result.HeaderWidth = headerStats
	if len(result.HeaderStats) == 0 || result.HeaderWidth == 0 {
		newErrs.AddError("header stats validation", "empty header stats map", cfg.ResourcePath, 0, 0)
		return nil, newErrs
	}
	if len(result.HeaderStats) > 1 {
		newErrs.AddWarning("header structure validation", "inconsistent header lengths", cfg.ResourcePath, 0, 0)
	}

	if len(result.Body) == 0 {
		newErrs.AddError("body validation", "no body rows found", cfg.ResourcePath, 0, 0)
		return nil, newErrs
	}
	bodyStats, err := ParseStats(result.BodyStats)
	if err != nil {
		newErrs.AddError("header stats validation", "unexpected empty stats map", cfg.ResourcePath, 0, 0)
	}
	result.BodyWidth = bodyStats
	if len(result.BodyStats) == 0 || result.BodyWidth == 0 {
		newErrs.AddError("body stats validation", "empty body stats map", cfg.ResourcePath, 0, 0)
		return nil, newErrs
	}
	if len(result.BodyStats) > 1 {
		newErrs.AddWarning("body structure validation", "inconsistent body lengths", cfg.ResourcePath, 0, 0)
	}

	return &result, newErrs
}
