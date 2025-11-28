package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func rowCreator(pos, n int) ([]string, error) {

	row := make([]string, 0, n)

	if pos >= n {
		return nil, fmt.Errorf("q position: out of bounds")
	}

	for i := range n {
		if i == pos {
			row = append(row, "Q")
		} else {
			row = append(row, ".")
		}
	}
	return row, nil
}

func randPos(n int) int {
	return rand.IntN(n)
}

func separator(n int, s string) string {

	var result strings.Builder

	for range n {
		result.WriteString(s)
	}

	return result.String()
}

func main() {

	n := 2

	// board := [][]string{}

	for i := range n {
		// result, err := rowCreator(randPos(n), n)
		result, err := rowCreator(i, n)
		if err != nil {
			fmt.Printf("loop error")
		}
		fmt.Printf("%d: %v\n", i, result)
		// board = append(board, result)
	}

	// n = 6; l = 16; d = 10
	// n = 5; l = 14; d = 9
	// n = 4; l = 12; d = 8
	// n = 3; l = 10; d = 7
	// n = 2; l = 8;  d = 6
	fmt.Println(separator(n+6, "-"))

	for i := range n {
		result, err := rowCreator(randPos(n), n)
		// result, err := rowCreator(i, n)
		if err != nil {
			fmt.Printf("loop error")
		}
		fmt.Printf("%d: %v\n", i, result)
		// board = append(board, result)
	}

	// fmt.Println(board)
}
