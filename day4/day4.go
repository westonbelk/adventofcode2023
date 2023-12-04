package day4

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/westonbelk/adventofcode/util"
)

var input []string
var cardCount = 0
var mutex sync.Mutex
var wg sync.WaitGroup

func Execute() {
	input = util.ReadLines("day4/cali2.txt")
	// input = util.ReadLines("day4/input.txt")

	for i := range input {
		chain := ""
		Process(i+1, chain)
	}

	wg.Wait()
	fmt.Println("total:", cardCount)

}

func Inc() {
	mutex.Lock()
	cardCount = cardCount + 1
	mutex.Unlock()
}

func Process(n int, chain string) int {
	chain += fmt.Sprintf("%d -> ", n)
	// fmt.Println(chain)
	Inc()
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

	for i := 0; i < total; i++ {
		Process(n+(i+1), chain)
	}
	return 0
}

// func seedWinnerLookup() map[int]int {
// 	ret := make(map[int]int, 0)
// 	for i := 0; i < 101; i++ {
// 		ret[i] = 1
// 	}
// 	return ret
// }
