package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/inancgumus/screen"
	"github.com/nsf/termbox-go"
)

const (
	debug     = false
	demo      = false
	inputFile = "input.txt"
)

var inputData []string

func main() {

	if demo {
		err := termbox.Init()
		if err != nil {
			//fmt.Println(err)
			os.Exit(1)
		}
	}

	inputData = aocutils.ReadInput(inputFile)
	originalTrack := initializePuzzle()

	part1(originalTrack)
	part2(originalTrack)
	termbox.Close()
	if debug {
		debugGrid(racetrack, gridSize)
	}
	//fmt.Print(output)
}

type (
	entity     rune
	trackSpace struct {
		position    gridCell
		cellType    entity
		difficulty  int
		p2shortCuts map[gridCell]int
	}
	gridCell [2]int
)

var (
	racetrack = make(map[gridCell]*trackSpace)
	gridSize  = aocutils.NewGridSize()
	goal      gridCell
	start     gridCell
	player    gridCell
)

const (
	WALL  entity = '#'
	START entity = 'S'
	END   entity = 'E'
	FREE  entity = '.'
	ROUTE entity = 'O'
)

/* Do some puzzle initialization */

func initializePuzzle() (route []gridCell) {
	start = gridCell{0, 0}
	goal = gridCell{0, 0}
	player = start
	gridSize = &aocutils.Gridsize{MinX: 0, MaxX: len(inputData[0]) - 1, MinY: 0, MaxY: len(inputData) - 1}

	for y, line := range inputData {
		for x, cell := range line {
			pos := gridCell{x, y}
			racetrack[pos] = &trackSpace{position: pos, cellType: entity(cell), difficulty: math.MaxInt, p2shortCuts: map[gridCell]int{}}

			switch entity(cell) {
			case START:
				racetrack[pos].difficulty = 0
				start = pos
				player = pos
			case END:
				goal = pos
			}
		}

	}
	_ = inputData
	if debug {
		debugGrid(racetrack, gridSize)
	}

	// solve the maze first
	route = solveMaze(racetrack)
	if debug {
		fmt.Printf("Length of initial route: %d\n", len(route)-1)
	}
	return

}

