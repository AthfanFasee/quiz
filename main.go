package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// This flag is something we provide in command line
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the formate of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	shuffle := flag.String("shuffle", "no", "Shuffle the questions randomly")
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
	
	// Suffle the problems order if needed
	if *shuffle == "yes" {
		Shuffle(problems)
	}

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
			reader := bufio.NewReader(os.Stdin)
			answer, _ := reader.ReadString('\n')
			answerChan <- strings.TrimSpace(answer) // This will ignore case white spaces in answers
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerChan:							
			if strings.EqualFold(answer, p.a) {  // This will ignore case sensitive in answers
				correct ++
			}
		}	
	}
	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
}

// Parse the 2D slice into a slice with problem struct
func parseLines(lines [][]string) ([]problem) {
	ret:= make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), // trim whitespaces in csv file
		}
	}
	return ret
}

// Shuffle the parsedSlice
func Shuffle(slice []problem) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(slice); n > 0; n-- {
	   randIndex := r.Intn(n)
	   slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
	}
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}