package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 10, "the time limit for quiz is seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the csv file")
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	

	correct := 0
	for i, p := range problems {
		fmt.Printf("Q.%d What is %s \n", i+1, p.q)
		answerch := make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerch <- answer
		}()
		select {
		case <-timer.C: 
			fmt.Printf("You scored %d out of %d. \n", correct, len(problems))
			return
		case answer := <-answerch:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d \n", correct, len(lines))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
