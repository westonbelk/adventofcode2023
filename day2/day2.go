package day2

import (
	"fmt"

	"github.com/westonbelk/adventofcode/util"
)

func Execute() {
	data := util.ReadLines("day2/input.txt")

	sum := 0

	for _, line := range data {
		answer, err := process(line)
		if err != nil {
			panic(err)
		}
		sum += answer
	}

	fmt.Println("total sum:", sum)
}

func process(input string) (int, error) {
	return 0, nil
}
