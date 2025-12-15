package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/warlock016/projects/data_miner/client"
)

func dumpData(data []byte) error {

	cache := "./cache"

	if len(data) == 0 {
		return fmt.Errorf("invalid buffer length: 0")
	}

	if _, err := os.Stat(cache); os.IsNotExist(err) {
		os.Mkdir(cache, 0o644)
	}

	dir, err := os.ReadDir(cache)
	if err != nil {
		return fmt.Errorf("unexpected error: %w", err)
	}

	if len(dir) != 0 {
		for _, file := range dir {
			if file.Name() == cache {
				filePath, err := os.Getwd()
				if err != nil {
					return fmt.Errorf("unexpected directory path error: %w", err)
				}
				path.Join(filePath, "cache", file.Name())
				return fmt.Errorf("cache already exists: %s", filePath)
			}
		}
	}

	os.WriteFile(path.Join(cache, "cache.json"), data, 0o644)

	return nil
}

func main() {
	// Before creating http requests, we typically need to set up some configurations
	// such as API endpoints, headers, authentication tokens, etc.
	// These configurations can be hardcoded, read from a config file, or passed as environment variables.
	envConfig := "./.env"

	env, err := client.LoadEnv(envConfig)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	// We first build the base url for:
	baseUrl := fmt.Sprintf("%s:%s", env["HOME_ASSISTANT_URL"], env["TSDB_PORT"])
	haClient := client.NewHAClient(baseUrl, 30*time.Second)

	statusReq, err := haClient.RequestQueryRange("home_assistant", "", "", "V_value", "2025-12-13", "2025-12-14")
	// statusReq, err := haClient.RequestQuery("", "", "total_power", "")
	// statusReq, err := haClient.RequestLabels()
	if err != nil {
		log.Fatalf("failed to construct query: %v", err)
	}

	resp, err := haClient.Client.Do(statusReq)
	if err != nil {
		if resp != nil {
			log.Fatalf("request failed w/Status (%d): %v", resp.StatusCode, err)
		} else {
			log.Fatalf("unexpected nil response error: %v", err)
		}
	}
	if resp == nil {
		log.Fatal("unexpected nil response with nil error")
	}
	defer resp.Body.Close()

	var vmResponse client.VMEnvelope
	var rawData client.QueryResponse

	var instantData client.InstantResponse
	var rangeData client.RangeResponse

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}
	err = json.Unmarshal(body, &vmResponse)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON: %v", err)
	}

	err = json.Unmarshal(vmResponse.Data, &rawData)
	if err != nil {
		log.Fatalf("failed to unmarshal data: %v", err)
	}

	switch rawData.ResultType {
	case "vector":
		var results []client.InstantMetric
		err = json.Unmarshal(rawData.Result, &results)
		instantData.Result = results
	case "matrix":
		var results []client.RangeMetric
		err = json.Unmarshal(rawData.Result, &results)
		rangeData.Result = results
	}
	if err != nil {
		log.Fatalf("failed to unmarshal results: %v", err)
	}

	if vmResponse.Error != "" {
		fmt.Printf(" Error: %s", vmResponse.Error)
	}
	fmt.Print("\n")

	fmt.Printf("%v\n", instantData.Result)
	fmt.Printf("%v\n", rangeData.Result)
	// dumpData(data.Bytes())

}
