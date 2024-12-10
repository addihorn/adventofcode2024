package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"slices"
)

const inputFile = "input.txt"

var inputData []string

var bfs_enables = true

func main(enable_bfs bool) {
	inputData = aocutils.ReadInput(inputFile)
	initializePuzzle()
	part1()
	part2()
}

/* Do some puzzle initialization */

type waypoint struct {
	x, y              int
	elevation, trails int //pt1
	rating            int //pt2
}

var grid = make(map[[2]int]*waypoint)
var summits = make(map[[2]int]*waypoint)
var trailheads = make(map[[2]int]*waypoint)

func initializePuzzle() {
	for y, line := range inputData {
		for x, spot := range line {
			elevation := int(spot - '0')
			if spot == '.' {
				elevation = math.MinInt8
			}
			wp := &waypoint{elevation: elevation, trails: 0, rating: 0, x: x, y: y}

			switch elevation {
			case 9:
				wp.trails = 1
				summits[[2]int{x, y}] = wp
			case 0:
				trailheads[[2]int{x, y}] = wp
			}

			grid[[2]int{x, y}] = wp

		}
	}
	floodFill()
}

/* Solve here */

func floodFill() {
	for _, descend := range summits {
		queue := []*waypoint{descend}
		seen := []*waypoint{}

		for len(queue) > 0 {
			// as queue (BFS)
			waypoint := queue[0]
			queue = queue[1:]

			// as stack (DFS)
			// waypoint := queue[len(queue)-1]
			// queue = queue[:len(queue)-1]

			if !slices.Contains(seen, waypoint) {
				//we have never been here, so add the waypoint to our seen list
				seen = append(seen, waypoint)
				//we found an additional descend to this waypoint
				waypoint.trails++
			}

			//we found an additional route to descend to this waypoint
			waypoint.rating++

			//if elevation is 0, we found a trailhead and don't need descend any further
			if waypoint.elevation == 0 {
				continue
			}
			x, y := waypoint.x, waypoint.y

			//check neigbours
			for _, delta_x := range []int{1, -1} {
				if neigh, ok := grid[[2]int{x + delta_x, y}]; ok && neigh.elevation == waypoint.elevation-1 {
					//add neighbour to queue
					queue = append(queue, neigh)
				}
			}

			for _, delta_y := range []int{1, -1} {
				if neigh, ok := grid[[2]int{x, y + delta_y}]; ok && neigh.elevation == waypoint.elevation-1 {
					//add neighbour to queue
					queue = append(queue, neigh)
				}
			}

		}

	}
}

func part1() {

	score := 0
	for _, trailhead := range trailheads {
		score += trailhead.trails
	}

	fmt.Printf("Solution for part 1: %d\n", score)
}

func part2() {

	score := 0
	for _, trailhead := range trailheads {
		score += trailhead.rating
	}
	fmt.Printf("Solution for part 2: %d\n", score)
}
