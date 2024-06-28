package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-jwt/controllers"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("users/signup", controllers.SignUp())
	router.POST("users/login", controllers.Login())
}
