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

func convertStringToRune(s string) rune {
	if s == "" {
		return 0
	}
	result := []rune(s)
	return result[0]
}

func ValidateRawFile(cfg *config.ParserConfig) (*config.ParserConfig, *types.RawData, apiErrors.ProcessingErrors) {

	result := types.RawData{
		Header: make([][]string, 0, cfg.HeaderRows),
		Body:   make([][]string, 0, 730), // ~730 rows in a month of hourly data
		Source: cfg.ResourcePath,
	}

	newErrs := apiErrors.ProcessingErrors{
		Errors:   make([]apiErrors.FieldError, 0, 10),
		Warnings: make([]apiErrors.FieldError, 0, 10),
	}

	file, err := os.Open(cfg.ResourcePath)
	if err != nil {
		newErrs.Add("file open", "unable to open file", cfg.ResourcePath, 0, 0)
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
		newErrs.Add("file encoding", "unsupported file encoding", cfg.Encoding, 0, 0)
		return nil, newErrs
	}
	csvReader := csv.NewReader(reader)
	csvReader.Comma = convertStringToRune(cfg.Delimiter)
	if cfg.Comment != "" {
		csvReader.Comment = convertStringToRune(cfg.Comment)
	}

	rowIdx := 0
	for {
		// a record defines a row in CSV
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			newErrs.AddWarning("csv read", fmt.Sprintf("error reading CSV record: %v", err), cfg.ResourcePath, rowIdx, 0)
		}

		if rowIdx < cfg.SkipRows { // if true, then skip the row
			rowIdx++
			continue
		}

		// this section naively checks for consistent row widths,
		// assuming that the first row of a type defines the width
		switch { // ... else continue here
		case rowIdx < cfg.HeaderRows+cfg.SkipRows:
			if rowIdx-cfg.SkipRows == 0 { //first header row, set width
				result.HeaderWidth = len(record)
			} else if len(record) != result.HeaderWidth {
				newErrs.AddWarning("header validation", "inconsistent header width", cfg.ResourcePath, rowIdx, 0)
			}
			result.Header = append(result.Header, record)
		default:
			if rowIdx-cfg.HeaderRows-cfg.SkipRows == 0 { // first body row, set width
				result.BodyWidth = len(record)
			} else if len(record) != result.BodyWidth {
				newErrs.AddWarning("body validation", "inconsistent body width", cfg.ResourcePath, rowIdx, 0)
			}
			result.Body = append(result.Body, record)
		}
		rowIdx++
	}

	if len(result.Header) == 0 {
		newErrs.Add("header validation", "no header rows found", cfg.ResourcePath, 0, 0)
	}
	if len(result.Body) == 0 {
		newErrs.Add("body validation", "no body rows found", cfg.ResourcePath, 0, 0)
	}

	return cfg, &result, newErrs
}
