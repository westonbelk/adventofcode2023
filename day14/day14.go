package day14

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/westonbelk/adventofcode/util"
)

//go:generate stringer -type=iterations
type iterations int

const (
	three      iterations = 3
	billion    iterations = 1000000000
	onepercent iterations = billion / 100
)

func NewGridCacheEntry(g []string) Grid {
	grid := Grid{
		Hash: strings.ReplaceAll(ZipGrid(g), ".", ""),
		Grid: g,
	}
	return grid
}

type Grid struct {
	Hash string
	Grid []string
}

var cache = make(map[string]Grid, 0)
var cacheHits int = 0

func ZipGrid(grid []string) string {
	return strings.Join(grid, "")
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

func Cycle(grid Grid) Grid {
	next, ok := cache[grid.Hash]
	if ok {
		cacheHits++
		return next
	}

	res := slices.Clone(grid.Grid)
	res = FallEastGrid(FallSouthGrid(FallWestGrid(FallNorthGrid(res))))
	nextCacheEntry := NewGridCacheEntry(res)
	cache[grid.Hash] = nextCacheEntry
	return nextCacheEntry
}

func FallNorthGrid(grid []string) []string {
	return Transposed(FallLeftGrid(Transposed(grid)))
}

func FallWestGrid(grid []string) []string {
	return FallLeftGrid(grid)
}

func FallSouthGrid(res []string) []string {
	slices.Reverse(res)
	res = FallNorthGrid(res)
	slices.Reverse(res)
	return res
}

func FallEastGrid(res []string) []string {
	for row := range res {
		r := []rune(res[row])
		slices.Reverse(r)
		res[row] = string(r)
	}
	res = FallWestGrid(res)
	for row := range res {
		r := []rune(res[row])
		slices.Reverse(r)
		res[row] = string(r)
	}
	return res
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

func FallLeftGrid(res []string) []string {
	for i, row := range res {
		row = FallLeft(row)
		res[i] = row
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
	input := util.ReadLines("day14/input.txt")
	iterations := onepercent
	fmt.Printf("running %s iterations\n", iterations)
	start := time.Now()
	grid := NewGridCacheEntry(input)
	for i := 1; i <= int(iterations); i++ {
		grid = Cycle(grid)
		// fmt.Println("After", i, "cycles:")
		// fmt.Println(strings.Join(input, "\n"))
		// fmt.Println()
	}
	if iterations == onepercent {
		dur := time.Since(start)
		fmt.Printf("estimated duration for %s: %s\n", billion, dur*100)
	}

	fmt.Println("cache entries:", len(cache))
	fmt.Println("cache hits:", cacheHits)
	fmt.Println("load", WeighGrid(input))
}
