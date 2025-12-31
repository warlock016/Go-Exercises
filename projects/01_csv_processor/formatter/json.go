package formatter

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/warlock016/csv_processor/types"
)

func createRowJSON(data *types.ParsedData) types.JsonData {
	result := types.JsonData{
		Source:   data.Source,
		Timezone: data.Timezone.String(),
		Count:    len(data.Time),
		Records:  formatRowJSON(data),
	}

	return result
}

func formatRowJSON(data *types.ParsedData) []types.RowRecord {

	result := make([]types.RowRecord, 0, len(data.Labels)+1)

	for idx, t := range data.Time {
		res := make(types.RowRecord, len(data.Labels)+1)
		res["time"] = t.Format(time.RFC3339)
		for _, label := range data.Labels {
			if idx < len(data.Datapoints[label]) {
				val := data.Datapoints[label][idx]
				if math.IsNaN(val) {
					res[label] = nil
				} else {
					res[label] = val
				}
			}
		}
		result = append(result, res)
	}

	return result
}

/*
Row-Oriented JSON:

	{
		"source":"file.csv",
		"timezone":"America/Costa_Rica",
		"count":"4464",
		"records":[
			{"time":"2025-01-01T00:00:00", "Temperature":20.5, "Humidity":65.0},
			{"time":"2025-01-01T00:10:00", "Temperature":21.0, "Humidity":64.5},
			{..., ..., ...},
			...
		]
	}
*/

func timeToString(x []time.Time) []string {
	result := make([]string, 0, len(x))

	for _, v := range x {
		result = append(result, v.Format(time.RFC3339))
	}

	return result
}

func createColJSON(data *types.ParsedData) types.JsonData {

	return types.JsonData{
		Source:     data.Source,
		Timezone:   data.Timezone.String(),
		Count:      len(data.Labels),
		Time:       timeToString(data.Time),
		Labels:     data.Labels,
		Datapoints: formatColJSON(data),
	}
}

func formatColJSON(data *types.ParsedData) types.ColRecord {
	/*
			Col-Oriented JSON:
		{
			"source":"file.csv",
			"timezone":"America/Costa_Rica",
			"count":"4464",
			"time":["2025-01-01T00:00:00", "2025-01-01T00:10:00"]
			labels:["Temperature", ...],
			"datapoints": {
				"Temperature": 	[20.5, 65.0, ...],
				"Humidity": 	[65.0, 64.5, ...],
				...
			}

		}
	*/

	result := make(types.ColRecord)

	for _, label := range data.Labels {
		for _, v := range data.Datapoints[label] {
			if math.IsNaN(v) {
				result[label] = append(result[label], nil)
			} else {
				result[label] = append(result[label], v)
			}
		}
	}

	return result
}

// returns a string with a pretty json
func StringifyJSON(data types.JsonData) string {

	res, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return string(res)
}
