package main

import (
	"AdventOfCode/runner"
	"AdventOfCode/util"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	args := os.Args[1:]
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	yearStr := args[1]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		panic(err)
	}

	dayStr := args[2]
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		panic(err)
	}

	if args[0] == "solve" {
		solvers := runner.GetAll()
		input := util.GetInputStringFor(yearStr, dayStr)
		fn := solvers[yearStr][dayStr]
		ans := fn.Part1(input)
		fmt.Println(ans)
	} else if args[0] == "download" {
		util.DownloadInput(yearStr, dayStr)
	} else if args[0] == "create" {
		runner.GenerateSolutionFileFor(year, day)
		util.DownloadInput(yearStr, dayStr)
	}
}
