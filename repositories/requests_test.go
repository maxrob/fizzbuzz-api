package repositories

import (
	"context"
	"testing"

	"github.com/fizzbuzz-api/models"
	"github.com/fizzbuzz-api/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestInsertWithUpdateRequest(t *testing.T) {
	storer := store.NewDummy()
	collection := storer.GetCollection("requests")

	if err := collection.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}

	request := &models.Request{
		FirstInt:     3,
		SecondInt:    7,
		Limit:        50,
		FirstString:  "first",
		SecondString: "second",
	}

	err := UpdateRequest(storer, request)
	if err != nil {
		t.Fatal(err)
	}

	result := &models.Request{}
	opts := options.FindOne()
	filter := bson.D{
		{"first_int", request.FirstInt},
		{"second_int", request.SecondInt},
		{"limit", request.Limit},
		{"first_string", request.FirstString},
		{"second_string", request.SecondString},
	}
	err = collection.FindOne(context.TODO(), filter, opts).Decode(result)
	if err != nil {
		t.Fatal(err)
	}

	if result.FirstInt != request.FirstInt {
		t.Errorf("Unexpected first_int: got %v want %v", result.FirstInt, request.FirstInt)
	}

	if result.SecondInt != request.SecondInt {
		t.Errorf("Unexpected second_int: got %v want %v", result.SecondInt, request.SecondInt)
	}

	if result.Limit != request.Limit {
		t.Errorf("Unexpected limit: got %v want %v", result.Limit, request.Limit)
	}

	if result.FirstString != request.FirstString {
		t.Errorf("Unexpected first_string: got %v want %v", result.FirstString, request.FirstString)
	}

	if result.SecondString != request.SecondString {
		t.Errorf("Unexpected second_string: got %v want %v", result.SecondString, request.SecondString)
	}

	if result.Iteration != 1 {
		t.Errorf("Unexpected second_string: got %v want %v", result.SecondString, 1)
	}

	err = UpdateRequest(storer, request)
	if err != nil {
		t.Fatal(err)
	}

	err = collection.FindOne(context.TODO(), filter, opts).Decode(result)
	if err != nil {
		t.Fatal(err)
	}

	if result.Iteration != 2 {
		t.Errorf("Unexpected second_string: got %v want %v", result.SecondString, 2)
	}
}

func TestUpdateWithUpdateRequest(t *testing.T) {
	storer := store.NewDummy()
	collection := storer.GetCollection("requests")

	if err := collection.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}

	request := &models.Request{
		FirstInt:     3,
		SecondInt:    7,
		Limit:        50,
		FirstString:  "first",
		SecondString: "second",
	}

	err := UpdateRequest(storer, request)
	if err != nil {
		t.Fatal(err)
	}

	err = UpdateRequest(storer, request)
	if err != nil {
		t.Fatal(err)
	}

	result := &models.Request{}
	opts := options.FindOne()
	filter := bson.D{
		{"first_int", request.FirstInt},
		{"second_int", request.SecondInt},
		{"limit", request.Limit},
		{"first_string", request.FirstString},
		{"second_string", request.SecondString},
	}
	err = collection.FindOne(context.TODO(), filter, opts).Decode(result)
	if err != nil {
		t.Fatal(err)
	}

	if result.FirstInt != request.FirstInt {
		t.Errorf("Unexpected first_int: got %v want %v", result.FirstInt, request.FirstInt)
	}

	if result.SecondInt != request.SecondInt {
		t.Errorf("Unexpected second_int: got %v want %v", result.SecondInt, request.SecondInt)
	}

	if result.Limit != request.Limit {
		t.Errorf("Unexpected limit: got %v want %v", result.Limit, request.Limit)
	}

	if result.FirstString != request.FirstString {
		t.Errorf("Unexpected first_string: got %v want %v", result.FirstString, request.FirstString)
	}

	if result.SecondString != request.SecondString {
		t.Errorf("Unexpected second_string: got %v want %v", result.SecondString, request.SecondString)
	}

	if result.Iteration != 2 {
		t.Errorf("Unexpected second_string: got %v want %v", result.SecondString, 2)
	}
}

