package main

import (
	"bytes"
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/inancgumus/screen"
	"github.com/nsf/termbox-go"
)

const (
	demo       = true
	debug      = false
	inputFile  = "1.txt"
	memorySize = 212
	runtime    = 1024
)

var inputData []string
var output = &bytes.Buffer{}

func main() {
	if demo {
		err := termbox.Init()
		if err != nil {
			//fmt.Println(err)
			os.Exit(1)
		}
	}

	inputData = aocutils.ReadInput(inputFile)
	initializePuzzle()
	part1()
	part2()

	termbox.Close()
	fmt.Print(output)
}

type (
	entity      rune
	memorySpace struct {
		position   gridCell
		cellType   entity
		difficulty int
	}
	gridCell [2]int
)

var (
	memory   = make(map[gridCell]*memorySpace)
	gridSize = aocutils.NewGridSize()
	goal     gridCell
	start    gridCell
	player   gridCell
)

const (
	CORRUPT entity = '#'
	START   entity = 'S'
	END     entity = 'E'
	FREE    entity = '.'
	ROUTE   entity = 'O'
)

/* Do some puzzle initialization */

func initializePuzzle() {
	start = gridCell{0, 0}
	player = start
	goal = gridCell{memorySize, memorySize}
	gridSize = &aocutils.Gridsize{MinX: 0, MaxX: memorySize, MinY: 0, MaxY: memorySize}

	for x := 0; x <= memorySize; x++ {
		for y := 0; y <= memorySize; y++ {
			pos := gridCell{x, y}
			memory[pos] = &memorySpace{position: pos, cellType: FREE, difficulty: math.MaxInt}

			switch pos {
			case start:
				memory[pos].cellType = START
				memory[pos].difficulty = 0
			case goal:
				memory[pos].cellType = END
			}
		}
	}
	_ = inputData
	if debug {
		debugGrid(memory, gridSize)
	}

}

func debugGrid(maze map[gridCell]*memorySpace, size *aocutils.Gridsize) {

	if !termbox.IsInit {
		screen.Clear()
		output := ""
		for y := size.MinY; y <= size.MaxY; y++ {
			for x := size.MinX; x <= size.MaxX; x++ {
				if val, ok := maze[[2]int{x, y}]; ok {

					switch true {
					case [2]int{x, y} == player:
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

		time.Sleep(time.Millisecond * 10)
		return
	}

	for y := size.MinY; y <= size.MaxY; y++ {
		for x := size.MinX; x <= size.MaxX; x++ {
			if val, ok := maze[[2]int{x, y}]; ok {

				fg_color := termbox.ColorDefault
				out := val.cellType

				switch true {
				case [2]int{x, y} == player:
					out = '@'
				default:
				}

				switch entity(out) {
				case FREE:
					out = '.'
				case START, END:
					fg_color = termbox.ColorLightGreen
				case ROUTE:
					out = '█'
					fg_color = termbox.ColorLightYellow
				case CORRUPT:
					out = '█'
					fg_color = termbox.ColorRed
				}
				termbox.SetCell(x, y, rune(out), fg_color, termbox.ColorDefault)

			} else {
				termbox.SetCell(x, y, ' ', termbox.ColorLightGreen, termbox.ColorDefault)
			}
		}
	}
	termbox.Flush()

	time.Sleep(time.Nanosecond)

}

func rain() {

	for i, dataPackage := range inputData {
		if i == runtime {
			break
		}

		raindrop(getCellFromDataPackage(dataPackage))

	}
}

func getCellFromDataPackage(dataPackage string) gridCell {
	data := strings.Split(dataPackage, ",")

	x, _ := strconv.Atoi(data[0])
	y, _ := strconv.Atoi(data[1])
	return gridCell{x, y}
}

func raindrop(cell gridCell) {

	memory[cell].cellType = CORRUPT
}

/* Solve here */

func solve() (route []gridCell) {

	route = nil
	// bfs search
	queue := []*memorySpace{memory[start]}
	for len(queue) > 0 {
		toCheck := queue[0]
		queue = queue[1:]
		player = toCheck.position

		if debug {
			debugGrid(memory, gridSize)
		}

		x := toCheck.position[0]
		y := toCheck.position[1]

		for _, dx := range []int{-1, 1} {
			newScore := math.MaxInt
			if neigh, ok := memory[gridCell{x + dx, y}]; ok {
				if memory[gridCell{x + dx, y}].cellType != CORRUPT {
					newScore = toCheck.difficulty + 1
				}
				if newScore < neigh.difficulty {
					neigh.difficulty = newScore
					queue = append(queue, neigh)
				}
			}
		}

		for _, dy := range []int{-1, 1} {
			newScore := math.MaxInt
			if neigh, ok := memory[gridCell{x, y + dy}]; ok {
				if memory[gridCell{x, y + dy}].cellType != CORRUPT {
					newScore = toCheck.difficulty + 1
				}
				if newScore < neigh.difficulty {
					neigh.difficulty = newScore
					queue = append(queue, neigh)
				}
			}

		}
	}

	// do not calculate way back, if the goal was not reached/blocked
	if memory[goal].difficulty == math.MaxInt {
		return
	}

	// reverse way calculation

	route = make([]gridCell, memory[goal].difficulty+1)
	next := goal
	for i := 0; i <= memory[goal].difficulty; i++ {
		route[i] = next
		if memory[next].cellType == FREE {
			memory[next].cellType = ROUTE
		}

		x, y := next[0], next[1]

		for _, dx := range []int{-1, 1} {
			neigh, ok := memory[gridCell{x + dx, y}]
			if !ok {
				continue
			}

			if neigh.difficulty < memory[next].difficulty {
				next = gridCell{x + dx, y}
				break
			}
		}

		for _, dy := range []int{-1, 1} {
			neigh, ok := memory[gridCell{x, y + dy}]
			if !ok {
				continue
			}
			if neigh.difficulty < memory[next].difficulty {
				next = gridCell{x, y + dy}
				break
			}
		}

	}
	if debug || demo {
		debugGrid(memory, gridSize)
	}
	return

}

func part1() {
	rain()
	solve()
	fmt.Fprintf(output, "Solution for part 1: %d\n", memory[goal].difficulty)
}

func reset() {
	player = start
	for x := 0; x <= memorySize; x++ {
		for y := 0; y <= memorySize; y++ {
			pos := gridCell{x, y}
			memory[pos].difficulty = math.MaxInt
			if memory[pos].cellType == ROUTE {
				memory[pos].cellType = FREE
			}

			switch pos {
			case start:
				memory[pos].cellType = START
				memory[pos].difficulty = 0
			case goal:
				memory[pos].cellType = END
			}
		}
	}
}

func part2() {

	reset()

	i := runtime
	route := solve()
	for memory[goal].difficulty < math.MaxInt {

		newDrop := getCellFromDataPackage(inputData[i])
		raindrop(newDrop)
		if slices.Contains(route, newDrop) {
			reset()
			route = solve()
		}
		i++

		/*
			if demo {
				debugGrid(memory, gridSize)
			}
		*/

	}

	if demo {
		reverseDrop := getCellFromDataPackage(inputData[i-1])
		memory[reverseDrop].cellType = FREE
		reset()
		solve()

		time.Sleep(time.Second * 10)
	}

	fmt.Fprintf(output, "Solution for part 2: %d - %s\n", i-1, inputData[i-1])
}
