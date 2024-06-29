package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prateek69/fetcher-service/database"
	"github.com/prateek69/fetcher-service/middlewares"
	"github.com/prateek69/fetcher-service/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.Use(middlewares.AuthMiddleware)

	db := database.GetNewDB()
	collection := database.GetCollection(db)
	cache := database.GetNewCache()
	routers.AppRouter(r, collection, cache)

	r.Run(":" + os.Getenv("APP_PORT"))
}
