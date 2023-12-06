package day6

import (
	"fmt"
)

func Execute() {
	// start := time.Now()
	// defer fmt.Println("time:", time.Since(start).Seconds())
	// input := util.ReadLines("day6/input2.txt")
	// input := util.ReadLines("day6/calibration.txt")

	timeLimit := 40828492
	distanceGoal := 233101111101487

	correctSolutions := 0

	for buttonTime := 0; buttonTime <= timeLimit; buttonTime += 1 {
		raceTime := timeLimit - buttonTime
		distance := raceTime * buttonTime
		if distance > distanceGoal {
			correctSolutions++
			// fmt.Printf("Solution: hold button for %d seconds\n", buttonTime)
		}
	}
	fmt.Println()

	fmt.Println("total:", correctSolutions)
}
