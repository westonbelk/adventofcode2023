package day7

import "strconv"

func (i Card) String() string {
	switch i {
	case T:
		return "T"
	case J:
		return "J"
	case K:
		return "K"
	case Q:
		return "Q"
	case A:
		return "A"
	default:
		return strconv.FormatInt(int64(i), 10)
	}
}
