package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-users-api/controllers"
)

func TaskRouter(r *gin.Engine) {
	r.GET("/user/:id/tasks", controllers.GetTasks)
	r.POST("/user/:id/task", controllers.CreateTask)
	r.DELETE("/user/:id/task/:taskid", controllers.DeleteTask)
	r.PUT("/user/:id/task/:taskid", controllers.UpdateTask)
}
