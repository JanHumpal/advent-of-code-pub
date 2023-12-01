package main

import (
	"bufio"
	"fmt"
	"os"
)

func scoreMap() map[string]int {
	return map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}
}

func scoreMap2() map[string]int {
	return map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}
}

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
		panic(e)
	}
}

func main() {

	textLines := readInput("./input")
	fmt.Printf("Number of lines: %v\n", len(textLines))

	linesProcessed := 0
	lineScore := 0
	total := 0
	scoreSheet := scoreMap2()

	for _, line := range textLines {
		lineScore = scoreSheet[line]
		if lineScore < 1 {
			panic("unexpected input: " + line)
		}
		total += lineScore
		linesProcessed++
	}

	fmt.Printf("Total score: %v", total)
}
