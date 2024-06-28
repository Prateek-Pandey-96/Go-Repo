package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-jwt/controllers"
	"github.com/prateek69/go-jwt/middleware"
)

func UserRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())
	router.GET("users", controllers.GetAllUsers())
	router.GET("users/:user_id", controllers.GetUser())
}
