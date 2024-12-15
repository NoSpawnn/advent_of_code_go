package main

import (
	"AdventOfCode/util"
	"AdventOfCode/y2024"
	"fmt"
)

func main() {
	input := util.ReadFileAsString("inputs/2024/day_3.txt")
	ans := y2024.Day3Part2(input)
	fmt.Println(ans)
}
