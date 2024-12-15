// https://adventofcode.com/2024/day/2

package y2024

import (
	"strconv"
	"strings"
)

type reportsState int

const (
	increasing reportsState = iota
	decreasing
)

func Day2Part1(input string) int {
	safeCount := 0

	for _, line := range strings.Split(input, "\n") {
		if isSafe(lineToNums(line)) {
			safeCount++
		}
	}

	return safeCount
}

func Day2Part2(input string) int {
	safeCount := 0

	for _, line := range strings.Split(input, "\n") {
		if existsSafePermutation(lineToNums(line)) {
			safeCount++
		}
	}

	return safeCount
}

func isSafe(nums []int) bool {
	state := increasing
	if nums[0] > nums[1] {
		state = decreasing
	}

	prev := nums[0]
	for _, num := range nums[1:] {
		if num == prev ||
			(state == increasing && (num < prev || num > prev+3)) ||
			(state == decreasing && (num > prev || num < prev-3)) {
			return false
		}

		prev = num
	}

	return true
}

func existsSafePermutation(nums []int) bool {
	if isSafe(nums) {
		return true
	}

	// Not allocating an entire new array results in the original being modified,
	// weird Go behaviour, I guess?
	for i := range nums {
		newNums := make([]int, len(nums))
		copy(newNums, nums)

		if isSafe(append(newNums[:i], newNums[i+1:]...)) {
			return true
		}
	}

	return false
}

func lineToNums(line string) []int {
	numsStrs := strings.Split(strings.TrimSpace(line), " ")

	nums := make([]int, len(numsStrs))
	for i, numStr := range numsStrs {
		nums[i], _ = strconv.Atoi(numStr)
	}

	return nums
}
