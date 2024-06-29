package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/prateek69/fetcher-service/controllers"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func AppRouter(r *gin.Engine, collection *mongo.Collection, cache *redis.Client) {
	r.POST("/fetcher/fetch", controllers.FetchSitemap(collection, cache))
}
