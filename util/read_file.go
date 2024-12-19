package util

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	FilePath = "./inputs/%s/day_%s.txt"
)

func GetInputStringFor(year, day string) string {
	file := fmt.Sprintf(FilePath, year, day)
	return ReadFileAsString(file)
}

func ReadFileAsString(path string) string {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(contents)
}

func ReadFileAsLines(path string) []string {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(contents), "\n")
}
