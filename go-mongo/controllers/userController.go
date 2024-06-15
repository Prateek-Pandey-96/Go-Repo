package controllers

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-mongo/database"
	"github.com/prateek69/go-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client = database.GetConnection()
var usersCollection *mongo.Collection = db.Database(os.Getenv("DATABASE")).Collection(os.Getenv("COLLECTION"))

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User

		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		insertionNum, err := usersCollection.InsertOne(ctx, newUser)
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

		err := usersCollection.FindOne(ctx, bson.M{"userid": intid}).Decode(foundUser)
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

		result, err := usersCollection.DeleteOne(ctx, bson.M{"userid": intid})
		deletedCount := strconv.Itoa(int(result.DeletedCount))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "No user found with id:" + id})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"user": deletedCount + " user deleted"})
	}
}
