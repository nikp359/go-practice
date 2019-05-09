package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func startQuiz(quiz [][]string, timeLimit int) {
	reader := bufio.NewReader(os.Stdin)
	correctAnswer := 0
	lenQuiz := len(quiz)
	fmt.Printf("Quiz. Time limit: %v. To start the test press ENTER", timeLimit)
	reader.ReadString('\n')
	scanner := bufio.NewScanner(os.Stdin)

	for _, row := range quiz {
		isCorrect, err := question(row[0], row[1], timeLimit, scanner)
		if err != nil {
			break
		}

		if isCorrect {
			correctAnswer++
		}
	}

	fmt.Printf("Total questions: %v, Correct answers: %v\n", lenQuiz, correctAnswer)
}

func question(question string, answer string, timeLimit int, scanner *bufio.Scanner) (bool, error) {
	fmt.Printf("%s => ", question)
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	input := make(chan bool)

	go func() {
		if scanner.Scan() {
			userAnswer := scanner.Text()
			userAnswer = strings.ToLower(strings.Trim(userAnswer, "\n "))
			isCorrect := userAnswer == answer
			input <- isCorrect
		}
	}()

	select {
	case timeEnd := <-timer.C:
		fmt.Printf("\nTime is over %v\n", timeEnd)
		return false, errors.New("time is ower")
	case answer := <-input:
		return answer, nil
	}
}
