package day3

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/westonbelk/adventofcode/util"
)

type TokenType int

const (
	Digit TokenType = iota
	Symbol
	Space
	Adjacent
)

var gearCounter = 1
var infectedValues []*InfectedValue

type InfectedValue struct {
	orig        string
	value       int
	digitGearID int
}

type Pos struct {
	tokenType    TokenType
	value        string
	symbolGearID int
	digitGearID  int
}

func (p Pos) String() string {
	switch p.tokenType {
	case Digit:
		return "D"
	case Space:
		return "."
	case Adjacent:
		if p.digitGearID > 0 {
			return "G"
		}
		return "A"
	default:
		if p.symbolGearID > 0 {
			return "*"
		}
		return "#"
	}
}

type Grid struct {
	cols      int
	positions []*Pos
}

func (g *Grid) At(x, y int) *Pos {
	return g.positions[(y*g.cols)+x]
	// 1, 0
}

func (g *Grid) Rows() int {
	return len(g.positions) / g.cols
}

func (g *Grid) Size() int {
	return len(g.positions)
}

func (g Grid) String() string {
	str := ""

	for i := 0; i < g.Rows(); i++ {
		for j := 0; j < g.cols; j++ {
			str += g.At(j, i).String()
		}
		str += "\n"
	}
	return str
}

func Execute() {
	// data := util.ReadLines("day3/calibration1.txt")
	data := util.ReadLines("day3/input.txt")

	c := len(data[0])
	grid := Grid{
		cols:      c,
		positions: make([]*Pos, 0, c*len(data)),
	}

	for _, line := range data {
		addToGrid(&grid, line)
	}
	grid.symbolInfect()

	for i := 0; i < grid.Rows(); i++ {
		grid.digitInfect()
	}

	fmt.Println(grid.String())
	_ = grid.infectedNumberValues()

	for gearID := 1; gearID < gearCounter; gearID++ {
		gears := make([]*InfectedValue, 0)
		for _, x := range infectedValues {
			if x.digitGearID == gearID {
				gears = append(gears, x)
			}
		}
		if len(gears) == 2 {
			gears[0].value = gears[0].value * gears[1].value
			gears[1].value = 0
		} else {
			for _, e := range gears {
				e.value = 0
			}
		}
	}

	sum := 0
	for _, x := range infectedValues {
		if x.digitGearID == 0 {
			x.value = 0
		}
		sum += x.value
		fmt.Printf("'%s': %d\n", x.orig, x.value)
	}

	fmt.Println("sum:", sum)

	// fmt.Println("sum", sum)
}

func addToGrid(grid *Grid, line string) {
	for _, c := range line {
		if unicode.IsDigit(c) {
			grid.positions = append(grid.positions, &Pos{Digit, string(c), 0, 0})
		} else if c == '.' {
			grid.positions = append(grid.positions, &Pos{Space, string(c), 0, 0})
		} else {
			n := 0
			if c == '*' {
				n = gearCounter
				gearCounter = gearCounter + 1
			}
			grid.positions = append(grid.positions, &Pos{Symbol, string(c), n, 0})
		}
	}
}

func (g *Grid) symbolInfect() {
	for y := 0; y < g.Rows(); y++ {
		for x := 0; x < g.cols; x++ {
			symbol := g.At(x, y)
			if symbol.tokenType != Symbol {
				continue
			}

			// we now have a symbol at x,y
			adj := adjacentPositions(x, y, g.cols, g.Rows())
			for _, p := range adj {
				pos := g.At(p[0], p[1])
				if pos.tokenType == Digit {
					if symbol.symbolGearID > 0 {
						pos.digitGearID = symbol.symbolGearID
					}
					pos.tokenType = Adjacent
				}
			}
			// next row
		}
	}
}

func (g *Grid) digitInfect() {
	for y := 0; y < g.Rows(); y++ {
		for x := 0; x < g.cols; x++ {
			inf := g.At(x, y)
			if inf.tokenType != Adjacent {
				continue
			}

			// we now have an infected symbol at x,y
			lr := leftRight(x, y, g.cols, g.Rows())
			for _, p := range lr {
				pos := g.At(p[0], p[1])
				if pos.tokenType == Digit {
					pos.tokenType = Adjacent
					pos.digitGearID = inf.digitGearID
				}
			}
			// next row
		}
	}
}

func (g *Grid) infectedNumberValues() []int {
	str := ""
	currentNumberStr := ""
	var lastDigitGearID int
	for y := 0; y < g.Rows(); y++ {
		for x := 0; x < g.cols; x++ {
			p := g.At(x, y)
			switch p.tokenType {
			case Adjacent:
				lastDigitGearID = p.digitGearID
				currentNumberStr += p.value
				str += p.value
			default:
				if currentNumberStr != "" {
					n, err := strconv.Atoi(currentNumberStr)
					if err != nil {
						panic(err)
					}
					infectedValues = append(infectedValues, &InfectedValue{
						value:       n,
						digitGearID: lastDigitGearID,
						orig:        currentNumberStr,
					})
				}
				currentNumberStr = ""
				str += " "
			}
		}
		if currentNumberStr != "" {
			n, err := strconv.Atoi(currentNumberStr)
			if err != nil {
				panic(err)
			}
			infectedValues = append(infectedValues, &InfectedValue{
				value:       n,
				digitGearID: lastDigitGearID,
				orig:        currentNumberStr,
			})
		}
		currentNumberStr = ""
		str += " "
	}
	numbers := make([]int, 0)
	for _, s := range strings.Fields(str) {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}
	return numbers
}

func leftRight(x, y, cols, rows int) [][2]int {
	lr := [2][2]int{
		{x - 1, y}, {x + 1, y},
	}
	lrFixed := make([][2]int, 0, len(lr))
	for _, c := range lr {
		if c[0] != -1 && c[1] != -1 && c[0] != cols && c[1] != rows {
			lrFixed = append(lrFixed, c)
		}
	}
	return lrFixed
}

func adjacentPositions(x, y, cols, rows int) [][2]int {
	adj := [][2]int{
		{x - 1, y - 1}, {x - 1, y}, {x - 1, y + 1},
		{x, y - 1}, {x, y + 1},
		{x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1},
	}
	adjFixed := make([][2]int, 0, len(adj))
	for _, c := range adj {
		if c[0] != -1 && c[1] != -1 && c[0] != cols && c[1] != rows {
			adjFixed = append(adjFixed, c)
		}
	}
	return adjFixed
}

func Ratios(input string) (int, error) {
	return 0, nil
}
