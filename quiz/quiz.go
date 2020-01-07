package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const QuestionsCSVPath = "./problems.csv"
const TimeAllowanceSeconds = 10

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
	timer := time.NewTimer(time.Second * TimeAllowanceSeconds)
	list := getQuestionAndAnswerList(QuestionsCSVPath)
	questionsChannel := make(chan QuestionAndAnswer, len(list))
	for _, questionAndAnswer := range list {
		questionsChannel <- questionAndAnswer
	}
	for {
		select {
		case gameOver := <-timer.C:
			fmt.Println("Game over!\n")
			fmt.Printf("You scored %d out of %d at %s", score.Correct, len(list), gameOver)
			return
		case nextQuestion := <-questionsChannel:
			fmt.Printf("What is %s?\n> ", nextQuestion.Question)
			answer, _ := inputReader.ReadString('\n')
			answer = strings.TrimSuffix(answer, "\n")
			answerInt, err := strconv.Atoi(answer)
			if err != nil {
				fmt.Printf("That's not an integer :(")
			}
			if answerInt == nextQuestion.Answer {
				score.Correct += 1
				fmt.Println("Correct!")
			} else {
				score.Incorrrect += 1
				fmt.Println("Incorrect!")
			}
		default:
			fmt.Printf("You scored %d out of %d", score.Correct, len(list))
			return
		}
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
