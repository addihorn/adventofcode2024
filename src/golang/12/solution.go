package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"maps"
	"math/rand"
	"os"
	"slices"
)

const debug = true
const inputFile = "xmas.txt"

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

var colorPallet = map[int]color.RGBA{
	0:  color.RGBA{R: 102, G: 194, B: 165, A: 1},
	1:  color.RGBA{R: 252, G: 141, B: 98, A: 1},
	2:  color.RGBA{R: 141, G: 160, B: 203, A: 1},
	3:  color.RGBA{R: 231, G: 138, B: 195, A: 1},
	4:  color.RGBA{R: 166, G: 216, B: 84, A: 1},
	5:  color.RGBA{R: 255, G: 217, B: 47, A: 1},
	6:  color.RGBA{R: 229, G: 196, B: 148, A: 1},
	7:  color.RGBA{R: 179, G: 179, B: 179, A: 1},
	8:  color.RGBA{R: 0, G: 0, B: 0, A: 1},
	9:  color.RGBA{R: 158, G: 1, B: 66, A: 1},
	10: color.RGBA{R: 213, G: 62, B: 79, A: 1},
	11: color.RGBA{R: 244, G: 109, B: 67, A: 1},
	12: color.RGBA{R: 253, G: 174, B: 97, A: 1},
	13: color.RGBA{R: 254, G: 224, B: 139, A: 1},
	14: color.RGBA{R: 255, G: 255, B: 191, A: 1},
	15: color.RGBA{R: 230, G: 245, B: 152, A: 1},
	16: color.RGBA{R: 171, G: 221, B: 164, A: 1},
	17: color.RGBA{R: 102, G: 194, B: 165, A: 1},
	18: color.RGBA{R: 50, G: 136, B: 189, A: 1},
	19: color.RGBA{R: 94, G: 79, B: 162, A: 1},
}

var garden = make(map[Cell]rune)
var gardenData = make(map[Cell]*CellData)
var gridSize *aocutils.Gridsize
var colorGrid = make(map[rune]color.RGBA)

/* Do some puzzle initialization */

func initializePuzzle() {
	takenColors := []int{}
	gridSize = &aocutils.Gridsize{MinX: 0, MinY: 0, MaxX: len(inputData[0]), MaxY: len(inputData)}

	for y, line := range inputData {
		for x, plant := range line {
			garden[[2]int{x, y}] = plant
			gardenData[Cell{x, y}] = &CellData{plant: plant, original: plant, left: false, right: false, up: false, down: false}

			if len(takenColors) == len(colorPallet) {
				//reset taken colors
				takenColors = []int{}
			}

			if _, ok := colorGrid[plant]; !ok {
				r := rand.Intn(len(colorPallet))
				for slices.Contains(takenColors, r) {
					r = rand.Intn(len(colorPallet))
				}
				takenColors = append(takenColors, r)
				colorGrid[plant] = colorPallet[r]
			}

		}
	}
}

var images = []*image.Paletted{}
var delays = []int{}

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

		if debug {

			images = append(images, PaintGrid(gridSize, garden))
			delays = append(delays, 2)

		}

	}

	return

}

var cost_pt1 = 0
var cost_pt2 = 0

/* Solve here */
func solve() {
	gardenCopy := maps.Clone(gardenData)
	cellsSeen := 0
	maxCells := gridSize.MaxX * gridSize.MaxY

	for x := gridSize.MinX; x < gridSize.MaxX; x++ {
		for y := gridSize.MinY; y < gridSize.MaxY; y++ {
			cellsSeen++
			plant := string(gardenCopy[Cell{x, y}].plant)
			_ = plant
			if !(gardenCopy[Cell{x, y}].plant == 0 || gardenCopy[Cell{x, y}].plant == '.') {
				area, perimeter, sides, corners := FloodPt2GardenPatchFrom(gardenCopy, Cell{x, y})

				cost_pt1 += area * perimeter
				//costByEdges += area * sides
				cost_pt2 += area * corners

				if corners != sides {
					fmt.Printf("Found some disrepancy while calculating edges on area %s with starting point %d %d\n", string(plant), x, y)
				}

			}

			if debug && cellsSeen%100 == 0 {
				fmt.Printf("Seen cells: %d / %d\n", cellsSeen, maxCells)
			}

		}
	}

}

func part1() {

	fmt.Printf("Solution for part 1: %d\n", cost_pt1)
}

func part2() {

	f, _ := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image:     images,
		Delay:     delays,
		LoopCount: -1,
	})

	fmt.Printf("Solution for part 2: %d\n", cost_pt2)
}

func PaintGrid(this *aocutils.Gridsize, grid map[Cell]*CellData) (canvas *image.Paletted) {

	var palette = []color.Color{
		color.RGBA{R: 102, G: 194, B: 165, A: 1},
		color.RGBA{R: 252, G: 141, B: 98, A: 1},
		color.RGBA{R: 141, G: 160, B: 203, A: 1},
		color.RGBA{R: 231, G: 138, B: 195, A: 1},
		color.RGBA{R: 166, G: 216, B: 84, A: 1},
		color.RGBA{R: 255, G: 217, B: 47, A: 1},
		color.RGBA{R: 229, G: 196, B: 148, A: 1},
		color.RGBA{R: 179, G: 179, B: 179, A: 1},
		color.RGBA{R: 0, G: 0, B: 0, A: 1},
		color.White,
		color.RGBA{R: 158, G: 1, B: 66, A: 1},
		color.RGBA{R: 213, G: 62, B: 79, A: 1},
		color.RGBA{R: 244, G: 109, B: 67, A: 1},
		color.RGBA{R: 253, G: 174, B: 97, A: 1},
		color.RGBA{R: 254, G: 224, B: 139, A: 1},
		color.RGBA{R: 255, G: 255, B: 191, A: 1},
		color.RGBA{R: 230, G: 245, B: 152, A: 1},
		color.RGBA{R: 171, G: 221, B: 164, A: 1},
		color.RGBA{R: 102, G: 194, B: 165, A: 1},
		color.RGBA{R: 50, G: 136, B: 189, A: 1},
		color.RGBA{R: 94, G: 79, B: 162, A: 1},
	}

	canvas = image.NewPaletted(image.Rectangle{image.Point{0, 0}, image.Point{gridSize.MaxX, gridSize.MaxY}}, palette)
	for y := this.MinY; y <= this.MaxY; y++ {
		for x := this.MinX; x <= this.MaxX; x++ {
			if val, ok := grid[[2]int{x, y}]; ok {
				if val.plant == '.' {
					canvas.Set(x, y, colorGrid[val.original])
				} else {
					canvas.Set(x, y, color.White)
				}

			} else {
				canvas.Set(x, y, color.White)
			}

		}
	}
	return

}
