package day5

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/westonbelk/adventofcode/util"
)

type Mapping struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func (m Mapping) Contains(n int) bool {
	return n >= m.SourceRangeStart && n <= m.SourceRangeStart+m.RangeLength
}

func (m Mapping) Transform(n int) int {
	diff := n - m.SourceRangeStart
	return m.DestinationRangeStart + diff
}

var input []string

func Execute() {
	start := time.Now()

	// input = util.ReadLines("day5/calibration.txt")
	input = util.ReadLines("day5/input.txt")

	seeds := readNums(strings.Fields((input[0]))[1:])
	maps := mappings()
	numWorkers := 10
	ch := make(chan int, numWorkers)
	var wg sync.WaitGroup

	fmt.Println("seeds:", seeds)

	for i := 0; i < len(seeds); i += 2 {
		rangeStart := seeds[i]
		rangeEnd := seeds[i] + seeds[i+1]

		// do work
		wg.Add(1)
		go func(rangeStart, rangeEnd int, ch chan int, wg *sync.WaitGroup) {
			fmt.Println("started seed")
			defer wg.Done()
			lowest := 9999999999999999
			for seedOriginal := rangeStart; seedOriginal < rangeEnd; seedOriginal++ {
				seed := seedOriginal
				for _, kind := range maps {
					for _, mapping := range kind {
						if mapping.Contains(seed) {
							seed = mapping.Transform(seed)
							break
						}
					}
				}
				if seed < lowest {
					lowest = seed
				}
			}
			fmt.Println("finished seed")
			ch <- lowest
		}(rangeStart, rangeEnd, ch, &wg)
	}

	wg.Wait()
	close(ch)
	overallLowest := 9999999999999999
	for result := range ch {
		if result < overallLowest {
			overallLowest = result
		}
	}
	fmt.Println("overall lowest:", overallLowest)
	fmt.Println("duration:", time.Since(start).Seconds())
}

func mappings() [][]Mapping {
	kind := 0
	entry := 0
	res := make([][]Mapping, 0)
	res = append(res, make([]Mapping, 0))

	for i := 3; i < len(input); i++ {
		s := input[i]

		if s == "" {
			continue
		}
		if unicode.IsLetter(rune(s[0])) {
			res = append(res, make([]Mapping, 0))
			kind++
			entry = 0
			continue
		}

		fields := strings.Fields(s)
		nums := readNums(fields)
		if len(nums) == 0 {
			fmt.Printf("%v: %q\n", i, s)
		}
		res[kind] = append(res[kind], newMap(nums))

		entry++
	}
	return res
}

func readNums(s []string) []int {
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

func newMap(arr []int) Mapping {
	if len(arr) != 3 {
		panic(fmt.Sprintf("invalid number of values in newMap. want %v; got %v", 3, len(arr)))
	}

	return Mapping{
		DestinationRangeStart: arr[0],
		SourceRangeStart:      arr[1],
		RangeLength:           arr[2],
	}
}
