package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-users-api/controllers"
)

func TaskRouter(r *gin.Engine, db *sql.DB) {
	r.GET("/user/:id/tasks", controllers.GetTasks(db))
	r.POST("/user/:id/task", controllers.CreateTask(db))
	r.PUT("/user/:id/task/:taskid", controllers.FinishTask(db))
}
