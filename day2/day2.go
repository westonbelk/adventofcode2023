package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/glog"
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

// var bag = Set{
// 	red:   12,
// 	green: 13,
// 	blue:  14,
// }

// func isPossible(set Set) bool {
// 	return set.red <= bag.red && set.green <= bag.green && set.blue <= bag.blue
// }

func (s Set) power() int {
	return s.red * s.green * s.blue
}

func Conundrum(input string) (int, error) {
	possible := true
	inputSplit := strings.Split(input, ": ")
	game, setsRaw := inputSplit[0], strings.Split(inputSplit[1], ";")
	sets := []Set{}
	for _, setRaw := range setsRaw {
		sets = append(sets, toSet(setRaw))
	}

	wendysBiggieBag := makeBiggieBag(sets)

	// debug //
	if glog.V(2) {
		fmt.Println("game:", game)
		for _, s := range sets {
			fmt.Println(s)
			fmt.Println(possible)
		}
		fmt.Println(wendysBiggieBag)
		fmt.Println()
	}
	// debug //

	return wendysBiggieBag.power(), nil
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

func makeBiggieBag(sets []Set) Set {
	biggieBag := Set{0, 0, 0}
	for _, s := range sets {
		biggieBag.red = max(biggieBag.red, s.red)
		biggieBag.green = max(biggieBag.green, s.green)
		biggieBag.blue = max(biggieBag.blue, s.blue)
	}
	return biggieBag
}

func max(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}
