package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-jwt/routes"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = "8000"
	}

	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)
}
