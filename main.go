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
	// Show help menu
	csvFilename := flag.String("csv", "problems.csv", "a csv file in a format of 'question,answer'")
	timeLimit := flag.Int("limit", 10, "the time limit for the quiz in seconds")
	flag.Parse()

	// Read questions from csv file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)

	// Set timer (default 10 seconds)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

problemLoop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.q)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == problem.a {
				correct++
			}
		}
	}

	fmt.Printf("\nYou scored %d out of %d \n", correct, len(problems))

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// Return i.e. [{q a}, {q a}, ...]
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}
