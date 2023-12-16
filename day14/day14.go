package day14

import (
	"fmt"
	"slices"
	"strings"

	"github.com/westonbelk/adventofcode/util"
)

const (
	billion    int = 1000000000
	onepercent int = billion / 100
)

var cache = make(map[string]string, 0)

func ZipGrid(grid []string) string {
	return strings.Join(grid, "\n")
}

func UnzipGrid(zipped string) []string {
	return strings.Split(zipped, "\n")
}

func Transposed(grid []string) []string {
	res := make([]string, 0)
	for x := range grid[0] {
		row := make([]byte, 0)
		for y := range grid {
			row = append(row, grid[y][x])
		}
		res = append(res, string(row))
	}
	return res
}

func Cycle(grid []string) []string {
	gridZipped := ZipGrid(grid)
	cached, ok := cache[gridZipped]
	if ok {
		return UnzipGrid(cached)
	}

	res := FallEastGrid(FallSouthGrid(FallWestGrid(FallNorthGrid(grid))))
	cache[gridZipped] = ZipGrid(res)
	return res
}

func FallNorthGrid(grid []string) []string {
	return Transposed(FallLeftGrid(Transposed(grid)))
}

func FallWestGrid(grid []string) []string {
	return FallLeftGrid(grid)
}

func FallSouthGrid(grid []string) []string {
	res := slices.Clone(grid)
	slices.Reverse(res)
	res = FallNorthGrid(res)
	slices.Reverse(res)
	return res
}

func FallEastGrid(grid []string) []string {
	res := slices.Clone(grid)
	for row := range res {
		r := []rune(res[row])
		slices.Reverse(r)
		res[row] = string(r)
	}
	fell := FallWestGrid(res)
	for row := range fell {
		r := []rune(fell[row])
		slices.Reverse(r)
		fell[row] = string(r)
	}
	return fell
}

func FallLeft(s string) string {
	r := []rune(s)
	rocks := 0
	startPos := 0
	for i, c := range r {
		switch c {
		case 'O':
			rocks++
			r[i] = '.'
		case '#':
			for n := 0; n < rocks; n++ {
				r[startPos+n] = 'O'
			}
			startPos = i + 1
			rocks = 0
		}
	}
	for n := 0; n < rocks; n++ {
		r[startPos+n] = 'O'
	}
	return string(r)
}

func FallLeftGrid(grid []string) []string {
	res := make([]string, 0)
	for _, r := range grid {
		r = FallLeft(r)
		res = append(res, r)
	}
	return res
}

func WeighGrid(grid []string) int {
	sum := 0
	for y := range grid {
		row := grid[y]
		for _, r := range row {
			if r == 'O' {
				sum += (len(grid) - y)
			}
		}
	}
	return sum
}

func Execute() {
	input := util.ReadLines("day14/calibration.txt")
	for i := 1; i <= billion; i++ {
		input = Cycle(input)
		// fmt.Println("After", i, "cycles:")
		// fmt.Println(strings.Join(input, "\n"))
		// fmt.Println()
	}

	fmt.Println("load", WeighGrid(input))
}
