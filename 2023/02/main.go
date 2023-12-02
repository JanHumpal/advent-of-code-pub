package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(fileName string) []string {
	inputFile, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)
	var textLines []string
	for scanner.Scan() {
		textLines = append(textLines, scanner.Text())
	}
	err = inputFile.Close()
	check(err)
	return textLines
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	solve()
}

func solve() {
	textLines := readInput("./input")
	fmt.Printf("Number of lines: %v\n", len(textLines))
	linesProcessed := 0
	totalId := 0
	totalPower := 0

	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	maxHand := Hand{12, 13, 14}

	for i, gameLine := range textLines {
		game := parseGame(gameLine, i+1)
		if isPossible(game, maxHand) {
			totalId += game.id
		}
		totalPower += powerOf(game)
		linesProcessed++
	}

	fmt.Printf("Processed %v lines. TotalId: %v, total power: %v", linesProcessed, totalId, totalPower)
}

func powerOf(game *Game) int {
	minBag := minBagOf(game)
	return minBag.red * minBag.green * minBag.blue
}

func minBagOf(game *Game) Hand {
	minBag := Hand{}
	for _, hand := range game.hands {
		if hand.red > minBag.red {
			minBag.red = hand.red
		}
		if hand.green > minBag.green {
			minBag.green = hand.green
		}
		if hand.blue > minBag.blue {
			minBag.blue = hand.blue
		}
	}
	return minBag
}

func isPossible(game *Game, max Hand) bool {
	for _, hand := range game.hands {
		if isImpossible(hand, max) {
			return false
		}
	}
	return true
}

func isImpossible(hand Hand, max Hand) bool {
	return hand.red > max.red ||
		hand.green > max.green ||
		hand.blue > max.blue
}

func parseGame(line string, id int) *Game {
	result := new(Game)
	result.id = id

	i := strings.IndexRune(line, ':')
	handsLine := line[i+2:]
	result.hands = parseHandsLine(handsLine)

	return result
}

func parseHandsLine(line string) []Hand {
	hands := strings.Split(line, "; ")
	return parseHands(hands)
}

func parseHands(hands []string) []Hand {
	result := make([]Hand, len(hands))
	for i, hand := range hands {
		result[i] = parseHand(hand)
	}
	return result
}

func parseHand(hand string) Hand {
	shows := strings.Split(hand, ", ")
	return toHand(shows)
}

func toHand(shows []string) Hand {
	result := Hand{}
	//fmt.Printf("%v", shows)
	for _, show := range shows {
		i := strings.Index(show, " red")
		if i > -1 {
			red, err := strconv.Atoi(show[:i])
			check(err)
			result.red = red
			continue
		}
		i = strings.Index(show, " green")
		if i > -1 {
			green, err := strconv.Atoi(show[:i])
			check(err)
			result.green = green
			continue
		}
		i = strings.Index(show, " blue")
		if i > -1 {
			blue, err := strconv.Atoi(show[:i])
			check(err)
			result.blue = blue
			continue
		}
	}
	//fmt.Printf(" => {%v red, %v green, %v blue}\n", result.red, result.green, result.blue)
	return result
}

type Hand struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id    int
	hands []Hand
}
