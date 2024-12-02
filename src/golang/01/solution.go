package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"regexp"
)

const inputFile = "input_test.txt"

func part1() {
	inputData := aocutils.ReadInput(inputFile)
	r, _ := regexp.Compile(`\d+`)

	for _, line := range inputData {
		fmt.Println(line)
		matches := r.FindAllString(line, -1)
		fmt.Println(matches)
	}

}

func main() {
	part1()
}
