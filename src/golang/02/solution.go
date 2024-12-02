package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

const inputFile = "input.txt"

func isSafe(ele_slice sort.IntSlice) bool {
	if !sort.IsSorted(ele_slice) {
		//check of slice is sorted in decreasing order

		if !sort.IsSorted(sort.Reverse(ele_slice)) {
			// still not sorted in increasing order
			return false
		}
	}

	//at this point we have a sorted slices (increasinf or decreasing)
	for i := 0; i < len(ele_slice)-1; i++ {
		diff := aocutils.Abs(ele_slice[i+1] - ele_slice[i])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func part1() {
	inputData := aocutils.ReadInput(inputFile)

	r, _ := regexp.Compile(`\d+`)

	safe := 0
	for _, line := range inputData {
		matches := r.FindAllString(line, -1)
		num := make([]int, len(matches))
		for i, m := range matches {
			num[i], _ = strconv.Atoi(m)
		}
		ele := sort.IntSlice(num)
		if isSafe(ele) {
			safe++
		}
	}

	fmt.Printf("Solution for part 1: %d\n", safe)

}

func part2() {

	inputData := aocutils.ReadInput(inputFile)

	r, _ := regexp.Compile(`\d+`)

	safe := 0
	for _, line := range inputData {
		matches := r.FindAllString(line, -1)
		num := make([]int, len(matches))
		for i, m := range matches {
			num[i], _ = strconv.Atoi(m)
		}

		if isSafe(sort.IntSlice(num)) {
			safe++
		} else {
			for i := range num {
				new := append(make([]int, 0), num[:i]...)
				new = append(new, num[i+1:]...)
				//fmt.Println(new)
				if isSafe(new) {
					safe++
					break
				}
			}
		}

		//fmt.Printf("[02] After %d rows %d safe \n", i, safe)

	}

	fmt.Printf("Solution for part 2: %d\n", safe)

}

func main() {
	part1()
	part2()
}
