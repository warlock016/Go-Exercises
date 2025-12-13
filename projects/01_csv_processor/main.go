package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	apiErrors "github.com/warlock016/csv_processor/errors"
)

type Record struct {
	Time      []time.Time
	Variables []string
	Records   map[string][]float64
}

func parseDatetime(datetime string) (*time.Time, error) {
	res, err := time.Parse("02/01/2006 15:04:05", datetime)
	if err != nil {
		return nil, fmt.Errorf("invalid timestamp conversion %v", err)
	}
	return &res, nil
}

func parseValue(value string) (float64, error) {

	var cleanedString string

	if value == "" {
		return 0, fmt.Errorf("invalid input (empty): %w", apiErrors.ErrInvalidInput) //fmt.Errorf("invalid empty input")
	}

	if strings.Contains(value, ",") {
		cleanedString = strings.ReplaceAll(value, ",", "")
	} else {
		cleanedString = value
	}

	fmt.Println(cleanedString)
	res, err := strconv.ParseFloat(cleanedString, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid float string conversion %v", err)
	}
	return res, nil
}

// func main() {
// 	filePath := "./testdata/Weathercloud Pupuseria El Mirador 2025-11.csv"
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatalf("error: %v at path %s", err, filePath)
// 	}
// 	defer file.Close()
// 	// Wrap file in UTF-16 decoder (handles BOM automatically)
// 	decoder := unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()
// 	utf8Reader := transform.NewReader(file, decoder)

// 	csv := csv.NewReader(utf8Reader)
// 	csv.Comma = ';'
// 	csv.Comment = '#'

// 	newErr := apiErrors.ProcessingErrors{
// 		Errors: make([]apiErrors.FieldError, 0),
// 	}

// 	rawRecord := [][]string{}
// 	newRecord := Record{
// 		Time:      make([]time.Time, 0),
// 		Variables: make([]string, 0),
// 		Records:   make(map[string][]float64),
// 	}
// 	i := 0
// 	for {
// 		record, err := csv.Read()
// 		if err == io.EOF {
// 			fmt.Println("EOF reached!")
// 			break
// 		} else if err != nil {
// 			newErr.Add(fmt.Sprintf("%v: %d", err, len(record)), i, 0)
// 		}
// 		rawRecord = append(rawRecord, record)
// 		i++
// 	}

// 	if len(rawRecord) == 0 {
// 		newErr.Add("csv reader returned empty result", 0, 0)
// 	}
// 	for i, v := range rawRecord[0] {
// 		if len([]rune(v)) == 0 {
// 			newErr.Add("invalid empty label", 0, i)
// 		}
// 		newRecord.Variables = append(newRecord.Variables, v)
// 	}

// 	body := rawRecord[1:]
// 	bodyMap := map[int]int{}

// 	for idx, slice := range body {
// 		bodyMap[len(slice)]++

// 		for jdx, field := range slice {
// 			switch jdx {
// 			case 0:
// 				res, err := parseDatetime(field)
// 				if err != nil {
// 					newErr.Add(fmt.Sprintf("invalid timestamp conversion %s", field), idx, jdx)
// 				} else {
// 					newRecord.Time = append(newRecord.Time, *res)
// 				}

// 			default:
// 				res, err := parseValue(field)
// 				if err != nil {
// 					newErr.Add(fmt.Sprintf("invalid value conversion %s, %v at ln: %d, pos: %d", field, err, idx, jdx), idx, jdx)
// 				} else {
// 					newRecord.Records[newRecord.Variables[jdx]] = append(newRecord.Records[newRecord.Variables[jdx]], res)
// 				}

// 			}
// 		}
// 	}

// 	if newErr.HasFatalErrors() || newErr.HasWarnings() {
// 		fmt.Println(newErr.Error())
// 	}

// 	// fmt.Printf("\nBody Map: %v\n\n", bodyMap)
// }
