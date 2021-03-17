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
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the formt of 'question,answer'")
	timeLimit := flag.Int64("limit", 30, "a time limit to give an answer for the question in seconds")
	flag.Parse()

	// Open file
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	// read all the content of the file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	// Parse the lines from file to problem struct
	problems := parseLines(lines)

	//Set the timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)

		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)

			// Send answer to the chanel
			answerCh <- answer
		}()

		select {
		// when timeout
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
		// when answer come from the chanel
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

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

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
