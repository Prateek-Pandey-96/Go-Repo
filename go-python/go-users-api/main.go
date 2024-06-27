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
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// initialize connections to db and cache
	database.ConnectDatabase()
	db := database.GetNewDbConn()

	database.ConnectCache()
	client := database.GetNewCacheConn()

	r := gin.Default()
	routers.UserRouter(r, db)
	routers.TaskRouter(r, db, client)

	r.Run(os.Getenv("APP_PORT"))
}
