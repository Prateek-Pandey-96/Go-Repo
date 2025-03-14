package main

import "sync"

type WorkerPool struct {
	wg          *sync.WaitGroup
	concurrency int
	jobsChan    chan Task
	resultChan  chan Task
}

func (wp *WorkerPool) Init(concurrency int) {
	wp.concurrency = concurrency
	wp.wg = &sync.WaitGroup{}
	wp.jobsChan = make(chan Task)
	wp.resultChan = make(chan Task)
}

func (wp *WorkerPool) Work() {
	defer wp.wg.Done()
	for task := range wp.jobsChan {
		task.ProcessTask()
		wp.resultChan <- task
	}
}

func (wp *WorkerPool) Process(tasks []Task) {

	go func() {
		for _, task := range tasks {
			wp.jobsChan <- task
		}
		close(wp.jobsChan)
	}()

	go func() {
		for range wp.resultChan {

		}
		close(wp.resultChan)
	}()

	for range wp.concurrency {
		wp.wg.Add(1)
		go wp.Work()
	}
	wp.wg.Wait()
}
