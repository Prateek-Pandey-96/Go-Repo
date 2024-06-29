package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prateek69/fetcher-service/database"
	"github.com/prateek69/fetcher-service/models"
	"github.com/prateek69/fetcher-service/sitemap"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchSitemap(collection *mongo.Collection, cache *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		userReq := &models.UserReq{}
		if err := c.Bind(&userReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "improper request"})
		}

		key := fmt.Sprintf("%s_%d", userReq.Url, userReq.Depth)
		val, _ := database.GetCacheKey(cache, key)

		var pages []string
		if val == "not_found" {
			foundSitemap := &models.Sitemap{}
			filter := bson.M{"url": userReq.Url, "depth": userReq.Depth}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			err := collection.FindOne(ctx, filter).Decode(foundSitemap)
			if err == mongo.ErrNoDocuments {
				pages = sitemap.BFS(userReq.Url, userReq.Depth)
				// put in mongo
				if err := SaveToMongo(collection, userReq, pages); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err})
					return
				}
				// put in cache
				if err := SaveToCache(cache, key, pages); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err})
					return
				}
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}

			pages = foundSitemap.Destinations
			if err := SaveToCache(cache, key, pages); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
		} else {
			if err := json.Unmarshal([]byte(val), &pages); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
		}

		resp := CreateResponse(pages, userReq.Url)
		c.JSON(http.StatusAccepted, resp)
	}
}
