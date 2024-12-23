package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"strconv"
)

const (
	inputFile = "input.txt"
	iter      = 2000
)

var inputData []string

func mixAndPrune(secret, value int) (out int) {
	out = secret ^ value
	//fmt.Println(out)
	out = out % 16777216
	//fmt.Println(out)
	return

}

func solve(secret, iterations int) (out int) {
	originalSecret := secret
	for i := 0; i < iterations; i++ {
		out := secret * 64 //bitwise shift right 6
		//fmt.Println(out)
		secret = mixAndPrune(secret, out)

		out = secret / 32 // bitwise shift left 5
		//fmt.Println(out)
		secret = mixAndPrune(secret, out)

		out = secret * 2048 // bitwise shift  right 11
		//fmt.Println(out)
		secret = mixAndPrune(secret, out)

	}
	fmt.Printf("Solved for %d: %d - binary: %b\n", originalSecret, secret, secret)
	return secret
}

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
	sum := 0
	for _, secret := range inputData {
		secretNum, _ := strconv.Atoi(secret)
		sum += solve(secretNum, iter)
	}
	fmt.Printf("Solution for part 1: %d\n", sum)
}

func part2() {
	_ = inputData
	fmt.Printf("Solution for part 2: %d\n", 2)
}
