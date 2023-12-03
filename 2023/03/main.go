package main

import (
	"fmt"
	"strings"
	"unicode"
	"utl"
)

func isDot(c rune) bool {
	return c == '.'
}

func isDigit(c rune) bool {
	return unicode.IsDigit(c)
}

func isNotDigit(c rune) bool {
	return !isDigit(c)
}

func main() {
	solve()
}

type Part struct {
	value  int
	symbol rune
}

type Line struct {
	numbers map[int]int
	symbols map[int]rune
}

func solve() {
	input := utl.ReadInput("./input")
	fmt.Printf("Number of lines: %v\n", len(input))
	linesProcessed := 0
	totalParts := 0
	totalGears := 0

	lines := toLines(input)
	linesProcessed = len(lines)
	for i := 0; i < linesProcessed; i++ {
		surroundingLines := utl.GetSurrounding(lines, i)
		parts := getParts(lines[i], surroundingLines)
		for _, part := range parts {
			totalParts += part.value
		}
		gears := getGears(lines[i], surroundingLines)
		for _, gear := range gears {
			totalGears += gear.a * gear.b
		}
	}

	fmt.Printf("Processed %v lines. Total parts: %v Total gears: %v", linesProcessed, totalParts, totalGears)
	//for _, line := range lines {
	//	line.reconstruct(140)
	//	fmt.Println()
	//}
}

func getGears(line Line, surroundingLines []Line) []Gear {
	result := make([]Gear, 0)
	for i, symbol := range line.symbols {
		if symbol != '*' {
			continue
		}
		a, b := getGearParts(i, surroundingLines)
		if a != 0 && b != 0 {
			result = append(result, Gear{a, b})
		}
	}
	return result
}

func getGearParts(gearIndex int, surroundingLines []Line) (int, int) {
	a, b := 0, 0
	for _, line := range surroundingLines {
		for numberIndex, number := range line.numbers {
			length := utl.DigitCount(number)
			if gearIndex >= numberIndex-1 && gearIndex <= numberIndex+length {
				if a == 0 {
					a = number
				} else {
					b = number
				}
			}
		}
	}
	return a, b
}

type Gear struct {
	a int
	b int
}

func toLines(input []string) []Line {
	result := make([]Line, len(input))

	for i, inputLine := range input {
		result[i] = toLine(inputLine)
	}
	return result
}

func toLine(line string) Line {
	result := Line{make(map[int]int), make(map[int]rune)}
	length := len(line)
	char := '.'
	//fmt.Println(line)
	for i := 0; i < length; i++ {
		char = int32(line[i])
		if isDot(char) {
			continue
		}
		if isDigit(char) {
			//fmt.Printf("examining %c at %v\n", char, i)
			//fmt.Printf("examining %s\n", line[i:])
			end := strings.IndexFunc(line[i:], isNotDigit)
			if end == -1 {
				end = length
			} else {
				end += i
			}
			//fmt.Printf("it ends at %v\n", end)
			value := utl.IntOf(line[i:end])
			//fmt.Printf("the full value is %v\n", value)
			result.numbers[i] = value
			i = end - 1
			//fmt.Printf("continuing at %v with %v\n", i, line[i:])
			continue
		}
		result.symbols[i] = char
	}
	return result
}

func getParts(line Line, surroundingLines []Line) []Part {
	result := make([]Part, 0, len(line.numbers))
	//fmt.Printf("=== Getting parts of \n%v\n", line)
	for numberI, number := range line.numbers {
		isAPart, symbol := isPart(number, numberI, surroundingLines)
		if isAPart {
			result = append(result[:], Part{number, symbol})
		}
	}
	return result
}

func isPart(number int, start int, surroundingLines []Line) (bool, rune) {
	end := start + utl.DigitCount(number)
	//fmt.Printf("Checking number %v between %v and %v among lines:\n", number, start, end)
	//for _, line := range surroundingLines {
	//	line.printout()
	//}
	for _, line := range surroundingLines {
		for i := start - 1; i < end+1; i++ {
			symbol := line.symbols[i]
			if symbol != 0 {
				//fmt.Printf("This number %v is a part with symbol %c\n", number, symbol)
				return true, symbol
			}
		}
	}
	return false, 0
}

func (line *Line) printout() {
	fmt.Printf("\tnumbers:\t")
	for i, num := range line.numbers {
		fmt.Printf("%v: %v, ", i, num)
	}
	fmt.Printf("\n\tsymbols:\t")
	for i, sym := range line.symbols {
		fmt.Printf("%v: %c, ", i, sym)
	}
	fmt.Printf("\n\n")
}

// todo this could fail if input has zeroes as numbers
func (line *Line) reconstruct(length int) {
	for i := 0; i < length; i++ {
		symbol := line.symbols[i]
		if symbol != 0 {
			fmt.Printf("%c", symbol)
			continue
		}
		number := line.numbers[i]
		if number != 0 {
			fmt.Printf("%v", number)
			i += utl.DigitCount(number) - 1
			continue
		}
		fmt.Print(".")
	}
}
