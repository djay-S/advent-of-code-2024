package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2024/utils"
)

func main() {
	fileLines := utils.GetInputFile("day1/test.txt")

	answer := Two(fileLines)
	fmt.Println(answer)
}

func Two(fileLines []string) int {
	var leftLocationIdList, rightLocationIdList []int
	answer := 0

	for _, line := range fileLines {
		line := strings.ReplaceAll(line, "   ", ",")
		letters := strings.Split(line, ",")

		leftId, _ := strconv.Atoi(letters[0])
		rightId, _ := strconv.Atoi(letters[1])

		leftLocationIdList = append(leftLocationIdList, leftId)
		rightLocationIdList = append(rightLocationIdList, rightId)
	}
	locationMap := getRightLocationIdMap(rightLocationIdList)

	for _, locationId := range leftLocationIdList {
		similarityScore := locationId * locationMap[locationId]

		answer += similarityScore
	}

	return answer
}

func getRightLocationIdMap(locationIdList []int) map[int]int {
	locationMap := make(map[int]int)
	for _, locationId := range locationIdList {
		freq := locationMap[locationId]
		freq++
		locationMap[locationId] = freq
	}

	return locationMap
}
