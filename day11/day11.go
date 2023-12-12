package day11

import (
	"fmt"
	"image"
	"slices"
	"sort"
	"strings"

	xmaps "golang.org/x/exp/maps"

	"github.com/westonbelk/adventofcode/util"
)

// var input []string

const (
	Million int = 1000000 - 1
	Hundred int = 100 - 1
	Ten     int = 10 - 1
)

func ExpandLocations(grid []string) ([]int, []int) {
	colExpansions := make([]int, 0)
	rowExpansions := make([]int, 0)

	for y := len(grid) - 1; y >= 0; y-- {
		allStars := !strings.ContainsFunc(grid[y], func(r rune) bool {
			return r != '.'
		})
		if allStars {
			rowExpansions = append(rowExpansions, y)
		}
	}

	for x := range grid[0] {
		allStars := true
		for y := range grid {
			if grid[y][x] != '.' {
				allStars = false
			}
		}
		if allStars {
			colExpansions = append(colExpansions, x)
		}
	}
	slices.Sort(colExpansions)
	slices.Sort(rowExpansions)
	slices.Reverse(colExpansions)
	slices.Reverse(rowExpansions)
	return rowExpansions, colExpansions
}

func FindPoints(grid []string, target rune) []image.Point {
	count := 1
	points := make([]image.Point, 0)
	for y, s := range grid {
		for x, r := range s {
			if r == target {
				points = append(points, image.Point{x, y})
				count++
			}
		}
	}
	return points
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func StepsBetween(start, end image.Point) int {
	d := start.Sub(end)
	d.X, d.Y = Abs(d.X), Abs(d.Y)
	fmt.Println(d)
	return d.X + d.Y
}

// generates all pairs of 0-(n-1)
func Permuatations(n int) []image.Point {
	ret := make(map[image.Point]struct{})
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			lower := i
			higher := j
			if lower > higher {
				lower = j
				higher = i
			}
			if lower == higher {
				continue
			}
			ret[image.Point{lower, higher}] = struct{}{}
		}
	}

	retS := xmaps.Keys(ret)
	sort.Slice(retS, func(i, j int) bool {
		if retS[i].X == retS[j].X {
			return retS[i].Y < retS[j].Y
		}
		return retS[i].X < retS[j].X
	})
	return retS
}

func Execute() {
	input := util.ReadLines("day11/input.txt")

	extraRows, extraCols := ExpandLocations(input)
	points := FindPoints(input, '#')

	// perform expansion of the universe
	for i := range points {
		p := &points[i]
		for _, ex := range extraRows {
			if ex < p.Y {
				p.Y += Million
			}
		}

		for _, ex := range extraCols {
			if ex < p.X {
				p.X += Million
			}
		}
	}

	fmt.Println()
	for i, p := range points {
		fmt.Println(i, "=>", p)
	}

	total := 0
	for _, p := range Permuatations(len(points)) {
		total += StepsBetween(points[p.X], points[p.Y])
	}

	fmt.Println("total", total)
}
