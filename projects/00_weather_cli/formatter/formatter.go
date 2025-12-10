package formatter

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/warlock016/weather_cli/errors"
	"github.com/warlock016/weather_cli/types"
)

type FormatOptions struct {
	Format   string
	Timezone string
}

type FormatterInput struct {
	Weather  *types.WeatherData
	Location *types.GeoData
	Options  FormatOptions
}

type FormattedOutput struct {
	Content    string
	DataPoints int
}

func Format(input FormatterInput) (*FormattedOutput, error) {
	if input.Weather == nil || input.Location == nil {
		return nil, errors.ErrInvalidInput
	}

	var result string
	var err error

	switch input.Options.Format {
	case "table":
		result, err = FormatTable(input)
		if err != nil {
			return nil, errors.ErrInvalidInput
		}
	case "json":
		result, err = FormatJSON(input)
		if err != nil {
			return nil, errors.ErrInvalidInput
		}
	default:
		return nil, errors.ErrInvalidInput
	}

	return &FormattedOutput{
		Content:    result,
		DataPoints: len(input.Weather.Hourly["time"]),
	}, nil
}

func Max(inputs []any) float64 {
	if len(inputs) == 0 {
		return 0
	}

	floats := make([]float64, 0, len(inputs))
	for _, v := range inputs {
		if res, ok := v.(float64); ok {
			floats = append(floats, res)
		}
	}

	switch len(floats) {
	case 0:
		return 0
	case 1:
		return floats[0]
	default:
		return slices.Max(floats)
	}
}

func Min(inputs []any) float64 {
	if len(inputs) == 0 {
		return 0
	}

	floats := make([]float64, 0, len(inputs))
	for _, val := range inputs {
		if res, ok := val.(float64); ok {
			floats = append(floats, res)
		}
	}
	switch len(floats) {
	case 0:
		return 0
	case 1:
		return floats[0]
	default:
		return slices.Min(floats)
	}
}

func Avg(inputs []any) float64 {
	if len(inputs) == 0 {
		return 0
	}

	floats := make([]float64, 0, len(inputs))
	for _, val := range inputs {
		if res, ok := val.(float64); ok {
			floats = append(floats, res)
		}
	}

	switch len(floats) {
	case 0:
		return 0
	case 1:
		return floats[0]
	default:
		sum := float64(0)
		count := float64(len(floats))
		for _, v := range floats {
			sum += v
		}
		return sum / count
	}
}

func Median(inputs []any) float64 {
	if len(inputs) == 0 {
		return 0
	}
	floats := make([]float64, 0, len(inputs))
	for _, v := range inputs {
		if res, ok := v.(float64); ok {
			floats = append(floats, res)
		}
	}

	switch len(floats) {
	case 0:
		return 0
	case 1:
		return floats[0]
	default:
		slices.Sort(floats)
		return floats[len(floats)/2]
	}
}

