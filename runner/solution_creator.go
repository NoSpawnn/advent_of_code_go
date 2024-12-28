package runner

import (
	"errors"
	"fmt"
	"html/template"
	"os"
)

const (
	TEMPLATE_FILE = "solution_template.tmpl"
	SOLUTION_DIR  = "y%d/"
	SOLUTION_FILE = "day_%02d.go"
	SOLUTION_FP   = SOLUTION_DIR + SOLUTION_FILE
)

type solutionDate struct {
	Year int
	Day  int
}

func GenerateSolutionFileFor(year, day int) {
	tmpl, err := template.New(TEMPLATE_FILE).ParseFiles(TEMPLATE_FILE)
	if err != nil {
		panic(err)
	}

	dir := fmt.Sprintf(SOLUTION_DIR, year)
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		os.Mkdir(dir, os.ModePerm)
	}

	solFileName := fmt.Sprintf(SOLUTION_FP, year, day)
	solFile, err := os.Create(solFileName)
	if err != nil {
	}

	err = tmpl.Execute(solFile, solutionDate{year, day})
	if err != nil {
		panic(err)
	}
}
