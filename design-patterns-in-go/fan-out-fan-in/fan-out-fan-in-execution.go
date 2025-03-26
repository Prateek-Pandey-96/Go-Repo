package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type WorkerPool struct {
	wg          *sync.WaitGroup
	concurrency int
	jobs        chan Person
	resultChan  chan Person
}

func (wp *WorkerPool) init(concurrency int) {
	wp.wg = &sync.WaitGroup{}
	wp.concurrency = concurrency
	wp.jobs = make(chan Person)
	wp.resultChan = make(chan Person)
}

func (wp *WorkerPool) mock1(personChan chan Person) chan Person {
	out := make(chan Person)
	go func() {
		defer close(out)
		for person := range personChan {
			out <- person
		}
	}()
	return out
}

func (wp *WorkerPool) mock2(personChan chan Person) chan Person {
	out := make(chan Person)
	go func() {
		defer close(out)
		for person := range personChan {
			out <- person
		}
	}()
	return out
}

func (wp *WorkerPool) work() {
	defer wp.wg.Done()
	for person := range wp.mock2(wp.mock1(wp.jobs)) {
		if person.Salary > 100000 {
			time.Sleep(1 * time.Microsecond)
			wp.resultChan <- person
		}
	}
}

func readerFunction(r *csv.Reader) chan Person {
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

func fanOutFanInExecution() {
	start := time.Now()
	file := Must(os.Open("data.csv"))
	defer file.Close()

	reader := csv.NewReader(file)
	workerPool := &WorkerPool{}
	workerPool.init(100)

	go func() {
		for person := range readerFunction(reader) {
			workerPool.jobs <- person
		}
		close(workerPool.jobs)
	}()

	for range workerPool.concurrency {
		workerPool.wg.Add(1)
		go workerPool.work()
	}

	cities := make(map[string]struct{})
	go func(cities map[string]struct{}) {
		for person := range workerPool.resultChan {
			cities[person.City] = struct{}{}
		}
		close(workerPool.resultChan)
	}(cities)

	workerPool.wg.Wait()
	fmt.Printf("City count for these people is %d\n", len(cities))
	fmt.Printf("Time taken in fan out fan in execution is %d ms\n", time.Since(start).Milliseconds())
}
