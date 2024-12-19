package main

import (
	"AdventOfCode/util"
	"AdventOfCode/y2024"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	if args[0] == "solve" {
		input := util.GetInputStringFor(args[1], args[2])
		ans := y2024.Day3Part1(input)
		fmt.Println(ans)
	} else if args[0] == "download" {
		util.DownloadInput(args[1], args[2])
	}
}
