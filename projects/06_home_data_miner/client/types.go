package client

import "encoding/json"

type VMEnvelope struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
	Error  string          `json:"error,omitempty"`
}
type QueryResponse struct {
	ResultType string          `json:"resultType"`
	Result     json.RawMessage `json:"result"`
}
type InstantResponse struct {
	ResultType string          `json:"resultType"`
	Result     []InstantMetric `json:"result"`
}
type InstantMetric struct {
	Metadata Metric `json:"metric"`
	Value    [2]any `json:"value"`
}
type RangeResponse struct {
	ResultType string        `json:"resultType"`
	Result     []RangeMetric `json:"result"`
}
type RangeMetric struct {
	Metadata Metric   `json:"metric"`
	Values   [][2]any `json:"values"`
}
type Metric struct {
	Name     string `json:"__name__"`
	DB       string `json:"db"`
	Domain   string `json:"domain"`
	EntityId string `json:"entity_id"`
}

type Labels []string

type Status struct {
	TotalSeries             int          `json:"totalSeries"`
	TotalLabelValuePairs    int          `json:"totalLabelValuePairs"`
	SeriesCountByMetricName []SeriesName `json:"seriesCountByMetricName"` // XXX!!
}

type SeriesName struct {
	Name                 string `json:"name"`
	Value                int    `json:"value"`
	RequestsCount        int    `json:"requestsCount"`
	LastRequestTimestamp int64  `json:"lastRequestTimestamp"`
}

type SeriesLabel struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type SeriesValuePair struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// type LabelValueCountByLabelName
