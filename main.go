package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/westonbelk/adventofcode/day1"
	"github.com/westonbelk/adventofcode/day2"
	"github.com/westonbelk/adventofcode/day3"
	"github.com/westonbelk/adventofcode/day4"
	"github.com/westonbelk/adventofcode/day5"
	"github.com/westonbelk/adventofcode/day6"
	"github.com/westonbelk/adventofcode/day7"
	"github.com/westonbelk/adventofcode/day8"
	"github.com/westonbelk/adventofcode/day9"
)

var day int

func main() {
	flag.Set("logtostderr", "true")
	flag.Set("v", "1")
	flag.IntVar(&day, "day", 0, "day of advent of code to run")
	flag.Parse()

	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	today := time.Now().In(est).Day()

	if day == 0 {
		day = today
	}

	dayFuncMap := map[int]interface{}{
		1: day1.Execute,
		2: day2.Execute,
		3: day3.Execute,
		4: day4.Execute,
		5: day5.Execute,
		6: day6.Execute,
		7: day7.Execute,
		8: day8.Execute,
		9: day9.Execute,
	}

	f, ok := dayFuncMap[day].(func())
	if !ok {
		panic(fmt.Sprintln("day not found:", day))
	}

	fmt.Printf("Running day %d:\n", day)
	f()
}
