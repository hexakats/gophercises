package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	questions, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return questions
}

func askRandomQuestion(questions [][]string) {
	points := 0
	question_number := rand.IntN(len(questions) - 1)
	question := questions[question_number][0]
	answer := questions[question_number][1]

	loop := true

	for loop {
		fmt.Println("What's " + question + "?")
		var response string
		fmt.Scanln(&response)

		if string(response) != string(answer) {
			points--
			fmt.Printf("This answer is incorrect, -1 point, You now have %d points!\n", points)
			if points < 0 {
				println("GAME OVER")
				loop = false
			}
			question_number = rand.IntN(len(questions) - 1)
			question = questions[question_number][0]
			answer = questions[question_number][1]
		} else {
			points++
			fmt.Printf("This answer is correct! +1 point, You now have %d points! \n", points)
			question_number = rand.IntN(len(questions) - 1)
			question = questions[question_number][0]
			answer = questions[question_number][1]
		}
	}
	return
}

func main() {
	filePath := flag.String("csv", "problems.csv", "Reads a CSV file with questions in the first column and answers in the second column.")
	quiz := readCsvFile(*filePath)

	time := flag.Int("timer", 0, "Sets a timer of n seconds")
	flag.Parse()
	_ = time

	askRandomQuestion(quiz)

	println("Goodbye!")
}
