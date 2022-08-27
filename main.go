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
	// This flag is like something we provide in command line
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the formate of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the csv file.")
	}
	problems := parseLines(lines)
	
	// Getting timer ready
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0

	// This lable is to break the code in select case statement (this is used instead of return)
	problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerChan := make(chan string)
		// This is just a way to code a function and call it right there, plus we dont need a name for it
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerChan:								
			if answer == p.a {
				correct ++
			}
		}
		
	}

	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) ([]problem) {
	ret:= make([]problem, len(lines))
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

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}