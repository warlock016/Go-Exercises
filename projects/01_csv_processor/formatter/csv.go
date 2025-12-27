package formatter

import (
	"strconv"
	"strings"

	"github.com/warlock016/csv_processor/types"
)

func FormatCSV(parsed *types.ParsedData) string {
	sep := ";"
	var result strings.Builder
	header := make([]string, 0, 1+len(parsed.Labels))
	header = append(header, "timestamp")
	header = append(header, parsed.Labels...)
	result.WriteString(strings.Join(header, sep) + "\n")

	for i, v := range parsed.Time {
		row := []string{}
		row = append(row, v.Format("2006-01-02 15:04:05"))
		for _, w := range parsed.Labels {
			row = append(row, strconv.FormatFloat(parsed.Datapoints[w][i], 'f', 2, 64))
		}
		result.WriteString(strings.Join(row, sep) + "\n")
	}

	return result.String()
}
