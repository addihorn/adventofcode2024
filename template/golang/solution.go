package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
)

const inputFile = "input.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInput(inputFile)
	initializePuzzle()
	part1()
	part2()
}

/* Do some puzzle initialization */

func initializePuzzle() {
	_ = inputData
}

/* Solve here */

func part1() {
	_ = inputData
	fmt.Printf("Solution for part 1: %d\n", 1)
}

func part2() {
	_ = inputData
	fmt.Printf("Solution for part 2: %d\n", 2)
}
