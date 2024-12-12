package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"maps"
	"slices"
	"time"

	"github.com/inancgumus/screen"
)

const debug = false
const inputFile = "bug.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInput(inputFile)
	initializePuzzle()
	solve()
	part1()
	part2()
}

type Cell [2]int

type CellData struct {
	plant, original rune
	//fences
	left, right, up, down bool
}

var garden = make(map[Cell]rune)
var gardenData = make(map[Cell]*CellData)
var gridSize *aocutils.Gridsize

/* Do some puzzle initialization */

func initializePuzzle() {

	gridSize = &aocutils.Gridsize{MinX: 0, MinY: 0, MaxX: len(inputData[0]), MaxY: len(inputData)}

	for y, line := range inputData {
		for x, plant := range line {
			garden[[2]int{x, y}] = plant
			gardenData[Cell{x, y}] = &CellData{plant: plant, original: plant, left: false, right: false, up: false, down: false}
		}
	}
}

func floodGardenPatchFrom(garden map[Cell]rune, startCell Cell) (area, perimeter, sideCount int) {

	perimeter = 0
	area = 0

	queue := []Cell{startCell}
	seen := []Cell{}
	sides := make(map[string]int)

	for len(queue) > 0 {
		cell := queue[0]
		x, y := cell[0], cell[1]
		area++
		seen = append(seen, cell)

		//delete cell from search tree
		queue = queue[1:]

		//check neighbors
		for _, delta_x := range []int{1, -1} {
			neigh, ok := garden[Cell{x + delta_x, y}]
			neigSeen := slices.Contains(seen, Cell{x + delta_x, y})

			if !ok || (neigh != garden[cell] && !neigSeen) {
				perimeter++
				sides[fmt.Sprintf("X %d %d", x, delta_x)]++
				continue
			}
			if !slices.Contains(queue, Cell{x + delta_x, y}) && !neigSeen {
				queue = append(queue, Cell{x + delta_x, y})
			}

		}

		for _, delta_y := range []int{1, -1} {

			neigh, ok := garden[Cell{x, y + delta_y}]
			neigSeen := slices.Contains(seen, Cell{x, y + delta_y})
			if !ok || (neigh != garden[cell] && !neigSeen) {
				perimeter++
				sides[fmt.Sprintf("Y %d %d", y, delta_y)]++
				continue
			}
			if !slices.Contains(queue, Cell{x, y + delta_y}) && !slices.Contains(seen, Cell{x, y + delta_y}) {
				queue = append(queue, Cell{x, y + delta_y})
			}
		}
		garden[cell] = '.'
		if debug {
			PaintGrid(gridSize, garden)
		}

	}
	sideCount = len(sides)

	return

}

