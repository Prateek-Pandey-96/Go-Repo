package database

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client
var mutex *sync.Mutex = &sync.Mutex{}

// singelton to get db connection
func GetNewDB() *mongo.Client {
	if db == nil {
		mutex.Lock()
		defer mutex.Unlock()

		if db == nil {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("CONNECTION_URI")))
			if err != nil {
				panic(err)
			}
			db = client
			log.Println("connection with mongo: successful")
		}
	}
	return db
}

func GetCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("Collection_Name"))
}
