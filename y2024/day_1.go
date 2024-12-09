package y2024

import (
	"math"
	"regexp"
	"slices"
	"strconv"
)

func Day1Part1(input string) int {
	leftNums, rightNums := getColumns(input)

	slices.Sort(leftNums)
	slices.Sort(rightNums)

	totalDistance := float64(0)
	for i := 0; i < len(leftNums); i++ {
		totalDistance += math.Abs(float64(leftNums[i] - rightNums[i]))
	}

	return int(totalDistance)
}

func Day1Part2(input string) int {
	similarityScore := 0
	leftNums, rightNums := getColumns(input)

	occurrences := map[int]int{}
	for _, num := range rightNums {
		occurrences[num]++
	}

	for _, num := range leftNums {
		similarityScore += num * occurrences[num] // else 0
	}

	return similarityScore
}

func getColumns(input string) ([]int, []int) {
	rgx := regexp.MustCompile(`(\d+)   (\d+)`)
	matches := rgx.FindAllStringSubmatch(input, -1)

	// Split into columns
	leftNums, rightNums := make([]int, len(matches)), make([]int, len(matches))
	for i, match := range matches {
		l, r := match[1], match[2]
		leftNums[i], _ = strconv.Atoi(l)
		rightNums[i], _ = strconv.Atoi(r)
	}

	return leftNums, rightNums
}
