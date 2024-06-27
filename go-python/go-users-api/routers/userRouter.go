package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-users-api/controllers"
)

func UserRouter(r *gin.Engine, db *sql.DB) {
	r.GET("/user/:id", controllers.GetUser(db))
	r.POST("/user", controllers.CreateUser(db))
	r.DELETE("/user/:id", controllers.DeleteUser(db))
	r.PUT("/user/:id", controllers.UpdateUser(db))
}
