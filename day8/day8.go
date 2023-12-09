package day8

import (
	"fmt"
	"strings"

	"github.com/westonbelk/adventofcode/util"
)

func Execute() {
	input := util.ReadLines("day8/input.txt")
	// input := util.ReadLines("day8/calibration2.txt")

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

	var count int
	found := false
	node := lookup["AAA"] // start node
	endNode := lookup["ZZZ"]
	for count = 0; !found; count++ {
		i := (count) % len(steps)
		r := steps[i]
		fmt.Printf("%s [%s] => ", node.Value, string(r))
		switch r {
		case 'R':
			fmt.Printf("%s\n", node.Right.Value)
			node = node.Right
		case 'L':
			fmt.Printf("%s\n", node.Left.Value)
			node = node.Left
		default:
			panic(fmt.Sprintf("unhandled instruction: %v", r))
		}
		if node == endNode {
			found = true
		}
	}

	fmt.Println("number of steps: ", count)
}

type Node struct {
	Line  int
	Value string
	Left  *Node
	Right *Node
}
