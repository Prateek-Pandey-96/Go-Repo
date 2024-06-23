package models

type TaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
