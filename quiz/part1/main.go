package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Populate this struct and then use this to generate quiz.
// Flexible in case we have different source than csv
type problem struct {
	q string
	a string
}
func main()  {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
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
	correct := 0
	problems := parseLines(lines)
	for i, p := range problems {
		fmt.Printf("Problem: #%d: %s = \n", i+1, p.q)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == p.a {
			correct++
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