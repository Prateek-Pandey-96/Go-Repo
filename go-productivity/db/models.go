package db

type Task struct {
	Id    uint64
	Title string
}

type FinishedTask struct {
	Timestamp uint64
	Title     string
}