func FloodPt2GardenPatchFrom(garden map[Cell]*CellData, startCell Cell) (area, perimeter, edges, corners int) {

	perimeter = 0
	area = 0
	edges = 0
	corners = 0

	queue := []Cell{startCell}
	seen := []Cell{}

	for len(queue) > 0 {

		cell := queue[0]

		if garden[cell].original == 'X' {
			_ = "debug"
		}

		x, y := cell[0], cell[1]
		area++
		seen = append(seen, cell)

		//delete cell from search tree
		queue = queue[1:]

		//check neighbors
		for _, delta_x := range []int{1, -1} {
			neigh, ok := garden[Cell{x + delta_x, y}]
			neigSeen := slices.Contains(seen, Cell{x + delta_x, y})

			if !ok || (neigh.plant != garden[cell].plant && !neigSeen) {
				perimeter++

				down := garden[Cell{x, y + 1}]
				up := garden[Cell{x, y - 1}]

				switch delta_x {
				case 1:
					garden[cell].right = true
					// same edge if
					if !(up != nil && up.right && up.original == garden[cell].original) && !(down != nil && down.right && down.original == garden[cell].original) {
						edges++
					}

				case -1:
					garden[cell].left = true
					if !(up != nil && up.left && up.original == garden[cell].original) && !(down != nil && down.left && down.original == garden[cell].original) {
						edges++
					}
				}
				continue
			}
			if !slices.Contains(queue, Cell{x + delta_x, y}) && !neigSeen {
				queue = append(queue, Cell{x + delta_x, y})
			}

		}

		for _, delta_y := range []int{1, -1} {

			neigh, ok := garden[Cell{x, y + delta_y}]
			neigSeen := slices.Contains(seen, Cell{x, y + delta_y})
			if !ok || (neigh.plant != garden[cell].plant && !neigSeen) {
				perimeter++

				left := garden[Cell{x - 1, y}]
				right := garden[Cell{x + 1, y}]

				switch delta_y {
				case 1:
					garden[cell].down = true
					if !(left != nil && left.down && left.original == garden[cell].original) && !(right != nil && right.down && right.original == garden[cell].original) {
						edges++
					}
				case -1:
					garden[cell].up = true
					if !(left != nil && left.up && left.original == garden[cell].original) && !(right != nil && right.up && right.original == garden[cell].original) {
						edges++
					}
				}
				continue
			}
			if !slices.Contains(queue, Cell{x, y + delta_y}) && !slices.Contains(seen, Cell{x, y + delta_y}) {
				queue = append(queue, Cell{x, y + delta_y})
			}
		}

		for _, delta_x := range []int{1, -1} {
			for _, delta_y := range []int{1, -1} {

				dx := garden[Cell{x + delta_x, y}]
				dy := garden[Cell{x, y + delta_y}]
				dxy := garden[Cell{x + delta_x, y + delta_y}]
				//outer corner (from X)
				/*
					...
					.Xx
					.xx
				*/
				if (dx == nil || dx.original != garden[cell].original) && (dy == nil || dy.original != garden[cell].original) {
					corners++
				}

				//inner corner (from X)
				/*
					X..
					X..
					XXX
				*/

				if (dx != nil && dx.original == garden[cell].original) && (dy != nil && dy.original == garden[cell].original) && (dxy == nil || dxy.original != garden[cell].original) {
					corners++
				}

			}
		}

		garden[cell].plant = '.'

	}

	return

}

var cost_pt1 = 0
var cost_pt2 = 0

/* Solve here */
func solve() {
	gardenCopy := maps.Clone(garden)

	for x := gridSize.MinX; x < gridSize.MaxX+1; x++ {
		for y := gridSize.MinY; y < gridSize.MaxY+1; y++ {
			plant := string(gardenCopy[Cell{x, y}])
			_ = plant
			if !(gardenCopy[Cell{x, y}] == 0 || gardenCopy[Cell{x, y}] == '.') {
				area, perimeter, sides := floodGardenPatchFrom(gardenCopy, Cell{x, y})
				cost_pt1 += area * perimeter
				cost_pt2 += area * sides
			}
		}
	}
}

func part1() {

	fmt.Printf("Solution for part 1: %d\n", cost_pt1)
}

func part2() {

	gardenCopy := maps.Clone(gardenData)
	costByFences := 0
	costByEdges := 0
	costByCorners := 0

	for x := gridSize.MinX; x < gridSize.MaxX; x++ {
		for y := gridSize.MinY; y < gridSize.MaxY; y++ {
			plant := string(gardenCopy[Cell{x, y}].plant)
			_ = plant
			if !(gardenCopy[Cell{x, y}].plant == 0 || gardenCopy[Cell{x, y}].plant == '.') {
				area, perimeter, sides, corners := FloodPt2GardenPatchFrom(gardenCopy, Cell{x, y})

				costByFences += area * perimeter
				costByEdges += area * sides
				costByCorners += area * corners

				if corners != sides {
					fmt.Printf("Found some disrepancy while calculating edges on area %s with starting point %d %d\n", string(plant), x, y)
				}

			}
		}
	}

	fmt.Printf("Solution for part 2: %d %d %d\n", costByFences, costByEdges, costByCorners)
}

func PaintGrid(this *aocutils.Gridsize, grid map[Cell]rune) {
	screen.Clear()
	output := ""
	for y := this.MinY; y <= this.MaxY; y++ {
		for x := this.MinX; x <= this.MaxX; x++ {
			if val, ok := grid[[2]int{x, y}]; ok {
				output = output + string(val)
			} else {
				output = output + "."
			}

		}
		output = output + "\n"
	}
	fmt.Println(output)
	time.Sleep(time.Millisecond * 20)

}
