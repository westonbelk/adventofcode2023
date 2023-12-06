package day6

import (
	"fmt"
	"strings"
	"time"

	"github.com/westonbelk/adventofcode/util"
)

func Execute() {
	start := time.Now()
	input := util.ReadLines("day6/input2.txt")
	// input := util.ReadLines("day6/calibration.txt")

	times := util.ReadNums(strings.Fields(strings.Split(input[0], ":")[1]))
	distances := util.ReadNums(strings.Fields(strings.Split(input[1], ":")[1]))

	correctSolutions := make([]int, len(times))

	for i := range distances {
		timeLimit := times[i]
		distanceGoal := distances[i]

		for buttonTime := 0; buttonTime <= timeLimit; buttonTime++ {
			raceTime := timeLimit - buttonTime
			speed := buttonTime
			distance := raceTime * speed
			if distance > distanceGoal {
				correctSolutions[i] = correctSolutions[i] + 1
			}
		}
		fmt.Println()
	}

	fmt.Println(correctSolutions)
	total := 1
	for _, solution := range correctSolutions {
		total *= solution
	}

	fmt.Println("time:", time.Since(start).Seconds())
	fmt.Println("total:", total)
}
