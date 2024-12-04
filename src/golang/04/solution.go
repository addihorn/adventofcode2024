package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
)

const inputFile = "input.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInput(inputFile)

	initializePuzzle()

	part1()
	part2()
}

/* initialize data if neccessary here */

var monitor map[[2]int]rune = make(map[[2]int]rune)
var sizeData *aocutils.Gridsize
var xStarts = make([][2]int, 0)
var aStarts = make([][2]int, 0)

func initializePuzzle() {

	for y, row := range inputData {
		line := []rune(row)
		for x, l := range line {
			// all XMAS starts with X, so to save time, we can save, all positions of X for later user
			if l == 'X' {
				xStarts = append(xStarts, [2]int{x, y})
			}
			if l == 'A' {
				aStarts = append(aStarts, [2]int{x, y})
			}
			monitor[[2]int{x, y}] = l
		}
	}

	sizeData = &aocutils.Gridsize{MinX: 0, MinY: 0, MaxX: len(inputData[0]) - 1, MaxY: len(inputData) - 1}

	sizeData.PaintGrid(monitor)
}

/* Solve here */

func part1() {

	sum := 0
	for _, point := range xStarts {
		x := point[0]
		y := point[1]

		// first check boundaries
		left_ok := x-3 >= sizeData.MinX
		right_ok := x+3 <= sizeData.MaxX
		up_ok := y-3 >= sizeData.MinY
		down_ok := y+3 <= sizeData.MaxY

		//check directly left
		if left_ok {
			// add 1 if true
			sum += aocutils.CBool2Int(monitor[[2]int{x - 1, y}] == 'M' && monitor[[2]int{x - 2, y}] == 'A' && monitor[[2]int{x - 3, y}] == 'S')
		}

		// check diagonal up left
		if left_ok && up_ok {
			sum += aocutils.CBool2Int(monitor[[2]int{x - 1, y - 1}] == 'M' && monitor[[2]int{x - 2, y - 2}] == 'A' && monitor[[2]int{x - 3, y - 3}] == 'S')
		}

		// check diagonal down left
		if left_ok && down_ok {
			sum += aocutils.CBool2Int(monitor[[2]int{x - 1, y + 1}] == 'M' && monitor[[2]int{x - 2, y + 2}] == 'A' && monitor[[2]int{x - 3, y + 3}] == 'S')
		}

		//check directly right
		if right_ok {
			sum += aocutils.CBool2Int(monitor[[2]int{x + 1, y}] == 'M' && monitor[[2]int{x + 2, y}] == 'A' && monitor[[2]int{x + 3, y}] == 'S')
		}

		// check diagonal up right
		if right_ok && up_ok {
			sum += aocutils.CBool2Int(monitor[[2]int{x + 1, y - 1}] == 'M' && monitor[[2]int{x + 2, y - 2}] == 'A' && monitor[[2]int{x + 3, y - 3}] == 'S')
		}

		// check diagonal down right
		if right_ok && down_ok {
			sum += aocutils.CBool2Int(monitor[[2]int{x + 1, y + 1}] == 'M' && monitor[[2]int{x + 2, y + 2}] == 'A' && monitor[[2]int{x + 3, y + 3}] == 'S')
		}

		//check directly up
		if up_ok {
			sum += aocutils.CBool2Int(monitor[[2]int{x, y - 1}] == 'M' && monitor[[2]int{x, y - 2}] == 'A' && monitor[[2]int{x, y - 3}] == 'S')
		}

		//check directly down
		if down_ok {
			sum += aocutils.CBool2Int(monitor[[2]int{x, y + 1}] == 'M' && monitor[[2]int{x, y + 2}] == 'A' && monitor[[2]int{x, y + 3}] == 'S')
		}

	}

	fmt.Printf("Solution for part 1: %d\n", sum)
}

func part2() {
	sum := 0

	for _, point := range aStarts {

		x := point[0]
		y := point[1]

		left_ok := x-1 >= sizeData.MinX
		right_ok := x+1 <= sizeData.MaxX
		up_ok := y-1 >= sizeData.MinY
		down_ok := y+1 <= sizeData.MaxY

		//only check of a can be reached as a middle point
		if !(left_ok && right_ok && up_ok && down_ok) {
			continue
		}
		sum += aocutils.CBool2Int(((monitor[[2]int{x - 1, y - 1}] == 'M' && monitor[[2]int{x + 1, y + 1}] == 'S') || (monitor[[2]int{x - 1, y - 1}] == 'S' && monitor[[2]int{x + 1, y + 1}] == 'M')) &&
			((monitor[[2]int{x - 1, y + 1}] == 'M' && monitor[[2]int{x + 1, y - 1}] == 'S') || (monitor[[2]int{x - 1, y + 1}] == 'S' && monitor[[2]int{x + 1, y - 1}] == 'M')))

	}

	fmt.Printf("Solution for part 2: %d\n", sum)
}
