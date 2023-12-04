package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/westonbelk/adventofcode/util"
)

var input []string

func Execute() {
	// input := util.ReadLines("day4/calibration.txt")
	input = util.ReadLines("day4/input.txt")

	sum := 0
	for i := range input {
		res := Process(i)
		sum += res
	}
	fmt.Println("sum:", sum)

}

func Process(n int) int {
	line := input[n]

	total := 0
	winnerLookup := make(map[int]int, 0)
	// myNumbers := make([]int, 0, 0)

	cardData := strings.Split(line, ": ")[1]
	cardDataSplit := strings.Split(cardData, " | ")
	winningNumbersRaw, myNumbersRaw := cardDataSplit[0], cardDataSplit[1]

	for _, s := range strings.Fields(winningNumbersRaw) {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		winnerLookup[n] = 2
	}

	for _, s := range strings.Fields(myNumbersRaw) {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		if v, ok := winnerLookup[n]; ok {
			if total == 0 {
				total = 1
			} else {
				total = total * v
			}
		}
	}
	return total
}

// func seedWinnerLookup() map[int]int {
// 	ret := make(map[int]int, 0)
// 	for i := 0; i < 101; i++ {
// 		ret[i] = 1
// 	}
// 	return ret
// }
