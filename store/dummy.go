package store

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Dummy struct {
	client *mongo.Client
}

func NewDummy() Dummy {
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

	return Dummy{
		client: client,
	}
}

func (d Dummy) GetCollection(dbName string) *mongo.Collection {
	return d.client.Database(os.Getenv("TEST_DATABASE")).Collection(dbName)
}
