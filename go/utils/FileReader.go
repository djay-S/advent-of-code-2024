package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

func GetInputFile(fileName string) []string {
	baseFolderPath := "../../input/"

	baseFolderPath += fileName

	//	fmt.Println(baseFolderPath)

	absPath, err := filepath.Abs(baseFolderPath)
	if err != nil {
		panic(err)
	}

	var lines []string

	file, err := os.Open(absPath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
