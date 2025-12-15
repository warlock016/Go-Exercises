package client

import (
	"net"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/go-envparse"
)

type HAClient struct {
	Client  *http.Client
	BaseURL string
}

func NewHAClient(baseURL string, timeout time.Duration) *HAClient {
	tr := http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second, // mDNS (.local) resolution needs more time than regular DNS
			KeepAlive: 30 * time.Second,
		}).DialContext,

		TLSHandshakeTimeout:   800 * time.Millisecond,
		ResponseHeaderTimeout: 2 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		IdleConnTimeout:       30 * time.Second,
		MaxIdleConns:          10,
		MaxConnsPerHost:       10,
	}

	return &HAClient{
		Client: &http.Client{
			Transport: &tr,
			Timeout:   timeout,
		},
		BaseURL: baseURL,
	}
}

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
