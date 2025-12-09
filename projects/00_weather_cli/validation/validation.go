package validation

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/warlock016/weather_cli/errors"
)

type CLIInput struct {
	Latitude  string
	Longitude string
	StartDate string
	EndDate   string
	Timezone  string
	Variables string
	Format    string
}

type ValidatedInput struct {
	Latitude  float64
	Longitude float64
	StartDate time.Time
	EndDate   time.Time
	Timezone  string
	Variables []string
	Format    string
}

func Validate(input CLIInput) (*ValidatedInput, error) {
	result := ValidatedInput{}
	lat, err := validateLatitude(input.Latitude)
	if err != nil {
		return &result, fmt.Errorf("invalid latitude: %v", err)
	}
	result.Latitude = lat

	lon, err := validateLongitude(input.Longitude)
	if err != nil {
		return &result, fmt.Errorf("invalid longitude: %v", err)
	}
	result.Longitude = lon

	dt, err := validateDate(input.StartDate, "startDate")
	if err != nil {
		return &result, fmt.Errorf("invalid start date %v", err)
	}
	result.StartDate = dt

	df, err := validateDate(input.EndDate, "endDate")
	if err != nil {
		return &result, fmt.Errorf("invalid end date %v", err)
	}
	result.EndDate = df

	if err = validateDateRange(dt, df); err != nil {
		return &result, fmt.Errorf("invalid range %v", err)
	}

	format, err := validateFormat(input.Format)
	if err != nil {
		return &result, fmt.Errorf("invalid format %v", err)
	}
	result.Format = format

	variables, err := validateVariables(input.Variables)
	if err != nil {
		return &result, fmt.Errorf("invalid variables %v", err)
	}
	result.Variables = append(result.Variables, variables...)

	return &result, nil
}

func validateLatitude(lat string) (float64, *errors.FieldError) {
	res, err := strconv.ParseFloat(lat, 64)
	newErr := errors.FieldError{Field: "latitude"}
	if err != nil {
		if lat == "" {
			newErr.Message = "empty string"
		} else {
			newErr.Message = "invalid float conversion"
		}
		return 0, &newErr
	}
	if res > 90 || res < -90 {
		newErr.Message = "out of range"
		return 0, &newErr
	}
	return res, nil
}

func validateLongitude(lon string) (float64, *errors.FieldError) {
	res, err := strconv.ParseFloat(lon, 64)
	newErr := errors.FieldError{Field: "longitude"}

	if err != nil {
		if lon == "" {
			newErr.Message = "empty string"
		} else {
			newErr.Message = "invalid float conversion"
		}
		return 0, &newErr
	}
	if res > 180 || res < -180 {
		newErr.Message = "out of range"
		return 0, &newErr
	}
	return res, nil
}

func validateDate(date, fieldName string) (time.Time, *errors.FieldError) {
	res, err := time.Parse("2006-01-02", date)
	newErr := errors.FieldError{}
	if err != nil {
		newErr.Field = fieldName
		newErr.Message = fmt.Sprintf("date parsing failed: %v", err)
		return time.Time{}, &newErr
	}
	return res, nil
}

func validateDateRange(start, end time.Time) *errors.FieldError {
	if start.After(end) {
		return &errors.FieldError{
			Field:   "start",
			Message: "begins after end",
		}
	}
	return nil
}

func validateFormat(format string) (string, *errors.FieldError) {
	if format == "json" || format == "table" {
		return format, nil
	}

	return "", &errors.FieldError{
		Field:   "format",
		Message: "invalid !(json/table)",
	}
}
func validateVariables(variables string) ([]string, *errors.FieldError) {

	newErr := errors.FieldError{
		Field: "variables",
	}

	if len(variables) == 0 || variables == "" {
		newErr.Message = "invalid variable"
		return nil, &newErr
	}

	fields := strings.Split(variables, ",")

	if len(fields) == 0 {
		newErr.Message = "empty variables"
		return nil, &newErr
	}

	result := make([]string, 0, cap(fields))

	for _, f := range fields {
		clean := strings.Trim(f, " ")
		if len(clean) != 0 {
			result = append(result, clean)
		}
	}

	for _, v := range result {
		if len(v) < 1 {
			newErr.Message = "!variable len > 0"
			return nil, &newErr
		}
	}

	return result, nil
}
