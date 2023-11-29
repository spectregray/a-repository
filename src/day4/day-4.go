package main

import (
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"sort"
	utils "github.com/spectregray/aoc-2022/src/utils"
)

func part1(scanner *bufio.Scanner) int {
	overlapCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		elves := strings.Split(line, ",")
		var intervals [][]int

		for _, element := range elves {
			interval := strings.Split(element, "-")
			start, err := strconv.Atoi(interval[0])
			utils.Check(err)
			end, err := strconv.Atoi(interval[1])
			utils.Check(err)
			intervals = append(intervals, []int{start, end})
		}

		sort.Slice(intervals, func(i, j int) bool {
			// alternatively, can check for equal starts in the later if statement
			if intervals[i][0] == intervals[j][0] {
				return intervals[i][1] > intervals[j][1]
			}
			return intervals[i][0] < intervals[j][0]
		})
		
		if intervals[0][1] >= intervals[1][1] {
			overlapCount += 1
		}
	}

	return overlapCount
}

func part2(scanner *bufio.Scanner) int {
	return 0
}

func main() {
	inputFile := "input-day-4.txt"
	// inputFile := "ex.txt"
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
