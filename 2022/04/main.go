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
		log.Fatalln(e)
	}
}

func main() {
	inputLines := readInput("./input")
	fmt.Printf("Number of lines: %v\n", len(inputLines))
	total, linesProcessed := solve(inputLines)
	fmt.Printf("01: Processed %v lines. Total: %v\n", linesProcessed, total)
	total, linesProcessed = solve2(inputLines)
	fmt.Printf("02: Processed %v lines. Total: %v", linesProcessed, total)
}

func solve2(lines []string) (int, int) {
	linesProcessed := 0
	total := 0

	for _, line := range lines {

		pair := parsePair(line)
		if hasOverlap(pair) {
			total++
		}

		linesProcessed++
	}

	return total, linesProcessed
}

func hasOverlap(pair Pair) bool {
	first := pair.first
	second := pair.second
	fmt.Printf("First %v-%v \tSecond %v-%v\n", first.from, first.to, second.from, second.to)
	firstSpillsToSecond := first.to >= second.from && first.to <= second.to
	secondSpillsToFirst := second.to >= first.from && second.to <= first.to
	return firstSpillsToSecond || secondSpillsToFirst
}

type Pair struct {
	first  Assignment
	second Assignment
}

type Assignment struct {
	from int
	to   int
}

func solve(inputLines []string) (int, int) {
	linesProcessed := 0
	total := 0

	for _, line := range inputLines {

		pair := parsePair(line)
		if hasFullOverlap(pair) {
			total++
		}

		linesProcessed++
	}

	return total, linesProcessed
}

func hasFullOverlap(pair Pair) bool {
	first := pair.first
	second := pair.second
	return isWithin(first, second) || isWithin(second, first)
}

func isWithin(first Assignment, second Assignment) bool {
	return first.from >= second.from && first.to <= second.to
}

func parsePair(inputLine string) Pair {
	assignments := strings.Split(inputLine, ",")
	return Pair{parseAssignment(assignments[0]), parseAssignment(assignments[1])}
}

func parseAssignment(assignment string) Assignment {
	bounds := strings.Split(assignment, "-")
	return Assignment{intOf(bounds[0]), intOf(bounds[1])}
}

func intOf(s string) int {
	res, err := strconv.Atoi(s)
	check(err)
	return res
}
