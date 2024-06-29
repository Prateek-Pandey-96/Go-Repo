package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var c *redis.Client

// singelton to get cache connection
func GetNewCache() *redis.Client {
	if c == nil {
		mutex.Lock()
		defer mutex.Unlock()

		if c == nil {
			opt, err := redis.ParseURL(os.Getenv("REDIS_CONNECTION_STRING"))
			if err != nil {
				panic(err)
			}
			log.Println("Successfully connected to cache!")
			cache := redis.NewClient(opt)
			c = cache
		}
	}
	return c
}

func SetCacheKey(client *redis.Client, key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := client.Set(ctx, key, value, 15*time.Minute).Err()
	if err != nil {
		log.Printf("unable to set value for key: %s", key)
		log.Printf("error: %v", err)
	}

	return nil
}

func GetCacheKey(client *redis.Client, key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	value, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "not_found", nil
	}
	if err != nil {
		log.Printf("unable to get value for key: %s", key)
		log.Printf("error: %v", err)
	}

	return value, nil
}
