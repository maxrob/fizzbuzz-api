package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fizzbuzz-api/store"
)

func TestGetFizzbuzzHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/fizzbuzzs", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("int1", "3")
	q.Add("int2", "5")
	q.Add("limit", "15")
	q.Add("str1", "fizz")
	q.Add("str2", "buzz")
	req.URL.RawQuery = q.Encode()

	res := httptest.NewRecorder()

	storer := store.NewDummy()
	if err = storer.GetCollection("requests").Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(GetFizzbuzzHandler(storer))
	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedResult := []string{
		"1",
		"2",
		"fizz",
		"4",
		"buzz",
		"fizz",
		"7",
		"8",
		"fizz",
		"buzz",
		"11",
		"fizz",
		"13",
		"14",
		"fizzbuzz",
	}
	expected, err := json.Marshal(expectedResult)
	if err != nil {
		t.Fatal(err)
	}
	if res.Body.String() != string(expected) {
		t.Errorf("Unexpected body: got %v want %v", res.Body.String(), expected)
	}
}
