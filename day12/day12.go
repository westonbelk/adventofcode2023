package day12

import (
	"fmt"
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

type Entry struct {
	Original *Entry
	Pattern  []rune
	Nums     []int
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

	fmt.Printf("generating %d permutations\n", e.Unknown())
	perms := Product([]rune("#."), e.Unknown())
	fmt.Printf("checking %d iterations for %s\n", len(perms), e)
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
	// fmt.Println("Checking:", string(pattern), CountPattern(pattern))
	return slices.Equal(e.Nums, CountPattern(pattern))
}

var PermutationMap = make(map[int][][]rune, 0)

func Execute() {
	// input := util.ReadLines("day12/input.txt")
	input := []string{"???.### 1,1,3"}

	entries := make([]Entry, 0, len(input))
	for _, line := range input {
		split := strings.Fields(line)
		pattern, numsRaw := split[0], split[1]
		nums := util.ReadNums(strings.Split(numsRaw, ","))
		expandedNums := make([]int, 0, len(nums)*5)
		expandedPattern := make([]rune, 0, (len(pattern)*5)+4)

		for i := 0; i < 5; i++ {
			expandedNums = append(expandedNums, nums...)
			expandedPattern = append(expandedPattern, []rune(pattern)...)
			if i != 4 {
				expandedPattern = append(expandedPattern, '?')
			}
		}
		e := Entry{
			Original: &Entry{
				Pattern:  []rune(pattern),
				Nums:     nums,
				Original: nil,
			},
			Pattern: []rune(expandedPattern),
			Nums:    expandedNums,
		}
		fmt.Println(e)
		entries = append(entries, e)
	}

	// check the largest
	largest := 0
	for _, x := range entries {
		if x.Unknown() > largest {
			largest = x.Unknown()
		}
	}
	fmt.Println("largest:", largest)
	// end check the largest

	sum := 0
	for i, e := range entries {
		sum += e.ValidIterations()
		fmt.Printf("processed %d entries out of %d\n", i+1, len(entries))
	}

	fmt.Println("sum:", sum)
}
