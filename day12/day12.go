package day12

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/westonbelk/adventofcode/util"
)

func fill[T any](slice []T, val T) {
	for i := range slice {
		slice[i] = val
	}
}

func Product(a []rune, k int) [][]rune {
	indexes := make([]int, k)
	var ps [][]rune

	for indexes != nil {
		p := make([]rune, k)
		for i, x := range indexes {
			p[i] = a[x]
		}

		for i := len(indexes) - 1; i >= 0; i-- {
			indexes[i]++
			if indexes[i] < len(a) {
				break
			}
			indexes[i] = 0
			if i <= 0 {
				indexes = nil
				break
			}
		}
		ps = append(ps, p)
	}
	return ps
}

func TestProduct() {
	s := Product([]rune("#."), 18)

	// for _, s := range o {
	// 	fmt.Println(string(s))
	// }
	if 1 == 1 {
		os.Exit(1)
	}
}

func Hashify(strLen int) [][]rune {
	res := make([][]rune, 0)
	starter := make([]rune, strLen)
	fill(starter, '.')
	res = append(res, starter)

	for n := 1; n <= len(starter); n++ {
		for pos := 0; pos < len(starter); pos++ {
			if pos+n > len(starter) {
				continue
			}
			s := slices.Clone(starter)
			for i := 0; i < n; i++ {
				s[pos+i] = '#'
			}
			res = append(res, s)
		}
	}
	return res
}

type Entry struct {
	Pattern []rune
	Nums    []int
}

func (e Entry) String() string {
	s := make([]string, len(e.Nums))
	for i := range e.Nums {
		s[i] = strconv.Itoa(e.Nums[i])
	}
	return fmt.Sprintf("%s %s", string(e.Pattern), strings.Join(s, ","))
}

func (e *Entry) Unknown() int {
	return strings.Count(string(e.Pattern), "?")
}

func (e *Entry) ValidIterations() int {
	total := 0

	perms := PermutationMap[e.Unknown()]
	for _, p := range perms {
		if e.CheckIteration(p) {
			total++
		}
	}

	return total
}

func CountPattern(pattern []rune) []int {
	res := make([]int, 0)
	count := 0
	for _, p := range pattern {
		if p == '#' {
			count++
		} else {
			if count > 0 {
				res = append(res, count)
			}
			count = 0
		}
	}
	if count > 0 {
		res = append(res, count)
	}
	return res
}

func (e *Entry) CheckIteration(replacements []rune) bool {
	unknownIdx := 0
	pattern := slices.Clone(e.Pattern)
	for i := range pattern {
		if pattern[i] == '?' {
			pattern[i] = replacements[unknownIdx]
			unknownIdx++
		}
	}

	// check if pattern is valid
	fmt.Println("Checking:", string(pattern), CountPattern(pattern))
	return slices.Equal(e.Nums, CountPattern(pattern))
}

var PermutationMap = make(map[int][][]rune, 0)

func Execute() {
	// input := util.ReadLines("day12/calibration.txt")
	input := []string{"???.### 1,1,3"}
	TestProduct()

	entries := make([]Entry, 0, len(input))
	uniqueUnknown := make(map[int]struct{}, 0)
	for _, line := range input {
		split := strings.Fields(line)
		pattern, numsRaw := split[0], split[1]
		nums := util.ReadNums(strings.Split(numsRaw, ","))
		e := Entry{
			Pattern: []rune(pattern),
			Nums:    nums,
		}

		entries = append(entries, e)
		uniqueUnknown[e.Unknown()] = struct{}{}
	}

	for k := range uniqueUnknown {
		PermutationMap[k] = Hashify(k)
	}

	sum := 0
	for _, e := range entries {
		sum += e.ValidIterations()
	}

	fmt.Println("sum:", sum)
}
