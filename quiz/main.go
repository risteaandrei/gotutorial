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
	filename := flag.String("csv", "problems.csv", "csv file containing questions/answers")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println("Failed to open file", filename, ":", err.Error())
		os.Exit(1)
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to read csv file", filename)
	}

	problems := parseLines(records)
	correct := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem %d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var a string
			fmt.Scanf("%s", &a)
			answerCh <- a
		}()
		select {
		case <-timer.C:
			fmt.Println("\nYou scored", correct, "out of", len(problems))
			return
		case a := <-answerCh:
			if a == p.a {
				correct++
			}
		}

	}

	fmt.Println("You scored", correct, "out of", len(problems))
}

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
