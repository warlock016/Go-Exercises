package channel_reinforcement

import (
	"time"
)

// Echo sends a value to a goroutine which doubles it and returns the result
func Echo(value int) int {
	// TODO(human): Implement
	ch := make(chan int)

	go func() {
		val := <-ch
		val *= 2
		ch <- val
	}()

	ch <- value
	return <-ch
}

// Countdown returns a channel that sends n, n-1, ..., 1, then closes
func Countdown(n int) <-chan int {
	// TODO(human): Implement
	ch := make(chan int)

	go func() {
		for i := n; i > 0; i-- {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

// Relay passes a value through n goroutines, each adding 1
func Relay(value int, stages int) int {
	// TODO(human): Implement

	if stages == 0 {
		return value
	}

	result := make(chan int)
	comm := make([]chan int, stages)
	for i := range comm {
		comm[i] = make(chan int)
	}

	for i := range stages {
		go func() {
			val := <-comm[i]
			val++
			if i < stages-1 {
				comm[i+1] <- val
			} else {
				result <- val
			}
		}()
	}

	comm[0] <- value
	return <-result
}

// FanIn launches one goroutine per value, each sends to a shared channel
func FanIn(values []int) int {
	// TODO(human): Implement

	if len(values) == 0 {
		return 0
	}

	result := make(chan int)

	for _, v := range values {
		go func(val int) {
			result <- v
		}(v)
	}

	var val int
	for range values {
		val += <-result
	}
	return val
}

// Ticker returns a channel that sends 0, 1, 2, ..., n-1 with interval delays
func Ticker(n int, interval time.Duration) <-chan int {
	// TODO(human): Implement
	ch := make(chan int)

	go func() {
		for i := range n {
			ch <- i
			if i < n {
				time.Sleep(interval)
			}
		}
		close(ch)
	}()

	return ch
}

// Pipeline: Generate(1..n) → FilterEven → Double → AddOne → FilterOdd → Sum
func Pipeline(n int) int {
	// TODO(human): Implement

	if n == 0 {
		return 0
	}

	// we need 5 channels for communication between stages:
	comm := make([]chan int, 5)
	for i := range 5 {
		comm[i] = make(chan int)
	}

	// channel for generating (1...n)
	go func() {
		// loops over 1...n
		for i := 1; i <= n; i++ {
			comm[0] <- i
		}
		close(comm[0])
	}()

	// goroutine for filtering even numbers (keeps even, removes odd)
	go func() {
		for val := range comm[0] {
			if val%2 == 0 {
				comm[1] <- val
			}
		}
		close(comm[1])
	}()

	// goroutine for doubling values
	go func() {
		for val := range comm[1] {
			val *= 2
			comm[2] <- val
		}
		close(comm[2])
	}()

	// goroutine for incrementing 1
	go func() {
		for val := range comm[2] {
			val += 1
			comm[3] <- val
		}
		close(comm[3])
	}()

	// goroutine for filtering odd numbers (keeps odd, eliminates even)
	go func() {
		for val := range comm[3] {
			if val%2 != 0 {
				comm[4] <- val
			}
		}
		close(comm[4])
	}()

	var res int
	for {
		val, ok := <-comm[4]
		if !ok {
			break
		}
		res += val
	}

	return res
}
