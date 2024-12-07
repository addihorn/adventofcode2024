package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"strconv"
	"strings"
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

func isSolvable(target int, numbers []int) bool {

	listLength := len(numbers)

	if listLength == 0 {
		return target == 0
	}

	//calculate from the back
	newNumbers := numbers[:listLength-1]
	numberToCheck := numbers[listLength-1]

	//target is devisable by the number (so it could be a *)
	if target%numberToCheck == 0 {
		return isSolvable(target/numberToCheck, newNumbers) || isSolvable(target-numberToCheck, newNumbers)
	}

	return isSolvable(target-numberToCheck, newNumbers)

}

func isSolvableByConcatination(target int, numbers []int) bool {
	listLength := len(numbers)

	if listLength == 1 {
		return isSolvable(target, numbers)
	}
	if listLength == 0 {
		return target == 0
	}

	//calculate from the back
	newNumbers := numbers[:listLength-1]
	numberToCheck := numbers[listLength-1]

	//check if the numberTopCheck is the same as the last n disgits of the target
	numberOfDigits := aocutils.OrderOfMagnitude(numberToCheck) + 1

	concatEvaluation := false
	if target%int(math.Pow(float64(10), float64(numberOfDigits))) == numberToCheck {
		newTarget := target / int(math.Pow(float64(10), float64(numberOfDigits)))
		concatEvaluation = isSolvableByConcatination(newTarget, newNumbers)
	}

	//target is devisable by the number (so it could be a *)
	if target%numberToCheck == 0 {
		return concatEvaluation ||
			isSolvableByConcatination(target/numberToCheck, newNumbers) ||
			isSolvableByConcatination(target-numberToCheck, newNumbers)
	}

	return concatEvaluation || isSolvableByConcatination(target-numberToCheck, newNumbers)
}

func part1() {
	sum := 0
	for _, equation := range inputData {
		e := strings.Split(equation, ": ")
		t := e[0]
		target, _ := strconv.Atoi(t)
		n_str := strings.Split(e[1], " ")
		numbers := make([]int, len(n_str))
		for i, n := range n_str {
			numbers[i], _ = strconv.Atoi(n)
		}

		if isSolvable(target, numbers) {
			//fmt.Printf("[DEBUG] Target %d was solvable \n", target)
			sum += target
		}
	}

	fmt.Printf("Solution for part 1: %d\n", sum)
}

func part2() {
	sum := 0
	for _, equation := range inputData {
		e := strings.Split(equation, ": ")
		t := e[0]
		target, _ := strconv.Atoi(t)
		n_str := strings.Split(e[1], " ")
		numbers := make([]int, len(n_str))
		for i, n := range n_str {
			numbers[i], _ = strconv.Atoi(n)
		}

		if isSolvableByConcatination(target, numbers) {
			//fmt.Printf("[DEBUG] Target %d was solvable \n", target)
			sum += target
		}
	}

	fmt.Printf("Solution for part 2: %d\n", sum)
}
