package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-users-api/models"
)

func GetUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		user := models.User{}

		query := `SELECT * from users where id = $1;`
		err := db.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.Email, &user.Created_at)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			} else {
				fmt.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to get user!"})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

func CreateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := models.UserRequest{}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "send proper request!"})
			return
		}

		query := `INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id;`
		var userId int
		if err := db.QueryRow(query, req.Username, req.Email).Scan(&userId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to create a new user!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user created with id: " + strconv.Itoa(userId)})
	}
}

func UpdateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		req := models.UserRequest{}
		if err := c.BindJSON(&req); err != nil {
			log.Printf("unable to map sent request to user request model %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "send proper request!"})
			return
		}

		query := `UPDATE users SET username = $1 where id = $2 RETURNING id;`
		var userId int
		if err := db.QueryRow(query, req.Username, id).Scan(&userId); err != nil {
			log.Printf("unable to update user in database %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to update user!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user updated with id: " + strconv.Itoa(userId)})
	}
}

func DeleteUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		query := `DELETE FROM users where id = $1 RETURNING id;`
		var deletedId int
		err := db.QueryRow(query, id).Scan(&deletedId)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			} else {
				log.Printf("unable to delete user: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to delete user!"})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user deleted with id: " + strconv.Itoa(deletedId)})
	}
}
