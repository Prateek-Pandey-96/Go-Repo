package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prateek69/go-mongo/router.go"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	r := gin.New()
	router.UserRouter(r)
	r.Run(":" + os.Getenv("PORT"))
}
