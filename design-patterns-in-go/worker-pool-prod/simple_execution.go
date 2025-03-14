package main

func SimpleExecution(tasks []Task) {
	for _, task := range tasks {
		task.ProcessTask()
	}
}
