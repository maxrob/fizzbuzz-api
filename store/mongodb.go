package store

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

func NewMongoDB() MongoDB {
	clientOptions := options.Client().ApplyURI("mongodb://fizzbuzz-mongo:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return MongoDB{
		client: client,
	}
}

func (mgo MongoDB) GetCollection(dbName string) *mongo.Collection {
	return mgo.client.Database(os.Getenv("PRODUCTION_DATABASE")).Collection(dbName)
}
