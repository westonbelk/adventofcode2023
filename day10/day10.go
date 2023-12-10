package day10

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

var input []string

var Red = color.RGBA64{65535, 0, 0, 65535}

func Execute() {
	inputBytes, err := os.ReadFile("day10/input.txt")
	// inputBytes, err := os.ReadFile("day10/calibration4.txt")
	if err != nil {
		fmt.Println("error reading file: exiting")
		os.Exit(1)
	}
	input = strings.Fields(string(inputBytes))

	startLoc := findS(input)
	player := Player{Location: startLoc, Direction: Down}
	steps := 0
	bounds := image.Rect(0, 0, len(input[0]), len(input))
	img := image.NewRGBA64(bounds)

	for steps == 0 || player.Location != startLoc {
		img.SetRGBA64(player.Location.X, player.Location.Y, Red)
		player.Advance()
		steps++
		fmt.Println()
		if player.Location == startLoc {
			break
		}
	}

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
