package day7

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/golang/glog"
	"github.com/westonbelk/adventofcode/util"
)

type Hand struct {
	Type     Type
	Original string
	Cards    []Card
	Wager    int
}

//go:generate stringer -type=Type
type Type int

const (
	HighCard Type = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

//go:generate stringer -type=Card
type Card int

const (
	// J Card = 1
	T Card = iota + 10
	J
	Q
	K
	A
)

func Execute() {
	input := util.ReadLines("day7/input.txt")
	// input := util.ReadLines("day7/calibration.txt")

	hands := make([]Hand, 0)
	for _, s := range input {
		fields := strings.Fields(s)
		hand := toHand(fields[0], fields[1])
		hand.Original = s
		hand.Type = calculateType(hand.Cards)
		hands = append(hands, hand)
	}

	sort.SliceStable(hands, func(i, j int) bool {
		lhs, rhs := hands[i], hands[j]
		if lhs.Type == rhs.Type {
			for n := range lhs.Cards {
				if lhs.Cards[n] == rhs.Cards[n] {
					continue
				}
				return lhs.Cards[n] < rhs.Cards[n]
			}
		}
		if lhs.Type == rhs.Type {
			glog.Fatalf("unable to properly sort hands: %v, %v\n", lhs, rhs)
		}
		return lhs.Type < rhs.Type
	})
	// slices.Reverse(hands)
	fmt.Println(hands)

	total := 0
	for i, h := range hands {
		total += (i + 1) * h.Wager
		fmt.Printf("%d * %d + ", (i + 1), h.Wager)
	}
	fmt.Println()
	fmt.Println("total:", total)
}

func calculateType(cardsUnsorted []Card) Type {
	cards := slices.Clone(cardsUnsorted)
	slices.Sort(cards)

	matches := make([]int, 0)
	count := 0
	// numJs := 0
	last := Card(-1)
	for i := 0; i < len(cards); i++ {
		card := cards[i]

		if i == 0 {
			count = 1
			last = card
			continue
		}

		if card == last {
			count = count + 1
		}

		if card != last {
			matches = append(matches, count)
			count = 1
		}

		if i == (len(cards) - 1) {
			matches = append(matches, count)
		}

		last = card
	}
	slices.Sort(matches)
	slices.Reverse(matches)
	longest := matches[0]
	secondLongest := 0
	if len(matches) > 1 {
		secondLongest = matches[1]
	}

	fmt.Println(cards)
	fmt.Println(matches)
	fmt.Println()

	// glog.Exit(1)

	if longest == 5 {
		return FiveOfAKind
	}
	if longest == 4 {
		return FourOfAKind
	}
	if longest == 3 && secondLongest == 2 {
		return FullHouse
	}
	if longest == 3 {
		return ThreeOfAKind
	}
	if longest == 2 && secondLongest == 2 {
		return TwoPair
	}
	if longest == 2 {
		return OnePair
	}
	if longest == 1 && secondLongest == 1 {
		return HighCard
	}
	glog.Fatalf("failed to find a type for hand: %v", cardsUnsorted)
	panic("failed to find a type for hand")
}

func toHand(cardsRaw, wagerRaw string) Hand {
	hand := Hand{}
	wager, err := strconv.Atoi(wagerRaw)
	if err != nil {
		glog.Fatalf("failed to convert wager to number: %v", wagerRaw)
	}
	hand.Wager = wager

	cards := make([]Card, len(cardsRaw))
	for i, r := range cardsRaw {
		if unicode.IsDigit(r) {
			c, err := strconv.Atoi(string(r))
			if err != nil {
				glog.Fatalf("failed to convert card to number: %v", string(r))
			}
			cards[i] = Card(c)
			continue
		}
		switch string(r) {
		case "T":
			cards[i] = T
		case "J":
			cards[i] = J
		case "Q":
			cards[i] = Q
		case "K":
			cards[i] = K
		case "A":
			cards[i] = A
		}
	}
	hand.Cards = cards
	return hand
}
