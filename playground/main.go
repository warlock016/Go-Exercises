package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type WorkResult struct {
	Output string
	Error  error
}

func doWork(s string, milliseconds time.Duration) WorkResult {

	var output strings.Builder
	result := WorkResult{}

	start := time.Now()
	output.WriteString(fmt.Sprintf("began work: %s\n", start.Format(time.RFC3339)))
	time.Sleep(milliseconds)
	if output.String() == "" {
		result.Error = fmt.Errorf("invalid work output %s", s)
	} else {
		output.WriteString(fmt.Sprintf("finished work. Elapsed time %v ms\n", time.Since(start)))
	}
	result.Output = output.String()
	return result
}

func main() {
	results := make(chan WorkResult, 3)
	params := map[string]time.Duration{"fast": 200 * time.Millisecond, "normal": 400 * time.Millisecond, "slow": 700 * time.Millisecond}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	for key, dur := range params {
		go func(key string, dur time.Duration) {
			results <- doWork(key, dur)
		}(key, dur)
	}
	count := 0
	for range params {
		select {
		case res := <-results:
			if res.Error != nil {
				fmt.Printf("Error: %v\n", res.Error)
			} else {
				count++
				fmt.Println(res.Output)
			}
		case <-ctx.Done():
			fmt.Printf("Timeout! Only received %d/%d results\n", count, len(params))
			return
		}
	}
	fmt.Println("All workers completed")
}
