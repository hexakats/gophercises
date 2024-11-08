package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func getQuestionsFromCsv(filePath string, shuffle bool) [][]string {
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

	if shuffle {
		for i := range questions {
			j := rand.Intn(i + 1)
			questions[i], questions[j] = questions[j], questions[i]
		}
	}
	return questions
}

func askQuestion(questions [][]string) {
	points := 0
	question_number := 0

	for true {
		fmt.Println("What's " + questions[question_number][0] + "?")
		var response string
		fmt.Scanln(&response)

		if string(response) != string(questions[question_number][1]) {
			fmt.Println("This answer is incorrect.")
		} else {
			points++
			fmt.Printf("This answer is correct! +1 point, You now have %d points! \n", points)
		}
		question_number++
		if question_number == len(questions) {
			break
		}
	}
	fmt.Printf("You scored %1d out of %2d in the quiz.\n", points, len(questions))
	return
}

func gameTimer(duration int) {
	timer := time.NewTimer(time.Duration(duration) * time.Second)
	<-timer.C
	fmt.Println("Timeout, You fail.")
	os.Exit(1)
}

func main() {
	filePath := flag.String("csv", "problems.csv", "Reads a CSV file with questions in the first column and answers in the second column.")
	limit := flag.Int("limit", 30, "Sets a timer of n seconds. default is -limit=30")
	shuffle := flag.Bool("shuffle", false, "Shuffles the question order when used.")

	flag.Parse()
	questions := getQuestionsFromCsv(*filePath, *shuffle)

	go gameTimer(*limit)
	askQuestion(questions)
}
