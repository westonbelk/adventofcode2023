package day10

import (
	"fmt"
	"image"
)

var Up = image.Point{0, -1}
var Left = image.Point{-1, 0}
var Down = image.Point{0, 1}
var Right = image.Point{1, 0}

type Player struct {
	Direction image.Point
	Location  image.Point
}

// advance one step in the direction the player is facing
// interpret the new tile underfoot and turn as necessary
func (p *Player) Advance() {
	fmt.Printf("Current direction: %v\n", p.Direction)
	prevLoc := p.Location
	p.Location = p.Location.Add(p.Direction)
	fmt.Printf("Advanced from %v to %v\n", prevLoc, p.Location)

	underFoot := rune(input[p.Location.Y][p.Location.X])

	switch underFoot {
	case '|':
		fmt.Printf("Detected %q. Continuing forward.\n", underFoot)
	case 'L':
		fmt.Printf("Detected %q. Turning.\n", underFoot)
		switch p.Direction {
		case Left:
			p.Direction = Up
		case Down:
			p.Direction = Right
		default:
			fmt.Printf("encountered error with %q\n", underFoot)
		}
	case 'J':
		fmt.Printf("Detected %q. Turning.\n", underFoot)
		switch p.Direction {
		case Down:
			p.Direction = Left
		case Right:
			p.Direction = Up
		default:
			fmt.Printf("encountered error with %q", underFoot)
		}
	case '7':
		fmt.Printf("Detected %q. Turning.\n", underFoot)
		switch p.Direction {
		case Up:
			p.Direction = Left
		case Right:
			p.Direction = Down
		default:
			fmt.Printf("encountered error with %q\n", underFoot)
		}
	case 'F':
		fmt.Printf("Detected %q. Turning.\n", underFoot)
		switch p.Direction {
		case Up:
			p.Direction = Right
		case Left:
			p.Direction = Down
		default:
			fmt.Printf("encountered error with %q\n", underFoot)
		}
	default:
		fmt.Printf("unable to parse underfoot: %q\n", underFoot)
	}

}

func (p *Player) TurnLeft() {
	switch p.Direction {
	case Up:
		p.Direction = Left
	case Left:
		p.Direction = Down
	case Down:
		p.Direction = Right
	case Right:
		p.Direction = Up
	default:
		panic(fmt.Errorf("unable to figure out direction: %v", p.Direction))
	}
}

func (p *Player) TurnRight() {
	switch p.Direction {
	case Up:
		p.Direction = Right
	case Left:
		p.Direction = Down
	case Down:
		p.Direction = Left
	case Right:
		p.Direction = Up
	default:
		panic(fmt.Errorf("unable to figure out direction: %v", p.Direction))
	}
}
