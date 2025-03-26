package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func readerFunc(r *csv.Reader) <-chan Person {
	out := make(chan Person)
	go func() {
		defer close(out)
		lineNum := 0
		for {
			row, err := r.Read()
			if err != nil {
				break // Stop reading on EOF or error
			}
			if lineNum == 0 {
				lineNum += 1
				continue
			}
			salary := Must(strconv.Atoi(row[2]))
			person := Person{
				Name:   row[0],
				City:   row[1],
				Salary: salary,
			}
			out <- person
		}
	}()
	return out
}

func processFunc(personChan <-chan Person) <-chan Person {
	out := make(chan Person)
	go func() {
		defer close(out)
		for person := range personChan {
			time.Sleep(1 * time.Microsecond)
			if person.Salary > 100000 {
				out <- person
			}
		}
	}()
	return out
}

func pipelinedExecution() {
	start := time.Now()
	file := Must(os.Open("data.csv"))
	defer file.Close()

	reader := csv.NewReader(file)
	cities := make(map[string]struct{})
	for person := range processFunc(readerFunc(reader)) {
		cities[person.City] = struct{}{}
	}
	fmt.Printf("City count for these people is %d\n", len(cities))
	fmt.Printf("Time taken in pipelined execution is %d ms\n", time.Since(start).Milliseconds())
}
