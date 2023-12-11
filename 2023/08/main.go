package main

import (
	"fmt"
	"strings"
	"utl"
)

func main() {
	solve2()
}

type Transitions struct {
	left  string
	right string
}

func solve() {
	input := utl.ReadInput("./input")
	fmt.Printf("Number of lines: %v\n", len(input))

	totalSteps := 0

	directions := input[0]
	theMap := parseMap(input[2:])
	startingNode := "AAA"

	totalSteps = findPathLength(directions, startingNode, theMap, isExit)

	fmt.Printf("Total: %v\n", totalSteps)
}

type testExit func(string) bool

func findPathLength(
	directions string,
	startingNode string,
	theMap map[string]Transitions,
	isExitNode testExit) int {

	totalSteps := 0
	dirL := len(directions)
	foundExit := false
	currentNode := startingNode

	for !foundExit {
		direction := string(directions[totalSteps%dirL])
		if direction == "L" {
			currentNode = theMap[currentNode].left
		} else {
			currentNode = theMap[currentNode].right
		}
		totalSteps++
		if isExitNode(currentNode) {
			foundExit = true
		}
	}
	return totalSteps
}

func isExit(node string) bool {
	return node == "ZZZ"
}

func solve2() {
	input := utl.ReadInput("./input")
	fmt.Printf("Number of lines: %v\n", len(input))

	directions := input[0]
	theMap := parseMap(input[2:])
	currentNodes := getStartingNodes(theMap)

	pathLengths := findPathLengths(currentNodes, directions, theMap)

	fmt.Printf("Path lengths: %v\n", pathLengths)
	fmt.Printf("Path lengths LCM: %v",
		utl.LCM(pathLengths[0], pathLengths[1], pathLengths[2:]...))

	// Naively it does not finish in reasonable time
	//dirL := len(directions)
	//foundExit := false
	//direction := ""
	//foundAllExits := true
	//for i := 0; !foundExit; i++ {
	//	if i == dirL {
	//		i -= dirL
	//	}
	//	direction = string(directions[i])
	//	if direction == "L" {
	//		for i2, currentNode := range currentNodes {
	//			currentNodes[i2] = theMap[currentNode].left
	//		}
	//	} else {
	//		for i2, currentNode := range currentNodes {
	//			currentNodes[i2] = theMap[currentNode].right
	//		}
	//	}
	//
	//	totalSteps++
	//
	//	foundAllExits = true
	//	for _, currentNode := range currentNodes {
	//		if !isExit2(currentNode) {
	//			foundAllExits = false
	//			break
	//		}
	//	}
	//	if foundAllExits {
	//		foundExit = true
	//	}
	//}
}

func isExit2(currentNode string) bool {
	return strings.HasSuffix(currentNode, "Z")
}

func findPathLengths(startingNodes []string, directions string, theMap map[string]Transitions) []int {
	result := make([]int, len(startingNodes))
	for i, node := range startingNodes {
		result[i] = findPathLength(directions, node, theMap, isExit2)
	}
	return result
}

func getStartingNodes(theMap map[string]Transitions) []string {
	result := make([]string, 0)
	for node := range theMap {
		if strings.HasSuffix(node, "A") {
			result = append(result, node)
		}
	}
	return result
}

func parseMap(lines []string) map[string]Transitions {
	result := make(map[string]Transitions, len(lines))

	for _, line := range lines {
		result[line[0:3]] = Transitions{line[7:10], line[12:15]}
	}

	return result
}
