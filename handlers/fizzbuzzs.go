package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/fizzbuzz-api/helpers"
	"github.com/fizzbuzz-api/models"
	"github.com/fizzbuzz-api/repositories"
	"github.com/fizzbuzz-api/store"
)

func GetFizzbuzzHandler(s store.Storer) func(http.ResponseWriter, *http.Request) {
	getParams := func(queryParams url.Values) (*models.Request, error) {
		firstInt, err := helpers.GetIntParams(queryParams, "int1")
		if err != nil {
			return nil, err
		}

		secondInt, err := helpers.GetIntParams(queryParams, "int2")
		if err != nil {
			return nil, err
		}

		limit, err := helpers.GetIntParams(queryParams, "limit")
		if err != nil {
			return nil, err
		}

		firstString, err := helpers.GetStringParams(queryParams, "str1")
		if err != nil {
			return nil, err
		}

		secondString, err := helpers.GetStringParams(queryParams, "str2")
		if err != nil {
			return nil, err
		}

		return &models.Request{
			FirstInt:     firstInt,
			SecondInt:    secondInt,
			Limit:        limit,
			FirstString:  firstString,
			SecondString: secondString,
		}, nil
	}
	getFizzbuzzList := func(request *models.Request) []string {
		var results []string

		for number := 1; number <= request.Limit; number++ {
			result := helpers.Fizzbuzzer(number, request.FirstInt, request.FirstString)
			result = result + helpers.Fizzbuzzer(number, request.SecondInt, request.SecondString)
			if len(result) == 0 {
				result = strconv.Itoa(number)
			}
			results = append(results, result)
		}

		return results
	}
	return func(w http.ResponseWriter, r *http.Request) {

		request, err := getParams(r.URL.Query())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		results := getFizzbuzzList(request)
		json_response, err := json.Marshal(results)
		if err != nil {
			http.Error(w, "json-encode-error", http.StatusInternalServerError)
			return
		}

		err = repositories.UpdateRequest(s, request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json_response)
	}
}
