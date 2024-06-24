package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-users-api/models"
)

func GetTasks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userid := c.Param("id")
		query := `SELECT name, description, task_status from tasks where user_id = $1;`

		rows, err := db.Query(query, userid)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "No tasks found"})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to fetch tasks for the user!"})
			return
		}

		var tasks []models.Task
		for rows.Next() {
			task := models.Task{}
			if err := rows.Scan(&task.Name, &task.Description, &task.TaskStatus); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to fetch tasks for the user!"})
				return
			}
			tasks = append(tasks, task)
		}

		c.JSON(http.StatusOK, gin.H{"tasks": tasks})
	}
}

func CreateTask(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userid := c.Param("id")
		taskreq := models.TaskRequest{}
		if err := c.BindJSON(&taskreq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "please send proper task!"})
			return
		}

		query := `INSERT INTO tasks (user_id, name, description) VALUES ($1, $2, $3) RETURNING task_id;`
		var taskid int
		if err := db.QueryRow(query, userid, taskreq.Name, taskreq.Description).Scan(&taskid); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to create a new task!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "task created with id: " + strconv.Itoa(taskid)})
	}
}

func FinishTask(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskid := c.Param("taskid")
		userid := c.Param("id")

		query := `UPDATE tasks SET task_status = TRUE WHERE task_id = $1 AND user_id = $2 RETURNING task_id;`
		var deletedid int
		err := db.QueryRow(query, taskid, userid).Scan(&deletedid)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "this user doesnt own the task or task doesn't exist!"})
			return
		}
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "unable to update task status!"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "task finished with id: " + strconv.Itoa(deletedid)})
	}
}
