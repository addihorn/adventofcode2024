package main

import (
	"bytes"
	"example/hello/src/golang/aocutils"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/nsf/termbox-go"
)

const debug = false
const demo = false
const inputFile = "input.txt"

var inputData []string

var output = &bytes.Buffer{}

func main() {
	inputData = aocutils.ReadInputWithDelimeter(inputFile, "\n\n")

	if demo {
		err := termbox.Init()

		if err != nil {
			//fmt.Println(err)
			os.Exit(1)
		}

	}

	initializePuzzle()
	part1()
	initializePuzzleP2()
	part2()

	termbox.Close()
	fmt.Print(output)

}

/* Do some puzzle initialization */

type (
	gridCell  [2]int
	direction rune
	entity    rune
)

var (
	warehouse = make(map[gridCell]entity)
	gridSize  aocutils.Gridsize
	movement  []rune
	robot     gridCell
)

const (
	ROBOT     entity = '@'
	WALL      entity = '#'
	FREE      entity = '.'
	BOX       entity = 'O'
	BOX_LEFT  entity = '['
	BOX_RIGHT entity = ']'
)

const (
	UP    direction = '^'
	DOWN  direction = 'v'
	LEFT  direction = '<'
	RIGHT direction = '>'
)

func initializePuzzle() {

	// generate grid
	whInput := strings.Split(inputData[0], "\n")
	gridSize = aocutils.Gridsize{MaxX: len(whInput[0]) - 1, MaxY: len(whInput) - 1}
	for y, row := range whInput {
		for x, cell := range row {

			warehouse[gridCell{x, y}] = entity(cell)

			if entity(cell) == ROBOT {
				robot = gridCell{x, y}
			}
		}
	}

	if debug {
		aocutils.Paint(warehouse, &gridSize)
	}

	moveInputs := strings.ReplaceAll(inputData[1], "\n", "")
	movement = []rune(moveInputs)
}

func initializePuzzleP2() {
	// generate grid

	whLayout := inputData[0]
	whLayout = strings.ReplaceAll(whLayout, "#", "##")
	whLayout = strings.ReplaceAll(whLayout, "O", "[]")
	whLayout = strings.ReplaceAll(whLayout, ".", "..")
	whLayout = strings.ReplaceAll(whLayout, "@", "@.")

	whInput := strings.Split(whLayout, "\n")
	gridSize = aocutils.Gridsize{MaxX: len(whInput[0]) - 1, MaxY: len(whInput) - 1}
	for y, row := range whInput {
		for x, cell := range row {

			warehouse[gridCell{x, y}] = entity(cell)

			if entity(cell) == ROBOT {
				robot = gridCell{x, y}
			}
		}
	}

	if debug || demo {
		aocutils.Paint(warehouse, &gridSize)
	}

	moveInputs := strings.ReplaceAll(inputData[1], "\n", "")
	movement = []rune(moveInputs)
}

/* Solve here */

func moveRobotFrom(wh map[gridCell]entity, cell gridCell, dir direction) (newRobotPosition gridCell) {

	dx, dy := 0, 0

	switch dir {
	case RIGHT:
		dx = 1
	case LEFT:
		dx = -1
	case DOWN:
		dy = 1
	case UP:
		dy = -1
	}
	newRobotPosition = cell
	end := gridCell{cell[0] + dx, cell[1] + dy}
	pushedBoxes := 0
	//check next free or wall-space in movement direction
	for wh[end] == BOX {
		end[0] += dx
		end[1] += dy
		pushedBoxes++
	}

	switch true {
	case entity(wh[end]) == FREE && pushedBoxes == 0:
		// no boxes in the way, robot can just move
		newRobotPosition = end
		wh[newRobotPosition] = ROBOT
		wh[cell] = FREE
	case entity(wh[end]) == FREE && pushedBoxes > 0:
		newRobotPosition = gridCell{cell[0] + dx, cell[1] + dy}
		wh[newRobotPosition] = ROBOT
		wh[cell] = FREE
		wh[end] = BOX
	default:
		// do nothing since a wall was hit
		if debug && !demo {

			fmt.Println("Robot pushed things into a wall")
		}
	}

	return

}

func getGPS(wh map[gridCell]entity, cell gridCell) int {
	switch wh[cell] {
	case BOX:
		return 100*cell[1] + cell[0]
	default:
		return 0
	}
}

func part1() {

	for i, move := range movement {
		robot = moveRobotFrom(warehouse, robot, direction(move))

		if debug || demo {
			aocutils.Paint(warehouse, &gridSize)
			_ = i
			//fmt.Printf("After %d moves and Move %v", i+1, move)
		}
	}

	gpsSum := 0
	for x := 1; x < gridSize.MaxX; x++ {
		for y := 1; y < gridSize.MaxY; y++ {
			gpsSum += getGPS(warehouse, gridCell{x, y})
		}
	}

	if termbox.IsInit {
		aocutils.TBprint(0, gridSize.MaxY+2, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("Solution for part 1: %d\n", gpsSum))
		termbox.Flush()

	}

	fmt.Fprintf(output, "Solution for part 1: %d\n", gpsSum)
}

