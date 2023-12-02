package advent_of_code

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	fmt.Printf("Processed %v lines. Total: %v", linesProcessed, total)
}

func solve(inputLines []string) (int, int) {
	linesProcessed := 0
	total := 0

	return total, linesProcessed
}
