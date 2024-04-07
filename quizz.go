package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readProblems() [][]string {
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Unable to read file", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, _ := csvReader.ReadAll()
	return records
}

func main() {
	records := readProblems()
	length := len(records)
	fmt.Println(records, length)

	var input string
	var score int = 0

	for i, v := range records {
		question := v[0]
        answer := v[1]
		fmt.Printf("Problem %v: %v\n", i, question)
		fmt.Scanln(&input)
		
		if input == answer {
			score += 1
		}

	}

	fmt.Printf("You finished the test! You scored %v/%v", score, length)
}