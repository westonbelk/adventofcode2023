package day3

import (
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/westonbelk/adventofcode/util"
)

func Execute() {
	// input, err := os.ReadFile("day3/calibration1.txt")
	input, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	symbolGrid := util.Grid(input, func(r rune) bool {
		return !unicode.IsDigit(r) && r != '.'
	})
	parts := make(map[image.Point][]int)

	for y, s := range strings.Fields(string(input)) {

		// for each line, parse the numbers on the line, find the adjacent bounds,
		// and check if
		matches := regexp.MustCompile(`\d+`).FindAllStringIndex(s, -1)
		for _, m := range matches {
			n, err := strconv.Atoi(s[m[0]:m[1]])
			if err != nil {
				panic(err)
			}
			bounds := util.AdjacentToRect(image.Rectangle{
				Min: image.Point{m[0], y},
				Max: image.Point{m[1] - 1, y},
			})

			for _, p := range bounds {
				if _, ok := symbolGrid[p]; ok {
					parts[p] = append(parts[p], n)
				}
			}
		}
	}

	part1 := 0
	for _, v := range parts {
		for _, n := range v {
			part1 += n
		}
	}
	fmt.Println("part1: ", part1)

	part2 := 0
	for _, v := range parts {
		if len(v) == 2 {
			part2 += v[0] * v[1]
		}
	}
	fmt.Println("part2: ", part2)
}
