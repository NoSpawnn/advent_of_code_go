package runner

import (
	"AdventOfCode/y2024"
)

type Solver struct {
	Part1 func(string) int
	Part2 func(string) int
}

type AllSolverMap map[string]map[string]Solver

func GetAll() AllSolverMap {
	solvers := AllSolverMap{
		"2024": map[string]Solver{
			"01": {Part1: y2024.Day01Part1, Part2: y2024.Day01Part2},
			"02": {Part1: y2024.Day02Part1, Part2: y2024.Day02Part2},
			"03": {Part1: y2024.Day03Part1, Part2: y2024.Day03Part2},
		},
	}

	return solvers
}
