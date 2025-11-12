package fizzbuzz

import "strconv"

// FizzBuzz returns a slice of strings from 1 to n with FizzBuzz rules:
// - "Fizz" for multiples of 3
// - "Buzz" for multiples of 5
// - "FizzBuzz" for multiples of both
// - The number as a string otherwise
func FizzBuzz(n int) []string {
	// TODO(human): Implement FizzBuzz

	output := make([]string, 0)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 != 0 {
			output = append(output, "Fizz")
		} else if i%3 != 0 && i%5 == 0 {
			output = append(output, "Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			output = append(output, "FizzBuzz")
		} else {
			output = append(output, strconv.Itoa(i))
		}
	}

	return output
}
