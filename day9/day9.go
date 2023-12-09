package day9

import (
	"fmt"
	"strings"

	"github.com/westonbelk/adventofcode/util"
)

func nextNum(rowStr string) int {
	numbers := make([][]int, 0)
	numbers = append(numbers, util.ReadNums(strings.Fields(rowStr)))
	for {
		lastRow := numbers[len(numbers)-1]
		if lastRow[0] == 0 && lastRow[len(lastRow)-1] == 0 {
			break // if the first and last are zero then the whole thing should be
		}

		nextRow := make([]int, 0, len(lastRow)-1)

		for i := 0; i < (len(lastRow) - 1); i++ {
			nextRow = append(nextRow, lastRow[i+1]-lastRow[i])
		}
		numbers = append(numbers, nextRow)
	}

	// placeholder 0
	numbers[len(numbers)-1] = append([]int{0}, numbers[len(numbers)-1]...)
	for i := len(numbers) - 1; i > 0; i-- {
		bottomRow := numbers[i]
		topRow := numbers[i-1]
		loc := 0
		sum := topRow[loc] - bottomRow[loc]
		numbers[i-1] = append([]int{sum}, numbers[i-1]...)
	}

	for _, r := range numbers {
		fmt.Println(r)
	}
	return numbers[0][0]
}

func Execute() {
	input := util.ReadLines("day9/input.txt")
	// input := util.ReadLines("day9/calibration.txt")

	sum := 0
	for _, s := range input {
		sum += nextNum(s)
		fmt.Println()
	}

	fmt.Println(sum)
}
