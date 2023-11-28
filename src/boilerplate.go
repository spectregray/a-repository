package main

import (
	"bufio"
	"fmt"
	utils "github.com/spectregray/aoc-2022/src/utils"
)

func part1(scanner *bufio.Scanner) int {
	return 0
}

func part2(scanner *bufio.Scanner) int {
	return 0
}

func main() {
	inputFile := ""
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
