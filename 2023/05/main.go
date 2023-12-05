package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"utl"
)

func main() {
	solve2()
}

type ElfMap struct {
	source int
	target int
	length int
	shift  int // = target - source
}

func (m *ElfMap) maps(in int) bool {
	return in >= m.source && in < m.source+m.length
}

func (m *ElfMap) doMap(in int) int {
	return in + m.shift
}

type Category struct {
	from    string
	to      string
	elfMaps []ElfMap
}

func (c *Category) findTarget(source int) int {
	for _, elfMap := range c.elfMaps {
		if elfMap.maps(source) {
			return elfMap.doMap(source)
		}
	}
	return source
}

func (c *Category) findTargets(targets *[]int) []int {
	result := make([]int, len(*targets))
	for i, target := range *targets {
		result[i] = c.findTarget(target)
	}
	return result
}

func solve() {
	input := utl.ReadInput("./input")
	fmt.Printf("Number of lines: %v\n", len(input))

	seeds := parseSeeds(input[0])
	categories := parseCategories(input[2:])
	locations := make([]int, len(seeds))

	for i, seed := range seeds {
		target := seed
		for _, category := range categories {
			target = category.findTarget(target)
		}
		locations[i] = target
	}

	fmt.Printf("Lowest target: %v\n", slices.Min(locations))
}

type Seed struct {
	from int
	to   int
}

func solve2() {
	input := utl.ReadInput("./input")
	fmt.Printf("Number of lines: %v\n", len(input))

	seeds := parseSeeds2(input[0])
	categories := parseCategories(input[2:])
	allTargets := make([][]int, len(seeds))

	for i, seed := range seeds {
		targets := make([]int, seed.to-seed.from+1)
		for _, category := range categories {
			targets = category.findTargets(&targets)
		}
		allTargets[i] = targets
	}

	lowestTargets := make([]int, len(allTargets))
	for i := range lowestTargets {
		lowestTargets[i] = slices.Min(allTargets[i])
	}

	fmt.Printf("Lowest target: %v\n", slices.Min(lowestTargets))
}

func parseCategories(lines []string) []Category {
	parsingHeader := true
	result := make([]Category, 0, len(lines))
	for _, line := range lines {
		if parsingHeader {
			from, to := parseHeader(line)
			result = append(result, Category{from, to, make([]ElfMap, 0)})
			parsingHeader = false
		} else if line == "" {
			parsingHeader = true
		} else {
			lastResult := last(result)
			lastResult.elfMaps = append(lastResult.elfMaps, parseMap(line))
		}
	}
	return result
}

func parseMap(line string) ElfMap {
	result := ElfMap{}
	numbers := strings.Split(line, " ")
	if len(numbers) != 3 {
		log.Fatalf("Unexpected map input %v\n", line)
	}

	target, err := strconv.Atoi(numbers[0])
	utl.Check(err)
	result.target = target

	source, err := strconv.Atoi(numbers[1])
	utl.Check(err)
	result.source = source

	length, err := strconv.Atoi(numbers[2])
	utl.Check(err)
	result.length = length

	result.shift = result.target - result.source
	return result
}

func parseHeader(line string) (string, string) {
	firstDash := strings.IndexRune(line, '-')
	from := line[:firstDash]
	lastDash := strings.LastIndex(line, "-")
	to := line[lastDash+1 : len(line)-5] // 5 = len(" map:")
	return from, to
}

func last(result []Category) *Category {
	return &result[len(result)-1]
}

func parseSeeds(line string) []int {
	strSeeds := strings.Split(line[7:], " ")
	result := make([]int, len(strSeeds))
	for i, seed := range strSeeds {
		numSeed, err := strconv.Atoi(seed)
		utl.Check(err)
		result[i] = numSeed
	}
	return result
}

func parseSeeds2(line string) []Seed {
	strSeeds := strings.Split(line[7:], " ")
	resultNumbers := make([]int, len(strSeeds))
	for i, seed := range strSeeds {
		numSeed, err := strconv.Atoi(seed)
		utl.Check(err)
		resultNumbers[i] = numSeed
	}
	result := make([]Seed, len(resultNumbers)/2)
	for i := range result {
		from := resultNumbers[2*i]
		result[i] = Seed{from, from + resultNumbers[2*i+1] - 1}
	}
	return result
}
