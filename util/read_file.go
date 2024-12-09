package util

import (
	"log"
	"os"
)

func ReadFileAsString(path string) string {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(contents)
}
