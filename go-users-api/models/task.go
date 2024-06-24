package models

type TaskRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type Task struct {
	Name        string
	Description string
	TaskStatus  bool
}
