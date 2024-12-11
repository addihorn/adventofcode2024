package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"slices"
	"strconv"
)

const inputFile = "input.txt"

var inputData []string
var maxGenerations = 1
var totalStones = 0

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

	totalStones = 0

	maxGenerations = 25
	for _, stone := range stones {
		evolve(stone, 0)
	}

	fmt.Printf("Solution for part 1: %d\n", totalStones)
}

func evolve(stone, gen int) {

	if gen == maxGenerations {
		totalStones++
		return
	}

	stoneNumberLength := aocutils.OrderOfMagnitude(stone) + 1

	switch true {
	case stone == 0:
		evolve(1, gen+1)
	case stoneNumberLength%2 == 0:
		//split stone
		left := stone / int(math.Pow10(stoneNumberLength/2))
		right := stone % int(math.Pow10(stoneNumberLength/2))

		evolve(left, gen+1)
		evolve(right, gen+1)

	default:
		evolve(stone*2024, gen+1)
	}

}

func part2() {
	//takes to long

	stones := slices.Clone(initialStones)

	totalStones = 0
	maxGenerations = 45
	for n, stone := range stones {
		evolve(stone, 0)

		fmt.Printf("After %d stones: %d\n", n, totalStones)

	}

	fmt.Printf("Solution for part 2: %d\n", totalStones)
}
