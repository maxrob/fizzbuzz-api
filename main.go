package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/fizzbuzz-api/models"
	"github.com/fizzbuzz-api/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var storer store.Storer
	storer = store.NewMongoDB()

	http.HandleFunc("/fizzbuzzs", InsertStatisticHandler(storer))
	http.HandleFunc("/fizzbuzzs/most_popular_request", SelectAllStatisticsHandler(storer))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func InsertStatisticHandler(s store.Storer) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := s.GetCollection("requests")
		result, err := collection.InsertOne(context.TODO(), models.Request{
			FirstInt:     3,
			SecondInt:    4,
			Limit:        5,
			FirstString:  "test",
			SecondString: "test2",
			Iteration:    6,
		})
		if err != nil {
			log.Fatal(err)
		}

		js, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func SelectAllStatisticsHandler(s store.Storer) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		findOptions := options.Find()
		var results []*models.Request

		collection := s.GetCollection("requests")
		cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
		if err != nil {
			log.Fatal(err)
		}

		for cur.Next(context.TODO()) {
			var elem models.Request
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, &elem)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		cur.Close(context.TODO())

		js, err := json.Marshal(results)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
