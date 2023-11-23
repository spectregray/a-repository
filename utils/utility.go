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

func openFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	Check(err)
	return file, nil
}

func GetBufferedScanner(filePath string) (*bufio.Scanner, *os.File, error) {
	file, err := openFile(filePath)
	Check(err)

	scanner := bufio.NewScanner(file)
	return scanner, file, nil
}

func GetPath(relativePath string) (fullPath string) {
	currentDir, err := os.Getwd()
	if err != nil {
		return
	}

	absolutePath := currentDir + "/" + relativePath
	return absolutePath
}