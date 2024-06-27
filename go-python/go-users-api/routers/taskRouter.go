package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-users-api/controllers"
	"github.com/redis/go-redis/v9"
)

func TaskRouter(r *gin.Engine, db *sql.DB, cache *redis.Client) {
	r.GET("/user/:id/tasks", controllers.GetTasks(db, cache))
	r.POST("/user/:id/task", controllers.CreateTask(db, cache))
	r.PUT("/user/:id/task/:taskid", controllers.FinishTask(db, cache))
}
