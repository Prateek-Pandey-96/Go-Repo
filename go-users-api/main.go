package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prateek69/go-users-api/database"
	"github.com/prateek69/go-users-api/routers"
)

func main() {
	/*
		Adding cache
		Tasks unit tests
		Changing schema for task status
		Apply authentication
		Automate cleaning finished tasks
	*/
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	database.ConnectDatabase()
	db := database.GetNewDbConn()
	r := gin.Default()
	routers.UserRouter(r, db)
	routers.TaskRouter(r, db)
	r.Run(os.Getenv("APP_PORT"))
}
