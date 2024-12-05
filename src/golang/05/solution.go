package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInputWithDelimeter(inputFile, "\n\n")
	initializePuzzle()
	part1()
	part2() //depends on part1
}

// hashmap to show all devendents to a given page
var pageOrder = make(map[int][]int)
var updateList []string
var nokList = make([]int, 0)

func initializePuzzle() {

	pageOrderAsList := strings.Split(inputData[0], "\n")

	for _, pageRule := range pageOrderAsList {
		rule := strings.Split(pageRule, "|")

		firstPage, _ := strconv.Atoi(rule[0])
		secondPage, _ := strconv.Atoi(rule[1])

		if _, ok := pageOrder[firstPage]; !ok {
			pageOrder[firstPage] = make([]int, 0)
		}
		pageOrder[firstPage] = append(pageOrder[firstPage], secondPage)
	}

	updateList = strings.Split(inputData[1], "\n")

}

/* Solve here */

func pagesAreOkay(pages []string) bool {

	for i, page := range pages[1:] {
		pageNo, _ := strconv.Atoi(page)
		for _, pageBeforeActivePage := range pages[:i+1] {
			pageNo_before, _ := strconv.Atoi(pageBeforeActivePage)

			if slices.Contains(pageOrder[pageNo], pageNo_before) {
				return false
			}

		}
	}
	return true
}

func part1() {
	sum := 0
	for p, pagesToUpdate := range updateList {
		_ = p // for debugging
		pages := strings.Split(pagesToUpdate, ",")

		if pagesAreOkay(pages) {

			//add middle Page number to sum
			middleNumberAsString := pages[len(pages)/2]
			middleNumber, _ := strconv.Atoi(middleNumberAsString)
			//fmt.Printf("[DEBUG] Adding Page Number %d from set %d\n", middleNumber, p)
			sum += middleNumber

		} else {
			nokList = append(nokList, p)
		}

	}
	fmt.Printf("Solution for part 1: %d\n", sum)
}

func part2() {

	sum := 0
	for _, p := range nokList {

		pages := strings.Split(updateList[p], ",")
		newPageList := []string{pages[0]}
		for _, page := range pages[1:] {

			for i := range newPageList {

				// slide the page number through the array, until it is at a valid position
				checkList := []string{}
				checkList = append(checkList, newPageList[:i]...)
				checkList = append(checkList, page)
				checkList = append(checkList, newPageList[i:]...)

				//fmt.Println(checkList)

				if pagesAreOkay(checkList) {
					newPageList = checkList
					break
				}
				//handle end of array-situation
				if i == len(newPageList)-1 {
					newPageList = append(newPageList, page)
					break
				}

			}
		}

		//fmt.Println(newPageList)

		middleNumberAsString := newPageList[len(newPageList)/2]
		middleNumber, _ := strconv.Atoi(middleNumberAsString)
		//fmt.Printf("[DEBUG] Adding Page Number %d from set %d\n", middleNumber, p)
		sum += middleNumber

	}
	fmt.Printf("Solution for part 2: %d\n", sum)
}
