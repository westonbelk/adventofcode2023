package util

import "strconv"

// takes numbers as strings and returns them as ints
func ReadNums(s []string) []int {
	numbers := make([]int, 0, len(s))
	for _, s := range s {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}
	return numbers
}
