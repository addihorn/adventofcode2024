package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

const inputFile = "input.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInputWithDelimeter(inputFile, "\n\n")
	initializePuzzle()
	part1()
	part2()
}

/* Do some puzzle initialization */

type clawMachine struct {
	prize_x, prize_y       float64
	button_a_x, button_a_y float64
	button_b_x, button_b_y float64
	p1_moves_a, p1_moves_b float64
	price                  float64
}

var machines []clawMachine

func initializePuzzle() {
	_ = inputData
}

/* Solve here */

func part1() {
	tokens := 0
	machines = make([]clawMachine, len(inputData))

	for i, machineInfo := range inputData {
		r, _ := regexp.Compile(`(.+):.+?(\d+).+?(\d+)`)
		matches := r.FindAllStringSubmatch(machineInfo, -1)
		_, _ = i, matches
		machine := clawMachine{}
		for _, match := range matches {
			x, _ := strconv.Atoi(match[2])
			y, _ := strconv.Atoi(match[3])
			switch match[1] {
			case "Prize":
				machine.prize_x, machine.prize_y = float64(x), float64(y)
			case "Button A":
				machine.button_a_x, machine.button_a_y = float64(x), float64(y)
			case "Button B":
				machine.button_b_x, machine.button_b_y = float64(x), float64(y)
			}
		}

		machine.p1_moves_b = ((machine.button_a_x * machine.prize_y) - (machine.button_a_y * machine.prize_x)) / ((machine.button_a_x * machine.button_b_y) - (machine.button_a_y * machine.button_b_x))
		machine.p1_moves_a = (machine.prize_x - (machine.button_b_x * machine.p1_moves_b)) / machine.button_a_x

		machine.price = float64(machine.p1_moves_a*3 + machine.p1_moves_b*1)

		if machine.price == math.Trunc(machine.price) {
			tokens += int(machine.price)
		}

		machines[i] = machine

	}
	fmt.Printf("Solution for part 1: %d\n", tokens)
}

func part2() {
	tokens := 0
	machines = make([]clawMachine, len(inputData))

	for i, machineInfo := range inputData {
		r, _ := regexp.Compile(`(.+):.+?(\d+).+?(\d+)`)
		matches := r.FindAllStringSubmatch(machineInfo, -1)
		_, _ = i, matches
		machine := clawMachine{}
		for _, match := range matches {
			x, _ := strconv.Atoi(match[2])
			y, _ := strconv.Atoi(match[3])
			switch match[1] {
			case "Prize":
				machine.prize_x, machine.prize_y = float64(x+10000000000000), float64(y+10000000000000)
			case "Button A":
				machine.button_a_x, machine.button_a_y = float64(x), float64(y)
			case "Button B":
				machine.button_b_x, machine.button_b_y = float64(x), float64(y)
			}
		}

		machine.p1_moves_b = ((machine.button_a_x * machine.prize_y) - (machine.button_a_y * machine.prize_x)) / ((machine.button_a_x * machine.button_b_y) - (machine.button_a_y * machine.button_b_x))
		machine.p1_moves_a = (machine.prize_x - (machine.button_b_x * machine.p1_moves_b)) / machine.button_a_x

		machine.price = float64(machine.p1_moves_a*3 + machine.p1_moves_b*1)

		if machine.price == math.Trunc(machine.price) {
			tokens += int(machine.price)
		}

		machines[i] = machine

	}
	fmt.Printf("Solution for part 2: %d\n", tokens)
}
