package main

import (
	"bufio"
	"fmt"
	"strings"
	utils "github.com/spectregray/aoc-2022/utils"
)

func part1(scanner *bufio.Scanner) (int) {
	pointsMap := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
		"A": 1,
		"B": 2,
		"C": 3,
	}
	// a rock, b paper, c scissors
	// x rock, y paper, z scissors
	victoryMap := map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}
	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, " ")

		totalScore += pointsMap[inputs[1]]
		// victory conditions score
		if pointsMap[inputs[0]] == pointsMap[inputs[1]] {
			totalScore += 3
		} else {
			if victoryMap[inputs[1]] == inputs[0] {
				totalScore += 6
			}
		}
	}

	return totalScore
}

func part2(scanner *bufio.Scanner) (int) {
	return 0
}

// I hate the system I used in day 1 I am just going to do everything in main :)
// update: nvm I think I still might use the part1 part2 structure
func main() {
	inputFile := "input-day-2.txt"
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