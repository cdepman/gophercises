package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type QuestionAndAnswer struct {
	Question string
	Answer   int
}

func main() {
	readCSV("./problems.csv")
}

func readCSV(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error reading file %v", err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			fmt.Println("End of quiz!")
			return
		}
		questionString := row[0]
		answerString := row[1]
		answerInt, err := strconv.Atoi(answerString)

		nextQuestionAndAnswer := QuestionAndAnswer{Question: questionString, Answer: answerInt}
		fmt.Printf("question: %s, answer: %d \n", nextQuestionAndAnswer.Question, nextQuestionAndAnswer.Answer)
	}

}
