package formatter

import (
	// "go/types"
	// "errors"

	"encoding/json"
	"fmt"
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

func FormatTable(input FormatterInput) (string, error) {
	var output strings.Builder

	output.WriteString(fmt.Sprintf("\nWeather Report: %s\n", input.Location.DisplayName))
	output.WriteString(fmt.Sprintf("Lat, Lon: %.2f, %.2f\n", input.Weather.Latitude, input.Weather.Longitude))

	times := input.Weather.Hourly["time"]
	output.WriteString(fmt.Sprintf("Period: %v to %v\n\n", times[0], times[len(times)-1]))

	cols := []int{}
	keys := []string{}
	output.WriteString("Time" + strings.Repeat(" ", 13) + "|") // 18 character space == 4 + 13 + 1

	for key := range input.Weather.Hourly {
		if key != "time" {
			cols = append(cols, len(key)+len(input.Weather.HourlyUnits[key])+1) // calculate required space by characters
			output.WriteString(fmt.Sprintf("%25s |", key))                      // 27 chars long!
			keys = append(keys, key)
		}
	}
	output.WriteString("\n")
	output.WriteString(strings.Repeat("=", 18+(len(input.Weather.Hourly)-1)*27) + "\n")

	for i, val := range times {

		output.WriteString(fmt.Sprintf("%-16s ", val))
		output.WriteString("|")

		for _, key := range keys {
			var comp strings.Builder
			comp.WriteString(fmt.Sprintf("%v %s", input.Weather.Hourly[key][i], input.Weather.HourlyUnits[key]))
			output.WriteString(fmt.Sprintf("%25s", comp.String()))
			output.WriteString(" |")
		}
		output.WriteString("\n")
	}

	// output.WriteString(fmt.Sprintf("Summary: %d hourly data points", len(times)))

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
