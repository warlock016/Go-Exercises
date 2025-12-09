package closure_syntax

import (
	"fmt"
	"time"
)

// Exercise 1: StatefulOperation
// Returns two closures sharing state: add modifies, getValue retrieves
func StatefulOperation(initial int) (add func(int), getValue func() int) {
	// TODO(human): Implement

	init := initial
	add = func(x int) {
		init += x
	}

	getValue = func() int {
		return init
	}

	return add, getValue
}

// Exercise 2: UndoRedo
// Returns three closures for undo/redo functionality
func UndoRedo() (do func(int), undo func() (int, bool), redo func() (int, bool)) {
	// TODO(human): Implement

	history := []int{}
	undone := []int{}

	do = func(x int) {
		history = append(history, x)
		undone = []int{}
	}

	undo = func() (int, bool) {
		switch len(history) {
		case 0:
			return 0, false
		case 1:
			last := history[len(history)-1]
			undone = append(undone, last)
			history = history[:len(history)-1]
			return 0, true
		default:
			last := history[len(history)-1]
			undone = append(undone, last)
			history = history[:len(history)-1]
			return history[len(history)-1], true
		}
	}

	redo = func() (int, bool) {
		switch len(undone) {
		case 0:
			return 0, false
		default:
			last := undone[len(undone)-1]
			undone = undone[:len(undone)-1]
			history = append(history, last)
			return last, true
		}
	}
	return do, undo, redo
}

// Exercise 3: Logger
// Wraps a function with logging of inputs and outputs
func Logger(prefix string, fn func(int) int) func(int) int {
	// TODO(human): Implement
	return func(i int) int {
		fmt.Printf("[%s] input: %d\n", prefix, i)
		result := fn(i)
		fmt.Printf("[%s] output: %d\n", prefix, result)
		return result
	}
}

// Exercise 4: Compose
// Returns f(g(x)) - applies g first, then f
func Compose(f, g func(int) int) func(int) int {
	// TODO(human): Implement
	return func(i int) int {
		gRes := g(i)
		fRes := f(gRes)
		return fRes
	}
}

// Exercise 5: TimingMiddleware
// Returns middleware that logs slow function execution
func TimingMiddleware(threshold time.Duration) func(func(int) int) func(int) int {
	// TODO(human): Implement

	return func(f func(int) int) func(int) int {

		return func(i int) int {
			start := time.Now()
			result := f(i)
			elapsed := time.Since(start)

			if elapsed > threshold {
				fmt.Printf("SLOW! Threshold exceeded: %v\n", elapsed)
			}
			return result
		}
	}
}

// Exercise 6: PartialApply
// Pre-fills the first argument of a two-argument function
func PartialApply(fn func(int, int) int, first int) func(int) int {
	// TODO(human): Implement
	return func(i int) int {
		return fn(first, i)
	}
}

// Exercise 7: Pipeline
// Returns closures to build and execute a processing pipeline
func Pipeline() (add func(func(int) int), execute func(int) int) {
	// TODO(human): Implement
	pipeline := []func(int) int{}

	add = func(x func(int) int) {
		pipeline = append(pipeline, x)
	}

	execute = func(i int) int {
		result := i
		for _, fn := range pipeline {
			result = fn(result)
		}
		return result
	}

	return add, execute
}

// Exercise 8: Builder
// OptionBuilder is a self-referential function type for fluent building
type OptionBuilder func(string) OptionBuilder

// NewOptionBuilder returns a chainable builder and a function to get results
func NewOptionBuilder() (builder OptionBuilder, getOptions func() []string) {
	// TODO(human): Implement

	options := []string{}

	builder = func(s string) OptionBuilder {
		options = append(options, s)
		return builder
	}

	getOptions = func() []string {
		return options
	}

	return builder, getOptions
}

// Exercise 9: EventBus
// Returns pub/sub closures coordinated via channel
func EventBus(bufferSize int) (publish func(string), subscribe func() <-chan string, closeBus func()) {
	// TODO(human): Implement

	ch := make(chan string, bufferSize)
	var closed bool

	publish = func(s string) {
		if !closed { // prevents panic
			ch <- s
		}
	}

	subscribe = func() <-chan string {
		return ch
	}

	closeBus = func() {
		closed = true
		close(ch)
	}

	return publish, subscribe, closeBus
}

// Exercise 10: MiddlewareChain
// Handler processes a string and returns a string
type Handler func(string) string

// Middleware wraps a Handler and returns a new Handler
type Middleware func(Handler) Handler

// Chain combines multiple middleware into a single middleware
func Chain(middlewares ...Middleware) Middleware {
	// TODO(human): Implement

	return func(h Handler) Handler {
		handler := h
		for i := len(middlewares) - 1; i >= 0; i-- {
			handler = middlewares[i](handler)
		}
		return handler
	}

}
