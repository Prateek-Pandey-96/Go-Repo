package controllers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/prateek69/fetcher-service/database"
	"github.com/prateek69/fetcher-service/models"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateResponse(pages []string, reqUrl string) models.Response {
	return models.Response{
		Url:          reqUrl,
		Destinations: pages,
	}
}

func SaveToCache(cache *redis.Client, key string, pages []string) error {
	data, err := json.Marshal(pages)
	if err != nil {
		return err
	}
	if err := database.SetCacheKey(cache, key, string(data)); err != nil {
		return err
	}
	return nil
}

func SaveToMongo(collection *mongo.Collection, userReq *models.UserReq, pages []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	newMap := &models.Sitemap{
		Username:     userReq.Username,
		Url:          userReq.Url,
		Depth:        userReq.Depth,
		Destinations: pages,
	}
	_, err := collection.InsertOne(ctx, newMap)
	if err != nil {
		return err
	}
	return nil
}
