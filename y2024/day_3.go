// https://adventofcode.com/2024/day/3

package y2024

import (
	"regexp"
	"strconv"
	"strings"
)

func Day3Part1(input string) int {
	nums := getMulNums(input)

	result := 0
	for _, nums := range nums {
		result += nums[0] * nums[1]
	}

	return result
}

func Day3Part2(input string) int {
	nums := getMulNumsHonouringToggle(input)

	result := 0
	for _, nums := range nums {
		result += nums[0] * nums[1]
	}

	return result
}

func getMulNums(input string) [][]int {
	rgx := regexp.MustCompile(`mul\((\d{1,4}),(\d{1,4})\)`)
	muls := rgx.FindAllStringSubmatch(input, -1)

	nums := make([][]int, len(muls))
	for i, match := range muls {
		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])
		nums[i] = []int{n1, n2}
	}

	return nums
}

func getMulNumsHonouringToggle(input string) [][]int {
	rgx := regexp.MustCompile(`(mul\((\d{1,4}),(\d{1,4})\))|(do\(\)|don't\(\))`)
	ops := rgx.FindAllStringSubmatch(input, -1)

	var nums [][]int
	enabled := true
	for _, op := range ops {
		if op[0] == "do()" {
			enabled = true
		} else if op[0] == "don't()" {
			enabled = false
		}

		if enabled && strings.HasPrefix(op[0], "mul") {
			n1, _ := strconv.Atoi(op[2])
			n2, _ := strconv.Atoi(op[3])
			nums = append(nums, []int{n1, n2})
		}
	}

	return nums
}
