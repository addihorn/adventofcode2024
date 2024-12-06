package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"slices"

	"github.com/inancgumus/screen"
)

const inputFile = "input.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInput(inputFile)
	initializePuzzle()
	part2()
	initializePuzzle()
	part1()

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

func getUniquePositions() int {
	isFree, nextSpace := pathIsBlocked()
	grid[currentLocation] = 'X'
	positions := 1
	steps := 1

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
				positions++
			}
			grid[currentLocation] = 'X'
			steps++

		}

		thisSpace := spaceVisit{position: currentLocation, facing: currentFacing}
		if slices.Contains(visitedSpaces, thisSpace) {
			//we found a circular way
			return -1
		}
		visitedSpaces = append(visitedSpaces, thisSpace)

		//gridSize.PaintGrid(grid)

		isFree, nextSpace = pathIsBlocked()

	}

	return positions
}

func part1() {

	fmt.Printf("Solution for part 1: %d\n", getUniquePositions())
}

func part2() {

	validRoadblocks := 0

	// just bruteforce your way through the puzzle
	for x := 0; x < gridSize.MaxX+1; x++ {
		for y := 0; y < gridSize.MaxY+1; y++ {
			// reinitialize the puzzle input
			initializePuzzle()

			//place obstacle
			if grid[[2]int{x, y}] == '.' {
				grid[[2]int{x, y}] = 'O'
			} else {
				continue
			}

			if getUniquePositions() == -1 {
				validRoadblocks++
			}
			screen.Clear()
			fmt.Printf("Checked positions of overall grid: %.2f percent \n", float64(x*gridSize.MaxY+y)/float64(gridSize.MaxX*gridSize.MaxY)*100)

		}
	}

	fmt.Printf("Solution for part 2: %d\n", validRoadblocks)
}
