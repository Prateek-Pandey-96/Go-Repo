package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-auth/controllers"
	"github.com/prateek69/go-auth/initializers"
	"github.com/prateek69/go-auth/middlewares"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)
	r.Run("localhost:" + os.Getenv("PORT"))
}
