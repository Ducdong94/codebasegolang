package db

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"codebase.sample/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var lock = &sync.Mutex{}

var client *mongo.Client
var database *mongo.Database

func InitConnection() {
	getInstance()
}

func getInstance() *mongo.Database {
	if client == nil {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("Creating database connection now.")
		client, err := mongo.NewClient(options.Client().ApplyURI(configs.EnvMongoURI()))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		} else {
			database = client.Database("codebasegolang")
			fmt.Println("Connected to MongoDB")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return database
}

func GetCollection(name string) *mongo.Collection {
	return getInstance().Collection(name)
}
