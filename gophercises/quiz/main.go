package main

import (
	"encoding/csv"
	"flag"
	"os"
)

func main() {
	csvPath := flag.String("f", "problems.csv", "test file name")
	timeLimit := flag.Int("t", 30, "test time limit")
	flag.Parse()

	csvFile, err := os.Open(*csvPath)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	startQuiz(records, *timeLimit)
}
