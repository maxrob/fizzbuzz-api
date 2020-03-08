package main

import (
	"net/http"
	"os"

	"github.com/fizzbuzz-api/handlers"
	"github.com/fizzbuzz-api/store"
)

func main() {
	var storer store.Storer
	storer = store.NewMongoDB()

	http.HandleFunc("/fizzbuzzs", handlers.GetFizzbuzzHandler(storer))
	http.HandleFunc("/fizzbuzzs/most_popular_request", handlers.GetMostPopularRequestHandler(storer))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
