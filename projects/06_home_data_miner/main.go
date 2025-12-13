package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/hashicorp/go-envparse"
)

func LoadEnv(path string) (map[string]string, error) {
	// Dummy implementation for loading environment variables from a file
	// In a real implementation, you would read the file and parse key-value pairs

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	env, err := envparse.Parse(file)
	if err != nil {
		return nil, err
	}

	return env, nil
}

func main() {
	// Before creating http requests, we typically need to set up some configurations
	// such as API endpoints, headers, authentication tokens, etc.
	// These configurations can be hardcoded, read from a config file, or passed as environment variables.
	envConfig := "./.env"

	env, err := LoadEnv(envConfig)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	url, err := url.JoinPath(fmt.Sprintf("%s:%s", env["HOME_ASSISTANT_URL"], env["HOME_ASSISTANT_PORT"]), "api/")
	if err != nil {
		log.Fatalf("invalid URL: %v", err)
	}

	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalf("failed to create request: %v", err)
	}

	// Step 2: Set headers
	req.Header.Set("Authorization", "Bearer "+env["HOME_ASSISTANT_TOKEN"])
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("Response status: %s\n", resp.Status)
}
