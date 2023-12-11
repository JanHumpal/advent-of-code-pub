package main

import (
	"fmt"
	"utl"
)

func main() {
	solve()
}

const expansionRate = 1000 * 1000
const galaxyRune = '#'

type Galaxy struct {
	x, y uint64
}

type Universe struct {
	galaxies     []Galaxy
	emptyColumns []uint64
	emptyRows    []uint64
}

func (u Universe) getManhattan(a Galaxy, b Galaxy) uint64 {
	dX, dY := dist(a.x, b.x), dist(a.y, b.y)
	extraColumns := u.getExtraCs(a.x, b.x)
	extraRows := u.getExtraRs(a.y, b.y)
	dX += extraColumns * (expansionRate - 1)
	dY += extraRows * (expansionRate - 1)
	return dX + dY
}

func (u Universe) getExtraCs(x1 uint64, x2 uint64) uint64 {
	if x1 < x2 {
		return countBetween(x1, x2, u.emptyColumns)
	} else {
		return countBetween(x2, x1, u.emptyColumns)
	}
}

func (u Universe) getExtraRs(y1 uint64, y2 uint64) uint64 {
	if y1 < y2 {
		return countBetween(y1, y2, u.emptyRows)
	} else {
		return countBetween(y2, y1, u.emptyRows)
	}
}

func countBetween(lBound uint64, uBound uint64, values []uint64) uint64 {
	result := uint64(0)
	for _, value := range values {
		if value > lBound && value < uBound {
			result++
		}
	}
	return result
}

func dist(a uint64, b uint64) uint64 {
	if a < b {
		return b - a
	}
	return a - b
}

func solve() {
	input := utl.ReadInput("./input")
	fmt.Printf("Number of lines: %v\n", len(input))

	totalDistance := uint64(0)
	universe := parse(input)
	galaxies := universe.galaxies
	for i, galaxy := range galaxies {
		for j := len(galaxies) - 1; j > i; j-- {
			totalDistance += universe.getManhattan(galaxy, galaxies[j])
		}
	}

	fmt.Printf("Total: %v\n", totalDistance)
}

func parse(input []string) Universe {
	result := Universe{make([]Galaxy, 0), make([]uint64, 0), make([]uint64, 0)}
	takenColumns := make([]bool, len(input[0]))
	for y, line := range input {
		lineClear := true
		for x, char := range line {
			if char == galaxyRune {
				result.galaxies = append(result.galaxies, Galaxy{uint64(x), uint64(y)})
				lineClear = false
				takenColumns[x] = true
			}
		}
		if lineClear {
			result.emptyRows = append(result.emptyRows, uint64(y))
		}
	}
	for i, taken := range takenColumns {
		if !taken {
			result.emptyColumns = append(result.emptyColumns, uint64(i))
		}
	}
	return result
}
