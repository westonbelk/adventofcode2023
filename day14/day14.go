package day14

import (
	"fmt"
	"strings"

	"github.com/westonbelk/adventofcode/util"
)

func Rotated(grid []string) []string {
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
	input := util.ReadLines("day14/input.txt")
	grid := Rotated(input)
	grid = FallLeftGrid(grid)
	grid = Rotated(grid)
	fmt.Println(strings.Join(grid, "\n"))
	fmt.Println(WeighGrid(grid))
}
