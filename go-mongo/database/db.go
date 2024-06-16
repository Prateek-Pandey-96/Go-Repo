package database

import (
	"context"
	"sync"
	"time"

	"github.com/prateek69/go-mongo/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client
var mutex *sync.Mutex = &sync.Mutex{}

// Singelton
func GetConnection() *mongo.Client {
	if db == nil {
		mutex.Lock()
		defer mutex.Unlock()

		if db == nil {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			mongoUri := config.AppConfig.ConnectionURI

			client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
			if err != nil {
				panic(err)
			}
			db = client
		}
	}
	return db
}

func GetCollection() *mongo.Collection {
	return GetConnection().Database(config.AppConfig.DatabaseName).Collection(config.AppConfig.CollectionName)
}
