package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 15, "the time limit for the question in seconds")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintln("Failed to open CSV file:", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse the provided CSV file")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)

		answerCh := make(chan string)

		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)

			answerCh <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d \n", correct, len(problems))
			return
		case ans := <- answerCh:
			if ans == p.answer {
				correct ++
			}
		}

	}
	fmt.Printf("You scored %d out of %d \n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
