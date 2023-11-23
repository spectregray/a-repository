
package main 

import (
	"bufio"
	"fmt"
	"strconv"
	"container/heap" // I know the heap is of size 3 but just for the lols :)
	utils "github.com/spectregray/aoc/utils"
)

// trying different solution setups lol
func Day1(inputFile string) {
	// Setup
	if inputFile == "" {
		inputFile = "input-day-1.txt"
	}
	path := utils.GetPath(inputFile)
	file, err := utils.OpenFile(path)
	utils.Check(err)
	scanner, err := utils.GetBufferedScanner(file)
	utils.Check(err)
	defer file.Close() 

	// Solving
	fmt.Printf("Part 1: %d\n", Part1(scanner))
	scanner, err = utils.GetBufferedScanner(file)
	utils.Check(err)
	fmt.Printf("Part 2: %d\n", Part2(scanner))

	// Error handling
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}
}

func Part1(scanner *bufio.Scanner) (int) {
	mostCalories := 0
	currentCalories := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			if currentCalories > mostCalories {
				mostCalories = currentCalories
			}
			currentCalories = 0
		} else {
			num, err := strconv.Atoi(line)
			utils.Check(err)
			currentCalories += num
		}
	}

	return mostCalories
}

func Part2(scanner *bufio.Scanner) (int) {
	topThreeCalories := 0
	minHeap := &utils.MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, 3)
	heap.Push(minHeap, 4)
	heap.Push(minHeap, 1)
	// for scanner.Scan() {

	// }

	// sum values in heap

	return topThreeCalories
}

func main() {
	Day1("")
}
