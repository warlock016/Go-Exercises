package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func Formatter(
	openMeteoResult *OpenMeteoResponse,
	geoApiResult *GeoApiResonse,
	ts []string,
	vars map[string][]float64,
	units map[string]string,
	format string,
) (string, error) {

	var result strings.Builder

	if len(ts) == 0 || len(vars) == 0 {
		return result.String(), fmt.Errorf("missing time series data")
	}

	varKeys := make([]string, 0, len(vars))
	for k := range vars {
		varKeys = append(varKeys, k)
	}

	switch format {
	case "table":
		result.WriteString("Weather for " + geoApiResult.Name + " " + geoApiResult.Address.Country + " (" + geoApiResult.Latitude[:5] + ", " + geoApiResult.Longitude[:5] + ")\n")
		result.WriteString(strings.Repeat("=", len(result.String())) + "\n")

		for i, v := range ts {
			v = strings.ReplaceAll(v, "T", " ")
			result.WriteString((v))
			result.WriteString("  |  ")

			// continue here --> WIP!
			for j, w := range varKeys {
				readings := strconv.FormatFloat(vars[w][i], 'f', 2, 64)

				if j == len(varKeys)-1 {
					result.WriteString(readings + " " + units[w])
				} else {
					result.WriteString(readings + " " + units[w] + "  | ")
				}
			}

			result.WriteString("\n")
			// WIP!
		}
	case "json":
		byteGeoApi, err := json.MarshalIndent(geoApiResult, "", "  ")
		if err != nil {
			return result.String(), fmt.Errorf("Marshal error %s", err)
		}

		result.WriteString(string(byteGeoApi))

		byteOpenMeteo, err := json.MarshalIndent(openMeteoResult, "", "  ")
		if err != nil {
			return result.String(), fmt.Errorf("Marshal error %s", err)
		}
		result.WriteString(string(byteOpenMeteo))
	default:
		return result.String(), fmt.Errorf("invalid output format: %s", format)
	}

	return result.String(), nil
}
