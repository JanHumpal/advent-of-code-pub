package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func digitNames() map[string]int {
	return map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
}

func main() {
	textLines := readInput("./input")
	fmt.Printf("Number of lines: %v\n", len(textLines))

	linesProcessed := 0
	total := 0

	for _, line := range textLines {
		firstNum, lastNum := parseDigits(line)

		fmt.Printf("%v%v :\t%v\n", firstNum, lastNum, line)
		total += firstNum*10 + lastNum
		linesProcessed++
	}

	fmt.Printf("Processed %v lines. Total is %v\n", linesProcessed, total)
}

func parseDigits(line string) (int, int) {
	firstStringI, lastStringI, firstWordNum, lastWordNum := -1, -1, -1, -1

	for digit := range digitNames() {
		currentFirstStringI := strings.Index(line, digit)
		if firstStringI == -1 {
			firstStringI = currentFirstStringI
		}
		if (currentFirstStringI >= 0) && (currentFirstStringI <= firstStringI) {
			firstStringI = currentFirstStringI
			firstWordNum = digitNames()[digit]
		}

		currentLastStringI := strings.LastIndex(line, digit)
		if (currentLastStringI >= 0) && (currentLastStringI > lastStringI) {
			lastStringI = currentLastStringI
			lastWordNum = digitNames()[digit]
		}
	}

	firstIntI := strings.IndexFunc(line, unicode.IsDigit)
	lastIntI := strings.LastIndexFunc(line, unicode.IsDigit)

	firstNum := -1
	lastNum := -1

	if (firstStringI >= 0) && (firstStringI < firstIntI) {
		firstNum = firstWordNum
	} else {
		firstNum = int(line[firstIntI] - '0')
	}

	if lastStringI > lastIntI {
		lastNum = lastWordNum
	} else {
		lastNum = int(line[lastIntI] - '0')
	}
	return firstNum, lastNum
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
