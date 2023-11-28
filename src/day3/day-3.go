package main

import (
	"bufio"
	"fmt"
	utils "github.com/spectregray/aoc-2022/src/utils"
)

func part1(scanner *bufio.Scanner) (int) {
	prioritySum := 0

	for scanner.Scan() {
		line := scanner.Text()
		halfLen := len(line) / 2
		firstHalfSet := make(map[byte]bool)

		// iterate through first half, add to set
		for i := 0; i < halfLen; i++ {
			firstHalfSet[line[i]] = true
		}

		// iterate through second half, check against set
		for i := halfLen; i < len(line); i++ {
			_, ok := firstHalfSet[line[i]]
			if ok {
				ascii := int(line[i])
				if ascii >= 97 {
					prioritySum += ascii - 96
				} else if ascii >= 65 {
					prioritySum += ascii - 64 + 26
				}
				break
			}
		}
	}

	return prioritySum
}

// assumes total no. of lines is a multiple of 3
func part2(scanner *bufio.Scanner) (int) {
	prioritySum := 0
	sharedMap := make(map[byte]bool)
	i := 0
	
	for scanner.Scan() {
		line := scanner.Text()

		// for the first elf, adds all of its bag contents to map
		if i == 0 {
			for j := 0; j < len(line); j++ {
				sharedMap[line[j]] = true
			}
		// for i > 0, checks against the map
		} else {
			tempMap := make(map[byte]bool)
			for j := 0; j < len(line); j++ {
				tempMap[line[j]] = true
			}

			for key := range sharedMap {
				_, ok := tempMap[key]
				if !ok {
					delete(sharedMap, key)
				} 
			}
		}

		i += 1
		if i == 3 {
			// there should only be one value
			if len(sharedMap) > 1 {
				panic("some reason, multiple values are left")
			}
			for key := range sharedMap {
				ascii := int(key)
				if ascii >= 97 {
					prioritySum += ascii - 96
				} else if ascii >= 65 {
					prioritySum += ascii - 64 + 26
				}
			}
			sharedMap = make(map[byte]bool)
			i = 0
		}
	} 

	return prioritySum
}

func main() {
	inputFile := "input-day-3.txt"
	path := utils.GetPath(inputFile)

	// Part 1
	file, err := utils.OpenFile(path)
	utils.Check(err)
	scanner, err := utils.GetBufferedScanner(file)
	utils.Check(err)
	defer file.Close()

	fmt.Printf("Part 1: %d\n", part1(scanner))

	// Part 2
	file, err = utils.OpenFile(path)
	utils.Check(err)
	scanner, err = utils.GetBufferedScanner(file)
	utils.Check(err)

	fmt.Printf("Part 2: %d\n", part2(scanner))

	// Error handling
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}
}