func p2MoveRobotFrom(wh map[gridCell]entity, cell gridCell, dir direction) (newRobotPosition gridCell) {

	dx, dy := 0, 0

	switch dir {
	case RIGHT:
		dx = 1
	case LEFT:
		dx = -1
	case DOWN:
		dy = 1
	case UP:
		dy = -1
	}
	newRobotPosition = cell
	end := gridCell{cell[0] + dx, cell[1] + dy}

	if wh[end] == FREE {
		newRobotPosition = end
		wh[newRobotPosition] = ROBOT
		wh[cell] = FREE
		return
	}

	if dir == LEFT || dir == RIGHT {
		for wh[end] == BOX_LEFT || wh[end] == BOX_RIGHT {
			end[0] += dx
		}
		if wh[end] == WALL {
			return
		}

		// here we hit a free space
		spaces := 0
		for i := end[0]; i != cell[0]+dy; i += (dx * -1) {
			switch true {
			case i < cell[0]+dx && spaces%2 == 0:
				wh[gridCell{i, end[1]}] = BOX_LEFT
			case i < cell[0]+dx && spaces%2 != 0:
				wh[gridCell{i, end[1]}] = BOX_RIGHT
			case i > cell[0]+dx && spaces%2 == 0:
				wh[gridCell{i, end[1]}] = BOX_RIGHT
			case i > cell[0]+dx && spaces%2 != 0:
				wh[gridCell{i, end[1]}] = BOX_LEFT
			}
			spaces++

		}
		newRobotPosition = gridCell{cell[0] + dx, cell[1]}
		wh[newRobotPosition] = ROBOT
		wh[cell] = FREE

	}

	if dir == UP || dir == DOWN {
		pushedTiles := checkAndPushBoxes(wh, []gridCell{cell}, dir)
		if debug && !demo {
			fmt.Println(pushedTiles)
		}

		for _, tile := range pushedTiles {
			if wh[tile] == ROBOT {
				newRobotPosition = gridCell{cell[0], cell[1] + dy}
			}

			wh[gridCell{tile[0], tile[1] + dy}] = wh[tile]
			wh[tile] = FREE

		}

	}

	if debug || demo {
		aocutils.Paint(wh, &gridSize)
	}

	return

}

func checkAndPushBoxes(wh map[gridCell]entity, originalCells []gridCell, dir direction) (cellsToMove []gridCell) {

	dy := 0

	switch dir {
	case UP:
		dy = -1
	case DOWN:
		dy = 1
	}

	nextLevel := []gridCell{}

	for _, cell := range originalCells {
		cell[1] += dy
		switch wh[cell] {
		case BOX_LEFT:
			for _, dx := range []int{0, 1} {
				if !slices.Contains(nextLevel, gridCell{cell[0] + dx, cell[1]}) {
					nextLevel = append(nextLevel, gridCell{cell[0] + dx, cell[1]})
				}
			}
		case BOX_RIGHT:
			for _, dx := range []int{0, -1} {
				if !slices.Contains(nextLevel, gridCell{cell[0] + dx, cell[1]}) {
					nextLevel = append(nextLevel, gridCell{cell[0] + dx, cell[1]})
				}
			}
		case WALL:
			//stack is already pushed to a wall, can't push any further
			return []gridCell{}
		}

	}

	if len(nextLevel) == 0 {
		return originalCells
	}

	next := checkAndPushBoxes(wh, nextLevel, dir)
	if len(next) == 0 {
		return []gridCell{}
	}
	return append(next, originalCells...)

}

func p2GetGPS(wh map[gridCell]entity, cell gridCell) int {
	switch wh[cell] {
	case BOX_LEFT:
		return 100*cell[1] + cell[0]
	default:
		return 0
	}
}

func part2() {

	for i, move := range movement {
		robot = p2MoveRobotFrom(warehouse, robot, direction(move))

		if debug || demo {
			aocutils.Paint(warehouse, &gridSize)
			_ = i
			//fmt.Printf("After %d moves and Move %v", i+1, move)
		}
	}

	gpsSum := 0
	for x := 1; x < gridSize.MaxX; x++ {
		for y := 1; y < gridSize.MaxY; y++ {
			gpsSum += p2GetGPS(warehouse, gridCell{x, y})
		}
	}

	if termbox.IsInit {
		aocutils.TBprint(0, gridSize.MaxY+3, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("Solution for part 2: %d\n", gpsSum))
		termbox.Flush()
	}

	fmt.Fprintf(output, "Solution for part 2: %d\n", gpsSum)

}
