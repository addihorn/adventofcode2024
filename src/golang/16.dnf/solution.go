package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"slices"
	"time"

	"github.com/inancgumus/screen"
)

const debug = false
const inputFile = "input.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInput(inputFile)
	initializePuzzle()
	part1()
	part2()
}

/* Do some puzzle initialization */

type (
	entity              rune
	facing              int
	mazeCellInformation struct {
		position      gridCell
		cellType      entity
		difficulty    int
		minimalFacing facing
		steps, turns  int
	}
	reindeerInformation struct {
		position  gridCell
		lookingTo facing
	}
	gridCell [2]int
)

var (
	maze     = make(map[gridCell]*mazeCellInformation)
	gridSize = aocutils.NewGridSize()
	goal     gridCell
	reindeer *reindeerInformation
)

const (
	WALL  entity = '#'
	START entity = 'S'
	END   entity = 'E'
	FREE  entity = '.'
	SEAT  entity = 'O'
)

const (
	WEST facing = iota
	SOUTH
	EAST
	NORTH
)

func initializePuzzle() {

	gridSize = &aocutils.Gridsize{MinX: 0, MinY: 0, MaxX: len(inputData[0]) - 1, MaxY: len(inputData) - 1}

	for y, line := range inputData {
		for x, cell := range line {
			mazeCell := &mazeCellInformation{position: gridCell{x, y}, cellType: entity(cell), difficulty: math.MaxInt, minimalFacing: -1}

			switch entity(cell) {
			case START:
				mazeCell.difficulty = 0
				mazeCell.minimalFacing = EAST
				reindeer = &reindeerInformation{position: gridCell{x, y}, lookingTo: EAST}

			case END:
				goal = gridCell{x, y}
			}

			maze[gridCell{x, y}] = mazeCell
		}
	}

	if debug {
		debugGrid(maze, gridSize)
	}
}

func debugGrid(maze map[gridCell]*mazeCellInformation, size *aocutils.Gridsize) {
	screen.Clear()
	output := ""
	for y := size.MinY; y <= size.MaxY; y++ {
		for x := size.MinX; x <= size.MaxX; x++ {
			if val, ok := maze[[2]int{x, y}]; ok {

				switch true {
				case [2]int{x, y} == reindeer.position:
					output += "@"
				default:
					output = output + string(val.cellType)
				}

			} else {
				output = output + "."
			}

		}
		output = output + "\n"
	}
	fmt.Println(output)
	fmt.Printf("Rudolf: %+v\n", reindeer)

	time.Sleep(time.Millisecond * 10)

}

/* Solve here */

