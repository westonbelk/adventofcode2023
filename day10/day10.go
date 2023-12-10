package day10

import (
	"fmt"
	"image"
	"os"
	"strings"
)

var input []string

func Execute() {
	inputBytes, err := os.ReadFile("day10/input.txt")
	// inputBytes, err := os.ReadFile("day10/calibration2.txt")
	if err != nil {
		fmt.Println("error reading file: exiting")
		os.Exit(1)
	}
	input = strings.Fields(string(inputBytes))

	startLoc := findS(input)
	player := Player{Location: startLoc, Direction: Down}
	steps := 0

	for steps == 0 || player.Location != startLoc {
		player.Advance()
		steps++
		fmt.Println()
		if player.Location == startLoc {
			break
		}
	}

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