func FormatTable(input FormatterInput) (string, error) {
	var output strings.Builder

	// Summary
	times := input.Weather.Hourly["time"]
	output.WriteString(fmt.Sprintf("\nWeather Report: %s\n", input.Location.DisplayName))
	output.WriteString(fmt.Sprintf("Lat, Lon: %.2f, %.2f (%s) UTC Offset (s): %d\n", input.Weather.Latitude, input.Weather.Longitude, input.Weather.Timezone, input.Weather.Offset))

	start := input.Weather.Hourly["time"][0].(string)
	end := input.Weather.Hourly["time"][len(input.Weather.Hourly["time"])-1].(string)

	output.WriteString(fmt.Sprintf("Period: %v to %v\n\n", strings.ReplaceAll(start, "T", " "), strings.ReplaceAll(end, "T", " ")))

	// Table headers (Time, ...Variable names)
	keyCharCount := make(map[string]int)
	// keys := []string{}                                         // ordered list of keys
	output.WriteString("Time" + strings.Repeat(" ", 13) + "|") // 18 character space == 4 + 13 + 1
	keyCharCount["time"] = 4 + 13 + 1                          // > timestamp length (15 chars for YYYY-MM-DD HH:MM string)

	// for key := range input.Weather.Hourly {
	// 	if key != "time" {
	// 		keys = append(keys, key)
	// 	}
	// }
	// slices.Sort(keys)

	for _, v := range input.Weather.Variables {
		if len([]rune(v)) < 10 {
			output.WriteString(fmt.Sprintf(" %8v |", v)) // len substring == 10
			keyCharCount[v] = 11                         // sets key length to len(key) + 2 white spaces and one pipe == " <key> |"
		} else {
			output.WriteString(fmt.Sprintf(" %s |", v)) // len substring == keyCharCount[key] + 3
			keyCharCount[v] = len(v) + 3                // sets key length to len(key) + 2 white spaces and one pipe == " <key> |"
		}

	}
	output.WriteString("\n")

	// "*=" Separator
	repeat := 0
	for _, v := range keyCharCount {
		repeat += v
	}
	output.WriteString(strings.Repeat("=", repeat) + "\n")

	// Table values
	if len(times) > 12 { // if table holds over 12 rows, then compact
		capA := 3
		capB := len(times) - 3
		truncated := false

		for i, val := range times {
			if i < capA || i >= capB {
				ts := strings.ReplaceAll(val.(string), "T", " ")
				output.WriteString(fmt.Sprintf("%-16s ", ts))
				output.WriteString("|")

				for _, key := range input.Weather.Variables {
					var comp strings.Builder
					var result string
					comp.WriteString(fmt.Sprintf("%v %s |", input.Weather.Hourly[key][i], input.Weather.HourlyUnits[key]))
					result = comp.String()
					output.WriteString(strings.Repeat(" ", keyCharCount[key]-len([]rune(result))))
					output.WriteString(result)
				}
				output.WriteString("\n")
			} else if !truncated {
				truncated = true
				output.WriteString(fmt.Sprintf("%-17s", "..."))
				output.WriteString("|")

				for _, key := range input.Weather.Variables {
					var comp strings.Builder
					var result string
					comp.WriteString(fmt.Sprintf("%v |", "..."))
					result = comp.String()
					output.WriteString(strings.Repeat(" ", keyCharCount[key]-len([]rune(result))))
					output.WriteString(result)
				}
				output.WriteString("\n")
			}
		}
	} else {
		for i, val := range times {
			ts := strings.ReplaceAll(val.(string), "T", " ")
			output.WriteString(fmt.Sprintf("%-16s ", ts))
			output.WriteString("|")

			for _, key := range input.Weather.Variables {
				var comp strings.Builder
				var result string
				comp.WriteString(fmt.Sprintf("%v %s |", input.Weather.Hourly[key][i], input.Weather.HourlyUnits[key]))
				result = comp.String()
				output.WriteString(strings.Repeat(" ", keyCharCount[key]-len([]rune(result))))
				output.WriteString(result)
			}
			output.WriteString("\n")
		}
	}
	output.WriteString(strings.Repeat("=", repeat) + "\n")

	stats := []string{"min", "max", "avg", "median"}

	for _, v := range stats {
		output.WriteString(fmt.Sprintf("%16v |", v))

		var result float64
		for j, varible := range input.Weather.Variables {
			switch v {
			case "min":
				result = Min(input.Weather.Hourly[varible])
			case "max":
				result = Max(input.Weather.Hourly[varible])
			case "avg":
				result = Avg(input.Weather.Hourly[varible])
			case "median":
				result = Median(input.Weather.Hourly[varible])
			default:
			}

			output.WriteString(strings.Repeat(" ", keyCharCount[varible]-len([]rune(fmt.Sprintf("%.2f %s |", result, input.Weather.HourlyUnits[varible])))))
			output.WriteString(fmt.Sprintf("%.2f %s |", result, input.Weather.HourlyUnits[varible]))

			if j == len(input.Weather.Variables)-1 {
				output.WriteString("\n")
			}
		}
	}

	output.WriteString(strings.Repeat("=", repeat) + "\n")
	output.WriteString(fmt.Sprintf("Summary: %d hourly data points\n", len(times)))

	if output.String() == "" {
		return "", fmt.Errorf("invalid string output: got %s, want some string", output.String())
	}

	return output.String(), nil
}

func FormatJSON(input FormatterInput) (string, error) {

	bytes, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return "", fmt.Errorf("invalid struct marshalling %+v, got err", input)
	}
	return string(bytes), nil
}
