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
		Users crud and unit tests
		Tasks crud and unit tests
		Adding cache
		Changing schema for task status
		Apply authentication
	*/
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	database.ConnectDatabase()
	db := database.GetNewDbConn()
	r := gin.Default()
	routers.UserRouter(r, db)
	r.Run(os.Getenv("APP_PORT"))
}
