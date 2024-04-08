package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"flag"
	"time"
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

func Timer(timeLimit int) {
	time.Sleep(time.Duration(timeLimit) * time.Second)
}

func Score(score int, length int) {
	fmt.Printf("You finished the test! You scored %v/%v", score, length)
}

func main() {
	var timeFlag = flag.Int("limit", 30, "Duration of the quiz")
	flag.Parse()
	timeLimit := *timeFlag
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	timerSignal := make(chan bool)

	go func() {
		<-timer.C
		timerSignal <- true
	}()

	records := readProblems()
	if records == nil {
		return
	}
	length := len(records)

	var input string
	var score int = 0
	for i, v := range records {
		question := v[0]
		answer := v[1]
		fmt.Printf("Problem %v: %v = ", i, question)

		answerCh := make(chan string)
		go func() {
			var input string
			fmt.Scanln(&input)
			answerCh <- input
		}()

		select {
		case <-timerSignal:
			fmt.Println("\nTime's up!")
			Score(score, length)
			return
		case input = <-answerCh:
			if input == answer {
				score++
			}
		}
	}

	Score(score, length)
}