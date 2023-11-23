package utils

import (
	"bufio"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e) // exits with nonzero status
	}
}

func OpenFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	Check(err)
	return file, nil
}

func GetBufferedScanner(file *os.File) (*bufio.Scanner, error) {
	scanner := bufio.NewScanner(file)
	return scanner, nil
}

func GetPath(relativePath string) (fullPath string) {
	currentDir, err := os.Getwd()
	if err != nil {
		return
	}

	absolutePath := currentDir + "/" + relativePath
	return absolutePath
}