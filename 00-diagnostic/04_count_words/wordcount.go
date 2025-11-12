package wordcount

import "strings"

// CountWords returns a map of each word and how many times it appears
func CountWords(text string) map[string]int {
	// TODO(human): Implement word counting

	wordcount := make(map[string]int, 0)

	fields := strings.Fields(text)

	for idx := range fields {
		wordcount[fields[idx]]++
	}

	// fmt.Println(slice)

	return wordcount
}
