package day10

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"slices"
	"strings"

	xmaps "golang.org/x/exp/maps"
)

var input []string

var Red = color.RGBA64{65535, 0, 0, 65535}
var Blue = color.RGBA64{0, 0, 65535, 65535}
var Green = color.RGBA64{0, 65535, 0, 65535}

func fill[T any](slice []T, val T) {
	for i := range slice {
		slice[i] = val
	}
}

func AddSpaceBetweenInput() {
	for y := range input {
		for x := len(input[y]) - 1; x > 0; x-- {
			row := []byte(input[y])
			rowModified := slices.Insert(row, x, byte('-'))
			input[y] = string(rowModified)
		}
	}

	for y := len(input) - 1; y > 0; y-- {
		row := make([]byte, len(input[y]))
		fill(row, '|')
		input = slices.Insert(input, y, string(row))
	}
}

// returns (bounds []image.Point, steps int)
func FollowPath(startDirection image.Point) ([]image.Point, int) {
	startLoc := findS(input)
	player := Player{Location: startLoc, Direction: startDirection}
	steps := 0
	bounds := make([]image.Point, 0)

	for steps == 0 || player.Location != startLoc {
		bounds = append(bounds, image.Point{X: player.Location.X, Y: player.Location.Y})
		player.Advance()
		steps++
		if player.Location == startLoc {
			return bounds, steps
		}
	}
	panic("reached end of FollowPath")
}

func ReplaceNonBounds(bounds []image.Point) {
	for y := range input {
		for x := range input[y] {
			if !slices.Contains(bounds, image.Point{x, y}) {
				l := []rune(input[y])
				l[x] = '.'
				input[y] = string(l)
			}
		}
	}
}

func CountInside(innerPoints []image.Point, bounds []image.Point) int {
	// convert bounds to map
	fence := make(map[image.Point]struct{}, 0)
	for _, b := range bounds {
		fence[b] = struct{}{}
	}

	count := 0

	for _, p := range innerPoints {
		_, tl := fence[p]
		_, tr := fence[p.Add(Right)]
		_, bl := fence[p.Add(Down)]
		_, br := fence[p.Add(Right).Add(Down)]
		if !tl && !tr && !bl && !br {
			count++
		}
	}

	return count
}

func Flood(img *image.RGBA64, bounds []image.Point, fillStartDir image.Point, fillColor color.RGBA64) []image.Point {
	// convert bounds to map
	fence := make(map[image.Point]struct{}, 0)
	for _, b := range bounds {
		fence[b] = struct{}{}
	}

	infected := make(map[image.Point]struct{}, 0)
	firstBlood := findS(input).Add(fillStartDir)
	img.SetRGBA64(firstBlood.X, firstBlood.Y, fillColor)
	infected[firstBlood] = struct{}{}
	replaced := true
	for replaced {
		replaced = false
		for p := range infected {
			for _, direction := range Directions {
				patient := p.Add(direction)
				_, alreadyInfected := infected[patient]
				_, onFence := fence[patient]
				if !onFence && !alreadyInfected {
					img.SetRGBA64(patient.X, patient.Y, fillColor)
					infected[patient] = struct{}{}
					replaced = true
				}
			}
		}

	}
	return xmaps.Keys(infected)
}

func Execute() {
	inputBytes, err := os.ReadFile("day10/input.txt")
	// inputBytes, err := os.ReadFile("day10/calibration5.txt")
	if err != nil {
		fmt.Println("error reading file: exiting")
		os.Exit(1)
	}
	input = strings.Fields(string(inputBytes))

	fmt.Println(strings.Join(input, "\n"))
	fmt.Println()
	fmt.Println("doing some stuff")
	fmt.Println()

	preExpandedBounds, _ := FollowPath(Down)
	ReplaceNonBounds(preExpandedBounds)
	fmt.Println(strings.Join(input, "\n"))

	// fmt.Println(strings.Join(input, "\n"))
	// os.Exit(1)

	AddSpaceBetweenInput()
	img := image.NewRGBA64(image.Rect(0, 0, len(input[0]), len(input)))
	expandedBounds, steps := FollowPath(Down)

	for _, p := range expandedBounds {
		img.SetRGBA64(p.X, p.Y, Red)
	}
	s := findS(input)
	img.SetRGBA64(s.X, s.Y, Blue)

	// fill
	innerPoints := Flood(img, expandedBounds, Left, Green)
	area := CountInside(innerPoints, expandedBounds)

	// draw image
	f, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)

	// part1 output
	fmt.Println(strings.Join(input, "\n"))
	fmt.Println("steps:", steps/2)
	fmt.Println("area:", area/4)
}

func findS(input []string) image.Point {
	for i := range input {
		sPosX := strings.IndexRune(input[i], 'S')
		if sPosX != -1 {
			return image.Point{sPosX, i}
		}
	}
	panic("didn't find S")
}
