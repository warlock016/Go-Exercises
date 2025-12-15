package client

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type HAQuery struct {
	Database string // db: "home_assistant"
	Domain   string // domain: sensor
	EntityID string // entity_id: current_total, voltage, frequeny, total_power, forward_energy_total
	Name     string // __name__: mA_value, V_value, Hz_value, W_value, kWh_value
	From, To string
}

// (baseUrl, "home_assistant", "", "", "")
func (c *HAClient) RequestQuery(database, domain, entity_id, name string) (*http.Request, error) {
	// baseURL/prometheus/api/v1/query
	result, err := url.JoinPath(c.BaseURL, "prometheus", "api", "v1", "query")
	if err != nil {
		return nil, fmt.Errorf("Join URL failed: %v", err)
	}

	params := url.Values{}
	var constructor strings.Builder
	constructor.WriteString(name)
	elements := make([]string, 0, 4)
	if database != "" {
		elements = append(elements, fmt.Sprintf("db=\"%s\"", database))
	}
	if domain != "" {
		elements = append(elements, fmt.Sprintf("domain=\"%s\"", domain))
	}
	if entity_id != "" {
		elements = append(elements, fmt.Sprintf("entity_id=\"%s\"", entity_id))
	}
	if len(elements) != 0 {
		constructor.WriteString("{")
		for i, entry := range elements {
			constructor.WriteString(entry)
			if i != len(elements)-1 {
				constructor.WriteString(",")
			}
		}
		constructor.WriteString("}")
	}
	if len(constructor.String()) == 0 {
		return nil, fmt.Errorf("failed to construct query: empty string")
	}
	params.Set("query", constructor.String())
	fullURL := result + "?" + params.Encode()

	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}
	return request, nil
}

// (baseUrl, "home_assistant", "", "", "W_value", "2025-12-12", "2025-12-13")
func (c *HAClient) RequestQueryRange(database, domain, entityId, metricName, from, to string) (*http.Request, error) {

	if metricName == "" {
		return nil, fmt.Errorf("empty \"name\" input")
	}
	if from == "" {
		return nil, fmt.Errorf("empty \"from\" input")
	}
	if to == "" {
		// to = time.Now().Format("2006/02/01 15:04:05")
		return nil, fmt.Errorf("empty \"to\" input")
	}
	// baseURL/prometheus/api/v1/query_range
	result, err := url.JoinPath(c.BaseURL, "prometheus", "api", "v1", "query_range")
	if err != nil {
		return nil, fmt.Errorf("Join URL failed: %v", err)
	}

	params := url.Values{}
	var constructor strings.Builder
	elements := make([]string, 0, 4)

	constructor.WriteString(metricName)
	if database != "" {
		elements = append(elements, fmt.Sprintf("db=\"%s\"", database))
	}
	if domain != "" {
		elements = append(elements, fmt.Sprintf("domain=\"%s\"", domain))
	}
	if entityId != "" {
		elements = append(elements, fmt.Sprintf("entity_id=\"%s\"", entityId))
	}

	if len(elements) != 0 {
		constructor.WriteString("{")
		for i, entry := range elements {
			constructor.WriteString(entry)
			if i != len(elements)-1 {
				constructor.WriteString(",")
			}
		}
		constructor.WriteString("}")
	}

	if len(constructor.String()) == 0 {
		return nil, fmt.Errorf("failed to construct query: empty string")
	}

	params.Set("query", constructor.String())
	params.Set("start", from)
	params.Set("end", to)
	params.Set("step", "5m")
	fullURL := result + "?" + params.Encode()

	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %s %w", fullURL, err)
	}
	return request, nil
}
