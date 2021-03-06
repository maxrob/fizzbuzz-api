package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fizzbuzz-api/helpers"
	"github.com/fizzbuzz-api/repositories"
	"github.com/fizzbuzz-api/store"
)

func GetMostPopularRequestHandler(s store.Storer) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, helpers.ErrorMethodNotAllowed, http.StatusMethodNotAllowed)
			return
		}

		result, err := repositories.GetMostPopularRequest(s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json_response, err := json.Marshal(result)
		if err != nil {
			http.Error(w, helpers.ErrorJSONEncode, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json_response)
	}
}
