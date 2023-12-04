package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/westonbelk/adventofcode/util"
)

var input []string
var cardCount = 0
var cardLookup = make(map[int]int, 0)

func Execute() {
	// input = util.ReadLines("day4/cali2.txt")
	input = util.ReadLines("day4/input.txt")

	for i := range input {
		preprocess(i + 1)
	}

	for i := range input {
		Process(i + 1)
	}

	fmt.Println("total:", cardCount)

}

func preprocess(n int) {
	line := input[n-1]

	total := 0
	winnerLookup := make(map[int]int, 0)

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
		if _, ok := winnerLookup[n]; ok {
			total++
		}
	}
	cardLookup[n] = total
}

func Process(n int) {
	cardCount++

	for i := 0; i < cardLookup[n]; i++ {
		Process(n + (i + 1))
	}
}
