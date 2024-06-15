package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-mongo/controllers"
)

func UserRouter(r *gin.Engine) {
	r.GET("/user/:id", controllers.GetUserById())
	r.POST("/user", controllers.CreateUser())
	r.DELETE("/user/:id", controllers.DeleteUserById())
}
