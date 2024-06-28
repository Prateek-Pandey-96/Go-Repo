package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/prateek69/go-jwt/database"
	"github.com/prateek69/go-jwt/helpers"
	"github.com/prateek69/go-jwt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if validationErr := validate.Struct(user); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		}

		count, err := UserCollection.CountDocuments(context, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "error occured while doing email check"})
		}

		if count > 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "email already exists"})
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()

		if err := helpers.GenerateTokens(&user); err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "error occured while signing up"})
		}
		if err := helpers.HashPassword(&user); err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "error occured while signing up"})
		}
		insertionNum, err := UserCollection.InsertOne(context, user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "error occured while signing up"})
		}
		ctx.JSON(http.StatusOK, insertionNum)
	}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")

		if err := helpers.IsAdmin(ctx, userId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User
		err := UserCollection.FindOne(context, bson.M{"user_id": userId}).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := UserCollection.FindOne(context, bson.M{"email": user.Email}).Decode(&foundUser); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		if err := helpers.VerifyPassword(&user, &foundUser); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		// refreshes the token
		if err := helpers.GenerateTokens(&foundUser); err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "error occured while logging in"})
		}

		foundUser.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		if _, err := UserCollection.UpdateByID(context, user.User_id, foundUser); err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "error occured while logging in"})
		}
		ctx.JSON(http.StatusOK, foundUser)
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
