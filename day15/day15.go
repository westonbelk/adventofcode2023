package day15

import (
	"fmt"
	"strings"

	"github.com/westonbelk/adventofcode/util"
)

func Hash(s string) int {
	value := 0
	for _, r := range s {
		value += int(r)
		value *= 17
		value = value % 256
	}
	return value
}

func Execute() {
	input := util.ReadLines("day15/input.txt")[0]

	sum := 0
	for _, e := range strings.Split(input, ",") {
		sum += Hash(e)
	}
	fmt.Println("Sum:", sum)
}
