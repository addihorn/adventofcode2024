package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"slices"
	"strconv"
)

const debug = true
const inputFile = "input.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInputWithDelimeter(inputFile, " ")
	initializePuzzle()
	part1()
	part2()
}

/* Do some puzzle initialization */

var initialStones = []int{}

func initializePuzzle() {
	initialStones = make([]int, len(inputData))
	for i, stone := range inputData {
		initialStones[i], _ = strconv.Atoi(stone)
	}
}

/* Solve here */

func part1() {

	stones := slices.Clone(initialStones)

	sumOfStones := 0
	for n, stone := range stones {
		istones := []int{stone}
		for blinks := 0; blinks < 25; blinks++ {
			offset := 0
			for i, istone := range istones {

				stoneNumberLength := aocutils.OrderOfMagnitude(istone) + 1

				switch true {
				case istone == 0:
					istones[i+offset] = 1
				case stoneNumberLength%2 == 0:
					//split stone
					left := istone / int(math.Pow10(stoneNumberLength/2))
					right := istone % int(math.Pow10(stoneNumberLength/2))
					istones = slices.Concat(istones[:i+offset], []int{left, right}, istones[i+offset+1:])
					offset++
				default:
					istones[i+offset] = istone * 2024
				}
			}
			if debug {
				//fmt.Printf("	Length of stone %d after %d blinks: %d stones\n", n, blinks, len(istones))
			}

		}
		sumOfStones += len(istones)
		if debug {
			fmt.Printf("Stone %d increased to %d stones. \n", n, len(istones))
		}
	}

	fmt.Printf("Solution for part 1: %d\n", sumOfStones)
}

func part2() {
	_ = inputData
	fmt.Printf("Solution for part 2: %d\n", 2)
}
