// https://adventofcode.com/2024/day/3

package y2024

import (
	"regexp"
	"strconv"
)

func Day03Part1(input string) int {
	nums := parse(input, `(mul\((\d{1,3}),(\d{1,3})\))`)

	result := 0
	for _, nums := range nums {
		result += nums[0] * nums[1]
	}

	return result
}

func Day03Part2(input string) int {
	nums := parse(input, `(mul\((\d{1,3}),(\d{1,3})\))|(do\(\)|don't\(\))`)

	result := 0
	for _, nums := range nums {
		result += nums[0] * nums[1]
	}

	return result
}

func parse(input string, rx string) [][]int {
	rgx := regexp.MustCompile(rx)
	ops := rgx.FindAllStringSubmatch(input, -1)

	var nums [][]int
	enabled := true
	for _, op := range ops {
		switch op[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				n1, _ := strconv.Atoi(op[2])
				n2, _ := strconv.Atoi(op[3])
				nums = append(nums, []int{n1, n2})
			}
		}
	}

	return nums
}
