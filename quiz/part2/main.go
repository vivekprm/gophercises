package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Populate this struct and then use this to generate quiz.
// Flexible in case we have different source than csv
type problem struct {
	q string
	a string
}
// https://pkg.go.dev/time#Timer
// Timer: Fired once
// Ticket: Fired continuously after a duration
func main()  {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	expiry := flag.Int("timeout", 30, "timeout in second")
	
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("error in opening the csv file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse provided csv file.")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*expiry) * time.Second)
	
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem: #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func()  {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerCh <- ans
		}()
		select {
		// Waiting for message from the channel
		case <- timer.C:
			fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
			return
		case ans := <- answerCh:
			if ans == p.a {
				correct++
			}
		// default:
			// var ans string
			// Issue is we block here at Scanf even if we have ran out of time. And count this
			// answer even if we provided it after the timeout. So move it to different
			// go routine
			// fmt.Scanf("%s\n", &ans)
			
			// if ans == p.a {
			//	correct++
			//}
		}
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

func parseLines(lines[][]string) []problem{
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}