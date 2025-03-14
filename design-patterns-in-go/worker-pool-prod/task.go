package main

import "time"

type Task struct {
	Id int
}

func (t *Task) ProcessTask() {
	time.Sleep(1 * time.Microsecond)
}
