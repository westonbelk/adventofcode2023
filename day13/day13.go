package day13

import (
	"fmt"
	"slices"
	"strings"

	"github.com/westonbelk/adventofcode/util"
)

// int  => -1 if not found
// otherwise index = value
func rowMirror(grid []string, ignore int) int {
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
				if ignore == (i + 1) {
					continue
				}
				return i + 1
			}
		}
	}
	return -1
}

func flipBit(s string, idx int) string {
	res := []rune(s)
	if res[idx] == '.' {
		res[idx] = '#'
	} else if res[idx] == '#' {
		res[idx] = '.'
	} else {
		panic("can't flip bit")
	}
	return string(res)
}

func flipBitGrid(originalGrid []string, x, y int) []string {
	grid := slices.Clone(originalGrid)
	row := grid[y]
	row = flipBit(row, x)
	grid[y] = row
	return grid
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

func GridValue(grid []string, ignore int) int {
	value := rowMirror(grid, ignore/100)
	if value != -1 {
		return value * 100
	}

	rotatedValue := rowMirror(RotatedGrid(grid), ignore)
	if rotatedValue != -1 {
		return rotatedValue
	}

	return -1
	fmt.Println(strings.Join(grid, "\n"))
	panic("unable to find reflection for grid")
}

func Execute() {
	inputRaw := util.ReadLines("day13/input.txt")
	gridsRaw := strings.Split(strings.Join(inputRaw, "\n"), "\n\n")

	sum := 0
	for g := range gridsRaw {
		grid := strings.Fields(gridsRaw[g])
		originalValue := GridValue(grid, -1)
		value := -1

	Check:
		for y := range grid {
			for x := range grid[y] {
				flippedBitGrid := flipBitGrid(grid, x, y)
				v := GridValue(flippedBitGrid, originalValue)
				if v != -1 {
					value = v
					break Check
				}
			}
		}
		if value == -1 {
			fmt.Println(strings.Join(grid, "\n"))
			panic("unable to determine new grid reflection value")
		}
		sum += value
	}

	fmt.Println("sum:", sum)
}