func solveMaze(track map[gridCell]*trackSpace) (route []gridCell) {

	route = nil
	// bfs search
	queue := []*trackSpace{track[start]}
	for len(queue) > 0 {
		toCheck := queue[0]
		queue = queue[1:]
		player = toCheck.position

		if debug {
			debugGrid(track, gridSize)
		}

		x := toCheck.position[0]
		y := toCheck.position[1]

		for _, dx := range []int{-1, 1} {
			newScore := math.MaxInt
			if neigh, ok := track[gridCell{x + dx, y}]; ok {
				if track[gridCell{x + dx, y}].cellType != WALL {
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
			if neigh, ok := track[gridCell{x, y + dy}]; ok {
				if track[gridCell{x, y + dy}].cellType != WALL {
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
	if track[goal].difficulty == math.MaxInt {
		return
	}

	// reverse way calculation

	route = make([]gridCell, track[goal].difficulty+1)
	next := goal
	for i := 0; i <= track[goal].difficulty; i++ {

		route[i] = next

		if track[next].cellType == FREE {
			track[next].cellType = ROUTE
		}

		x, y := next[0], next[1]

		for _, dx := range []int{-1, 1} {
			neigh, ok := track[gridCell{x + dx, y}]
			if !ok {
				continue
			}

			if neigh.difficulty < track[next].difficulty {
				next = gridCell{x + dx, y}
				break
			}
		}

		for _, dy := range []int{-1, 1} {
			neigh, ok := track[gridCell{x, y + dy}]
			if !ok {
				continue
			}
			if neigh.difficulty < track[next].difficulty {
				next = gridCell{x, y + dy}
				break
			}
		}

	}

	if debug {
		debugGrid(track, gridSize)
	}

	return

}

func debugGrid(maze map[gridCell]*trackSpace, size *aocutils.Gridsize) {

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
					fg_color = termbox.ColorLightYellow
				case WALL:
					fg_color = termbox.ColorRed
				}
				termbox.SetCell(x, y, rune(out), fg_color, termbox.ColorDefault)

			} else {
				termbox.SetCell(x, y, ' ', termbox.ColorLightGreen, termbox.ColorDefault)
			}
		}
	}
	termbox.Flush()

	time.Sleep(time.Microsecond * 10)

}

/* Solve here */

type shortcut struct {
	from, to gridCell
	save     int
}

func part1(route []gridCell) []shortcut {

	shortCuts := []shortcut{}
	for _, cell := range route {
		x, y := cell[0], cell[1]

		for _, dx := range []int{1, -1} {
			to, to_ok := racetrack[gridCell{x + 2*dx, y}]
			if neigh, ok := racetrack[gridCell{x + dx, y}]; ok && to_ok && neigh.cellType == WALL && to.cellType != WALL {
				s := shortcut{
					from: gridCell{x, y},
					to:   to.position,
					save: to.difficulty - racetrack[gridCell{x, y}].difficulty - 2,
				}

				if s.save > 0 {
					shortCuts = append(shortCuts, s)
				}
			}
		}

		for _, dy := range []int{1, -1} {
			to, to_ok := racetrack[gridCell{x, y + 2*dy}]
			if neigh, ok := racetrack[gridCell{x, y + dy}]; ok && to_ok && neigh.cellType == WALL && to.cellType != WALL {
				s := shortcut{
					from: gridCell{x, y},
					to:   to.position,
					save: to.difficulty - racetrack[gridCell{x, y}].difficulty - 2,
				}

				if s.save > 0 {
					shortCuts = append(shortCuts, s)
				}
			}
		}

	}

	count := 0
	for _, sc := range shortCuts {
		if sc.save > 99 {
			count++
		}
	}

	fmt.Printf("Solution for part 1: %d\n", count)

	return shortCuts
}

func part2(route []gridCell) {

	shortCuts := []shortcut{}

	for _, cell := range route {
		x, y := cell[0], cell[1]

		for dx := 0; dx <= 20; dx++ {
			for dy := 0; dy <= 20-dx; dy++ {
				from := racetrack[gridCell{x, y}]

				if to, ok := racetrack[gridCell{x + dx, y + dy}]; ok && to.cellType != WALL {
					s := shortcut{
						from: from.position,
						to:   to.position,
						save: to.difficulty - from.difficulty - (dx + dy),
					}

					if s.save > 0 && from.p2shortCuts[to.position] == 0 {
						shortCuts = append(shortCuts, s)
						from.p2shortCuts[to.position] = s.save
					}
				}

				if to, ok := racetrack[gridCell{x - dx, y + dy}]; ok && to.cellType != WALL {
					s := shortcut{
						from: gridCell{x, y},
						to:   to.position,
						save: to.difficulty - from.difficulty - (dx + dy),
					}

					if s.save > 0 && from.p2shortCuts[to.position] == 0 {
						shortCuts = append(shortCuts, s)
						from.p2shortCuts[to.position] = s.save
					}
				}

				if to, ok := racetrack[gridCell{x + dx, y - dy}]; ok && to.cellType != WALL {
					s := shortcut{
						from: gridCell{x, y},
						to:   to.position,
						save: to.difficulty - from.difficulty - (dx + dy),
					}

					if s.save > 0 && from.p2shortCuts[to.position] == 0 {
						shortCuts = append(shortCuts, s)
						from.p2shortCuts[to.position] = s.save
					}
				}

				if to, ok := racetrack[gridCell{x - dx, y - dy}]; ok && to.cellType != WALL {
					s := shortcut{
						from: gridCell{x, y},
						to:   to.position,
						save: to.difficulty - from.difficulty - (dx + dy),
					}

					if s.save > 0 && from.p2shortCuts[to.position] == 0 {
						shortCuts = append(shortCuts, s)
						from.p2shortCuts[to.position] = s.save
					}
				}

			}
		}
	}

	count := 0
	for _, sc := range shortCuts {
		if sc.save > 99 {
			count++
		}
	}
	_ = inputData
	fmt.Printf("Solution for part 2: %d\n", count)
}
