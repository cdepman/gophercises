package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type QuestionAndAnswer struct {
	question string
	answer   int
}

func main() {

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

		if err != nil {
			fmt.Printf("Error reading CSV lines from file %v", err)
			return
		}
	}

}
