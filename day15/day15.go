package day15

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/westonbelk/adventofcode/util"
)

type Box struct {
	Items []*Item
}

func (b *Box) Pop(label string) *Item {
	remainingItems := make([]*Item, 0)
	foundIdx := slices.IndexFunc(b.Items, func(item *Item) bool {
		return item.Label == label
	})

	if foundIdx == -1 {
		return nil
	}
	found := b.Items[foundIdx]

	for i := range b.Items {
		if i != foundIdx {
			remainingItems = append(remainingItems, b.Items[i])
		}
	}
	b.Items = remainingItems
	return found
}

func (b *Box) Push(item Item) {
	for _, v := range b.Items {
		if v.Label == item.Label {
			v.FocalLength = item.FocalLength
			return
		}
	}
	b.Items = append(b.Items, &item)
}

type Item struct {
	Label       string
	FocalLength int
}

func Hash(s string) int {
	value := 0
	for _, r := range s {
		value += int(r)
		value *= 17
		value = value % 256
	}
	return value
}

func operationsFunc(c rune) bool {
	return c == '=' || c == '-'
}

func Execute() {
	input := util.ReadLines("day15/input.txt")[0]
	steps := strings.Split(input, ",")

	// initialize map
	boxes := make(map[int]*Box, 0)
	for i := 0; i < 256; i++ {
		boxes[i] = &Box{Items: make([]*Item, 0)}
	}

	sum := 0
	for _, stepStr := range steps {
		step := strings.FieldsFunc(stepStr, operationsFunc)
		switch len(step) {
		case 1: // '-'
			boxNum := Hash(step[0])
			box := boxes[boxNum]
			box.Pop(step[0])
		case 2: // '='
			boxNum := Hash(step[0])
			box := boxes[boxNum]
			focalLength, err := strconv.Atoi(step[1])
			if err != nil {
				panic(err)
			}
			box.Push(Item{step[0], focalLength})
		default:
			panic(fmt.Sprintln("unable to handle step: ", stepStr))
		}
		// debug
		// fmt.Printf("\nAfter %q:\n", stepStr)
		// for k, v := range boxes {
		// 	if len(v.Items) > 0 {
		// 		fmt.Printf("Box %d: ", k)
		// 		for _, item := range v.Items {
		// 			fmt.Printf("[%v %v]", item.Label, item.FocalLength)
		// 		}
		// 		fmt.Println()
		// 	}
		// }
	}

	// add up all the stuff
	for boxNum, box := range boxes {
		for slotNum, slot := range box.Items {
			sum += (1 + boxNum) * (1 + slotNum) * (slot.FocalLength)
		}
	}

	fmt.Println("Focusing power:", sum)
}
