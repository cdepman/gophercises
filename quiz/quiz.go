package main

import "os"

import "fmt"

import "encoding/csv"

func main() {

}

func readCSV(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error reading file %v", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file).ReadAll()

}
