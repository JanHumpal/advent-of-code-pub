package main

import (
	"bufio"
	"fmt"
	"os"
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
		panic(e)
	}
}

func main() {

	//fmt.Printf("a: %v\n", int('a'))
	//fmt.Printf("A: %v\n", int('A'))
	//fmt.Printf("A-a: %v\n", int('A'-'a'))
	//fmt.Printf("a-96: %v\n", int('a'-96))
	//fmt.Printf("z-96: %v\n", int('z'-96))
	//fmt.Printf("A-96: %v\n", int('A'-96))
	//fmt.Printf("Z-96: %v\n", int('Z'-96))
	//fmt.Printf("A-96+58: %v\n", int('A'-96+58))
	//fmt.Printf("Z-96+58: %v\n", int('Z'-96+58))

	inputLines := readInput("./input")
	fmt.Printf("Number of lines: %v\n", len(inputLines))
	total, linesProcessed := solve(inputLines)
	total, linesProcessed = solve2(inputLines)
	fmt.Printf("Processed %v lines. Total: %v", linesProcessed, total)

}

func solve2(inputLines []string) (int, int) {
	linesProcessed := 0
	total := 0

	for i := 0; i < len(inputLines); i += 3 {
		elf1 := inputLines[i]
		elf2 := inputLines[i+1]
		elf3 := inputLines[i+2]
		badge := '0'

		for _, char := range elf1 {
			if strings.ContainsRune(elf2, char) && strings.ContainsRune(elf3, char) {
				badge = char
				break
			}
		}
		total += priorityFrom(badge)
		linesProcessed += 3
	}

	return total, linesProcessed
}

func solve(inputLines []string) (int, int) {
	linesProcessed := 0
	total := 0

	for _, rucksack := range inputLines {
		total += getPriority(rucksack)
		linesProcessed++
	}

	return total, linesProcessed
}

func getPriority(rucksack string) int {
	left, right := getCompartments(rucksack)
	priority := 0
	for _, char := range left {
		if strings.ContainsRune(right, char) {
			priority = priorityFrom(char)
			fmt.Printf("%v:%v", string(char), priority)
			break
		}
	}
	fmt.Printf("\t%v = %v + %v\n", rucksack, left, right)
	return priority
}

func priorityFrom(char int32) int {
	result := int(char - 96)
	if result < 1 {
		result += 58
	}
	return result
}

func getCompartments(rucksack string) (string, string) {
	half := len(rucksack) / 2
	return rucksack[:half], rucksack[half:]
}
