package main

import example/hello/src/golang/aocutils

const inputFile = "input_test.txt"

func part1() {
	inputData.ReadInput(inputFile)
	r, _ := regexp.Compile(`\d+`)

	for _, line := range inputData {
		fmt.Println(line)
		matches := r.FindAllString(line, -1)
	}


}

func main() {
	part1()
}