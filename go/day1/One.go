package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"aoc2024/utils"
)

func main() {
	fileLines := utils.GetInputFile("day1/test.txt")

	answer := One(fileLines)
	fmt.Println(answer)
}

func One(fileLines []string) int {
	var leftLocationIdList, rightLocationIdList []int
	answer := 0

	for _, line := range fileLines {
		line = strings.ReplaceAll(line, "   ", ",")
		letters := strings.Split(line, ",")

		leftStr, _ := strconv.Atoi(letters[0])
		rightStr, _ := strconv.Atoi(letters[1])

		leftLocationIdList = append(leftLocationIdList, leftStr)
		rightLocationIdList = append(rightLocationIdList, rightStr)
	}
	sort.Ints(leftLocationIdList)
	sort.Ints(rightLocationIdList)

	for i := range leftLocationIdList {
		distance := Abs(leftLocationIdList[i] - rightLocationIdList[i])
		answer += distance
	}

	return answer
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func ReadFile() []string {
	absPath, err := filepath.Abs("../../input/day1/test.txt")
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
