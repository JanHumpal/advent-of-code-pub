package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"utl"
)

func main() {
	solve()
}

const (
	HIGH       = iota
	PAIR       = iota
	TWO_PAIR   = iota
	THREE_OAK  = iota
	FULL_HOUSE = iota
	POKER      = iota
	FIVE_OAK   = iota
)

type Hand struct {
	kind  int
	cards string
}

func cardValues() map[uint8]int {
	return map[uint8]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}
}

func cardValues2() map[uint8]int {
	return map[uint8]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}
}

func (h Hand) isLessThan(other Hand) bool {
	if h.kind != other.kind {
		return h.kind < other.kind
	}
	//values := cardValues()
	values := cardValues2()
	for i := 0; i < 5; i++ { // 5 cards in a hand
		if h.cards[i] == other.cards[i] {
			continue
		}
		isLess := values[h.cards[i]] < values[other.cards[i]]
		return isLess
	}
	log.Fatalf("Failed to compare hands %v and %v\n", h, other)
	panic(1)
}

type Game struct {
	bid  int
	hand Hand
}

func solve() {
	input := utl.ReadInput("./input")
	fmt.Printf("Number of lines: %v\n", len(input))

	total := 0

	games := parseGames(input)

	sort.Slice(games, func(i, j int) bool {
		return games[i].hand.isLessThan(games[j].hand)
	})

	for i, game := range games {
		total += (i + 1) * game.bid
	}

	fmt.Printf("Total: %v\n", total)
}

func parseGames(lines []string) []Game {
	result := make([]Game, len(lines))
	for i, line := range lines {
		result[i] = parseGame2(line)
	}
	return result
}

func parseGame(line string) Game {
	parts := strings.Split(line, " ")
	hand := Hand{0, parts[0]}
	result := Game{utl.IntOf(parts[1]), hand}
	result.hand.kind = getKindOf(hand.cards)
	return result
}

func parseGame2(line string) Game {
	parts := strings.Split(line, " ")
	hand := Hand{0, parts[0]}
	result := Game{utl.IntOf(parts[1]), hand}
	result.hand.kind = getKindOf2(hand.cards)
	return result
}

func getKindOf(cards string) int {
	multi1, multi2 := findMultiples(cards)

	switch multi1 {
	case 0:
		return HIGH
	case 2:
		if multi2 == 2 {
			return TWO_PAIR
		}
		if multi2 == 3 {
			return FULL_HOUSE
		}
		return PAIR
	case 3:
		if multi2 == 2 {
			return FULL_HOUSE
		}
		return THREE_OAK
	case 4:
		return POKER
	case 5:
		return FIVE_OAK
	}
	log.Fatalln("Could not get kind of " + cards)
	return -1
}

func getKindOf2(cards string) int {
	multi1, multi2, jokers := findMultiples2(cards)

	switch multi1 {
	case 0:
		switch jokers {
		case 0:
			return HIGH
		case 1:
			return PAIR
		case 2:
			return THREE_OAK
		case 3:
			return POKER
		case 4, 5:
			return FIVE_OAK
		}
	case 2:
		if multi2 == 2 {
			if jokers == 1 {
				return FULL_HOUSE
			}
			return TWO_PAIR
		}
		if multi2 == 3 {
			return FULL_HOUSE
		}
		// multi2 == 0
		switch jokers {
		case 0:
			return PAIR
		case 1:
			return THREE_OAK
		case 2:
			return POKER
		case 3:
			return FIVE_OAK
		}
	case 3:
		if multi2 == 2 {
			return FULL_HOUSE
		}
		// multi2 == 0
		switch jokers {
		case 0:
			return THREE_OAK
		case 1:
			return POKER
		case 2:
			return FIVE_OAK
		}
	case 4:
		if jokers == 1 {
			return FIVE_OAK
		}
		return POKER
	case 5:
		return FIVE_OAK
	}
	log.Fatalln("Could not get kind of " + cards)
	return -1
}

func findMultiples(cards string) (int, int) {
	blank := "*"
	var multi1, multi2 int
	for i := 0; i < 4; i++ { // 4 = 5 cards - 1 last card that can no longer have multiple
		card := cards[i]
		if card == blank[0] {
			continue
		}
		count := strings.Count(cards[i+1:], string(card))
		if count > 0 {
			if multi1 > 0 {
				multi2 = count + 1
			} else {
				multi1 = count + 1
			}
			cards = strings.ReplaceAll(cards, string(card), blank)
		}
	}
	return multi1, multi2
}

func findMultiples2(cards string) (int, int, int) {
	blank := "*"
	jokers := strings.Count(cards, "J")
	cards = strings.ReplaceAll(cards, "J", blank)
	var multi1, multi2 int
	for i := 0; i < 4; i++ { // 4 = 5 cards - 1 last card that can no longer have multiple
		card := cards[i]
		if card == blank[0] {
			continue
		}
		count := strings.Count(cards[i+1:], string(card))
		if count > 0 {
			if multi1 > 0 {
				multi2 = count + 1
			} else {
				multi1 = count + 1
			}
			cards = strings.ReplaceAll(cards, string(card), blank)
		}
	}
	return multi1, multi2, jokers
}
