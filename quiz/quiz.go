package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

const QuestionsCSVPath = "./problems.csv"

type Score struct {
	Right int
	Wrong int
}

type QuestionAndAnswer struct {
	Question string
	Answer   int
}

func main() {
	// score := Score{}
	list := getQuestionAndAnswerList(QuestionsCSVPath)
	for _, questionAndAnswer := range list {
		fmt.Printf("What is %s?\n", questionAndAnswer.Question)
	}
}

func getQuestionAndAnswerList(filePath string) []QuestionAndAnswer {
	file, err := os.Open(filePath)
	questionsAndAnswersList := []QuestionAndAnswer{}

	if err != nil {
		fmt.Printf("Error reading file %v", err)
		return questionsAndAnswersList
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			fmt.Println("End of quiz!")
			return questionsAndAnswersList
		}
		questionString := row[0]
		answerString := row[1]
		answerInt, err := strconv.Atoi(answerString)

		nextQuestionAndAnswer := QuestionAndAnswer{Question: questionString, Answer: answerInt}
		questionsAndAnswersList = append(questionsAndAnswersList, nextQuestionAndAnswer)
	}
}
