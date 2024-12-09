package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"maps"
	"slices"
)

const debug = false
const inputFile = "input.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInput(inputFile)
	initializePuzzle()
	part1()
	part2()
}

var diskMap []int // the disk as a whole

// stores the start and length for each block of file and free space
var fileStarts = make(map[int]int)
var fileBlocks = make(map[int]int)

var freeStarts = make(map[int]int)
var freeBlocks = make(map[int]int)

/* Do some puzzle initialization */

func initializePuzzle() {

	//first get size of disk
	diskSize := 0
	for _, block := range inputData[0] {
		blockSize := block - '0'
		diskSize += int(blockSize)
	}

	diskMap = make([]int, diskSize)

	fileId := 0
	freeSpaceIndex := 0
	diskSpaceIndex := 0

	for i, block := range inputData[0] {
		blockSize := block - '0'
		for chunk := 0; chunk < int(blockSize); chunk++ {
			if i%2 != 0 {

				diskMap[diskSpaceIndex] = -1

				if _, ok := freeStarts[fileId-1]; !ok {
					freeStarts[fileId-1] = diskSpaceIndex
				}

				diskSpaceIndex++
				freeSpaceIndex++
				freeBlocks[fileId-1]++

				continue
			}

			diskMap[diskSpaceIndex] = fileId

			if _, ok := fileStarts[fileId]; !ok {
				fileStarts[fileId] = diskSpaceIndex
			}
			fileBlocks[fileId]++

			diskSpaceIndex++
		}
		if i%2 == 0 {
			fileId++
		}
	}
	if debug {
		fmt.Printf("Evaluating disk of size %d\n", diskSize)
		printDebugInfo(diskMap)
		//fmt.Println(freeSpaces)
	}

}

/* Solve here */

func printDebugInfo(diskMap []int) {
	for _, chunk := range diskMap {
		if chunk != -1 {
			fmt.Printf("%d", chunk)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
	fmt.Printf("File-Starts: %+v\n", fileStarts)
	fmt.Printf("File-Blocks: %+v\n", fileBlocks)

	fmt.Printf("Free-Starts: %+v\n", freeStarts)
	fmt.Printf("Free-Blocks: %+v\n", freeBlocks)
}

func part1() {
	// copy original map of disk
	diskCopy := slices.Clone(diskMap)
	lastIndex := len(diskMap) - 1

	for i, fileId := range diskCopy {

		if fileId >= 0 {
			continue
		}

		for diskCopy[lastIndex] == -1 {
			lastIndex--
		}

		if i >= lastIndex {
			break
		}
		diskCopy[i], diskCopy[lastIndex] = diskCopy[lastIndex], diskCopy[i]

		if debug {
			printDebugInfo(diskCopy)
		}

	}
	checksum := 0
	for i, fileId := range diskCopy[:lastIndex+1] {
		checksum += fileId * i
	}

	fmt.Printf("Solution for part 1: %d\n", checksum)
}

func part2() {

	diskCopy := slices.Clone(diskMap)
	files := slices.Collect(maps.Keys(fileBlocks))
	slices.Sort(files)
	slices.Reverse(files)

	freeSpaces := slices.Collect(maps.Keys(freeBlocks))
	slices.Sort(freeSpaces)

	for _, file := range files {
		if file == 0 {
			break
		}

		fileLength := fileBlocks[file]
		for _, free := range freeSpaces {

			//only run this file if the file is (potentially) moved to the left
			if free >= file {
				break
			}

			if freeBlocks[free] >= fileLength {

				start := freeStarts[free]

				//move on disk
				for i := 0; i < fileLength; i++ {
					diskCopy[start+i] = file
					diskCopy[fileStarts[file]+i] = -1
				}

				//make free block smaller
				freeStarts[free] += fileLength
				freeBlocks[free] -= fileLength

				//move file start
				fileStarts[file] = start

				break
			}

		}
		if debug {
			printDebugInfo(diskCopy)
		}

	}

	checksum := 0
	for i, fileId := range diskCopy {
		if fileId > -1 {
			checksum += fileId * i
		}

	}
	fmt.Printf("Solution for part 2: %d\n", checksum)
}