func TestOneRequestWithGetMostPopularRequest(t *testing.T) {
	storer := store.NewDummy()
	collection := storer.GetCollection("requests")

	if err := collection.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}

	request := &models.Request{
		FirstInt:     3,
		SecondInt:    7,
		Limit:        50,
		FirstString:  "first",
		SecondString: "second",
	}

	err := UpdateRequest(storer, request)
	if err != nil {
		t.Fatal(err)
	}

	expectedResult := &models.Request{}
	opts := options.FindOne()
	filter := bson.D{
		{"first_int", request.FirstInt},
		{"second_int", request.SecondInt},
		{"limit", request.Limit},
		{"first_string", request.FirstString},
		{"second_string", request.SecondString},
	}
	err = collection.FindOne(context.TODO(), filter, opts).Decode(expectedResult)
	if err != nil {
		t.Fatal(err)
	}

	result, err := GetMostPopularRequest(storer)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != expectedResult.ID {
		t.Errorf("Unexpected id: got %v want %v", result.FirstInt, expectedResult.FirstInt)
	}

	if result.FirstInt != expectedResult.FirstInt {
		t.Errorf("Unexpected first_int: got %v want %v", result.FirstInt, expectedResult.FirstInt)
	}

	if result.SecondInt != expectedResult.SecondInt {
		t.Errorf("Unexpected second_int: got %v want %v", result.SecondInt, expectedResult.SecondInt)
	}

	if result.Limit != expectedResult.Limit {
		t.Errorf("Unexpected limit: got %v want %v", result.Limit, expectedResult.Limit)
	}

	if result.FirstString != expectedResult.FirstString {
		t.Errorf("Unexpected first_string: got %v want %v", result.FirstString, expectedResult.FirstString)
	}

	if result.SecondString != expectedResult.SecondString {
		t.Errorf("Unexpected second_string: got %v want %v", result.SecondString, expectedResult.SecondString)
	}

	if result.Iteration != expectedResult.Iteration {
		t.Errorf("Unexpected second_string: got %v want %v", result.SecondString, expectedResult.Iteration)
	}
}

func TestMultipleRequestWithGetMostPopularRequest(t *testing.T) {
	storer := store.NewDummy()
	collection := storer.GetCollection("requests")

	if err := collection.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}

	firstRequest := &models.Request{
		FirstInt:     3,
		SecondInt:    7,
		Limit:        50,
		FirstString:  "first",
		SecondString: "request",
	}

	err := UpdateRequest(storer, firstRequest)
	if err != nil {
		t.Fatal(err)
	}

	err = UpdateRequest(storer, firstRequest)
	if err != nil {
		t.Fatal(err)
	}

	secondRequest := &models.Request{
		FirstInt:     4,
		SecondInt:    7,
		Limit:        50,
		FirstString:  "second",
		SecondString: "request",
	}

	err = UpdateRequest(storer, secondRequest)
	if err != nil {
		t.Fatal(err)
	}

	expectedResult := &models.Request{}
	opts := options.FindOne()
	filter := bson.D{
		{"first_int", firstRequest.FirstInt},
		{"second_int", firstRequest.SecondInt},
		{"limit", firstRequest.Limit},
		{"first_string", firstRequest.FirstString},
		{"second_string", firstRequest.SecondString},
	}
	err = collection.FindOne(context.TODO(), filter, opts).Decode(expectedResult)
	if err != nil {
		t.Fatal(err)
	}

	result, err := GetMostPopularRequest(storer)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != expectedResult.ID {
		t.Errorf("Unexpected id: got %v want %v", result.FirstInt, expectedResult.FirstInt)
	}

	if result.FirstInt != expectedResult.FirstInt {
		t.Errorf("Unexpected first_int: got %v want %v", result.FirstInt, expectedResult.FirstInt)
	}

	if result.SecondInt != expectedResult.SecondInt {
		t.Errorf("Unexpected second_int: got %v want %v", result.SecondInt, expectedResult.SecondInt)
	}

	if result.Limit != expectedResult.Limit {
		t.Errorf("Unexpected limit: got %v want %v", result.Limit, expectedResult.Limit)
	}

	if result.FirstString != expectedResult.FirstString {
		t.Errorf("Unexpected first_string: got %v want %v", result.FirstString, expectedResult.FirstString)
	}

	if result.SecondString != expectedResult.SecondString {
		t.Errorf("Unexpected second_string: got %v want %v", result.SecondString, expectedResult.SecondString)
	}

	if result.Iteration != expectedResult.Iteration {
		t.Errorf("Unexpected second_string: got %v want %v", result.SecondString, expectedResult.Iteration)
	}
}
