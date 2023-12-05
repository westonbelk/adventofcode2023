package day5

import (
	"fmt"

	"github.com/westonbelk/adventofcode/util"
)

var input []string

func Execute() {
	// input = util.ReadLines("day4/cali2.txt")
	input = util.ReadLines("day4/input.txt")

	for i := range input {
		process(i)
	}

	fmt.Println("total:", 0)

}

func process(n int) {
}
