package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	ques string
	ans  string
}

func main() {
	fileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *fileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the provided file: %s\n", *fileName))
	}

	problems := parseLines(lines)

	correct := 0
	for i, prob := range problems {
		fmt.Printf("question number %d: %s\n", i+1, prob.ques)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == prob.ans {
			fmt.Println("Correct")
			correct++
		} else {
			fmt.Println("Incorrect")
		}
	}

	fmt.Printf("%d out of %d correct\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	res := make([]problem, len(lines))

	for i, line := range lines {
		res[i] = problem{
			ques: strings.TrimSpace(line[0]),
			ans:  strings.TrimSpace(line[1]),
		}
	}

	return res
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
