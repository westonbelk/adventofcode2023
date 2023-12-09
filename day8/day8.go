package day8

import (
	"fmt"
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

	var count int
	found := false
	node1 := lookup["MJA"] // start node
	node2 := lookup["RGA"]
	node3 := lookup["JMA"]
	node4 := lookup["XHA"]
	node5 := lookup["DQA"]
	node6 := lookup["AAA"]
	// node1 := lookup["11A"] // start node
	// node2 := lookup["11A"]
	// node3 := lookup["11A"]
	// node4 := lookup["22A"]
	// node5 := lookup["22A"]
	// node6 := lookup["22A"]

	for count = 0; !found; count++ {
		i := (count) % len(steps)
		r := steps[i]
		switch r {
		case 'R':
			// fmt.Printf("%s [%s] => %s\n", node1.Value, string(r), node1.Right.Value)
			// fmt.Printf("%s [%s] => %s\n", node4.Value, string(r), node4.Right.Value)
			node1 = node1.Right
			node2 = node2.Right
			node3 = node3.Right
			node4 = node4.Right
			node5 = node5.Right
			node6 = node6.Right
		case 'L':
			// fmt.Printf("%s [%s] => %s\n", node1.Value, string(r), node1.Left.Value)
			// fmt.Printf("%s [%s] => %s\n", node4.Value, string(r), node4.Left.Value)
			node1 = node1.Left
			node2 = node2.Left
			node3 = node3.Left
			node4 = node4.Left
			node5 = node5.Left
			node6 = node6.Left
		default:
			panic(fmt.Sprintf("unhandled instruction: %v", r))
		}
		// fmt.Printf("%v", node1.Value[2] == 'Z')
		// fmt.Println()
		if node1.Value[2] == 'Z' && node2.Value[2] == 'Z' && node3.Value[2] == 'Z' && node4.Value[2] == 'Z' && node5.Value[2] == 'Z' && node6.Value[2] == 'Z' {
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
