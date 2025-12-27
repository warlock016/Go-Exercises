package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Datapoint struct {
	Label     string
	Timestamp time.Time
	Value     float64
}

func generateDataPoint(label string) Datapoint {
	return Datapoint{
		Label:     label,
		Timestamp: time.Now(),
		Value:     rand.Float64(),
	}
}

func (d Datapoint) String() string {
	return fmt.Sprintf("%s: %v %.2f", d.Label, d.Timestamp, d.Value)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Millisecond * 500)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

/*
rough concept:
implement an endless queue that returns a preconfigured set of timestamp/value pairs via an API until the user stops the service.
The service should shut down gracefully when "Ctrl+C"
Ensure no memory leaks and correct status response via API
*/

func main() {

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan struct{})
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		<-gracefulShutdown
		close(done)
	}()

	for {
		select {
		case <-done:
			fmt.Println("shutting down gracefully...")
			return
		case <-ticker.C:
			v := generateDataPoint("Voltage")
			fmt.Println(v)
		}
	}
}
