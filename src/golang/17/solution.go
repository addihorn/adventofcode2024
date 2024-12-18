package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const debug = false
const inputFile = "1.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInputWithDelimeter(inputFile, "\n\n")
	initializePuzzle()
	part1()
	part2()
}

/* Do some puzzle initialization */

type (
	instructionSet int64
	operand        int64
	memory         struct {
		register_a, register_b, register_c int64
	}
)

var (
	mem     = memory{}
	program []int64
)

const (
	adv instructionSet = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

const (
	pointer_a operand = 4
	pointer_b operand = 5
	pointer_c operand = 6
)

func initializePuzzle() {
	r, _ := regexp.Compile(`Register (.): (\d+)`)

	matches := r.FindAllStringSubmatch(inputData[0], -1)
	for _, match := range matches {
		switch match[1] {
		case "A":
			mem.register_a, _ = strconv.ParseInt(match[2], 10, 64)
		case "B":
			mem.register_b, _ = strconv.ParseInt(match[2], 10, 64)
		case "C":
			mem.register_c, _ = strconv.ParseInt(match[2], 10, 64)

		}

	}

	r, _ = regexp.Compile(`\d+`)
	matches = r.FindAllStringSubmatch(inputData[1], -1)
	program = make([]int64, len(matches))
	for i, match := range matches {
		program[i], _ = strconv.ParseInt(match[0], 10, 64)
	}

	if debug {
		fmt.Printf("Memory State: %+v\n", mem)
		fmt.Printf("Program: %+v\n", program)
	}

}

func division(numerator, divisor int64) int64 {

	return numerator / int64(math.Pow(2, float64(divisor)))

}

func solve() string {
	inst_pointer := 0
	output := []string{}
	for inst_pointer < len(program) {

		instruction := instructionSet(program[inst_pointer])
		combo_value := program[inst_pointer+1]
		liteal_value := combo_value

		switch operand(combo_value) {
		case pointer_a:
			combo_value = mem.register_a
		case pointer_b:
			combo_value = mem.register_b
		case pointer_c:
			combo_value = mem.register_c
		default:
			// value can be treated as literal
		}

		switch instructionSet(instruction) {
		case adv:
			mem.register_a = division(mem.register_a, combo_value)
		case bxl:
			mem.register_b = mem.register_b ^ liteal_value
		case bst:
			mem.register_b = combo_value % 8
		case jnz:
			if mem.register_a != 0 {
				inst_pointer = int(liteal_value)
				continue
			}
		case bxc:
			mem.register_b = mem.register_b ^ mem.register_c
		case out:
			output = append(output, fmt.Sprint(combo_value%8))
		case bdv:
			mem.register_b = division(mem.register_a, combo_value)
		case cdv:
			mem.register_c = division(mem.register_a, combo_value)

		}

		inst_pointer += 2

	}
	return strings.Join(output, ",")
}

/* Solve here */

func part1() {

	fmt.Printf("Solution for part 1: %s\n", solve())
}

func part2() {

	possibleProgram := ""
	intialA := int64(1)
	matchingNum := 1
	totalNumbersChecked := 0
	for fmt.Sprintf("Program: %s", possibleProgram) != inputData[1] && len(possibleProgram) < len(program)*2 {

		mem.register_a = int64(intialA)
		mem.register_b = 0
		mem.register_c = 0

		possibleProgram = solve()
		if debug {
			fmt.Printf("Program with inital A %d : %s\n", intialA, possibleProgram)
		}

		if possibleProgram[len(possibleProgram)-matchingNum:] == inputData[1][len(inputData[1])-matchingNum:] {
			intialA *= 8
			matchingNum += 2
		} else {
			intialA += 1
		}
		totalNumbersChecked++
	}

	if fmt.Sprintf("Program: %s", possibleProgram) != inputData[1] {
		fmt.Printf("No solution for Part2 found")
	} else {
		fmt.Printf("Solution for part 2: %d, total Numbers checked: %d\n", intialA/8, totalNumbersChecked)
	}

}
