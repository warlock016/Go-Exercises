package formatter

import (
	"fmt"

	"github.com/warlock016/csv_processor/types"
)

func FormatData(data *types.ParsedData, format string) (*types.FormattedOutput, error) {

	result := types.FormattedOutput{}

	switch format {
	case "csv":
		result.Content = FormatCSV(data)
		result.Type = format
	case "json-col":
		result.Content = StringifyJSON(CreateColJSON(data))
		result.Type = "json"
	case "json-row":
		result.Content = StringifyJSON(CreateRowJSON(data))
		result.Type = "json"
	default:
		return nil, fmt.Errorf("format not implemented yet")
	}

	return &result, nil
}
