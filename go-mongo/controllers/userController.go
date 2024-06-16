package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-mongo/database"
	"github.com/prateek69/go-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User

		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		insertionNum, err := database.GetCollection().InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusAccepted, insertionNum)
	}
}

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		intid, _ := strconv.Atoi(id)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		foundUser := &models.User{}

		err := database.GetCollection().FindOne(ctx, bson.M{"userid": intid}).Decode(foundUser)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "No user found with id:" + id})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"user": foundUser})
	}
}

func DeleteUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		intid, _ := strconv.Atoi(id)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		result, err := database.GetCollection().DeleteOne(ctx, bson.M{"userid": intid})
		deletedCount := strconv.Itoa(int(result.DeletedCount))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "No user found with id:" + id})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"user": deletedCount + " user deleted"})
	}
}

func UpdateUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		intid, _ := strconv.Atoi(id)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		newUser := &models.User{}
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		foundUser := &models.User{}
		if err := database.GetCollection().FindOne(ctx, bson.M{"userid": intid}).Decode(foundUser); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "No user found with id:" + id})
			return
		}

		foundUser.Age = newUser.Age
		foundUser.Name = newUser.Name
		foundUser.Gender = newUser.Gender

		if _, err := database.GetCollection().UpdateOne(ctx, bson.M{"userid": intid}, bson.M{"$set": foundUser}); err != nil {
			fmt.Print(err)
			c.JSON(http.StatusNotFound, gin.H{"message": "Error while updating"})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"userUpdated": foundUser})
	}
}
