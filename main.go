package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/westonbelk/adventofcode/util"
	"golang.org/x/text/language"
	"golang.org/x/text/search"
)

func main() {
	data := util.ReadLines("input/day1.txt")
	sum := 0

	for _, line := range data {
		answer, err := day1(line)
		if err != nil {
			panic(err)
		}
		sum += answer
	}

	fmt.Println("total sum:", sum)
}

func day1(input string) (int, error) {
	front, back := "", ""
	frontDigitPos, backDigitPos := -1, -1

	for i := 0; i < len(input); i++ {
		c := rune(input[i])
		if unicode.IsDigit(c) {
			frontDigitPos = i
			break
		}
	}

	for i := len(input) - 1; i >= 0; i-- {
		c := rune(input[i])
		if unicode.IsDigit(c) {
			backDigitPos = i
			break
		}
	}

	frontSonPos, frontSonValue := spelledOutNumber(input)
	backSonPos, backSonValue := spelledOutNumberEnd(input)
	backSonPos = len(input) - backSonPos

	if frontDigitPos != -1 {
		front = string(input[frontDigitPos])
	}

	if frontSonPos != -1 && (frontSonPos < frontDigitPos || frontDigitPos == -1) {
		front = frontSonValue
	}

	if backDigitPos != -1 {
		back = string(input[backDigitPos])
	}

	if backSonPos > backDigitPos && backSonPos != len(input)+1 {
		back = backSonValue
	}

	fmt.Println(input)
	fmt.Println("len: ", len(input))
	fmt.Println(frontDigitPos, frontSonPos)
	fmt.Println(backDigitPos, backSonPos)
	fmt.Println(front, back)
	fmt.Println(strconv.Atoi(front + back))
	fmt.Println()

	return strconv.Atoi(front + back)

}

var spelledOut = map[string]string{
	"zero": "0", "one": "1", "two": "2", "three": "3", "four": "4",
	"five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
}

// returns the starting or ending index of the first spelled out number
// index found at or -1, number value as a string
func spelledOutNumberEnd(s string) (int, string) {
	s = util.Reverse(s)
	srch := search.New(language.English)
	earliestPos := -1
	numberValue := "-1"
	for k, v := range spelledOut {
		_, pos := srch.IndexString(s, util.Reverse(k))

		if pos >= 0 && (pos < earliestPos || earliestPos == -1) {
			earliestPos = pos
			numberValue = v
		}
	}
	return earliestPos, numberValue
}

// returns the starting or ending index of the first spelled out number
// index found at or -1, number value as a string
func spelledOutNumber(s string) (int, string) {
	srch := search.New(language.English)
	earliestPos := -1
	numberValue := "-1"
	for k, v := range spelledOut {
		pos, _ := srch.IndexString(s, k)

		if pos >= 0 && (pos < earliestPos || earliestPos == -1) {
			earliestPos = pos
			numberValue = v
		}
	}
	return earliestPos, numberValue
}
