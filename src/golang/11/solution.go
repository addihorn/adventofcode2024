package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"strconv"
)

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

func simulateBlinks(stones []int, cycles int) int {

	// store only the tally of each stone engraving
	tallyList := make(map[int]int)
	for _, stone := range stones {
		tallyList[stone]++
	}

	for i := 0; i < cycles; i++ {

		nextEvolution := make(map[int]int)

		// each stone engraving behaves the same, so we can bulk-change all of these instances
		for stone, count := range tallyList {
			stoneNumberLength := aocutils.OrderOfMagnitude(stone) + 1
			switch true {
			case stone == 0:
				nextEvolution[1] += count
			case stoneNumberLength%2 == 0:
				//split stone
				left := stone / int(math.Pow10(stoneNumberLength/2))
				right := stone % int(math.Pow10(stoneNumberLength/2))

				nextEvolution[left] += count
				nextEvolution[right] += count
			default:
				nextEvolution[stone*2024] += count
			}
		}
		tallyList = nextEvolution
	}

	// tally up all stone engravings
	total := 0
	for _, count := range tallyList {
		total += count
	}

	return total
}

/* Solve here */

func part1() {
	fmt.Printf("Solution for part 1: %d\n", simulateBlinks(initialStones, 25))
}

func part2() {

	fmt.Printf("Solution for part 2: %d\n", simulateBlinks(initialStones, 75))
}