func part1() {

	stack := []*mazeCellInformation{maze[reindeer.position]}

	for len(stack) > 0 {
		toCheck := stack[0]
		stack = stack[1:]

		reindeer.lookingTo = toCheck.minimalFacing
		reindeer.position = toCheck.position

		if reindeer.position == goal {
			continue
		}
		if debug {
			debugGrid(maze, gridSize)
		}

		x, y := toCheck.position[0], toCheck.position[1]

		neighbors := []*mazeCellInformation{}
		//west
		if neigh := maze[gridCell{x - 1, y}]; entity(neigh.cellType) != WALL {

			newScore := math.MaxInt
			turns, steps := toCheck.turns, toCheck.steps
			switch reindeer.lookingTo {
			case NORTH, SOUTH:
				newScore = toCheck.difficulty + 1001
				turns += 1
				steps += 1
			case EAST:
				newScore = toCheck.difficulty + 2001
				turns += 2
				steps += 1
			default:
				newScore = toCheck.difficulty + 1
				steps += 1

			}
			if newScore <= neigh.difficulty && newScore <= maze[goal].difficulty {
				neigh.difficulty = newScore
				neigh.minimalFacing = WEST
				neigh.steps = steps
				neigh.turns = turns
				neighbors = append(neighbors, neigh)
			}
		}

		//east
		if neigh := maze[gridCell{x + 1, y}]; entity(neigh.cellType) != WALL {

			newScore := math.MaxInt
			turns, steps := toCheck.turns, toCheck.steps
			switch reindeer.lookingTo {
			case NORTH, SOUTH:
				newScore = toCheck.difficulty + 1001
				turns += 1
				steps += 1
			case WEST:
				newScore = toCheck.difficulty + 2001
				turns += 2
				steps += 1
			default:
				newScore = toCheck.difficulty + 1
				steps += 1

			}
			if newScore <= neigh.difficulty && newScore <= maze[goal].difficulty {
				neigh.difficulty = newScore
				neigh.minimalFacing = EAST
				neigh.steps = steps
				neigh.turns = turns
				neighbors = append(neighbors, neigh)
			}
		}

		//south
		if neigh := maze[gridCell{x, y + 1}]; entity(neigh.cellType) != WALL {

			newScore := math.MaxInt
			turns, steps := toCheck.turns, toCheck.steps
			switch reindeer.lookingTo {
			case WEST, EAST:
				newScore = toCheck.difficulty + 1001
				turns += 1
				steps += 1
			case NORTH:
				newScore = toCheck.difficulty + 2001
				turns += 2
				steps += 1
			default:
				newScore = toCheck.difficulty + 1
				steps += 1

			}
			if newScore <= neigh.difficulty && newScore <= maze[goal].difficulty {
				neigh.difficulty = newScore
				neigh.minimalFacing = SOUTH
				neigh.steps = steps
				neigh.turns = turns
				neighbors = append(neighbors, neigh)
			}
		}

		//north
		if neigh := maze[gridCell{x, y - 1}]; entity(neigh.cellType) != WALL {

			newScore := math.MaxInt
			turns, steps := toCheck.turns, toCheck.steps
			switch reindeer.lookingTo {
			case WEST, EAST:
				newScore = toCheck.difficulty + 1001
				turns += 1
				steps += 1
			case SOUTH:
				newScore = toCheck.difficulty + 2001
				turns += 2
				steps += 1
			default:
				newScore = toCheck.difficulty + 1
				steps += 1

			}
			if newScore <= neigh.difficulty && newScore <= maze[goal].difficulty {
				neigh.difficulty = newScore
				neigh.minimalFacing = NORTH
				neigh.steps = steps
				neigh.turns = turns
				neighbors = append(neighbors, neigh)
			}
		}

		stack = append(stack, neighbors...)

	}

	fmt.Printf("Solution for part 1: %+v\n", maze[goal].difficulty)
}

func part2() {

	for pos, cell := range maze {
		if cell.cellType == WALL {
			maze[pos].steps = maze[goal].steps
			maze[pos].turns = maze[goal].turns
		}
	}

	queue := []gridCell{goal}
	seen := []gridCell{}
	sum := 0
	for len(queue) > 0 {

		toCheck := queue[0]
		queue = queue[1:]
		if slices.Contains(seen, toCheck) {
			continue
		}

		sum++

		seen = append(seen, toCheck)

		if maze[toCheck].cellType == START {
			continue
		}

		maze[toCheck].cellType = SEAT

		x, y := toCheck[0], toCheck[1]
		//west

		for _, dx := range []int{-1, 1} {
			foo := maze[gridCell{x + dx, y}]
			if foo.difficulty < maze[goal].difficulty && foo.cellType != SEAT && (foo.difficulty < maze[toCheck].difficulty || foo.difficulty == maze[toCheck].difficulty-1001 || foo.difficulty == maze[toCheck].difficulty-2001 || foo.difficulty == maze[toCheck].difficulty+999) {
				queue = append(queue, foo.position)
			}
		}

		for _, dy := range []int{-1, 1} {
			foo := maze[gridCell{x, y + dy}]
			if foo.difficulty < maze[goal].difficulty && foo.cellType != SEAT && (foo.difficulty < maze[toCheck].difficulty || foo.difficulty == maze[toCheck].difficulty+999) {
				queue = append(queue, foo.position)
			}
		}

		if debug {
			debugGrid(maze, gridSize)
		}

	}
	//debugGrid(maze, gridSize)
	fmt.Printf("Solution for part 2: %d\n", sum)

}
