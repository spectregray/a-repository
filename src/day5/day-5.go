package main

import (
	"bufio"
	"fmt"
	"strings"
	"strconv"
	utils "github.com/spectregray/aoc-2022/src/utils"
)

// git push -u -f origin main
func part1(scanner *bufio.Scanner) string {
	topCrates := ""
	crates := make([][]byte, 9)
	crateStrings := make([]string, 0)
	initialized := false

	for scanner.Scan() {
		line := scanner.Text()

		if !initialized {
			if line == "" {
				// now go through the lines backwards, initializing 2d array
				for i := len(crateStrings) - 2; i >= 0; i-- {
					slice := crateStrings[i]
					for j := 1; j < len(slice); j += 4 {
						if slice[j] != 32 {
							index := j / 4 // index of the crate
							crates[index] = append(crates[index], slice[j])
						}
					}
				}
				// debugging to make sure crates parsed right
				// for _, crate := range crates {
				// 	for _, item := range crate {
				// 		fmt.Printf("%c", item)
				// 	}
				// 	fmt.Println("")
				// }
				initialized = true
			} else {
				crateStrings = append(crateStrings, line)
			}
		} else if line[:4] == "move" {
			fields := strings.Fields(line)
			count, err := strconv.Atoi(fields[1])
			utils.Check(err)
			from, err := strconv.Atoi(fields[3])
			utils.Check(err)
			from -= 1
			to, err := strconv.Atoi(fields[5])
			utils.Check(err)
			to -= 1

			removed := crates[from][len(crates[from]) - count:]
			for i, j := 0, len(removed)-1; i < j; i, j = i+1, j-1 {
				removed[i], removed[j] = removed[j], removed[i]
			}
			crates[from] = crates[from][:len(crates[from]) - count]
			crates[to] = append(crates[to], removed...)
		}
	}

	// find top Crates 
	for _, crate := range crates {
		last := crate[len(crate) - 1]
		topCrates = topCrates + string(last)
	}
	return topCrates
}

// same as part 1, without reversing the removed slice
func part2(scanner *bufio.Scanner) string {
	topCrates := ""
	crates := make([][]byte, 9)
	crateStrings := make([]string, 0)
	initialized := false

	for scanner.Scan() {
		line := scanner.Text()

		if !initialized {
			if line == "" {
				// now go through the lines backwards, initializing 2d array
				for i := len(crateStrings) - 2; i >= 0; i-- {
					slice := crateStrings[i]
					for j := 1; j < len(slice); j += 4 {
						if slice[j] != 32 {
							index := j / 4 // index of the crate
							crates[index] = append(crates[index], slice[j])
						}
					}
				}
				// debugging to make sure crates parsed right
				// for _, crate := range crates {
				// 	for _, item := range crate {
				// 		fmt.Printf("%c", item)
				// 	}
				// 	fmt.Println("")
				// }
				initialized = true
			} else {
				crateStrings = append(crateStrings, line)
			}
		} else if line[:4] == "move" {
			fields := strings.Fields(line)
			count, err := strconv.Atoi(fields[1])
			utils.Check(err)
			from, err := strconv.Atoi(fields[3])
			utils.Check(err)
			from -= 1
			to, err := strconv.Atoi(fields[5])
			utils.Check(err)
			to -= 1

			removed := crates[from][len(crates[from]) - count:]
			crates[from] = crates[from][:len(crates[from]) - count]
			crates[to] = append(crates[to], removed...)
		}
	}

	// find top Crates 
	for _, crate := range crates {
		last := crate[len(crate) - 1]
		topCrates = topCrates + string(last)
	}
	return topCrates
}

func main() {
	inputFile := "input-day-5.txt"
	path := utils.GetPath(inputFile)

	// Part 1
	file, err := utils.OpenFile(path)
	utils.Check(err)
	scanner, err := utils.GetBufferedScanner(file)
	utils.Check(err)
	defer file.Close()

	fmt.Printf("Part 1: %s\n", part1(scanner))

	// Part 2
	file, err = utils.OpenFile(path)
	utils.Check(err)
	scanner, err = utils.GetBufferedScanner(file)
	utils.Check(err)

	fmt.Printf("Part 2: %s\n", part2(scanner))

	// Error handling
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}
}
