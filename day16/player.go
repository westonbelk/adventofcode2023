package day16

import (
	"fmt"
	"image"
)

var Up = image.Point{0, -1}
var Left = image.Point{-1, 0}
var Down = image.Point{0, 1}
var Right = image.Point{1, 0}

var Directions = [4]image.Point{Up, Left, Down, Right}

type Player struct {
	Direction image.Point
	Location  image.Point
}

func (p *Player) InBounds() bool {
	x := p.Location.X >= 0 && p.Location.X < len(input[0])
	y := p.Location.Y >= 0 && p.Location.Y < len(input)
	return x && y
}

// advance one step in the direction the player is facing
// interpret the new tile underfoot and turn as necessary
// modifies the current player and returns a new player or nil
func (p *Player) Advance() *Player {
	var newPlayer *Player
	// fmt.Printf("Current direction: %v\n", p.Direction)
	// prevLoc := p.Location
	p.Location = p.Location.Add(p.Direction)
	// fmt.Printf("Advanced from %v to %v\n", prevLoc, p.Location)
	if p.InBounds() {
		travelReplaced := []rune(traveled[p.Location.Y])
		travelReplaced[p.Location.X] = '#'
		traveled[p.Location.Y] = string(travelReplaced)

		underFoot := rune(input[p.Location.Y][p.Location.X])

		switch underFoot {
		case '|':
			switch p.Direction {
			case Left:
				p.Direction = Up
				newPlayer = &Player{
					Location:  p.Location,
					Direction: Down,
				}
			case Right:
				p.Direction = Up
				newPlayer = &Player{
					Location:  p.Location,
					Direction: Down,
				}
			}
			// fmt.Printf("Detected %q. Continuing forward.\n", underFoot)
		case '-':
			// fmt.Printf("Detected %q. Turning.\n", underFoot)
			switch p.Direction {
			case Up:
				p.Direction = Left
				newPlayer = &Player{
					Location:  p.Location,
					Direction: Right,
				}
			case Down:
				p.Direction = Left
				newPlayer = &Player{
					Location:  p.Location,
					Direction: Right,
				}
			default:
				// fmt.Printf("encountered error with %q\n", underFoot)
			}
		case '\\':
			// fmt.Printf("Detected %q. Turning.\n", underFoot)
			switch p.Direction {
			case Left:
				p.Direction = Up
			case Up:
				p.Direction = Left
			case Right:
				p.Direction = Down
			case Down:
				p.Direction = Right
			default:
				// fmt.Printf("encountered error with %q\n", underFoot)
			}
		case '/':
			// fmt.Printf("Detected %q. Turning.\n", underFoot)
			switch p.Direction {
			case Down:
				p.Direction = Left
			case Right:
				p.Direction = Up
			case Up:
				p.Direction = Right
			case Left:
				p.Direction = Down
			default:
				// fmt.Printf("encountered error with %q", underFoot)
			}
		case '.':
		default:
			// fmt.Printf("unable to parse underfoot: %q\n", underFoot)
		}

	}
	return newPlayer
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
