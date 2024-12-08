package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"maps"
)

const inputFile = "input.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInput(inputFile)
	initializePuzzle()
	part1()
	part2()
}

/* Do some puzzle initialization */

type antenna struct {
	frequency rune
	X, Y      int
}

var signals = make(map[rune][]antenna)
var gridSize *aocutils.Gridsize

// for visualization
var grid = make(map[[2]int]rune)

func initializePuzzle() {
	gridSize = &aocutils.Gridsize{MinX: 0, MinY: 0, MaxX: len(inputData[0]) - 1, MaxY: len(inputData) - 1}

	for y, line := range inputData {
		for x, position := range line {
			if position == '.' {
				continue
			}

			grid[[2]int{x, y}] = position
			if _, ok := signals[position]; !ok {
				signals[position] = make([]antenna, 0)
			}
			signals[position] = append(signals[position], antenna{frequency: position, X: x, Y: y})
		}
	}
}

/* Solve here */

func part1() {
	p1Grid := maps.Clone(grid)
	antinodes := make(map[[2]int]int)
	for _, antennas := range signals {
		//fmt.Printf("Evaluating frequency %v \n", frequency)

		for i, a1 := range antennas {
			for _, a2 := range antennas[i+1:] {
				difX := a1.X - a2.X
				difY := a1.Y - a2.Y

				if !(a1.X+difX < gridSize.MinX || a1.X+difX > gridSize.MaxX || a1.Y+difY < gridSize.MinY || a1.Y+difY > gridSize.MaxY) {
					antinodes[[2]int{a1.X + difX, a1.Y + difY}]++
					if _, ok := p1Grid[[2]int{a1.X + difX, a1.Y + difY}]; !ok {
						p1Grid[[2]int{a1.X + difX, a1.Y + difY}] = '#'
					}

				}

				if !(a2.X-difX < gridSize.MinX || a2.X-difX > gridSize.MaxX || a2.Y-difY < gridSize.MinY || a2.Y-difY > gridSize.MaxY) {
					antinodes[[2]int{a2.X - difX, a2.Y - difY}]++

					if _, ok := p1Grid[[2]int{a2.X - difX, a2.Y - difY}]; !ok {
						p1Grid[[2]int{a2.X - difX, a2.Y - difY}] = '#'
					}

				}

			}
		}
	}

	//gridSize.PaintGrid(p1Grid)
	fmt.Printf("Solution for part 1: %d\n", len(antinodes))
}

func part2() {
	p2Grid := maps.Clone(grid)
	antinodes := make(map[[2]int]int)
	for _, antennas := range signals {
		//fmt.Printf("Evaluating frequency %v \n", frequency)

		for i, a1 := range antennas {
			for _, a2 := range antennas[i+1:] {

				a1_copy := antenna{frequency: a1.frequency, X: a1.X, Y: a1.Y}

				difX := a1_copy.X - a2.X
				difY := a1_copy.Y - a2.Y

				//also add antenna to antinode-List
				antinodes[[2]int{a1_copy.X, a1_copy.Y}]++
				antinodes[[2]int{a2.X, a2.Y}]++

				for !(a1_copy.X+difX < gridSize.MinX || a1_copy.X+difX > gridSize.MaxX || a1_copy.Y+difY < gridSize.MinY || a1_copy.Y+difY > gridSize.MaxY) {
					antinodes[[2]int{a1_copy.X + difX, a1_copy.Y + difY}]++
					if _, ok := p2Grid[[2]int{a1_copy.X + difX, a1_copy.Y + difY}]; !ok {
						p2Grid[[2]int{a1_copy.X + difX, a1_copy.Y + difY}] = '#'
					}
					//gridSize.PaintGrid(p2Grid)
					a1_copy.X += difX
					a1_copy.Y += difY

				}

				for !(a2.X-difX < gridSize.MinX || a2.X-difX > gridSize.MaxX || a2.Y-difY < gridSize.MinY || a2.Y-difY > gridSize.MaxY) {
					antinodes[[2]int{a2.X - difX, a2.Y - difY}]++

					if _, ok := p2Grid[[2]int{a2.X - difX, a2.Y - difY}]; !ok {
						p2Grid[[2]int{a2.X - difX, a2.Y - difY}] = '#'
					}
					//gridSize.PaintGrid(p2Grid)
					a2.X -= difX
					a2.Y -= difY
				}

			}
		}
	}

	//gridSize.PaintGrid(p2Grid)
	fmt.Printf("Solution for part 2: %d\n", len(antinodes))
}
