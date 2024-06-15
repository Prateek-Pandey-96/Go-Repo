package database

import (
	"context"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client
var mutex *sync.Mutex = &sync.Mutex{}

var mongoUri string = os.Getenv("CONNECTION_URI")

// Singelton
func GetConnection() *mongo.Client {
	if db == nil {
		mutex.Lock()
		defer mutex.Unlock()

		if db == nil {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
			if err != nil {
				panic(err)
			}
			db = client
		}
	}
	return db
}
