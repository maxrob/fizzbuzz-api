package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fizzbuzz-api/models"
	"github.com/fizzbuzz-api/repositories"
	"github.com/fizzbuzz-api/store"
)

func TestGetMostPopularRequestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/fizzbuzzs/most_popular_request", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	storer := store.NewDummy()
	if err = storer.GetCollection("requests").Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}

	testRequestModel := &models.Request{
		FirstInt:     3,
		SecondInt:    7,
		Limit:        50,
		FirstString:  "first",
		SecondString: "second",
	}

	err = repositories.UpdateRequest(storer, testRequestModel)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(GetMostPopularRequestHandler(storer))
	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}

	result := &models.Request{}
	err = json.NewDecoder(res.Body).Decode(result)
	if err != nil {
		t.Fatal(err)
	}

	if testRequestModel.FirstInt != result.FirstInt {
		t.Errorf("Unexpected first_int: got %v want %v", result.FirstInt, testRequestModel.FirstInt)
	}

	if testRequestModel.SecondInt != result.SecondInt {
		t.Errorf("Unexpected second_int: got %v want %v", result.SecondInt, testRequestModel.SecondInt)
	}

	if testRequestModel.Limit != result.Limit {
		t.Errorf("Unexpected limit: got %v want %v", result.Limit, testRequestModel.Limit)
	}

	if testRequestModel.FirstString != result.FirstString {
		t.Errorf("Unexpected first_string: got %v want %v", result.FirstString, testRequestModel.FirstString)
	}

	if testRequestModel.SecondString != result.SecondString {
		t.Errorf("Unexpected second_string: got %v want %v", result.SecondString, testRequestModel.SecondString)
	}
}
