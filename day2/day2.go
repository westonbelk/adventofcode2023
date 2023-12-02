package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/westonbelk/adventofcode/util"
)

func Execute() {
	// data := util.ReadLines("day2/calibration1.txt")
	data := util.ReadLines("day2/input.txt")

	sum := 0

	for _, line := range data {
		answer, err := Conundrum(line)
		if err != nil {
			panic(err)
		}
		sum += answer
	}

	fmt.Println("total sum:", sum)
}

type Set struct {
	red   int
	green int
	blue  int
}

var bag = Set{
	red:   12,
	green: 13,
	blue:  14,
}

func Conundrum(input string) (int, error) {
	possible := true
	inputSplit := strings.Split(input, ": ")
	game, setsRaw := inputSplit[0], strings.Split(inputSplit[1], ";")
	sets := []Set{}
	for _, setRaw := range setsRaw {
		sets = append(sets, toSet(setRaw))
		set := toSet(setRaw)
		if !isPossible(set) {
			possible = false
		}
	}

	// debug //
	fmt.Println("game:", game)
	for _, s := range sets {
		fmt.Println(s)
		fmt.Println(possible)
	}
	fmt.Println()
	// debug //

	if possible {
		return strconv.Atoi(strings.TrimPrefix(game, "Game "))
	}
	return 0, nil
}

func isPossible(set Set) bool {
	return set.red <= bag.red && set.green <= bag.green && set.blue <= bag.blue
}

func toSet(set string) Set {
	retSet := Set{0, 0, 0}
	set = strings.Trim(set, " ")
	events := strings.Split(set, ", ")
	for _, eventRaw := range events {
		num, color := parseEvent(eventRaw)
		switch color {
		case "red":
			retSet.red += num
		case "green":
			retSet.green += num
		case "blue":
			retSet.blue += num
		}
	}
	return retSet
}

// output: (number, color)
func parseEvent(rawEvent string) (int, string) {
	splitEvent := strings.Split(rawEvent, " ")
	numStr, color := splitEvent[0], splitEvent[1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num, color
}
