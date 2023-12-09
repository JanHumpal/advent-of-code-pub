package main

import (
	"fmt"
	"log"
	"strings"
	"utl"
)

func main() {
	solve2()
}

func solve() {
	lines := utl.ReadInput("./input")
	fmt.Printf("Number of lines: %v\n", len(lines))

	total := 0
	histories := parseInput(lines)

	for _, history := range histories {
		total += utl.Last(extrapolate(history))
	}

	fmt.Printf("Total: %v\n", total)
}

func solve2() {
	lines := utl.ReadInput("./input")
	fmt.Printf("Number of lines: %v\n", len(lines))

	fmt.Printf("10  13  16  21  30  45: %v\n", extrapolateBackwards([]int{10, 13, 16, 21, 30, 45}))

	total := 0
	histories := parseInput(lines)

	for _, history := range histories {
		total += utl.Last(extrapolateBackwards(history))
	}

	fmt.Printf("Total: %v\n", total)
}

func extrapolateBackwards(history []int) []int {
	diffs := getDiffs(history)
	if areAllSame(diffs) {
		return []int{diffs[0], history[0] - diffs[0]}
	}
	extrapolations := extrapolateBackwards(diffs)
	return append(extrapolations, history[0]-utl.Last(extrapolations))
}

func parseInput(lines []string) [][]int {
	result := make([][]int, len(lines))
	for i, line := range lines {
		result[i] = parse(line)
	}
	return result
}

func parse(history string) []int {
	fields := strings.Fields(history)
	result := make([]int, len(fields))
	for i, field := range fields {
		result[i] = utl.IntOf(field)
	}
	return result
}

func extrapolate(history []int) []int {
	diffs := getDiffs(history)
	if areAllSame(diffs) {
		return []int{diffs[0], utl.Last(history) + diffs[0]}
	}
	extrapolations := extrapolate(diffs)
	return append(extrapolations, utl.Last(history)+utl.Last(extrapolations))
}

func areAllSame(diffs []int) bool {
	if len(diffs) == 0 {
		log.Fatalln("illegal state")
	}
	first := diffs[0]
	for _, diff := range diffs {
		if diff != first {
			return false
		}
	}
	return true
}

func getDiffs(history []int) []int {
	results := make([]int, len(history)-1)
	for i := range results {
		results[i] = history[i+1] - history[i]
	}
	return results
}
