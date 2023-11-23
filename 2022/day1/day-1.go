
package main 

import (
	"fmt"
	utils "github.com/spectregray/aoc/utils"
)

func Day1(inputFile string) {
	// Setup
	if inputFile == "" {
		inputFile = "input-day-1.txt"
	}
	path := utils.GetPath(inputFile)
	scanner, file, err := utils.GetBufferedScanner(path)
	defer file.Close() 
	utils.Check(err)

	// Solution
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Error handling
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}
}

func main() {
	Day1()
}
