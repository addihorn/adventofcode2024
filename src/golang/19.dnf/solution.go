package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"regexp"
	"strings"
)

const debug = false
const inputFile = "input.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInputWithDelimeter(inputFile, "\n\n")
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
	stripes := strings.Split(inputData[0], ", ")
	r := regexp.MustCompile(fmt.Sprintf(`(?m)^(%s)+$`, strings.Join(stripes, "|")))
	matches := r.FindAllString(inputData[1], -1)
	fmt.Printf("Solution for part 1: %d\n", len(matches))
}

var memory map[string]int

func validSubstrings(word string, towels map[string]int) {
	if len(word) == 0 {
		return
	}

	for i := 0; i < len(word); i++ {
		if debug {
			fmt.Printf("Substring = %s %s\n", word[:i+1], word[i+1:])
		}

		substring := word[:i+1]

		if towels[substring] > 0 {
			if count, ok := memory[substring]; ok {
				memory[substring] += max(1, count)
				return
			}

			if count, ok := memory[word[i+1:]]; ok {
				memory[substring+word[i+1:]] += max(1, count)
				return
			}

			if word == substring {
				memory[word] = max(1, memory[word])
				return
			}
			towels[substring]--
			validSubstrings(word[i+1:], towels)
			memory[substring+word[i+1:]] = max(1, memory[word[i+1:]])
			towels[substring]++
		}
	}

}

func part2() {

	sum := 0
	stripes := strings.Split(inputData[0], ", ")
	r := regexp.MustCompile(fmt.Sprintf(`(?m)^(%s)+$`, strings.Join(stripes, "|")))
	matches := r.FindAllString(inputData[1], -1)
	for _, match := range matches {
		memory = map[string]int{}
		possibleStripes := map[string]int{}
		for _, towel := range stripes {
			if count := strings.Count(match, towel); count > 0 {
				possibleStripes[towel] = count
			}
		}

		validSubstrings(match, possibleStripes)
		foo := memory[match]
		sum += foo
		fmt.Printf("Valid solutions for match %s : %d\n", match, foo)

	}

	fmt.Printf("Solution for part 2: %d\n", sum)

}
