package day12

import (
	"fmt"
	"slices"
	"strings"
)

func fill[T any](slice []T, val T) {
	for i := range slice {
		slice[i] = val
	}
}

func Hashify(strLen int) []string {
	res := make([]string, 0)
	starter := make([]rune, strLen)
	fill(starter, '.')
	res = append(res, string(starter))

	for n := 1; n <= len(starter); n++ {
		for pos := 0; pos < len(starter); pos++ {
			if pos+n > len(starter) {
				continue
			}
			s := slices.Clone(starter)
			for i := 0; i < n; i++ {
				s[pos+i] = '#'
			}
			res = append(res, string(s))
		}
	}

	return res
}

func Execute() {
	//inputRaw := util.ReadLines("day12/input.txt")

	fmt.Println(strings.Join(Hashify(30), "\n"))
}
