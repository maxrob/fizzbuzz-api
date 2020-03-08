package repositories

import (
	"context"
	"errors"

	"github.com/fizzbuzz-api/models"
	"github.com/fizzbuzz-api/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateRequest(s store.Storer, request *models.Request) error {
	collection := s.GetCollection("requests")
	opts := options.FindOne()
	filter := bson.D{
		{"first_int", request.FirstInt},
		{"second_int", request.SecondInt},
		{"limit", request.Limit},
		{"first_string", request.FirstString},
		{"second_string", request.SecondString},
	}

	result := &models.Request{}
	err := collection.FindOne(context.TODO(), filter, opts).Decode(result)
	if err != nil && err != mongo.ErrNoDocuments {
		return errors.New("database-find-error")
	}

	if err != mongo.ErrNoDocuments {
		opts := options.Update().SetUpsert(false)
		filter := bson.D{{"_id", result.ID}}
		update := bson.D{{"$set", bson.D{{"iteration", result.Iteration + 1}}}}

		_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
		if err != nil {
			return errors.New("database-update-error")
		}

		return nil
	}

	request.Iteration = 1
	_, err = collection.InsertOne(context.TODO(), request)
	if err != nil {
		return errors.New("database-insert-error")
	}

	return nil
}

func GetMostPopularRequest(s store.Storer) (*models.Request, error) {
	collection := s.GetCollection("requests")
	opts := options.FindOne().SetSort(bson.D{{"iteration", -1}})

	result := &models.Request{}
	err := collection.FindOne(context.TODO(), bson.D{}, opts).Decode(&result)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, errors.New("database-find-error")
	}

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	return result, nil
}
