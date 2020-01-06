package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const QuestionsCSVPath = "./problems.csv"

type Score struct {
	Correct    int
	Incorrrect int
}

type QuestionAndAnswer struct {
	Question string
	Answer   int
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	score := Score{}
	list := getQuestionAndAnswerList(QuestionsCSVPath)
	for _, questionAndAnswer := range list {
		fmt.Printf("What is %s?\n> ", questionAndAnswer.Question)
		answer, _ := inputReader.ReadString('\n')
		answer = strings.TrimSuffix(answer, "\n")
		answerInt, err := strconv.Atoi(answer)
		if err != nil {
			fmt.Printf("That's not an integer :(, %s", err)
		}
		if answerInt == questionAndAnswer.Answer {
			score.Correct += 1
		} else {
			score.Incorrrect += 1
		}
	}
	fmt.Printf("Correct: %d\nIncorrrect: %d\n", score.Correct, score.Incorrrect)
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
