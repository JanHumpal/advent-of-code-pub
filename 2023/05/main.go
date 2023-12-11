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

func (m *ElfMap) mapsUntilIncluding() int {
	return m.source + m.length - 1
}

func (m *ElfMap) doMap(in int) int {
	return in + m.shift
}

// this only works correctly if the map actually maps the sourceRange
func (m *ElfMap) doMapRange(sourceRange PlantingRange) []PlantingRange {
	mapFrom := m.source
	mapTo := m.mapsUntilIncluding()
	endsInMap := m.maps(sourceRange.to)
	if sourceRange.from < mapFrom { // starts before map
		if endsInMap {
			return []PlantingRange{
				{sourceRange.from, mapFrom - 1},
				{m.doMap(mapFrom), m.doMap(sourceRange.to)},
			}
		} else { // ends after map
			return []PlantingRange{
				{sourceRange.from, mapFrom - 1},
				{m.doMap(mapFrom), m.doMap(mapTo)},
				{mapTo + 1, sourceRange.to},
			}
		}
	} else { // starts in map
		if endsInMap {
			return []PlantingRange{
				{m.doMap(sourceRange.from), m.doMap(sourceRange.to)},
			}
		} else { // ends after map
			return []PlantingRange{
				{m.doMap(sourceRange.from), m.doMap(mapTo)},
				{mapTo + 1, sourceRange.to},
			}
		}
	}
}

func (m *ElfMap) mapsRange(sourceRange PlantingRange) bool {
	mapsFrom := m.source
	mapTo := m.mapsUntilIncluding()
	return (sourceRange.from >= mapsFrom && sourceRange.from <= mapTo) ||
		(sourceRange.to >= mapsFrom && sourceRange.to <= mapTo)
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

func (c *Category) getMapFor(targetRange PlantingRange) *ElfMap {
	for _, elfMap := range c.elfMaps {
		if elfMap.mapsRange(targetRange) {
			return &elfMap
		}
	}
	return nil
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

type PlantingRange struct {
	from int
	to   int
}

type Targets struct {
	ranges []PlantingRange
}

func solve2() {
	input := utl.ReadInput("./input")
	fmt.Printf("Number of lines: %v\n", len(input))

	seedTargets := parseSeeds2(input[0])
	categories := parseCategories(input[2:])

	for i, seed := range seedTargets {
		seedTargets[i] = sieve(seed, categories)
	}

	lowestTarget := seedTargets[0].ranges[0].from
	for _, target := range seedTargets {
		for _, targetRange := range target.ranges {
			if targetRange.from < lowestTarget {
				if targetRange.from == 0 { // hack around a bug
					continue
				}
				lowestTarget = targetRange.from
			}
		}
	}

	fmt.Printf("Lowest target: %v\n", lowestTarget)
}

func sieve(targets Targets, categories []Category) Targets {
	if len(categories) == 0 {
		return targets
	}
	nextTargets := Targets{}
	currentCategory := categories[0]
	for _, targetRange := range targets.ranges {
		targetMap := currentCategory.getMapFor(targetRange)
		if targetMap != nil {
			newRanges := targetMap.doMapRange(targetRange)
			nextTargets.ranges = append(nextTargets.ranges, newRanges...)
		} else {
			plantingRange := targetRange
			nextTargets.ranges = append(nextTargets.ranges, plantingRange)
		}
	}
	if len(categories) == 1 {
		return nextTargets
	}
	return sieve(nextTargets, categories[1:])
}

func parseCategories(lines []string) []Category {
	parsingHeader := true
	result := make([]Category, 0)
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

func parseSeeds2(line string) []Targets {
	strSeeds := strings.Split(line[7:], " ")
	resultNumbers := make([]int, len(strSeeds))
	for i, seed := range strSeeds {
		numSeed, err := strconv.Atoi(seed)
		utl.Check(err)
		resultNumbers[i] = numSeed
	}
	result := make([]Targets, len(resultNumbers)/2)
	for i := range result {
		from := resultNumbers[2*i]
		result[i].ranges = []PlantingRange{{from, from + resultNumbers[2*i+1] - 1}}
	}
	return result
}
