package store

import "go.mongodb.org/mongo-driver/mongo"

type Storer interface {
	GetCollection(string) *mongo.Collection
}
