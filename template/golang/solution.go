package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
)

const inputFile = "input.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInput(inputFile)
	part1()
	part2()
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
