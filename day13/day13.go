package day13

import (
	"fmt"
	"strings"

	"github.com/westonbelk/adventofcode/util"
)

// int  => -1 if not found
// otherwise index = value
func rowMirror(grid []string) int {
	for i := 0; i < len(grid)-1; i++ {
		if grid[i] == grid[i+1] {
			t, b := i, i+1
			mismatch := false
			for t > 0 && b < len(grid)-1 {
				t--
				b++
				if grid[t] != grid[b] {
					mismatch = true
				}
			}
			if grid[t] == grid[b] && !mismatch {
				return i + 1
			}
		}
	}
	return -1
}

func RotatedGrid(grid []string) []string {
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

func GridValue(grid []string) int {
	value := rowMirror(grid)
	if value != -1 {
		return value * 100
	}

	rotatedValue := rowMirror(RotatedGrid(grid))
	if rotatedValue != -1 {
		return rotatedValue
	}

	panic("unable to find reflection for grid")
}

func Execute() {
	inputRaw := util.ReadLines("day13/input.txt")
	gridsRaw := strings.Split(strings.Join(inputRaw, "\n"), "\n\n")

	sum := 0
	for g := range gridsRaw {
		grid := strings.Fields(gridsRaw[g])
		value := GridValue(grid)
		fmt.Println(value)
		sum += value
	}

	fmt.Println("sum:", sum)
}
