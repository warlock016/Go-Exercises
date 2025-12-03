package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func Formatter(openMeteoResult *OpenMeteoResponse, geoApiResult *GeoApiResonse, ts []string, vars map[string][]float64, units map[string]string, format string) (string, error) {

	var result strings.Builder

	switch format {
	case "table":
		result.WriteString("Weather for ")
		result.WriteString(geoApiResult.Name + " " + geoApiResult.Address.Country + " (" + geoApiResult.Latitude[:5] + ", " + geoApiResult.Longitude[:5] + ")")
		result.WriteString(strings.Repeat("=", len(result.String())))
		for i, v := range ts {
			result.WriteString(v)
			result.WriteString("  |  ")

			// continue here --> WIP!
			for j, dp := range vars {
				strconv.FormatFloat(vars[j][i])
			}
			// WIP!
		}
	case "json":
		byteRep, err := json.MarshalIndent(openMeteoResult, "", "  ")
		if err != nil {
			return result.String(), fmt.Errorf("Marshal error %s", err)
		}
		result.WriteString(string(byteRep))
	default:
		return result.String(), fmt.Errorf("invalid output format: %s", format)
	}

	return result.String(), nil
}
