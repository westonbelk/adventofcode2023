package day16

import (
	"fmt"
	"image"
	"slices"
	"strings"

	"github.com/westonbelk/adventofcode/util"
)

var input []string
var traveled []string

func Execute() {
	input = util.ReadLines("day16/calibration.txt")
	traveled = slices.Clone(input)

	players := make([]*Player, 0)
	players = append(players, &Player{
		Location:  image.Point{0, 0},
		Direction: Right,
	})
	originalPlayer := players[0]

	for len(players) > 0 {
		if players[0] == originalPlayer {
			fmt.Println(originalPlayer)
		}
		afterAdvance := make([]*Player, 0)
		for _, p := range players {
			newPlayer := p.Advance()
			if p.InBounds() {
				fmt.Println(p.Location)
				afterAdvance = append(afterAdvance, p)
			}
			if newPlayer != nil && newPlayer.InBounds() {
				afterAdvance = append(afterAdvance, newPlayer)
			}
		}
		players = afterAdvance
	}

	strings.Join(traveled, "\n")
}
