package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"slices"

	"github.com/gosuri/uilive"
)

const inputFile = "input.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInput(inputFile)

	part1()
	part2()

}

/* Do some puzzle initialization */

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

const (
	FREE = iota
	END
	BLOCKED
)

var grid = make(map[[2]int]rune)
var currentLocation = [2]int{0, 0}
var currentFacing = UP // 0 up, 1 right, 2 down, 3, left

var gridSize *aocutils.Gridsize

func initializePuzzle() {

	gridSize = &aocutils.Gridsize{MinX: 0, MinY: 0, MaxX: len(inputData[0]), MaxY: len(inputData)}

	for y, line := range inputData {
		for x, char := range line {
			grid[[2]int{x, y}] = char

			if char == '^' {
				currentLocation = [2]int{x, y}
				currentFacing = UP
			}

		}
	}

}

/* Solve here */

func pathIsBlocked() (int, [2]int) {

	curX := currentLocation[0]
	curY := currentLocation[1]

	var spaceAhead [2]int

	switch currentFacing {
	case UP:
		spaceAhead = [2]int{curX, curY - 1}
	case DOWN:
		spaceAhead = [2]int{curX, curY + 1}
	case RIGHT:
		spaceAhead = [2]int{curX + 1, curY}
	case LEFT:
		spaceAhead = [2]int{curX - 1, curY}
	default:
	}

	switch grid[spaceAhead] {
	case '#', 'O':
		return BLOCKED, spaceAhead
	case '.', 'X':
		return FREE, spaceAhead
	default:
		return END, spaceAhead
	}

}

func getUniquePositions() (bool, [][2]int) {
	isFree, nextSpace := pathIsBlocked()
	grid[currentLocation] = 'X'
	positions := make([][2]int, 0)
	positions = append(positions, currentLocation)

	type spaceVisit struct {
		position [2]int
		facing   int
	}

	visitedSpaces := []spaceVisit{}

	for !(isFree == END) {

		switch isFree {
		case BLOCKED:
			currentFacing = (currentFacing + 1) % 4
		default:
			currentLocation = nextSpace
			if !(grid[currentLocation] == 'X') {
				positions = append(positions, currentLocation)
			}
			grid[currentLocation] = 'X'

		}

		thisSpace := spaceVisit{position: currentLocation, facing: currentFacing}
		if slices.Contains(visitedSpaces, thisSpace) {
			//we found a circular way
			return false, nil
		}
		visitedSpaces = append(visitedSpaces, thisSpace)

		//gridSize.PaintGrid(grid)

		isFree, nextSpace = pathIsBlocked()

	}

	return true, positions
}

func part1() {
	initializePuzzle()
	_, positions := getUniquePositions()
	fmt.Printf("Solution for part 1: %d\n", len(positions))
}

func part2() {
	writer := uilive.New()
	writer.Start()
	validRoadblocks := 0

	// just bruteforce your way through the puzzle
	initializePuzzle()
	_, firstPositions := getUniquePositions()

	for i, possiblePos := range firstPositions[1:] {

		// reinitialize the puzzle input

		//place obstacle

		initializePuzzle()
		grid[possiblePos] = 'O'
		if f, _ := getUniquePositions(); !f {
			validRoadblocks++
		}

		fmt.Fprintf(writer, "Checked possible positions: %2.2f percent\n", float64(i)/float64(len(firstPositions)-1)*100)

	}
	writer.Stop()
	fmt.Printf("Solution for part 2: %d\n", validRoadblocks)
}
