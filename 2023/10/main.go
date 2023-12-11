package main

import (
	"fmt"
	"slices"
	"utl"
)

func main() {
	solve()
}

type Direction struct {
	dX int
	dY int
}

func (d Direction) opposite() Direction {
	return Direction{d.dX * -1, d.dY * -1}
}

var UP = Direction{0, -1}
var DOWN = Direction{0, 1}
var LEFT = Direction{-1, 0}
var RIGHT = Direction{1, 0}
var NOOP = Direction{0, 0}

type GridPoint struct {
	x int
	y int
}

func (p GridPoint) plus(dir Direction) GridPoint {
	return GridPoint{p.x + dir.dX, p.y + dir.dY}
}

type PipeTile struct {
	pos  GridPoint
	pipe Pipe
}

type Pipe struct {
	a, b Direction
}

func (p PipeTile) walk(pipeMap TileMap, direction Direction) PipeTile {
	return pipeMap.at(p.pos.plus(direction))
}

func (p PipeTile) isConnectedTo(tile PipeTile) bool {
	return p.pos.plus(p.pipe.a) == tile.pos || p.pos.plus(p.pipe.b) == tile.pos
}

func pipesByRune() map[rune]Pipe {
	return map[rune]Pipe{
		'|': {UP, DOWN},
		'-': {LEFT, RIGHT},
		'L': {UP, RIGHT},
		'J': {UP, LEFT},
		'7': {LEFT, DOWN},
		'F': {RIGHT, DOWN},
		'.': {NOOP, NOOP},
	}
}

func pipePics() map[Pipe]rune {
	return map[Pipe]rune{
		Pipe{UP, DOWN}:    '│',
		Pipe{LEFT, RIGHT}: '━',
		Pipe{UP, RIGHT}:   '╰',
		Pipe{UP, LEFT}:    '╯',
		Pipe{LEFT, DOWN}:  '╮',
		Pipe{RIGHT, DOWN}: '╭',
		Pipe{NOOP, NOOP}:  '.',
	}
}

type TileMap [][]PipeTile

func (m TileMap) at(point GridPoint) PipeTile {
	return m[point.y][point.x]
}

type PipeLoop []PipeTile

func (l PipeLoop) contains(tile PipeTile) bool {
	return slices.Contains(l, tile)
}

func newPipeLoop(cap int) PipeLoop {
	return make([]PipeTile, 0, cap)
}

func solve() {
	lines := utl.ReadInput("./input")
	fmt.Printf("Number of lines: %v\n", len(lines))

	// for my input, S is F
	startingPipe := pipesByRune()['F']
	tileMap, startingTile := parse(lines, startingPipe)
	steps, loop := walkLoop(startingTile, tileMap)
	furthest := steps / 2

	fmt.Printf("Steps: %v, Furthest: %v\n", steps, furthest)

	//printLoop(tileMap, loop)

	nestArea := calculateNestArea(tileMap, loop)

	fmt.Printf("Nest area: %v\n", nestArea)
}

func walkLoop(start PipeTile, pipeMap TileMap) (int, PipeLoop) {
	steps := 0
	loop := newPipeLoop(14194) // cap known from solution #1
	currentTile := start
	direction := RIGHT
	next := currentTile.walk(pipeMap, direction)
	loop = append(loop, next)
	currentTile = next
	steps++

	for currentTile != start {
		cameFrom := direction.opposite()
		if currentTile.pipe.a == cameFrom {
			direction = currentTile.pipe.b
		} else {
			direction = currentTile.pipe.a
		}
		next = currentTile.walk(pipeMap, direction)
		loop = append(loop, next)
		currentTile = next
		steps++
	}
	return steps, loop
}

func printLoop(pipeMap TileMap, loop PipeLoop) {
	fmt.Println("Printing the loop shape:")
	fmt.Println()
	pics := pipePics()
	for _, tiles := range pipeMap {
		for _, tile := range tiles {
			if loop.contains(tile) {
				fmt.Printf("%c", pics[tile.pipe])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// calculateNestArea takes map line by line and walks along the "bottom half" of each line
// when it hits a wall (a vertical part of pipe) it starts/stops counting nest tiles
func calculateNestArea(pipeMap TileMap, loop PipeLoop) int {
	nestArea := 0
	pipes := pipesByRune()
	pF, p7, pI := pipes['F'], pipes['7'], pipes['|'] // transitions from/to nest
	var lastPipe Pipe
	inNest, fromLoop, onLoop := false, false, false

	for _, tileLine := range pipeMap {
		for _, currentTile := range tileLine {
			fromLoop = onLoop
			onLoop = loop.contains(currentTile)

			if fromLoop && (lastPipe == pF || lastPipe == p7 || lastPipe == pI) {
				inNest = !inNest
			}

			lastPipe = currentTile.pipe
			if inNest && !onLoop {
				nestArea++
			}
		}
	}
	return nestArea
}

func parse(lines []string, startingPipe Pipe) (TileMap, PipeTile) {
	result := make([][]PipeTile, len(lines))
	startingTile := PipeTile{GridPoint{}, startingPipe}
	pipeTypes := pipesByRune()
	for i, line := range lines {
		result[i] = make([]PipeTile, len(lines[0]))
		for j, char := range line {
			current := PipeTile{GridPoint{j, i}, Pipe{}}
			if char == 'S' {
				current.pipe = startingPipe
				startingTile.pos = current.pos
			} else {
				current.pipe = pipeTypes[char]
			}
			result[i][j] = current
		}
	}
	return result, startingTile
}
