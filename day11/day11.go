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

// @todo instead of allocating a bunch of memroy we should just determine the indexes of where to
// expand and add it to the coords...
func Expand(in []string) []string {
	arr := slices.Clone(in)
	for i := len(arr) - 1; i >= 0; i-- {
		allStars := !strings.ContainsFunc(arr[i], func(r rune) bool {
			return r != '.'
		})
		if allStars {
			for g := 0; g < Million; g++ {
				arr = slices.Insert(arr, i, arr[i])
			}
		}
	}
	fmt.Println("errrrrr")

	// rotated := make([]string, 0)
	allStarColumns := make([]int, 0)
	for x := range arr[0] {
		// row := make([]byte, len(arr))
		allStars := true
		for y := range arr {
			if arr[y][x] != '.' {
				allStars = false
			}
			// row = append(row, arr[y][x])
		}
		if allStars {
			allStarColumns = append(allStarColumns, x)
		}
		// rotated = append(rotated, string(row))
	}
	// fmt.Println(rotated)
	fmt.Println()
	slices.Reverse(allStarColumns)

	for y := range arr {
		for _, x := range allStarColumns { // sort reverse a range of the index where it needs to be inserted
			row := []byte(arr[y])
			for g := 0; g < Million; g++ {
				row = slices.Insert(row, x, byte('.'))
			}
			arr[y] = string(row)
		}
		fmt.Println("yeesh")
	}

	return arr
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
	input := util.ReadLines("day11/calibration.txt")

	expanded := Expand(input)

	points := FindPoints(expanded, '#')

	fmt.Println("expanded")
	fmt.Println(strings.Join(expanded, "\n"))
	for i, p := range points {
		fmt.Println(i, "=>", p)
	}

	total := 0
	for _, p := range Permuatations(len(points)) {
		total += StepsBetween(points[p.X], points[p.Y])
	}

	fmt.Println("total", total)
}
