package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func parseMul(command string) int {
	reg, _ := regexp.Compile(`\d+`)
	m := reg.FindAllString(command, -1)

	l, _ := strconv.Atoi(m[0])
	r, _ := strconv.Atoi(m[1])

	return l * r
}

func parseInput(input string) int {

	r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)

	matches := r.FindAllString(input, -1)

	sum := 0
	for _, m := range matches {
		sum += parseMul(m)
	}
	return sum
}

func part1() {
	lines := aocutils.ReadInput(inputFile)
	inputAsString := strings.Join(lines, "")

	fmt.Printf("Solution for part 1: %d\n", parseInput(inputAsString))

}

func part2() {

	lines := aocutils.ReadInput(inputFile)
	inputAsString := strings.Join(lines, "")

	// find everything between a don't() and a do()
	// find the information lazu (.*?), otherwise it would potentially overshoot the range
	dontDoReg, _ := regexp.Compile(`don't\(\).*?do\(\)`)

	// delete all parts that a disabled (they don't serve a purpose, so we don't need to handle them later)
	for _, disabledPart := range dontDoReg.FindAllString(inputAsString, -1) {
		inputAsString = strings.Replace(inputAsString, disabledPart, "", 1)
	}

	//make sure the strings doesn't end with a disabled part
	inputAsString = strings.Split(inputAsString, `don't()`)[0]

	fmt.Printf("Solution for part 2: %d\n", parseInput(inputAsString))

}

func main() {
	part1()
	part2()
}
