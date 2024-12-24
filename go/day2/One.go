package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2024/utils"
)

func main() {
	fileLines := utils.GetInputFile("day2/puzzle.txt")
	answer := One(fileLines)
	fmt.Println(answer)
}

func One(fileLines []string) int {
	var answer int

	for _, line := range fileLines {
		letters := strings.Split(line, " ")

		var nums []int
		for _, numStr := range letters {
			num, _ := strconv.Atoi(numStr)

			nums = append(nums, num)
		}
		isLineSafe := isLineSafe(nums)
		if isLineSafe {
			answer++
		}
	}

	return answer
}

func isLineSafe(nums []int) bool {
	diff := nums[0] - nums[1]

	if diff == 0 {
		return false
	}
	if diff < -3 || diff > 3 {
		return false
	}

	for i := 1; i < len(nums)-1; i++ {
		newDiff := nums[i] - nums[i+1]
		if newDiff == 0 {
			return false
		}
		if newDiff > 3 || newDiff < -3 {
			return false
		}
		if newDiff*diff < 0 {
			return false
		}
	}
	return true
}
