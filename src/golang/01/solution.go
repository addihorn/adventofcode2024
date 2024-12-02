package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

const inputFile = "input.txt"

func part1() {
	inputData := aocutils.ReadInput(inputFile)
	r, _ := regexp.Compile(`\d+`)

	left := make([]int, len(inputData))
	right := make([]int, len(inputData))

	for index, line := range inputData {
		//fmt.Println(line)
		matches := r.FindAllString(line, -1)
		//fmt.Println(matches)
		left[index], _ = strconv.Atoi(matches[0])
		right[index], _ = strconv.Atoi(matches[1])
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := range left {
		sum += aocutils.Abs(left[i] - right[i])
	}

	fmt.Printf("Solution for part 1: %d\n", sum)

}

func part2() {
	inputData := aocutils.ReadInput(inputFile)
	r, _ := regexp.Compile(`\d+`)

	left := make([]int, len(inputData))
	right := make(map[int]int)
	for i, line := range inputData {
		//fmt.Println(line)
		matches := r.FindAllString(line, -1)
		//fmt.Println(matches)
		num, _ := strconv.Atoi(matches[1])
		left[i], _ = strconv.Atoi(matches[0])
		right[num] += 1
	}

	sum := 0
	for _, num := range left {
		sum += num * right[num]
	}
	fmt.Printf("Solution for part 2: %d\n", sum)

}

func main() {
	part1()
	part2()
}
