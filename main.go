package main

import (
	"AdventOfCode/util"
	"AdventOfCode/y2024"
	"fmt"
)

func main() {
	input := util.ReadFileAsString("inputs/2024/day_2.txt")
	ans := y2024.Day2Part1(input)
	fmt.Println(ans)
}
