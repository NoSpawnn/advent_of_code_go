package main

import (
	"AdventOfCode/util"
	"AdventOfCode/y2024"
	"fmt"
)

func main() {
	input := util.ReadFile("inputs/2024/day_1.txt")
	ans := y2024.Part2(input)
	fmt.Println(ans)
}
