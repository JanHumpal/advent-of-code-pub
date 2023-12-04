package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"utl"
)

func main() {
	solve()
}

type Card struct {
	id             int
	winningNumbers []int
	myNumbers      []int
}

func solve() {
	input := utl.ReadInput("./input")
	fmt.Printf("Number of lines: %v\n", len(input))
	linesProcessed := 0
	totalScore := 0
	totalScore2 := 0
	totalLines := 203

	cards := toCards(input)

	copies := make([]int, totalLines)
	for i := range copies {
		copies[i] = 1
	}

	for i, card := range cards {
		totalScore += card.getScore()

		wins := card.getScore2()
		for j := i + 1; j < min(i+wins+1, totalLines); j++ {
			copies[j] += copies[i]
		}

		linesProcessed++
	}

	for _, copyCount := range copies {
		totalScore2 += copyCount
	}

	fmt.Printf("Processed %v lines. Total score: %v Total score 2: %v\n", linesProcessed, totalScore, totalScore2)
}

func toCards(input []string) []Card {
	result := make([]Card, len(input))
	for i, line := range input {
		result[i] = toCard(line)
	}
	return result
}

func toCard(line string) Card {
	result := Card{}
	idIndex := strings.IndexFunc(line, unicode.IsDigit)
	colon := strings.IndexRune(line, ':')
	id, err := strconv.Atoi(line[idIndex:colon])
	utl.Check(err)
	result.id = id

	line = line[colon+2:]

	parts := strings.Split(line, "|")

	winningStrings := strings.Split(parts[0], " ")
	result.winningNumbers = toNumbers(winningStrings)

	myStrings := strings.Split(parts[1], " ")
	result.myNumbers = toNumbers(myStrings)

	return result
}

func toNumbers(stringNumbers []string) []int {
	result := make([]int, 0, len(stringNumbers))
	for _, number := range stringNumbers {
		if len(number) > 0 {
			x, err := strconv.Atoi(number)
			utl.Check(err)
			result = append(result, x)
		}
	}
	return result
}

func (card *Card) getScore() int {
	winCount := 0
	for _, number := range card.winningNumbers {
		if slices.Contains(card.myNumbers, number) {
			winCount++
		}
	}
	if winCount == 0 {
		return 0
	}
	if winCount == 1 {
		return 1
	}
	score := 1
	for i := 0; i < winCount-1; i++ {
		score *= 2
	}
	return score
}

func (card *Card) getScore2() int {
	winCount := 0
	for _, number := range card.winningNumbers {
		if slices.Contains(card.myNumbers, number) {
			winCount++
		}
	}
	return winCount
}
