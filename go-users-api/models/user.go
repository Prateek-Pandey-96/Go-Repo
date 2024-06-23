package models

import "time"

type User struct {
	Id         int
	Username   string
	Email      string
	Created_at time.Time
}

type UserRequest struct {
	Username string `json:"username"  binding:"required"`
	Email    string `json:"email"  binding:"required"`
}
