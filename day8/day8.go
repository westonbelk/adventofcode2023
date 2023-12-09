package day8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/westonbelk/adventofcode/util"
)

func Execute() {
	input := util.ReadLines("day8/input.txt")
	// input := util.ReadLines("day8/calibration3.txt")

	lookup := make(map[string]*Node)

	steps := input[0]

	// create blank nodes
	for i, s := range input[2:] {
		value := strings.Fields(s)[0]
		lookup[value] = &Node{Value: value, Line: i + 2, Left: nil, Right: nil}
	}

	// fill in
	for _, s := range input[2:] {
		fields := strings.Fields(s)
		leftValue := strings.TrimRight(strings.TrimLeft(fields[2], "("), ",")
		rightValue := strings.TrimRight(fields[3], ")")
		node := lookup[fields[0]]
		node.Left = lookup[leftValue]
		node.Right = lookup[rightValue]
	}

	for _, v := range lookup {
		fmt.Printf("%s = (%s, %s)\n", v.Value, v.Left.Value, v.Right.Value)
	}
	fmt.Println(len(lookup))
	fmt.Println("proceeding with walking")
	fmt.Println()

	nodes := []*Node{
		lookup["MJA"],
		lookup["RGA"],
		lookup["JMA"],
		lookup["XHA"],
		lookup["DQA"],
		lookup["AAA"],
	}

	answers := make([]string, len(nodes))
	for i, node := range nodes {
		answers[i] = strconv.Itoa(countSteps(node, steps))
	}
	fmt.Printf("LCM(%s)\n", strings.Join(answers, ", "))

}

func countSteps(node *Node, steps string) int {
	var count int
	found := false
	for count = 0; !found; count++ {
		i := (count) % len(steps)
		r := steps[i]
		switch r {
		case 'R':
			// fmt.Printf("%s [%s] => %s\n", node.Value, string(r), node.Right.Value)
			node = node.Right
		case 'L':
			// fmt.Printf("%s [%s] => %s\n", node.Value, string(r), node.Left.Value)
			node = node.Left
		default:
			panic(fmt.Sprintf("unhandled instruction: %v", r))
		}
		// fmt.Printf("%v", node.Value[2] == 'Z')
		// fmt.Println()
		if node.Value[2] == 'Z' {
			found = true
		}
	}
	return count
}

type Node struct {
	Line  int
	Value string
	Left  *Node
	Right *Node
}
