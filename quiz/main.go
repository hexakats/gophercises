package main

import (
	"encoding/csv"
	"flag"
	"log"
	"math/rand"
	"os"
)

func readCsvFile(filePath String) [][]string {

	return quiz
}

func main() {
	file := flag.String("csv", "problems.csv", "Reads a csv file with questions in the first column and answers in the second column.")
	time := flag.Int("timer", 0, "Sets a timer of n seconds")
	flag.Parse()

	_ = file
	_ = time

	println("hello world")
}
