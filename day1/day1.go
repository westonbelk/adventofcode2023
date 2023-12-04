package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/golang/glog"
	"github.com/westonbelk/adventofcode/util"
	"golang.org/x/text/language"
	"golang.org/x/text/search"
)

func Execute() {
	data, err := os.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0

	for _, line := range strings.Fields(string(data)) {
		answer, err := Trebuchet(line)
		if err != nil {
			panic(err)
		}
		sum += answer
	}

	fmt.Println("total sum:", sum)
}

func Trebuchet(input string) (int, error) {
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

	frontDigitValue := string(input[frontDigitPos])
	backDigitValue := string(input[backDigitPos])
	frontSonPos, frontSonValue := spelledOutNumber(input)
	backSonPos, backSonValue := spelledOutNumberEnd(input)

	if frontDigitPos != -1 {
		front = frontDigitValue
	}

	if frontSonPos != -1 && frontSonPos < frontDigitPos {
		front = frontSonValue
	}

	if backDigitPos != -1 {
		back = backDigitValue
	}

	if backSonPos > backDigitPos {
		back = backSonValue
	}

	// debug
	if glog.V(2) {
		for i := 0; i < len(input); i++ {
			iStr := strconv.Itoa(i)
			fmt.Printf("%s", string(iStr[len(iStr)-1]))
		}
		fmt.Println()
		fmt.Println(input)
		fmt.Println("len: ", len(input))
		fmt.Printf("f_digit [%d] | f_son [%d]\n", frontDigitPos, frontSonPos)
		fmt.Printf("b_digit [%d] | b_son [%d]\n", backDigitPos, backSonPos)
		fmt.Printf("front   [%s] | back  [%s]\n", front, back)
		fmt.Println(strconv.Atoi(front + back))
		fmt.Println()
	}

	return strconv.Atoi(front + back)

}

var spelledOut = map[string]string{
	"zero": "0", "one": "1", "two": "2", "three": "3", "four": "4",
	"five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
}

// returns the index of the first character of the first spelled out number
// index -1 if not found, numberValue is something that will panic Atoi
func spelledOutNumber(s string) (int, string) {
	srch := search.New(language.English)
	earliestPos := -1
	numberValue := "x"
	for k, v := range spelledOut {
		pos, _ := srch.IndexString(s, k)

		if pos >= 0 && (pos < earliestPos || earliestPos == -1) {
			earliestPos = pos
			numberValue = v
		}
	}
	return earliestPos, numberValue
}

// returns the index of the first character of the last spelled out number
// index -1 if not found, numberValue is something that will panic Atoi
func spelledOutNumberEnd(s string) (int, string) {
	s = util.ReverseString(s)
	srch := search.New(language.English)
	earliestPos := -1
	numberValue := "x"
	for k, v := range spelledOut {
		_, pos := srch.IndexString(s, util.ReverseString(k))

		if pos >= 0 && (pos < earliestPos || earliestPos == -1) {
			earliestPos = pos
			numberValue = v
		}
	}
	if earliestPos == -1 {
		return -1, numberValue
	}
	return len(s) - earliestPos, numberValue
}
