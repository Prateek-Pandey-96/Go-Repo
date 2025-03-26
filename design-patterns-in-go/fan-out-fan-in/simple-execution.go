package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func simpleExecution() {
	start := time.Now()

	file := Must(os.Open("data.csv"))
	defer file.Close()

	reader := csv.NewReader(file)

	records := Must(reader.ReadAll())
	people := make([]Person, 0, 50000)
	for i, row := range records {
		if i == 0 { // Skip header
			continue
		}

		salary := Must(strconv.Atoi(row[2]))
		person := Person{
			Name:   row[0],
			City:   row[1],
			Salary: salary,
		}
		people = append(people, person)
	}

	filteredPeople := make([]Person, 0, 25000)
	for _, person := range people {
		if person.Salary > 100000 {
			filteredPeople = append(filteredPeople, person)
		}
	}

	cities := make(map[string]struct{})
	for _, person := range filteredPeople {
		cities[person.City] = struct{}{}
	}

	fmt.Printf("People with salary more than 100k are %d\n", len(filteredPeople))
	fmt.Printf("City count for these people is %d\n", len(cities))
	fmt.Printf("Time taken is %d ms\n", time.Since(start).Milliseconds())
}
